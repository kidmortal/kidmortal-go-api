package models

// ClienteFornecedorAlterado estrutura do webhook
type ClienteFornecedorAlterado struct {
	Messageid string `json:"messageId"`
	Topic     string `json:"topic"`
	Event     struct {
		Bairro                  string `json:"bairro"`
		Bloqueado               string `json:"bloqueado"`
		BloquearFaturamento     string `json:"bloquear_faturamento"`
		Cep                     string `json:"cep"`
		Cidade                  string `json:"cidade"`
		Cnae                    string `json:"cnae"`
		CnpjCpf                 string `json:"cnpj_cpf"`
		CodigoClienteIntegracao string `json:"codigo_cliente_integracao"`
		CodigoClienteOmie       int    `json:"codigo_cliente_omie"`
		CodigoPais              string `json:"codigo_pais"`
		Complemento             string `json:"complemento"`
		Contato                 string `json:"contato"`
		Contribuinte            string `json:"contribuinte"`
		Dadosbancarios          struct {
			Agencia       string `json:"agencia"`
			CodigoBanco   string `json:"codigo_banco"`
			ContaCorrente string `json:"conta_corrente"`
			DocTitular    string `json:"doc_titular"`
			NomeTitular   string `json:"nome_titular"`
		} `json:"dadosBancarios"`
		Email                  string `json:"email"`
		Endereco               string `json:"endereco"`
		EnderecoNumero         string `json:"endereco_numero"`
		Estado                 string `json:"estado"`
		FaxDdd                 string `json:"fax_ddd"`
		FaxNumero              string `json:"fax_numero"`
		Homepage               string `json:"homepage"`
		Inativo                string `json:"inativo"`
		InscricaoEstadual      string `json:"inscricao_estadual"`
		InscricaoMunicipal     string `json:"inscricao_municipal"`
		InscricaoSuframa       string `json:"inscricao_suframa"`
		Logradouro             string `json:"logradouro"`
		Nif                    string `json:"nif"`
		NomeFantasia           string `json:"nome_fantasia"`
		Observacao             string `json:"observacao"`
		OptanteSimplesNacional string `json:"optante_simples_nacional"`
		PessoaFisica           string `json:"pessoa_fisica"`
		ProdutorRural          string `json:"produtor_rural"`
		RazaoSocial            string `json:"razao_social"`
		RecomendacaoAtraso     string `json:"recomendacao_atraso"`
		Recomendacoes          struct {
			CodigoVendedor int    `json:"codigo_vendedor"`
			EmailFatura    string `json:"email_fatura"`
			GerarBoletos   string `json:"gerar_boletos"`
			NumeroParcelas string `json:"numero_parcelas"`
		} `json:"recomendacoes"`
		Tags []struct {
			Tag string `json:"tag"`
		} `json:"tags"`
		Telefone1Ddd       string `json:"telefone1_ddd"`
		Telefone1Numero    string `json:"telefone1_numero"`
		Telefone2Ddd       string `json:"telefone2_ddd"`
		Telefone2Numero    string `json:"telefone2_numero"`
		TipoAtividade      string `json:"tipo_atividade"`
		ValorLimiteCredito string `json:"valor_limite_credito"`
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
