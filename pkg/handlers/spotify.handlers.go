package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/victorbischoff/GOFIBER-TMPL/pkg/config"
)

const redirectURI = "http://localhost:8080/callback"

func loadenv() config.Config {
	config, err := config.LoadConfig("./app.env")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	return config
}

var (
	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithClientID(loadenv().SPOTIFY_ID), spotifyauth.WithClientSecret(loadenv().SPOTIFY_SECRET), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func CompleteAuth(w http.ResponseWriter, r *http.Request) {

	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

func RegHand(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request for:", r.URL.String())

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}
