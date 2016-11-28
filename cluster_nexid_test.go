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

  // check if we can get "200" as next available id
//*
// Looks like it's not possible to pass this optional value to the API call
// via GET as API does not implement normal query params parsing
//*
//	vmid2, err := goproxmoxapi.GetNextVMID(c, 200)
//  if err != nil {
//    t.Error(err)
//  }
//  if vmid2 != "200" {
//    t.Error("Next available ID has to be 200")
//  }
}
