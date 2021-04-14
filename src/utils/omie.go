package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kidmortal/kidmortal-go-api/src/models"
)

type clienteRequest struct {
	Call         string                `json:"call"`
	AppKey       string                `json:"app_key"`
	AppSecret    string                `json:"app_secret"`
	ClienteParam []models.ClienteParam `json:"param,omitempty"`
}

type nfRequest struct {
	Call      string           `json:"call"`
	AppKey    string           `json:"app_key"`
	AppSecret string           `json:"app_secret"`
	NfParam   []models.NfParam `json:"param,omitempty"`
}

var baseUrl = "https://app.omie.com.br/api/v1/"
var contentType = "application/json"

// OmieGetClienteById OMIE busca cliente pelo ID, e Retorna os dados
func OmieGetClienteByNcodcli(ID int) models.ClienteOmie {
	endpoint := "geral/clientes/"
	param := clienteRequest{Call: "ConsultarCliente",
		AppKey:    os.Getenv("OMIE_PYRAMID_APPKEY"),
		AppSecret: os.Getenv("OMIE_PYRAMID_APPSECRET"),
		ClienteParam: []models.ClienteParam{
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
	var response models.ClienteOmie
	json.Unmarshal(body, &response)
	return response
}

func OmieGetNfSaidaByPeriod(startDate string, endDate string) []models.NfOmie {
	endpoint := "produtos/nfconsultar/"
	param := nfRequest{Call: "ListarNF",
		AppKey:    os.Getenv("OMIE_PYRAMID_APPKEY"),
		AppSecret: os.Getenv("OMIE_PYRAMID_APPSECRET"),
		NfParam: []models.NfParam{
			{
				Pagina:             1,
				RegistrosPorPagina: 50,
				ApenasImportadoAPI: "N",
				OrdenarPor:         "DATA_LANCAMENTO",
				Tpnf:               "1",
				Demiinicial:        startDate,
				Demifinal:          endDate,
			},
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
	var response models.NfConsultarRetorno
	json.Unmarshal(body, &response)

	notas := response.Notas
	if response.TotalDePaginas > 1 {
		for i := 2; i < response.TotalDePaginas; i++ {
			fmt.Println("Pagina atual: ", i, "/", response.TotalDePaginas)
			param.NfParam[0].Pagina = i
			postBody, _ := json.Marshal(param)
			request := bytes.NewBuffer(postBody)
			resp, err := http.Post(baseUrl+endpoint, contentType, request)
			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			var response models.NfConsultarRetorno
			json.Unmarshal(body, &response)
			notas = append(notas, response.Notas...)

		}
	}

	return notas
}

func OmieGetNfUrlBynIdNF(ID int) models.NfUrlRetorno {
	endpoint := "produtos/notafiscalutil/"
	param := nfRequest{Call: "GetUrlDanfe",
		AppKey:    os.Getenv("OMIE_PYRAMID_APPKEY"),
		AppSecret: os.Getenv("OMIE_PYRAMID_APPSECRET"),
		NfParam: []models.NfParam{
			{
				NCodNF: ID,
			},
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
	var response models.NfUrlRetorno
	json.Unmarshal(body, &response)
	return response
}
