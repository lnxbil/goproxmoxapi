package goproxmoxapi

import (
  "strconv"
)

// structure representing Proxmox Node
type Node struct {
  Id       string   //  "node/pve"
  Cpu      float64  //  0.0165765696901382
  Disk     float64  //  1336041472
  Level    string   //  ""
  MaxCpu   int      //  1
  MaxDisk  float64  //  26023899136
  MaxMem   float64  //  4143636480
  Mem      float64  //  705904640
  Node     string   //  "pve"
  Type     string   //  "node"
  UpTime   int      //  29488
}

// Structure representing Lxc CT to create or restore from backup
type LxcConfig struct {
  Node                 string  // required - pve node name to create CT on
  VMId                 int     // The (unique) ID of the VM
  OsTemplate           string  // required - OS Template or backup file
  Password             string  // optional - Sets root password inside container.
  Storage              string  // Default Storage. ( default: 'local' )
  Force                int     // optional - Allow to overwrite existing container.
  Restore              int     // optional - Mark this as restore task.
  Pool                 string  // optional - Add the VM to the specified pool.
  //  IgnoreUnpackErrors   int     // optional - Ignore errors when extracting the template.
  //  SshPublicKeys        string  // optional - Setup public SSH keys (one key per line, OpenSSH format).
  Lock                 string  // optional - Lock/unlock the VM. ( enum: migrate|backup|snapshot|rollback )
  Onboot               int     // optional - Specifies whether a VM will be started during system bootup. ( default: 0 )
  Startup              string  // optional - [[order=]\d+][,up=\d+][,down=\d+]
  Arch                 string  // optional - OS architecture type. ( enum: amd64|i386, default: amd64 )
  OsType               string  // optional - OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup.  ( enum: debian|ubuntu|centos|fedora|opensuse|archlinux|alpine|gentoo|unmanaged )
  Console              int     // optional - Attach a console device (/dev/console) to the container. ( default: 1 )
  TTY                  int     // optional - Specify the number of tty available to the container ( default: 2 )
  CpuLimit             int     // optional - Limit of CPU usage (0-128). NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit. ( default: 0 )
  CpuUnits             int // optional - CPU weight (0-500000) for a VM. Argument is used in the kernel fair scheduler. The larger the number is, the more CPU time this VM gets. Number is relative to the weights of all the other running VMs. NOTE: You can disable fair-scheduler configuration by setting this to 0. ( default: 1024 )
  Memory               int     // optional - Amount of RAM for the VM in MB. ( minimum: 16, default: 512 )
  Swap                 int     // optional - Amount of SWAP for the VM in MB. ( default: 512 )
  HostName             string  // optional - Set a host name for the container.
  Description          string  // optional - Container description. Only used on the configuration web interface.
  SearchDomain         string  // optional - Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
  NameServer           string  // optional - Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
  //  RootFS               string  // [volume=]<volume>[,acl=<1|0>][,quota=<1|0>][,ro=<1|0>][,size=<DiskSize>]
  Parent               string  // optional - Parent snapshot name. This is used internally, and should not be modified.
}

// Sets Defaults for the LxcConfig
func NewLxcConfig(obj *LxcConfig) *LxcConfig {
  if obj == nil {
    obj = &LxcConfig{}
  }

  if obj.Storage == "" {
    obj.Storage = "local"
  }
  if obj.Arch == "" {
    obj.Arch = "amd64"
  }
  if obj.Console == 0 {
    obj.Console = 1
  }
  if obj.TTY == 0 {
    obj.TTY = 2
  }
  if obj.CpuUnits == 0 {
    obj.CpuUnits = 1024
  }
  if obj.Memory == 0 {
    obj.Memory = 512
  }
  if obj.Swap == 0 {
    obj.Swap = 512
  }
  if obj.HostName == "" {
    obj.HostName = strconv.Itoa( obj.VMId )
  }
  //  if obj.RootFS == "" {
  //    obj.RootFS = "local-lvm"
  //  }
  return obj
}

type Lxc struct {
  Cpu        int      //  0
  Cpus       string   //  "1"
  Disk       int      //  0
  Diskread   int      //  0
  Diskwrite  int      //  0
  Lock       string   //  ""
  Maxdisk    float64  //  8589934592
  Maxmem     float64  //  536870912
  Maxswap    float64  //  536870912
  Mem        int      //  0
  Name       string   //  "100"
  Netin      int      //  0
  Netout     int      //  0
  Status     string   //  "stopped"
  Swap       int      //  0
  Template   string   //  ""
  Type       string   //  "lxc"
  Uptime     float64  //  0
  Vmid       string   //  "100"
}

// Destroy the container (also delete all uses files). Returns Task PuId.
func (ct LxcConfig) DeleteLxc(c *Client) (string, error) {
  tskinfo := ""
  // check definition of required fields
  if ct.Node == "" {
    return tskinfo, &errorString{ "Node name (Node.Node) is required to destroy container", }
  }
  if ct.VMId == 0 {
    return tskinfo, &errorString{ "VMId is required to destroy container", }
  }

  _, rbody, err := c.NewRequest("DELETE", "/api2/json/nodes/" + ct.Node + "/lxc/" + strconv.Itoa( ct.VMId ), nil )
  if err != nil {
    return tskinfo, err
  } else {
    err = dataUnmarshal( rbody, &tskinfo )

    return tskinfo, err
  }
}

// Create or restore a LXC container. Returns Task PuId.
func (ct LxcConfig) CreateLxc(c *Client) (string, error) {
  tskinfo := ""
  // check definition of required fields
  if ct.Node == "" {
    return tskinfo, &errorString{ "Node name is required to create container", }
  }

  // POST parameters
  pbody := structToMap(&ct, []string{}, []string{ "Node" } )

  _, rbody, err := c.NewRequest("POST", "/api2/json/nodes/" + ct.Node + "/lxc", pbody )
  if err != nil {
    return tskinfo, err
  } else {
    err = dataUnmarshal( rbody, &tskinfo )

    return tskinfo, err
  }

}

// Returns all Lxc containers per node
func GetAllLxc(c *Client, nodeName string) ([]Lxc, error) {
  cts := make([]Lxc, 0)
  if nodeName == "" {
    return cts, &errorString{ "Node name (Node.Node) is required to get list of containers", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/nodes/" + nodeName + "/lxc", nil )
  if err != nil {
    return cts, err
  } else {
    err = dataUnmarshal( rbody, &cts )

    return cts, err
  }
}

// Returns all defined cluster nodes
func GetAllNodes(c *Client) ([]Node, error) {
  nodes := make([]Node, 0)

  _, rbody, err := c.NewRequest("GET", "/api2/json/nodes", nil )
  if err != nil {
    return nodes, err
  } else {
    err = dataUnmarshal( rbody, &nodes )

    return nodes, err
  }
}
