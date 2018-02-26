package goproxmoxapi

// Test Proxmox VE instance configuration
// as a singleton according to http://marcio.io/2015/07/singleton-pattern-in-go/

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// ProxmoxConfig represents all needed login information
type ProxmoxConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Realm    string `json:"realm"`
	Node     string `json:"node"`
}

var instance *ProxmoxConfig
var once sync.Once

// GetProxmoxConfigInstance loads default config
func GetProxmoxConfigInstance() *ProxmoxConfig {
	once.Do(func() {
		content, err := ioutil.ReadFile("testconfig.json")

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		var config ProxmoxConfig
		json.Unmarshal(content, &config)
		instance = &config
	})
	return instance
}

// GetProxmoxAccess gets working proxmox ve access information
func GetProxmoxAccess() (string, string, string, string) {
	i := GetProxmoxConfigInstance()
	return i.User, i.Password, i.Realm, i.Host
}

// GetProxmoxNode return the defined test node
func GetProxmoxNode() string {
	i := GetProxmoxConfigInstance()
	return i.Node
}

// GetProxmoxRealm return the defined test realm
func GetProxmoxRealm() string {
	i := GetProxmoxConfigInstance()
	return i.Realm
}
