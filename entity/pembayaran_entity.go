package entity

import (
	"github.com/gofrs/uuid"
)

const (
	PembayaranTableName = "pembayaran"
)

// ArticleModel is a model for entity.Article
type Pembayaran struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	TglBayar        string    `gorm:"type:date;not_null" json:"tgl_bayar"`
	Total int64    `gorm:"type:integer;not_null" json:"total"`
	TransaksiID uuid.UUID `sql:"type:uuid REFERENCES transaksis(id)"`
	Transaksi  Transaksi `gorm:"Foreignkey:TransaksiID;association_foreignkey:id_transaksi"`
	//Auditable
}

func NewPembayaran(id uuid.UUID, tgl_bayar string, total int64, transaksi_id uuid.UUID) *Pembayaran {
	return &Pembayaran{
		Id:          id,
		TglBayar:  tgl_bayar,
		Total:    int64(total),
		TransaksiID:       transaksi_id,
		//Auditable:  NewAuditable(),
	}
}

// TableName specifies table name for ArticleModel.
func (model *Pembayaran) TableName() string {
	return PembayaranTableName
}
