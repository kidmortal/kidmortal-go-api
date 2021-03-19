package utils

import (
	"fmt"
	"net/http"
)

func SendWhatsAppMessage(number string, apikey string, message string) {
	url := fmt.Sprintf("https://api.callmebot.com/whatsapp.php?phone=%s&text=%s&apikey=%s", number, message, apikey)
	http.Get(url)
}
