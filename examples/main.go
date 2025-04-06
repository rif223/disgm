package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/rif223/disgm"
	"github.com/rif223/disgm/store"
)

type Test struct {
	file string
}

var _ store.TokenStore = (*Test)(nil)

func NewTestStore(file string) (t *Test) {
	return &Test{file}
}

func (t *Test) Store(tokens map[string]string) (err error) {
	f, err := os.Create(t.file)
	if err != nil {
		return
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(tokens)
	return
}

func (t *Test) Load() (tokens map[string]string, err error) {
	tokens = map[string]string{}
	f, err := os.Open(t.file)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&tokens)
	return
}

func main() {

	session, err := discordgo.New("Bot " + "1234567890")
	if err != nil {
		log.Fatal("Error creating Discord session,", err)
	}

	disgmInstance, err := disgm.New(session, disgm.Options{
		DisableStartupMessage: true,
		TokenStore:            NewTestStore(".tokens.json"),
	})
	if err != nil {
		log.Fatal("Error initializing Disgm,", err)
	}

	disgmInstance.RegisterApiRouter()
	disgmInstance.RegisterWebSocket()

	disgmInstance.Listen()

	err = session.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer session.Close()

	// keep bot running untill there is NO os interruption (ctrl + C)
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}
