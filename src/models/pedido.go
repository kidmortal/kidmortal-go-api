package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Pedido struct para pedido
type Pedido struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" `
	Numero     int                `bson:"numero,omitempty" json:"numero,omitempty" validate:"required"`
	Data       time.Time          `bson:"data,omitempty" json:"data,omitempty" validate:"required"`
	Nome       string             `bson:"nome,omitempty" json:"nome,omitempty" validate:"required"`
	Cnpj       string             `bson:"cnpj,omitempty" json:"cnpj,omitempty" validate:"required"`
	Valor      float64            `bson:"valor,omitempty" json:"valor,omitempty" validate:"required"`
	Cidade     string             `bson:"cidade,omitempty" json:"cidade,omitempty" `
	Estado     string             `bson:"estado,omitempty" json:"estado,omitempty"`
	Status     string             `bson:"status,omitempty" json:"status,omitempty"`
	DataStatus time.Time          `bson:"dataStatus,omitempty" json:"dataStatus,omitempty"`
	Vendedor   string             `bson:"vendedor,omitempty" json:"vendedor,omitempty"`
}
