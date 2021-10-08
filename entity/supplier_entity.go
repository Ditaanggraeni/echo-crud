package entity

import "github.com/google/uuid"

const (
	SupplierTableName = "supplier"
)

//SupplierModel is a model for entity.Supplier
type Supplier struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"id_supplier"`
	NamaSupplier string    `gorm:"type:varchar(200);not_null" json:"nama_supplier"`
	Telepon      string    `gorm:"type:varchar(200);null" json:"telepon"`
	Alamat       string    `gorm:"type:varchar(200);null" json:"alamat"`
	//Auditable
}

func NewSupplier(id uuid.UUID, nama_supplier, telepon, alamat string) *Supplier {
	return &Supplier{
		ID:           id,
		NamaSupplier: nama_supplier,
		Telepon:      telepon,
		Alamat:       alamat,
		//Auditable:    NewAuditable(),
	}
}

func (model *Supplier) TableName() string {
	return SupplierTableName
}
