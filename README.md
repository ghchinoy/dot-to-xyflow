# dot-to-xyflow

A high-performance translator that converts Graphviz `.dot` files into XYFlow-compatible JSON. This tool bridges the gap between static architecture diagrams and interactive web-based flow visualizations.

## Core Features

- **Go-powered**: Written in Go for maximum speed and zero dependencies.
- **XYFlow Compatible**: Outputs JSON structure ready for `@xyflow/react` and `@xyflow/system`.
- **LitFlow Support**: Designed to work seamlessly with the sibling `litflow` project.
- **Intelligent Mapping**: Translates Graphviz shapes, labels, and directions into appropriate XYFlow node types and properties.

## Getting Started

### Prerequisites
- [Go](https://go.dev/) installed on your system.

### Usage
Run the translator against any `.dot` file:

```bash
go run dot_to_xyflow.go path/to/your/file.dot > output.json
```

## Integration Examples

### React (XYFlow)
Import the generated JSON directly into your React components:

```tsx
import { ReactFlow } from '@xyflow/react';
import flowData from './output.json';

// Use flowData.nodes and flowData.edges in your <ReactFlow /> component
```

### Lit (LitFlow)
Pass the nodes and edges as properties to your `lit-flow` element:

```ts
import flowData from './output.json';

// In your render method:
html`<lit-flow .nodes="${flowData.nodes}" .edges="${flowData.edges}"></lit-flow>`
```

## Development & HMR

To work on new features for `litflow` (like adding edge labels) while viewing them live in the test-bench:

1. Open `examples/test-bench/main.tsx`.
2. Uncomment the relative import: `import '../../../litflow/src/lit-flow'`.
3. Comment out the npm import: `import '@ghchinoy/litflow'`.
4. Update `vite.config.ts` to allow the parent directory in `server.fs.allow`.

This enables cross-project Hot Module Replacement.

## Project Structure

- `dot_to_xyflow.go`: The primary Go translator.
- `dot_to_xyflow.py`: A legacy Python reference implementation.
- `examples/`:
    - `sample.dot`: A representative Graphviz workflow.
    - `sample_output.json`: The generated result of the translation.
    - `ReactFlowExample.tsx`: Reference React implementation.
    - `LitFlowExample.ts`: Reference Lit implementation.
