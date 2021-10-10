package service

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilTransaksi occurs when a nil transaksi is passed.
	ErrNilTransaksi = errors.New("Transaksi is nil")
)

// TransaksiService responsible for any flow related to transaksi.
// It also implements TransaksiService.
type TransaksiService struct {
	transaksiRepo TransaksiRepository
}

func NewTransaksiService(transaksiRepo TransaksiRepository) *TransaksiService {
	return &TransaksiService{
		transaksiRepo: transaksiRepo,
	}
}

type TransaksiUseCase interface {
	Create(ctx context.Context, transaksi *entity.Transaksi) error
	GetListTransaksi(ctx context.Context, limit, offset string) ([]*entity.Transaksi, error)
	GetDetailTransaksi(ctx context.Context, ID uuid.UUID) (*entity.Transaksi, error)
	UpdateTransaksi(ctx context.Context, transaksi *entity.Transaksi) error
	DeleteTransaksi(ctx context.Context, ID uuid.UUID) error
}

type TransaksiRepository interface {
	Insert(ctx context.Context, transaksi *entity.Transaksi) error
	GetListTransaksi(ctx context.Context, limit, offset string) ([]*entity.Transaksi, error)
	GetDetailTransaksi(ctx context.Context, ID uuid.UUID) (*entity.Transaksi, error)
	UpdateTransaksi(ctx context.Context, transaksi *entity.Transaksi) error
	DeleteTransaksi(ctx context.Context, ID uuid.UUID) error
}

func (svc TransaksiService) Create(ctx context.Context, transaksi *entity.Transaksi) error {
	// Checking nil file
	if transaksi == nil {
		return ErrNilTransaksi
	}

	// Generate id if nil
	if transaksi.ID == uuid.Nil {
		transaksi.ID = uuid.New()
	}

	if err := svc.transaksiRepo.Insert(ctx, transaksi); err != nil {
		return errors.Wrap(err, "[TransaksiService-Create]")
	}
	return nil
}

func (svc TransaksiService) GetListTransaksi(ctx context.Context, limit, offset string) ([]*entity.Transaksi, error) {
	transaksi, err := svc.transaksiRepo.GetListTransaksi(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[TransaksiService-Create]")
	}
	return transaksi, nil
}

func (svc TransaksiService) GetDetailTransaksi(ctx context.Context, ID uuid.UUID) (*entity.Transaksi, error) {
	transaksi, err := svc.transaksiRepo.GetDetailTransaksi(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[TransaksiService-Create]")
	}
	return transaksi, nil
}

func (svc TransaksiService) UpdateTransaksi(ctx context.Context, transaksi *entity.Transaksi) error {
	// Checking nil transaksi
	if transaksi == nil {
		return ErrNilTransaksi
	}

	// Generate id if nil
	if transaksi.ID == uuid.Nil {
		transaksi.ID = uuid.New()
	}

	if err := svc.transaksiRepo.UpdateTransaksi(ctx, transaksi); err != nil {
		return errors.Wrap(err, "[TransaksiService-Create]")
	}
	return nil
}

func (svc TransaksiService) DeleteTransaksi(ctx context.Context, ID uuid.UUID) error {
	err := svc.transaksiRepo.DeleteTransaksi(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[TransaksiService-Create]")
	}
	return nil
}
