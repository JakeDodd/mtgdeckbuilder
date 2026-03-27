import "./SearchBar.css";
import axios from "axios";

function SearchBar() {
  return (
    <div class="grid-container">
      <div></div>
      <div class="grid-item">
        <input type="text" placeholder="Ex. lightning bolt"></input>
      </div>
      <div class="grid-item-button">
        <button type="submit">Beseech</button>
      </div>
    </div>
  );
}

export default SearchBar;
