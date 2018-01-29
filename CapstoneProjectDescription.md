## Capstone Project: Geofenced

### Jesse Cochran

## Project Overview
The idea of this project is to create a functionable and scalable solution to geofencing. Implementing a geofence requires lookups using CPU-intensive point-in-polygon algorithms in order to determine if an object exist in a geofence. One can brute force the search to determine which geofences the current latitude and longitude point is contained in with something such as a ray-casting algorithm, but this solution is much to slow. Another option is to use an RTree or S2, a spatial indexing library. 

Additionally, one can further design a multi-tiered architecture to optimize geofence indexing within certain locations. This would be a two-level hierarchy where the first level is a city or populated place’s geofence, and the second is the geofences within this city or populated area. While this design does not reduce the runtime complexity O(N), it has been implemented and shown to reduce N from the order of 10,000 to order of 100.

Initially in this project, I will plan to implement a server-side service with the most reasonable solution to the point-in-polygon algorithm to ensure I am able to achieve a working solution. Time permitting, I will plan to explore optimal solutions and implementations. 

This project could theoretically be implemented in any popular “backend” language, although from research it appears Go (Golang) is best suited for this type of service. Background refreshing can hold up CPU resources for “long” periods of time. Goroutines will allow us to improve this bottleneck by running background jobs in parallel with foreground queries.

The service with endpoints will be novel code that will be created during this project. Additionally, the implementation of a variety of data structures to handle the geofencing portion of the backend does no exist in any open-source projects. 

## Project Features and Implementation
We will use a priority system to rank the importance of the features. There will be three levels of priority, with A being essential, B being want to implement, and C being extensions or extras.

#### Features (A)
The most rudimentary and basic implementation of this project will be a simple service with one endpoint, hosted on DigitalOcean, that returns a "hello world" type of response. The next step of the project will be to implement a simple interface that will enable us to easily interact with the API. Next we will connect the service to a database and then test to ensure the service is correctly uploading data from our FrontEnd to a database.

The next step in the process will implement more endpoints that will allow the client to query the api and retrieve a response that contains the latitude, longitude and id of all points contained within a set geofence. 

#### Features (B)
Adding to this implementation the endpoint will also allow the request to have a distance parameter to determine the radius of the geofence. 

#### Features (C)

As mentioned there are many point-in-polygon algorithms that can be used. Exploring different algorithms and architectures would allow one to optimize this part of the service. 

Implementing geofencing by consistent use of the user’s GPS will drain the battery on their devices very quickly, and thus is not optimal. While this isn’t directly part of the service, it would be useful to research and test the most optimal ways to use the backend service to optimize battery resources. 


## Similar Work
Applications and services that implement geofencing exist, but most are proprietary systems that charge for use. There are many libraries that help with location data and implementation of common data strucutres that will be useful in the implementation.


## Previous Experience
Previous experience at my internship working on web apps with assist me in easily producing a front end to test and demonstrate the functionality of the microservice. Additionally, I have built a number of RESTful APIs during the internship, for personal projects and during hackathons. My solid background in data structures and algorithms will help me navigate through the implemenation of the specific algorithms outlined in the project description. 

# Testing
Testing is an important, if not essential, part of software development. Golang has a testing library that will be used to test this project. I will also plan to do performance testing on the microservice to determine the load the microservice is able to handle.

## Technology
The technology stack that will be used is listed below:
-Golang: Backend language
-React(JS/HTML/CSS): Frontend language used to demonstrate functionality of backend
-Jasmine/Karma: Testing frontend
-Docker: Container for application deployment
-Travis CI: Build automation
-Git/Github: Source Version Control
-DigitalOcean: Droplet to host the service
-Google Maps API: Helps gather coordinates for testing

## Libraries

Solving a problem that has already been solved, is not a good use of development time. For this reason, I will be uitilizing a wide array of libraries. Many of these libraries will be included in the Go standard library. Below is a list of libraries that I plan to use and their use case, although more may be added in the future as problems arise:

gorilla/mux -
cors - Used to handle cross origin domain sharing that is blocked by the browser


I will be using the built in Go libraries to test my project, in addition to Jasmine/Karma for testing on the frontend. All of the builds and test can be automated through Travis CI. The front end portion of the project uses npm and yarn to manage dependences. dep will be included on the backend to manage dependencies in Go. Go has a built in formatting tool that can be used by adding it to a build script.

## Risk Areas

The challenge of this project isn't in developer the API, but rather the implementation of the advanced data structures and point-in-polygon algorithms used to detect the user's location in relation to other relevant points.

In every project I have worked on dealing with HTTP request, I have always run into CORS (Cross Origin Resource Sharing) issues. For security reasons, browsers restrict cross-origin HTTP requests initiated from within scripts. Problems can quickly arise from this security precaution while testing communication between the front-end and backend of an application.

Almost all of the implemenations of geofencing use data structures that I am aware of, but unfamiliar with. This poses a risk due to the fact I will need to do thorough research to ensure my understanding of the data structure mechanisms are sufficient to allow me to succeed in properly implementing the data structure in the geofencing algorithm. 



