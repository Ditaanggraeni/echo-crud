package repository

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ProdukRepository connects entity.Produk with database.
type ProdukRepository struct {
	db *gorm.DB
}

// NewProdukRepository creates an instance of RoleRepository.
func NewProdukRepository(db *gorm.DB) *ProdukRepository {
	return &ProdukRepository{
		db: db,
	}
}

// Insert inserts Produk data to database.
func (repo *ProdukRepository) Insert(ctx context.Context, ent *entity.Produk) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Produk{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[ProdukRepository-Insert]")
	}
	return nil
}

func (repo *ProdukRepository) GetListProduk(ctx context.Context, limit, offset string) ([]*entity.Produk, error) {
	var models []*entity.Produk
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Produk{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ProdukRepository-FindAll]")
	}
	return models, nil
}

func (repo *ProdukRepository) GetDetailProduk(ctx context.Context, Id uuid.UUID) (*entity.Produk, error) {
	var models *entity.Produk
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Produk{}).
		Take(&models, Id).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ProdukRepository-FindById]")
	}
	return models, nil

}

func (repo *ProdukRepository) UpdateProduk(ctx context.Context, ent *entity.Produk) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Produk{Id: ent.Id}).
		Select("kode_produk", "nama_produk", "harga", "stok").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[ProdukRepository-Update]")
	}
	return nil
}

func (repo *ProdukRepository) DeleteProduk(ctx context.Context, Id uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Produk{Id: Id}).Error; err != nil {
		return errors.Wrap(err, "[ProdukRepository-Delete]")
	}
	return nil
}
