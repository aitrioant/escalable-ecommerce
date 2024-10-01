package model

// user represents data about a record user.
type Offer struct {
	ID     string
	CardId string
	Price  float32
}

// users slice to seed record user data.
var OffersData = []Offer{
	{ID: "1", CardId: "1", Price: 5.9},
	{ID: "2", CardId: "2", Price: 15.9},
	{ID: "3", CardId: "3", Price: 25.9},
	{ID: "4", CardId: "1", Price: 15.9},
	{ID: "5", CardId: "2", Price: 25.9},
	{ID: "6", CardId: "3", Price: 35.9},
}
