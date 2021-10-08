package repository

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// PelangganRepository connects entity.Pelanggan with database.
type PelangganRepository struct {
	db *gorm.DB
}

// NewPelangganRepository creates an instance of RoleRepository.
func NewPelangganRepository(db *gorm.DB) *PelangganRepository {
	return &PelangganRepository{
		db: db,
	}
}

// Insert inserts pelanggan data to database.
func (repo *PelangganRepository) Insert(ctx context.Context, ent *entity.Pelanggan) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Pelanggan{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[PelangganRepository-Insert]")
	}
	return nil
}

func (repo *PelangganRepository) GetListPelanggan(ctx context.Context, limit, offset string) ([]*entity.Pelanggan, error) {
	var models []*entity.Pelanggan
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Pelanggan{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[PelangganRepository-FindAll]")
	}
	return models, nil
}

func (repo *PelangganRepository) GetDetailPelanggan(ctx context.Context, ID uuid.UUID) (*entity.Pelanggan, error) {
	var models *entity.Pelanggan
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Pelanggan{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[PelangganRepository-FindById]")
	}
	return models, nil

}

func (repo *PelangganRepository) UpdatePelanggan(ctx context.Context, ent *entity.Pelanggan) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Pelanggan{ID: ent.ID}).
		Select("name", "kode", "description", "quantity", "price").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[PelangganRepository-Update]")
	}
	return nil
}

func (repo *PelangganRepository) DeletePelanggan(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Pelanggan{ID: ID}).Error; err != nil {
		return errors.Wrap(err, "[PelangganRepository-Delete]")
	}
	return nil
}
