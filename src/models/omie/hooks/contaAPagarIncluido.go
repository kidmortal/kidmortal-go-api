package models

// ContaAPagarIncluido estrutura do webhook
type ContaAPagarIncluido struct {
	Messageid string `json:"messageId"`
	Topic     string `json:"topic"`
	Event     struct {
		BaixaBloqueada             string  `json:"baixa_bloqueada"`
		Bloqueado                  string  `json:"bloqueado"`
		BoletoGerado               string  `json:"boleto_gerado"`
		ChaveNfe                   string  `json:"chave_nfe"`
		CodigoCategoria            string  `json:"codigo_categoria"`
		CodigoClienteFornecedor    int     `json:"codigo_cliente_fornecedor"`
		CodigoCmc7Cheque           string  `json:"codigo_cmc7_cheque"`
		CodigoLancamentoIntegracao string  `json:"codigo_lancamento_integracao"`
		CodigoLancamentoOmie       int     `json:"codigo_lancamento_omie"`
		CodigoProjeto              int     `json:"codigo_projeto"`
		CodigoTipoDocumento        string  `json:"codigo_tipo_documento"`
		CodigoVendedor             int     `json:"codigo_vendedor"`
		DataEmissao                string  `json:"data_emissao"`
		DataPrevisao               string  `json:"data_previsao"`
		DataRegistro               string  `json:"data_registro"`
		DataVencimento             string  `json:"data_vencimento"`
		IDContaCorrente            int     `json:"id_conta_corrente"`
		IDOrigem                   string  `json:"id_origem"`
		Nsu                        string  `json:"nsu"`
		NumeroDocumento            string  `json:"numero_documento"`
		NumeroDocumentoFiscal      string  `json:"numero_documento_fiscal"`
		NumeroParcela              string  `json:"numero_parcela"`
		NumeroPedido               string  `json:"numero_pedido"`
		Observacao                 string  `json:"observacao"`
		Operacao                   string  `json:"operacao"`
		PixGerado                  string  `json:"pix_gerado"`
		RetemCofins                string  `json:"retem_cofins"`
		RetemCsll                  string  `json:"retem_csll"`
		RetemInss                  string  `json:"retem_inss"`
		RetemIr                    string  `json:"retem_ir"`
		RetemIss                   string  `json:"retem_iss"`
		RetemPis                   string  `json:"retem_pis"`
		Situacao                   string  `json:"situacao"`
		ValorCofins                int     `json:"valor_cofins"`
		ValorCsll                  int     `json:"valor_csll"`
		ValorDocumento             float64 `json:"valor_documento"`
		ValorInss                  int     `json:"valor_inss"`
		ValorIr                    int     `json:"valor_ir"`
		ValorIss                   int     `json:"valor_iss"`
		ValorPis                   int     `json:"valor_pis"`
	} `json:"event"`
	Author struct {
		Email  string `json:"email"`
		Name   string `json:"name"`
		Userid int    `json:"userId"`
	} `json:"author"`
	Appkey  string `json:"appKey"`
	Apphash string `json:"appHash"`
	Origin  string `json:"origin"`
}
