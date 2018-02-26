package goproxmoxapi_test

import (
  "testing"
  "github.com/lnxbil/goproxmoxapi"
)

func TestPoolAPI(t *testing.T) {
  t.Parallel()
  c, err := goproxmoxapi.New(goproxmoxapi.GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test that we can create pools
  err = (goproxmoxapi.Pool{ PoolId: "pl3", Comment: "test3" }).CreatePool(c)
  if err != nil {
    t.Error(err)
  }
  err = (goproxmoxapi.Pool{ PoolId: "pl4" }).CreatePool(c)
  if err != nil {
    t.Error(err)
  }

  // test that we can fetch created pools
  _, err = (goproxmoxapi.Pool{ PoolId: "pl3" }).GetPool(c)
  if err != nil {
    t.Error(err)
  }
  _, err = (goproxmoxapi.Pool{ PoolId: "pl4" }).GetPool(c)
  if err != nil {
    t.Error(err)
  }

  // test that we can update pool
  err = ( goproxmoxapi.Pool{ PoolId: "pl4", Comment: "test4" } ).UpdatePool(c)
  if err != nil {
    t.Error(err)
  }

  // test the fact that pool updated
  pl4, err := (goproxmoxapi.Pool{ PoolId: "pl4" }).GetPool(c)
  if err != nil {
    t.Error(err)
  }
  if pl4.Comment != "test4" && pl4.PoolId != "pl4" {
    t.Error("Updating pl4 comment field has failed")
  }

  // Test number of pools is at least 2
  pp1, err := goproxmoxapi.GetAllPools(c)
  if err != nil {
    t.Log(pp1)
    t.Error(err)
  }
  if len(pp1) < 2 {
    t.Error("The number of pools must be at least 2")
  }

  // Test we can delete pools
  err = ( goproxmoxapi.Pool{ PoolId: "pl3", } ).DeletePool(c)
  if err != nil {
    t.Error(err)
  }

  err = ( goproxmoxapi.Pool{ PoolId: "pl4", } ).DeletePool(c)
  if err != nil {
    t.Error(err)
  }

  // test the fact that both pools are removed
  _, err = (goproxmoxapi.Pool{ PoolId: "pl3" }).GetPool(c)
  if err == nil {
    t.Error(err)
  }

  _, err = (goproxmoxapi.Pool{ PoolId: "pl4" }).GetPool(c)
  if err == nil {
    t.Error(err)
  }

  // Test number of pools should be 2 less after removal of 2 pools
  pp2, err := goproxmoxapi.GetAllPools(c)
  if err != nil {
    t.Log(pp2)
    t.Error(err)
  }
  if len(pp1) != 2+len(pp2) {
    t.Log(pp1)
    t.Log(pp2)
    t.Error("Number of pools before removal should be greater than after")
  }
}
