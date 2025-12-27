import re
import json
import sys

def parse_dot(dot_content):
    # Remove comments and global settings
    dot_content = re.sub(r'//.*', '', dot_content)
    dot_content = re.sub(r'rankdir=.*;', '', dot_content)
    
    nodes = {}
    edges = []
    
    # 1. Extract Explicit Nodes
    # Matches: ID [attr1="...", attr2=...]
    # We use a non-greedy match for attributes to avoid over-matching
    node_pattern = re.compile(r'^\s*(\w+)\s*\[([^\]]+)\]', re.MULTILINE)
    label_pattern = re.compile(r'label="([^"]+)"')
    shape_pattern = re.compile(r'shape=(\w+)')

    for match in node_pattern.finditer(dot_content):
        node_id, attr_string = match.groups()
        if node_id in ['node', 'graph', 'edge']: continue
        
        label_match = label_pattern.search(attr_string)
        label = label_match.group(1).replace('\\n', '\n') if label_match else node_id
        
        shape_match = shape_pattern.search(attr_string)
        shape = shape_match.group(1) if shape_match else "box"

        nodes[node_id] = {
            "id": node_id,
            "data": {"label": label},
            "position": {"x": 0, "y": 0},
            "type": "input" if "input" in node_id.lower() else ("output" if shape == "component" else "default"),
        }

    # 2. Extract Edges
    # Matches: Source -> Target [label="..."]
    edge_pattern = re.compile(r'(\w+)\s*->\s*(\w+)(?:\s*\[([^\]]+)\])?')
    for i, match in enumerate(edge_pattern.finditer(dot_content)):
        src, tgt, attr_string = match.groups()
        if src in ['node', 'graph', 'edge']: continue
        
        # Add nodes if they were used in edges but not explicitly defined
        if src not in nodes: nodes[src] = {"id": src, "data": {"label": src}, "position": {"x": 0, "y": 0}, "type": "default"}
        if tgt not in nodes: nodes[tgt] = {"id": tgt, "data": {"label": tgt}, "position": {"x": 0, "y": 0}, "type": "default"}

        edge = {
            "id": f"e{i}",
            "source": src,
            "target": tgt,
            "markerEnd": {"type": "arrowclosed"}
        }
        
        if attr_string:
            label_match = label_pattern.search(attr_string)
            if label_match:
                edge["label"] = label_match.group(1)
        
        edges.append(edge)

    return {"nodes": list(nodes.values()), "edges": edges}

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python dot_to_xyflow.py <filename.dot>")
        sys.exit(1)
        
    with open(sys.argv[1], 'r') as f:
        content = f.read()
        
    result = parse_dot(content)
    
    # Vertical layout (simpler for React Flow preview)
    for i, node in enumerate(result['nodes']):
        node['position'] = {"x": 250, "y": i * 120}

    print(json.dumps(result, indent=2))
