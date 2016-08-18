# Neighborhood Manager
## Service Registry
[![Build Status](https://travis-ci.org/RackHD/NeighborhoodManager.svg?branch=master)](https://travis-ci.org/RackHD/NeighborhoodManager)
The [Service Registry] is a utility that will sit on the network and listen for services that advertise themselves. The list of services a user wants to know about can be configured. When an advertisement is seen from a matching service, the Service Registry collects information about the service, and registers it with a storage backend of the users choosing. Other applications can access this storage for aggregated information about all services on a network. 

[Service Registry]: https://github.com/RackHD/NeighborhoodManager/tree/master/registry
