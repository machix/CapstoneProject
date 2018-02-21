# Geofencing Service

This service is designed to allow queries of geospatial points and return information regarding the points location relative to other geofences and points within a specific distance.


## Installation

To install this project you will first need to ensure that you have Golang(Go) on your system. If you do not have Go on your system you can [visit the site](https://golang.org/doc/install) to download the most recent version and set up the proper system settings.

Next you will need to ensure that you have Docker installed on your system. If you do not have docker installed you can [visit the docker website](https://docs.docker.com/install/#supported-platforms) for instructions on how to setup docker on your system.

## Usage

This code is primarly meant to be used as a standalone service in a container.

To build the service and save it as ```geo-api-linux``` you can run ```make geo-api-linux```.

To build the docker image for the service ```geo-api-linux``` that was just built, you would run ```make geo-service1.docker```

To remove the latest ```.docker``` built, simply use the ```make clean``` command.

All other ```make``` commands are specific to accesing a private DigitalOcean droplet. You can switch out the values in the makefile with your own private digital ocean droplet IP. This would allow you to transfer your service docker image with ```make rsync``` and you can restart your current docker container running with the new one on your remote droplet using the ```make restart``` command.


