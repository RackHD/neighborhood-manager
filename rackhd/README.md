# Neighborhood-Manager RackHD Proxy

## Host System Requirements

You will need to install the below programs on your host system:

1. **Docker** version [1.12] or greater is required.
2. **Docker-Compose** [latest]

Once these two programs are installed a reboot may be needed.

## Installation

    git clone https://github.com/RackHD/neighborhood-manager.git

then 

    cd neighborhood-manager
  

## Deployment

    make run-proxy

## Usage

Some test commands

    curl http://localhost:10001/test

returns that the proxy is alive
 
    curl http://localhost:10001/array

returns an array of objects
 
    curl http://localhost:10001/object

returns an object


[1.12]: https://docs.docker.com/
[latest]: https://docs.docker.com/compose/install/
