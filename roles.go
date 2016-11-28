package goproxmoxapi

// structure representing Proxmox Role
type Role struct {
	RoleId  string     // Required
  Privs   string     // Optional
}

// Create new Proxmox Role
func (r Role) CreateRole(c *Client) error {
  // check definition of required fields
  if r.RoleId == "" {
    return &errorString{ "RoleId is required to create Proxmox role", }
  }

  // POST parameters
  pbody := structToMap(&r, []string{}, []string{} )

	_, _, err := c.NewRequest("POST", "/api2/json/access/roles", pbody )
  if err != nil {
    return err
  }
  return nil
}

// Delete Proxmox Role
func (r Role) DeleteRole(c *Client) error {
  // check definition of required fields
  if r.RoleId == "" {
    return &errorString{ "RoleId is required to delete Proxmox role", }
  }

  _, _, err := c.NewRequest("DELETE", "/api2/json/access/roles/" + r.RoleId, nil )
  if err != nil {
    return err
  }
  return err
}

// Update Proxmox role
func (r Role) UpdateRole(c *Client) error {
  // check definition of required fields
  if r.RoleId == "" {
    return &errorString{ "RoleId is required to update Proxmox role", }
  }

  // POST parameters
  pbody := structToMap(&r, []string{}, []string{ "roleid" } )

  _, _, err := c.NewRequest("PUT", "/api2/json/access/roles/" + r.RoleId, pbody )
  return err
}

// Get Proxmox role
func (r Role) GetRole(c *Client) (Role, error) {
  rl := Role{}
  // check definition of required fields
  if r.RoleId == "" {
    return rl, &errorString{ "RoleId is required to get Proxmox role", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/access/roles/" + r.RoleId, nil )
  if err != nil {
	  return rl, err
  } else {
    var privs map[string]interface{}
    err = dataUnmarshal( rbody, &privs )

    var privsString string
    for k,_ := range privs {
      privsString = privsString + k + ","
    }

    rl.RoleId = r.RoleId
    rl.Privs = privsString
    // Any Error parsing json ?
		return rl, err
  }
}

// Gets all defined roles
func GetAllRoles(c *Client) ([]Role, error) {
  rr := make([]Role, 0)

  // GET parameters
	_, rbody, err := c.NewRequest("GET", "/api2/json/access/roles", nil )
  if err != nil {
	  return rr, err
  } else {
    err = dataUnmarshal( rbody, &rr )

    // Any Error parsing json ?
		return rr, err
  }
}
