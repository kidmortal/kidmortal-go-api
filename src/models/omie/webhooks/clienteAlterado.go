package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Produto struct para Produto
type Produto struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Nome       string             `bson:"nome" json:"nome,omitempty"`
	CondicaoNF string             `bson:"condicaoNF" json:"condicaoNF,omitempty"`
	Cnpj       []string           `bson:"cnpj" json:"cnpj,omitempty"`
}
