# Final Technical Report

Jesse Cochran

April 28th, 2018

## Abstract

This project involves building a geofencing microservice, with the goal of allowing clients to query the service and determine if a point is contained within a polygon/geofence. These types of services are common in industry, although are generally proprietary systems. 

Keywords: Geofence, Point-In-Polygon, Microservice

## Table of Contents

## Introduction

In our ever increasingly mobile first world, applications with location awareness have become increasingly popular and useful. 
Within this technology trend, an application of geofences has arisen. A geofence is a part of a software program 
that uses GPS, WiFi or RFID to define geographical boundaries. With pushes in technology towards augmented reality, self-driving 
cars and IoT the need for location aware machines is becoming increasingly necessary.

This project is intended to provide a microservice that allows a client to query the service and determine if a point is contained within a polygon or geofence. A microservice architecture is an architecture style that is structured a collection of loosely coupled services that generally implement some type of business capabilities [2]. 

Point-in-Polygon detection algorithms are common tools used to implement geofences. 

## Project Overview



## Design, Development and Test

### Design


The design of this microservice can be seen in the diagram below. This microservice is designed to be part of a larger system of microservices, in which each is decoupled from the other and has its own functionality.


### Development



### Test

Testing and building was automated using TravisCI.

My initial design of the http handler methods didn’t allow for easy integration testing. The first implementation would have required setting up a temporary http server with proper environment variables to test the handlers. Given most of my testing had been automated through TravisCI, this wasn’t easily feasible. Instead I refactored the database dependency out of my handlers using and interface. This allowed me to decouple my handlers from my database interaction and create a mock database connections for my testing.

Testing the database required using a library for mock databases. I was unable to find any libraries for a mock PostGIS database. Due to limited time, I was unable to build my own version of a mock database for testing the PostGIS. Given the model package contains no functionality there are no unit test for this package.

## Results

A rudimentary service was created at first with Go with a single endpoint was created and tested locally. Once the basic app had been tested locally, it was moved to a DigitalOcean droplet (server). Docker was used to help facilitate easy deployment of new service iterations. Using a makefile and a few make commands, one is able to load up a new docker image and run it on the droplet. 

To help demonstrate the working API, a bootstrapped React App was created. Within this React App some basic http request were created and the results displayed on the UI to help test the functionality of the service. 

The next feature implemented was the addition of database interactions within the service. A Postgres database hosted on Amazon Web Services (AWS) was used to host the database. A database package was created to handle interactions between the service and database. In order to allow the client to retrieve data from the database, endpoints were created that allowed the client to perform GET, POST, and DELETE http request. 

As each feature was added to the service, a demonstration of this functionality was added to the React App. 

In order to save polygons in the database the PostGIS extension was needed. This extension allows one to save geometrical shapes in the database. 


## Conclusion

The algorithm implemented in this project is not the most efficient solution. 

## References

[1] Geofencing definition

[2] http://microservices.io

## Appendices
