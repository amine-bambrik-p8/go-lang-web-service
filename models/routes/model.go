package routes

// Represents the Information about the Route
type RouteInfo struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

// Represents all possible Routes starting from source
type AllRoutes struct {
	Source string      `json:"source"`
	Routes []RouteInfo `json:"routes"`
}
