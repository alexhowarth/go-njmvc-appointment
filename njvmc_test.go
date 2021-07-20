package main

import (
	"encoding/json"
	"log"
	"testing"
)

func Test_UnmarshalJSON_TimeDataType(t *testing.T) {
	td := `[{"LocationId":101,"FirstOpenSlot":"1222 Appointments Available <br/> Next Available: 08/24/2021 10:20 AM"},{"LocationId":102,"FirstOpenSlot":"930 Appointments Available <br/> Next Available: 08/30/2021 11:20 AM"}]`
	timeData := TimeDataType{}

	err := json.Unmarshal([]byte(td), &timeData)
	if err != nil {
		log.Fatal(err)
	}

	got := len(timeData)
	expected := 2
	if got != expected {
		t.Errorf("expected %v, got %v", got, expected)
	}

	got = timeData[0].LocationID
	expected = 101
	if got != expected {
		t.Errorf("expected %v, got %v", got, expected)
	}

	got = timeData[1].LocationID
	expected = 102
	if got != expected {
		t.Errorf("expected %v, got %v", got, expected)
	}
}

func Test_UnmarshalJSON_LocationDataType(t *testing.T) {
	ld := `[{"Name":"Bakers Basin - License or Non Driver ID Renewal","Street1":"3200 Brunswick Pike","Street2":null,"City":"Lawrenceville","State":"NJ","Zip":"08648","PhoneNumber":"(609) 292-6500","FaxNumber":null,"Lat":"676","Long":"788","TimeZone":"America/New_York","IpAddress":"","LocationGroupId":"BB","MapId":0,"NoOfWindows":3,"Status":true,"LunchStartTime":null,"LunchEndTime":null,"LocAppointments":[{"LocationId":101,"AppointmentType":null,"AppointmentTypeId":11,"ApiType":10,"DateCreated":"2020-10-20T12:40:56.2589709","DateModified":"2020-10-20T12:40:56.2589709","Id":148,"ErrorMessage":null,"HasError":false}],"LocationHours":[{"Day":1,"StartTime":"2020-10-26T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":101,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2020-10-16T13:45:18.8304398","DateModified":"2020-10-16T13:45:18.8304398","Id":1034,"ErrorMessage":null,"HasError":false},{"Day":2,"StartTime":"2020-10-26T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":101,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2020-10-16T13:45:18.8304458","DateModified":"2020-10-16T13:45:18.8304458","Id":1035,"ErrorMessage":null,"HasError":false},{"Day":3,"StartTime":"2020-10-26T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":101,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2020-10-16T13:45:18.8304499","DateModified":"2020-10-16T13:45:18.8304499","Id":1036,"ErrorMessage":null,"HasError":false},{"Day":4,"StartTime":"2020-10-26T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":101,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2020-10-16T13:45:18.8304538","DateModified":"2020-10-16T13:45:18.8304538","Id":1037,"ErrorMessage":null,"HasError":false},{"Day":5,"StartTime":"2020-10-26T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":101,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2020-10-16T13:45:18.8304577","DateModified":"2020-10-16T13:45:18.8304577","Id":1038,"ErrorMessage":null,"HasError":false},{"Day":6,"StartTime":"2020-10-26T08:20:00","EndTime":"2020-10-26T14:30:00","Status":true,"LocationId":101,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2020-10-16T13:45:18.8304618","DateModified":"2020-10-16T13:45:18.8304618","Id":1039,"ErrorMessage":null,"HasError":false}],"AppointmentTypes":null,"LunchStartTimeString":null,"LunchEndTimeString":null,"ApiType":8,"TenantId":3,"Tenant":null,"DateCreated":"2020-10-22T09:27:15.68681","DateModified":"2020-10-22T09:27:15.68681","Id":101,"ErrorMessage":null,"HasError":false},{"Name":"Bayonne - License or Non Driver ID Renewal","Street1":"RT 440 & 1347 Kennedy Blvd","Street2":"Family Dollar Plaza","City":"Bayonne","State":"NJ","Zip":"07002","PhoneNumber":"(609) 292-6500","FaxNumber":null,"Lat":"1028","Long":"501","TimeZone":"America/New_York","IpAddress":"","LocationGroupId":"BA","MapId":0,"NoOfWindows":3,"Status":true,"LunchStartTime":null,"LunchEndTime":null,"LocAppointments":[{"LocationId":102,"AppointmentType":null,"AppointmentTypeId":11,"ApiType":10,"DateCreated":"2021-02-18T15:23:25.2963344","DateModified":"2021-02-18T15:23:25.2963344","Id":252,"ErrorMessage":null,"HasError":false}],"LocationHours":[{"Day":1,"StartTime":"2021-02-18T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":102,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2021-02-18T15:23:25.2963363","DateModified":"2021-02-18T15:23:25.2963363","Id":1660,"ErrorMessage":null,"HasError":false},{"Day":2,"StartTime":"2021-02-18T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":102,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2021-02-18T15:23:25.2963397","DateModified":"2021-02-18T15:23:25.2963397","Id":1661,"ErrorMessage":null,"HasError":false},{"Day":3,"StartTime":"2021-02-18T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":102,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2021-02-18T15:23:25.2963417","DateModified":"2021-02-18T15:23:25.2963417","Id":1662,"ErrorMessage":null,"HasError":false},{"Day":4,"StartTime":"2021-02-18T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":102,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2021-02-18T15:23:25.2963437","DateModified":"2021-02-18T15:23:25.2963437","Id":1663,"ErrorMessage":null,"HasError":false},{"Day":5,"StartTime":"2021-02-18T08:00:00","EndTime":"2020-10-26T15:45:00","Status":true,"LocationId":102,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2021-02-18T15:23:25.2963456","DateModified":"2021-02-18T15:23:25.2963456","Id":1664,"ErrorMessage":null,"HasError":false},{"Day":6,"StartTime":"2020-10-26T08:20:00","EndTime":"2021-02-18T14:30:00","Status":true,"LocationId":102,"StartTimeString":null,"EndTimeString":null,"ApiType":11,"DateCreated":"2021-02-18T15:23:25.2963477","DateModified":"2021-02-18T15:23:25.2963477","Id":1665,"ErrorMessage":null,"HasError":false}],"AppointmentTypes":null,"LunchStartTimeString":null,"LunchEndTimeString":null,"ApiType":8,"TenantId":3,"Tenant":null,"DateCreated":"2021-02-18T15:23:25.2963315","DateModified":"2021-02-18T15:23:25.2963315","Id":102,"ErrorMessage":null,"HasError":false}]`
	locationData := LocationDataType{}

	err := json.Unmarshal([]byte(ld), &locationData)
	if err != nil {
		log.Fatal(err)
	}

	if val, ok := locationData[101]; ok {
		got := val.City
		expected := "Lawrenceville"
		if got != expected {
			t.Errorf("expected %s, got %s", got, expected)
		}
	} else {
		t.Error("expected locationData[101]")
	}

	if val, ok := locationData[102]; ok {
		got := val.City
		expected := "Bayonne"
		if got != expected {
			t.Errorf("expected %s, got %s", got, expected)
		}
	} else {
		t.Error("expected locationData[102]")
	}
}
