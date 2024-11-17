package models

type NviRequest struct {
	ID        string `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	BirthYear string `json:"birth_year" binding:"required"`
}

type ForeignNviRequest struct {
	ID         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	BirthDay   int    `json:"birth_day" binding:"required"`
	BirthMonth int    `json:"birth_month" binding:"required"`
	BirthYear  int    `json:"birth_year" binding:"required"`
}

type NviResponse struct {
	TCKimlikNoDogrulaResult string `xml:"Body>TCKimlikNoDogrulaResponse>TCKimlikNoDogrulaResult"`
}

type ForeignNviResponse struct {
	YabanciKimlikNoDogrulaResult string `xml:"Body>YabanciKimlikNoDogrulaResponse>YabanciKimlikNoDogrulaResult"`
}
