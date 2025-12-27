# dot-to-xyflow Examples

This directory contains reference implementations for using the JSON output of the `dot-to-xyflow` translator.

## Contents

- `sample.dot`: A sample Graphviz workflow.
- `sample_output.json`: The result of running the Go translator on `sample.dot`.
- `ReactFlowExample.tsx`: A standard React component using `@xyflow/react`.
- `LitFlowExample.ts`: A Web Component using `lit-flow`.

## Siblings

These examples are designed to work with [LitFlow](https://github.com/ghchinoy/litflow), a custom Lit-based implementation of the XYFlow system.

For a live side-by-side comparison, see the [test-bench](./test-bench) directory.
