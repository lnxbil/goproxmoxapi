package goproxmoxapi_test

import (
  "testing"
  "github.com/ncerny/goproxmoxapi"
)

func TestRolesAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New("root", "P@ssw0rd", "pam", "10.255.0.5")
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // Test that we can get role (clone existing with different name)
  xx, err := (goproxmoxapi.Role{ RoleId: "Administrator" }).GetRole(c)
  if err != nil {
    t.Log(xx)
    t.Error(err)
  }

  xx.RoleId = "SysAdm"
  err = xx.CreateRole(c)
  if err != nil {
    t.Error(err)
  }
  err = (goproxmoxapi.Role{ RoleId: "SysAdm" }).DeleteRole(c)
  if err != nil {
    t.Error(err)
  }

  // test that we can create groups
  rr1 := goproxmoxapi.Role{ RoleId: "SuperAdmin", Privs: "Pool.Allocate,VM.Audit" }
  err = rr1.CreateRole(c)
  if err != nil {
    t.Error(err)
  }

  // test removal of created role
  err = rr1.DeleteRole(c)
  if err != nil {
    t.Error(err)
  }

  rr2, err := goproxmoxapi.GetAllRoles(c)
  if err != nil {
    t.Log(rr2)
    t.Error(err)
  }
  if len(rr2) != 12 {
    t.Log(rr2)
    t.Log(len(rr2))
    t.Error("Number of default groups should be 12")
  }
}
