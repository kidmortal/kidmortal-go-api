package models

// ClienteOmie estrutura de retorno de um cliente do OMIE
type ClienteParam struct {
	CodigoClienteOmie       int    `json:"codigo_cliente_omie,omitempty"`
	CodigoClienteIntegracao string `json:"codigo_cliente_integracao,omitempty"`
}

type ClienteOmie struct {
	CodigoClienteOmie       int    `json:"codigo_cliente_omie"`
	CodigoClienteIntegracao string `json:"codigo_cliente_integracao"`
	RazaoSocial             string `json:"razao_social"`
	CnpjCpf                 string `json:"cnpj_cpf"`
	NomeFantasia            string `json:"nome_fantasia"`
	Telefone1Ddd            string `json:"telefone1_ddd"`
	Telefone1Numero         string `json:"telefone1_numero"`
	Contato                 string `json:"contato"`
	Endereco                string `json:"endereco"`
	EnderecoNumero          string `json:"endereco_numero"`
	Bairro                  string `json:"bairro"`
	Complemento             string `json:"complemento"`
	Estado                  string `json:"estado"`
	Cidade                  string `json:"cidade"`
	Cep                     string `json:"cep"`
	CodigoPais              string `json:"codigo_pais"`
	Telefone2Ddd            string `json:"telefone2_ddd"`
	Telefone2Numero         string `json:"telefone2_numero"`
	FaxDdd                  string `json:"fax_ddd"`
	FaxNumero               string `json:"fax_numero"`
	Email                   string `json:"email"`
	Homepage                string `json:"homepage"`
	InscricaoEstadual       string `json:"inscricao_estadual"`
	InscricaoMunicipal      string `json:"inscricao_municipal"`
	InscricaoSuframa        string `json:"inscricao_suframa"`
	OptanteSimplesNacional  string `json:"optante_simples_nacional"`
	TipoAtividade           string `json:"tipo_atividade"`
	Cnae                    string `json:"cnae"`
	ProdutorRural           string `json:"produtor_rural"`
	Contribuinte            string `json:"contribuinte"`
	Observacao              string `json:"observacao"`
	ObsDetalhadas           string `json:"obs_detalhadas"`
	RecomendacaoAtraso      string `json:"recomendacao_atraso"`
	Tags                    []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
	PessoaFisica        string `json:"pessoa_fisica"`
	Exterior            string `json:"exterior"`
	Logradouro          string `json:"logradouro"`
	ImportadoAPI        string `json:"importado_api"`
	Bloqueado           string `json:"bloqueado"`
	CidadeIbge          string `json:"cidade_ibge"`
	ValorLimiteCredito  int    `json:"valor_limite_credito"`
	BloquearFaturamento string `json:"bloquear_faturamento"`
	Recomendacoes       struct {
	} `json:"recomendacoes"`
	Enderecoentrega struct {
	} `json:"enderecoEntrega"`
	Nif            string `json:"nif"`
	Inativo        string `json:"inativo"`
	Dadosbancarios struct {
		Agencia       string `json:"agencia"`
		CodigoBanco   string `json:"codigo_banco"`
		ContaCorrente string `json:"conta_corrente"`
		DocTitular    string `json:"doc_titular"`
		NomeTitular   string `json:"nome_titular"`
	} `json:"dadosBancarios"`
	Info struct {
		Cimpapi string `json:"cImpAPI"`
		Dalt    string `json:"dAlt"`
		Dinc    string `json:"dInc"`
		Halt    string `json:"hAlt"`
		Hinc    string `json:"hInc"`
		Ualt    string `json:"uAlt"`
		Uinc    string `json:"uInc"`
	} `json:"info"`
	BloquearExclusao string `json:"bloquear_exclusao"`
}
