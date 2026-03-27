import { useState } from "react";
import "./RandomCardPage.css";
import Button from "./Button.jsx";
import CardDetail from "./CardDetail.jsx";

function RandomCardPage() {
  const [card, setCard] = useState({});
  const [imageUrl, setImageUrl] = useState("");
  const [backImageUrl, setBackImageUrl] = useState("");

  const getImageUrl = (card) => {
    const prints = card.Prints
    const english = prints.filter(p => { p.Lang === 'en' })
    const highres = english.filter(p => p.HiresImage)
    var print
    if (highres.length > 0) {
      print = highres[0]
    } else if (english.length > 0) {
      print = english[0]
    } else {
      print = prints[0]
    }
    return print.CardFaces?.length > 0 ? print.CardFaces[0].NormalUri : print.NormalUri
  }

  const getCard = (card) => {
    setCard(card);
    setImageUrl(getImageUrl(card));
    setBackImageUrl(
      "https://backs.scryfall.io/normal/2/2/222b7a3b-2321-4d4c-af19-19338b134971.jpg?1677416389",
    );
  };

  return (
    <div>
      <Button onClick={getCard} />
      <CardDetail card={card} imageUrl={imageUrl} backImageUrl={backImageUrl} />
    </div>
  );
}
export default RandomCardPage;
