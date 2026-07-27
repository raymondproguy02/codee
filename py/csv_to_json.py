import csv
import json
import sys
from pathlib import Path

def convert_csv_to_json(csv_path, json_path=None, pretty=True):
    """
    Convert CSV file to JSON format.
    
    Args:
        csv_path: Path to input CSV file
        json_path: Path to output JSON file (optional)
        pretty: Pretty print JSON output
    
    Returns:
        List of dictionaries representing the CSV data
    """
    csv_path = Path(csv_path)
    
    if not csv_path.exists():
        raise FileNotFoundError(f"CSV file not found: {csv_path}")
    
    # Read CSV
    with open(csv_path, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        data = list(reader)
    
    # Determine output path
    if json_path is None:
        json_path = csv_path.with_suffix('.json')
    else:
        json_path = Path(json_path)
    
    # Write JSON
    with open(json_path, 'w', encoding='utf-8') as f:
        if pretty:
            json.dump(data, f, indent=2, ensure_ascii=False)
        else:
            json.dump(data, f, ensure_ascii=False)
    
    print(f"✅ Converted {csv_path} → {json_path}")
    print(f"📊 {len(data)} records processed")
    
    return data

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python csv_to_json.py <input.csv> [output.json]")
        sys.exit(1)
    
    csv_file = sys.argv[1]
    json_file = sys.argv[2] if len(sys.argv) > 2 else None
    
    try:
        convert_csv_to_json(csv_file, json_file)
    except Exception as e:
        print(f"❌ Error: {e}")
        sys.exit(1)
