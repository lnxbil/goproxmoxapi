package goproxmoxapi_test

import (
  "testing"
  "github.com/ncerny/goproxmoxapi"
)

func TestRecentLogsAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New(GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test version number of the PVE
  logs, err := goproxmoxapi.GetRecentLogs(c)
  if err != nil {
    t.Error(err)
  }
  if len(logs) < 1 {
    t.Error("Recent log entry list can't be empty")
  }
}
