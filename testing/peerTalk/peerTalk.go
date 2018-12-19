package main

import (
	"bytes"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
	"github.com/DOSNetwork/core/testing/peerTalk/msg"

	"github.com/bshuster-repo/logrus-logstash-hook"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

func main() {
	id := uuid.New()
	fmt.Println(id[:])
	log := logrus.New()
	tcpAddr, err := net.ResolveTCPAddr("tcp", "163.172.36.173:9500")
	if err != nil {
		log.Error(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error(err)
	}

	hook, err := logrustash.NewHookWithFieldsAndConn(conn, "匹凸匹test", logrus.Fields{
		"startingTimestamp": time.Now(),
		"id":                new(big.Int).SetBytes(id[:]).String(),
	})
	if err != nil {
		log.Error(err)
	}

	log.AddHook(hook)
	logger := log.WithFields(logrus.Fields{})
	p, peerEvent, err := p2p.CreateP2PNetwork(id[:], 44460, logger)
	if err != nil {
		logger.Fatal(err)
	}
	if err := p.Listen(); err != nil {
		logger.Error(err)
	}

	if os.Getenv("ROLE") == "bootstrap" {
		logger.Data["role"] = "bootstrap"
		logger.Info()

		peerNumStr := os.Getenv("PEERNUM")
		peerNum, err := strconv.Atoi(peerNumStr)
		if err != nil {
			logger.Fatal(err)
		}

		groupSizeStr := os.Getenv("GROUPSIZE")
		groupSize, err := strconv.Atoi(groupSizeStr)
		if err != nil {
			logger.Fatal(err)
		}

		var ids [][]byte
		idMap := make(map[string]string)
		lastAddress := p.GetIP()

		for event := range peerEvent {
			switch content := event.Msg.Message.(type) {
			case *peerTalk.Register:
				if len(ids) < peerNum {
					if err = p.SendMessage(content.GetId(), &peerTalk.Bootstrap{Address: lastAddress}); err != nil {
						logger.WithFields(logrus.Fields{
							"event": "sendMessage",
							"toId":  new(big.Int).SetBytes(content.GetId()).String(),
						}).Warn(err)
					} else {
						logger.WithFields(logrus.Fields{
							"event": "sendMessage",
							"toId":  new(big.Int).SetBytes(content.GetId()).String(),
						}).Info()
					}
					lastAddress = content.GetAddress()
					ids = append(ids, content.GetId())
					idMap[new(big.Int).SetBytes(content.GetId()).String()] = ""
					logger.WithFields(logrus.Fields{
						"event":       "receiveId",
						"id":          new(big.Int).SetBytes(content.GetId()).String(),
						"ip":          content.GetAddress(),
						"currentSize": len(ids),
						"targetSize":  peerNum,
					}).Info()
					if len(ids) == peerNum {
						logger.WithField("event", "receiveAllHello").Info()
						go func() {
							for _, id := range ids {
								fmt.Println("grouping cmd", new(big.Int).SetBytes(id).String())
								if err = p.SendMessage(id, &peerTalk.Grouping{Size: uint32(groupSize), Ids: ids}); err != nil {
									logger.WithFields(logrus.Fields{
										"event": "sendMessage",
										"toId":  new(big.Int).SetBytes(id).String(),
									}).Warn(err)
								} else {
									logger.WithFields(logrus.Fields{
										"event": "sendMessage",
										"toId":  new(big.Int).SetBytes(id).String(),
									}).Info()
								}
							}
						}()
					}
				}
			case *peerTalk.Done:
				delete(idMap, new(big.Int).SetBytes(content.GetId()).String())
				logger.WithFields(logrus.Fields{
					"event":       "receiveDone",
					"id":          new(big.Int).SetBytes(content.GetId()).String(),
					"currentSize": len(idMap),
					"targetSize":  peerNum,
				}).Info()
				if len(idMap) == 0 {
					logger.WithField("event", "allDone").Info()
					//os.Exit(-100)
				}
			}
		}
	} else {
		logger.Data["role"] = "node"

		suite := suites.MustFind("bn256")
		peerEventForDKG := make(chan p2p.P2PMessage, 1)
		defer close(peerEventForDKG)
		p2pDkg, err := dkg.CreateP2PDkg(p, suite, peerEventForDKG, 3, logger)
		if err != nil {
			logger.Fatal(err)
		}
		go p2pDkg.EventLoop()
		dkgEvent := make(chan string, 1)
		p2pDkg.SubscribeEvent(dkgEvent)
		defer close(dkgEvent)

		logger.Info(os.Getenv("BOOTSTRAPIP"))
		bootstrapId, err := p.ConnectTo(os.Getenv("BOOTSTRAPIP"))
		if err != nil {
			logger.Fatal(err)
		}

		if err = p.SendMessage(bootstrapId, &peerTalk.Register{Id: p.GetID(), Address: p.GetIP()}); err != nil {
			logger.WithFields(logrus.Fields{
				"event": "sendMessage",
				"toId":  new(big.Int).SetBytes(bootstrapId).String(),
			}).Warn(err)
		} else {
			logger.WithFields(logrus.Fields{
				"event": "sendMessage",
				"toId":  new(big.Int).SetBytes(bootstrapId).String(),
			}).Info()
		}

		for event := range peerEvent {
			switch content := event.Msg.Message.(type) {
			case *vss.PublicKey:
				peerEventForDKG <- event
			case *dkg.Deal:
				peerEventForDKG <- event
			case *dkg.Responses:
				peerEventForDKG <- event
			case *peerTalk.Bootstrap:
				if err = p.Join(content.GetAddress()); err != nil {
					logger.WithFields(logrus.Fields{
						"event":     "join",
						"bootstrap": content.GetAddress(),
					}).Warn(err)
				} else {
					logger.WithFields(logrus.Fields{
						"event":     "join",
						"bootstrap": content.GetAddress(),
					}).Info()
				}
			case *peerTalk.Grouping:
				go func() {
					result := <-dkgEvent
					if result == "certified" {
						if err = p.SendMessage(bootstrapId, &peerTalk.Done{Id: p.GetID()}); err != nil {
							logger.WithField("event", "allDone").Error(err)
						} else {
							logger.WithField("event", "allDone").Info()
						}
						//os.Exit(-100)
					}
				}()

				logger.WithField("event", "receiveGrouping").Info()

				ids := content.GetIds()
				groupSize := content.GetSize()
				var group [][]byte
				for idx, id := range ids {
					if bytes.Compare(p.GetID(), id) == 0 {
						start := uint32(idx) / groupSize * groupSize
						fmt.Println("????????????", start, start+groupSize)
						group = ids[start : start+groupSize]
						break
					}
				}
				for _, id := range group {
					fmt.Println("!!!!!!!!!!!", new(big.Int).SetBytes(id).String())
				}

				p2pDkg.SetNbParticipants(int(groupSize))
				p2pDkg.SetGroupMembers(group)
				p2pDkg.RunDKG()
			}
		}
	}
}
