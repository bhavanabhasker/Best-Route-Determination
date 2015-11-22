package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"PostTrip",
		"POST",
		"/trips",
		PostTrip,
	},
	Route{
		"GetTrip",
		"GET",
		"/trips/{trip_id}",
		GetTrip,
	},
	Route{
		"PutTrip",
		"PUT",
		"/trips/{trip_id}/request",
		PutTrip,
	},
}
