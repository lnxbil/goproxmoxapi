package goproxmoxapi_test

import (
  "testing"
  "github.com/ncerny/goproxmoxapi"
)

func TestRecentTasksAPI(t *testing.T) {
  c, err := goproxmoxapi.New(GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  /*
  // test version number of the PVE
  // this test fails intermitently - disabled for now, this call is a duplication anyways
  tasks, err := goproxmoxapi.GetRecentTasks(c)
  if err != nil {
    t.Error(err)
  }
  if len(tasks) < 1 {
    t.Error("Recent Task list can't be empty")
  }
  */
}
