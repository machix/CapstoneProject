# Final Technical Report

Jesse Cochran

April 28th, 2018

## Abstract

The goal of this project was to build a geofencing microservice, with the goal of allowing clients to query the service and determine if a point is contained within a polygon/geofence. These types of services are common in industry, although are generally proprietary systems. 

Throughout the rest of this report the term polygon and geofence will be used interchangeably. 


Keywords: Geofence, Point-In-Polygon, Microservice

## Table of Contents

[Introduction](#introduction)
	[Problem](#problem)
	[Objectives](#objectives)
	[Potential Users](#potential-users)
[Project Overview](#project-overview)
[Design, Development and Test](#design,-development-and-test)
	[Design](#design)
	[Development](#development)
	[Test](#test)
[Results](#results)
[Conclusion](#conclusion)
[References](#references)
	

## Introduction

In our ever increasingly mobile first world, applications with location awareness have become increasingly popular and necessary. Whether it is a coupon manufacturing service or a ride-sharing application, they all use location based services for improving the customer experience.
Within this technology trend, an application of geofences has arisen. A geofence is a part of a software program 
that uses GPS, WiFi or RFID to define geographical boundaries. With pushes in technology towards augmented reality, self-driving 
cars and IoT the need for location aware devices is becoming increasingly important

### Problem

Apps that almost everyone uses on a daily basis such as Uber, Google Maps, Waze, Yelp, and the list goes on and on, use some type of geofencing in their product. This is a core part of their product, so these implementation are very valuable and not open to the public. Most of the information regarding similar services are in the form of high-level overviews of their engineering architecture. There are many libraries that help with location data and implementation of common data structures that will be useful in the implementation of this service.

Along with the convenience that mobile applications provide, users also expect relevant information. For example, if a user wants to use a mobile app to find local coupons, they don’t want to see coupons for stores over an hour away. The geolocation services offered as part of the applications allow commerce platforms to provide relevant and more valuable information to their customers. 

While there are geolocation applications you can use as a service, they are costly and you have to forfeit your customers data to a third-party. Additionally, with a high number of users the service will need to process tens of thousands of transactions per second. With such high transaction volumes, outsourcing this responsibility to a third party would again prove costly. 

With the help of open source geospatial libraries, tools and data structures one can build their out geolocation service.  

### Background

A microservice is one service in part of a microservice architecture that structures an application as a group of loosely coupled services which generally implement business capabilities. Benefits of this architecture style allows for scalability, flexibility, and portability[4]. The microservice designed and implemented for this project is just one service that is part of a larger system, as can be seen in the figure below.

Point-in-Polygon detection algorithms are common tools used to implement geofences. 

### Objectives

The idea of this project is to create a functionable and scalable solution to geofencing. Implementing a geofence requires lookups using CPU-intensive point-in-polygon algorithms in order to determine if an object exist in a geofence.

A microservice architecture is an architecture style that is structured a collection of loosely coupled services that generally implement some type of business capabilities [2].  This project is intended to provide a microservice that allows a client to query the service and determine if a point is contained within a polygon or geofence.  The microservice architecture would handle all server-side components of an augmented reality based marketing application. 

### Potential Users

While this microservice could be used alone, the functionality of the service fit better into a larger system that uses the service for geolocation purposes. 



## Project Overview



## Design, Development and Test

The development of this project was performed using 1 week sprints. Each sprint we planned feature(s) that were to be completed by the end of each sprint. Each of these features should be a testable unit of code. 

### Design

The design of this microservice can be seen in the diagram below. This microservice is designed to be part of a larger system of microservices, in which each is decoupled from the other and has its own functionality.

Golang was chosen as the development language for this microservice for a variety of reasons:


* Point in polygon lookups require CPU-intensive algorithms. Golang is a systems language that is designed to be fast and efficient.
* Low latency and high throughput. The service needs to be able to handle thousands of request, with reach request taking less than 100 milliseconds.
* Concurrent Design. This service must constantly refresh in-memory geofences in the background. Background refreshing can tie up the CPU and slow query response time. Goroutines can be executed on multiple cores and allow the service to run background queries in parallel with foreground queries.


### Development

One the service had been designed, the development sections were split into three main categories: Users, Clients, and Polygons. The polygon is the geofence, the client “owns” a polygon, and the service is used to determine if the user is inside of one of the client’s polygons. The service was built one endpoint at a time.Once the endpoint and HTTP handler was implemented, the database interaction was completed. 

### Test

Testing and building was automated using TravisCI. 

My initial design of the http handler methods didn’t allow for easy unit testing. The first implementation would have required setting up a temporary http server with proper environment variables to test the handlers. Given most of my testing had been automated through TravisCI, this wasn’t easily feasible. Instead I refactored the database dependencies out of my handlers using an interface. This allowed me to decouple my handlers from my database interaction and create a mock datastores for testing.

Testing the database required using a library for mock databases. I was unable to find any libraries for a mock PostGIS database. Due to limited time, I was unable to build my own version of a mock database for testing the PostGIS. Given the model package contains no functionality there are no unit test for this package.

Go-carpet was used to determine testing coverage. This metric was only mildly useful for this specific project as some of the packages had no functionality, and as mentioned some of the specific database interactions could not be testing. 

Once each feature was complete integration testing was also performed using calls from the client. The correct implementation was testing by verifying the correct response from the http request. 

## Results

A rudimentary service was created at first with Go with a single endpoint was created and tested locally. Once the basic app had been tested locally, it was moved to a DigitalOcean droplet (server). Docker was used to help facilitate easy deployment of new service iterations. Using a makefile and a few make commands, one is able to load up a new docker image and run it on the droplet. 

To help demonstrate the working API, a bootstrapped React App was created. Within this React App some basic http request were created and the results displayed on the UI to help test the functionality of the service. 

The next feature implemented was the addition of database interactions within the service. A Postgres database hosted on Amazon Web Services (AWS) was used to host the database. A database package was created to handle interactions between the service and database. In order to allow the client to retrieve data from the database, endpoints were created that allowed the client to perform GET, POST, and DELETE http request. 

As each feature was added to the service, a demonstration of this functionality was added to the React App. 

In order to save polygons in the database the PostGIS extension was needed. This extension allows one to save geometrical shapes in the database. 
This was my first large project using Go as the development language. Coming from an OOP background, learning best practices for the language was a valuable learning process.



## Conclusion

After refactoring my code to allow for proper unit testing, one can see the benefits of taking a Test Driven Development (TDD) approach. In future project, I believe using a TDD approach to development would improve design decisions and save time over the course of the project.


The algorithm implemented in this project is not the most efficient solution. 

## References

[1] Geofencing definition

[2] http://microservices.io

[4] https://icpe.spec.org/icpe_proceedings/2017/companion/p223.pdf

## Appendices

