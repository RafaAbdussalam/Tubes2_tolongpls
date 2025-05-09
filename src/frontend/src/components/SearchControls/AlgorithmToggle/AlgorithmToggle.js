import React from 'react';
import './AlgorithmToggle.css';

function AlgorithmToggle({ algorithm, onChange }) {
    return (
        <div className="algorithm-toggle">
            <button
                className={`toggle-button ${algorithm === 'bfs' ? 'active' : ''}`}
                onClick={() => onChange('bfs')}
            >
                BFS
            </button>
            <button
                className={`toggle-button ${algorithm === 'dfs' ? 'active' : ''}`}
                onClick={() => onChange('dfs')}
            >
                DFS
            </button>
        </div>
    );
};

export default AlgorithmToggle;