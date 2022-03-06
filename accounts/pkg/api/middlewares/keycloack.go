package middlewares

import "github.com/Nerzal/gocloak/v7"

type keycloak struct {
	gocloak      gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}

const KEYCLOACK_URL = "http://localhost:8080"

func newKeycloak() *keycloak {
	return &keycloak{
		gocloak:      gocloak.NewClient(KEYCLOACK_URL),
		clientId:     "nodejsapp",
		clientSecret: "PvDC5d6WgEEjJk9WPWXxe3UcNEAIlySe",
		realm:        "morpheus",
	}
}
