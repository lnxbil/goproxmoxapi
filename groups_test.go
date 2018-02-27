package goproxmoxapi_test

import (
  "testing"
  "github.com/lnxbil/goproxmoxapi"
)

func TestGroupAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New(GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test that we can create groups
  err = (goproxmoxapi.Group{ GroupId: "gr3", Comment: "test3" }).CreateGroup(c)
  if err != nil {
    t.Error(err)
  }
  err = (goproxmoxapi.Group{ GroupId: "gr4" }).CreateGroup(c)
  if err != nil {
    t.Error(err)
  }

  // test that we can fetch created groups
  _, err = (goproxmoxapi.Group{ GroupId: "gr3" }).GetGroup(c)
  if err != nil {
    t.Error(err)
  }
  _, err = (goproxmoxapi.Group{ GroupId: "gr4" }).GetGroup(c)
  if err != nil {
    t.Error(err)
  }

  // test that we can update group
  err = ( goproxmoxapi.Group{ GroupId: "gr4", Comment: "test4" } ).UpdateGroup(c)
  if err != nil {
    t.Error(err)
  }

  // test the fact that group updated
  gr4, err := (goproxmoxapi.Group{ GroupId: "gr4" }).GetGroup(c)
  if err != nil {
    t.Error(err)
  }
  if gr4.Comment != "test4" && gr4.GroupId != "gr4" {
    t.Error("Updating gr4 comment field has failed")
  }

  // Test number of groups is at least 2
  gg1, err := goproxmoxapi.GetAllGroups(c)
  if err != nil {
    t.Log(gg1)
    t.Error(err)
  }
  if len(gg1) < 2 {
    t.Error("The number of groups must be at least 2")
  }

  // Test we can delete groups
  err = ( goproxmoxapi.Group{ GroupId: "gr3", } ).DeleteGroup(c)
  if err != nil {
    t.Error(err)
  }

  err = ( goproxmoxapi.Group{ GroupId: "gr4", } ).DeleteGroup(c)
  if err != nil {
    t.Error(err)
  }

  // test the fact that both groups are removed
  _, err = (goproxmoxapi.Group{ GroupId: "gr3" }).GetGroup(c)
  if err == nil {
    t.Error(err)
  }

  _, err = (goproxmoxapi.Group{ GroupId: "gr4" }).GetGroup(c)
  if err == nil {
    t.Error(err)
  }

  // Test number of groups should be 2 less after removal of 2 groups
  gg2, err := goproxmoxapi.GetAllGroups(c)
  if err != nil {
    t.Log(gg2)
    t.Error(err)
  }
  if len(gg1) != 2+len(gg2) {
    t.Log(gg1)
    t.Log(gg2)
    t.Error("Number of groups before removal should be greater than after")
  }
}
