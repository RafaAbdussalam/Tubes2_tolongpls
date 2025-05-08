// Container for SearchBar
import React from 'react';
import './SearchControls.css';

function SearchControls({ children }) {
    return <div className="search-controls">{children}</div>;
};

export default SearchControls;