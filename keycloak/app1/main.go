package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	clientID     = "app"
	clientSecret = "637bac42-74b1-4099-9bc8-5e127a4b796c"
)

func main() {

	//variavel de contexto - caso de errado ele para para nao continuar
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8081/auth/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8082/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	//garante qu e ninguem esta tentando te enganar
	state := "exemplo"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		//protecao para garantir que nao e outro servidor forjando o redirecionamento
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State doesnt match", http.StatusBadRequest)
			return
		}
		//o Exchange vai trocar o codigo de autorizacao pelo callback pelo token que a gente quer
		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "problem to change token", http.StatusInternalServerError)
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "problem to take token", http.StatusInternalServerError)
			return
		}

		res := struct {
			OAuth2Token *oauth2.Token
			IDToken     string
		}{
			oauth2Token, rawIDToken,
		}

		data, _ := json.MarshalIndent(res, "", "   ")
		w.Write(data)

	})

	log.Fatal(http.ListenAndServe(":8082", nil))

}
