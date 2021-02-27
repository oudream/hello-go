package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

type Person struct {
	name string
	age  int
}

type ALlPropertyInfos struct {
	picAmount              int
	ID                     int
	UID                    int
	isType                 int
	TID                    int
	City                   string
	Area                   string
	Street                 string
	PropertyName           string
	Building               string
	Floor                  string
	Unit                   string
	ParkingSpaces          string
	SaleType               int
	RentalPrice            float64
	SellingPrice           float64
	BuildingArea           float64
	SalableArea            float64
	IntervalRoom           int
	IntervalHall           int
	IntervalToilet         int
	IntervalTerrace        int
	IntervalStoreRooms     int
	isParkingSpaces        int
	isLift                 int
	Landscape              string
	LandscapeDescription   string
	Direction              int
	Furniture              string
	isNew                  int
	Decoration             int
	DecorationDescription  string
	Detail                 string
	Developers             string
	PropertyCorp           string
	LinkMan                string
	Tel                    string
	Mobile                 string
	Email                  string
	isStatus               int
	PubDate                time.Time
	Hits                   int
	ExpireDate             time.Time
	EditDate               time.Time
	ContactHits            int
	Map_X                  string
	Map_Y                  string
	UnitPrice              float64
	Sequence               int
	AdditionalFeaturesType string
	Source                 int
	RentReservePrice       float64
	SellReservePrice       float64
	BrokerageType          int
	BrokeragePercent       int
	Cooperation            int
	Remark                 string
	LoftDescribe           int
	Lease                  int
	HandPick               int
	MHandPick              int
	GoodPick               int
	Selected               bool
	Age                    float64
	Other                  int
	RoomType               int
	Coverkey               int
	SaleDate               time.Time
	IsTop                  int
	ui_isStatus            int
	Cover                  string
	ui_Other               int
}

func helloJson2() {
	time.Now()
	fp1 := "D:\\twant\\go853\\referto\\PropertyInfos1.json"
	content, _ := ioutil.ReadFile(fp1)
	var person []ALlPropertyInfos
	json.Unmarshal([]byte(content), &person)

}

func main() {
	//helloJson1()
	helloJson3()
}

func helloJson3() {
	type response1 struct {
		Page   int
		Fruits []string
	}

	type response2 struct {
		Page   int      `json:"page"`
		Fruits []string `json:"fruits"`
	}

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func helloJson1() {
	blob := `[{"name": "a","age":1},{"name": null,"age":1},{"name": "a","age":null}]`
	//blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])
}
