package main

import (
	"fmt"
	"log"

	calendar "google.golang.org/api/calendar/v3"

	"net/http"

	"time"

	"strings"

	"github.com/abiosoft/ishell"
)

var monthMap = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}
var client *http.Client
var availableCalendars map[string]string
var usedCalendar = "primary"

func main() {
	shell := ishell.New()

	shell.Println("\n\t\t>>> GoCal <<<")
	connect()
	listCalendars()

	shell.AddCmd(&ishell.Cmd{
		Name: "choose",
		Help: "choose calendar",
		Func: func(c *ishell.Context) {
			chooseCalendar(c)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "add",
		Help: "add event",
		Func: func(c *ishell.Context) {
			createEvent(c.Args)
		},
	})

	shell.Start()
}

func connect() {
	client = passAuthAndGetClient()
	fmt.Println("\nConnected!")
}

func listCalendars() {
	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar Client %v", err)
	}

	calendars, err := srv.CalendarList.List().Do()

	availableCalendars = make(map[string]string)

	if len(calendars.Items) > 0 {
		for _, i := range calendars.Items {
			availableCalendars[i.Summary] = i.Id
		}
	}
}

func chooseCalendar(c *ishell.Context) {
	c.Print("Which calendar to use?: ")
	listOfCalendars := make([]string, 0, len(availableCalendars))
	for k := range availableCalendars {
		listOfCalendars = append(listOfCalendars, k)
	}

	choice := c.MultiChoice(listOfCalendars, "Which calendar to use?")
	usedCalendar = availableCalendars[listOfCalendars[choice]]
	c.Println("Use calendar: ", listOfCalendars[choice])
}

func createEvent(args []string) {
	title := fmt.Sprintf("%v %v's birthday", args[0], args[1])
	date := args[2]

	startTime := buildDate(date) + "10:00:00"
	endTime := buildDate(date) + "10:30:00"

	event := &calendar.Event{
		Summary:     title,
		Location:    "",
		Description: "",
		Start: &calendar.EventDateTime{
			DateTime: startTime,
			TimeZone: "Europe/Warsaw",
		},
		End: &calendar.EventDateTime{
			DateTime: endTime,
			TimeZone: "Europe/Warsaw",
		},
		Recurrence: []string{"RRULE:FREQ=YEARLY"},
	}

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar Client %v", err)
	}

	calendarId := usedCalendar
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
}

func buildDate(date string) string {
	year := time.Now().Year()
	pd := strings.Split(date, "-")
	month := strings.ToLower(pd[0])
	day := pd[1]
	return fmt.Sprintf("%v-%v-%vT", year, monthMap[month], day)
}
