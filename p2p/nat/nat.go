// Package nat implements NAT handling facilities
package nat

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/jackpal/gateway"
)

const (
	mapTimeout        = 20 * time.Minute
	mapUpdateInterval = 15 * time.Minute
)

var ErrNoExternalAddress = errors.New("no external address")
var ErrNoInternalAddress = errors.New("no internalMsg address")
var ErrNoNATFound = errors.New("no NAT found")

// protocol is either "udp" or "tcp"
type NAT interface {
	// Type returns the kind of NAT port mapping service that is used
	getType() string

	// GetDeviceAddress returns the internalMsg address of the gateway device.
	getDeviceAddress() (addr net.IP, err error)

	// GetExternalAddress returns the external address of the gateway device.
	getExternalAddress() (addr net.IP, err error)

	// GetInternalAddress returns the address of the local host.
	getInternalAddress() (addr net.IP, err error)

	// AddPortMapping maps a port on the local host to an external port.
	addPortMapping(protocol string, externalPort, internalPort int, description string, timeout time.Duration) error

	// DeletePortMapping removes a port mapping.
	deletePortMapping(protocol string, internalPort int) (err error)
}

type natController struct {
	natDevice NAT
	c chan struct{}
}

func SetMapping(protocol string, extport, intport int, name string)  (*natController, error) {
	nat, err := DiscoverGateway()
	if err != nil {
		return nil, err
	}

	if err := nat.addPortMapping(protocol, extport, intport, name, mapTimeout); err != nil {
		return nil, err
	}

	c := make(chan struct{})

	go func() {

		refresh := time.NewTimer(mapUpdateInterval)
		defer func() {
			refresh.Stop()
			nat.deletePortMapping(protocol, intport)
		}()

		for {
			select {
			case _, ok := <-c:
				if !ok {
					return
				}
			case <-refresh.C:
				if err := nat.addPortMapping(protocol, extport, intport, name, mapTimeout); err != nil {
					log.Fatal(err)
				}
				refresh.Reset(mapUpdateInterval)
			}
		}

	}()

	return &natController{
		natDevice:nat,
		c:c,
	}, nil
}

// DiscoverGateway attempts to find a gateway device.
func DiscoverGateway() (NAT, error) {
	select {
	case nat := <-discoverUPNP_IG1():
		return nat, nil
	case nat := <-discoverUPNP_IG2():
		return nat, nil
	case nat := <-discoverNATPMP():
		return nat, nil
	case <-time.After(10 * time.Second):
		return nil, ErrNoNATFound
	}
}

func (natController *natController) CloseMapping() {
	close(natController.c)
	time.Sleep(2*time.Second)
}

func (natController *natController) GetDeviceAddress() (addr net.IP, err error) {
	return natController.natDevice.getDeviceAddress()
}

func (natController *natController) GetExternalAddress() (addr net.IP, err error) {
	return natController.natDevice.getExternalAddress()
}

func (natController *natController) GetInternalAddress() (addr net.IP, err error) {
	return natController.natDevice.getInternalAddress()
}

func (natController *natController) GetType() string {
	return natController.natDevice.getType()
}

func IsPrivateIp() (bool, error) {
	ip, err := gateway.DiscoverGateway()
	if err != nil {
		return false, err
	}

	return netutil.IsLAN(ip), nil
}

func RandomPort() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(math.MaxUint16-10000) + 10000
}

