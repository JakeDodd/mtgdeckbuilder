import "./Card.css";

function Card({ url, index }) {
  return (
    <div>
      <img
        key={index} // Using index as key
        src={url}
        style={{ width: "150px", height: "150px", objectFit: "cover" }}
      />
    </div>
  );
}

export default Card;
