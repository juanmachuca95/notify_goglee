package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	fileName        string = "chat.txt"
	errNotFoundFile error  = errors.New("open chat.txt: no such file or directory")
)

type MessageDTO struct {
	SenderName         string   `json:"sender_name"`
	FormattedTimestamp string   `json:"formatted_timestamp"`
	Messages           []string `json:"messages"`
}

// Format message from google meet script javascript
// {"sender-name":"Tú","formatted-timestamp":"19:38","messages":["lkjljl"]}

// funciona
func main() {
	http.HandleFunc("/data", data)

	// Server on port 8000 initialized
	log.Println("Server initialized on port 8000 🚀")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err.Error())
	}
}

func data(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := new(MessageDTO)
	err = json.Unmarshal(data, m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check file exists or create
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			log.Fatal("Cannot create file - err ", err)
		}
	}

	_, _ = file.WriteString("User: " + m.SenderName + " - " + m.FormattedTimestamp + "\n")
	for _, value := range m.Messages {
		_, err := file.WriteString(value + "\n")
		if err != nil {
			log.Fatal("Cannot create row text on file ", err)
		}
	}
	_, _ = file.WriteString("\n")
}
