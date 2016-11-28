package goproxmoxapi_test

import (
  "testing"
  "github.com/isindir/goproxmoxapi"
)

func TestStorageAPI(t *testing.T) {
  t.Parallel()
	c, err := goproxmoxapi.New("root", "P@ssw0rd", "pam", "10.255.0.5")
  if err != nil {
    t.Log(c)
    t.Error(err)
  }

  // test that we can create storage
  err = (goproxmoxapi.Storage{
                               Storage: "backuptst",
                               Type: "dir",
                               Content: "backup",
															 Path: "/var/lib/backuptst",
                               MkDir: "1",
                               MaxFiles: "5",
                             }).CreateStorage(c)
  if err != nil {
    t.Error(err)
  }

  // test that we can fetch new storage
	st1, err := (goproxmoxapi.Storage{ Storage: "backuptst" }).GetStorage(c)
  if err != nil {
    t.Error(err)
  }
  // Tes we can update it
  st1.MaxFiles = "6"
  err = st1.UpdateStorage(c)
  if err != nil {
    t.Error(err)
  }
	st2, err := (goproxmoxapi.Storage{ Storage: "backuptst" }).GetStorage(c)
  if err != nil {
    t.Error(err)
  }
  if st2.MaxFiles != "6" {
    t.Error("Update of the storage failed")
  }

  // test that we can fetch all instantiated storages
	_, err = goproxmoxapi.GetAllStorages(c)
  if err != nil {
    t.Error(err)
  }

  // test that we can delete new storage
	err = (goproxmoxapi.Storage{ Storage: "backuptst" }).DeleteStorage(c)
  if err != nil {
    t.Error(err)
  }

  // test that we can fetch existing storages with no issues
	_, err = (goproxmoxapi.Storage{ Storage: "local" }).GetStorage(c)
  if err != nil {
    t.Error(err)
  }
	_, err = (goproxmoxapi.Storage{ Storage: "local-lvm" }).GetStorage(c)
  if err != nil {
    t.Error(err)
  }
  //TODO test CRUD for all types of storages which is possible to test
}
