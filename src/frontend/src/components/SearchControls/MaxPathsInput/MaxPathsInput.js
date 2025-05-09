import React from 'react';
import './MaxPathsInput.css';

function MaxPathsInput({ value, onChange }) {
    return (
        <div className="max-paths-input">
            <label htmlFor="maxPaths">Max Paths:</label>
            <input
                type="number"
                id="maxPaths"
                min="1"
                max="20"
                value={value}
                onChange={(e) => onChange(e.target.value)}
            />
        </div>
    );
};

export default MaxPathsInput;