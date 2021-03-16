package models

// PedidoStatusAlterado BLING - estrutura do webhook de alteração de status de um pedido

type PedidoBling struct {
	Desconto          string       `json:"desconto"`
	Observacoes       string       `json:"observacoes"`
	Observacaointerna string       `json:"observacaointerna"`
	Data              string       `json:"data"`
	Numero            string       `json:"numero"`
	Numeroordemcompra string       `json:"numeroOrdemCompra"`
	Vendedor          string       `json:"vendedor"`
	Valorfrete        string       `json:"valorfrete"`
	Totalprodutos     string       `json:"totalprodutos"`
	Totalvenda        string       `json:"totalvenda"`
	Situacao          string       `json:"situacao"`
	Datasaida         string       `json:"dataSaida"`
	Loja              string       `json:"loja"`
	Numeropedidoloja  string       `json:"numeroPedidoLoja"`
	Tipointegracao    string       `json:"tipoIntegracao"`
	Cliente           ClienteBling `json:"cliente"`
}

type ClienteBling struct {
	ID          string      `json:"id"`
	Nome        string      `json:"nome"`
	Cnpj        string      `json:"cnpj"`
	Ie          string      `json:"ie"`
	Rg          string      `json:"rg"`
	Endereco    string      `json:"endereco"`
	Numero      string      `json:"numero"`
	Complemento string      `json:"complemento"`
	Cidade      string      `json:"cidade"`
	Bairro      string      `json:"bairro"`
	Cep         string      `json:"cep"`
	Uf          string      `json:"uf"`
	Email       string      `json:"email"`
	Celular     interface{} `json:"celular"`
	Fone        string      `json:"fone"`
}

type PedidoStatusAlterado struct {
	Retorno struct {
		Pedidos []struct {
			Pedido PedidoBling `json:"pedido"`
		} `json:"pedidos"`
	} `json:"retorno"`
}
