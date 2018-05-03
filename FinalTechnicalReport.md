# Final Technical Report

Jesse Cochran

April 28th, 2018

## Abstract

The goal of this project was to build a geofencing microservice that would allow clients to query the service and determine if a point is contained within a geofence. These types of services are common in industry, although are generally implemented as proprietary systems. The service built for this project is one service in a larger architecture for a location-based marketing application. The location-based functionality will allow for an accurate targeting of customers to help provide a personalized user experience. The geofencing service was built using [Golang](https://golang.org) and hosted in a [Docker](https://www.docker.com) container on a [Digital Ocean](https://www.digitalocean.com) droplet. A basic React application provides an easy to use interface to demonstrate the functionality of the service. This app allows the user to draw polygons using the Google Maps Application Programmable Interface (API) and click inside and outside of the polygons. Upon clicking, the application sends a Hyper Transfer Text Protocol (HTTP) request to the service and issues a notification on the screen to alert the results. If the click was outside of the polygon(s), then the user will be alerted; if it is inside the polygon(s), then the user will be alerted to which polygons the click is contained. All of the core features of the service outlined in the proposal were completed. [TravisCI](https://travis-ci.com) was used for build and testing automation. Git was used as the version control system and Github was used to host our projects as well as manage our feature sprints. 
Keywords: Geofence, Point-In-Polygon, Microservice

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
Within this technology trend, an application of geofences has arisen. A geofence is a part of a software program 
that uses GPS, WiFi or RFID to define geographical boundaries. With pushes in technology towards augmented reality, self-driving 
cars and IoT the need for location aware devices is becoming increasingly important.


### Problem

Applications that almost everyone uses on a daily basis such as Uber, Google Maps, Waze, and Yelp use some type of geofencing in their product. This is a core part of their product, so these implementations are very valuable and not open to the public. Most of the information regarding similar services are in the form of high-level overviews or brief engineering blogs describing their implementation of a geofencing service. 

While there are geolocation applications you can use as a service, they are costly and you have to forfeit your customers data to a third-party. Additionally, with a high number of users the service will need to process tens of thousands of transactions per second. With such high request volumes, outsourcing this responsibility to a third party would again prove costly. With the limited processing power on many of our devices, it is not feasible to use them to determine their location in relation to geofences. To attain a satisfactory user experience, secure the user’s data, and achieve reasonable response times this computation must be offloaded onto a server. The geofencing microservice built for this project will provide a solution to all of these issues.

### Background

Point-in-Polygon detection algorithms are common methods used to implement geofences. There are a wide array of point-polygon algorithm which all vary in complexity and efficiency. Ray-casting is one of the oldest and most well-known algorithms[6] for detecting if a point is contained within a polygon. 

![raycast](https://user-images.githubusercontent.com/13584530/39556011-35909918-4e4b-11e8-9e5d-902b00853753.png)
FIGURE 1. RAY-CASTING ALGORITHM [1]

This was the first algorithm implemented in the project. It was replaced to then implement the geohashing algorithm discussed below.

Another version is to use geohashes to recursively fill the polygon with geohashes and then hash a point to determine if it is contained in one of the hashes. To recursively fill the set you start with an empty fence. Then if the hash is completely contained within the fence completely then it is added to the set. If the hash intersects with the geofence, it isn’t added, but instead recurse with the hashes of its children. Once you reach the point in which the hash is equal to the maximum precision that is being checked, you add the hash to the set and stop recursing. Ultimately, if the hash doesn’t intersect with the geofence, then you don’t add it to the set, and you stop recursing. A visual representation of this process can be seen in the figure below.

![screenshotfromhashfill](https://user-images.githubusercontent.com/13584530/39506363-06c9eb7a-4da6-11e8-8f9a-ed61248c415a.png)  
FIGURE 2. RECURSIVE HASHING [2]

This was the next algorithm implemented for this project and is the current implementation in the geofencing microservice.

Within the large variety of algorithms used for geofencing, the most efficient ones all share some commonalities. Previous research has found many use R-Trees to organize their Minimum Bounding Rectangle(MBR) of polygons for the filtering stage [8]. 

An R-Tree is a spatial data structure based on a B-Tree that is used for spatial indexing methods. In the two-dimensional case of geofencing, the MBR is a simple bounding box (bbox) defined by a minimum and maximum coordinate. The check to determine if one object’s bbox is contained in another is a constant time operation.  Figure 1 below shows a good representation of an R-Tree.

![r-tree explanation](https://user-images.githubusercontent.com/13584530/39414423-5891e842-4c05-11e8-9f88-0a3dc050e339.png)  
FIGURE 3. R-TREE DATA STRUCTURE [3]

The average search time complexity for an R-Tree is O(log Mn) where M is the defined constant of the maximum number of children a node can have.

The other common spatial data structure used in the fastest point-in-polygon algorithms is a QuadTree. A QuadTree is a specialization of a generic kd-tree for 2-dimensional indexing. You take a flat project of a surface and divide the surface into quarters, generally called cells. Figure 2 below shows an example of QuadTree generation.

![bingmapquad](https://user-images.githubusercontent.com/13584530/39414425-5a95f818-4c05-11e8-93dd-cd7758246207.jpeg)  
FIGURE 4. QUADTREE EXAMPLE [4]

QuadTrees are used in popular mapping applications such as Google Maps and Bing Maps. Google Maps uses a S2 algorithm, which is a projection of the Earth’s sphere using cube mapping so each cell has a uniform surface area. The cells are arranged using a Hilbert Curve to conserve spatial locality in the cell label. A Hilbert Curve is a space filling curve that allows the range to cover the entire n-dimensional space. 

While R-Tree implementations often have a higher maximum throughput, heavy update activity of the geospatial data decreases the performance of the R-Tree implementations, but heavy updating of the Quad Tree will have no impact on performance [7].

A microservice is one service in part of a microservice architecture that structures an application as a group of loosely coupled services which generally implement business capabilities. Benefits of this architecture style allows for scalability, flexibility, and portability[4]. 

### Objectives

The idea of this project is to create a functionable and scalable solution to geofencing. Implementing a geofence requires lookups using CPU-intensive point-in-polygon algorithms in order to determine if an object exist in a geofence. The algorithms used in this project were not the optimal solutions for point-in-polygon detection. 

A microservice architecture is an architecture style that is a structured a collection of loosely coupled services that generally implement some type of business capabilities [2].  This project is intended to provide a microservice that allows a client to query the service and determine if a point is contained within a polygon or geofence.  The microservice architecture would handle all server-side components of an augmented reality based marketing application. 

There are many libraries that help with location data and implementation of common data structures that will be useful in the implementation of this service.

In the project features were ranked as A, B, and C using a priority system. There were three levels of priority, with A being essential, B being want to implement, and C being extensions or extras. They are outlined in the figure below.

![faeatsasa](https://user-images.githubusercontent.com/13584530/39502959-23fd59ee-4d91-11e8-81c0-7df6190d0762.png)  

FIGURE 4. FEATURE TABLE



### Potential Users

While this microservice could be used alone, the functionality of the service fit better into a larger system that uses the service for geolocation purposes. The microservice designed and implemented for this project is just one service that is part of a larger system, as can be seen in the figure below.

![microservicecapstone](https://user-images.githubusercontent.com/13584530/39459534-30f312dc-4cca-11e8-9dbc-da9541c494cd.png)  
FIGURE 6. MICROSERVICE ARCHITECTURE DIAGRAM

In theory, almost any company that needs to implement geofencing into their application could use this service. Specifically, this microsevice will be used to handle geolocation operations in a location-based augment reality application. 


## Design, Development and Test

The development of this project was performed using 1 week sprints. Each sprint we planned features that were to be completed by the end of each sprint. Each of these features should be a testable unit of code. The development of this project was performed using 1 week sprints. Each sprint, we planned features that were to be completed by the end of each sprint. Each of these features should be a testable unit of code. Github Projects and Issues were used to track each of the features and our current position in each sprint. Git was used as the version control throughout the project. 

### Design

This microservice is designed to be part of a larger system of microservices, in which each is decoupled from the other and has its own functionality. I chose to split the microservice into the three main categories User, Client, and Polygon. In a true microservice the functionality of the user and client would be separated into different services. But, due to time constraints and project scope they were coupled with the geofencing functionality as one microservice. The service is split into four main packages:
* model - Holds all of the structs and models of the service. This package contains no functionality.
* database - This package handles all of the database interactions.
* handlers - This package contains all of the handlers and methods for building the router containing the handlers.
* geofence - This package contains the implementation of the geofence algorithm and associated functionality.

Golang was chosen as the development language for this microservice for a variety of reasons:

* Point in polygon lookups require CPU-intensive algorithms. Golang is a systems language that is designed to be fast and efficient.
* Low latency and high throughput. The service needs to be able to handle thousands of requests, with reach request taking less than 100 milliseconds.
* Concurrent Design. This service must constantly refresh in-memory geofences in the background. Background refreshing can tie up the CPU and slow query response time. Goroutines can be executed on multiple cores and allow the service to run background queries in parallel with foreground queries.

My first design of the database had the database dependencies/connections in the database section. Once I tried to unit test the database methods, I quickly realized this dependency would prevent me from easily writing unit test. Refactoring the dependence out of the methods and passing a pointer to a database into these methods allowed me to pass in a mock database for unit testing. 

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

Testing and building was automated using TravisCI. My initial design of the HTTP handler methods didn’t allow for easy unit testing. The first implementation had a dependency of the database connection and would have required setting up a temporary HTTP server with proper environment variables to test the HTTP handlers. Given most of my testing had been automated through TravisCI, this wasn’t easily feasible. Instead I refactored the database dependencies out of my handlers using an interface. This allowed me to decouple my handlers from my database interaction and create mock datastores for testing.

Testing the database package required using a library for mock databases. The library used was called [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock). I was unable to find any libraries for a mock PostGIS database. Due to limited time, I was unable to build my own version of a mock database for testing the PostGIS database interactions. Given the model package contains no functionality there are no unit tests for this package.

[Go-carpet](https://github.com/msoap/go-carpet) was used to determine testing coverage. This metric was only mildly useful for this specific project as some of the packages had no functionality, and as mentioned, some of the specific database interactions could not be tested. 

Once each feature was complete integration testing was also performed using calls from the client. The correct implementation was tested by verifying the correct response from the HTTP requests. The React application interface allowed me to test request responses from a client application to the service. 


## Results

A rudimentary service was created at first with Go containing a single endpoint was created and tested locally. Once the basic app had been tested locally, it was moved to a DigitalOcean droplet (server). Docker was used to help facilitate easy deployment of new service iterations. Using a makefile and a few make commands, one is able to load up a new docker image and run it on the droplet. 

To help demonstrate the working API, a bootstrapped React App was created. Within this React App some basic HTTP requests were created and the results displayed on the UI to help test the functionality of the service. As each feature was added to the service, a demonstration of this functionality was added to the React App. 

The next feature implemented was the addition of database interactions within the service. A Postgres database hosted on Amazon Web Services (AWS) was used to host the database. A database package was created to handle interactions between the service and database. In order to allow the client to retrieve data from the database, endpoints were created that allowed the client to perform GET, POST, and DELETE http request. 

As each feature was added to the service, a demonstration of this functionality was added to the React App. 
In order to save polygons in the database the PostGIS extension was needed. This extension allows one to save geometrical shapes in the database. Query syntax for a PostGIS database is significantly different from a regular SQL database. I didn’t research this extension before using it, so I had to learn how to properly query the database and work with geometric shapes. Integrating PostGIS into the project took longer than anticipated. 

In order to save polygons in the database the [PostGIS extension](https://postgis.net/install/) was needed. This extension allows one to save geometrical shapes in the database. Query syntax for a PostGIS database is significantly different from a regular SQL database. I didn’t research this extension before using it, so I had to learn how to properly query the database and work with geometric shapes. Integrating PostGIS into the project took longer than anticipated. 

In the proposal project features were ranked as A, B, and C using a priority system. There were three levels of priority, with A being essential, B being want to implement, and C being extensions or extras. As displayed in the chart below, I was able to finish all of the critical aspects of the project. 

![featuressss](https://user-images.githubusercontent.com/13584530/39502875-a7f57c5a-4d90-11e8-9d39-890a188b9935.png)  
FIGURE 5. FEATURE RESULTS


This was my first large project using Go as the development language. The current service implementation does not take advantage of concurrent programming, but it is planned to add this to the service in future development. Learning best practices for the language was time consuming, but a valuable learning process. This had a small impact on my productivity. Working to make a user-friendly interface within the React application took longer than anticipated. Demonstrating a working project at the end of our development process was necessary, but in the future I would research more time-efficient ways to build a user interface to demonstrate the service functionality. 

The algorithm that is used in the geofence package is the recursive hash filling algorithm discussed in the background section. This algorithm works best for queries of a large number of points in a polygon. As the number of polygons increases, this algorithm will be faster than the ray casting algorithm, but will not scale as well as others discussed in the background section of this report. I was able to begin coding the QuadTree data structure, but didn’t have time to finish the data structure or an algorithm to utilize the tree.


## Conclusion

A geofencing microservice was needed as part of a larger architecture for a location-based augmented reality application. To keep cost low, protect user's data, and provide the user with relevant information I built a geofencing microservice as my project. 

After refactoring my code to allow for proper unit testing, I would take a Test Driven Development (TDD) approach in future projects. Using a TDD approach to development would improve design decisions and save time over the course of the project. I spent a large portion of my time making a good interface for demonstrating the functionality of my service. In the future I would focus more on the primary functionality of the project, such as implementing more advanced algorithms. Overall, I was able to complete all of the necessary features for a working project.

Future development of this microservice would involve removing the other functionality such as interacting with the user and client database. These database interactions would be factored out into separate services that handle client and user data. The current geofencing service would not efficiently handle tens of thousands of request per second. Implementing more sophisticated point-in-polygon algorithms and more appropriate spatial data structures would improve efficiency of the service. Another aspect of the service that could be improved is by adding concurrent design patterns. This would enable to run services in the foreground and background simultaneously. This should ultimately improve the effeciency of the service. Finally, load and performance testing on the microservice would allow me to have a concrete analysis of the functionality of the service. With enough data I could accurate determine the total number of request the service could handle and decide if one instance of the service will be enough to handle all of the incoming request. 


## References

[1] "Ray-casting graphic". Available: https://www.google.com/url?sa=i&rct=j&q=&esrc=s&source=images&cd=&cad=rja&uact=8&ved=2ahUKEwj4vv63rOjaAhULUt8KHf6mA-EQjRx6BAgBEAU&url=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F217578%2Fhow-can-i-determine-whether-a-2d-point-is-within-a-polygon&psig=AOvVaw2XwYqoJczzLynGHIKPLk8e&ust=1525395833496471 [April 29, 2018]

[2] "Filling a geofence with geohashes". Available:https://www.google.com/url?sa=i&rct=j&q=&esrc=s&source=images&cd=&cad=rja&uact=8&ved=2ahUKEwi7uI3rrOjaAhXng-AKHW-yAlsQjRx6BAgBEAU&url=http%3A%2F%2Fwilldemaine.ghost.io%2Ffilling-geofences-with-geohashes%2F&psig=AOvVaw3BkMzF9YtXsrUp8VaVX6RW&ust=1525395936043284 [April 29, 2018]

[3] “Simple example of an R-tree for 2D rectangles” Available: https://en.wikipedia.org/wiki/R-tree [April 27, 2018]

[4] Available: https://msdn.microsoft.com/en-us/library/bb259689.aspx [April 27, 2018]

[5] Add reference here for ray-casting paper

[6] Add reference here for microservice paper


[7] M. M. Sardadi, M. S. bin Mohd Rahim, Z. Jupri, and D. bin Daman. “Choosing R-tree or Quadtree Spatial Data Indexing in One Oracle Spatial Database System to Make Faster Showing Geographical Map in Mobile Geographical Information System Technology.” Internet: https://pdfs.semanticscholar.org/c86e/a522e7872c44359b00a4102d16e72bbed891.pdf, [April 26, 2018]. 

[8] S. Tang, Y. Yu, R. Zimmerman, S. Obana. “39 Efficient Geo-fencing via Hybrid Hashing: A Combination of Bucket Selection and In-bucket Binary Search.” Internet: http://research.nii.ac.jp/~yiyu/GeoFence-20150512-FirstLook.pdf, [April 26, 2018]



