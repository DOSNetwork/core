package main

import (
	"fmt"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/testing/peerTalk/msg"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {
	log := logrus.New()
	id := uuid.New()
	fmt.Println(id[:])
	p, peerEvent, err := p2p.CreateP2PNetwork(id[:], 44460, log)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.Listen(); err != nil {
		log.Error(err)
	}

	if os.Getenv("ROLE") == "bootstrap" {
		logger := log.WithField("role", "bootstrap")
		logger.Info()
		peerNumStr := os.Getenv("PEERNUM")
		peerNum, err := strconv.Atoi(peerNumStr)
		if err != nil {
			logger.Fatal(err)
		}
		var addresses []string
		addressMap := make(map[string]string)
		for event := range peerEvent {
			switch content := event.Msg.Message.(type) {
			case *peerTalk.Address:
				if len(addresses) < peerNum {
					addresses = append(addresses, content.GetAddress())
					addressMap[content.GetAddress()] = ""
					logger.WithFields(logrus.Fields{
						"event":       "receiveIp",
						"currentSize": len(addresses),
						"targetSize":  peerNum,
					}).Info(content.GetAddress())
					if len(addresses) == peerNum {
						logger.WithField("event", "receiveAllHello").Info()
						go func() {
							for _, address := range addresses {
								fmt.Println(address)
								id, err := p.NewPeer(address)
								if err != nil {
									logger.Warn(err)
								}
								fmt.Println(id)

								if err = p.SendMessage(id, &peerTalk.Addresses{Address: addresses}); err != nil {
									logger.Warn(err)
								}
							}
						}()
					}
				}
			case *peerTalk.Done:
				delete(addressMap, content.GetAddress())
				logger.WithFields(logrus.Fields{
					"event":       "receiveDone",
					"currentSize": len(addressMap),
					"targetSize":  peerNum,
				}).Info(content.GetAddress())
				if len(addressMap) == 0 {
					logger.WithField("event", "allDone").Info()
					os.Exit(0)
				}
			}
		}
	} else {
		logger := log.WithField("role", "node")
		logger.Info()
		bootstrapId, err := p.NewPeer(os.Getenv("BOOTSTRAPIP"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(os.Getenv("BOOTSTRAPIP"))

		if err != nil {
			logger.Warn(err)
		}
		if err = p.SendMessage(bootstrapId, &peerTalk.Address{Address: p.GetIP()}); err != nil {
			logger.Warn(err)
		}

		for event := range peerEvent {
			switch content := event.Msg.Message.(type) {
			case *peerTalk.Addresses:
				logger.WithField("event", "receiveList").Info(len(content.GetAddress()))
				if err = p.SendMessage(bootstrapId, &peerTalk.Done{Address: p.GetIP()}); err != nil {
					logger.Warn(err)
				} else {
					logger.WithField("event", "allDone").Info()
					os.Exit(0)
				}
			}
		}
	}
}
