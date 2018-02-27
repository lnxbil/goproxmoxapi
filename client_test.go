package goproxmoxapi_test

import (
  "testing"
  "os"
  "github.com/lnxbil/goproxmoxapi"
)

func TestClientBaseAPI(t *testing.T) {
  t.Parallel()
  // test new connection
  c, err := goproxmoxapi.New(GetProxmoxAccess())
  if err != nil {
    t.Log(c, err)
    t.Error()
  }
  // Test Token refresh
  if err := c.RefreshToken(); err != nil {
    t.Log(c, err)
    t.Error()
  }
}

func TestFail_New_WrongPass(t *testing.T) {
  t.Parallel()
  username, _, realm, host := GetProxmoxAccess()
  _, err := goproxmoxapi.New(username, "wrong_password", realm, host)
  if err == nil {
    t.Log(err)
    t.Error()
  }
}

func TestFail_New_WrongServer(t *testing.T) {
  if os.Getenv("LONG_RUN_TEST") == "" {
    t.Skip("skipping long running test in console (env variable LONG_RUN_TEST not defned)")
  }
  t.Parallel()
  username, password, realm, _ := GetProxmoxAccess()
  _, err := goproxmoxapi.New(username, password, realm, "10.255.0.6")
  if err == nil {
    t.Log(err)
    t.FailNow()
  }
}
