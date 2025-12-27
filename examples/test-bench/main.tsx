import React from 'react';
import { createRoot } from 'react-dom/client';
import { ReactFlow, Background, Controls } from '@xyflow/react';
import '@xyflow/react/dist/style.css';

// 1. DATA
import flowData from './sample_output.json';

// 2. REACT RENDER
function ReactApp() {
  return (
    <div style={{ width: '100%', height: '100%' }}>
      <ReactFlow
        nodes={flowData.nodes}
        edges={flowData.edges}
        fitView
      >
        <Background />
        <Controls />
      </ReactFlow>
    </div>
  );
}

const container = document.getElementById('react-root');
if (container) {
  const root = createRoot(container);
  root.render(<ReactApp />);
}

// 3. LIT RENDER (PUBLISHED LITFLOW)
import '@ghchinoy/litflow';

const litContainer = document.getElementById('lit-root');
if (litContainer) {
  litContainer.innerHTML = `
    <lit-flow 
      show-grid 
      show-controls 
      show-minimap
      style="width: 100%; height: 100%; display: block;"
    ></lit-flow>
  `;
  const el = litContainer.querySelector('lit-flow') as any;
  if (el) {
    // Map data to the real LitFlow expectations
    el.nodes = flowData.nodes;
    el.edges = flowData.edges;
  }
}