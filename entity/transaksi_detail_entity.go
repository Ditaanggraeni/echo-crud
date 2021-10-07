package entity

import "github.com/gofrs/uuid"

const (
	TransaksiDetailTableName = "transaksi_detail"
)

//TransaksiModel is a model for entity.TransaksiDetail
type TransaksiDetail struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	TransaksiID uuid.UUID `sql:"type:uuid REFERENCES transaksi(id)"`
	Transaksi   Transaksi `gorm:"foreign_key:TransaksiID;AssociationForeignKey:id_transaksi"`
	ProdukID uuid.UUID `sql:"type:uuid REFERENCES produk(id)"`
	Produk   Produk `gorm:"foreign_key:ProdukID;AssociationForeignKey:id_produk"`
	Kuantitas      int64     `gorm:"type:int;null" json:"kuantitas"`
	Total      int64     `gorm:"type:int;null" json:"total"`
	//Auditable
}

func NewTransaksiDetail(id, transaksi_id, produk_id uuid.UUID, kuantitas, total int) *TransaksiDetail {
	return &TransaksiDetail{
		ID:         id,
		TransaksiID: transaksi_id,
		ProdukID: produk_id,
		Kuantitas:       int64(kuantitas),
		Total:       int64(total),
		//Auditable:  NewAuditable(),
	}
}

func (model *TransaksiDetail) TableName() string {
	return TransaksiDetailTableName
}
