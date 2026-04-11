import "./CardList.css";
import { useState, useEffect } from "react";
import Card from "./Card.jsx";

// const perPage = 10;

const getPerPage = () => {
  const cols = Math.min(Math.floor(window.innerWidth / 250), 5);
  const rows = window.innerHeight < 700 || window.innerWidth < 800 ? 1 : 2;
  return cols * rows;
};

function CardList() {
  const [images] = useState([
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
    "https://cards.scryfall.io/large/front/3/5/35c7c392-6782-40c8-bb24-6aad24f14660.jpg?1562784760a",
  ]);
  const [perPage, setPerPage] = useState(getPerPage);
  const [page, setPage] = useState(0);

  useEffect(() => {
    const handleResize = () => {
      setPerPage(getPerPage());
      setPage(0);
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  const totalPages = Math.ceil(images.length / perPage);
  const start = page * perPage; //page 0 = 0-8, page 1 = 9-17 and so on
  const visible = images.slice(start, start + perPage); //get 9 items from index, slice hides the rest

  return (
    <div className="image-grid-wrapper">
      <p className="results-count">{images.length} results</p>

      <div className="image-grid">
        {visible.map((url, index) => (
          <div key={index} className="image-card">
            <img src={url} alt={`Card ${start + index + 1}`} loading="lazy" />
          </div>
        ))}
      </div>

      <div className="pagination">
        <button onClick={() => setPage((p) => p - 1)} disabled={page === 0}>
          Prev
        </button>
        <span className="page-info">
          Page {page + 1} of {totalPages}
        </span>
        <button
          onClick={() => setPage((p) => p + 1)}
          disabled={page >= totalPages - 1}
        >
          Next
        </button>
      </div>
    </div>
    // <div>
    //   {images.map((url, index) => (
    //     <Card url={url} index={index} />
    //   ))}
    // </div>
  );
}
export default CardList;
