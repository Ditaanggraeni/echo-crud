package entity

import "github.com/gofrs/uuid"

const (
	PelangganTableName = "pelanggan"
)

//PelangganModel is a model for entity.Pelanggan
type Pelanggan struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key" json:"id_pelanggan"`
	NamaPelanggan string    `gorm:"type:varchar(200);not_null" json:"nama_pelanggan"`
	Telepon       string    `gorm:"type:varchar(200);null" json:"telepon"`
	Alamat        string    `gorm:"type:varchar(200);null" json:"alamat"`

	//Auditable
}

func NewPelanggan(id uuid.UUID, nama_pelanggan, telepon, alamat string) *Pelanggan {
	return &Pelanggan{
		ID:            id,
		NamaPelanggan: nama_pelanggan,
		Telepon:       telepon,
		Alamat:        alamat,
		//Auditable:     NewAuditable(),
	}
}

func (model *Pelanggan) TableName() string {
	return PelangganTableName
}
