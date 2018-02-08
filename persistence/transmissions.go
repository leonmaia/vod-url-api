package persistence

import (
	"database/sql"

	"github.com/leonmaia/vod-api/model"
)

//Repository ...
type Repository struct {
	DB *sql.DB
}

//Insert ...
func (r Repository) Insert(t model.Transmission) error {
	stmtIns, err := r.DB.Prepare("INSERT INTO transmission VALUES ( ?, ? )")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(&t.ID, &t.URL)
	return err
}

//Get ...
func (r Repository) Get(id string) model.Transmission {
	transmission := model.Transmission{}
	err := r.DB.QueryRow("SELECT id, url FROM transmission LIMIT 1").Scan(&transmission.ID, &transmission.URL)
	if err != nil {
		panic(err.Error())
	}
	return transmission
}
