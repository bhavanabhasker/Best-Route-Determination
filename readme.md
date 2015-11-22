How to execute ?

Setup : 
Download all the files from github

Post Installation:

Run the following dependencies,

go get gopkg.in/mgo.v2

How to execute?
go run *.go

Pre req: The locations for this program is taken from the locations db on mongo lab.

You can run the program with the below sample locations :
Location ID
12348
12349
12345
12346
12347

To generate the new locations ,
Follow below,
go get github.com/bhavanabhasker/cmpe273-assignment2/rest

and create new locations using POST http://localhost:8080/locations
EG: Request : { "name" : "John Smith", "address" : "123 Main St", "city" : "San Francisco", "state" : "CA", "zip" : "94113" }

take the location id in the response

Execution :
The following endpoints are implemented
1. Post /trips to plan a trip 

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

2. GET trip/{trip_id}
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
3. PUT  /trips/{trip_id}/request
Please note that,for securoty reasons,the access token for this has been emailed to the professor. This would only work if the access token 
is specified in the function GenerateRequestId (GenerateRequestId.Go file ).
go to the function and 
change the value of access token 
accesstoken := "abcdef1g"


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

Please note that the current next_destination_location_id is 12347. Subsequent calls would change the next destination and eventually return back to the origin.

 




                

 


