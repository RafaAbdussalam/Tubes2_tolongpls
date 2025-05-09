import React from 'react';
import './StatsPanel.css';

function StatsPanel({ stats, mode }) {
    return (
        <div className="stats-panel">
            <h3>Search Statistics</h3>
            <div className="stats-tables">
                <table className="stats-table">
                    <tbody>
                        <tr>
                            <td className="stat-label">Algorithm :</td>
                            <td className="stat-value">{stats.algorithm.toUpperCase()}</td>
                        </tr>
                        <tr>
                            <td className="stat-label">Depth :</td>
                            <td className="stat-value">{stats.depth}</td>
                        </tr>
                        <tr>
                            <td className="stat-label">Nodes :</td>
                            <td className="stat-value">{stats.node_count}</td>
                        </tr>
                    </tbody>
                </table>

                <table className="stats-table">
                    <tbody>
                        <tr>
                            <td className="stat-label">Mode :</td>
                            <td className="stat-value">
                            {mode === 'single' ? 'Single Path' : 'Multiple Paths'}
                            </td>
                        </tr>
                        <tr>
                            <td className="stat-label">Recipes :</td>
                            <td className="stat-value">{stats.recipe_count}</td>
                        </tr>
                        <tr>
                            <td className="stat-label">Time :</td>
                            <td className="stat-value">{stats.time} ms</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default StatsPanel;