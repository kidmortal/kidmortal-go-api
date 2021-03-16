package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	requests "github.com/kidmortal/kidmortal-go-api/src/models/omie/requests"
)

type clienteRequest struct {
	Call      string  `json:"call"`
	AppKey    string  `json:"app_key"`
	AppSecret string  `json:"app_secret"`
	Param     []param `json:"param"`
}

type param struct {
	CodigoClienteOmie       int    `json:"codigo_cliente_omie"`
	CodigoClienteIntegracao string `json:"codigo_cliente_integracao"`
}

// GetClienteById OMIE busca cliente pelo ID, e Retorna os dados
func GetClienteById(ID int) requests.ClienteOmie {

	cliente := clienteRequest{Call: "ConsultarCliente",
		AppKey:    os.Getenv("OMIE_PYRAMID_APPKEY"),
		AppSecret: os.Getenv("OMIE_PYRAMID_APPSECRET"),
		Param: []param{
			{CodigoClienteOmie: ID, CodigoClienteIntegracao: ""},
		},
	}

	postBody, _ := json.Marshal(cliente)
	request := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://app.omie.com.br/api/v1/geral/clientes/", "application/json", request)

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
