package repository

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// PembayaranRepository connects entity.Pembayaran with database.
type PembayaranRepository struct {
	db *gorm.DB
}

// NewPembayaranRepository creates an instance of RoleRepository.
func NewPembayaranRepository(db *gorm.DB) *PembayaranRepository {
	return &PembayaranRepository{
		db: db,
	}
}

// Insert inserts Pembayaran data to database.
func (repo *PembayaranRepository) Insert(ctx context.Context, ent *entity.Pembayaran) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Pembayaran{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[PembayaranRepository-Insert]")
	}
	return nil
}

func (repo *PembayaranRepository) GetListPembayaran(ctx context.Context, limit, offset string) ([]*entity.Pembayaran, error) {
	var models []*entity.Pembayaran
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Pembayaran{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[PembayaranRepository-FindAll]")
	}
	return models, nil
}

func (repo *PembayaranRepository) GetDetailPembayaran(ctx context.Context, Id uuid.UUID) (*entity.Pembayaran, error) {
	var models *entity.Pembayaran
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Pembayaran{}).
		Take(&models, Id).
		Error; err != nil {
		return nil, errors.Wrap(err, "[PembayaranRepository-FindById]")
	}
	return models, nil

}

func (repo *PembayaranRepository) UpdatePembayaran(ctx context.Context, ent *entity.Pembayaran) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Pembayaran{Id: ent.Id}).
		Select("tgl_bayar", "total").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[PembayaranRepository-Update]")
	}
	return nil
}

func (repo *PembayaranRepository) DeletePembayaran(ctx context.Context, Id uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Pembayaran{Id: Id}).Error; err != nil {
		return errors.Wrap(err, "[PembayaranRepository-Delete]")
	}
	return nil
}
