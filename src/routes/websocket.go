package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/kidmortal/kidmortal-go-api/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type message struct {
	event   string
	message string
}

func WebsocketManager(r *fiber.App, db *mongo.Database) {
	socket := r.Group("/ws")
	socket.Use("/", func(c *fiber.Ctx) error {
		c.Request().Header.Set("Access-Control-Allow-Origin", "*")
		c.Request().Header.Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

		return c.Next()
	})
	socket.Get("/", websocket.New(func(c *websocket.Conn) {

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}

			switch string(msg) {
			case "REQUEST_NF_LIST":
				requestNfList(c, string(msg))

			}

		}

	}))

}

func requestNfList(c *websocket.Conn, msg string) {
	fmt.Println(string(msg))
	utils.OmieGetNfSaidaByPeriod("01/01/2021", "10/03/2021")

}
