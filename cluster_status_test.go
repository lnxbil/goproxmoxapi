package goproxmoxapi_test

import (
  "testing"
  "github.com/ncerny/goproxmoxapi"
)

func TestClusterStatusAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New("root", "P@ssw0rd", "pam", "10.255.0.5")
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test version number of the PVE
  tasks, err := goproxmoxapi.GetClusterStatus(c)
  if err != nil {
    t.Error(err)
  }
  if len(tasks) < 1 {
    t.Error("Record number must not be less than 1")
  }

  pts := goproxmoxapi.TaskEntry{ Node: "pve" }
  ftasks, err := pts.GetFinishedTasks(c)
  if err != nil {
    t.Error(err)
  }
  if len(ftasks) < 1 {
    t.Error("Record number must not be less than 1")
  }
}
