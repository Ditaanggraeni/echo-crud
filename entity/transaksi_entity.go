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
	Auditable
}

func NewTransaksi(id uuid.UUID, kode_produk, nama_produk string, harga, stok int) *Produk {
	return &Produk{
		ID:         id,
		KodeProduk: kode_produk,
		NamaProduk: nama_produk,
		Harga:      int64(harga),
		Stok:       int64(stok),
		Auditable:  NewAuditable(),
	}
}

func (model *Transaksi) TableName() string {
	return TransaksiTableName
}
