package database

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"strconv"

	"github.com/JakeDodd/mtgdeckbuilder/models"
)

var CardNotFound = &CardNotFoundError{"Card not found"}

type CardNotFoundError struct {
	Message string
}

func (e *CardNotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

const RANDOM_CARD_SQL = "SELECT * FROM cards ORDER BY RANDOM() limit 1"
const CARD_BY_KEY_SQL = "SELECT * FROM cards WHERE oracle_id = $1 and card_name = $2"
const COLORS_BY_CARD_SQL = "SELECT color FROM card_color WHERE oracle_id = $1 and card_name = $2"
const COLOR_IDENTITIES_BY_CARD_SQL = "SELECT color FROM card_color_identity WHERE oracle_id = $1 and card_name = $2"
const PRODUCED_MANA_BY_CARD_SQL = "SELECT color FROM card_produced_mana WHERE oracle_id = $1 and card_name = $2"
const COLOR_INDICATOR_BY_CARD_SQL = "SELECT color FROM card_color_indicator WHERE oracle_id = $1 and card_name = $2"
const KEYWORDS_BY_CARD_SQL = "SELECT keyword FROM card_keyword WHERE oracle_id = $1 and card_name = $2"

func GetRandomCard(db *sql.DB) (models.Cards, error) {
	var card models.Cards = models.Cards{}
	row := db.QueryRow(RANDOM_CARD_SQL)

	err := row.Scan(&card.OracleId, &card.Object, &card.CardName, &card.Layout, &card.ManaCost, &card.Cmc, &card.TypeLine, &card.Power, &card.Toughness, &card.Reserved, &card.StandardF, &card.FutureF, &card.HistoricF, &card.TimelessF, &card.GladiatorF, &card.PioneerF, &card.ExplorerF, &card.ModernF, &card.LegacyF, &card.PauperF, &card.VintageF, &card.PennyF, &card.CommanderF, &card.OathbreakerF, &card.StandardbrawlF, &card.BrawlF, &card.AlchemyF, &card.PaupercommanderF, &card.DuelF, &card.PremodernF, &card.PredhF, &card.Defense, &card.Loyalty, &card.EdhrecRank, &card.HandModifier, &card.LifeModifier)

	oracleId := card.OracleId
	name := card.CardName

	if err != nil {
		if err == sql.ErrNoRows {
			return card, CardNotFound
		}
		return card, fmt.Errorf("GetCardByoracleId: %s: %v", oracleId, err)
	}
	rows, err := db.Query(COLORS_BY_CARD_SQL, oracleId, name)
	card.Colors, err = GetListFromRows[string](rows, err)

	rows, err = db.Query(COLOR_IDENTITIES_BY_CARD_SQL, oracleId, name)
	card.ColorIdentity, err = GetListFromRows[string](rows, err)

	rows, err = db.Query(PRODUCED_MANA_BY_CARD_SQL, oracleId, name)
	card.ProducedMana, err = GetListFromRows[string](rows, err)

	rows, err = db.Query(COLOR_INDICATOR_BY_CARD_SQL, oracleId, name)
	card.ColorIndicator, err = GetListFromRows[string](rows, err)

	rows, err = db.Query(KEYWORDS_BY_CARD_SQL, oracleId, name)
	card.Keywords, err = GetListFromRows[string](rows, err)

	return card, nil

}

func GetCardByOracleIdAndName(oracleId string, name string, db *sql.DB) (models.Cards, error) {
	var card models.Cards = models.Cards{}
	row := db.QueryRow(CARD_BY_KEY_SQL, oracleId, name)

	err := row.Scan(&card.OracleId, &card.Object, &card.CardName, &card.Layout, &card.ManaCost, &card.Cmc, &card.TypeLine, &card.Power, &card.Toughness, &card.Reserved, &card.StandardF, &card.FutureF, &card.HistoricF, &card.TimelessF, &card.GladiatorF, &card.PioneerF, &card.ExplorerF, &card.ModernF, &card.LegacyF, &card.PauperF, &card.VintageF, &card.PennyF, &card.CommanderF, &card.OathbreakerF, &card.StandardbrawlF, &card.BrawlF, &card.AlchemyF, &card.PaupercommanderF, &card.DuelF, &card.PremodernF, &card.PredhF, &card.Defense, &card.Loyalty, &card.EdhrecRank, &card.HandModifier, &card.LifeModifier)

	if err != nil {
		if err == sql.ErrNoRows {
			return card, CardNotFound
		}
		return card, fmt.Errorf("GetCardByoracleId: %s: %v", oracleId, err)
	}
	rows, err := db.Query(COLORS_BY_CARD_SQL, oracleId, name)
	card.Colors, err = GetListFromRows[string](rows, err)

	rows, err = db.Query(COLOR_IDENTITIES_BY_CARD_SQL, oracleId, name)
	card.ColorIdentity, err = GetListFromRows[string](rows, err)

	rows, err = db.Query(PRODUCED_MANA_BY_CARD_SQL, oracleId, name)
	card.ProducedMana, err = GetListFromRows[string](rows, err)

	rows, err = db.Query(COLOR_INDICATOR_BY_CARD_SQL, oracleId, name)
	card.ColorIndicator, err = GetListFromRows[string](rows, err)

	rows, err = db.Query(KEYWORDS_BY_CARD_SQL, oracleId, name)
	card.Keywords, err = GetListFromRows[string](rows, err)

	return card, nil

}

func SaveCard(card models.Cards, db *sql.DB) error {

	row, err := db.Query("INSERT into cards (oracle_id ,\"object\", card_name, layout, mana_cost, cmc, type_line, power,"+
		" toughness, reserved, standard_f, future_f, historic_f, timeless_f, gladiator_f, pioneer_f, explorer_f, modern_f, legacy_f, pauper_f, vintage_f, penny_f, commander_f, oathbreaker_f, standardbrawl_f, brawl_f, alchemy_f, paupercommander_f, duel_f, premodern_f, predh_f, defense, loyalty, edhrec_rank, "+
		"hand_modifier, life_modifier) "+
		"VALUES ($1,$2,$3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36)",
		card.OracleId, card.Object, card.CardName, card.Layout, card.ManaCost, strconv.FormatFloat(card.Cmc, 'f', -1, 64), card.TypeLine, card.Power,
		card.Toughness, card.Reserved,
		card.StandardF, card.FutureF, card.HistoricF, card.TimelessF, card.GladiatorF, card.PioneerF,
		card.ExplorerF, card.ModernF, card.LegacyF,
		card.PauperF, card.VintageF, card.PennyF, card.CommanderF, card.OathbreakerF,
		card.StandardbrawlF, card.BrawlF, card.AlchemyF, card.PaupercommanderF, card.DuelF,
		card.PremodernF, card.PredhF, card.Defense, card.Loyalty,
		card.EdhrecRank, card.HandModifier, card.LifeModifier)

	if err != nil {
		return err
	}
	row.Close()

	for i := 0; i < len(card.Colors); i++ {
		row, err = db.Query("INSERT INTO card_color (card_name, oracle_id, color) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.Colors[i])
		if err != nil {
			return err
		}
		row.Close()
	}
	for i := 0; i < len(card.ColorIdentity); i++ {
		row, err = db.Query("INSERT INTO card_color_identity (card_name, oracle_id, color) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.ColorIdentity[i])
		if err != nil {
			return err
		}
		row.Close()
	}
	for i := 0; i < len(card.ProducedMana); i++ {
		row, err = db.Query("INSERT INTO card_produced_mana (card_name, oracle_id, color) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.ProducedMana[i])
		if err != nil {
			return err
		}
		row.Close()
	}
	for i := 0; i < len(card.ColorIndicator); i++ {
		row, err = db.Query("INSERT INTO card_color_indicator (card_name, oracle_id, color) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.ColorIndicator[i])
		if err != nil {
			return err
		}
		row.Close()
	}
	/*
		for i := 0; i < len(card.AttractionLights); i++ {
			row, err = db.Query("INSERT INTO card_attraction_light (card_name, attraction_light) VALUES ($1, $2)", card.CardName, card.AttractionLights[i])
			if err != nil {
				return err
			}
			row.Close()
		}
	*/
	for i := 0; i < len(card.Keywords); i++ {
		row, err = db.Query("INSERT INTO card_keyword (card_name, oracle_id, keyword) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.Keywords[i])
		if err != nil {
			return err
		}
		row.Close()
	}

	/*
		for i := 0; i < len(card.CardFaces); i++ {
			cf := card.CardFaces[i]
			r := db.QueryRow("SELECT * FROM card_faces WHERE card_name = $1", cf.Name)
			if err := r.Scan(); err == sql.ErrNoRows {
				row, err = db.Query("INSERT INTO card_faces (card_name, artist, artist_id, cmc, defense, flavor_text, illustration_id, png_uri, boarder_crop_uri, art_crop_uri, large_uri, normal_uri, small_uri, layout, loyalty, mana_cost, object_type, oracle_id, oracle_text, power, printed_name, printed_text, printed_type_line, toughness, type_line, watermark) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26)", cf.Name, cf.Artist, cf.ArtistId, cf.Cmc, cf.Defense, cf.FlavorText, cf.IllustrationId, cf.PngUri, cf.BoarderCropUri, cf.ArtCropUri, cf.LargeUri, cf.NormalUri, cf.SmallUri, cf.Layout, cf.Loyalty, cf.ManaCost, cf.Object, cf.OracleId, cf.OracleText, cf.Power, cf.PrintedName, cf.PrintedText, cf.PrintedTypeLine, cf.Toughness, cf.TypeLine, cf.Watermark)
				if err != nil {
					return err
				}
				row.Close()
				row, err = db.Query("INSERT INTO card_card_faces (card_card_name, card_faces_card_name) VALUES ($1, $2)", card.CardName, cf.Name)
				if err != nil {
					return err
				}
				row.Close()
				for i := 0; i < len(cf.Colors); i++ {
					row, err = db.Query("INSERT INTO card_faces_color (card_name, color) VALUES ($1, $2)", cf.Name, cf.Colors[i])
					if err != nil {
						return err
					}
					row.Close()
				}
				for i := 0; i < len(cf.ColorIndicator); i++ {
					row, err = db.Query("INSERT INTO card_faces_color_indicator (card_name, color) VALUES ($1, $2)", cf.Name, cf.ColorIndicator[i])
					if err != nil {
						return err
					}
					row.Close()
				}
			}
		}
	*/
	return nil
}
