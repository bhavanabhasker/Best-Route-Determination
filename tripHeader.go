package main

type Trip struct {
	Id                        int      `json:"id"`
	Status                    string   `json:"status"`
	Startingfromlocationid    string   `json:"starting_from_location_id"`
	Nextdestinationlocationid string   `json:"next_destination_location_id,omitempty"`
	Bestroutelocationids      []string `json:"location_ids,omitempty"`
	Locationids               []string `json:"best_route_location_ids,omitempty"`
	Totalubercosts            int      `json:"total_uber_costs"`
	Totaluberduration         int      `json:"total_uber_duration"`
	Totaldistance             float64  `json:"total_distance"`
	Uberwaittime              int      `json:"uber_wait_time_eta,omitempty"`
}

type Trips []Trip
