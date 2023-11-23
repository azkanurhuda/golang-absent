package entity

type AbsentUser struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"foreignKey"`
	IPAddress string
	Latitude  float64
	Longitude float64
	Status    string
}

type Location struct {
	IP        string  `json:"ip"`
	City      string  `json:"city"`
	Region    string  `json:"region_name"`
	Country   string  `json:"country_name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
