import { useState } from 'react'
import './App.css'
import axios from 'axios'

function App() {
  const [card, setCard] = useState({})

  const getCard = async () => {
    const response = await axios.get("http://localhost:8000/random-card")
    setCard(response.data)
  }

  return (
    <div className="border-style" >
      <div className="container">
        <div className="button-container">
          <button className="random-button" onClick={() => getCard()}>
            Random Card
          </button>
        </div>
        <div className="card-box">
          <p className="card-name">
            {card.CardName}
          </p>
          <img className="card-img" src={card.NormalUri} />
        </div>

      </div>
    </div>
  )
}

export default App
