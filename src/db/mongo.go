package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kidmortal/kidmortal-go-api/src/models"
	"github.com/kidmortal/kidmortal-go-api/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var err error

type database struct {
	DB *mongo.Client
}

// GetInstance return copy of db session
func GetInstance() *mongo.Database {

	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		os.Getenv("dbuser"),
		os.Getenv("dbpassword"),
		os.Getenv("dbserver"),
		os.Getenv("dbname"))

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	db := client.Database("pyramid")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InsertOrUpdateNota(nota models.NfOmie, db *mongo.Database) {
	if nota.Ide.Tpnf == "0" {
		//	InsertOrUpdateNotaEntrada(nota,db)
	}
	if nota.Ide.Tpnf == "1" {
		insertOrUpdateNotaSaida(nota, db)
	}

}

func insertOrUpdateNotaSaida(nota models.NfOmie, db *mongo.Database) {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Collection("notas")

	numero, err := strconv.ParseInt(nota.Ide.Nnf, 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	cliente := utils.OmieGetClienteByNcodcli(nota.Nfdestint.Ncodcli)
	nfUrl := utils.OmieGetNfUrlBynIdNF(nota.Compl.Nidnf).Curldanfe
	imposto := nota.Total.Icmstot.Vnf - nota.Total.Icmstot.Vprod
	boleto := nota.Titulos[0].Ddtemissao == nota.Titulos[0].Ddtvenc

	filter := bson.M{"numero": numero}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	data, err := time.Parse("02/01/2006", nota.Info.Dinc)

	if err != nil {
		log.Fatal(err)
	}

	update := bson.M{
		"$set": models.Nota{
			Numero:        numero,
			Data:          data,
			Boleto:        boleto,
			Cliente:       cliente.RazaoSocial,
			Cnpj:          cliente.CnpjCpf,
			Empresa:       "PYRAMID",
			NfUrl:         nfUrl,
			ValorProdutos: nota.Total.Icmstot.Vprod,
			ValorTotal:    nota.Total.Icmstot.Vnf,
			ValorImposto:  imposto},
	}
	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		log.Fatal(result.Err())
	}
	fmt.Println("Nf Inserida/Atualizada ", numero)
}
