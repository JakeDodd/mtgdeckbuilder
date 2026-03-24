import { useState } from "react";
import "./RandomCardPage.css";
import Button from "./Button.jsx";
import CardDetail from "./CardDetail.jsx";

function RandomCardPage() {
  const [card, setCard] = useState({});
  const [imageUrl, setImageUrl] = useState("");
  const [backImageUrl, setBackImageUrl] = useState("");

  const getCard = (card) => {
    setCard(card);
    setImageUrl(
      card.CardFaces?.length > 0 ? card.CardFaces[0].NormalUri : card.NormalUri,
    );
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
