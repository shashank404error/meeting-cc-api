package data

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"encoding/json"
	//"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/api/calendar/v3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	log "github.com/sirupsen/logrus"
	"reflect"
)

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func EventDetailCallback (r *http.Request) {
	_, token, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		log.Error("EventDetailCallback ERROR:")
		log.Error(err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config, token)

	srv, err := calendar.New(client)
	if err != nil {
			log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
			fmt.Println("No upcoming events found.")
	} else {
			for _, item := range events.Items {
					date := item.Start.DateTime
					if date == "" {
							date = item.Start.Date
					}
					fmt.Printf("%v (%v)\n", item.Summary, date)
			}
	}
}

func getUserInfo(state string, code string) ([]byte, *oauth2.Token, error) {
	if state != oauthStateString {
		return nil, nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, nil ,fmt.Errorf("code exchange failed: %s", err.Error())
	}
	fmt.Println("token = ", reflect.TypeOf(token))

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, token, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, tok *oauth2.Token) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	return config.Client(context.Background(), tok)
}