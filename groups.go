package goproxmoxapi

import (
  "net/url"
)

// structure representing Proxmox Group
type Group struct {
	GroupId   string   // Required
  Comment   string   // Optional
}

// Create new Proxmox Group
func (g Group) CreateGroup(c *Client) error {
  // check definition of required fields
  if g.GroupId == "" {
    return &errorString{ "groupid is required to create Proxmox group", }
  }

  // POST parameters
  pbody := structToMap(&g, []string{}, []string{})

	_, _, err := c.NewRequest("POST", "/api2/json/access/groups", pbody )
  if err != nil {
    return err
  }
  return nil
}

// Delete Proxmox group
func (g Group) DeleteGroup(c *Client) error {
  // check definition of required fields
  if g.GroupId == "" {
    return &errorString{ "GroupId is required to delete Proxmox group", }
  }

  _, _, err := c.NewRequest("DELETE", "/api2/json/access/groups/" + g.GroupId, nil )
  if err != nil {
    return err
  }
  return err
}

// Update Proxmox group
func (g Group) UpdateGroup(c *Client) error {
  // check definition of required fields
  if g.GroupId == "" {
    return &errorString{ "GroupId is required to update Proxmox group", }
  }

  // POST parameters
	pbody := url.Values{}
  pbody.Add( "comment", g.Comment )

  _, _, err := c.NewRequest("PUT", "/api2/json/access/groups/" + g.GroupId, pbody )
  return err
}

// Get Proxmox group
func (g Group) GetGroup(c *Client) (Group, error) {
  rg := Group{}
  // check definition of required fields
  if g.GroupId == "" {
    return rg, &errorString{ "groupid is required to get Proxmox group", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/access/groups/" + g.GroupId, nil )
  if err != nil {
	  return rg, err
  } else {
    err = dataUnmarshal( rbody, &rg )

    // Any Error parsing json ?
		return rg, err
  }
}

// Gets all defined groups
func GetAllGroups(c *Client) ([]Group, error) {
	groups := make([]Group, 0)

  // GET parameters
	_, rbody, err := c.NewRequest("GET", "/api2/json/access/groups", nil )
  if err != nil {
	  return groups, err
  } else {
    err = dataUnmarshal( rbody, &groups )

    // Any Error parsing json ?
		return groups, err
  }
}
