# Final Technical Report

Jesse Cochran

April 28th, 2018

## Abstract

The goal of this project was to build a geofencing microservice the would allow clients to query the service and determine if a point is contained within a polygon/geofence. These types of services are common in industry, although are generally implemented as proprietary systems. The service built for this project will be once service in a larger architecture for a location-based marketing application. A geofencing service was build using Golang and hosted in a docker container on a Digital Ocean droplet. To demonstrate the functionality of the service a basic React app was built. This app allows the user to draw polygons using the Google Maps API and then click inside and outside of the polygons. Upon clicking, the app sends a http request to the service and issues a toast notification on the screen to alert the results. If the click was outside of the polygon, then the user will be alerted, and if it is inside the polygon(s), then the user will be alerts to which polygons the click is contained. 


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

In our ever increasingly mobile first world, applications with location awareness have become more popular and necessary. Whether it is a coupon manufacturing service or a ride-sharing application, they all use location based services for improving the customer experience.
Within this technology trend, an application of geofences has arisen. A geofence is a part of a software program 
that uses GPS, WiFi or RFID to define geographical boundaries. With pushes in technology towards augmented reality, self-driving 
cars and IoT the need for location aware devices is becoming increasingly important.

Throughout the rest of this report the term polygon and geofence will be used interchangeably. 

### Problem

Apps that almost everyone uses on a daily basis such as Uber, Google Maps, Waze, Yelp, and the list goes on and on, use some type of geofencing in their product. This is a core part of their product, so these implementation are very valuable and not open to the public. Most of the information regarding similar services are in the form of high-level overviews of their engineering architecture. There are many libraries that help with location data and implementation of common data structures that will be useful in the implementation of this service.

Along with the convenience that mobile applications provide, users also expect relevant information. For example, if a user wants to use a mobile app to find local coupons, they don’t want to see coupons for stores over an hour away. The geolocation services offered as part of the applications allow commerce platforms to provide relevant and more valuable information to their customers. 

While there are geolocation applications you can use as a service, they are costly and you have to forfeit your customers data to a third-party. Additionally, with a high number of users the service will need to process tens of thousands of transactions per second. With such high transaction volumes, outsourcing this responsibility to a third party would again prove costly. 

With the help of open source geospatial libraries, tools and data structures one can build their out geolocation service.  

### Background

A microservice is one service in part of a microservice architecture that structures an application as a group of loosely coupled services which generally implement business capabilities. Benefits of this architecture style allows for scalability, flexibility, and portability[4]. The microservice designed and implemented for this project is just one service that is part of a larger system, as can be seen in the figure below.

Point-in-Polygon detection algorithms are common tools used to implement geofences. The algorithms used in this project were not the optimal solutions for point-in-polygon detection. Within the large variety of algorithms used for geofencing, the most efficient ones all share some commonalities. Previous studies have found they all use R-Trees to organize their Minimum Bounding Rectangle(MBR) of polygons for the filtering stage [8]. 

An R-Tree is a spatial data structure based on a B-Tree that is used for spatial indexing methods. In the two-dimensional case of geofencing, the MBR is a simple bounding box (bbox) defined by a minimum and maximum coordinate. The check to determine if one object’s bbox is contained in another is a constant time operation.  Figure 1 below shows a good representation of an R-Tree.

![r-tree explanation](https://user-images.githubusercontent.com/13584530/39414423-5891e842-4c05-11e8-9f88-0a3dc050e339.png)
FIGURE 1. R-TREE DATA STRUCTURE

The average search time complexity for an R-Tree is O(log Mn) where M is the defined constant of the maximum number of children a node can have.

The other common spatial data structure used in the fastest point-in-polygon algorithms is a QuadTree. A QuadTree is a specialization of a generic kd-tree for 2-dimensional indexing. You take a flat project of a surface and divide the surface into quarters, generally called cells. Figure 2 below shows an example of QuadTree generation.

![bingmapquad](https://user-images.githubusercontent.com/13584530/39414425-5a95f818-4c05-11e8-93dd-cd7758246207.jpeg)
FIGURE 1. QUADTREE EXAMPLE

QuadTrees are used in popular mapping applications such as Google Maps and Bing Maps. Google Maps uses a S2 algorithm, which is a projection of the Earth’s sphere using cube mapping so each cell has a uniform surface area. The cells are arranged using a Hilbert Curve to conserve spatial locality in the cell label. A Hilbert Curve is a space filling curve that allows the range to cover the entire n-dimensional space. 

While R-Tree implementations often have a higher maximum throughput, heavy update activity of the geospatial data decreases the performance of the R-Tree implementations, but heavy updating of the Quad Tree will have no impact on performance [7].




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

[8] S. Tang, Y. Yu, R. Zimmerman, S. Obana. “39 Efficient Geo-fencing via Hybrid Hashing: A Combination of Bucket Selection and In-bucket Binary Search.” Internet: http://research.nii.ac.jp/~yiyu/GeoFence-20150512-FirstLook.pdf, [April 26, 2018]

[7] M. M. Sardadi, M. S. bin Mohd Rahim, Z. Jupri, and D. bin Daman. “Choosing R-tree or Quadtree Spatial Data Indexing in One Oracle Spatial Database System to Make Faster Showing Geographical Map in Mobile Geographical Information System Technology.” Internet: https://pdfs.semanticscholar.org/c86e/a522e7872c44359b00a4102d16e72bbed891.pdf, [April 28, 2018]. 

## Appendices


