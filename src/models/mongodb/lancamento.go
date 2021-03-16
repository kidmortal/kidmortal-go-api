package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Lancamento struct para Lancamento
type Lancamento struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Cliente   primitive.ObjectID `bson:"cliente" json:"Cliente,omitempty"`
	Valor     float64            `bson:"nome" json:"nome,omitempty"`
	Descricao string             `bson:"condicaoNF" json:"condicaoNF,omitempty"`
	Data      time.Time          `bson:"cnpj" json:"cnpj,omitempty"`
	CC        bool               `bson:"CC" json:"CC,omitempty"`
}
