package service

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilPelanggan occurs when a nil pelanggan is passed.
	ErrNilPelanggan = errors.New("Pelanggan is nil")
)

// PelangganService responsible for any flow related to pelanggan.
// It also implements PelangganService.
type PelangganService struct {
	PelangganRepo PelangganRepository
}

func NewPelangganService(PelangganRepo PelangganRepository) *PelangganService {
	return &PelangganService{
		PelangganRepo: PelangganRepo,
	}
}

type PelangganUseCase interface {
	Create(ctx context.Context, pelanggan *entity.Pelanggan) error
	GetListPelanggan(ctx context.Context, limit, offset string) ([]*entity.Pelanggan, error)
	GetDetailPelanggan(ctx context.Context, ID uuid.UUID) (*entity.Pelanggan, error)
	UpdatePelanggan(ctx context.Context, pelanggan *entity.Pelanggan) error
	DeletePelanggan(ctx context.Context, ID uuid.UUID) error
}

type PelangganRepository interface {
	Insert(ctx context.Context, pelanggan *entity.Pelanggan) error
	GetListPelanggan(ctx context.Context, limit, offset string) ([]*entity.Pelanggan, error)
	GetDetailPelanggan(ctx context.Context, ID uuid.UUID) (*entity.Pelanggan, error)
	UpdatePelanggan(ctx context.Context, pelanggan *entity.Pelanggan) error
	DeletePelanggan(ctx context.Context, ID uuid.UUID) error
}

func (svc PelangganService) Create(ctx context.Context, pelanggan *entity.Pelanggan) error {
	// Checking nil file
	if pelanggan == nil {
		return ErrNilPelanggan
	}

	// Generate id if nil
	if pelanggan.ID == uuid.Nil {
		pelanggan.ID = uuid.New()
	}

	if err := svc.PelangganRepo.Insert(ctx, pelanggan); err != nil {
		return errors.Wrap(err, "[PelangganService-Create]")
	}
	return nil
}

func (svc PelangganService) GetListPelanggan(ctx context.Context, limit, offset string) ([]*entity.Pelanggan, error) {
	pelanggan, err := svc.PelangganRepo.GetListPelanggan(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[PelangganService-Create]")
	}
	return pelanggan, nil
}

func (svc PelangganService) GetDetailPelanggan(ctx context.Context, ID uuid.UUID) (*entity.Pelanggan, error) {
	pelanggan, err := svc.PelangganRepo.GetDetailPelanggan(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[PelangganrService-Create]")
	}
	return pelanggan, nil
}

func (svc PelangganService) UpdatePelanggan(ctx context.Context, pelanggan *entity.Pelanggan) error {
	// Checking nil pelanggan
	if pelanggan == nil {
		return ErrNilPelanggan
	}

	// Generate id if nil
	if pelanggan.ID == uuid.Nil {
		pelanggan.ID = uuid.New()
	}

	if err := svc.PelangganRepo.UpdatePelanggan(ctx, pelanggan); err != nil {
		return errors.Wrap(err, "[PelangganService-Create]")
	}
	return nil
}

func (svc PelangganService) DeletePelanggan(ctx context.Context, ID uuid.UUID) error {
	err := svc.PelangganRepo.DeletePelanggan(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[PelangganService-Create]")
	}
	return nil
}
