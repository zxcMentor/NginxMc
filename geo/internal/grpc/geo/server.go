package geo

import (
	"context"
	"fmt"
	"geo/internal/grpc/grpcclients"
	"geo/internal/service"
	pbgeo "geo/protos/gen/go"
)

type Geo interface {
	GeoSearch(input string) ([]byte, error)
	GeoCode(lat, lng string) ([]byte, error)
}

type ServerGeo struct {
	pbgeo.UnimplementedGeoServiceServer
	gcl *grpcclients.ClientAuth
	geo service.GeoService
}

func (s *ServerGeo) SearchAddress(context context.Context, req *pbgeo.SearchRequest) (*pbgeo.SearchResponse, error) {

	address, err := s.geo.GeoSearch(req.Input)
	if err != nil {
		return nil, fmt.Errorf("err service:%v", err)
	}

	return &pbgeo.SearchResponse{Data: address}, nil
}

func (s *ServerGeo) GeocodeAddress(context context.Context, req *pbgeo.GeocodeRequest) (*pbgeo.GeocodeResponse, error) {
	address, err := s.geo.GeoCode(req.Lat, req.Lon)
	if err != nil {
		return nil, fmt.Errorf("err service:%v", err)
	}
	return &pbgeo.GeocodeResponse{Data: address}, nil
}
