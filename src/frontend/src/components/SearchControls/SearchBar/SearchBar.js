import React from 'react';
import './SearchBar.css';

function SearchBar({ value, onChange, onSubmit }) {
    function handleSubmit(e) {
        e.preventDefault();  // Prevent default form submission
        onSubmit(e);         // Teruskan event ke parent
    };

    return (
    <form className="search-bar" onSubmit={handleSubmit}>  {/* Gunakan handler lokal */}
        <input
            type="text"
            value={value}
            onChange={onChange}
            placeholder="Search element (e.g. Brick, Water)..."
            className="search-input"
        />
        <button type="submit" className="search-button">
            Search
        </button>
    </form>
    );
};

export default SearchBar;