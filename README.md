# Neighborhood Manager
[![Build Status](https://travis-ci.org/RackHD/neighborhood-manager.svg?branch=master)](https://travis-ci.org/RackHD/neighborhood-manager)  

## Service Registry
The [Service Registry] is a utility that will sit on the network and listen for services that advertise themselves. The list of services a user wants to know about can be configured. When an advertisement is seen from a matching service, the Service Registry collects information about the service, and registers it with a storage backend of the users choosing. Other applications can access this storage for aggregated information about all services on a network.

## RackHD Proxy
The [RackHD Proxy] is a utility that acts as a proxy to the RackHD API.


[Service Registry]: https://github.com/RackHD/neighborhood-manager/tree/master/registry
[RackHD Proxy]: https://github.com/RackHD/neighborhood-manager/tree/master/rackhd
