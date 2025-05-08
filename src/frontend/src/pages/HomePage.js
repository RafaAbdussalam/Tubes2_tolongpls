import React, { useState } from 'react';
import SearchBar from '../components/SearchBar/SearchBar';
import AlgorithmToggle from '../components/AlgorithmToggle/AlgorithmToggle';
import ModeToggle from '../components/ModeToggle/ModeToggle';
import MaxPathsInput from '../components/MaxPathsInput/MaxPathsInput';
import SearchControls from '../components/SearchControls/SearchControls';
import '../styles/HomePage.css';

const HomePage = () => {
  const [algorithm, setAlgorithm] = useState('bfs');
  const [mode, setMode] = useState('shortest');
  const [maxPaths, setMaxPaths] = useState(3);

  const handleSearch = (searchTerm) => {
    console.log('Search Params:', {
      element: searchTerm,
      algorithm,
      mode,
      maxPaths: mode === 'multiple' ? maxPaths : undefined
    });
    // Akan dihubungkan ke backend nanti
  };

  return (
    <div className="home-page">
      <h1>Little Alchemy Recipe Finder</h1>
      
      <SearchControls>
        <SearchBar onSubmit={handleSearch} />
        <AlgorithmToggle 
          algorithm={algorithm} 
          onAlgorithmChange={setAlgorithm} 
        />
        <ModeToggle 
          mode={mode} 
          onModeChange={setMode} 
        />
        {mode === 'multiple' && (
          <MaxPathsInput 
            maxPaths={maxPaths} 
            onMaxPathsChange={setMaxPaths} 
          />
        )}
      </SearchControls>
    </div>
  );
};

export default HomePage;