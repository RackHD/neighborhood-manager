[![Build Status](https://travis-ci.org/RackHD/neighborhood-manager.svg?branch=master)](https://travis-ci.org/RackHD/neighborhood-manager)


# Neighborhood Manager Service Registry
The Service Registry is a utility that will sit on the network and listen for services that advertise themselves. The list of services a user wants to know about can be configured. When an advertisement is seen from a matching service, the Service Registry collects information about the service, and registers it with a storage backend of the users choosing. Other applications can access this storage for aggregated information about all services on a network. 


## Background Information
This section details how to set up a production environment for the Service Registry. If you are interested in a quick demo of the functionality, skip to the Try It Out section to start everything up quickly in containers. 

### Backend
Some type of storage backend is needed for the Service Registry to store information. At this time, the only backend supported is [Consul].  

The best practice for using Consul is to set up a cluster of at least three servers on different hosts, and run a client on the same host that the Service Registry will run on. This allows uninterrupted service if one of the servers in the cluster dies, and is more robust than running a server locally on the same host as the Service Registry. 

#### Server
To create a cluster of Consul servers, reference this article on [configuring consul], and see the section **Creating the Bootstrap Configuration**, then follow into the next section **Creating the Regular Server Configuration**

#### Client
When a cluster is in place, a consul client must be started on the same host that Service Registry will run on.  The most extensible way to do this is to put all parameters in a config file, such as `/etc/consul.d/client/config.json`. 
```
{
    "server":false,
    "datacenter": "dc1",
    "data_dir": "/var/consul",
    "log_level": "DEBUG",
    "enable_syslog": true,
    "start_join": ["192.168.1.1", "192.168.1.2", "192.168.1.3"],
    "bind_addr": "192.168.1.4"
}
```
Where the three addresses in the `start_join` field are the addresses of the three clustered servers, and the `bind_addr` field is the address of the host machine that this client will run on.  

The consul client can then be started with `consul agent -config-dir /etc/consul.d/client/config.json`

## Configuration of URN list
The Service Registry listens for services advertised by [SSDP] messages. These messages contain a URN that identifies the service, and an IP address where the service can be found.  

To configure the list of services that the Service Registry should listen for, open the `registry.json` file in the source code.
The schema for the configuration is:
```
{
    "ssdp":{
        "tags":{
            "Service-Tag":[
                "ServiceUrn",
                "ServiceUrn2"
            ],
            "ServiceTag2":[
                "ServiceUrn3",
                "ServiceUrn4"
            ]
        }
    }
}
```
In this file, the `ServiceTag` and `ServiceTag2` identifiers under the `tags` field describe the services, and the `ServiceUrn` listings within them are the URNs found in matching SSDP messages.  

The `registry.json` file found in the source code is currently configured to listen to advertisements from [Inservice-Agent] and [RackHD]

## Building
Download the [source] from GitHub  
`go get -u github.com/rackhd/neighborhood-manager`  
`cd registry`

Build the dependencies  
`make deps`

Build the Service Registry  
`make build`

## Running
After building, the Service Registry binary (named `registry`) will be in the `bin/` folder of the source directory. To run it from there, copy `registry.json` to that location.  
`cp registry.json ../bin/`  

Move to that directory and run the Service Registry.  
```
cd ../bin
./registry
``` 


## Try It Out
The steps in this section will guide you through getting a test/dev environment up and running. Use these steps if you want to see the Service Registry in action. 

### Prerequisites
* [Git]  
* [Docker]  
* [Docker-compose]  

### Starting the environment
1. Change to a directory where you can clone the source code
2. `git clone https://github.com/RackHD/neighborhood-manager.git`
3. `cd registry`
4. `make run-reg` (Note that this will take some time when running for the first time. Docker has to pull several images from the internet. If any of the downloads fail, simply re-do the `make run-reg` command)

### Interacting with the environment
You now have four docker containers running: 
* Service Registry
* Consul Client
* Consul Server
* SSDP Spoofer

The SSDP Spoofer sends out dummy advertisement messages made to look like Inservice-Agent and RackHD. They are received by the Service Registry, passed to the Consul client, and sent to the Consul server for registration. You can interact with the Consul client in the same way the Service Registry does, to retrieve information that has been stored.

1. Open a new terminal prompt
2. Change to the `neighborhood-manager` source directory
3. `make consul-shell` 

Now you can use the [Consul Catalog API] to interact with the backend storage. For example, to retrieve all services that have been registered: 
```
root@c1208c34725f:/go/src/github.com/RackHD/neighborhood-manager# curl -s http://consulclient:8500/v1/catalog/services | python -mjson.tool
{
    "Inservice-service:agent:0.1": [],
    "Inservice-service:catalog-compute:0.1": [],
    "Inservice-service:lldp:0.1": [],
    "RackHD-device:on-http:1": [],
    "RackHD-service:api:1.1": [],
    "RackHD-service:api:2.0": [],
    "RackHD-service:redfish-rest:1.0": [],
    "consul": []
}
```
Notice the URL to which you are sending the request. The Docker containers have been configured in a way such that you can reference the container running the Consul client by name, even though the IP address of the container will change between runs. Port 8500 is the default port for HTTP communications in Consul. 


As another example, to see all nodes that are offering the RackHD 2.0 API service:
```
root@50283b1b6a03:/go/src/github.com/RackHD/neighborhood-manager# curl http://consulclient:8500/v1/catalog/service/"RackHD-service:api:2.0" | python -mjson.tool
[
    {
        "Address": "192.168.1.1",
        "CreateIndex": 11,
        "ModifyIndex": 11,
        "Node": "228617159241066160",
        "ServiceAddress": "192.168.1.1",
        "ServiceEnableTagOverride": false,
        "ServiceID": "RackHD-service:api:2.0",
        "ServiceName": "RackHD-service:api:2.0",
        "ServicePort": 65535,
        "ServiceTags": []
    }
]
```
When you are done, run `exit` to stop the shell container, and `docker-compose kill` to stop the four containers running the Service Registry and its environment. 



## Dependency Management  
### Adding a library dependency  
When a new library is used for the first time, it must be added to Glide. Open the file `glide.yaml` and add a new line under the `import` heading:
```
import:
- package: github.com/king-jam/gossdp
```
replacing the example GitHub link with the path that `go get` needs to grab the library. 

When this is done, save `glide.yaml` and run `glide update` to copy the dependent source code to the `vendor/` folder.  

After this, the steps in the next section should be done to lock in the specific version of the dependent library. 

### Updating a library dependency
When a dependent library has an updated commit that is desired (such as a big fix, added feature, etc), the library's entry in `glide.lock` should be updated. Open the file and find its entry, such as 
```
- name: github.com/hashicorp/consul
  version: 36dc9201f2e006d4b5db1f0446b17357811297bf
```
Replace the existing commit hash with the hash of the new desired commit.  
Save and close the file, then run `glide install`  

[Git]: https://git-scm.com/
[Docker]: https://docs.docker.com/engine/installation/
[Docker-compose]: https://docs.docker.com/compose/install/
[Consul Catalog API]: https://www.consul.io/docs/agent/http/catalog.html
[Consul]: https://www.consul.io/downloads.html
[joined together]: https://www.consul.io/docs/guides/servers.html
[agent]: https://www.consul.io/docs/agent/basics.html
[SSDP]: https://en.wikipedia.org/wiki/Simple_Service_Discovery_Protocol
[Inservice-Agent]: https://github.com/RackHD/InService
[RackHD]: https://github.com/RackHD
[configuring consul]: https://www.digitalocean.com/community/tutorials/how-to-configure-consul-in-a-production-environment-on-ubuntu-14-04
[source]: https://github.com/RackHD/neighborhood-manager
