package user

// user represents data about a record user.
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// users slice to seed record user data.
var Users = []User{
	{ID: "1", Username: "SkylineExplorer", Password: "safe-password"},
	{ID: "2", Username: "DigitalNomad42", Password: "super-safe-password"},
	{ID: "3", Username: "QuantumJumper", Password: "ultra-safe-password"},
}
