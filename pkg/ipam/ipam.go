package ipam

import (
	"errors"
	"fmt"
	"net/netip"

	"go4.org/netipx"
	"k8s.io/klog"
)

// Manager - handles the addresses for each namespace/vip
var Manager []ipManager

// ipManager defines the mapping to a namespace and address pool
type ipManager struct {
	// Identifies the manager
	namespace string

	// The network configuration
	cidr    string
	ipRange string

	// todo - This confuses me ...
	poolIPSet *netipx.IPSet
}

// FindAvailableHostFromRange - will look through the cidr and the address Manager and find a free address (if possible)
func FindAvailableHostFromRange(namespace, ipRange string, inUseIPSet *netipx.IPSet) (string, error) {

	// Look through namespaces and update one if it exists
	for x := range Manager {
		if Manager[x].namespace == namespace {
			// Check that the address range is the same
			if Manager[x].ipRange != ipRange {
				klog.Infof("Updating IP address range from [%s] to [%s]", Manager[x].ipRange, ipRange)

				// If not rebuild the available hosts
				poolIPSet, err := buildAddressesFromRange(ipRange)
				if err != nil {
					return "", err
				}
				Manager[x].poolIPSet = poolIPSet
				Manager[x].ipRange = ipRange
			}

			// TODO - currently we search (incrementally) through the list of hosts
			addr, err := FindFreeAddress(Manager[x].poolIPSet, inUseIPSet)
			if err != nil {
				return "", fmt.Errorf("no addresses available in [%s] range [%s]", namespace, ipRange)
			}
			return addr.String(), nil
		}
	}
	poolIPSet, err := buildAddressesFromRange(ipRange)
	if err != nil {
		return "", err
	}

	// If it doesn't exist then it will need adding
	newManager := ipManager{
		namespace: namespace,
		poolIPSet: poolIPSet,
		ipRange:   ipRange,
	}

	Manager = append(Manager, newManager)

	// TODO - currently we search (incrementally) through the list of hosts
	addr, err := FindFreeAddress(poolIPSet, inUseIPSet)
	if err != nil {
		return "", fmt.Errorf("no addresses available in [%s] range [%s]", namespace, ipRange)
	}
	return addr.String(), nil
}

// FindAvailableHostFromCidr - will look through the cidr and the address Manager and find a free address (if possible)
func FindAvailableHostFromCidr(namespace, cidr string, inUseIPSet *netipx.IPSet) (string, error) {

	// Look through namespaces and update one if it exists
	for x := range Manager {
		if Manager[x].namespace == namespace {
			// Check that the address range is the same
			if Manager[x].cidr != cidr {
				// If not rebuild the available hosts
				poolIPSet, err := buildHostsFromCidr(cidr)
				if err != nil {
					return "", err
				}
				Manager[x].poolIPSet = poolIPSet
				Manager[x].cidr = cidr

			}
			// TODO - currently we search (incrementally) through the list of hosts
			addr, err := FindFreeAddress(Manager[x].poolIPSet, inUseIPSet)
			if err != nil {
				return "", fmt.Errorf("no addresses available in [%s] cidr [%s]", namespace, cidr)
			}
			return addr.String(), nil

		}
	}
	poolIPSet, err := buildHostsFromCidr(cidr)
	if err != nil {
		return "", err
	}
	// If it doesn't exist then it will need adding
	newManager := ipManager{
		namespace: namespace,
		poolIPSet: poolIPSet,
		cidr:      cidr,
	}
	Manager = append(Manager, newManager)

	// TODO - currently we search (incrementally) through the list of hosts
	addr, err := FindFreeAddress(poolIPSet, inUseIPSet)
	if err != nil {
		return "", fmt.Errorf("no addresses available in [%s] cidr [%s]", namespace, cidr)
	}
	return addr.String(), nil

}

// // RenewAddress - removes the mark on an address
// func RenewAddress(namespace, address string) {
// 	for x := range Manager {
// 		if Manager[x].namespace == namespace {
// 			// Make sure we update the address manager to mark this address in use.
// 			Manager[x].addressManager[address] = true
// 		}
// 	}
// }

// // ReleaseAddress - removes the mark on an address
// func ReleaseAddress(namespace, address string) error {
// 	for x := range Manager {
// 		if Manager[x].namespace == namespace {
// 			Manager[x].addressManager[address] = false
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("unable to release address [%s] in namespace [%s]", address, namespace)
// }

// FindFreeAddress returns the next free IP Address in a range based on a set of existing addresses.
// It will skip assumed gateway ip or broadcast ip for IPv4 address
func FindFreeAddress(poolIPSet *netipx.IPSet, inUseIPSet *netipx.IPSet) (netip.Addr, error) {
	for _, iprange := range poolIPSet.Ranges() {
		ip := iprange.From()
		for {
			if !inUseIPSet.Contains(ip) && (!ip.Is4() || !isNetworkIDOrBroadcastIP(ip.As4())) {
				return ip, nil
			}
			if ip == iprange.To() {
				break
			}
			ip = ip.Next()
		}
	}
	return netip.Addr{}, errors.New("no address available")
}

func isNetworkIDOrBroadcastIP(ip [4]byte) bool {
	return ip[3] == 0 || ip[3] == 255
}
