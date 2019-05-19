package nat

import (
	"net"
	"time"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/internetgateway1"
	"github.com/huin/goupnp/dcps/internetgateway2"
)

var (
	_ NAT = (*uPnPNAT)(nil)
)

func discoverUPNPIG1() <-chan NAT {
	res := make(chan NAT, 1)
	go func() {

		// find devices
		devs, err := goupnp.DiscoverDevices(internetgateway1.URN_WANConnectionDevice_1)
		if err != nil {
			return
		}

		for _, dev := range devs {
			if dev.Root == nil {
				continue
			}

			dev.Root.Device.VisitServices(func(srv *goupnp.Service) {
				switch srv.ServiceType {
				case internetgateway1.URN_WANIPConnection_1:
					client := &internetgateway1.WANIPConnection1{ServiceClient: goupnp.ServiceClient{
						SOAPClient: srv.NewSOAPClient(),
						RootDevice: dev.Root,
						Service:    srv,
					}}
					_, isNat, err := client.GetNATRSIPStatus()
					if err == nil && isNat {
						res <- &uPnPNAT{client, make(map[int]int), "UPNP (IG1-IP1)", dev.Root}
						return
					}

				case internetgateway1.URN_WANPPPConnection_1:
					client := &internetgateway1.WANPPPConnection1{ServiceClient: goupnp.ServiceClient{
						SOAPClient: srv.NewSOAPClient(),
						RootDevice: dev.Root,
						Service:    srv,
					}}
					_, isNat, err := client.GetNATRSIPStatus()
					if err == nil && isNat {
						res <- &uPnPNAT{client, make(map[int]int), "UPNP (IG1-PPP1)", dev.Root}
						return
					}

				}
			})
		}

	}()
	return res
}

func discoverUPNPIG2() <-chan NAT {
	res := make(chan NAT, 1)
	go func() {

		// find devices
		devs, err := goupnp.DiscoverDevices(internetgateway2.URN_WANConnectionDevice_2)
		if err != nil {
			return
		}

		for _, dev := range devs {
			if dev.Root == nil {
				continue
			}

			dev.Root.Device.VisitServices(func(srv *goupnp.Service) {
				switch srv.ServiceType {
				case internetgateway2.URN_WANIPConnection_1:
					client := &internetgateway2.WANIPConnection1{ServiceClient: goupnp.ServiceClient{
						SOAPClient: srv.NewSOAPClient(),
						RootDevice: dev.Root,
						Service:    srv,
					}}
					_, isNat, err := client.GetNATRSIPStatus()
					if err == nil && isNat {
						res <- &uPnPNAT{client, make(map[int]int), "UPNP (IG2-IP1)", dev.Root}
						return
					}

				case internetgateway2.URN_WANIPConnection_2:
					client := &internetgateway2.WANIPConnection2{ServiceClient: goupnp.ServiceClient{
						SOAPClient: srv.NewSOAPClient(),
						RootDevice: dev.Root,
						Service:    srv,
					}}
					_, isNat, err := client.GetNATRSIPStatus()
					if err == nil && isNat {
						res <- &uPnPNAT{client, make(map[int]int), "UPNP (IG2-IP2)", dev.Root}
						return
					}

				case internetgateway2.URN_WANPPPConnection_1:
					client := &internetgateway2.WANPPPConnection1{ServiceClient: goupnp.ServiceClient{
						SOAPClient: srv.NewSOAPClient(),
						RootDevice: dev.Root,
						Service:    srv,
					}}
					_, isNat, err := client.GetNATRSIPStatus()
					if err == nil && isNat {
						res <- &uPnPNAT{client, make(map[int]int), "UPNP (IG2-PPP1)", dev.Root}
						return
					}

				}
			})
		}

	}()
	return res
}

type uPnPNATClient interface {
	GetExternalIPAddress() (string, error)
	AddPortMapping(string, uint16, string, uint16, string, bool, string, uint32) error
	DeletePortMapping(string, uint16, string) error
}

type uPnPNAT struct {
	c          uPnPNATClient
	ports      map[int]int
	typ        string
	rootDevice *goupnp.RootDevice
}

func (u *uPnPNAT) getExternalAddress() (addr net.IP, err error) {
	ipString, err := u.c.GetExternalIPAddress()
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(ipString)
	if ip == nil {
		return nil, ErrNoExternalAddress
	}

	return ip, nil
}

func mapProtocol(s string) string {
	switch s {
	case "udp":
		return "UDP"
	case "tcp":
		return "TCP"
	default:
		panic("invalid protocol: " + s)
	}
}

func (u *uPnPNAT) addPortMapping(protocol string, externalPort, internalPort int, description string, timeout time.Duration) error {
	u.deletePortMapping(protocol, internalPort)

	ip, err := u.getInternalAddress()
	if err != nil {
		return err
	}

	timeoutInSeconds := uint32(timeout / time.Second)

	if externalPort := u.ports[internalPort]; externalPort > 0 {
		err = u.c.AddPortMapping("", uint16(externalPort), mapProtocol(protocol), uint16(internalPort), ip.String(), true, description, timeoutInSeconds)
		if err == nil {
			return nil
		}
	}

	for i := 0; i < 3; i++ {
		err = u.c.AddPortMapping("", uint16(externalPort), mapProtocol(protocol), uint16(internalPort), ip.String(), true, description, timeoutInSeconds)
		if err == nil {
			u.ports[internalPort] = externalPort
			return nil
		}
	}

	return err
}

func (u *uPnPNAT) deletePortMapping(protocol string, internalPort int) error {
	if externalPort := u.ports[internalPort]; externalPort > 0 {
		delete(u.ports, internalPort)
		return u.c.DeletePortMapping("", uint16(externalPort), mapProtocol(protocol))
	}

	return nil
}

func (u *uPnPNAT) getDeviceAddress() (net.IP, error) {
	addr, err := net.ResolveUDPAddr("udp4", u.rootDevice.URLBase.Host)
	if err != nil {
		return nil, err
	}

	return addr.IP, nil
}

func (u *uPnPNAT) getInternalAddress() (net.IP, error) {
	devAddr, err := u.getDeviceAddress()
	if err != nil {
		return nil, err
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			switch x := addr.(type) {
			case *net.IPNet:
				if x.Contains(devAddr) {
					return x.IP, nil
				}
			}
		}
	}

	return nil, ErrNoInternalAddress
}

func (u *uPnPNAT) getType() string { return u.typ }
