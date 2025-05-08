import React from 'react';
import './AlgorithmToggle.css';

function AlgorithmToggle({ algorithm, onAlgorithmChange }) {
    return (
        <div className="algorithm-toggle">
            <button
                className={`toggle-button ${algorithm === 'bfs' ? 'active' : ''}`}
                onClick={() => onAlgorithmChange('bfs')}
            >
                BFS
            </button>
            <button
                className={`toggle-button ${algorithm === 'dfs' ? 'active' : ''}`}
                onClick={() => onAlgorithmChange('dfs')}
            >
                DFS
            </button>
        </div>
    );
};

export default AlgorithmToggle;