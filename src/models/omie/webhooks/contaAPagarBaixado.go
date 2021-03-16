package models

// ContaAPagarBaixado estrutura do webhook
type ContaAPagarBaixado struct {
	Messageid string `json:"messageId"`
	Topic     string `json:"topic"`
	Event     struct {
		CodigoBaixa           int    `json:"codigo_baixa"`
		CodigoBaixaIntegracao string `json:"codigo_baixa_integracao"`
		CodigoContaCorrente   int    `json:"codigo_conta_corrente"`
		ContaAPagar           []struct {
			CodigoLancamentoIntegracao string  `json:"codigo_lancamento_integracao"`
			CodigoLancamentoOmie       int     `json:"codigo_lancamento_omie"`
			DataEmissao                string  `json:"data_emissao"`
			DataEntrada                string  `json:"data_entrada"`
			DataVencimento             string  `json:"data_vencimento"`
			NumeroDocumento            string  `json:"numero_documento"`
			NumeroDocumentoFiscal      string  `json:"numero_documento_fiscal"`
			NumeroParcela              string  `json:"numero_parcela"`
			ValorDocumento             float64 `json:"valor_documento"`
		} `json:"conta_a_pagar"`
		Data       string  `json:"data"`
		DataCred   string  `json:"data_cred"`
		Desconto   int     `json:"desconto"`
		Juros      int     `json:"juros"`
		Multa      int     `json:"multa"`
		Observacao string  `json:"observacao"`
		Tarifa     int     `json:"tarifa"`
		Valor      float64 `json:"valor"`
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
