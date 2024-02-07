package model

type ConsultancyForAlgolia struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
}

type ConsultancyInput struct {
	ID                   ID       `json:"_id"`
	Name                 string   `json:"name"`
	Address              string   `json:"address"`
	City                 string   `json:"city"`
	DestinationCountries []string `json:"destination_countries"`
	Description          string   `json:"description"`
	Email                string   `json:"email"`
	Phone                []string `json:"phone"`
	Website              string   `json:"website"`
	Images               []string `json:"images"`
	Services             []string `json:"services"`
	Reviews              []string `json:"reviews"`
	SocialMedia          []string `json:"social_media"`
	TestPreparation      []string `json:"test_preparation"`
	Languages            []string `json:"languages"`
	IsBranch             bool     `json:"isBranch"`
	Featured             bool     `json:"featured"`
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	AddedAt              Date     `json:"added_at"`
}

type ID struct {
	OID string `json:"$oid"`
}

type Date struct {
	Date string `json:"$date"`
}
