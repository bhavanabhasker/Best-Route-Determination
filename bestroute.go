package main

import (
	"encoding/json"
	"log"
	"sort"
	"strconv"
)

type RouteEstimate struct {
	Startlocation int
	Endlocation   int
	Costs         int
	Duration      int
	Distance      float64
}

type RouteEstimates []RouteEstimate

var pe PriceEstimates
var revpe PriceEstimates
var BestRoutes = make(RouteEstimates, 0)
var Charges = make(RouteEstimates, 0)
var itr = 0

func DetermineBest(id int, startlatitude float64, startlongitude float64,
	location []string) RouteEstimates {
	if len(location) != 0 {
		BestRoute := CalculateEst(id, startlatitude, startlongitude, location)
		newLocations := getNewLocations(location, BestRoute.Endlocation)
		BestRoutes = append(BestRoutes, BestRoute)
		nextlocation := BestRoute.Endlocation
		startingloccord := FindinLocationService(nextlocation)
		startlatitude = startingloccord.Coordinates.Lat
		startlongitude = startingloccord.Coordinates.Long
		DetermineBest(nextlocation, startlatitude, startlongitude, newLocations)
	}
	return BestRoutes
}
func clearBestRoutes() {
	BestRoutes = nil
}
func Routescharges(startlocation int, startlatitude float64, startlongitude float64,
	endlocation int, endlatitude float64, endlongitude float64) {
	data := getResponse(startlatitude, startlongitude, endlatitude, endlongitude)
	if e := json.Unmarshal(data, &pe); e != nil {
		log.Fatal(e)
	}
	var route RouteEstimate
	route.Startlocation = startlocation
	route.Endlocation = endlocation
	route.Costs = pe.Prices[0].HighEstimate
	route.Duration = pe.Prices[0].Duration
	route.Distance = pe.Prices[0].Distance
	Charges = append(Charges, route)
}
func FromPrim() RouteEstimate {
	bestroute := PrimsImpl(Charges)
	Charges = nil
	return bestroute
}

func (charges RouteEstimates) Len() int {
	return len(charges)
}

func (charges RouteEstimates) Less(i, j int) bool {
	return charges[i].Costs < charges[j].Costs
}

func (charges RouteEstimates) Swap(i, j int) {
	charges[i], charges[j] = charges[j], charges[i]
}
func PrimsImpl(Charges RouteEstimates) RouteEstimate {
	sort.Sort(Charges)
	return Charges[0]

}
func CalculateEst(id int, startlatitude float64, startlongitude float64, location []string) RouteEstimate {
	for i := 0; i < len(location); i++ {
		locationid, _ := strconv.Atoi(location[i])
		result := FindinLocationService(locationid)
		endlatitude := result.Coordinates.Lat
		endlongitude := result.Coordinates.Long
		Routescharges(id, startlatitude, startlongitude,
			locationid, endlatitude, endlongitude)
	}
	BestRoute := FromPrim()
	return BestRoute
}
