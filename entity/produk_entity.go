package entity

import (
	"github.com/google/uuid"
)

const (
	ProdukTableName = "produk"
)

// ArticleModel is a model for entity.Article
type Produk struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	KodeProduk        string    `gorm:"type:varchar;not_null" json:"kode_produk"`
	NamaProduk       string    `gorm:"type:varchar;not_null" json:"nama_produk"`
	Harga int    `gorm:"type:integer;not_null" json:"harga"`
	Stok    int64     `gorm:"type:integer;not_null" json:"stok"`
	//Auditable
}

func NewProduk(id uuid.UUID, kode_produk, nama_produk string, harga int, stok int64) *Produk {
	return &Produk{
		Id:          id,
		KodeProduk:  kode_produk,
		NamaProduk:  nama_produk,
		Harga: int(harga),
		Stok:    int64(stok),
		//Auditable:  NewAuditable(),
	}
}

// TableName specifies table name for ArticleModel.
func (model *Produk) TableName() string {
	return ProdukTableName
}
