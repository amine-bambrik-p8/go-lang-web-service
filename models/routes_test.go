package models

import "testing"

func TestSortByDuration(t *testing.T) {
	someRoutes := AllRoutes{
		Routes: []RouteInfo{
			{
				Destination: "A",
				Duration:    5555.00,
				Distance:    6666.23,
			},
			{
				Destination: "B",
				Duration:    8026.00,
				Distance:    13324.23,
			},
			{
				Destination: "C",
				Duration:    8226.00,
				Distance:    13324.23,
			},
			{
				Destination: "D",
				Duration:    4525.12,
				Distance:    6555.23,
			},
		},
	}
	expectedResult := AllRoutes{
		Routes: []RouteInfo{
			{
				Destination: "D",
				Duration:    4525.12,
				Distance:    6555.23,
			},
			{
				Destination: "A",
				Duration:    5555.00,
				Distance:    6666.23,
			},
			{
				Destination: "B",
				Duration:    8026.00,
				Distance:    13324.23,
			},
			{
				Destination: "C",
				Duration:    8226.00,
				Distance:    13324.23,
			},
		},
	}

	someRoutes.SortByDuration()
	for idx, expectedRoute := range expectedResult.Routes {
		if someRoute := someRoutes.Routes[idx]; expectedRoute.Destination != someRoute.Destination {
			t.Errorf("Expected Destination %s but found %s", expectedRoute.Destination, someRoute.Destination)
		}
	}
}
