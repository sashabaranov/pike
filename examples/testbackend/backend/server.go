package backend

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	UnimplementedBackendServer
	storage *PostgreStorage
}

func (s *Server) Cleanup() {
	s.storage.db.Close()
}

func (s *Server) CreateAnimal(ctx context.Context, req *CreateAnimalRequest) (*CreateAnimalResponse, error) {
	created, err := s.storage.CreateAnimal(req.Animal)
	if err != nil {
		reportError("Error in CreateAnimal", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &CreateAnimalResponse{
		Created: created,
	}, nil
}

func (s *Server) GetAnimal(ctx context.Context, req *GetAnimalRequest) (*GetAnimalResponse, error) {
	ret, err := s.storage.GetAnimal(req.Id)
	if err != nil {
		reportError("Error in GetAnimal", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &GetAnimalResponse{
		Animal: ret,
	}, nil
}

func (s *Server) UpdateAnimal(ctx context.Context, req *UpdateAnimalRequest) (*UpdateAnimalResponse, error) {
	updated, err := s.storage.UpdateAnimal(req.Updated)
	if err != nil {
		reportError("Error in UpdateAnimal", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &UpdateAnimalResponse{
		Result: updated,
	}, nil
}

func (s *Server) DeleteAnimal(ctx context.Context, req *DeleteAnimalRequest) (*DeleteAnimalResponse, error) {
	err := s.storage.DeleteAnimal(req.Id)
	if err != nil {
		reportError("Error in DeleteAnimal", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &DeleteAnimalResponse{}, nil
}
