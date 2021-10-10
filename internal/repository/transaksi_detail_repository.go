package repository

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// TransaksiRepository connects entity.Transaksi with database.
type DetailTransaksiRepository struct {
	db *gorm.DB
}

// NewTransaksiRepository creates an instance of RoleRepository.
func NewDetailTransaksiRepository(db *gorm.DB) *DetailTransaksiRepository {
	return &DetailTransaksiRepository{
		db: db,
	}
}

// Insert inserts transaksi data to database.
func (repo *DetailTransaksiRepository) Insert(ctx context.Context, ent *entity.TransaksiDetail) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.TransaksiDetail{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[DetailTransaksiRepository-Insert]")
	}
	return nil
}

func (repo *DetailTransaksiRepository) GetListTransaksi_Detail(ctx context.Context, limit, offset string) ([]*entity.TransaksiDetail, error) {
	var models []*entity.TransaksiDetail
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.TransaksiDetail{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[DetailTransaksiRepository-FindAll]")
	}
	return models, nil
}

func (repo *DetailTransaksiRepository) GetDetailTransaksi_Detail(ctx context.Context, ID uuid.UUID) (*entity.TransaksiDetail, error) {
	var models *entity.TransaksiDetail
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.TransaksiDetail{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[DetailTransaksiRepository-FindById]")
	}
	return models, nil

}

func (repo *DetailTransaksiRepository) UpdateDetailTransaksi(ctx context.Context, ent *entity.TransaksiDetail) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.TransaksiDetail{ID: ent.ID}).
		Select("produk", "kuantitas", "total").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[DetailTransaksiRepository-Update]")
	}
	return nil
}

func (repo *DetailTransaksiRepository) DeleteDetailTransaksi(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.TransaksiDetail{ID: ID}).Error; err != nil {
		return errors.Wrap(err, "[DetailTransaksiRepository-Delete]")
	}
	return nil
}
