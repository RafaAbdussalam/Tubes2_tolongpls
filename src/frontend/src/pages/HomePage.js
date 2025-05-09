import React, { useState } from 'react';
import SearchControls from '../components/SearchControls/SearchControls';
import RecipeResults from '../components/RecipeResults/RecipeResults';
import { fetchRecipe } from '../utils/api';
import '../styles/HomePage.css';

// untuk uji coba tanpa API
// import data1 from '../utils/contoh1.json';
// import data2 from '../utils/contoh2.json';

function HomePage() {
    // State management
    const [searchTerm, setSearchTerm] = useState('');
    const [algorithm, setAlgorithm] = useState('bfs');
    const [mode, setMode] = useState('single');
    const [modeSubmitted, setModeSubmitted] = useState('single');
    const [maxPaths, setMaxPaths] = useState(2);
    const [results, setResults] = useState(null);
    const [isLoading, setIsLoading] = useState(false);
    
    // Handlers
    // ini siap dicoba sama API
    function handleSearchSubmit(e) {
        e.preventDefault();
        if (!searchTerm.trim()) return;
        setModeSubmitted(mode);
        setIsLoading(true);
        fetchRecipe({
            element: searchTerm,
            algorithm,
            maxPaths: mode === 'multiple' ? maxPaths : 1
        })
        .then(setResults)
        .catch(console.error)
        .finally(() => setIsLoading(false));
    };
    
    // dummy data
    // function handleSearchSubmit(e) {
        //     e.preventDefault();
        //     if (!searchTerm.trim()) return;
        
        //     setIsLoading(true);
        //     if (searchTerm.length % 2 === 1) {
            //         setResults(data1);
            //     } else {
                //         setResults(data2);
                //     }
                //     setIsLoading(false);
                // }
                
                function handleReset() {
                    setSearchTerm('');
                    setAlgorithm('bfs');
                    setMode('single');
                    setModeSubmitted('single');
                    setMaxPaths(2);
                    setResults(null);
                    setIsLoading(false);
                };
                
                return (
                    <div className="home-page">
            <h1>Little Alchemy Recipe Finder</h1>
            
            <SearchControls
                // SearchBar props
                searchTerm={searchTerm}
                onSearchChange={(e) => setSearchTerm(e.target.value)}
                onSearchSubmit={handleSearchSubmit}
                
                // AlgorithmToggle props
                algorithm={algorithm}
                onAlgorithmChange={setAlgorithm}
                
                // ModeToggle props
                mode={mode}
                onModeChange={setMode}
                
                // MaxPathsInput props
                maxPaths={maxPaths}
                onMaxPathsChange={(value) => setMaxPaths(Number(value))}
            />

            {isLoading ? (
                <div className="loading">Searching recipes...</div>
            ) : (
                results && 
                <RecipeResults 
                    results={results}
                    mode={modeSubmitted} 
                    onReset={handleReset} 
                />
            )}
        </div>
    );
};

export default HomePage;