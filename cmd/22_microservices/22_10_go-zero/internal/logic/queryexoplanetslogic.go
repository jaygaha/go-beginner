package logic

import (
	"context"

	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_10_go-zero/internal/svc"
	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_10_go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryExoplanetsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryExoplanetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryExoplanetsLogic {
	return &QueryExoplanetsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ExoplanetData struct {
	Name         string
	DistanceLy   int64
	Habitability float64
}

func (l *QueryExoplanetsLogic) QueryExoplanets(req *types.ExoplanetQueryRequest) (resp *types.ExoplanetQueryResponse, err error) {
	// Sample exoplanet data
	exoplanets := []ExoplanetData{
		{Name: "Kepler-442b", DistanceLy: 1200, Habitability: 0.84},
		{Name: "Proxima Centauri b", DistanceLy: 4, Habitability: 0.65},
		{Name: "TRAPPIST-1e", DistanceLy: 40, Habitability: 0.77},
	}

	// Filter exoplanets based on request criteria
	var filtered []types.Exoplanet
	for _, ep := range exoplanets {
		if ep.DistanceLy <= req.MaxDistanceLy && ep.Habitability >= req.MinHabitability {
			filtered = append(filtered, types.Exoplanet{
				Name:         ep.Name,
				DistanceLy:   ep.DistanceLy,
				Habitability: ep.Habitability,
			})
		}
	}

	return &types.ExoplanetQueryResponse{
		Exoplanets: filtered,
	}, nil
}
