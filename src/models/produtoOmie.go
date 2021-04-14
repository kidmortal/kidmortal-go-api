package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ProdutoOmie estrutura de retorno de um produto do OMIE
type ProdutoOmie struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Nome       string             `bson:"nome" json:"nome,omitempty"`
	CondicaoNF string             `bson:"condicaoNF" json:"condicaoNF,omitempty"`
	Cnpj       []string           `bson:"cnpj" json:"cnpj,omitempty"`
}
