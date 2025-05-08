import React, { useState } from 'react';

// SearchControls
import SearchBar from '../components/SearchControls/SearchBar/SearchBar';
import AlgorithmToggle from '../components/SearchControls/AlgorithmToggle/AlgorithmToggle';
import ModeToggle from '../components/SearchControls/ModeToggle/ModeToggle';
import MaxPathsInput from '../components/SearchControls/MaxPathsInput/MaxPathsInput';
import SearchControls from '../components/SearchControls/SearchControls';

import '../styles/HomePage.css';

function HomePage() {
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
        // Nanti dihubungin sama backend
    };

    return (
        <div className="home-page">
            <h1>Little Alchemy Recipe Finder</h1>
            
            <SearchControls>
                <SearchBar 
                    onSubmit={handleSearch} 
                />
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