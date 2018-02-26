package goproxmoxapi_test

import (
  "strings"
  "testing"
  "github.com/ncerny/goproxmoxapi"
)

func TestVersionAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New(goproxmoxapi.GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test version number of the PVE
  pvever, err := goproxmoxapi.GetVersion(c)
  if err != nil {
    t.Error(err)
  }
  if !strings.HasPrefix(pvever.Version, "5.1") {
    t.Error("Should run this test against pve version 5.x, yet you have '" + pvever.Version + "'")
  }
}
