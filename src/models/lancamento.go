package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Lancamento struct para Lancamento
type Lancamento struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Cliente   primitive.ObjectID `bson:"cliente,omitempty" json:"Cliente,omitempty"`
	Valor     float64            `bson:"nome,omitempty" json:"nome,omitempty"`
	Descricao string             `bson:"condicaoNF,omitempty" json:"condicaoNF,omitempty"`
	Data      time.Time          `bson:"cnpj,omitempty" json:"cnpj,omitempty"`
	CC        bool               `bson:"CC,omitempty" json:"CC,omitempty"`
}
