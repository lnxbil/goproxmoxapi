package goproxmoxapi

import (
  "net/url"
)

// structure representing Proxmox Pool
type Pool struct {
  PoolId   string   // Required
  Comment  string   // Optional
}

// Create new Proxmox Pool
func (p Pool) CreatePool(c *Client) error {
  // check definition of required fields
  if p.PoolId == "" {
    return &errorString{ "PoolId is required to create Proxmox Pool", }
  }

  // POST parameters
  pbody := structToMap(&p, []string{}, []string{} )

  _, _, err := c.NewRequest("POST", "/api2/json/pools", pbody )
  if err != nil {
    return err
  }
  return nil
}

// Delete Proxmox Pool
func (p Pool) DeletePool(c *Client) error {
  // check definition of required fields
  if p.PoolId == "" {
    return &errorString{ "PoolId is required to delete Proxmox Pool", }
  }

  _, _, err := c.NewRequest("DELETE", "/api2/json/pools/" + p.PoolId, nil )
  if err != nil {
    return err
  }
  return err
}

// Update Proxmox Pool
func (p Pool) UpdatePool(c *Client) error {
  // check definition of required fields
  if p.PoolId == "" {
    return &errorString{ "PoolId is required to update Proxmox Pool", }
  }

  // POST parameters
  pbody := url.Values{}
  pbody.Add( "comment", p.Comment )

  _, _, err := c.NewRequest("PUT", "/api2/json/pools/" + p.PoolId, pbody )
  return err
}

// Get Proxmox pool
func (p Pool) GetPool(c *Client) (Pool, error) {
  pp := Pool{}
  // check definition of required fields
  if p.PoolId == "" {
    return pp, &errorString{ "PoolId is required to get Proxmox Pool", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/pools/" + p.PoolId, nil )
  if err != nil {
    return pp, err
  } else {
    err = dataUnmarshal( rbody, &pp )

    return pp, err
  }
}

// Gets all defined pools
func GetAllPools(c *Client) ([]Pool, error) {
  pools := make([]Pool, 0)

  // GET parameters
  _, rbody, err := c.NewRequest("GET", "/api2/json/pools", nil )
  if err != nil {
    return pools, err
  } else {
    err = dataUnmarshal( rbody, &pools )

    return pools, err
  }
}
