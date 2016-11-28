# API to implement

* [+] access/domains
* [+] access/groups 
* [~] access/roles       roleid->PUT:append[optional]
* [~] access/users       userid->PUT:append[optional] , Users->GET:filter[optional]
* [ ] access/acl 
* [+] access/password 
* [+] access/ticket 
* [~] cluster
 * [ ] backup
 * [+] log
 * [~] nextid            optional operand to GET to check availability of a specific Id (looks like it's not possible)
 * [ ] options
 * [+] resources         optional operand to GET   : type (enum): vm|storage|node
 * [+] status
 * [+] tasks
 * [ ] ha
  * [ ] ha/groups
  * [ ] ha/resources
  * [ ] ha/status
 * [ ] firewall
  * [ ] firewall/aliases
  * [ ] firewall/groups
  * [ ] firewall/ipset
  * [ ] firewall/rules
  * [ ] firewall/macros
  * [ ] firewall/options
  * [ ] firewall/refs
* [ ] nodes
* [+] pools
* [~] storage            more tesing needed
* [+] version

* [ ] Create proxmox vagrant box, which can be used for testing and push it to public repository
