package routes

import "sort"

// Returns the AllRoutes to the client
func (r *AllRoutes) GetViewModel() interface{} {
	return map[string]interface{}{
		"source": r.Source,
		"routes": r.Routes,
	}
}

// Represents all possible Routes starting from source
type AllRoutes struct {
	Source string      `json:"source"`
	Routes []RouteInfo `json:"routes"`
	//Code   string      `json:"code"`
}

// Represents the Information about the Route
type RouteInfo struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

func (r *AllRoutes) SortByDuration() {
	sort.SliceStable(r.Routes, func(i, j int) bool {
		return r.Routes[i].Duration < r.Routes[j].Duration
	})
}

func (r *AllRoutes) SortByDistance() {
	sort.SliceStable(r.Routes, func(i, j int) bool {
		return r.Routes[i].Distance < r.Routes[j].Distance
	})
}
