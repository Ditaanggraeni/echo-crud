package entity

import "github.com/gofrs/uuid"

const (
	TransaksiTableName = "transaksi"
)

//TransaksiModel is a model for entity.Transaksi
type Transaksi struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id_transaksi"`
	Tanggal     string    `gorm:"type:date;null" json:"tanggal"`
	Keterangan  string    `gorm:"type:text;null" json:"keterangan"`
	Total       int64     `gorm:"type:int;null" json:"total"`
	PelangganID uuid.UUID `sql:"type:uuid REFERENCES pelanggan(id)"`
	Pelanggan   Pelanggan `gorm:"foreign_key:PelangganID;AssociationForeignKey:id_pelanggan"`
	//Auditable
}

func NewTransaksi(id_transaksi uuid.UUID, tanggal, keterangan string, total int, pelanggan_id uuid.UUID) *Transaksi {
	return &Transaksi{
		ID:         id_transaksi,
		Tanggal: tanggal,
		Keterangan: keterangan,
		Total:       int64(total),
		PelangganID: pelanggan_id,
		//Auditable:  NewAuditable(),
	}
}

func (model *Transaksi) TableName() string {
	return TransaksiTableName
}
