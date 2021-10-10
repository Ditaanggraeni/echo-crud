package service

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilProduk occurs when a nil Produk is passed.
	ErrNilProduk = errors.New("Produk is nil")
)

// ProdukService responsible for any flow related to Produk.
// It also implements ProdukService.
type ProdukService struct {
	ProdukRepo ProdukRepository
}

func NewProdukService(ProdukRepo ProdukRepository) *ProdukService {
	return &ProdukService{
		ProdukRepo: ProdukRepo,
	}
}

type ProdukUseCase interface {
	Create(ctx context.Context, produk *entity.Produk) error
	GetListProduk(ctx context.Context, limit, offset string) ([]*entity.Produk, error)
	GetDetailProduk(ctx context.Context, Id uuid.UUID) (*entity.Produk, error)
	UpdateProduk(ctx context.Context, Produk *entity.Produk) error
	DeleteProduk(ctx context.Context, Id uuid.UUID) error
}

type ProdukRepository interface {
	Insert(ctx context.Context, produk *entity.Produk) error
	GetListProduk(ctx context.Context, limit, offset string) ([]*entity.Produk, error)
	GetDetailProduk(ctx context.Context, Id uuid.UUID) (*entity.Produk, error)
	UpdateProduk(ctx context.Context, Produk *entity.Produk) error
	DeleteProduk(ctx context.Context, Id uuid.UUID) error
}

func (svc ProdukService) Create(ctx context.Context, produk *entity.Produk) error {
	// Checking nil file
	if produk == nil {
		return ErrNilProduk
	}

	// Generate id if nil
	if produk.Id == uuid.Nil {
		produk.Id = uuid.New()
	}

	if err := svc.ProdukRepo.Insert(ctx, produk); err != nil {
		return errors.Wrap(err, "[ProdukService-Create]")
	}
	return nil
}

func (svc ProdukService) GetListProduk(ctx context.Context, limit, offset string) ([]*entity.Produk, error) {
	produk, err := svc.ProdukRepo.GetListProduk(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[ProdukService-Create]")
	}
	return produk, nil
}

func (svc ProdukService) GetDetailProduk(ctx context.Context, Id uuid.UUID) (*entity.Produk, error) {
	produk, err := svc.ProdukRepo.GetDetailProduk(ctx, Id)
	if err != nil {
		return nil, errors.Wrap(err, "[ProdukrService-Create]")
	}
	return produk, nil
}

func (svc ProdukService) UpdateProduk(ctx context.Context, produk *entity.Produk) error {
	// Checking nil Produk
	if produk == nil {
		return ErrNilProduk
	}

	// Generate id if nil
	if produk.Id == uuid.Nil {
		produk.Id = uuid.New()
	}

	if err := svc.ProdukRepo.UpdateProduk(ctx, produk); err != nil {
		return errors.Wrap(err, "[ProdukService-Create]")
	}
	return nil
}

func (svc ProdukService) DeleteProduk(ctx context.Context, Id uuid.UUID) error {
	err := svc.ProdukRepo.DeleteProduk(ctx, Id)
	if err != nil {
		return errors.Wrap(err, "[ProdukService-Create]")
	}
	return nil
}
