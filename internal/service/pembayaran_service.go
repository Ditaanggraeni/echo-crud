package service

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilPembayaran occurs when a nil Pembayaran is passed.
	ErrNilPembayaran = errors.New("Pembayaran is nil")
)

// PembayaranService responsible for any flow related to Pembayaran.
// It also implements PembayaranService.
type PembayaranService struct {
	pembayaranRepo PembayaranRepository
}

func NewPembayaranService(pembayaranRepo PembayaranRepository) *PembayaranService {
	return &PembayaranService{
		pembayaranRepo: pembayaranRepo,
	}
}

type PembayaranUseCase interface {
	Create(ctx context.Context,pPembayaran *entity.Pembayaran) error
	GetListPembayaran(ctx context.Context, limit, offset string) ([]*entity.Pembayaran, error)
	GetDetailPembayaran(ctx context.Context, Id uuid.UUID) (*entity.Pembayaran, error)
	UpdatePembayaran(ctx context.Context, Pembayaran *entity.Pembayaran) error
	DeletePembayaran(ctx context.Context, Id uuid.UUID) error
}

type PembayaranRepository interface {
	Insert(ctx context.Context, pembayaran *entity.Pembayaran) error
	GetListPembayaran(ctx context.Context, limit, offset string) ([]*entity.Pembayaran, error)
	GetDetailPembayaran(ctx context.Context, Id uuid.UUID) (*entity.Pembayaran, error)
	UpdatePembayaran(ctx context.Context, Pembayaran *entity.Pembayaran) error
	DeletePembayaran(ctx context.Context, Id uuid.UUID) error
}

func (svc PembayaranService) Create(ctx context.Context, pembayaran *entity.Pembayaran) error {
	// Checking nil file
	if pembayaran == nil {
		return ErrNilPembayaran
	}

	// Generate id if nil
	if pembayaran.Id == uuid.Nil {
		pembayaran.Id = uuid.New()
	}

	if err := svc.pembayaranRepo.Insert(ctx, pembayaran); err != nil {
		return errors.Wrap(err, "[PembayaranService-Create]")
	}
	return nil
}

func (svc PembayaranService) GetListPembayaran(ctx context.Context, limit, offset string) ([]*entity.Pembayaran, error) {
	pembayaran, err := svc.pembayaranRepo.GetListPembayaran(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[PembayaranService-Create]")
	}
	return pembayaran, nil
}

func (svc PembayaranService) GetDetailPembayaran(ctx context.Context, Id uuid.UUID) (*entity.Pembayaran, error) {
	pembayaran, err := svc.pembayaranRepo.GetDetailPembayaran(ctx, Id)
	if err != nil {
		return nil, errors.Wrap(err, "[PembayaranService-Create]")
	}
	return pembayaran, nil
}

func (svc PembayaranService) UpdatePembayaran(ctx context.Context, pembayaran *entity.Pembayaran) error {
	// Checking nil Pembayaran
	if pembayaran == nil {
		return ErrNilPembayaran
	}

	// Generate id if nil
	if pembayaran.Id == uuid.Nil {
		pembayaran.Id = uuid.New()
	}

	if err := svc.pembayaranRepo.UpdatePembayaran(ctx, pembayaran); err != nil {
		return errors.Wrap(err, "[PembayaranService-Create]")
	}
	return nil
}

func (svc PembayaranService) DeletePembayaran(ctx context.Context, Id uuid.UUID) error {
	err := svc.pembayaranRepo.DeletePembayaran(ctx, Id)
	if err != nil {
		return errors.Wrap(err, "[PembayaranService-Create]")
	}
	return nil
}
