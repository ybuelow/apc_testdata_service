package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var passengers int

func main() {
	/**Setup interval and publish/client*/
	var interval int
	fmt.Println("Please state the interval of which the telemtry is sent: ")
	fmt.Scan(&interval)
	err := godotenv.Load(".env")
	check(err)
	topic := os.Getenv("TOPIC")
	client, err := newClient()
	check(err)
	serr :=client.sub(topic, messagePubHandler)
	check(serr)
	cerr:=client.Publish(topic, interval)
	check(cerr)
	//client.Disconnect(250)
}

// Make the option to stop outside the planed route TODO: case ausserplanmässiger halt


func insertTelemetry(round int, max int) APC {
	f, err := os.ReadFile("./data/telemetry.json")
	check(err)
	var apt APC
	unerr := json.Unmarshal([]byte(f), &apt)
	check(unerr)
	prt := &passengers
	var passengersin int
	var passengersout int
	if round == 0 {
		passengers = 0
		in, out := allocatePassangers(0, "start")
		passengersin += in
		passengersout += out
		*prt += passengersin
	} else if round == max {
		passengersin, passengersout = allocatePassangers(*prt, "end")
	} else {
		cin, cout := allocatePassangers(passengers, "normal")
		passengersin += cin
		passengersout += cout
		*prt += passengersin
		*prt -= passengersout
	}
	h, err := os.ReadFile("./data/Haltestellen.json")
	check(err)
	var station Station
	herr := json.Unmarshal([]byte(h), &station)
	check(herr)
	ia, ib, ic := travers(passengersin)
	a, b, c := travers(passengersout)
	apt.Sensors[0].Counts[0].In = int64(ia)
	apt.Sensors[0].Counts[0].Out = int64(a)
	apt.Sensors[1].Counts[0].In = int64(ib)
	apt.Sensors[1].Counts[0].Out = int64(b)
	apt.Sensors[2].Counts[0].In = int64(ic)
	apt.Sensors[2].Counts[0].Out = int64(c)
	apt.Gps.Latitude = station.Haltepunkte[round].Latitude
	apt.Gps.Longitude = station.Haltepunkte[round].Longitude
	apt.Gps.Timestamp = time.Now().String()
	apt.Timestamp = time.Now().String()
	return apt
}
/** Helper Functions*/

func allocatePassangers(passengers int, station string) (int, int) {
	switch {
	case station == "start":
		return randRange(0, 30), 0
	case station == "end":
		return 0, passengers
	default:
		return randRange(0, 30), randRange(0, passengers)
	}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func travers(p int) (int, int, int) {
	switch{
	case p == 0:
		return 0, 0, 0
	case p%3 == 0:
		return p/3, p/3, p/3
	case p%2 == 0:
		return p/2, p/2, 0
	default:
		return travers(p + 1)
	}
}

type APC struct {
	Timestamp    string      `json:"Timestamp"`
	Gps          Gps         `json:"Gps"`
	VehicleNr    string      `json:"VehicleNr"`
	StopID       interface{} `json:"StopId"`
	TripID       interface{} `json:"TripId"`
	BasisVersion interface{} `json:"BasisVersion"`
	Sensors      []Sensor    `json:"Sensors"`
}

type Gps struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Angle     int64   `json:"Angle"`
	DeviceID  string  `json:"DeviceId"`
	Timestamp string  `json:"Timestamp"`
}

type Sensor struct {
	SensorID string  `json:"SensorId"`
	Counts   []Count `json:"Counts"`
}

type Count struct {
	ObjectClass string `json:"ObjectClass"`
	In          int64  `json:"In"`
	Out         int64  `json:"Out"`
}
type Station struct {
	Haltepunkte []Haltepunkte `json:"Haltepunkte"`
}

type Haltepunkte struct {
	Name      string  `json:"Name"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}
