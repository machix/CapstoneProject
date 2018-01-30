# Capstone Project: Geofenced

### Jesse Cochran


## __*Background and Motivation*__
In our ever increasingly mobile first world, applications with location awareness have become increasingly popular and useful. Within this technology trend, an application of geofences has arisen. A geofence is a part of a software program that uses GPS, WiFi or RFID to define geographical boundaries. With pushes in technology towards augmented reality, self-driving cars and IoT the need for location aware machines is becoming increasingly popular.

Consistently using systems resources to determine the device location can become computationally expensive. Offloading this work onto a server-side service will free up systems resources and save energy on all of these devices. 


## __*Project Overview*__
The idea of this project is to create a functionable and scalable solution to geofencing. Implementing a geofence requires lookups using CPU-intensive point-in-polygon algorithms in order to determine if an object exist in a geofence. One can brute force the search to determine which geofences the current latitude and longitude point is contained in with something such as a ray-casting algorithm, but this solution is much to slow. Another option is to use an RTree or S2, a spatial indexing library. 

Additionally, one can further design a multi-tiered architecture to optimize geofence indexing within certain locations. This would be a two-level hierarchy where the first level is a city or populated place’s geofence, and the second is the geofences within this city or populated area. While this design does not reduce the runtime complexity O(N), it has been implemented and shown to reduce N from the order of 10,000 to order of 100.

Initially in this project, I will plan to implement a server-side service with the most reasonable solution to the point-in-polygon algorithm to ensure I am able to achieve a working solution. Time permitting, I will plan to explore optimal solutions and implementations. Please see the Algorithm Appendix for further information on algorithm and data structure implementations.

This project could theoretically be implemented in any popular “backend” language, although from research it appears Go (Golang) is best suited for this type of service. Background refreshing can hold up CPU resources for “long” periods of time. Goroutines will allow us to improve this bottleneck by running background jobs in parallel with foreground queries.

The geofencing service with endpoints will be novel code that will be created during this project. Additionally, while there exist some libraries with the data structure implementations, the implementation of both the data structures and algorithms to handle the geofencing portion of the backend does not exist in any open-source projects.

I have discussed with the professor the potential use for this application, as well as described in the background and motivation section why this project/service would be useful on a commercial level. 

## __*Project Features and Implementation*__
I will use a priority system to rank the importance of the features. There will be three levels of priority, with A being essential, B being want to implement, and C being extensions or extras. *Please note that while these feature descriptions are in paragraph form, they will be divided into single-purpose, manageable, and testable features on Github.*

#### **Features (A)**
The most rudimentary and basic implementation of this project will be a simple service with one endpoint, hosted on DigitalOcean, that returns a "hello world" type of response. The next step of the project will be to implement a simple interface that will enable me to easily interact with the API. This interface will also help demonstrate completion of features on the backend. Next we will connect the service to a database and then test to ensure the service is correctly uploading data from our FrontEnd to a database.

The next step in the process will implement more endpoints that will allow the client to post coordinates, retrieve coordinates based on id, delete coordinates, and query the api and retrieve a response that contains the latitude, longitude and id of all points contained within a set geofence. All of these with be implemented through the standard HTTP request (GET, POST, DELETE, etc.)

Finally, the data structure and algorithm that performs the geofencing task will be implemented. First I will plan to implement the ray casting algorithm as the most simple implementation. Further improvements are discussed in the next feature section and details can be found in the Appendix.

#### **Features (B)**
Adding to this implementation one of the endpoints will also allow the request to have a distance parameter to determine the radius of the geofence. During the process of building and thinking about the service, there may be other useful functionality of the service that would nice, but not necessary to implement.

As mentioned there are many point-in-polygon algorithms that can be used. Exploring different algorithms and architectures would allow one to optimize this part of the service. 

#### **Features (C)**
Implementing geofencing by consistent use of the user’s GPS will drain the battery on their devices very quickly, and thus is not optimal. While this isn’t directly part of the service, it would be useful to research and test the most optimal ways to use the backend service to optimize battery resources. 

Although not essential the main application of this service, implementing a user friendly and improved UI for interacting with the service would be ideal to improve user experience.

Assuming I am able to implement multiple iterations and improvements of the service, performing a runtime and performance analysis on the service with different implementation would be interesting and very informative. 


## __*Similar Work*__
Applications and services that implement geofencing exist, but most are proprietary systems. Apps that almost everyone uses on a daily basis such as Uber, Google Maps, Waze, Yelp, and the list goes on and on, use some type of geofencing in their product. This is a core part of their product, so these implementation are very valuable and not open to the public. Most of the information regarding similar services are in the form of high-level overviews of their engineering architecture. There are many libraries that help with location data and implementation of common data structures that will be useful in the implementation of this service.

##### Blog Post
* [Uber blog about geofencing](https://eng.uber.com/go-geofence/)
* [Geofencing with Google Data Flow](https://cloudplatform.googleblog.com/2014/12/geofencing-340-million-nyc-taxi-with-google-cloud-dataflow.html)
* [Google’s S2, geometry on the sphere, cells and Hilbert curve](http://blog.christianperone.com/2015/08/googles-s2-geometry-on-the-sphere-cells-and-hilbert-curve/)
* [S2 Cells and space filling curves](https://medium.com/sidewalk-talk/s2-cells-and-space-filling-curves-keys-to-building-better-digital-map-tools-for-cities-a312aa5e2f59)

##### Products
* [Bluedot](https://bluedotinnovation.com)
* [Geomoby](https://www.geomoby.com)
* [Carto](https://carto.com)


## __*Previous Experience*__
Previous experience at my internship working on web apps with assist me in easily producing a front end to test and demonstrate the functionality of the microservice. Additionally, I have built a number of RESTful APIs during my internship, for personal projects and during hackathons. My solid background in data structures and algorithms will help me navigate through the implementation of the specific algorithms outlined in the project description and algorithm appendix. 


## __*Testing*__
Testing is an important, if not essential, part of software development. Golang has a testing library that will be used to test this project. Once the database is implement I will plan to populate the database with "dummy" data to enable testing of the service and queries. Testing the frontend will be done with Jasmine and Karma. Jasmine is a behavior driven framework for testing javascript and Karma is a test runner. Automated testing will be configured and performed using TravisCI. Any actions that I perform more than 2 or 3 times while building the project, I will plan to automate.

Time permitting, I will also plan to do performance testing on the microservice to determine the load the microservice is able to handle. If I reach this point in the project, then I will have to do more research on tools and libraries that will enable me to easily perform these test.

## __*Technology*__
The technology stack that will be used is listed below:
* Golang: Backend language
* React(JS/HTML/CSS): Frontend language used to demonstrate functionality of backend
* Jasmine/Karma: Testing frontend
* Docker: Container for application deployment
* Travis CI: Build and test automation
* Git/Github: Source Version Control
* DigitalOcean: Droplet to host the service
* Google Maps API: Helps gather coordinates for testing
* PostGIS: Database to connect to backend 

## __*Libraries*__
Solving a problem that has already been solved, is not a good use of development time. For this reason, I will be utilizing a wide array of libraries. Many of these libraries will be included in the Go standard library. Below is a list of libraries that I plan to use and their use case, although more may be added in the future as problems arise, some replaced if better alternatives are found, and others as inspiration:

* gorilla/mux - implements a request router and dispatcher for matching incoming request to their respective handler
* cors - CORS is a net/http handler implementing Cross Origin Resource Sharing W3 specification in Golang
* go-geo - This is a library for manipulating geometric shapes (Primary spherical). This has a partial implementation of the S2 solution mentioned previously. This isn't a complete implementation, so this may or may not be used upon further research.
* dep - go dependency management tool
* rtreego - A library for efficiently storing and querying spatial data


I will be using the built in Go libraries to test my project, in addition to Jasmine/Karma for testing on the frontend. All of the builds and test can be automated through Travis CI and scripts. The front end portion of the project uses npm and yarn to manage dependencies. dep will be included on the backend to manage dependencies in Go. Go has a built in formatting tool that can be used by adding it to a build script. I will plan to use JSCS for linting with Javascript.

## __*Risk Areas*__
The main personal challenges of this project isn't in developing the API, but rather the implementation of the advanced data structures and point-in-polygon algorithms used to detect the user's location in relation to other relevant points. Almost all of the implementations of geofencing use data structures that I am aware of, but unfamiliar with. This poses a risk due to the fact I will need to do thorough research to ensure my understanding of the data structure's mechanisms are sufficient to allow me to succeed in properly implementing the algorithm that utilizes the data structures. I will also investigate current existing implementations of these data structures (possibly in other languages) to get a better understanding of their implementations. 

In every project I have worked on dealing with HTTP request, I have always run into CORS (Cross Origin Resource Sharing) issues. For security reasons, browsers restrict cross-origin HTTP requests initiated from within scripts. Problems can quickly arise from this security precaution while testing communication between the front-end and backend of an application. I have solved this problem many times, and I am using a library mentioned previously to help solve this issue before it arises.

I lack experience in developing high concurrent services, so I expect to run into roadblocks dealing with concurrency within the application. I plan to overcome this through research and reading about concurrent design patterns. Also, by reaching out for help with the professor and any other teachers that may be able to help. Another option is to use stackoverflow. Ultimately, I will seek guidance from people more knowledgeable than myself if I run into any roadblocks. 

## *Algorithm and Data Structure Appendix*

#### Ray Casting Algorithm
One of the easiest ways to determine if a point is in a polygon is to test how many times a ray starting from the point moving in a fixed direction intersects with edges of the polygon. The the point is outside of the target polygon then the ray will intersect an even number of times and if it is inside the polygon then it will intersect and odd number of times. One edge case in which this algorithm fails is if the point lies directly on the edge of the polygon. 

There is another "naive" algorithm called the winding number algorithm. I do not plan on using this, but this is another effective, but slow, approach to finding points in a polygon.

#### R-Tree
An R-Tree is based on a B-Tree, but used for multi-dimensional objects. The R-Tree allows us to compare minimum bounding rectangles to other objects minimum bounding rectangles. The minimum bounding rectangles is a simple bounding box (bbox) which is defined by the minimum and maximum coordinate. Checking if an object's bbox is inside another is a constant time operation. 

#### Quad Tree
A quad tree is a specialized kd-tree used for 2D indexing. Generally you would take a flat projection of your "map" or space and divide into quarters or cells. You then divide those recursively until you hit a defined depth, and these will be the leaves of the tree.

#### S2
This is a specialized implementation of a quadtree. This is used in many mapping systems. The S2 project is done via cube mapping, this is where you use six faces of a cube as the map shape. This is done so that each cell has a similar surface area. The cells are also organized using a space spilling curve to conserve locality in the cell. 
