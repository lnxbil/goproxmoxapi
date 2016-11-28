package goproxmoxapi

// structure representing Proxmox Storage
type Storage struct {
  // Settings common to all storage types
  Storage               string // Common: PVE Storage Id
  Type                  string // Common: Storage type (dir|drbd|glusterfs|iscsi|iscsidirect|lvm|lvmthin|nfs|rbd|sheepdog|zfs|zfspool)
  Content               string // Common: Allowed content types. NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs. (images,rootdir; vztmpl,iso,backup)
  Disable               int    // Common: Flag to disable the storage. (boolean)
  MaxFiles              string // Common: Maximal number of backup files per VM. Use '0' for unlimted.
  Shared                int    // Common: optional - Mark storage as shared. (boolean)
  Format                string // Common: optional - Default image format.
  Nodes                 string // Common: Node list
  // Directory
  //   'Shared' is not listed for this type
  Path                  string // Dir: File system path.
  MkDir                 string // Dir: Create the directory if it doesn't exist. (boolean, default: true)
  // Glusterfs
  Volume                string // Glusterfs: Glusterfs Volume.
  Server2               string // Glusterfs: Backup volfile server IP or DNS name. Requires 'server'
  Transport             string // Glusterfs: Gluster transport: tcp or rdma (tcp|rdma|unix)
  // iSCSI
  Portal                string // iSCSI: iSCSI portal (IP or DNS name with optional port).
  Target                string // iSCSI: iSCSI target.
  // LVM
  VGName                string // LVM: Volume group name.
  Base                  string // LVM: Base volume. This volume is automatically activated.
  SafeRemove            int    // LVM: Zero-out data when removing LVs. (boolean)
  SafeRemove_ThroughPut string // LVM: Wipe throughput (cstream -t parameter value).
  Tagged_Only           int    // LVM: Only use logical volumes tagged with 'pve-vm-ID'. (boolean)
  ThinPool              string // LvmThin: LVM thin pool LV name.
  // NFS
  Export                string // NFS: NFS export path.
  Server                string // NFS: Server IP or DNS name.
  Options               string // NFS: NFS mount options (see 'man nfs')
  // DRBD
  Redundancy            int    // DRBD: The redundancy count specifies the number of nodes to which the resource should be deployed. It must be at least 1 and at most the number of nodes in the cluster.
  MonHost               string // RBDB: Monitors daemon ips.
  Pool                  string // RBDB: Pool.
  UserName              string // RBDB: RBD Id.
  AuthSupported         string // RBDB: Authsupported.
  KRBD                  int    // RBDB: Access rbd through krbd kernel module. (boolean)
  // ZFS
  ISCSIProvider         string // ZFS: iscsi provider
  NowriteCache          int    // ZFS: disable write caching on the target (boolean)
  Comstar_tg            string // ZFS: target group for comstar views
  Comstar_hg            string // ZFS: host group for comstar views
  Blocksize             string // ZFSPool: block size
  Sparse                int    // ZFSPool: use sparse volumes (boolean)
}

// Helper function:
// Returns a predefined list of wanted params depending on the type of the storage
func (s Storage) genWantedStorageParams() []string {
    commons := []string{ "storage", "type", "content", "disable", "format", "nodes" }
    switch {
    case s.Type == "dir":
      return append( commons, "maxfiles", "shared", "path", "mkdir" )
    case s.Type == "lvm":
      return append( commons, "shared", "vgname", "base", "saferemove", "saferemove_throughput", "tagged_only" )
    case s.Type == "lvmthin":
      return append( commons, "vgname", "thinpool" )
    case s.Type == "nfs":
      return append( commons, "maxfiles", "export", "server", "options" )
    case s.Type == "iscsi" || s.Type == "iscsidirect":
      return append( commons, "portal", "target" )
    case s.Type == "glusterfs":
      return append( commons, "maxfiles", "server", "server2", "volume", "transport" )
    case s.Type == "drbd" || s.Type == "rbd":
      return append( commons, "redundancy", "monhost", "pool", "username", "authsupported", "krbd" )
    // TODO: Pure guess here, more research needed
    case s.Type == "zfspool" || s.Type == "zfs":
      return append( commons, "iscsiprovider", "nowritecache", "comstar_tg", "comstar_hg", "blocksize", "sparse" )
    case s.Type == "sheepdog":
      return []string{  }
    }
    return []string{}
}

// Create new Proxmox Storage
func (s Storage) CreateStorage(c *Client) error {
  // check definition of required fields
  if s.Storage == "" {
    return &errorString{ "Storage field is required to create Proxmox Storage", }
  }

  // POST parameters
  pbody := structToMap(&s, s.genWantedStorageParams(), []string{} )

	_, _, err := c.NewRequest("POST", "/api2/json/storage", pbody )
  if err != nil {
    return err
  }
  return nil
}

// Delete Proxmox Storage
func (s Storage) DeleteStorage(c *Client) error {
  // check definition of required fields
  if s.Storage == "" {
    return &errorString{ "Storage field is required to delete Proxmox Storage", }
  }

  _, _, err := c.NewRequest("DELETE", "/api2/json/storage/" + s.Storage, nil )
  if err != nil {
    return err
  }
  return err
}

// Update Proxmox Storage
func (s Storage) UpdateStorage(c *Client) error {
  // check definition of required fields
  if s.Storage == "" {
    return &errorString{ "Storage field is required to update Proxmox Storage", }
  }

  // POST parameters
  pbody := structToMap(&s, s.genWantedStorageParams(), []string{} )
  //TODO: More constraints likely to be included here for other storage types,
  // potentialy some will be conditional
  pbody = structCleanUp(pbody, []string{}, []string{"storage", "path", "type"} )

  _, _, err := c.NewRequest("PUT", "/api2/json/storage/" + s.Storage, pbody )
  return err
}

// Get Proxmox Storage
func (s Storage) GetStorage(c *Client) (Storage, error) {
  st := Storage{}
  // check definition of required fields
  if s.Storage == "" {
    return st, &errorString{ "Storage field is required to get Proxmox Storage", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/storage/" + s.Storage, nil )
  if err != nil {
	  return st, err
  } else {
    err = dataUnmarshal( rbody, &st )

    // Any Error parsing json ?
		return st, err
  }
}

// Gets all defined Storages
func GetAllStorages(c *Client) ([]Storage, error) {
	st := make([]Storage, 0)

  // GET parameters
	_, rbody, err := c.NewRequest("GET", "/api2/json/storage", nil )
  if err != nil {
	  return st, err
  } else {
    err = dataUnmarshal( rbody, &st )

    // Any Error parsing json ?
		return st, err
  }
}
