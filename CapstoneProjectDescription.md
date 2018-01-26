#Capstone Project: Geofenced
###Jesse Cochran

##Project Overview
The idea of this project is to create a functionable and scalable solution to geofencing. Implementing a geofence requires lookups using CPU-intensive point-in-polygon algorithms in order to determine if an object exist in a geofence. One can brute force the search to determine which geofences the current latitude and longitude point is contained in with something such as a ray-casting algorithm, but this solution is much to slow. Another option is to use an RTree or S2, a spatial indexing library. 

Additionally, one can further design a multi-tiered architecture to optimize geofence indexing within certain locations. This would be a two-level hierarchy where the first level is a city or populated place’s geofence, and the second is the geofences within this city or populated area. While this design does not reduce the runtime complexity O(N), it has been implemented and shown to reduce N from the order of 10,000 to order of 100.

Initially in this project, I will plan to implement a server-side service with the most reasonable solution to the point-in-polygon algorithm to ensure I am able to achieve a working solution. Time permitting, I will plan to explore optimal solutions and implementations. 

This project could theoretically be implemented in any popular “backend” language, although from research it appears Go (Golang) is best suited for this type of service. Background refreshing can hold up CPU resources for “long” periods of time. Goroutines will allow us to improve this bottleneck by running background jobs in parallel with foreground queries.

This service is part of a larger system that will be discussed further with the professor in person. We can discuss use cases and how this service will be used within the system.

###Project Extensions

As mentioned there are many point-in-polygon algorithms that can be used. Exploring different algorithms and architectures would allow one to optimize this part of the service. 

Implementing geofencing by consistent use of the user’s GPS will drain the battery on their devices very quickly, and thus is not optimal. While this isn’t directly part of the service, it would be useful to research and test the most optimal ways to use the backend service to optimize battery resources. 

Other extensions of this project can be explored. These extensions deal with the larger system in which this server-side service will be contained, and thus will be discussed with the professor in person.


##Similar Work
Applicatoins and services that implement geofencing exist, but most are proprietary systems. 


##Previous Experience
Previous experience at my internship working on web apps with assist me in easily producing a front end to test and demonstrate the functionality of the microservice. Additionally, I have built a number of RESTful APIs for personal projects and during hackathons.

##Technology
The technology stack that will be used is listed below:
-Golang: Backend language
-React(JS/HTML/CSS): Frontend language used to demonstrate functionality of backend
-Docker: Container for application deployment
-Travis CI: Build automation
-Git/Github: Source Version Control
-DigitalOcean: Droplet to host the service
-Google Maps API: Helps gather coordinates for testing

##Risk Areas


#Product Backlog