package goproxmoxapi

// structure representing Proxmox Version
type Version struct {
  Keyboard string
  Release  string
  RepoId   string
  Version  string
}

// Get Proxmox version
func GetVersion(c *Client) (Version, error) {
  pvever := Version{}

  _, rbody, err := c.NewRequest("GET", "/api2/json/version", nil )
  if err != nil {
	  return pvever, err
  } else {
    err = dataUnmarshal( rbody, &pvever )

    // Any Error parsing json ?
		return pvever, err
  }
}
