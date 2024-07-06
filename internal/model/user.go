package model

type User struct {
	ID                      string `json:"id,omitempty" `
	PassportSeriesAndNumber string `json:"passport_series_and_number"`
	Name                    string `json:"name"`
	Surname                 string `json:"surname"`
	Address                 string `json:"address"`
}
