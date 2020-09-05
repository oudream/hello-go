package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)
import "encoding/json"

type PublicKey struct {
	Id int
	Key string
}

func helloJson1() {
	keysBody := []byte(`[{"id": 1,"key": "-"},{"id": 2,"key": "-"},{"id": 3,"key": "-"}]`)
	keys := make([]PublicKey,0)
	json.Unmarshal(keysBody, &keys)

	//fmt.Printf("%#v", keys)
	fmt.Print(keys)
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
	//PubDate time.Time
	PubDate string
	Hits int
	//ExpireDate time.Time
	ExpireDate string
	//EditDate time.Time
	EditDate string
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
	//SaleDate time.Time
	SaleDate string
	IsTop int
	ui_isStatus int
	Cover string
	ui_Other int
}

func helloJson2() int {
	content, err := ioutil.ReadFile("D:\\twant\\go853\\referto\\PropertyInfos1.json")
	if err != nil {
		return -1
	}
	numberOfa := strings.Count(string(content), "},{")
	if numberOfa > 0 {
		numberOfa ++
	}
	keys := make([]ALlPropertyInfos,numberOfa,numberOfa)
	for i := 0; i < len(keys); i++  {

	}

	json.Unmarshal(content, &keys)

	fmt.Print(keys)
	return 0
}

func main() {
	helloJson2()
}