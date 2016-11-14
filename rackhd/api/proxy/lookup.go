package proxy

import (
	"github.com/RackHD/neighborhood-manager/rackhd/models"
)

// GetAddresses fetches services for passed ServiceName.
// It then creates a map with the (ip:port)'s retrieved from the GetService call
func getAllAddresses() (map[string]struct{}, error) {
	entries, err := models.GetAllRhd()
	if err != nil {
		return nil, err
	}
	addresses := make(map[string]struct{})
	for _, entry := range entries {
		http := entry.HTTPConf
		addr := http.URL
		addresses[addr.Host] = struct{}{}
	}
	return addresses, nil
}

func getAddrForNode(uuid string) (string, error) {
	entry, err := models.GetRhdIDByNodeID(uuid)
	if err != nil {
		return "", err
	}
	return entry, nil
}
