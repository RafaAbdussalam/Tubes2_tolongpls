import React from 'react';
import './ModeToggle.css';

function ModeToggle({ mode, onModeChange }) {
    return (
        <div className="mode-toggle">
            <button
                className={`toggle-button ${mode === 'shortest' ? 'active' : ''}`}
                onClick={() => onModeChange('shortest')}
            >
                Shortest Path
            </button>
            <button
                className={`toggle-button ${mode === 'multiple' ? 'active' : ''}`}
                onClick={() => onModeChange('multiple')}
            >
                Multiple Paths
            </button>
        </div>
    );
};

export default ModeToggle;