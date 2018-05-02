# Final Technical Report

Jesse Cochran

April 28th, 2018

## Abstract

The goal of this project was to build a geofencing microservice that would allow clients to query the service and determine if a point is contained within a geofence. These types of services are common in industry, although are generally implemented as proprietary systems. The service built for this project is one service in a larger architecture for a location-based marketing application. The location-based functionality will allow for an accurate targeting of customers to help provide a personalized user experience. The geofencing service was built using Golang and hosted in a Docker container on a Digital Ocean droplet. A basic React application provides an easy to use interface to demonstrate the functionality of the service. This app allows the user to draw polygons using the Google Maps Application Programmable Interface (API) and click inside and outside of the polygons. Upon clicking, the application sends a Hyper Transfer Text Protocol (HTTP) request to the service and issues a notification on the screen to alert the results. If the click was outside of the polygon(s), then the user will be alerted; if it is inside the polygon(s), then the user will be alerted to which polygons the click is contained. 


Keywords: Geofence, Point-In-Polygon, Microservice

## Table of Contents

1. [Introduction](#introduction)  

    1. [Problem](#problem)  

    2. [Objectives](#objectives)  

    3. [Potential Users](#potential-users)  

2. [Project Overview](#project-overview)  

3. [Design, Development and Test](#design,-development-and-test)  

    1. [Design](#design)  

    2. [Development](#development)  

    3. [Test](#test)  

4. [Results](#results)  

5. [Conclusion](#conclusion)  

6. [References](#references)  
    

## Introduction and Project Overview

In our ever increasingly mobile first world, applications with location awareness have become more popular and necessary. Whether it is a coupon manufacturing service or a ride-sharing application, they all use location-based services for improving the customer experience.
Within this technology trend, an application of geofences has arisen. A geofence is a part of a software program 
that uses GPS, WiFi or RFID to define geographical boundaries. With pushes in technology towards augmented reality, self-driving 
cars and IoT the need for location aware devices is becoming increasingly important.


### Problem

Applications that almost everyone uses on a daily basis such as Uber, Google Maps, Waze, and Yelp use some type of geofencing in their product. This is a core part of their product, so these implementations are very valuable and not open to the public. Most of the information regarding similar services are in the form of high-level overviews or brief engineering blogs describing their implementation of a geofencing service. 

While there are geolocation applications you can use as a service, they are costly and you have to forfeit your customers data to a third-party. Additionally, with a high number of users the service will need to process tens of thousands of transactions per second. With such high request volumes, outsourcing this responsibility to a third party would again prove costly. With the limited processing power on many of our devices, it is not feasible to use them to determine their location in relation to geofences. To attain a satisfactory user experience, secure the user’s data, and achieve reasonable response times this computation must be offloaded onto a server. The geofencing microservice built for this project will provide a solution to all of these issues.

### Background

Point-in-Polygon detection algorithms are common methods used to implement geofences. There are a wide array of point-polygon algorithm which all vary in complexity and efficiency. A ray-casting is one of the oldest and most well-known algorithms for detecting if a point is contained within a polygon. 

Another version is to use geohashes to recursively fill the polygon with geohashes and then hash a point to determine if it is contained in one of the hashes. To recursively fill the set you start with an empty fence. Then if the hash is completely contained within the fence completely then it is added to the set. If the hash intersects with the geofence, we don’t add it but recurse with the hashes of its children.

Within the large variety of algorithms used for geofencing, the most efficient ones all share some commonalities. Previous studies have found many use R-Trees to organize their Minimum Bounding Rectangle(MBR) of polygons for the filtering stage [8]. 

An R-Tree is a spatial data structure based on a B-Tree that is used for spatial indexing methods. In the two-dimensional case of geofencing, the MBR is a simple bounding box (bbox) defined by a minimum and maximum coordinate. The check to determine if one object’s bbox is contained in another is a constant time operation.  Figure 1 below shows a good representation of an R-Tree.

![r-tree explanation](https://user-images.githubusercontent.com/13584530/39414423-5891e842-4c05-11e8-9f88-0a3dc050e339.png)  
FIGURE 1. R-TREE DATA STRUCTURE

The average search time complexity for an R-Tree is O(log Mn) where M is the defined constant of the maximum number of children a node can have.

The other common spatial data structure used in the fastest point-in-polygon algorithms is a QuadTree. A QuadTree is a specialization of a generic kd-tree for 2-dimensional indexing. You take a flat project of a surface and divide the surface into quarters, generally called cells. Figure 2 below shows an example of QuadTree generation.

![bingmapquad](https://user-images.githubusercontent.com/13584530/39414425-5a95f818-4c05-11e8-93dd-cd7758246207.jpeg)  
FIGURE 1. QUADTREE EXAMPLE

QuadTrees are used in popular mapping applications such as Google Maps and Bing Maps. Google Maps uses a S2 algorithm, which is a projection of the Earth’s sphere using cube mapping so each cell has a uniform surface area. The cells are arranged using a Hilbert Curve to conserve spatial locality in the cell label. A Hilbert Curve is a space filling curve that allows the range to cover the entire n-dimensional space. 

While R-Tree implementations often have a higher maximum throughput, heavy update activity of the geospatial data decreases the performance of the R-Tree implementations, but heavy updating of the Quad Tree will have no impact on performance [7].

A microservice is one service in part of a microservice architecture that structures an application as a group of loosely coupled services which generally implement business capabilities. Benefits of this architecture style allows for scalability, flexibility, and portability[4]. The microservice designed and implemented for this project is just one service that is part of a larger system, as can be seen in the figure below.


![microservicecapstone](https://user-images.githubusercontent.com/13584530/39459534-30f312dc-4cca-11e8-9dbc-da9541c494cd.png)  


### Objectives

The idea of this project is to create a functionable and scalable solution to geofencing. Implementing a geofence requires lookups using CPU-intensive point-in-polygon algorithms in order to determine if an object exist in a geofence. The algorithms used in this project were not the optimal solutions for point-in-polygon detection. 

A microservice architecture is an architecture style that is structured a collection of loosely coupled services that generally implement some type of business capabilities [2].  This project is intended to provide a microservice that allows a client to query the service and determine if a point is contained within a polygon or geofence.  The microservice architecture would handle all server-side components of an augmented reality based marketing application. 

There are many libraries that help with location data and implementation of common data structures that will be useful in the implementation of this service.


In the project features were ranked as A, B, and C using a priority system. There were three levels of priority, with A being essential, B being want to implement, and C being extensions or extras. They are outlined in the figure below.

![faeatsasa](https://user-images.githubusercontent.com/13584530/39502959-23fd59ee-4d91-11e8-81c0-7df6190d0762.png)  

FIGURE 4. FEATURE TABLE



### Potential Users

While this microservice could be used alone, the functionality of the service fit better into a larger system that uses the service for geolocation purposes. 


## Design, Development and Test

The development of this project was performed using 1 week sprints. Each sprint we planned feature(s) that were to be completed by the end of each sprint. Each of these features should be a testable unit of code. 

### Design

The design of this microservice can be seen in the diagram below. This microservice is designed to be part of a larger system of microservices, in which each is decoupled from the other and has its own functionality.

Golang was chosen as the development language for this microservice for a variety of reasons:


* Point in polygon lookups require CPU-intensive algorithms. Golang is a systems language that is designed to be fast and efficient.
* Low latency and high throughput. The service needs to be able to handle thousands of request, with reach request taking less than 100 milliseconds.
* Concurrent Design. This service must constantly refresh in-memory geofences in the background. Background refreshing can tie up the CPU and slow query response time. Goroutines can be executed on multiple cores and allow the service to run background queries in parallel with foreground queries.


In a true microservice the functionality of the user and client would be separated into different services, but due to time constraints and project scope they were coupled with the geofencing functionality as one microservice.

### Development

One the service had been designed, the development sections were split into three main categories: Users, Clients, and Polygons. The polygon is the geofence, the client “owns” a polygon, and the service is used to determine if the user is inside of one of the client’s polygons. The service was built one endpoint at a time. The user endpoints were implemented first. Each of the endpoints below was implemented and tested:

* getPosition - GET request handler that returns the user’s position
* savePosition - POST request handler that would save the user’s position in the database for later reference
* deletePosition - DELETE request handler the would remove a position(s) from the user’s database

After the user endpoints receive the request, they need to then make a database transaction. The go standard library contains an sql library that was used to help complete these database transaction. After decoding any information in the http request, an SQL statement was then formed and an made to an AWS PostGIS database. 

Next the following client endpoints were implemented and tested:

* getClient - GET request handler to returns the clients
* postClient - POST request handler to save a new client to the client
* deleteClient - DELETE request handler to remove a current client from the database

Once the client endpoints received the request, the database transactions were handled just as with the user database. 

Next the polygon endpoints were implemented:

* getPolygons - GET request handler that retrieves specific polygons in the database
* savePolygon - POST request handler that saves a polygon to the database
* deletePolygon - DELETE request handler that deletes a polygon from the database
* checkGeofence - POST request to determine if a point is within an existing polygon
* checkPolygon - POST request to determine if a polygon intersects with an existing polygon

**NOTE**: The checkPolygon endpoint is functional, but the implementation and response of the endpoint is not fully functional.

The polygon database interactions were more complicated due to the use of the PostGIS extension of the Postgres database. PostGIS SQL queries are more complex and use a completely different type of syntax for manipulating geometrical data within the database. 


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

As previously mentioned, project features were ranked as A, B, and C using a priority system. There were three levels of priority, with A being essential, B being want to implement, and C being extensions or extras.

![featuressss](https://user-images.githubusercontent.com/13584530/39502875-a7f57c5a-4d90-11e8-9d39-890a188b9935.png)  
FIGURE 5. FEATURE RESULTS



## Conclusion

After refactoring my code to allow for proper unit testing, one can see the benefits of taking a Test Driven Development (TDD) approach. In future project, I believe using a TDD approach to development would improve design decisions and save time over the course of the project.

Future development of this microservice would involve removing the other functionality such as interacting with the user and client database, implementing more sophisticated point-in-polygon algorithms, and performing performance testing on the microservice.


The algorithm implemented in this project is not the most efficient solution. 


## References

[1] “Simple example of an R-tree for 2D rectangles” Available: https://en.wikipedia.org/wiki/R-tree [April 27, 2018]

[2] Available: https://msdn.microsoft.com/en-us/library/bb259689.aspx [April 27, 2018]

[3] “Microservice Architecture pattern” Available: http://microservices.io/patterns/microservices.html [April 28. 2018]

[4] https://icpe.spec.org/icpe_proceedings/2017/companion/p223.pdf
[8] S. Tang, Y. Yu, R. Zimmerman, S. Obana. “39 Efficient Geo-fencing via Hybrid Hashing: A Combination of Bucket Selection and In-bucket Binary Search.” Internet: http://research.nii.ac.jp/~yiyu/GeoFence-20150512-FirstLook.pdf, [April 26, 2018]



[7] M. M. Sardadi, M. S. bin Mohd Rahim, Z. Jupri, and D. bin Daman. “Choosing R-tree or Quadtree Spatial Data Indexing in One Oracle Spatial Database System to Make Faster Showing Geographical Map in Mobile Geographical Information System Technology.” Internet: https://pdfs.semanticscholar.org/c86e/a522e7872c44359b00a4102d16e72bbed891.pdf, [April 26, 2018]. 


