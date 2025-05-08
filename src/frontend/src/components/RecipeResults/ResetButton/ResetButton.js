import React from 'react';
import './ResetButton.css';

function ResetButton({ onClick }) {
    return (
        <button className="reset-button" onClick={onClick}>
            Reset Search
        </button>
    );
};

export default ResetButton;