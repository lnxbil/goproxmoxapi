package goproxmoxapi

import (
  "testing"
  "net/url"
)

func TestClientInternalAPI(t *testing.T) {
  t.Parallel()
  if !( atobt("1")==true && atobt("0")==false && atobt("a")==false ) {
    t.Log(atobt("1"), atobt("0"), atobt("a"))
    t.Error()
  }

  if !( atobf("1")==true && atobf("0")==false && atobf("a")==true ) {
    t.Log(atobf("1"), atobf("0"), atobf("a"))
    t.Error()
  }

  if !(btoa(true)=="1" && btoa(false)=="0") {
    t.Log(btoa(true), btoa(false))
    t.Error()
  }

  // test structCleanUp function
  tvals := url.Values{}
  tvals.Set( "aaa", "a" )
  tvals.Set( "bbb", "b" )
  tvals.Set( "ccc", "c" )
  tvals.Set( "ddd", "d" )
  tvals1 := structCleanUp(tvals, []string{ "aaa", "bbb", "eee" }, []string{} )
  if tvals1.Encode() != "aaa=a&bbb=b" {
    t.Error()
  }
  tvals.Set( "aaa", "a" )
  tvals.Set( "bbb", "b" )
  tvals.Set( "ccc", "c" )
  tvals.Set( "ddd", "d" )
  tvals2 := structCleanUp(tvals, []string{}, []string{ "aaa", "bbb", "eee" } )
  if tvals2.Encode() != "ccc=c&ddd=d" {
    t.Error()
  }
  tvals.Set( "aaa", "a" )
  tvals.Set( "bbb", "b" )
  tvals.Set( "ccc", "c" )
  tvals.Set( "ddd", "d" )
  tvals3 := structCleanUp(tvals, []string{ "aaa", "bbb", "ccc", "ddd" }, []string{ "aaa", "bbb", "ccc", "ddd" } )
  if tvals3.Encode() != "aaa=a&bbb=b&ccc=c&ddd=d" {
    t.Log( tvals.Encode() )
    t.Log( tvals3.Encode() )
    t.Error()
  }

  //TODO write test for structToMap
}
