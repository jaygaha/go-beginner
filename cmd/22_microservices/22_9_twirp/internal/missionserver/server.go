package missionserver

import (
	"context"
	"fmt"

	"sync"
	"time"

	pb "github.com/jaygaha/go-beginner/cmd/22_microservices/22_9_twirp/rpc/mission"
	"github.com/twitchtv/twirp"
)

// Planet holds data for a planet.
type Planet struct {
	Name       string
	DistanceAU float64 // Distance from Earth in Astronomical Units
}

// Mission represents a planned space mission.
type Mission struct {
	ID             string
	PlanetName     string
	Spacecraft     string
	LaunchDate     time.Time
	TravelTimeDays int64
}

// MissionServiceServer implements the MissionService interface.
type MissionServiceServer struct {
	planets  map[string]Planet
	missions map[string]Mission
	mu       sync.Mutex
}

// NewMissionServiceServer initializes the server with sample planets.
func NewMissionServiceServer() *MissionServiceServer {
	return &MissionServiceServer{
		planets: map[string]Planet{
			"Mars":    {Name: "Mars", DistanceAU: 0.52},
			"Jupiter": {Name: "Jupiter", DistanceAU: 4.2},
			"Saturn":  {Name: "Saturn", DistanceAU: 8.5},
		},
		missions: make(map[string]Mission),
	}
}

// PlanMission handles the PlanMission RPC.
func (s *MissionServiceServer) PlanMission(ctx context.Context, req *pb.PlanMissionRequest) (*pb.PlanMissionResponse, error) {
	// Validate required fields
	if req.PlanetName == "" {
		return nil, twirp.RequiredArgumentError("planet_name")
	}

	// Check if planet exists
	s.mu.Lock()
	planet, exists := s.planets[req.PlanetName]
	if !exists {
		s.mu.Unlock()
		return nil, twirp.InvalidArgumentError("planet_name", fmt.Sprintf("unknown planet: %s", req.PlanetName))
	}

	// Simulate mission planning
	missionID := generateID()
	launchDate := time.Now().AddDate(0, 1, 0) // Launch in 1 month
	// Simplified travel time: assume 200 km/s speed, 1 AU = 149.6M km
	travelTimeDays := int64(planet.DistanceAU * 149600000 / 200 / 86400)

	mission := Mission{
		ID:             missionID,
		PlanetName:     req.PlanetName,
		Spacecraft:     req.Spacecraft,
		LaunchDate:     launchDate,
		TravelTimeDays: travelTimeDays,
	}
	s.missions[missionID] = mission
	s.mu.Unlock()

	return &pb.PlanMissionResponse{
		MissionId:      mission.ID,
		PlanetName:     mission.PlanetName,
		Spacecraft:     mission.Spacecraft,
		LaunchDate:     mission.LaunchDate.Format(time.RFC3339),
		TravelTimeDays: mission.TravelTimeDays,
	}, nil
}

// generateID creates a unique mission ID.
func generateID() string {
	return fmt.Sprintf("mission-%d", time.Now().UnixNano())
}
