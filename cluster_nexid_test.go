package goproxmoxapi_test

import (
  "testing"
  "github.com/isindir/goproxmoxapi"
)

func TestNextIdAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New("root", "P@ssw0rd", "pam", "10.255.0.5")
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // check if we can get next available id (expected value is 101 as 1 
  // CT has been created manually for the testing purposes)
  vmid1, err := goproxmoxapi.GetNextVMID(c)
  if err != nil {
    t.Error(err)
  }
  if vmid1 != "101" {
    t.Error("Next available ID has to be 101")
  }
}
