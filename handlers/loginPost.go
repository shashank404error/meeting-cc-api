package handlers

import (
	"net/http"
	//"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// swagger:route POST /meeting/login meeting createEvent
// Login route for meeting API
//
// responses:
//  422: errorValidation
//  501: errorResponse

func (p *Meeting) CreateEventLogin (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> meeting-cc-api Module")

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/meeting/create/callback",
		ClientID:     "113188653176-fjoovrjckjns6hk9p9nunnp3677omhb3.apps.googleusercontent.com",
		ClientSecret: "C2b3yeljmmSW-rn5WEGJ17kl",
		Scopes:       []string{"https://www.googleapis.com/auth/calendar.events"},
		Endpoint:     google.Endpoint,
	}
}