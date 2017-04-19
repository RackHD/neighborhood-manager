Neighborhood Manager
====================

[![Build Status](https://travis-ci.org/RackHD/neighborhood-manager.svg?branch=master)](https://travis-ci.org/RackHD/neighborhood-manager)
[![Coverage Status](https://coveralls.io/repos/github/RackHD/neighborhood-manager/badge.svg?branch=master)](https://coveralls.io/github/RackHD/neighborhood-manager?branch=master)

## Service Registry [![GoDoc](https://godoc.org/github.com/RackHD/neighborhood-manager/registry?status.svg)](https://godoc.org/github.com/RackHD/neighborhood-manager/registry)
The [Service Registry](https://github.com/RackHD/neighborhood-manager/tree/master/registry) is a utility that will listen on the network for services that advertise themselves. The specific list of services a user would like to be notified about is configurable. When an advertisement is seen on the network from a configured service, the Service Registry collects information about the service, and registers the details with a storage backend of the user's choosing. Other applications can access this storage for aggregated information about all services available on a given network.

## RackHD Proxy [![GoDoc](https://godoc.org/github.com/RackHD/neighborhood-manager/rackhd?status.svg)](https://godoc.org/github.com/RackHD/neighborhood-manager/rackhd)
The [RackHD Proxy](https://github.com/RackHD/neighborhood-manager/tree/master/rackhd) is a utility that acts as a proxy to the RackHD API.

## Libreg [![GoDoc](https://godoc.org/github.com/RackHD/neighborhood-manager/libreg?status.svg)](https://godoc.org/github.com/RackHD/neighborhood-manager/libreg)

[Libreg](https://github.com/RackHD/neighborhood-manager/tree/master/libreg) is an abstraction library to interface with a backend datastore.  Currently supports Consul.

Basic Architecture Diagram
------------------

![diagram](https://github.com/RackHD/neighborhood-manager/blob/gh-pages/NM_02OCT2016-1.png?raw=true)

Prerequisites
-------------

This project is written in golang and some commands in the documentation assume a working go installation.

This project also uses docker and for those commands we assume the latest docker version is installed.

Contribute
----------

Neighborhood Manager is a collection of libraries and applications housed at https://github.com/RackHD/neighborhood-manager. The code for Neighborhood Manager is written in Golang and makes use of Makefiles. It is available under the Apache 2.0 license (or compatible sublicences for library dependencies).

Code and bug submissions are handled on GitHub using the Issues tab for this repository above.

Community
---------

We also have a #RackHD Slack channel: You can request an invite at http://community.emccode.com.


Licensing
---------

Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

RackHD is a Trademark of Dell EMC

Support
-------

Please file bugs and issues at the GitHub issues page. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
