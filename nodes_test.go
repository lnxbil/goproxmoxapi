package goproxmoxapi_test

import (
  "testing"
  "time"
  "github.com/isindir/goproxmoxapi"
)

func TestNodesAPI(t *testing.T) {
  t.Parallel()

  // Establish new session
  c, err := goproxmoxapi.New("root", "P@ssw0rd", "pam", "10.255.0.5")
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // Test number of nodes in the cluster is at least 1
  nodes1, err := goproxmoxapi.GetAllNodes(c)
  if err != nil {
    t.Log(nodes1)
    t.Error(err)
  }
  if len(nodes1) < 1 {
    t.Error("Number of defined nodes must be at least 1 (our test pve server).")
  }

  // Test that we have LXC container configured on pve node
  lxcs1, err := goproxmoxapi.GetAllLxc(c, nodes1[0].Node)
  if err != nil {
    t.Error(err)
  }
  if len(lxcs1) != 1 {
    t.Error("Number of defined lxc containers must be 1 (on our test pve server).")
  }

  // define test lxc ct config
  //    pvesh create /nodes/pve/lxc -vmid 101 -hostname test -password qwerty12 \
  //      -storage local-lvm -ostemplate local:vztmpl/centos-7-default_20160205_amd64.tar.xz -memory 512 -swap 512
  ct1 := goproxmoxapi.NewLxcConfig( &goproxmoxapi.LxcConfig{
    Node: "pve",
    VMId: 300,
    Password: "qwerty12",
    Storage: "local-lvm",
    OsTemplate: "local:vztmpl/centos-7-default_20160205_amd64.tar.xz",
//    Pool       
//    Onboot     
//    Startup    
//    Template: "",
    Description: "Test LXC Container",
    //RootFS: "local-lvm,size=8G",
    Arch: "amd64",
    OsType: "centos",
    Memory: 1024,
    Swap: 512,
    HostName: "testct",
    SearchDomain: "example.com",
    NameServer: "4.4.4.4,10.255.0.5",
  })

  // test that we can create Lxc container
  ss, err := ct1.CreateLxc( c )
  if err != nil {
    t.Log( ss )
    t.Error(err)
  }

  // Wait for create operation to finish (Using Proxmox TaskStatus) and only then destroy container
  ch1 := make(chan int)
  pts := goproxmoxapi.TaskEntry{ Node: "pve", UpId: ss }
  tsts := goproxmoxapi.TaskEntry{}
  go func() {
    for tsts, err = pts.GetTaskStatus( c ); tsts.Status != "stopped"; {
      if err != nil {
        t.Error( err )
        ch1 <- 1
      }
      time.Sleep(time.Second * 1)
      tsts, err = pts.GetTaskStatus( c )
    }
    ch1 <- 1
  }()
  // wait for task to complete
  <-ch1

  // test that we can obtain a log entries for a given task
  lgs, err := pts.GetTaskLogEntries( c )
  if err != nil || len(lgs) != 22 {
    t.Log( lgs )
    t.Error(err)
  }

  // Test that we can destroy Lxc container
  if err != nil && tsts.ExitStatus != "OK" {
    t.Log( tsts.ExitStatus )
    t.Log( tsts.UpId )
    t.Error( err )
  } else {
    ss, err = ct1.DeleteLxc( c )
    if err != nil {
      t.Error(err)
    }
  }

  // Wait for Delete Container task to finish and only than return
  ch2 := make(chan int)
  pts = goproxmoxapi.TaskEntry{ Node: "pve", UpId: ss }
  tsts = goproxmoxapi.TaskEntry{}
  go func() {
    for tsts, err = pts.GetTaskStatus( c ); tsts.Status != "stopped"; {
      if err != nil {
        t.Error( err )
        ch2 <- 1
      }
      time.Sleep(time.Second * 1)
      tsts, err = pts.GetTaskStatus( c )
    }
    ch2 <- 1
  }()

  // Test that we can destroy Lxc container
  <-ch2

}
