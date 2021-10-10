package repository

import (
	"context"
	"echo-crud/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// SupplierRepository connects entity.Supplier with database.
type SupplierRepository struct {
	db *gorm.DB
}

// NewSupplierRepository creates an instance of RoleRepository.
func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

// Insert inserts supplier data to database.
func (repo *SupplierRepository) Insert(ctx context.Context, ent *entity.Supplier) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Supplier{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[SupplierRepository-Insert]")
	}
	return nil
}

func (repo *SupplierRepository) GetListSupplier(ctx context.Context, limit, offset string) ([]*entity.Supplier, error) {
	var models []*entity.Supplier
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Supplier{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[SupplierRepository-FindAll]")
	}
	return models, nil
}

func (repo *SupplierRepository) GetDetailSupplier(ctx context.Context, ID uuid.UUID) (*entity.Supplier, error) {
	var models *entity.Supplier
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Supplier{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[SupplierRepository-FindById]")
	}
	return models, nil

}

func (repo *SupplierRepository) UpdateSupplier(ctx context.Context, ent *entity.Supplier) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Supplier{ID: ent.ID}).
		Select("nama_supplier", "telepon", "alamat").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[SupplierRepository-Update]")
	}
	return nil
}

func (repo *SupplierRepository) DeleteSupplier(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Supplier{ID: ID}).Error; err != nil {
		return errors.Wrap(err, "[SupplierRepository-Delete]")
	}
	return nil
}
