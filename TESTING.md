# Testing Overview

You need a working Proxmox VE in version 5.x, just because it is the current version.
In this Proxmox VE instance, you need an user with admin priviledges, which login data
has to go in to the configuration file `testconfig.json`, for an example, please have
a look at `testconfig.json.example`.

You need the following:

* one container present with ID `100`
* Alpine Linux container template `alpine-3.5-default_20170504_amd64.tar.xz`
* two groups `gp1` and `gp2`.
