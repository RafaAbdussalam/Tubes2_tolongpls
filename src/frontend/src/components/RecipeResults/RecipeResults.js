import React from 'react';
import StatsPanel from './StatsPanel/StatsPanel';
import RecipeTree from './RecipeTree/RecipeTree';
import ResetButton from './ResetButton/ResetButton';
import './RecipeResults.css';

function RecipeResults({ results, onReset }) {
    if (!results) return null;

    return (
        <div className="recipe-results">
            <StatsPanel 
                stats={{
                    algorithm: results.algorithm,
                    mode: results.mode,
                    time: results.time,
                    nodes: results.nodes,
                    recipeCount: results.recipeCount
                }} 
            />
            <RecipeTree 
                data={results.treeData} 
            />
            <ResetButton 
                onClick={onReset} 
            />
        </div>
    );
};

export default RecipeResults;