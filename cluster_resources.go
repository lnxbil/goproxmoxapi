package goproxmoxapi

// structure representing Proxmox Task Log Entry
type ClusterResource struct {
  Cpu        float64  //  0.0888742436099737
  Disk       float64  //  718333280
  Diskread   int      //  0
  Diskwrite  int      //  0
  Id         string   //  "storage/pve/local-lvm"
  Level      string   //  ""
  Maxcpu     int      //  1
  Maxdisk    float64  //  8589934592
  Maxmem     float64  //  536870912
  Mem        float64  //  885272576
  Name       string   //  "100"
  Netin      int      //  0
  Netout     int      //  0
  Node       string   //  "pve"
  Pool       string   //  "linux_vms"
  Status     string   //  "stopped"
  Storage    string   //  "local-lvm"
  Template   int      //  0
  Type       string   //  "storage"
  Uptime     int      //  260716
  Vmid       int      //  100
}

// Get Proxmox recent tasks (cluster wide).
func GetClusterResources(c *Client) ([]ClusterResource, error) {
  rc := []ClusterResource{}

  _, rbody, err := c.NewRequest("GET", "/api2/json/cluster/resources", nil )
  if err != nil {
	  return rc, err
  } else {
    err = dataUnmarshal( rbody, &rc )
		return rc, err
  }
}
