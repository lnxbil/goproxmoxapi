package goproxmoxapi

//import (
//  "strconv"
//  "net/url"
//)

// Check availability of the VMID or get next available VMID from PVE
func GetNextVMID(c *Client) (string, error) {
//func GetNextVMID(c *Client, idtocheck ...int) (string, error) {
//	data := url.Values{}
  vmid := ""
  // if idtocheck is passed and is positive int
//  if len(idtocheck)>0 && idtocheck[0]>0 {
//    data.Add( "vmid", strconv.Itoa(idtocheck[0]) )
//  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/cluster/nextid", nil )
//  _, rbody, err := c.NewRequest("GET", "/api2/json/cluster/nextid", data )
  if err != nil {
	  return vmid, err
  } else {
    err = dataUnmarshal( rbody, &vmid )
		return vmid, err
  }
}
