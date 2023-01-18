package imagesrv

import (
	"context"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/TiregeRRR/image_service/internal/controllers/imagefile"
	"github.com/TiregeRRR/image_service/internal/model"
	imagev1 "github.com/TiregeRRR/image_service/proto/image/v1"
)

type In struct {
	fx.In

	LC              fx.Lifecycle
	Logger          *zap.Logger
	ImageController *imagefile.Controller
}

type Service struct {
	imagev1.UnimplementedImageServiceServer

	Logger          *zap.Logger
	ImageController *imagefile.Controller
}

func New(in In) *Service {
	srv := &Service{
		ImageController: in.ImageController,
		Logger:          in.Logger,
	}
	return srv
}

func (s *Service) UploadImage(ctx context.Context, r *imagev1.UploadImageRequest) (*imagev1.UploadImageResponse, error) {
	t := time.Now()
	m, err := s.ImageController.UploadImage(ctx, &model.Image{
		Name:      r.GetName(),
		CreatedAt: t,
		UpdatedAt: t,
		Data:      r.GetData(),
	})
	if err != nil {
		return nil, err
	}
	return &imagev1.UploadImageResponse{
		Name:      m.Name,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
		Data:      m.Data,
	}, nil
}