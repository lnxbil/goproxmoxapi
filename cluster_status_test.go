package goproxmoxapi_test

import (
  "testing"
  "github.com/isindir/goproxmoxapi"
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
    t.Error("We have at least 1 node in cluster, record number must not be less than 1")
  }
}
