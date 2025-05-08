import React from 'react';
import { Tree } from 'react-d3-tree';
import './RecipeTree.css';

function RecipeTree({ data }) {
    return (
        <div className="recipe-tree-container">
            <Tree
                data={data}
                orientation="vertical"
                pathFunc="step"
                collapsible={false}
                nodeSvgShape={{ shape: 'circle', shapeProps: { r: 10 } }}
                translate={{ x: 300, y: 50 }}
                styles={{
                    nodes: {
                        node: {
                            circle: {
                                fill: 'steelblue',
                                stroke: 'darkblue',
                            },
                            name: {
                                fill: 'black',
                                fontSize: '10px',
                            },
                        },
                        leafNode: {
                            circle: {
                                fill: 'darkgreen',
                                stroke: 'forestgreen',
                            },
                        },
                    },
                    links: {
                        stroke: 'slategray',
                        strokeWidth: 2,
                    },
                }}
            />
        </div>
    );
};

export default RecipeTree;