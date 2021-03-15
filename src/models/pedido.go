package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Pedido struct para pedido
type Pedido struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Numero     int32              `bson:"numero" json:"numero,omitempty"`
	Data       time.Time          `bson:"data" json:"data,omitempty"`
	Nome       string             `bson:"nome" json:"nome,omitempty"`
	Cnpj       string             `bson:"cnpj" json:"cnpj,omitempty"`
	Valor      float64            `bson:"valor" json:"valor,omitempty"`
	Cidade     string             `bson:"cidade" json:"cidade,omitempty"`
	Estado     string             `bson:"estado" json:"estado,omitempty"`
	Status     string             `bson:"status" json:"status,omitempty"`
	DataStatus time.Time          `bson:"dataStatus" json:"dataStatus,omitempty"`
	Vendedor   string             `bson:"vendedor" json:"vendedor,omitempty"`
}
