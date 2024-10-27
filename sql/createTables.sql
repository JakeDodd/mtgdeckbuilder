--{"object":"card","id":"0000419b-0bba-4488-8f7a-6194544ce91e","oracle_id":"b34bb2dc-c1af-4d77-b0b3-a0fb342a5fc6","multiverse_ids":[668564],"mtgo_id":129825,"arena_id":91829,"tcgplayer_id":558404,"name":"Forest","lang":"en","released_at":"2024-08-02","uri":"https://api.scryfall.com/cards/0000419b-0bba-4488-8f7a-6194544ce91e","scryfall_uri":"https://scryfall.com/card/blb/280/forest?utm_source=api","layout":"normal","highres_image":true,"image_status":"highres_scan","image_uris":{"small":"https://cards.scryfall.io/small/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.jpg?1721427487","normal":"https://cards.scryfall.io/normal/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.jpg?1721427487","large":"https://cards.scryfall.io/large/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.jpg?1721427487","png":"https://cards.scryfall.io/png/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.png?1721427487","art_crop":"https://cards.scryfall.io/art_crop/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.jpg?1721427487","border_crop":"https://cards.scryfall.io/border_crop/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.jpg?1721427487"},"mana_cost":"","cmc":0.0,"type_line":"Basic Land — Forest","oracle_text":"({T}: Add {G}.)","colors":[],"color_identity":["G"],"keywords":[],"produced_mana":["G"],"legalities":{"standard":"legal","future":"legal","historic":"legal","timeless":"legal","gladiator":"legal","pioneer":"legal","explorer":"legal","modern":"legal","legacy":"legal","pauper":"legal","vintage":"legal","penny":"legal","commander":"legal","oathbreaker":"legal","standardbrawl":"legal","brawl":"legal","alchemy":"legal","paupercommander":"legal","duel":"legal","oldschool":"not_legal","premodern":"legal","predh":"legal"},"games":["paper","mtgo","arena"],"reserved":false,"foil":true,"nonfoil":true,"finishes":["nonfoil","foil"],"oversized":false,"promo":false,"reprint":true,"variation":false,"set_id":"a2f58272-bba6-439d-871e-7a46686ac018","set":"blb","set_name":"Bloomburrow","set_type":"expansion","set_uri":"https://api.scryfall.com/sets/a2f58272-bba6-439d-871e-7a46686ac018","set_search_uri":"https://api.scryfall.com/cards/search?order=set&q=e%3Ablb&unique=prints","scryfall_set_uri":"https://scryfall.com/sets/blb?utm_source=api","rulings_uri":"https://api.scryfall.com/cards/0000419b-0bba-4488-8f7a-6194544ce91e/rulings","prints_search_uri":"https://api.scryfall.com/cards/search?order=released&q=oracleid%3Ab34bb2dc-c1af-4d77-b0b3-a0fb342a5fc6&unique=prints","collector_number":"280","digital":false,"rarity":"common","card_back_id":"0aeebaf5-8c7d-4636-9e82-8c27447861f7","artist":"David Robert Hovey","artist_ids":["22ab27e3-6476-48f1-a9f7-9a9e86339030"],"illustration_id":"fb2b1ca2-7440-48c2-81c8-84da0a45a626","border_color":"black","frame":"2015","full_art":true,"textless":false,"booster":true,"story_spotlight":false,"prices":{"usd":"0.20","usd_foil":"0.36","usd_etched":null,"eur":null,"eur_foil":null,"tix":"0.02"},"related_uris":{"gatherer":"https://gatherer.wizards.com/Pages/Card/Details.aspx?multiverseid=668564&printed=false","tcgplayer_infinite_articles":"https://tcgplayer.pxf.io/c/4931599/1830156/21018?subId1=api&trafcat=infinite&u=https%3A%2F%2Finfinite.tcgplayer.com%2Fsearch%3FcontentMode%3Darticle%26game%3Dmagic%26partner%3Dscryfall%26q%3DForest","tcgplayer_infinite_decks":"https://tcgplayer.pxf.io/c/4931599/1830156/21018?subId1=api&trafcat=infinite&u=https%3A%2F%2Finfinite.tcgplayer.com%2Fsearch%3FcontentMode%3Ddeck%26game%3Dmagic%26partner%3Dscryfall%26q%3DForest","edhrec":"https://edhrec.com/route/?cc=Forest"},"purchase_uris":{"tcgplayer":"https://tcgplayer.pxf.io/c/4931599/1830156/21018?subId1=api&u=https%3A%2F%2Fwww.tcgplayer.com%2Fproduct%2F558404%3Fpage%3D1","cardmarket":"https://www.cardmarket.com/en/Magic/Products/Search?referrer=scryfall&searchString=Forest&utm_campaign=card_prices&utm_medium=text&utm_source=scryfall","cardhoarder":"https://www.cardhoarder.com/cards/129825?affiliate_id=scryfall&ref=card-profile&utm_campaign=affiliate&utm_medium=card&utm_source=scryfall"}},
DROP TABLE IF EXISTS print_lang, prints, card_keywords, cards, mtg_set, keywords;
DROP TYPE IF EXISTS related_uris, prices, legality, image_uris;

CREATE TYPE image_uris as(
    png text,
    border_crop text,
    art_crop text,
    large text,
    normal text,
    small text
);
CREATE TYPE legality as(
    standard_f text,
    future_f text,
    historic_f text,
    timeless_f text,
    gladiator_f text,
    pioneer_f text,
    explorer_f text,
    modern_f text,
    legacy_f text,
    pauper_f text,
    vintage_f text,
    penny_f text,
    commander_f text,
    oathbreaker_f text,
    standardbrawl_f text,
    brawl_f text,
    alchemy_f text,
    paupercommander_f text,
    duel_f text,
    oldschool_f text,
    premodern_f text,
    predh_f text
);
CREATE TYPE prices as(
    usd text,
    usd_foil text,
    usd_etched text,
    eur text,
    eur_foil text,
    tix text

);
CREATE TYPE related_uris as(
    gatherer text,
    tcg_articles text,
    tcg_decks text,
    edhrec text,
    tcg_buy text,
    cardmarket_buy text,
    cardhoarder_buy text
);

CREATE TABLE keywords (
    keyword_id integer,
    keyword_name text not null,
    PRIMARY KEY (keyword_id)
);

CREATE TABLE mtg_set (
    set_id text,
    set_code text,
    set_name text,
    set_type text,
    set_uri text,
    set_search_uri text,
    scryfall_set_uri text,
    PRIMARY KEY (set_id)
);
CREATE TABLE cards (
    oracle_id text not null, 
    card_name text not null,
    scryfall_uri text not null,
    layout text not null,
    mana_cost text,
    cmc integer not null,
    type_line text not null,
    oracle_text text,
    colors text[],
    color_identity text[],
    produced_mana text[],
    reserved boolean not null,
    rulings_uri text,
    legalities legality,
    

    PRIMARY KEY (card_name)
);
CREATE TABLE card_keywords (
    keyword_id integer REFERENCES keywords (keyword_id),
    card_name text REFERENCES cards (card_name),
    PRIMARY KEY (keyword_id, card_name)
);
CREATE TABLE prints (
    print_id integer,
    multiverse_ids integer[],
    mtgo_id integer, 
    arena_id integer,
    tcgplayer_id integer,
    released_at text not null,
    images image_uris,
    games text[] not null,
    oversized boolean not null,
    set_id text REFERENCES mtg_set (set_id),
    collector_number text,
    digital boolean not null,
    rarity text,
    card_back_id text not null, --is the normal mtg card back considered a card in this database
    artist text,
    illustration_id text,
    border_color text,
    frame text,
    full_art boolean not null,
    textless boolean not null,
    booster boolean not null,
    story_spotlight boolean not null,
    related_articles related_uris,
    card_name text REFERENCES cards (card_name),

    PRIMARY KEY (print_id) --card_name, set_id, booster
);
CREATE TABLE print_lang (
    lang text not null,
    scryfall_uri_json text not null,
    highres_image boolean not null,
    image_status text not null,
    foil boolean not null,
    not_foil boolean not null,
    finishes text[],
    promo boolean not null,
    reprint boolean not null,
    variation boolean not null, --NOTE figure out if print can have different values foil-variation
    price prices not null,
    print_id integer REFERENCES prints (print_id),


    PRIMARY KEY (print_id, lang)
);