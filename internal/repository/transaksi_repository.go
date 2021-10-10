package repository

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// TransaksiRepository connects entity.Transaksi with database.
type TransaksiRepository struct {
	db *gorm.DB
}

// NewTransaksiRepository creates an instance of RoleRepository.
func NewTransaksiRepository(db *gorm.DB) *TransaksiRepository {
	return &TransaksiRepository{
		db: db,
	}
}

// Insert inserts transaksi data to database.
func (repo *TransaksiRepository) Insert(ctx context.Context, ent *entity.Transaksi) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Transaksi{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[TransaksiRepository-Insert]")
	}
	return nil
}

func (repo *TransaksiRepository) GetListTransaksi(ctx context.Context, limit, offset string) ([]*entity.Transaksi, error) {
	var models []*entity.Transaksi
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Transaksi{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[TransaksiRepository-FindAll]")
	}
	return models, nil
}

func (repo *TransaksiRepository) GetDetailTransaksi(ctx context.Context, ID uuid.UUID) (*entity.Transaksi, error) {
	var models *entity.Transaksi
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Transaksi{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[TransaksiRepository-FindById]")
	}
	return models, nil

}

func (repo *TransaksiRepository) UpdateTransaksi(ctx context.Context, ent *entity.Transaksi) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Transaksi{ID: ent.ID}).
		Select("tanggal", "keterangan", "total").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[TransaksiRepository-Update]")
	}
	return nil
}

func (repo *TransaksiRepository) DeleteTransaksi(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Transaksi{ID: ID}).Error; err != nil {
		return errors.Wrap(err, "[TransaksiRepository-Delete]")
	}
	return nil
}
