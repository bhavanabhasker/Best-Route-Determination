#Best Route Determination 

##Objective 
Given source and destination, determine the shortest and best possible route using minimum walk algoritm 
Uber api will be used to determine the costs for the travel and eta determination. 

##How to execute ?

###Setup : 

1. Download all the files from github

###Post Installation:

2. Run the following dependencies,

<pre>
go get gopkg.in/mgo.v2
go run *.go
</pre>

##Application Framework and Notes 

1. The locations for this program is taken from the locations db on mongo lab.

2. You can run the program with the below sample locations :
<pre>
Location ID
12348
12349
12345
12346
12347
</pre>


### For creating new location 
<pre>
go get github.com/bhavanabhasker/cmpe273-assignment2/rest
POST http://localhost:8080/locations
Request : { "name" : "John Smith", "address" : "123 Main St", "city" : "San Francisco", "state" : "CA", "zip" : "94113" }
Response : The location ID is returned 
</pre>

###Endpoints Information 

1. Post /trips to plan a trip 
<pre>
Request format : 
{
    "starting_from_location_id": "12345",
    "location_ids" : [ "12346", "12347", "12348", "12349" ] 
}
Response :
{
  "id": 1126,
  "status": "planning",
  "starting_from_location_id": "12345",
  "best_route_location_ids": [
    "12347",
    "12349",
    "12346",
    "12348"
  ],
  "total_uber_costs": 144,
  "total_uber_duration": 6203,
  "total_distance": 79.13
}
</pre>


2. GET trip/{trip_id}
<pre>
Response :
{
  "id": 1126,
  "status": "planning",
  "starting_from_location_id": "12345",
  "best_route_location_ids": [
    "12347",
    "12349",
    "12346",
    "12348"
  ],
  "total_uber_costs": 144,
  "total_uber_duration": 6203,
  "total_distance": 79.13
}
</pre>


3. PUT  /trips/{trip_id}/request
<pre>
Response : 
{
  "id": 1126,
  "status": "requesting",
  "starting_from_location_id": "12345",
  "next_destination_location_id": "12347",
  "best_route_location_ids": [
    "12347",
    "12349",
    "12346",
    "12348"
  ],
  "total_uber_costs": 144,
  "total_uber_duration": 6203,
  "total_distance": 79.13,
  "uber_wait_time_eta": 15
}
</pre>


 




                

 


