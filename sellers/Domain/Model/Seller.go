package model

// user represents data about a record user.
type Seller struct {
	ID     string
	Offers []Offer `gorm:"foreignKey:ID"`
}

// users slice to seed record user data.
var SellersData = []Seller{
	{ID: "1", Offers: []Offer{OffersData[1], OffersData[2]}},
	{ID: "2", Offers: []Offer{OffersData[2], OffersData[3]}},
	{ID: "3", Offers: []Offer{OffersData[3], OffersData[4]}},
}
