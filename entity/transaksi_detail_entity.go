package entity

import "github.com/google/uuid"

const (
	TransaksiDetailTableName = "transaksi_detail"
)

//TransaksiModel is a model for entity.TransaksiDetail
type TransaksiDetail struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Produk      string     `gorm:"type:string;null" json:"produk"`
	Kuantitas      int64     `gorm:"type:int;null" json:"kuantitas"`
	Total      int64     `gorm:"type:int;null" json:"total"`
	//Auditable
}

func NewTransaksiDetail(id uuid.UUID, produk string, kuantitas, total int) *TransaksiDetail {
	return &TransaksiDetail{
		ID:         id,
		Produk: produk,
		Kuantitas:       int64(kuantitas),
		Total:       int64(total),
		//Auditable:  NewAuditable(),
	}
}

func (model *TransaksiDetail) TableName() string {
	return TransaksiDetailTableName
}
