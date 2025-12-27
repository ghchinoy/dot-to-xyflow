# dot-to-xyflow

A high-performance translator that converts Graphviz `.dot` files into XYFlow-compatible JSON. This tool bridges the gap between static architecture diagrams and interactive web-based flow visualizations.

## Core Features

- **Go-powered**: Written in Go for maximum speed.
- **Robust Parsing**: Uses `go-graphviz` for full DOT specification support.
- **Automatic Layout**: Leverages Graphviz layout engines (dot, neato, etc.) to generate node positions.
- **XYFlow Compatible**: Outputs JSON structure ready for `@xyflow/react` and `@xyflow/system`.
- **LitFlow Support**: Designed to work seamlessly with the sibling `litflow` project.
- **Intelligent Mapping**: Translates Graphviz shapes, labels, and directions into appropriate XYFlow node types and properties.

## Getting Started

### Prerequisites
- [Go](https://go.dev/) installed on your system.
- [Graphviz](https://graphviz.org/) (optional, but recommended for layout engines).

### Installation
Clone the repository and install dependencies:

```bash
go mod tidy
```

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

## Project Structure

- `dot_to_xyflow.go`: The primary Go translator.
- `examples/`:
    - `sample.dot`: A representative Graphviz workflow.
    - `sample_output.json`: The generated result of the translation.
    - `ReactFlowExample.tsx`: Reference React implementation.
    - `LitFlowExample.ts`: Reference Lit implementation.
