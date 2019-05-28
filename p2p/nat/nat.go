// Package nat implements NAT handling facilities
package nat

import (
	"context"
	"errors"
	"log"
	"net"
	"time"
)

const (
	mapTimeout        = 20 * time.Minute
	mapUpdateInterval = 15 * time.Minute
)

var errNoExternalAddress = errors.New("no external address")
var errNoInternalAddress = errors.New("no internalMsg address")
var errNoNATFound = errors.New("no NAT found")

// NAT represents an interface for upnp and natpmp
type NAT interface {
	// Type returns the kind of NAT port mapping service that is used
	getType() string

	// GetDeviceAddress returns the internalMsg address of the gateway device.
	getDeviceAddress() (addr net.IP, err error)

	// GetExternalAddress returns the external address of the gateway device.
	GetExternalAddress() (addr net.IP, err error)

	// GetInternalAddress returns the address of the local host.
	getInternalAddress() (addr net.IP, err error)

	// AddPortMapping maps a port on the local host to an external port.
	addPortMapping(protocol string, externalPort, internalPort int, description string, timeout time.Duration) error

	// DeletePortMapping removes a port mapping.
	deletePortMapping(protocol string, internalPort int) (err error)
}

// SetMapping adds a port mapping on m and keeps it alive until c is closed.
func SetMapping(ctx context.Context, nat NAT, protocol string, extport, intport int, name string) error {
	if err := nat.addPortMapping(protocol, extport, intport, name, mapTimeout); err != nil {
		return err
	}

	go func() {

		refresh := time.NewTimer(mapUpdateInterval)
		defer func() {
			refresh.Stop()
			nat.deletePortMapping(protocol, intport)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case <-refresh.C:
				if err := nat.addPortMapping(protocol, extport, intport, name, mapTimeout); err != nil {
					log.Fatal(err)
				}
				refresh.Reset(mapUpdateInterval)
			}
		}

	}()

	return nil
}

// DiscoverGateway attempts to find a gateway device.
func DiscoverGateway() (NAT, error) {
	select {
	case nat := <-discoverUPNPIG1():
		return nat, nil
	case nat := <-discoverUPNPIG2():
		return nat, nil
	case nat := <-discoverNATPMP():
		return nat, nil
	case <-time.After(10 * time.Second):
		return nil, errNoNATFound
	}
}
