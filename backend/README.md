# Geofencing Service

This service is designed to allow queries of geospatial points and return information regarding the points location relative to other geofences and points within a specific distance. The service communicates with a Postgres and PostGIS database on the backend for storing data.


## Installation

To install this project you will first need to ensure that you have Golang(Go) on your system. If you do not have Go on your system you can [visit the site](https://golang.org/doc/install) to download the most recent version and set up the proper system settings.

Next you will need to ensure that you have Docker installed on your system. If you do not have docker installed you can [visit the docker website](https://docs.docker.com/install/#supported-platforms) for instructions on how to setup docker on your system.

## Usage

This application is primarly meant to be used as a standalone service in a container. In addition to building the project, you will need to setup your own database to connect to the service. The user database can be any type of SQL database, but the client database must be a Postgres database with the PostGIS extensions loaded. This extensions allows for storing of polygon and other geometrical points for efficient data storage and querying.

### Building

The makefile can be sure to streamline the building process:

* To build the service and save it as ```geo-api-linux``` you can run ```make geo-api-linux```.

* To build the docker image for the service ```geo-api-linux``` that was just built, you would run ```make geo-service1.docker```

* To remove the latest ```.docker``` built, simply use the ```make clean``` command.

All other ```make``` commands are specific to accesing a private DigitalOcean droplet. You can switch out the values in the makefile with your own private digital ocean droplet IP. This would allow you to transfer your service docker image with ```make rsync``` and you can restart your current docker container running with the new one on your remote droplet using the ```make restart``` command.

Addtionally, you will need to setup a ```env.list``` file on your droplet that contains environment variables for connecting to the database. The environment variables you will need to set will be DBHOST, DBPORT, DBUSER, DBPASS, and DBNAME where this is the database host, database port, database user, database password, and database name, respectively.

### Testing

Testing for the application includes ```sql-mock``` to install this run the command below:

```go get gopkg.in/DATA-DOG/go-sqlmock.v1```

Once this dependency is installed, ensure you are at the top of the backend directory and run:

```go test ./...```

