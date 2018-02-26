package goproxmoxapi_test

import (
  "testing"
  "time"
  "github.com/ncerny/goproxmoxapi"
)

func TestLxcOpAPI(t *testing.T) {
  t.Parallel()

  // Establish new session
  c, err := goproxmoxapi.New(goproxmoxapi.GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // define basic test lxc ct config for Lxc operations
  ct1 := goproxmoxapi.NewLxcConfig( &goproxmoxapi.LxcConfig{
    Node: goproxmoxapi.GetProxmoxNode(),
    VMId: 100,
  })

  // Test that test container is down
  cts, err := ct1.GetLxcStatus( c )
  if cts.Type != "lxc" {
    t.Log( cts )
    t.Error("Unexpected type of container")
  }
  if err != nil {
    t.Error(err)
  }

  // Test that operation we are going to perform is invalid
  _, err = ct1.LxcOp( c, "current" )
  if err == nil {
    t.Error("Expecting to fail this operation")
  }

  // Test that we can start test containder ( valid operation )
  ss, err := ct1.LxcOp( c, "start" )
  if err != nil {
    t.Error(err)
  }

  // Wait for create operation to finish (Using Proxmox TaskStatus) and only then destroy container
  ch1 := make(chan int)
  pts := goproxmoxapi.TaskEntry{Node: goproxmoxapi.GetProxmoxNode(), UpId: ss}
  tsts := goproxmoxapi.TaskEntry{}
  go func() {
    for tsts, err = pts.GetTaskStatus( c ); tsts.Status != "running"; {
      if err != nil {
        t.Error( err )
        ch1 <- 1
      }
      time.Sleep(time.Millisecond * 500)
      tsts, err = pts.GetTaskStatus( c )
    }
    ch1 <- 1
  }()
  // wait for task to complete
  <-ch1

  // Test that we can start test containder ( valid operation )
  ss, err = ct1.LxcOp( c, "stop" )
  if err != nil {
    t.Error(err)
  }

  // Wait for create operation to finish (Using Proxmox TaskStatus) and only then destroy container
  ch1 = make(chan int)
  pts = goproxmoxapi.TaskEntry{Node: goproxmoxapi.GetProxmoxNode(), UpId: ss}
  go func() {
    for tsts, err = pts.GetTaskStatus( c ); tsts.Status != "stopped"; {
      if err != nil {
        t.Error( err )
        ch1 <- 1
      }
      time.Sleep(time.Millisecond * 500)
      tsts, err = pts.GetTaskStatus( c )
    }
    ch1 <- 1
  }()
  // wait for task to complete
  <-ch1
}
