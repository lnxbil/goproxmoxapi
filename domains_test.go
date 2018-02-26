package goproxmoxapi_test

import (
  "testing"
  "github.com/lnxbil/goproxmoxapi"
)

func TestDomainAPI(t *testing.T) {
  t.Parallel()

  // Establish new session
  c, err := goproxmoxapi.New(goproxmoxapi.GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test getting all domains
  dlist, err := goproxmoxapi.GetAllDomains(c)
  if err != nil {
    t.Log(dlist)
    t.Error(err)
  }

  // Test getting one domain
  _, err = (goproxmoxapi.Domain{Realm: goproxmoxapi.GetProxmoxRealm()}).GetDomain(c)
  if err != nil {
    t.Error(err)
  }

  // Test AD creation
  dmn1 := goproxmoxapi.Domain{
    Realm: "tst",
    Type: "ad",
    Server1: "srv1",
    Domain: "example.com",
    Comment: "Test AD domain",
    Port: 389,
  }

  err = dmn1.CreateDomain( c )
  if err != nil {
    t.Log( dmn1 )
    t.Error(err)
  }

  // Test getting new domain
  dmn11, err := (goproxmoxapi.Domain{ Realm: "tst" }).GetDomain( c )
  if err != nil {
    t.Error(err)
  }
  if dmn11.Port != 389 {
    t.Error("Wrong port number obtained after creation")
  }

  // test update of existing domain
  dmn2 := goproxmoxapi.Domain{
    Realm: "tst",
    Server2: "srv2",
    Port: 636,
    Secure: 1,
  }
  err = dmn2.UpdateDomain( c )
  if err != nil {
    t.Log( dmn2 )
    t.Error(err)
  }

  // Test getting updated domain
  dmn12, err := (goproxmoxapi.Domain{ Realm: "tst" }).GetDomain( c )
  if err != nil {
    t.Error(err)
  }
  if dmn12.Port != 636 {
    t.Error("Wrong port number obtained after update")
  }

  // Test domain deletion
  err = dmn1.DeleteDomain( c )
  if err != nil {
    t.Error(err)
  }

  /*
  realm, bdn, srv1, srv2, port, ssl, tfa, username, default, comment
  */
  // Test LDAP auth domain creation
  ldap1 := goproxmoxapi.Domain{
    Realm: "tst",
    Type: "ldap",
    Server1: "srv1",
    Comment: "Test AD domain",
    Port: 389,
    Base_DN: "cn=users,dc=abc,dc=com",
    Bind_DN: "dc=abc,dc=com",
    User_Attr: "bindusername",
  }

  err = ldap1.CreateDomain( c )
  if err != nil {
    t.Log( ldap1 )
    t.Error(err)
  }

  // test update of existing domain
  ldap2 := goproxmoxapi.Domain{
    Realm: "tst",
    Server2: "srv2",
    Comment: "Test AD domain Updated",
    Port: 636,
    Secure: 1,
    Base_DN: "cn=users,dc=example,dc=com",
    Bind_DN: "dc=example,dc=com",
    User_Attr: "bindusername",
  }
  err = ldap2.UpdateDomain( c )
  if err != nil {
    t.Log( ldap2 )
    t.Error(err)
  }

  // Test domain deletion
  err = ldap1.DeleteDomain( c )
  if err != nil {
    t.Error(err)
  }
}
