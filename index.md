---
layout: post
title: Neighborhood Manager
---

====================

[![Build Status](https://travis-ci.org/RackHD/neighborhood-manager.svg?branch=master)](https://travis-ci.org/RackHD/neighborhood-manager)
[![Coverage Status](https://coveralls.io/repos/github/RackHD/neighborhood-manager/badge.svg?branch=master)](https://coveralls.io/github/RackHD/neighborhood-manager?branch=master)
##### Registry
[![GoDoc](https://godoc.org/github.com/RackHD/neighborhood-manager/registry?status.svg)](https://godoc.org/github.com/RackHD/neighborhood-manager/registry)
##### Libreg
[![GoDoc](https://godoc.org/github.com/RackHD/neighborhood-manager/libreg?status.svg)](https://godoc.org/github.com/RackHD/neighborhood-manager/libreg)

## Service Registry
The [Service Registry] is a utility that will listen on the network for services that advertise themselves. The specific list of services a user would like to be notified about is configurable. When an advertisement is seen on the network from a configured service, the Service Registry collects information about the service, and registers the details with a storage backend of the user's choosing. Other applications can access this storage for aggregated information about all services available on a given network.

## RackHD Proxy
The [RackHD Proxy] is a utility that acts as a proxy to the RackHD API.

[Service Registry]: https://github.com/RackHD/neighborhood-manager/tree/master/registry
[RackHD Proxy]: https://github.com/RackHD/neighborhood-manager/tree/master/rackhd

Contribute
----------

Neighborhood Manager is a collection of libraries and applications housed at https://github.com/RackHD/neighborhood-manager. The code for Neighborhood Manager is written in Golang and makes use of Makefiles. It is available under the Apache 2.0 license (or compatible sublicences for library dependencies).

Code and bug submissions are handled on GitHub using the Issues tab for this repository above.

Community
---------

We also have a #InfraEnablers Slack channel: You can get an invite by requesting one at http://community.emccode.com.

Documentation
-------------

You can find documentation for Registry here: [![GoDoc](https://godoc.org/github.com/RackHD/neighborhood-manager/registry?status.svg)](https://godoc.org/github.com/RackHD/neighborhood-manager/registry)

You can find documentation for Libreg here: [![GoDoc](https://godoc.org/github.com/RackHD/neighborhood-manager/libreg?status.svg)](https://godoc.org/github.com/RackHD/neighborhood-manager/libreg)

Licensing
---------

Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

RackHD is a Trademark of EMC Corporation

Support
-------

Please file bugs and issues at the GitHub issues page. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
