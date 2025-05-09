import React from 'react';
import './ModeToggle.css';

function ModeToggle({ mode, onChange }) {
    return (
        <div className="mode-toggle">
            <button
                className={`toggle-button ${mode === 'single' ? 'active' : ''}`}
                onClick={() => onChange('single')}
            >
                Single Path
            </button>
            <button
                className={`toggle-button ${mode === 'multiple' ? 'active' : ''}`}
                onClick={() => onChange('multiple')}
            >
                Multi Paths
            </button>
        </div>
    );
};

export default ModeToggle;