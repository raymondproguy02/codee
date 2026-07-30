import json
import sys
from pathlib import Path

def prettify(input_file, output_file=None, indent=2):
    with open(input_file, 'r') as f:
        data = json.load(f)
    
    pretty = json.dumps(data, indent=indent, sort_keys=True)
    
    if output_file:
        with open(output_file, 'w') as f:
            f.write(pretty)
        print(f"✅ Written to: {output_file}")
    else:
        print(pretty)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python json_pretty.py <input.json> [output.json]")
        sys.exit(1)
    
    input_path = sys.argv[1]
    output_path = sys.argv[2] if len(sys.argv) > 2 else None
    prettify(input_path, output_path)
