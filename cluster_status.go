package goproxmoxapi

// structure representing Proxmox Cluster Node Entry
type ClusterNodeEntry struct {
  Id     string // "node/pve",
  Ip     string // "10.255.0.5",
  Level  string // "",
  Local  int    // 1,
  Name   string // "pve",
  Nodeid int    // 0,
  Online int    // 1,
  Type   string // "node"
}

// Get Proxmox Cluster Status
func GetClusterStatus(c *Client) ([]ClusterNodeEntry, error) {
  status := []ClusterNodeEntry{}

  _, rbody, err := c.NewRequest("GET", "/api2/json/cluster/status", nil )
  if err != nil {
	  return status, err
  } else {
    err = dataUnmarshal( rbody, &status )

    // Any Error parsing json ?
		return status, err
  }
}
