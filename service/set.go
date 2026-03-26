package database

import (
	"database/sql"
	"fmt"
	_ "fmt"

	"github.com/JakeDodd/mtgdeckbuilder/models"
)

const SET_BY_ID_SQL = "SELECT * FROM mtg_set WHERE set_id = $1"

func GetSetById(db *sql.DB, id string) (models.MtgSet, error) {
	var set models.MtgSet

	row := db.QueryRow(SET_BY_ID_SQL, id)
	err := row.Scan(&set.SetId, &set.SetCode, &set.SetName, &set.SetType, &set.SetUri, &set.SetSearchUri, &set.ScryfallSetUri)
	if err != nil {
		fmt.Errorf("GetSetById: %s, %v", id, err)
	}

	return set, nil
}
