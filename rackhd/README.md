# Neighborhood-Manager RackHD Proxy

The proxy is a utility that exposes a http endpoint on the network for api calls to hit. It routes those calls to all RackHD instances that are registered in the backend store (consul).  It then takes all those responses and aggregates them together and serves them back to the initial caller as one completed http request.

To view the RackHD API that is currently supported click here => [RackHD API].


## Background Information
This section details how to set up a production environment for the proxy. If you are interested in a quick demo of the functionality, skip to the Try It Out section to start everything up quickly in containers.

### Backend
Some type of storage backend is needed for the proxy to lookup endpoint nodes to forward calls to. At this time, the only backend supported is [Consul]. The proxy also supports query string inputs for direct calls to known endpoints (see Configuration below).

The proxy assumes the user already has a consul environment setup.  Best practice is to run a consul agent-client on the same host as the proxy.  The proxy is configured to talk to this agent-client through a flag (see Configuration below).

#### Consul Server
To create a cluster of Consul servers, reference this article on [configuring consul], and see the section **Creating the Bootstrap Configuration**, then follow into the next section **Creating the Regular Server Configuration**

#### Consul Client
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

## Accessing Consul Registry
To retrieve a list of services from consul one can use curl (or a client like Postman) to hit the consul API.

This call: Lists services in a given DC:
`curl http://address:port/v1/catalog/services` where address:port is the address and port that the consul agent-client is listening on.

## Building
Download the [source] from GitHub  
`go get -u github.com/rackhd/neighborhood-manager`  
`cd rackhd`

Build the dependencies  
`make deps`

Build the proxy  
`make build`

## Running
After building, the proxy binary (named `rackhd`) will be in the `rackhd/bin/` folder of from the source directory.

Move to that directory and run the proxy.  
```
cd /rackhd/bin
./rackhd
```

This will likely error as no configuration parameters have been specified.  See Configuration below for config flags


## Configuration

The proxy has 4 configuration variables:

`-proxy-address=http://address:port` (Default is `http://0.0.0.0:10001`)

Sets the address and port that the proxy is bound to.

`-backend-address=address:port` (Default is `127.0.0.1:8500`)

Sets the proxy to connect to the backend store at address:port.

`-service-name=some-string` (Default is `RackHD-service:api:2.0`)

To change this variable see the Accessing Consul Registry section above.  

`-datacenter=some-string`  (Default is `dc1`)

Comes from the consul server setup and is the datacenter of the consul cluster.

These flags are passed at the time of starting the binary in a format as such:

`./rackhd -proxy-address=http://10.10.10.1:10001 -datacenter=My-Datacenter`



[configuring consul]: https://www.digitalocean.com/community/tutorials/how-to-configure-consul-in-a-production-environment-on-ubuntu-14-04
[RackHD API]: http://rackhd.readthedocs.io/en/latest/rackhd/index.html
