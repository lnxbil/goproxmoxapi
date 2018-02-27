package goproxmoxapi_test

import (
  "testing"
  "github.com/ncerny/goproxmoxapi"
)

func TestClusterResourcesAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New(GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test version number of the PVE
  tasks, err := goproxmoxapi.GetClusterResources(c)
  if err != nil {
    t.Error(err)
  }
  if len(tasks) < 1 {
    t.Error("More than 1 resource is defined on system by default.")
  }
}
