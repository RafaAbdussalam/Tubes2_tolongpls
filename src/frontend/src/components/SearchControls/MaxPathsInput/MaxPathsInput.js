import React from 'react';
import './MaxPathsInput.css';

function MaxPathsInput({ maxPaths, onMaxPathsChange }) {
    return (
        <div className="max-paths-input">
            <label htmlFor="maxPaths">Max Paths:</label>
            <input
                type="number"
                id="maxPaths"
                min="1"
                max="20"
                value={maxPaths}
                onChange={(e) => onMaxPathsChange(e.target.value)}
            />
        </div>
    );
};

export default MaxPathsInput;