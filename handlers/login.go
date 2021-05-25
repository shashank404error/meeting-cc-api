
package handlers

/*import (
	"net/http"
	//"github.com/gorilla/mux"
	//"github.com/shashank404error/meeting-cc-api/data"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// swagger:route GET /meeting/login meeting OAuth2
// Get OAuth2/google authorization from user.
//
// responses:
//  422: errorValidation
//  501: errorResponse

func (p *Meeting) MeetingLogin(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> meeting-cc-api Module")
	
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
		RedirectURL:  "http://localhost:8080/meeting/get/all/callback",
		ClientID:     "113188653176-fjoovrjckjns6hk9p9nunnp3677omhb3.apps.googleusercontent.com",
		ClientSecret: "C2b3yeljmmSW-rn5WEGJ17kl",
		Scopes:       []string{"https://www.googleapis.com/auth/calendar.events"},
		Endpoint:     google.Endpoint,
	}
}*/