package database

import (
	"database/sql"
	"fmt"
	_ "fmt"

	"github.com/JakeDodd/mtgdeckbuilder/models"
)

var PrintNotFound = &PrintNotFoundError{"Print not found"}

type PrintNotFoundError struct {
	Message string
}

func (e *PrintNotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func GetRandomPrint(db *sql.DB) (models.Prints, error) {
	var print models.Prints = models.Prints{}

	row := db.QueryRow("SELECT * FROM prints ORDER BY RANDOM() limit 1")

	err := row.Scan(&print.CardName, &print.SetId, &print.Lang, &print.OracleId, &print.MtgoId, &print.MtgoFoilId, &print.ArenaId, &print.TcgplayerId,
		&print.TcgplayerEtchedId, &print.ReleasedAt, &print.Oversized, &print.OracleText, &print.CollectorNumber, &print.Digital, &print.OldschoolF,
		&print.Rarity, &print.CardBackId, &print.Artist, &print.IllustrationId, &print.BorderColor, &print.Frame, &print.FullArt, &print.Textless, &print.Booster,
		&print.StorySpotlight, &print.TcgArticlesUri, &print.TcgDecksUri, &print.EdhrecUri, &print.TcgBuyUri, &print.CardmarketBuyUri, &print.CardhoarderBuyUri,
		&print.PrintsSearchUri, &print.FlavorName, &print.SecurityStamp, &print.PreviewedAt, &print.PreviewUri,
		&print.PreviewSource, &print.ContentWarning, &print.ScryfallUri, &print.RulingsUri, &print.GathererUri,
		&print.HighresImage, &print.ImageStatus, &print.Foil, &print.NotFoil, &print.Promo, &print.Reprint, &print.Variation, &print.VariationOf,
		&print.PriceUsd, &print.PriceUsdFoil, &print.PriceUsdEtched, &print.PriceEur, &print.PriceEurFoil, &print.PriceTix, &print.PrintedName,
		&print.PrintedText, &print.PrintedTypeLine, &print.CardmarketId, &print.Watermark, &print.PngUri, &print.BoarderCropUri,
		&print.ArtCropUri, &print.LargeUri, &print.NormalUri, &print.SmallUri)
	card_name := print.CardName
	set_id := print.SetId
	oracle_id := print.OracleId
	lang := print.Lang
	collector_number := print.CollectorNumber

	if err != nil {
		if err == sql.ErrNoRows {
			return print, PrintNotFound
		}
		return print, fmt.Errorf("GetPrintByCardName1: %s: %v", card_name, err)
	}

	rows, err := db.Query("SELECT attraction_light FROM print_attraction_light WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var attraction_lights []int

	if err == nil {
		for rows.Next() {
			var attraction_light int
			err = rows.Scan(&attraction_light)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName2: %s: %v", card_name, err)
			}
			attraction_lights = append(attraction_lights, attraction_light)
		}
	}
	print.AttractionLights = attraction_lights
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT game FROM print_game WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var games []string

	if err == nil {
		for rows.Next() {
			var game string
			err = rows.Scan(&game)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName3: %s: %v", card_name, err)
			}
			games = append(games, game)
		}
	}
	print.Games = games
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT border_effect FROM print_border_effect WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)

	var border_effects []string

	if err == nil {
		for rows.Next() {
			var border_effect string
			err = rows.Scan(&border_effect)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName4: %s: %v", card_name, err)
			}
			border_effects = append(border_effects, border_effect)
		}
	}
	print.BorderEffects = border_effects
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT frame_effect FROM print_frame_effect WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var frame_effects []string

	if err == nil {
		for rows.Next() {
			var frame_effect string
			err = rows.Scan(&frame_effect)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName5: %s: %v", card_name, err)
			}
			frame_effects = append(frame_effects, frame_effect)
		}
	}
	print.FrameEffects = frame_effects
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT related_id FROM print_related WHERE print_card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)

	var related_cards []models.Related

	if err == nil {
		for rows.Next() {
			var related_id string
			err = rows.Scan(&related_id)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %s: %v", card_name, set_id, err)
			}
			var related models.Related

			related_row := db.QueryRow("SELECT * FROM related WHERE id = $1", related_id)
			err = related_row.Scan(&related.Object, &related.Id, &related.Component, &related.Name, &related.TypeLine, &related.Uri)

			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %s: %v", card_name, set_id, err)

			}
			related_cards = append(related_cards, related)
		}
		print.Related = related_cards
		if rows != nil {
			rows.Close()
		}
	}

	rows, err = db.Query("SELECT finish FROM print_finish WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var finishs []string

	if err == nil {
		for rows.Next() {
			var finish string
			err = rows.Scan(&finish)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName6: %s: %v", card_name, err)
			}
			finishs = append(finishs, finish)
		}
	}
	print.Finishes = finishs
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT promo FROM print_promo WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var promos []string

	if err == nil {
		for rows.Next() {
			var promo string
			err = rows.Scan(&promo)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName7: %s: %v", card_name, err)
			}
			promos = append(promos, promo)
		}
	}
	print.PromoTypes = promos
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT multiverse_id FROM print_multiverse_id WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var multiverse_ids []int

	if err == nil {
		for rows.Next() {
			var multiverse_id int
			err = rows.Scan(&multiverse_id)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName8: %s: %v", card_name, err)
			}
			multiverse_ids = append(multiverse_ids, multiverse_id)
		}
	}
	print.MultiverseIds = multiverse_ids
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT card_faces_card_name FROM print_card_faces WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)

	var cardFaces []models.CardFaces

	if err == nil {
		for rows.Next() {
			var cardFacesCardName string
			err = rows.Scan(&cardFacesCardName)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %s: %v", card_name, set_id, err)
			}
			var cardFace models.CardFaces

			related_row := db.QueryRow("SELECT * FROM card_faces WHERE card_name = $1", cardFacesCardName)
			err = related_row.Scan(&cardFace.Name, &cardFace.Artist, &cardFace.ArtistId, &cardFace.Cmc, &cardFace.Defense, &cardFace.FlavorText,
				&cardFace.IllustrationId, &cardFace.PngUri, &cardFace.BoarderCropUri, &cardFace.ArtCropUri, &cardFace.LargeUri, &cardFace.NormalUri,
				&cardFace.SmallUri, &cardFace.Layout, &cardFace.Loyalty, &cardFace.ManaCost, &cardFace.Object, &cardFace.OracleId, &cardFace.OracleText,
				&cardFace.Power, &cardFace.PrintedName, &cardFace.PrintedText, &cardFace.PrintedTypeLine, &cardFace.Toughness, &cardFace.TypeLine, &cardFace.Watermark)

			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %s: %v", card_name, set_id, err)

			}
			cardFaces = append(cardFaces, cardFace)
		}
		print.CardFaces = cardFaces
		if rows != nil {
			rows.Close()
		}
	}

	return print, nil
}

func GetPrint(card_name string, oracle_id string, set_id string, lang string, collector_number string, db *sql.DB) (models.Prints, error) {
	var print models.Prints = models.Prints{}

	row := db.QueryRow("SELECT * FROM prints WHERE card_name = $1 and oracle_id = $2 and set_id = $3 and lang = $4 and collector_number = $5", card_name, oracle_id, set_id, lang, collector_number)

	err := row.Scan(&print.CardName, &print.SetId, &print.Lang, &print.OracleId, &print.MtgoId, &print.MtgoFoilId, &print.ArenaId, &print.TcgplayerId,
		&print.TcgplayerEtchedId, &print.ReleasedAt, &print.Oversized, &print.OracleText, &print.CollectorNumber, &print.Digital, &print.OldschoolF,
		&print.Rarity, &print.CardBackId, &print.Artist, &print.IllustrationId, &print.BorderColor, &print.Frame, &print.FullArt, &print.Textless, &print.Booster,
		&print.StorySpotlight, &print.TcgArticlesUri, &print.TcgDecksUri, &print.EdhrecUri, &print.TcgBuyUri, &print.CardmarketBuyUri, &print.CardhoarderBuyUri,
		&print.PrintsSearchUri, &print.FlavorName, &print.SecurityStamp, &print.PreviewedAt, &print.PreviewUri,
		&print.PreviewSource, &print.ContentWarning, &print.ScryfallUri, &print.RulingsUri, &print.GathererUri,
		&print.HighresImage, &print.ImageStatus, &print.Foil, &print.NotFoil, &print.Promo, &print.Reprint, &print.Variation, &print.VariationOf,
		&print.PriceUsd, &print.PriceUsdFoil, &print.PriceUsdEtched, &print.PriceEur, &print.PriceEurFoil, &print.PriceTix, &print.PrintedName,
		&print.PrintedText, &print.PrintedTypeLine, &print.CardmarketId, &print.Watermark, &print.PngUri, &print.BoarderCropUri,
		&print.ArtCropUri, &print.LargeUri, &print.NormalUri, &print.SmallUri)

	if err != nil {
		if err == sql.ErrNoRows {
			return print, PrintNotFound
		}
		return print, fmt.Errorf("GetPrintByCardName1: %s: %v", card_name, err)
	}

	rows, err := db.Query("SELECT attraction_light FROM print_attraction_light WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var attraction_lights []int

	if err == nil {
		for rows.Next() {
			var attraction_light int
			err = rows.Scan(&attraction_light)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName2: %s: %v", card_name, err)
			}
			attraction_lights = append(attraction_lights, attraction_light)
		}
	}
	print.AttractionLights = attraction_lights
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT game FROM print_game WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var games []string

	if err == nil {
		for rows.Next() {
			var game string
			err = rows.Scan(&game)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName3: %s: %v", card_name, err)
			}
			games = append(games, game)
		}
	}
	print.Games = games
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT border_effect FROM print_border_effect WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)

	var border_effects []string

	if err == nil {
		for rows.Next() {
			var border_effect string
			err = rows.Scan(&border_effect)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName4: %s: %v", card_name, err)
			}
			border_effects = append(border_effects, border_effect)
		}
	}
	print.BorderEffects = border_effects
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT frame_effect FROM print_frame_effect WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var frame_effects []string

	if err == nil {
		for rows.Next() {
			var frame_effect string
			err = rows.Scan(&frame_effect)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName5: %s: %v", card_name, err)
			}
			frame_effects = append(frame_effects, frame_effect)
		}
	}
	print.FrameEffects = frame_effects
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT related_id FROM print_related WHERE print_card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)

	var related_cards []models.Related

	if err == nil {
		for rows.Next() {
			var related_id string
			err = rows.Scan(&related_id)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %s: %v", card_name, set_id, err)
			}
			var related models.Related

			related_row := db.QueryRow("SELECT * FROM related WHERE id = $1", related_id)
			err = related_row.Scan(&related.Object, &related.Id, &related.Component, &related.Name, &related.TypeLine, &related.Uri)

			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %s: %v", card_name, set_id, err)

			}
			related_cards = append(related_cards, related)
		}
		print.Related = related_cards
		if rows != nil {
			rows.Close()
		}
	}

	rows, err = db.Query("SELECT finish FROM print_finish WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var finishs []string

	if err == nil {
		for rows.Next() {
			var finish string
			err = rows.Scan(&finish)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName6: %s: %v", card_name, err)
			}
			finishs = append(finishs, finish)
		}
	}
	print.Finishes = finishs
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT promo FROM print_promo WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var promos []string

	if err == nil {
		for rows.Next() {
			var promo string
			err = rows.Scan(&promo)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName7: %s: %v", card_name, err)
			}
			promos = append(promos, promo)
		}
	}
	print.PromoTypes = promos
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT multiverse_id FROM print_multiverse_id WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)
	var multiverse_ids []int

	if err == nil {
		for rows.Next() {
			var multiverse_id int
			err = rows.Scan(&multiverse_id)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName8: %s: %v", card_name, err)
			}
			multiverse_ids = append(multiverse_ids, multiverse_id)
		}
	}
	print.MultiverseIds = multiverse_ids
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT card_faces_card_name FROM print_card_faces WHERE card_name = $1 and set_id = $2 and oracle_id = $3 and lang = $4 and collector_number = $5", card_name, set_id, oracle_id, lang, collector_number)

	var cardFaces []models.CardFaces

	if err == nil {
		for rows.Next() {
			var cardFacesCardName string
			err = rows.Scan(&cardFacesCardName)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %s: %v", card_name, set_id, err)
			}
			var cardFace models.CardFaces

			related_row := db.QueryRow("SELECT * FROM card_faces WHERE card_name = $1", cardFacesCardName)
			err = related_row.Scan(&cardFace.Name, &cardFace.Artist, &cardFace.ArtistId, &cardFace.Cmc, &cardFace.Defense, &cardFace.FlavorText,
				&cardFace.IllustrationId, &cardFace.PngUri, &cardFace.BoarderCropUri, &cardFace.ArtCropUri, &cardFace.LargeUri, &cardFace.NormalUri,
				&cardFace.SmallUri, &cardFace.Layout, &cardFace.Loyalty, &cardFace.ManaCost, &cardFace.Object, &cardFace.OracleId, &cardFace.OracleText,
				&cardFace.Power, &cardFace.PrintedName, &cardFace.PrintedText, &cardFace.PrintedTypeLine, &cardFace.Toughness, &cardFace.TypeLine, &cardFace.Watermark)

			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %s: %v", card_name, set_id, err)

			}
			cardFaces = append(cardFaces, cardFace)
		}
		print.CardFaces = cardFaces
		if rows != nil {
			rows.Close()
		}
	}

	return print, nil
}

func SavePrint(print models.Prints, db *sql.DB) error {
	row, err := db.Query("INSERT into prints (lang, mtgo_id, mtgo_foil_id, arena_id, tcgplayer_id, tcgplayer_etched_id, released_at, oversized, set_id, oracle_text,"+
		"collector_number, digital, rarity, oldschool_f, card_back_id, artist, illustration_id, border_color, frame, full_art, textless, booster, story_spotlight, tcg_articles_uri,"+
		"tcg_decks_uri, edhrec_uri, tcg_buy_uri, cardmarket_buy_uri, cardhoarder_buy_uri, oracle_id, card_name, prints_search_uri, flavor_name, security_stamp, previewed_at, previewed_source_uri,"+
		"preview_source, content_warning, scryfall_uri, rulings_uri, gatherer_uri, highres_image, image_status, foil, not_foil, promo, reprint, variation,"+
		"variation_of, price_usd, price_usd_foil, price_usd_etched, price_eur, price_eur_foil, price_tix, printed_name, printed_next, printed_type_line, cardmarket_id,"+
		"watermark, png_uri, boarder_crop_uri, art_crop_uri, large_uri, normal_uri, small_uri) "+
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52,$53,$54,$55,$56,$57,$58,$59,$60,$61,$62,$63,$64,$65,$66)",
		print.Lang, print.MtgoId, print.MtgoFoilId, print.ArenaId, print.TcgplayerId, print.TcgplayerEtchedId, print.ReleasedAt, print.Oversized, print.SetId, print.OracleText,
		print.CollectorNumber, print.Digital, print.Rarity, print.OldschoolF, print.CardBackId, print.Artist, print.IllustrationId, print.BorderColor, print.Frame, print.FullArt, print.Textless, print.Booster,
		print.StorySpotlight, print.TcgArticlesUri, print.TcgDecksUri, print.EdhrecUri, print.TcgBuyUri, print.CardmarketBuyUri, print.CardhoarderBuyUri, print.OracleId, print.CardName,
		print.PrintsSearchUri, print.FlavorName, print.SecurityStamp, print.PreviewedAt, print.PreviewUri, print.PreviewSource, print.ContentWarning, print.ScryfallUri, print.RulingsUri, print.GathererUri,
		print.HighresImage, print.ImageStatus, print.Foil, print.NotFoil, print.Promo, print.Reprint, print.Variation, print.VariationOf, print.PriceUsd, print.PriceUsdFoil, print.PriceUsdEtched, print.PriceEur,
		print.PriceEurFoil, print.PriceTix, print.PrintedName, print.PrintedText, print.PrintedTypeLine, print.CardmarketId, print.Watermark, print.PngUri, print.BoarderCropUri, print.ArtCropUri, print.LargeUri, print.NormalUri, print.SmallUri)

	if err != nil {
		return err
	}
	row.Close()

	for i := 0; i < len(print.AttractionLights); i++ {
		row, err = db.Query("INSERT INTO print_attraction_light (card_name, oracle_id, set_id, attraction_light, lang, collector_number) VALUES ($1, $2, $3, $4, $5, $6)", print.CardName, print.OracleId, print.SetId, print.AttractionLights[i], print.Lang, print.CollectorNumber)
		if err != nil {
			return err
		}
		row.Close()
	}
	for i := 0; i < len(print.Games); i++ {
		row, err = db.Query("INSERT INTO print_game (card_name, oracle_id, set_id, game, lang, collector_number) VALUES ($1, $2, $3, $4, $5, $6)", print.CardName, print.OracleId, print.SetId, print.Games[i], print.Lang, print.CollectorNumber)
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.BorderEffects); i++ {
		row, err = db.Query("INSERT INTO print_border_effect (card_name, oracle_id, set_id, border_effect, lang, collector_number) VALUES ($1, $2, $3, $4, $5, $6)", print.CardName, print.OracleId, print.SetId, print.BorderEffects[i], print.Lang, print.CollectorNumber)
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.FrameEffects); i++ {
		row, err = db.Query("INSERT INTO print_frame_effect (card_name, oracle_id, set_id, frame_effect, lang, collector_number) VALUES ($1, $2, $3, $4, $5, $6)", print.CardName, print.OracleId, print.SetId, print.FrameEffects[i], print.Lang, print.CollectorNumber)
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.Related); i++ {
		related_card := print.Related[i]
		related_row := db.QueryRow("SELECT id FROM related WHERE id = $1", related_card.Id)
		var id string
		if err := related_row.Scan(&id); err == sql.ErrNoRows {
			row, err = db.Query("INSERT INTO related (object_parts, id, component, card_name, type_line, uri) VALUES ($1, $2, $3, $4, $5, $6)", related_card.Object, related_card.Id, related_card.Component, related_card.Name, related_card.TypeLine, related_card.Uri)
			if err != nil {
				return err
			}
			row.Close()
			row, err = db.Query("INSERT INTO print_related (print_card_name, oracle_id, set_id, related_id, lang, collector_number) VALUES ($1, $2, $3, $4, $5, $6)", print.CardName, print.OracleId, print.SetId, related_card.Id, print.Lang, print.CollectorNumber)
			if err != nil {
				return err
			}
			row.Close()
		}
	}

	for i := 0; i < len(print.CardFaces); i++ {
		card_faces := print.CardFaces[i]
		card_face_row := db.QueryRow("SELECT card_name FROM card_faces WHERE card_name = $1", card_faces.Name)
		var card_name string
		if err := card_face_row.Scan(&card_name); err == sql.ErrNoRows {
			row, err = db.Query("INSERT INTO card_faces (card_name, artist, artist_id, cmc,defense,flavor_text,illustration_id,png_uri,boarder_crop_uri,art_crop_uri,large_uri,normal_uri,small_uri,layout,loyalty,mana_cost,object_type,oracle_id,oracle_text,power,printed_name,printed_text,printed_type_line,toughness,type_line,watermark) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)",
				card_faces.Name, card_faces.Artist, card_faces.ArtistId, card_faces.Cmc, card_faces.Defense, card_faces.FlavorText,
				card_faces.IllustrationId, card_faces.PngUri, card_faces.BoarderCropUri, card_faces.ArtCropUri, card_faces.LargeUri, card_faces.NormalUri,
				card_faces.SmallUri, card_faces.Layout, card_faces.Loyalty, card_faces.ManaCost, card_faces.Object, card_faces.OracleId, card_faces.OracleText,
				card_faces.Power, card_faces.PrintedName, card_faces.PrintedText, card_faces.PrintedTypeLine, card_faces.Toughness, card_faces.TypeLine, card_faces.Watermark)
			if err != nil {
				return err
			}
			row.Close()
			row, err = db.Query("INSERT INTO print_card_faces (lang, card_name, oracle_id, set_id, card_faces_card_name, collector_number) VALUES ($1, $2, $3, $4, $5, $6)", print.Lang, print.CardName, print.OracleId, print.SetId, card_faces.Name, print.CollectorNumber)
			if err != nil {
				return err
			}
			row.Close()
		}
	}

	for i := 0; i < len(print.Finishes); i++ {
		row, err = db.Query("INSERT INTO print_finish (card_name, oracle_id, set_id, finish, lang, collector_number) VALUES ($1, $2, $3, $4, $5, $6)", print.CardName, print.OracleId, print.SetId, print.Finishes[i], print.Lang, print.CollectorNumber)
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.PromoTypes); i++ {
		row, err = db.Query("INSERT INTO print_promo (card_name, oracle_id, set_id, promo, lang, collector_number) VALUES ($1, $2, $3, $4, $5, $6)", print.CardName, print.OracleId, print.SetId, print.PromoTypes[i], print.Lang, print.CollectorNumber)
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.MultiverseIds); i++ {
		row, err = db.Query("INSERT INTO print_multiverse_id (card_name, oracle_id, set_id, multiverse_id, lang, collector_number) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (card_name, oracle_id, set_id, multiverse_id, lang, collector_number) DO NOTHING", print.CardName, print.OracleId, print.SetId, print.MultiverseIds[i], print.Lang, print.CollectorNumber)
		if err != nil {
			return err
		}
		row.Close()
	}

	return nil
}
