package main

import (
	"encoding/json"
	"log"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Coordinates struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}
}
type Locations []Location

// List of price estimates
type PriceEstimates struct {
	StartLatitude  float64
	StartLongitude float64
	EndLatitude    float64
	EndLongitude   float64
	Prices         []PriceEstimate `json:"prices"`
}

var iterator int
var prev_location string

func ConnectDB() *mgo.Session {
	session, err := mgo.Dial("mongodb://bhavana:bhavana@ds037244.mongolab.com:37244/tests")
	if err != nil {
		log.Fatal(err)
	}
	return session
}
func RepoCreateTrip(t Trip) Trip {
	session := ConnectDB()
	c := session.DB("tests").C("Trips")
	currentid := []Trip{}
	err := c.Find(nil).All(&currentid)
	if len(currentid) == 0 {
		t.Id = 1122
	} else {
		t.Id = 1122 + len(currentid)
	}
	t.Status = "planning"
	t.Startingfromlocationid = t.Startingfromlocationid
	//find starting Coordinates
	startinglocation := Location{}
	id, _ := strconv.Atoi(t.Startingfromlocationid)
	s := session.DB("tests").C("locations")
	err = s.Find(bson.M{"id": id}).One(&startinglocation)
	if err != nil {
		log.Fatal(err)
	}

	startlatitude := startinglocation.Coordinates.Lat
	startlongitude := startinglocation.Coordinates.Long
	location := t.Bestroutelocationids
	// Generic Minimum Walk algorithm used
	// determine the best routes frm the starting location
	// to all the points
	// visit the verticies from starting point to all locations
	var BestRoutes = make(RouteEstimates, 0)
	var locations []string
	BestRoutes = DetermineBest(id, startlatitude, startlongitude, location)
	for _, c := range BestRoutes {
		t.Totalubercosts += c.Costs
		t.Totaldistance += c.Distance
		t.Totaluberduration += c.Duration
		end := strconv.Itoa(c.Endlocation)
		locations = append(locations, end)
	}
	t.Bestroutelocationids = nil
	t.Locationids = locations
	clearBestRoutes()
	//save in the db
	Saveindb(t)
	BestRoutes = nil
	return t
}
func FindTrip(id int) Trip {
	session := ConnectDB()
	c := session.DB("tests").C("Trips")
	result := Trip{}
	err := c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
func UpdateTrip(id int) Trip {
	session := ConnectDB()
	c := session.DB("tests").C("Trips")
	result := Trip{}
	err := c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	s := NewStack()
	locations := len(result.Locationids)
	origin := result.Startingfromlocationid
	// Push the origin
	var node Node
	node.Value = origin
	s.Push(&node)
	locations = locations + 1
	for i := (locations - 2); i >= 0; i-- {
		var node Node
		node.Value = result.Locationids[i]
		s.Push(&node)
	}

	var next_location string
	var startLocation string
	startlocation := Location{}
	if iterator == 0 {
		element, count := s.Pop(0)
		iterator = count
		next_location = element.Value
		startLocation = result.Startingfromlocationid
		locationid, _ := strconv.Atoi(startLocation)
		startlocation = FindinLocationService(locationid)
		prev_location = next_location
	} else {
		element, count := s.Pop(iterator)
		iterator = count
		startLocation = prev_location
		locationid, _ := strconv.Atoi(startLocation)
		startlocation = FindinLocationService(locationid)
		next_location = element.Value
		prev_location = next_location
	}
	endLocation := Location{}
	endid, _ := strconv.Atoi(next_location)
	endLocation = FindinLocationService(endid)
	data := getResponse(startlocation.Coordinates.Lat,
		startlocation.Coordinates.Long,
		endLocation.Coordinates.Lat,
		endLocation.Coordinates.Long)
	var pe PriceEstimates
	if e := json.Unmarshal(data, &pe); e != nil {
		log.Fatal(e)
	}
	var request Request
	request.product_id = pe.Prices[0].ProductId
	request.start_latitude = startlocation.Coordinates.Lat
	request.start_longitude = startlocation.Coordinates.Long
	request.end_latitude = endLocation.Coordinates.Lat
	request.end_longitude = endLocation.Coordinates.Long
	rideRequest := GenerateRequestId(request)
	makerequest := MakeRequest(rideRequest.RequestId)
	// Fill in the new values
	var t Trip
	t.Id = id
	if origin == next_location {
		t.Status = "accepted"
	} else {
		t.Status = "requesting"
	}

	t.Startingfromlocationid = origin
	t.Nextdestinationlocationid = next_location
	t.Bestroutelocationids = nil
	t.Locationids = result.Locationids
	t.Totalubercosts = result.Totalubercosts
	t.Totaluberduration = result.Totaluberduration
	t.Totaldistance = result.Totaldistance
	t.Uberwaittime = makerequest.Eta

	Updateindb(t)
	return t
}
func FindinLocationService(id int) Location {
	result := Location{}
	session := ConnectDB()
	s := session.DB("tests").C("locations")
	err := s.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
func getNewLocations(location []string, Todel int) []string {
	var newLocations []string
	for i := 0; i < len(location); i++ {
		locationid, _ := strconv.Atoi(location[i])
		if locationid != Todel {
			newLocations = append(newLocations, location[i])
		}
	}
	return newLocations
}
