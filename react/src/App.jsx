import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [card, setCard] = useState({})

  const getCard = () => {

  }

  return (
    <>
      <div className="card">
        <button onClick={() => setCard(getCard())}>
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
    </>
  )
}

export default App
