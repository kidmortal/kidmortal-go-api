package models

// Pedido struct para pedido
type Pedido struct {
	ID       int     `sql:"id,omitempty" json:"id,omitempty"`
	Numero   int32   `sql:"numero,omitempty" json:"numero,omitempty"`
	Nome     string  `sql:"nome,omitempty" json:"nome,omitempty"`
	Cnpj     string  `sql:"cnpj,omitempty" json:"cnpj,omitempty"`
	Valor    float64 `sql:"valor,omitempty" json:"valor,omitempty"`
	Cidade   string  `sql:"cidade,omitempty" json:"cidade,omitempty"`
	Estado   string  `sql:"estado,omitempty" json:"estado,omitempty"`
	Status   string  `sql:"status,omitempty" json:"status,omitempty"`
	Vendedor string  `sql:"vendedor,omitempty" json:"vendedor,omitempty"`
}
