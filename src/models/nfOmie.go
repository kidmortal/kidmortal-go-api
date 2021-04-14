package models

// ClienteOmie estrutura de retorno de um cliente do OMIE

type NfParam struct {
	Pagina             int    `json:"pagina,omitempty"`
	RegistrosPorPagina int    `json:"registros_por_pagina,omitempty"`
	ApenasImportadoAPI string `json:"apenas_importado_api,omitempty"`            // S - N
	OrdenarPor         string `json:"ordenar_por,omitempty"`                     // CODIGO - INTEGRACAO - DATA_LANCAMENTO
	Tpnf               string `json:"tpNF,omitempty"`                            // "0" = entrada , "1" = saida
	Demiinicial        string `json:"dEmiInicial,omitempty" validate:"required"` // format -> "25/12/2020"
	Demifinal          string `json:"dEmiFinal,omitempty" validate:"required"`   // format -> "25/12/2020"
	NCodNF             int    `json:"nCodNF,omitempty"`
}

type NfConsultarRetorno struct {
	Pagina           int      `json:"pagina"`
	TotalDePaginas   int      `json:"total_de_paginas"`
	Registros        int      `json:"registros"`
	TotalDeRegistros int      `json:"total_de_registros"`
	Notas            []NfOmie `json:"nfCadastro"`
}

type NfUrlRetorno struct {
	Curldanfe string `json:"cUrlDanfe"`
	Dtvalurl  string `json:"dtValUrl"`
}

type NfOmie struct {
	Compl struct {
		Cchavenfe string `json:"cChaveNFe"`
		Ccodcateg string `json:"cCodCateg"`
		Cmodfrete string `json:"cModFrete"`
		Nidnf     int    `json:"nIdNF"`
		Nidpedido int    `json:"nIdPedido"`
		Nidreceb  int    `json:"nIdReceb"`
		Nidtransp int    `json:"nIdTransp"`
	} `json:"compl"`
	Det []struct {
		Nfprodint struct {
			Ccoditemint string `json:"cCodItemInt"`
			Ccodprodint string `json:"cCodProdInt"`
			Ncoditem    int    `json:"nCodItem"`
			Ncodprod    int    `json:"nCodProd"`
		} `json:"nfProdInt"`
		Prod struct {
			Cfop         string  `json:"CFOP"`
			Extipi       string  `json:"EXTIPI"`
			Ncm          string  `json:"NCM"`
			Cean         string  `json:"cEAN"`
			Ceantrib     string  `json:"cEANTrib"`
			Cprod        string  `json:"cProd"`
			Cprodorig    string  `json:"cProdOrig"`
			Indtot       string  `json:"indTot"`
			Ncmctotal    float64 `json:"nCMCTotal"`
			Ncmcunitario float64 `json:"nCMCUnitario"`
			Picmsst      float64 `json:"pICMSST"`
			Qcom         float64 `json:"qCom"`
			Qtrib        float64 `json:"qTrib"`
			Ucom         string  `json:"uCom"`
			Utrib        string  `json:"uTrib"`
			Vdesc        float64 `json:"vDesc"`
			Vfrete       float64 `json:"vFrete"`
			Vicmsst      float64 `json:"vICMSST"`
			Voutro       float64 `json:"vOutro"`
			Vprod        float64 `json:"vProd"`
			Vseg         float64 `json:"vSeg"`
			Vtotitem     float64 `json:"vTotItem"`
			Vuncom       float64 `json:"vUnCom"`
			Vuntrib      float64 `json:"vUnTrib"`
			Xprod        string  `json:"xProd"`
			Xprodorig    string  `json:"xProdOrig"`
		} `json:"prod"`
		Rastreabilidade struct {
		} `json:"rastreabilidade"`
	} `json:"det"`
	Ide struct {
		Dcan    string `json:"dCan"`
		Demi    string `json:"dEmi"`
		Dinut   string `json:"dInut"`
		Dreg    string `json:"dReg"`
		Dsaient string `json:"dSaiEnt"`
		Finnfe  string `json:"finNFe"`
		Hemi    string `json:"hEmi"`
		Hsaient string `json:"hSaiEnt"`
		Indpag  string `json:"indPag"`
		Mod     string `json:"mod"`
		Nnf     string `json:"nNF"`
		Serie   string `json:"serie"`
		Tpamb   string `json:"tpAmb"`
		Tpnf    string `json:"tpNF"`
	} `json:"ide"`
	Info struct {
		Cimpapi string `json:"cImpAPI"`
		Dalt    string `json:"dAlt"`
		Dinc    string `json:"dInc"`
		Halt    string `json:"hAlt"`
		Hinc    string `json:"hInc"`
		Ualt    string `json:"uAlt"`
		Uinc    string `json:"uInc"`
	} `json:"info"`
	Nfdestint struct {
		Ccodcliint string `json:"cCodCliInt"`
		Ncodcli    int    `json:"nCodCli"`
	} `json:"nfDestInt"`
	Nfemitint struct {
		Ccodempint string `json:"cCodEmpInt"`
		Ncodemp    int    `json:"nCodEmp"`
	} `json:"nfEmitInt"`
	Pedido struct {
	} `json:"pedido"`
	Titulos []struct {
		Ccodcateg     string  `json:"cCodCateg"`
		Ccodinttitulo string  `json:"cCodIntTitulo"`
		Cnumtitulo    string  `json:"cNumTitulo"`
		Ddtemissao    string  `json:"dDtEmissao"`
		Ddtprevisao   string  `json:"dDtPrevisao"`
		Ddtvenc       string  `json:"dDtVenc"`
		Dreg          string  `json:"dReg"`
		Ncodcomprador int     `json:"nCodComprador"`
		Ncodprojeto   int     `json:"nCodProjeto"`
		Ncodtitrepet  int     `json:"nCodTitRepet"`
		Ncodtitulo    int     `json:"nCodTitulo"`
		Ncodvendedor  int     `json:"nCodVendedor"`
		Nparcela      int     `json:"nParcela"`
		Ntotparc      int     `json:"nTotParc"`
		Nvalortitulo  float64 `json:"nValorTitulo"`
	} `json:"titulos"`
	Total struct {
		Icmstot struct {
			Vbc             float64 `json:"vBC"`
			Vbcst           float64 `json:"vBCST"`
			Vcofins         float64 `json:"vCOFINS"`
			Vdesc           float64 `json:"vDesc"`
			Vfrete          float64 `json:"vFrete"`
			Vicms           float64 `json:"vICMS"`
			Vicmsdesonerado float64 `json:"vICMSDesonerado"`
			Vii             float64 `json:"vII"`
			Vipi            float64 `json:"vIPI"`
			Vnf             float64 `json:"vNF"`
			Voutro          float64 `json:"vOutro"`
			Vpis            float64 `json:"vPIS"`
			Vprod           float64 `json:"vProd"`
			Vst             float64 `json:"vST"`
			Vseg            float64 `json:"vSeg"`
			Vtottrib        float64 `json:"vTotTrib"`
		} `json:"ICMSTot"`
		Issqntot struct {
			Vbc     float64 `json:"vBC"`
			Vcofins float64 `json:"vCOFINS"`
			Viss    float64 `json:"vISS"`
			Vpis    float64 `json:"vPIS"`
			Vserv   float64 `json:"vServ"`
		} `json:"ISSQNtot"`
		Rettrib struct {
			Vbcirrf    float64 `json:"vBCIRRF"`
			Vbcretprev float64 `json:"vBCRetPrev"`
			Virrf      float64 `json:"vIRRF"`
			Vretcofins float64 `json:"vRetCOFINS"`
			Vretcsll   float64 `json:"vRetCSLL"`
			Vretpis    float64 `json:"vRetPIS"`
			Vretprev   float64 `json:"vRetPrev "`
		} `json:"retTrib"`
	} `json:"total"`
}
