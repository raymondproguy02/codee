import hashlib
import argparse
from pathlib import Path

def calculate_checksum(filepath, algorithm='sha256'):
    """Calculate hash of file using specified algorithm"""
    hash_func = getattr(hashlib, algorithm)()
    
    with open(filepath, 'rb') as f:
        for chunk in iter(lambda: f.read(8192), b''):
            hash_func.update(chunk)
    
    return hash_func.hexdigest()

def verify_checksum(filepath, expected_hash, algorithm='sha256'):
    """Verify file against expected hash"""
    actual = calculate_checksum(filepath, algorithm)
    match = actual == expected_hash
    
    print(f"📁 File: {filepath}")
    print(f"🔐 Algorithm: {algorithm}")
    print(f"📊 Expected: {expected_hash}")
    print(f"✅ Actual:   {actual}")
    print(f"🔍 Match: {'✅ YES' if match else '❌ NO'}")
    
    return match

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="File checksum validator")
    parser.add_argument("file", help="File to checksum")
    parser.add_argument("-a", "--algorithm", default="sha256", 
                       choices=['md5', 'sha1', 'sha256', 'sha512'],
                       help="Hash algorithm")
    parser.add_argument("-c", "--check", help="Expected hash to verify against")
    
    args = parser.parse_args()
    
    if not Path(args.file).exists():
        print(f"❌ File not found: {args.file}")
        exit(1)
    
    if args.check:
        verify_checksum(args.file, args.check, args.algorithm)
    else:
        checksum = calculate_checksum(args.file, args.algorithm)
        print(f"{args.algorithm}: {checksum}")
