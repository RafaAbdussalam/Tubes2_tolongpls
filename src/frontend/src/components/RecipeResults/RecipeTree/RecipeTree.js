import React from 'react';
import { Tree } from 'react-d3-tree';
import { formatTree } from '../../../utils/formatTree';
import './RecipeTree.css';

// Render custom
function renderCustomNode({ nodeDatum }) {
    const isRecipe = nodeDatum.attributes?.type === 'recipe';
    const isBase = nodeDatum.attributes?.type === 'base-element';

    return (
        <g>
            <circle
                r="10"
                fill={isRecipe ? 'lightblue' : isBase ? 'lightgreen' : 'white'}
                stroke={isRecipe ? 'steelblue' : 'darkgreen'}
                strokeWidth="2"
            />
            <text
                x={isRecipe ? 25 : 20}
                dy={isRecipe ? '.31em' : '.35em'}
                fill={isRecipe ? 'navy' : 'black'}
                fontSize={isRecipe ? '20px' : '20px'}
            >
                {nodeDatum.name}
            </text>
        </g>
    );
};

function RecipeTree({ data }) {
    return (
        <div className="recipe-tree-container">
            <Tree
                data={formatTree(data)}
                orientation="vertical"
                pathFunc="step"
                translate={{ x: 200, y: 50 }}
                renderCustomNodeElement={renderCustomNode}
                className="recipe-tree"
            />
        </div>
    );
};

export default RecipeTree;