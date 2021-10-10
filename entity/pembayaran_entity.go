package entity

import (
	"github.com/google/uuid"
)

const (
	PembayaranTableName = "pembayaran"
)

// ArticleModel is a model for entity.Article
type Pembayaran struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	TglBayar        string    `gorm:"type:date;not_null" json:"tgl_bayar"`
	Total int64    `gorm:"type:integer;not_null" json:"total"`
	//Auditable
}

func NewPembayaran(id uuid.UUID, tgl_bayar string, total int64) *Pembayaran {
	return &Pembayaran{
		Id:          id,
		TglBayar:  tgl_bayar,
		Total:    int64(total),
		//Auditable:  NewAuditable(),
	}
}

// TableName specifies table name for ArticleModel.
func (model *Pembayaran) TableName() string {
	return PembayaranTableName
}
