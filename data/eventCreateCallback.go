package data

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"google.golang.org/api/calendar/v3"
	"golang.org/x/oauth2/google"
	log "github.com/sirupsen/logrus"
)

func EventCreateCallback (d *CreateEventRequest,r *http.Request) {
	_, token, err := GetUserInfo(r.FormValue("state"), r.FormValue("code"))
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
	client := GetClient(config, token)

	srv, err := calendar.New(client)
	if err != nil {
			log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	event := &calendar.Event{
		Summary: d.Summary,
		Location: d.Location,
		Description: d.Description,
		Start: &calendar.EventDateTime{
		  DateTime: d.StartTime,
		  TimeZone: d.TimeZone,
		},
		End: &calendar.EventDateTime{
		  DateTime: d.EndTime,
		  TimeZone: d.TimeZone,
		},
		Attendees: []*calendar.EventAttendee{
		  &calendar.EventAttendee{Email: d.Email},
		},
	  }
	  
	  calendarId := "primary"
	  event, err = srv.Events.Insert(calendarId, event).Do()
	  if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	  }
	  fmt.Printf("Event created: %s\n", event.HtmlLink)
	  
}