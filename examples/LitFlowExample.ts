import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import './lit-flow'; // Assuming lit-flow is in the same project

// Import the generated JSON (requires appropriate build setup like Vite)
import flowData from './sample_output.json';

@customElement('lit-flow-example')
export class LitFlowExample extends LitElement {
  static styles = css`
    :host {
      display: block;
      width: 100%;
      height: 600px;
    }
  `;

  render() {
    return html`
      <lit-flow 
        .nodes="${flowData.nodes}" 
        .edges="${flowData.edges}"
        show-grid
        show-controls
      ></lit-flow>
    `;
  }
}
