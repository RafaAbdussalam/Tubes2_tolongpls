import React from 'react';
import './StatsPanel.css';

function StatsPanel({ stats }) {
    return (
        <div className="stats-panel">
            <h3>Search Statistics</h3>
            <div className="stats-grid">
                <div className="stat-item">
                    <span className="stat-label">Algorithm:</span>
                    <span className="stat-value">{stats.algorithm}</span>
                </div>
                <div className="stat-item">
                    <span className="stat-label">Mode:</span>
                    <span className="stat-value">{stats.mode}</span>
                </div>
                <div className="stat-item">
                    <span className="stat-label">Time:</span>
                    <span className="stat-value">{stats.time} ms</span>
                </div>
                <div className="stat-item">
                    <span className="stat-label">Nodes Visited:</span>
                    <span className="stat-value">{stats.nodes}</span>
                </div>
                <div className="stat-item">
                    <span className="stat-label">Recipe Count:</span>
                    <span className="stat-value">{stats.recipeCount || 1}</span>
                </div>
            </div>
        </div>
    );
};

export default StatsPanel;