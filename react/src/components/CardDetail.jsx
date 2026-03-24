import "./CardDetail.css";
function CardDetail({ card, imageUrl, backImageUrl }) {
  return (
    <div class="grid-container">
      <div class="grid-item">1</div>
      <div class="grid-item">
        <p className="card-name">{card.CardName}</p>
      </div>
      <div class="grid-item">3</div>
      <div class="grid-item">
        {card.Artist && <p className="card-artist">Artist: {card.Artist}</p>}
        {card.PriceUsd && (
          <p className="card-price-usd">Price: ${card.PriceUsd}</p>
        )}
      </div>
      <div class="grid-item">
        <div className="card-box">
          <div
            className="image-container" /*onMouseEnter={setIsFlipped(true)} onMouseLeave={setIsFlipped(false)}*/
          >
            <div className="image-container-inner">
              <div className="image-front">
                {imageUrl && <img src={imageUrl} />}
              </div>

              <div className="image-back">
                {backImageUrl && <img src={backImageUrl} />}
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="grid-item">6</div>
    </div>
  );
}

export default CardDetail;
