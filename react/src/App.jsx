import { useState } from 'react'
import './App.css'
import axios from 'axios'

function App() {
  const [card, setCard] = useState({})

  const getCard = async () => {
    const response = await axios.get("http://localhost:8000/random-card")
    console.log(response.data)
    setCard(response.data)
  }

  const getNormalUri = () => {
    if (card.CardFaces != null) {
      return card.CardFaces.NormalUri
    }
    return card.NormalUri
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
          <img className="card-img" src={getNormalUri()} />
        </div>

      </div>
    </div>
  )
}

export default App
