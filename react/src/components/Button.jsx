import "./Button.css";
import axios from "axios";

function Button({ onClick }) {
  const getCard = async () => {
    const response = await axios.get("http://localhost:8000/random-card");
    console.log(response.data);
    onClick(response.data);
  };

  return (
    <div className="button-container">
      <button className="random-button" onClick={getCard}>
        Random Card
      </button>
    </div>
  );
}

export default Button;
