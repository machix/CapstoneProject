# Geofencing Microservice - Final Technical Report

Jesse Cochran

April 28th, 2018

## Abstract

The goal of this project was to build a geofencing microservice that would allow clients to query the service and determine if a point is contained within a geofence. These types of services are common in industry, although are generally implemented as proprietary systems. The service built for this project is one service in a larger architecture for a location-based marketing application. The location-based functionality will allow for an accurate targeting of customers to help provide a personalized user experience. The geofencing service was built using [Golang](https://golang.org) and hosted in a [Docker](https://www.docker.com) container on a [Digital Ocean](https://www.digitalocean.com) droplet. A basic React application provides an easy to use interface to demonstrate the functionality of the service. This app allows the user to draw polygons using the Google Maps Application Programmable Interface (API) and click inside and outside of the polygons. Upon clicking, the application sends a Hyper Transfer Text Protocol (HTTP) request to the service and issues a notification on the screen to alert the results. If the click was outside of the polygon(s), then the user will be alerted; if it is inside the polygon(s), then the user will be alerted to which polygons the click is contained. All of the core features of the service outlined in the proposal were completed. [TravisCI](https://travis-ci.com) was used for build and testing automation. Git was used as the version control system and Github was used to host our projects as well as manage our feature sprints. 


#### Keywords

Geofence, Point-In-Polygon, Microservice, Golang, AWS, Postgres, PostGIS, endpoint, http, SQL, interface.

## Table of Contents

1. [Introduction & Project Overview](#introduction-&-project-overview)  

    1. [Problem](#problem)  

    2. [Objectives](#objectives)  

    3. [Potential Users](#potential-users)  

2. [Design, Development and Test](#design,-development-and-test)  

    1. [Design](#design)  

    2. [Development](#development)  

    3. [Test](#test)  

3. [Results](#results)  

4. [Conclusion](#conclusion)  

5. [References](#references)  
    

## Introduction & Project Overview

In our ever increasingly mobile first world, applications with location awareness have become more popular and necessary. Whether it is a coupon manufacturing service or a ride-sharing application, they all use location-based services for improving the customer experience.
Within this technology trend, an application of geofences has arisen. A geofence is a part of a software program that uses Global Positioning Systems (GPS), Wi-Fi or Radio Frequency Identification Devices (RFID) to define geographical boundaries. With pushes in technology towards augmented reality, self-driving cars and IoT the need for location aware devices is becoming increasingly important.

### Problem

Applications that almost everyone uses on a daily basis such as Uber, Google Maps, Waze, and Yelp use some type of geofencing in their product. This is a core part of their product, so these implementations are very valuable and not open to the public. Most of the information regarding similar services are in the form of high-level overviews or brief engineering blogs describing their implementation of a geofencing service. 

While there are geolocation applications you can use as a service, they are costly and you have to forfeit your customers data to a third-party. Additionally, with a high number of users the service will need to process tens of thousands of transactions per second. With such high request volumes, outsourcing this responsibility to a third party would again prove costly. With the limited processing power on many of our devices, it is not feasible to use them to determine their location in relation to geofences. To attain a satisfactory user experience, secure the user’s data, and achieve reasonable response times this computation must be offloaded onto a server. The geofencing microservice built for this project will provide a solution to all of these issues.


### Background

Point-in-Polygon detection algorithms are common methods used to implement geofences. There are a wide array of point-polygon algorithm which all vary in complexity and efficiency. Ray-casting is one of the oldest and most well-known algorithms for detecting if a point is contained within a polygon. Starting from the point of concern, you count the number the number of times a ray, in any fixed direction, intersects with the edges of the polygon. If the point intersects the edges of the polygon an odd number of times, then it is inside the polygon [5]. Unfortunately, this algorithm doesn’t work if the point of concern lies on the edge of the polygon. This was the first algorithm implemented in the project. A visual representation of this process can be seen below.

![300px-ray casting algorithm illustration](https://user-images.githubusercontent.com/13584530/39556319-8d2abf1c-4e4d-11e8-8f9f-513044fb7ad1.png)  
FIGURE 1. RAY-CASTING ALGORITHM [1]


The geohashing algorithm was the next algorithm implemented for this project and is the current implementation in the geofencing microservice. Another version is to use geohashes to recursively fill the polygon with geohashes and then hash a point to determine if it is contained in one of the hashes [8]. A geohash is a way to express or encode a location using a short alphanumeric string. The steps to the algorithm using geohashes is outlined below:
1) To recursively fill the set you start with an empty fence. 
2) If the hash is completely contained within the fence completely then it is added to the set. If the hash intersects with the geofence, it isn’t added, but instead recurse with the hashes of its children. 
3) Once you reach the point in which the hash is equal to the maximum precision that is being checked, you add the hash to the set and stop recursing. 
4) Ultimately, if the hash doesn’t intersect with the geofence, then you don’t add it to the set, and you stop recursing. 
A visual representation of this process can be seen in the figure below. 


![screenshotfromhashfill](https://user-images.githubusercontent.com/13584530/39506363-06c9eb7a-4da6-11e8-8f9a-ed61248c415a.png)  
FIGURE 2. RECURSIVE HASHING [2]

The following algorithms are more optimal geofencing algorithms. The implementations of these algorithms in this project are partially completed or are planned to be implemented in future iterations. Within the large variety of algorithms used for geofencing, the most efficient ones all share some commonalities. Previous research has found many use R-Trees to organize their Minimum Bounding Rectangle(MBR) of polygons for the filtering stage [8]. 

An R-Tree is a spatial data structure based on a B-Tree that is used for spatial indexing methods. In the two-dimensional case of geofencing, the MBR is a simple bounding box (bbox) defined by a minimum and maximum coordinate. The check to determine if one object’s bbox is contained in another is a constant time operation.  Figure 1 below shows a good representation of an R-Tree.

![r-tree explanation](https://user-images.githubusercontent.com/13584530/39414423-5891e842-4c05-11e8-9f88-0a3dc050e339.png)  
FIGURE 3. R-TREE DATA STRUCTURE [3]

The average search time complexity for an R-Tree is O(log Mn) where M is the defined constant of the maximum number of children a node can have.

The other common spatial data structure used in the fastest point-in-polygon algorithms is a QuadTree. A QuadTree is a specialization of a generic kd-tree for 2-dimensional indexing. You take a flat project of a surface and divide the surface into quarters, generally called cells. Figure 2 below shows an example of QuadTree generation.

![bingmapquad](https://user-images.githubusercontent.com/13584530/39414425-5a95f818-4c05-11e8-93dd-cd7758246207.jpeg)  
FIGURE 4. QUADTREE EXAMPLE [4]

QuadTrees are used in popular mapping applications such as Google Maps and Bing Maps. Google Maps uses a S2 algorithm, which is a projection of the Earth’s sphere using cube mapping so each cell has a uniform surface area. The cells are arranged using a Hilbert Curve to conserve spatial locality in the cell label. A Hilbert Curve is a space filling curve that allows the range to cover the entire n-dimensional space. 

While R-Tree implementations often have a higher maximum throughput, heavy update activity of the geospatial data decreases the performance of the R-Tree implementations, but heavy updating of the Quad Tree will have no impact on performance [7].

 This project is intended to provide a microservice that allows a client to query the service and determine if a point is contained within a polygon or geofence. This microservice is one service in part of a [microservice architecture](http://microservices.io/patterns/microservices.html). A microservice architecture is an architecture style that is a structured collection of loosely coupled services that generally implement some type of business capabilities [6]. The microservice architecture would handle all server-side components of an augmented reality based marketing application. Benefits of this architecture style allows for scalability, flexibility, and portability[4].


### Objectives

The idea of this project is to create a functionable and scalable solution to geofencing. Implementing a geofence requires lookups using CPU-intensive point-in-polygon algorithms in order to determine if an object exist in a geofence. The algorithms used in this project were not the optimal solutions for point-in-polygon detection. 

In the project features were ranked as A, B, and C using a priority system. There were three levels of priority, with A being essential, B being want to implement, and C being extensions or extras. They are outlined in the figure below.


<img width="902" alt="screen shot 2018-05-02 at 9 21 11 pm" src="https://user-images.githubusercontent.com/13584530/39556508-cac656aa-4e4e-11e8-9e1d-f9832dd7b42a.png">

FIGURE 5. FEATURE TABLE



### Potential Users

While this microservice could be used alone, the functionality of the service fit better into a larger system that uses the service for geolocation purposes. 

In theory, almost any company that needs to implement geofencing into their application could use this service. As an example, a ride-sharing application such as Lyft would use this in their product. They want to optimize the number of passengers a car is able to service, and optimize the wait time for a rider to get a car. They can use a geofence around a car to determine if a rider is in close enough range to send the driver to pick up the rider. This helps the service be as efficient as possible for both users. This assumes there is only one point in the polygon, but if there were multiple points in the polygon, then choosing the optimal rider would be a different problem a bit outside of the scope of this project. Specifically, this microsevice will be used to handle geolocation operations in a location-based augmented reality application. 


## Design, Development and Test

The development of this project was performed using 1 week sprints. Each sprint, we planned features that were to be completed by the end of each sprint. Each of these features should be a testable unit of code. Github Projects and Issues were used to track each of the features and our current position in each sprint. Git was used as the version control throughout the project. 

### Design

This microservice is designed to be part of a larger system of microservices, in which each is decoupled from the other and has its own functionality. The microservice designed and implemented for this project is just one service that is part of a larger system, as can be seen in the figure below.

![microservicecapstone](https://user-images.githubusercontent.com/13584530/39459534-30f312dc-4cca-11e8-9dbc-da9541c494cd.png)  
FIGURE 6. MICROSERVICE ARCHITECTURE DIAGRAM

I chose to split the microservice into the three main categories User, Client, and Polygon. In a true microservice the functionality of the user and client would be separated into different services. But, due to time constraints and project scope they were coupled with the geofencing functionality as one microservice. The service is split into four main packages:
* model - Holds all of the structs and models of the service. This package contains no functionality.
* database - This package handles all of the database interactions.
* handlers - This package contains all of the handlers and methods for building the router containing the handlers.
* geofence - This package contains the implementation of the geofence algorithm and associated functionality.

Golang was chosen as the development language for this microservice for a variety of reasons:

* Point in polygon lookups require CPU-intensive algorithms. Golang is a systems language that is designed to be fast and efficient.
* Low latency and high throughput. The service needs to be able to handle thousands of requests, with reach request taking less than 100 milliseconds.
* Concurrent Design. This service must constantly refresh in-memory geofences in the background. Background refreshing can tie up the CPU and slow query response time. Goroutines can be executed on multiple cores and allow the service to run background queries in parallel with foreground queries.

Processing request in Go comprise of two main things, a request router and request handler. A router compares incoming request against a list of predefined Universal Resource Locators (URLs) paths, and then calls an associated handler for the path if a match is found. Handlers are responsible for writing response headers and bodies. In this project the library [gorilla/mux](https://github.com/gorilla/mux) was used to implement a request router and dispatcher for matching  incoming request to the appropriate handler.

Eact of the two categories user and client each had three endpoints. One handling a GET request and retrieving the user or client. Another endpoint handling a POST request to create a new user or client. Finally, the last endpoint handled DELETE request and deleted a specified user or client. The polygon category has 5 endpoints. Two of these endpoints are similar to the client and user endpoints, as they handle GET, POST, and DELETE request from retrieving and creating new polygons in the database. The last two endpoints check to see if a point is contained within a polygon and to check if a polygon overlaps with an already created polygon. 


A more detailed breakdown of the endpoints and their implementation can be found in the [development](#development) section.

In order to save polygons in the database the [PostGIS extension](https://postgis.net/install/) was needed. This extension is a spatial database extender for the Postgres database. It adds support for geographic objects and allows location queries to be run in SQL. 

My first design of the database had the database dependencies/connections in the database section. Once I tried to unit test the database methods, I quickly realized this dependency would prevent me from easily writing unit test. Refactoring the dependence out of the methods and passing a pointer to a database into these methods allowed me to pass in a mock database for unit testing. 

#### Libraries Used
 
| Library | Description |
| ---------- | ---------------- |
| pq | A pure Go Postgre driver for Go’s database/sql package |
| golang-geo | Go library for translating, geocoding, and calculating distances between geographical points|
| gorila/mux | A powerful URL router and dispatcher for golang |
| go-sqlmock| SQL mock driver for golang to test database interactions |
| stretchr/testify | A toolkit with common assertions and mocks that work with the standard library |
| go-carpet | Shows test coverage in terminal for Go source files |

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


### Test


Testing and building was automated using TravisCI. My initial design of the HTTP handler methods didn’t allow for easy unit testing. The first implementation had a dependency of the database connection and would have required setting up a temporary HTTP server with proper environment variables to test the HTTP handlers. Given most of my testing had been automated through TravisCI, this wasn’t easily feasible. Instead I refactored the database dependencies out of my handlers using an interface. This allowed me to decouple my handlers from my database interaction and create mock datastores for testing.

Testing the database package required using a library for mock databases. The library used was called [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock). I was unable to find any libraries for a mock PostGIS database. Due to limited time, I was unable to build my own version of a mock database for testing the PostGIS database interactions. Given the model package contains no functionality there are no unit tests for this package.

[Go-carpet](https://github.com/msoap/go-carpet) was used to determine testing coverage. This metric was only mildly useful for this specific project as some of the packages had no functionality, and as mentioned, some of the specific database interactions could not be tested. 

Once each feature was complete integration testing was also performed using calls from the client. The correct implementation was tested by verifying the correct response from the HTTP requests. The React application interface allowed me to test request responses from a client application to the service. 



## Results

A rudimentary service was created at first with Go containing a single endpoint was created and tested locally. Once the basic app had been tested locally, it was moved to a DigitalOcean droplet (server). Docker was used to help facilitate easy deployment of new service iterations. Using a makefile and a few make commands, one is able to load up a new docker image and run it on the droplet. 

To help demonstrate the working API, a bootstrapped React App was created. Within this React App some basic HTTP requests were created and the results displayed on the UI to help test the functionality of the service. As each feature was added to the service, a demonstration of this functionality was added to the React App. 

The next feature implemented was the addition of database interactions within the service. A Postgres database hosted on Amazon Web Services (AWS) was used to host the database. A database package was created to handle interactions between the service and database. In order to allow the client to retrieve data from the database, endpoints were created that allowed the client to perform GET, POST, and DELETE http request. 

In the proposal project features were ranked as A, B, and C using a priority system. There were three levels of priority, with A being essential, B being want to implement, and C being extensions or extras. As displayed in the chart below, I was able to finish all of the critical aspects of the project. 

<img width="899" alt="screen shot 2018-05-02 at 9 10 46 pm" src="https://user-images.githubusercontent.com/13584530/39556291-5605560a-4e4d-11e8-80f6-8567257a6284.png"> 
FIGURE 7. FEATURE RESULTS

The algorithm that is used in the geofence package is the recursive hash filling algorithm discussed in the background section. This algorithm works best for queries of a large number of points in a polygon. As the number of polygons increases, this algorithm will be faster than the ray casting algorithm, but will not scale as well as others discussed in the background section of this report. I was able to begin coding the QuadTree data structure, but didn’t have time to finish the data structure or an algorithm to utilize the tree.

## Conclusion

A geofencing microservice was needed as part of a larger architecture for a location-based augmented reality application. To keep cost low, protect user's data, and provide the user with relevant information I built a geofencing microservice as my project.

I would take a Test Driven Development (TDD) approach in future projects. Using a TDD approach to development would improve design decisions and save time over the course of the project. I spent a large portion of my time making a good interface for demonstrating the functionality of my service. In the future I would focus more on the primary functionality of the project, such as implementing more advanced algorithms. Overall, I was able to complete all of the necessary features for a working project.

This was my first large project using Go as the development language. The current service implementation does not take advantage of concurrent programming, but it is planned to add this to the service in future development. Learning best practices for the language was time consuming, but a valuable learning process. This had a small impact on my productivity. Working to make a user-friendly interface within the React application took longer than anticipated. Demonstrating a working project at the end of our development process was necessary, but in the future I would research more time-efficient ways to build a user interface to demonstrate the service functionality. 

Query syntax for a PostGIS database is significantly different from a regular SQL database. I didn’t research this extension before using it, so I had to learn how to properly query the database and work with geometric shapes. Integrating PostGIS into the project took longer than anticipated. 

Future development of this microservice would involve removing the other functionality such as interacting with the user and client database. These database interactions would be factored out into separate services that handle client and user data. The current geofencing service would not efficiently handle tens of thousands of request per second. Implementing more sophisticated point-in-polygon algorithms and more appropriate spatial data structures would improve efficiency of the service. Another aspect of the service that could be improved is by adding concurrent design patterns. This would enable the microservice to run processes in the foreground and background simultaneously. This should ultimately improve the efficiency of the service. Finally, load and performance testing on the microservice would allow me to have a concrete analysis of the functionality of the service. With enough data I could accurately determine the total number of requests the service could handle and decide if one instance of the service will be enough to handle all of the incoming request.

## References

[1] "Ray-casting graphic". Available: https://www.google.com/url?sa=i&rct=j&q=&esrc=s&source=images&cd=&cad=rja&uact=8&ved=2ahUKEwj4vv63rOjaAhULUt8KHf6mA-EQjRx6BAgBEAU&url=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F217578%2Fhow-can-i-determine-whether-a-2d-point-is-within-a-polygon&psig=AOvVaw2XwYqoJczzLynGHIKPLk8e&ust=152539583349647,  [April 29, 2018].

[2] "Filling a geofence with geohashes". Available:https://www.google.com/url?sa=i&rct=j&q=&esrc=s&source=images&cd=&cad=rja&uact=8&ved=2ahUKEwi7uI3rrOjaAhXng-AKHW-yAlsQjRx6BAgBEAU&url=http%3A%2F%2Fwilldemaine.ghost.io%2Ffilling-geofences-with-geohashes%2F&psig=AOvVaw3BkMzF9YtXsrUp8VaVX6RW&ust=1525395936043284, [April 29, 2018].

[3] “Simple example of an R-tree for 2D rectangles” Available: https://en.wikipedia.org/wiki/R-tree, [April 27, 2018].

[4] Available: https://msdn.microsoft.com/en-us/library/bb259689.aspx, [April 27, 2018]

[5] M. Galetzka, P. Galuner. “A Simple and Correct Even-Odd Algorithm for the Point-In-Polygon Problem for Complex Polygons.” Internet: https://arxiv.org/pdf/1207.3502.pdf, [April 30, 2018].

[6] A. Bucchiarone, N. Dragoni, S. Dustdar, S. T. Larsen, M. Mazzara. “From Monolithic to Microservices: An experience report.” Internet: https://www.researchgate.net/publication/318653629_From_Monolithic_to_Microservices_An_experience_report, [April 30, 2018].

[7] M. M. Sardadi, M. S. bin Mohd Rahim, Z. Jupri, and D. bin Daman. “Choosing R-tree or Quadtree Spatial Data Indexing in One Oracle Spatial Database System to Make Faster Showing Geographical Map in Mobile Geographical Information System Technology.” Internet: https://pdfs.semanticscholar.org/c86e/a522e7872c44359b00a4102d16e72bbed891.pdf, [April 26, 2018]. 

[8] S. Tang, Y. Yu, R. Zimmerman, S. Obana. “39 Efficient Geo-fencing via Hybrid Hashing: A Combination of Bucket Selection and In-bucket Binary Search.” Internet: http://research.nii.ac.jp/~yiyu/GeoFence-20150512-FirstLook.pdf, [April 26, 2018].



