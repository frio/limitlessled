package limitlessled

import (
	"database/sql"
)

type States struct {
	*sql.DB
}

func (db States) Retrieve() (*Bulb, error) {
	bulb := Bulb{}

	query := "SELECT brightness, temperature, isOn FROM bulb ORDER BY id DESC LIMIT 0, 1"
	err := db.QueryRow(query).Scan(&bulb.Brightness, &bulb.Temperature, &bulb.IsOn)

	if err != nil {
		return nil, err
	}

	return &bulb, nil
}

func (db States) Store(bulb Bulb) error {

	tx, err := db.Begin()
	defer tx.Commit()
	_, err = tx.Exec("INSERT INTO bulb (brightness, temperature, isOn) VALUES (?, ?, ?);",
		bulb.Brightness,
		bulb.Temperature,
		bulb.IsOn)

	return err

}
