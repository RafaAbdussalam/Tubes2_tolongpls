import React, { useState } from 'react';
import SearchBar from '../components/SearchBar/SearchBar';
import AlgorithmToggle from '../components/AlgorithmToggle/AlgorithmToggle';
import './HomePage.css';

const HomePage = () => {
  const [algorithm, setAlgorithm] = useState('bfs');
  
  const handleSearch = (searchTerm) => {
    console.log(`Searching for: ${searchTerm} using ${algorithm}`);
    // Akan dihubungkan ke backend nanti
  };

  return (
    <div className="home-page">
      <h1>Little Alchemy Recipe Finder</h1>
      <SearchBar onSubmit={handleSearch} />
      <AlgorithmToggle 
        algorithm={algorithm} 
        onAlgorithmChange={setAlgorithm} 
      />
    </div>
  );
};

export default HomePage;