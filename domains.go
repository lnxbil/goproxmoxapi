package goproxmoxapi

// structure representing Proxmox Domain
type Domain struct {
  // Common for all types
  Realm     string  // required - Authentication domain ID
  Type      string  // required - Realm type (ad|ldap|pam|pve), creation of pam,pve not allowed
  Comment   string  // optional - Description.
  // LDAP & AD
  Server1   string  // optional - Server IP address (or DNS name)
  Server2   string  // optional - Fallback Server IP address (or DNS name)
  Secure    string     // optional - Use secure LDAPS protocol.
  Default   int     // optional - Use this as default realm
  Port      string  // optional - Server port.
  TFA       string  // optional - Use Two-factor authentication.
  // AD
  Domain    string  // optional - AD domain name
  // LDAP
  Base_DN   string  // optional - LDAP base domain name
  Bind_DN   string  // optional - LDAP bind domain name
  User_Attr string  // optional - LDAP user attribute name
}

// Helper function:
// Returns a predefined list of wanted params depending on the type of the domain
func (d Domain) genWantedStorageParams() []string {
  commons := []string{ "realm", "type", "comment" }
  switch {
  case d.Type == "pam":
    return commons
  case d.Type == "pve":
    return commons
  case d.Type == "ad":
    return append( commons, "server1", "server2",  "secure",  "default",  "port",  "tfa", "domain"   )
  case d.Type == "ldap":
    return append( commons, "server1", "server2",  "secure",  "default",  "port",  "tfa", "base_dn", "bind_dn", "user_attr" )
  }
  return []string{}
}

// Create new Proxmox domain
func (d Domain) CreateDomain(c *Client) error {
  // check definition of required fields
  if d.Realm == "" {
    return &errorString{ "Realm is required to create Proxmox Domain", }
  }

  // POST parameters
  pbody := structToMap(&d, d.genWantedStorageParams(), []string{} )

  _, _, err := c.NewRequest("POST", "/api2/json/access/domains", pbody )
  if err != nil {
    return err
  }
  return nil
}

// Delete Proxmox Domain
func (d Domain) DeleteDomain(c *Client) error {
  // check definition of required fields
  if d.Realm == "" {
    return &errorString{ "Realm is required to delete Proxmox Domain", }
  }

  _, _, err := c.NewRequest("DELETE", "/api2/json/access/domains/" + d.Realm, nil )
  if err != nil {
    return err
  }
  return err
}

// Update Proxmox user
func (d Domain) UpdateDomain(c *Client) error {
  // check definition of required fields
  if d.Realm == "" {
    return &errorString{ "Realm is required to update Proxmox Domain", }
  }

  // POST parameters
  pbody := structToMap(&d, d.genWantedStorageParams(), []string{} )
  //TODO: Needs to be thorohly tested and all constraints to be included here for all domain types
  pbody = structCleanUp( pbody, []string{}, []string{ "realm", "type" } )

  _, _, err := c.NewRequest("PUT", "/api2/json/access/domains/" + d.Realm, pbody )
  return err
}

// Get Proxmox Domain information
func (d Domain) GetDomain(c *Client) (Domain, error) {
  dmn := Domain{}
  // check definition of required fields for this call
  if d.Realm == "" {
    return dmn, &errorString{ "Realm is required to get Proxmox Domain", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/access/domains/" + d.Realm, nil )
  if err != nil {
    return dmn, err
  } else {
    err = dataUnmarshal( rbody, &dmn )
    dmn.Realm = d.Realm

    return dmn, err
  }
}

// Gets all defined domains
func GetAllDomains(c *Client) ([]Domain, error) {
  domains := make([]Domain, 0)

  _, rbody, err := c.NewRequest("GET", "/api2/json/access/domains", nil )
  if err != nil {
    return domains, err
  } else {
    err = dataUnmarshal( rbody, &domains )

    return domains, err
  }
}
