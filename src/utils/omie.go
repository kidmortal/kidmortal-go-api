package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	requests "github.com/kidmortal/kidmortal-go-api/src/models/omie/requests"
)

type request struct {
	Call      string         `json:"call"`
	AppKey    string         `json:"app_key"`
	AppSecret string         `json:"app_secret"`
	Param     []clienteParam `json:"param,omitempty"`
}

type clienteParam struct {
	CodigoClienteOmie       int    `json:"codigo_cliente_omie,omitempty"`
	CodigoClienteIntegracao string `json:"codigo_cliente_integracao,omitempty"`
}

var baseUrl = "https://app.omie.com.br/api/v1/"
var contentType = "application/json"
var key = os.Getenv("OMIE_PYRAMID_APPKEY")
var secret = os.Getenv("OMIE_PYRAMID_APPSECRET")

// OmieGetClienteById OMIE busca cliente pelo ID, e Retorna os dados
func OmieGetClienteById(ID int) requests.ClienteOmie {
	endpoint := "geral/clientes/"
	param := request{Call: "ConsultarCliente",
		AppKey:    key,
		AppSecret: secret,
		Param: []clienteParam{
			{CodigoClienteOmie: ID, CodigoClienteIntegracao: ""},
		},
	}

	postBody, _ := json.Marshal(param)
	request := bytes.NewBuffer(postBody)
	resp, err := http.Post(baseUrl+endpoint, contentType, request)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var response requests.ClienteOmie
	json.Unmarshal(body, &response)
	return response
}
