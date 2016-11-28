package goproxmoxapi

// structure representing Proxmox User Password change
type Password struct {
  UserId   string
  Password string
}

// Update user password
func (p Password) UpdatePassword(c *Client) error {
  // check definition of required fields
  if p.UserId == "" || p.Password == "" {
    return &errorString{ "UserId and Password are required to update Proxmox user account", }
  }

  // POST parameters
  pbody := structToMap(&p, []string{}, []string{} )

  _, _, err := c.NewRequest("PUT", "/api2/json/access/password", pbody )
  if err != nil {
    return err
  }
  return err
}
