package goproxmoxapi

// Check availability of the VMID or get next available VMID from PVE
func GetNextVMID(c *Client) (string, error) {
  vmid := ""

  _, rbody, err := c.NewRequest("GET", "/api2/json/cluster/nextid", nil )
  if err != nil {
    return vmid, err
  } else {
    err = dataUnmarshal( rbody, &vmid )
    return vmid, err
  }
}
