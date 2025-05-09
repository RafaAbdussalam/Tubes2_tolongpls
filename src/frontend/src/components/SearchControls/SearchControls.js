import React from 'react';
import SearchBar from './SearchBar/SearchBar';
import AlgorithmToggle from './AlgorithmToggle/AlgorithmToggle';
import ModeToggle from './ModeToggle/ModeToggle';
import MaxPathsInput from './MaxPathsInput/MaxPathsInput';
import './SearchControls.css';

function SearchControls({ 
        searchTerm,
        onSearchChange,
        onSearchSubmit,
        algorithm,
        onAlgorithmChange,
        mode,
        onModeChange,
        maxPaths,
        onMaxPathsChange
    }) {
    return (
        <div className="search-controls">
            <SearchBar 
                value={searchTerm}
                onChange={onSearchChange}
                onSubmit={onSearchSubmit}
            />
            <div className="search-options">
                <AlgorithmToggle 
                    algorithm={algorithm} 
                    onChange={onAlgorithmChange} 
                />
                <ModeToggle 
                    mode={mode} 
                    onChange={onModeChange} 
                />
                {mode === 'multiple' && (
                    <MaxPathsInput 
                        value={maxPaths}
                        onChange={onMaxPathsChange}
                    />
                )}
            </div>
        </div>
    );
};

export default SearchControls;