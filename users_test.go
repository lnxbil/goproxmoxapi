package goproxmoxapi_test

import (
  "testing"
  "github.com/lnxbil/goproxmoxapi"
)

func TestUserAPI(t *testing.T) {
  // Establish new session
  c, err := goproxmoxapi.New(GetProxmoxAccess())
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // define test user
  tu1 := goproxmoxapi.User{
    UserId: "newtestuser1@pve",
    Comment: "User Created via API",
    Email: "newtestuser@example.com",
    Enable: 1,
    FirstName: "Newtest",
    Groups: []string{"gr1", "gr2"},
    Keys: "xx",
    LastName: "User",
    Password: "P@ssw0rd!",
  }
  // define test user
  tu2 := goproxmoxapi.User{
    UserId: "newtestuser2@pve",
    Comment: "User Created via API",
    Email: "newtestuser@example.com",
    Enable: 1,
    FirstName: "Newtest",
    Groups: []string{"gr1"},
    Keys: "xx",
    LastName: "User",
    Password: "P@ssw0rd!",
  }

  // test that we can create users
  err = tu1.CreateUser( c )
  if err != nil {
    t.Error(err)
  }

  err = tu2.CreateUser( c )
  if err != nil {
    t.Error(err)
  }

  // test that we can fetch created users
  _, err = tu1.GetUser( c )
  if err != nil {
    t.Error(err)
  }

  xxu1, err := tu2.GetUser( c )
  if err != nil {
    t.Error(err)
  }
  if len(xxu1.Groups) != 1 {
    t.Error("User should belong to 1 group")
  }

  // Update user and check that number of assigned groups increased
  tu2.Groups = []string{"gr1", "gr2"}
  tu2.LastName = "Hey Joe"
  err = tu2.UpdateUser( c )
  if err != nil {
    t.Error(err)
  }
  xxu2, err := tu2.GetUser( c )
  if err != nil {
    t.Error(err)
  }
  if len(xxu2.Groups) != 2 {
    t.Error("User should belong to 2 groups")
  }
  if xxu2.LastName != "Hey Joe" {
    t.Error("Lastname should be updated")
  }

  // Test number of users is at least 2
  uu1, err := goproxmoxapi.GetAllUsers(c)
  if err != nil {
    t.Log(uu1)
    t.Error(err)
  }
  if len(uu1) < 2 {
    t.Error("The number of users must be at least 2")
  }

  // Test user deletion
  err = tu1.DeleteUser( c )
  if err != nil {
    t.Error(err)
  }

  err = tu2.DeleteUser( c )
  if err != nil {
    t.Error(err)
  }

  // test the fact that both users are removed
  _, err = tu1.GetUser( c )
  if err == nil {
    t.Error("User " + tu1.UserId + " should be removed.")
  }

  _, err = tu2.GetUser( c )
  if err == nil {
    t.Error("User " + tu2.UserId + " should be removed.")
  }

  // Test number of users should be 2 less after removal of 2 users
  uu2, err := goproxmoxapi.GetAllUsers(c)
  if err != nil {
    t.Log(uu2)
    t.Error(err)
  }
  if len(uu1) != 2+len(uu2) {
    t.Log(uu1)
    t.Log(uu2)
    t.Error("Number of users before removal should be greater than after")
  }

}
