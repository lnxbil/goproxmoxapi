package goproxmoxapi_test

import (
  "testing"
  "github.com/lnxbil/goproxmoxapi"
)

func TestPasswordAPI(t *testing.T) {
  c, err := goproxmoxapi.New(goproxmoxapi.GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // define test user
  tu1 := goproxmoxapi.User{
    UserId: "usertotestpass@pve",
    Comment: "User Created via API",
    Email: "usertotestpass@example.com",
    Enable: 1,
    FirstName: "PassTest",
    Groups: []string{"gr1"},
    LastName: "User",
    Password: "inipassA",
  }

  // define test user password update
  pu1 := goproxmoxapi.Password{
    UserId: tu1.UserId,
    Password: "newpassA",
  }

  // Create user
  err = tu1.CreateUser( c )
  if err != nil {
    t.Error(err)
  }

  // Test connection with initial pass
  ctu1, err := goproxmoxapi.New("usertotestpass", tu1.Password, "pve", goproxmoxapi.GetProxmoxHost())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }
  _, err = goproxmoxapi.GetVersion(ctu1)
  if err != nil {
    t.Error(err)
  }

  // Update pass
  err = pu1.UpdatePassword(c)
  if err != nil {
    t.Error(err)
  }

  // Test connection with new pass
  cpu1, err := goproxmoxapi.New("usertotestpass", pu1.Password, "pve", goproxmoxapi.GetProxmoxHost())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  _, err = goproxmoxapi.GetVersion(cpu1)
  if err != nil {
    t.Error(err)
  }

  // Delete User
  err = tu1.DeleteUser( c )
  if err != nil {
    t.Error(err)
  }
}
