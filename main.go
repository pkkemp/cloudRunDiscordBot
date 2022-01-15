package main

import (
	ed255192 "crypto/ed25519"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"os"
)

const pubKey = "dc6f78b6668e74737ab0e80d30e13beb345dd74a8aa90b4e20b14a2bb2a4341f"

func main() {

	http.HandleFunc("/interactions/", interactions)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func interactions(w http.ResponseWriter, r *http.Request) {
	if verify(r) {

	} else {
		w.WriteHeader(401)
		return
	}

	//dg, err := discordgo.New("Bot " + Token)
	//dg.WebhookExecute()
	//if err != nil {
	//	fmt.Println("error creating Discord session,", err)
	//	return
	//}
	//http.Redirect(w, r, "/app/login", 307)
}
func verify(r *http.Request) bool {
	key := ed255192.PublicKey(pubKey)
	return discordgo.VerifyInteraction(r, key)
}
