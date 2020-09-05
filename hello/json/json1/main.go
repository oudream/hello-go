package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	age int
}

type ALlPropertyInfos struct {
	picAmount int  
	ID int  
	UID int  
	isType int  
	TID int  
	City string
	Area string
	Street string 
	PropertyName string 
	Building string 
	Floor string 
	Unit string 
	ParkingSpaces string 
	SaleType int  
	RentalPrice float64  
	SellingPrice float64  
	BuildingArea float64  
	SalableArea float64  
	IntervalRoom int  
	IntervalHall int  
	IntervalToilet int  
	IntervalTerrace int  
	IntervalStoreRooms int  
	isParkingSpaces int  
	isLift int  
	Landscape string 
	LandscapeDescription string 
	Direction int  
	Furniture string 
	isNew int  
	Decoration int  
	DecorationDescription string 
	Detail string 
	Developers string
	PropertyCorp string
	LinkMan string 
	Tel string 
	Mobile string 
	Email string
	isStatus int  
	PubDate time.Time
	Hits int  
	ExpireDate time.Time
	EditDate time.Time
	ContactHits int  
	Map_X string 
	Map_Y string 
	UnitPrice float64  
	Sequence int  
	AdditionalFeaturesType string 
	Source int  
	RentReservePrice float64  
	SellReservePrice float64  
	BrokerageType int  
	BrokeragePercent int  
	Cooperation int  
	Remark string
	LoftDescribe int  
	Lease int  
	HandPick int  
	MHandPick int  
	GoodPick int  
	Selected bool
	Age float64
	Other int 
	RoomType int 
	Coverkey int 
	SaleDate time.Time
	IsTop int 
	ui_isStatus int  
	Cover string 
	ui_Other int
}

func helloJson2()  {
	time.Now()
	fp1 := "D:\\twant\\go853\\referto\\PropertyInfos1.json"
	content, _ := ioutil.ReadFile(fp1)
	var person[] ALlPropertyInfos
	json.Unmarshal([]byte(content), &person)

}

func main() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
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
