package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Nota struct para Nota
type Nota struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Numero        int64              `bson:"numero,omitempty" json:"numero,omitempty"`
	Boleto        bool               `bson:"boleto,omitempty" json:"boleto,omitempty"`
	Cliente       string             `bson:"cliente,omitempty" json:"cliente,omitempty"`
	Cnpj          string             `bson:"cnpj,omitempty" json:"cnpj,omitempty"`
	Data          time.Time          `bson:"data,omitempty" json:"data,omitempty"`
	Empresa       string             `bson:"empresa,omitempty" json:"empresa,omitempty"`
	NfUrl         string             `bson:"nfUrl,omitempty" json:"nfUrl,omitempty"`
	ValorImposto  float64            `bson:"valorImposto" json:"valorImposto"`
	ValorProdutos float64            `bson:"valorProdutos" json:"valorProdutos"`
	ValorTotal    float64            `bson:"valorTotal" json:"valorTotal"`
}
