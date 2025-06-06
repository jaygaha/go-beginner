// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type Exoplanet struct {
	Name         string  `json:"name"`
	DistanceLy   int64   `json:"distance_ly"`
	Habitability float64 `json:"habitability"`
}

type ExoplanetQueryRequest struct {
	MaxDistanceLy   int64   `json:"max_distance_ly,range=[1:100000]"` // Max distance in light-years
	MinHabitability float64 `json:"min_habitability,range=[0:1]"`     // Habitability score (0-1)
}

type ExoplanetQueryResponse struct {
	Exoplanets []Exoplanet `json:"exoplanets"`
}
