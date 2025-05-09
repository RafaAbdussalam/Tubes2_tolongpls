import React from 'react';
import StatsPanel from './StatsPanel/StatsPanel';
import RecipeTree from './RecipeTree/RecipeTree';
import ResetButton from './ResetButton/ResetButton';
import './RecipeResults.css';

function RecipeResults({ results, onReset }) {
    if (!results) return null;

    return (
    <div className="recipe-results">
        <div className="results-header">
            <StatsPanel stats={results} />
            <ResetButton onClick={onReset} />
        </div>
        <RecipeTree data={results.tree_data} />
    </div>
    );
};

export default RecipeResults;