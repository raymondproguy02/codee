import os
from pathlib import Path
import argparse

def get_size(path):
    """Calculate total size of file or directory in bytes"""
    path = Path(path)
    
    if path.is_file():
        return path.stat().st_size
    
    total = 0
    for item in path.rglob('*'):
        if item.is_file():
            total += item.stat().st_size
    return total

def format_size(bytes):
    """Convert bytes to human readable format"""
    for unit in ['B', 'KB', 'MB', 'GB', 'TB']:
        if bytes < 1024.0:
            return f"{bytes:.2f} {unit}"
        bytes /= 1024.0
    return f"{bytes:.2f} PB"

def analyze(directory, top_n=10):
    """Analyze directory and show largest items"""
    path = Path(directory)
    
    if not path.exists():
        print(f"❌ Directory not found: {directory}")
        return
    
    # Get all items with their sizes
    items = []
    for item in path.iterdir():
        try:
            size = get_size(item)
            items.append((item, size))
        except PermissionError:
            continue
    
    # Sort by size (largest first)
    items.sort(key=lambda x: x[1], reverse=True)
    
    total_size = sum(size for _, size in items)
    
    print(f"📊 Directory: {path.absolute()}")
    print(f"📦 Total size: {format_size(total_size)}")
    print(f"📁 Items: {len(items)}")
    print("\n🔝 Largest items:")
    
    for item, size in items[:top_n]:
        if size > 0:
            pct = (size / total_size) * 100
            print(f"  {format_size(size):>10}  {pct:>5.1f}%  {item.name}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Analyze directory size")
    parser.add_argument("directory", default=".", nargs="?", help="Directory to analyze")
    parser.add_argument("-n", "--top", type=int, default=10, help="Number of top items")
    args = parser.parse_args()
    
    analyze(args.directory, args.top)
