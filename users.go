package goproxmoxapi

// structure representing Proxmox User
type User struct {
	UserId    string   // Required - user@realm
  Comment   string   // Optional
	Email     string   // Optional
  Enable    int      // Optional - Enable the account (default). You can set this to '0' to disable the account.
  Expire    int      // Optional - Account expiration date (seconds since epoch). '1' means no expiration date.
	FirstName string   // Optional
  Groups    []string // Optional
  Keys      string   // Optional - Keys for two factor authentication (yubico).
  LastName  string   // Optional
  Password  string   // Optional - Initial password.
}

// Create new Proxmox user
func (u User) CreateUser(c *Client) error {
  // check definition of required fields
  if u.UserId == "" {
    return &errorString{ "UserId is required to create Proxmox user", }
  }

  // POST parameters
  pbody := structToMap(&u, []string{}, []string{} )

	_, _, err := c.NewRequest("POST", "/api2/json/access/users", pbody )
  if err != nil {
    return err
  }
  return nil
}

// Delete Proxmox user
func (u User) DeleteUser(c *Client) error {
  // check definition of required fields
  if u.UserId == "" {
    return &errorString{ "UserId is required to delete Proxmox user", }
  }

  _, _, err := c.NewRequest("DELETE", "/api2/json/access/users/" + u.UserId, nil )
  if err != nil {
    return err
  }
  return err
}

// Update Proxmox user
func (u User) UpdateUser(c *Client) error {
  // check definition of required fields
  if u.UserId == "" {
    return &errorString{ "UserId is required to update Proxmox user", }
  }

  // POST parameters
  pbody := structToMap(&u, []string{}, []string{ "userid", "password" } )

  _, _, err := c.NewRequest("PUT", "/api2/json/access/users/" + u.UserId, pbody )
  return err
}

// Get Proxmox user
func (u User) GetUser(c *Client) (User, error) {
  ru := User{}
  // check definition of required fields
  if u.UserId == "" {
    return ru, &errorString{ "UserId is required to get Proxmox user", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/access/users/" + u.UserId, nil )
  if err != nil {
	  return ru, err
  } else {
    err = dataUnmarshal( rbody, &ru )

    // Any Error parsing json ?
		return ru, err
  }
}

// Gets all defined users
func GetAllUsers(c *Client) ([]User, error) {
	users := make([]User, 0)

	_, rbody, err := c.NewRequest("GET", "/api2/json/access/users", nil )
  if err != nil {
	  return users, err
  } else {
    err = dataUnmarshal( rbody, &users )

    // Any Error parsing json ?
		return users, err
  }
}
