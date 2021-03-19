package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Nota struct para Nota
type Nota struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nome       string             `bson:"nome,omitempty" json:"nome,omitempty"`
	CondicaoNF string             `bson:"condicaoNF,omitempty" json:"condicaoNF,omitempty"`
	Cnpj       []string           `bson:"cnpj,omitempty" json:"cnpj,omitempty"`
}
