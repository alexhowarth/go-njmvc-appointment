package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/slack-go/slack"
)

const (
	dateTimeForm = "01/02/2006 15:04 PM"
	renewalURL   = "https://telegov.njportal.com/njmvc/AppointmentWizard/11"
)

var (
	timeDataRegex      = regexp.MustCompile(`var timeData = (\[.*?\])`)
	locationDataRegex  = regexp.MustCompile(`var locationData = (\[.*\])`)
	nextAvailableRegex = regexp.MustCompile(`Next Available: (.*)$`)
)

type TimeData struct {
	LocationID    int    `json:"LocationId"`
	FirstOpenSlot string `json:"FirstOpenSlot"`
	NextAvailable time.Time
}

type LocationData struct {
	Name            string      `json:"Name"`
	Street1         string      `json:"Street1"`
	Street2         interface{} `json:"Street2"`
	City            string      `json:"City"`
	State           string      `json:"State"`
	Zip             string      `json:"Zip"`
	PhoneNumber     string      `json:"PhoneNumber"`
	FaxNumber       interface{} `json:"FaxNumber"`
	Lat             string      `json:"Lat"`
	Long            string      `json:"Long"`
	TimeZone        string      `json:"TimeZone"`
	IPAddress       string      `json:"IpAddress"`
	LocationGroupID string      `json:"LocationGroupId"`
	MapID           int         `json:"MapId"`
	NoOfWindows     int         `json:"NoOfWindows"`
	Status          bool        `json:"Status"`
	LunchStartTime  interface{} `json:"LunchStartTime"`
	LunchEndTime    interface{} `json:"LunchEndTime"`
	LocAppointments []struct {
		LocationID        int         `json:"LocationId"`
		AppointmentType   interface{} `json:"AppointmentType"`
		AppointmentTypeID int         `json:"AppointmentTypeId"`
		APIType           int         `json:"ApiType"`
		DateCreated       string      `json:"DateCreated"`
		DateModified      string      `json:"DateModified"`
		ID                int         `json:"Id"`
		ErrorMessage      interface{} `json:"ErrorMessage"`
		HasError          bool        `json:"HasError"`
	} `json:"LocAppointments"`
	LocationHours []struct {
		Day             int         `json:"Day"`
		StartTime       string      `json:"StartTime"`
		EndTime         string      `json:"EndTime"`
		Status          bool        `json:"Status"`
		LocationID      int         `json:"LocationId"`
		StartTimeString interface{} `json:"StartTimeString"`
		EndTimeString   interface{} `json:"EndTimeString"`
		APIType         int         `json:"ApiType"`
		DateCreated     string      `json:"DateCreated"`
		DateModified    string      `json:"DateModified"`
		ID              int         `json:"Id"`
		ErrorMessage    interface{} `json:"ErrorMessage"`
		HasError        bool        `json:"HasError"`
	} `json:"LocationHours"`
	AppointmentTypes     interface{} `json:"AppointmentTypes"`
	LunchStartTimeString interface{} `json:"LunchStartTimeString"`
	LunchEndTimeString   interface{} `json:"LunchEndTimeString"`
	APIType              int         `json:"ApiType"`
	TenantID             int         `json:"TenantId"`
	Tenant               interface{} `json:"Tenant"`
	DateCreated          string      `json:"DateCreated"`
	DateModified         string      `json:"DateModified"`
	ID                   int         `json:"Id"`
	ErrorMessage         interface{} `json:"ErrorMessage"`
	HasError             bool        `json:"HasError"`
}

type TimeDataType []TimeData

func (tdc *TimeDataType) UnmarshalJSON(b []byte) (err error) {
	var tmp []TimeData
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	for i, v := range tmp {
		if nextAvailableRegex.MatchString(v.FirstOpenSlot) {
			d := nextAvailableRegex.FindStringSubmatch(v.FirstOpenSlot)[1]
			loc, _ := time.LoadLocation("America/New_York")
			t, err := time.ParseInLocation(dateTimeForm, d, loc)
			if err != nil {
				return err
			}
			tmp[i].NextAvailable = t
		}
	}

	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i].NextAvailable.Before(tmp[j].NextAvailable)
	})

	*tdc = tmp

	return nil
}

type LocationDataType map[int]LocationData

func (md LocationDataType) UnmarshalJSON(b []byte) (err error) {
	var tmp []LocationData
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	for _, v := range tmp {
		md[v.ID] = v
	}

	return nil
}

// flags, required for handling comma-separated location flag
type location []string

var locationFlag location

func (i *location) String() string {
	return fmt.Sprint(*i)
}

func (i *location) Set(value string) error {
	for _, l := range strings.Split(value, ",") {
		*i = append(*i, l)
	}
	return nil
}

var slackChannel, slackToken string

func main() {

	flag.Var(&locationFlag, "location", "comma-separated list limits results to one or more locations")
	flag.StringVar(&slackChannel, "slack-channel", "", "slack channel id to post to")
	flag.StringVar(&slackToken, "slack-token", "", "slack oauth token for your bot")
	flag.Parse()

	c := colly.NewCollector()

	timeData := TimeDataType{}
	locationData := LocationDataType{}

	c.OnHTML("script", func(e *colly.HTMLElement) {
		// scrape time JSON
		if timeDataRegex.MatchString(e.Text) {
			td := timeDataRegex.FindStringSubmatch(e.Text)[1]
			// unmarshal it
			err := json.Unmarshal([]byte(td), &timeData)
			if err != nil {
				log.Fatal(err)
			}
		}
		// scrape location JSON
		if locationDataRegex.MatchString(e.Text) {
			ld := locationDataRegex.FindStringSubmatch(e.Text)[1]
			// unmarshal it
			err := json.Unmarshal([]byte(ld), &locationData)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	c.Visit(renewalURL)

	if len(timeData) == 0 || len(locationData) == 0 {
		log.Fatal("Unable to scrape data")
	}

	var sb strings.Builder

	if len(locationFlag) > 0 {
		// iterate over available times
		for _, v := range timeData {
			ld := locationData[v.LocationID]
			for _, location := range locationFlag {
				// if the location exists, write
				if ld.City == location {
					sb.WriteString(prettyPrint(ld.City, v.NextAvailable))
				}
			}
		}
	} else {
		// write all data
		for _, v := range timeData {
			sb.WriteString(prettyPrint(locationData[v.LocationID].City, v.NextAvailable))
		}
	}

	// slack or stdout
	if slackChannel != "" && slackToken != "" {
		postSlackMessage(sb.String())
	} else {
		fmt.Print(sb.String())
	}
}

func prettyPrint(city string, date time.Time) string {
	return fmt.Sprintf("%-17v %v\n", city, date)
}

func postSlackMessage(txt string) {
	OAUTH_TOKEN := slackToken
	CHANNEL_ID := slackChannel

	api := slack.New(OAUTH_TOKEN)
	attachment := slack.Attachment{
		Pretext: "Available appointments",
		Text:    txt,
	}

	channelId, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	fmt.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)
}
