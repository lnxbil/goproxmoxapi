package goproxmoxapi_test

import (
  "testing"
  "github.com/isindir/goproxmoxapi"
)

func TestVersionAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New("root", "P@ssw0rd", "pam", "10.255.0.5")
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test version number of the PVE
  pvever, err := goproxmoxapi.GetVersion(c)
  if err != nil {
    t.Error(err)
  }
  if pvever.Version != "4.3" {
    t.Error("Should run this test against pve version 4.3")
  }
}
