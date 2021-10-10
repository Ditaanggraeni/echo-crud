package service

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilTransaksi occurs when a nil transaksi is passed.
	ErrNilDetailTransaksi = errors.New("DetailTransaksi is nil")
)

// TransaksiService responsible for any flow related to transaksi.
// It also implements TransaksiService.
type DetailTransaksiService struct {
	transaksiDetailRepo DetailTransaksiRepository
}

func NewDetailTransaksiService(transaksiDetailRepo DetailTransaksiRepository) *DetailTransaksiService {
	return &DetailTransaksiService{
		transaksiDetailRepo: transaksiDetailRepo,
	}
}

type DetailTransaksiUseCase interface {
	Create(ctx context.Context, transaksiDetail *entity.TransaksiDetail) error
	GetListTransaksi_Detail(ctx context.Context, limit, offset string) ([]*entity.TransaksiDetail, error)
	GetDetailTransaksi_Detail(ctx context.Context, ID uuid.UUID) (*entity.TransaksiDetail, error)
	UpdateDetailTransaksi(ctx context.Context, transaksiDetail *entity.TransaksiDetail) error
	DeleteDetailTransaksi(ctx context.Context, ID uuid.UUID) error
}

type DetailTransaksiRepository interface {
	Insert(ctx context.Context, transaksiDetail *entity.TransaksiDetail) error
	GetListTransaksi_Detail(ctx context.Context, limit, offset string) ([]*entity.TransaksiDetail, error)
	GetDetailTransaksi_Detail(ctx context.Context, ID uuid.UUID) (*entity.TransaksiDetail, error)
	UpdateDetailTransaksi(ctx context.Context, transaksiDetail *entity.TransaksiDetail) error
	DeleteDetailTransaksi(ctx context.Context, ID uuid.UUID) error
}

func (svc DetailTransaksiService) Create(ctx context.Context, transaksiDetail *entity.TransaksiDetail) error {
	// Checking nil file
	if transaksiDetail == nil {
		return ErrNilDetailTransaksi
	}

	// Generate id if nil
	if transaksiDetail.ID == uuid.Nil {
		transaksiDetail.ID = uuid.New()
	}

	if err := svc.transaksiDetailRepo.Insert(ctx, transaksiDetail); err != nil {
		return errors.Wrap(err, "[DetailTransaksiService-Create]")
	}
	return nil
}

func (svc DetailTransaksiService) GetListTransaksi_Detail(ctx context.Context, limit, offset string) ([]*entity.TransaksiDetail, error) {
	transaksiDetail, err := svc.transaksiDetailRepo.GetListTransaksi_Detail(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[DetailTransaksiService-Create]")
	}
	return transaksiDetail, nil
}

func (svc DetailTransaksiService) GetDetailTransaksi_Detail(ctx context.Context, ID uuid.UUID) (*entity.TransaksiDetail, error) {
	transaksiDetail, err := svc.transaksiDetailRepo.GetDetailTransaksi_Detail(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[DetailTransaksiService-Create]")
	}
	return transaksiDetail, nil
}

func (svc DetailTransaksiService) UpdateDetailTransaksi(ctx context.Context, transaksiDetail *entity.TransaksiDetail) error {
	// Checking nil transaksi
	if transaksiDetail == nil {
		return ErrNilDetailTransaksi
	}

	// Generate id if nil
	if transaksiDetail.ID == uuid.Nil {
		transaksiDetail.ID = uuid.New()
	}

	if err := svc.transaksiDetailRepo.UpdateDetailTransaksi(ctx, transaksiDetail); err != nil {
		return errors.Wrap(err, "[DetailTransaksiService-Create]")
	}
	return nil
}

func (svc DetailTransaksiService) DeleteDetailTransaksi(ctx context.Context, ID uuid.UUID) error {
	err := svc.transaksiDetailRepo.DeleteDetailTransaksi(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[DetailTransaksiService-Create]")
	}
	return nil
}
