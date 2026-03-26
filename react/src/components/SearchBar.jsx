import "./SearchBar.css";
import axios from "axios";

function SearchBar() {
  return (
    <div class="card-search">
      <input type="text" placeholder="Ex. lightning bolt"></input>
      <button type="submit">Search</button>
    </div>
  );
}

export default SearchBar;
