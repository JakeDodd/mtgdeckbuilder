import { useState } from 'react'
import './App.css'
import axios from 'axios'
import Button from './components/Button.jsx'

function App() {
  const [card, setCard] = useState({})
  const [imageUrl, setImageUrl] = useState("")
  const [backImageUrl, setBackImageUrl] = useState("https://backs.scryfall.io/normal/2/2/222b7a3b-2321-4d4c-af19-19338b134971.jpg?1677416389")

  const getCard = async () => {
    const response = await axios.get("http://localhost:8000/random-card")
    console.log(response.data)
    setCard(response.data)
    setImageUrl(response.data.CardFaces?.length > 0 ? response.data.CardFaces[0].NormalUri : response.data.NormalUri)
  }

  return (
    <div className="border-style" >
      <div className="container">
        <Button onClick={getCard} />
        <div className="card-box">
          <p className="card-name">
            {card.CardName}
          </p>
          <div className="image-container" /*onMouseEnter={setIsFlipped(true)} onMouseLeave={setIsFlipped(false)}*/>
            <div className="image-container-inner">
              <div className="image-front">
                <img className="card-img" src={imageUrl} />
              </div>
              {/*
              <div className="image-back">
                <img className="card-img" src={backImageUrl} />
              </div>
      */}
            </div>
          </div>
        </div>

      </div>
    </div>
  )
}

export default App
