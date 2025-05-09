import React from 'react';
import './ModeToggle.css';

function ModeToggle({ mode, onChange }) {
    return (
        <div className="mode-toggle">
            <button
                className={`toggle-button ${mode === 'shortest' ? 'active' : ''}`}
                onClick={() => onChange('shortest')}
            >
                Shortest Path
            </button>
            <button
                className={`toggle-button ${mode === 'multiple' ? 'active' : ''}`}
                onClick={() => onChange('multiple')}
            >
                Multiple Paths
            </button>
        </div>
    );
};

export default ModeToggle;