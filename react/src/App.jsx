import { useState } from "react";
import "./App.css";
import Button from "./components/Button.jsx";
import CardDetail from "./components/CardDetail.jsx";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import RandomCardPage from "./components/RandomCardPage.jsx";
import CardSearchPage from "./components/CardSearchPage.jsx";

function App() {

  return (
    <BrowserRouter>
      <div className="border-style">
        <div className="container">
          <Routes>
            <Route path="/random" element={<RandomCardPage />} />
            <Route path="/search" element={<CardSearchPage />} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
