import os
import shutil
from pathlib import Path
from datetime import datetime
import argparse

class FileOrganizer:
    EXTENSION_MAP = {
        'Images': ['.jpg', '.jpeg', '.png', '.gif', '.bmp', '.svg', '.webp'],
        'Documents': ['.pdf', '.docx', '.doc', '.txt', '.md', '.rtf', '.odt'],
        'Spreadsheets': ['.xlsx', '.xls', '.csv', '.ods'],
        'Presentations': ['.pptx', '.ppt', '.odp'],
        'Code': ['.py', '.js', '.html', '.css', '.json', '.xml', '.yaml', '.yml', '.sh'],
        'Archives': ['.zip', '.tar', '.gz', '.rar', '.7z'],
        'Videos': ['.mp4', '.avi', '.mkv', '.mov', '.wmv'],
        'Audio': ['.mp3', '.wav', '.flac', '.aac', '.ogg'],
        'Executables': ['.exe', '.msi', '.appimage', '.deb', '.rpm'],
    }

    def __init__(self, directory, dry_run=False):
        self.directory = Path(directory)
        self.dry_run = dry_run
        self.stats = {'moved': 0, 'skipped': 0, 'errors': 0}

    def organize(self):
        """Main orchestration method"""
        if not self.directory.exists():
            raise FileNotFoundError(f"Directory not found: {self.directory}")
        
        print(f"📁 Organizing: {self.directory}")
        print(f"🧪 {'DRY RUN' if self.dry_run else 'LIVE'}\n")
        
        for item in self.directory.iterdir():
            if item.is_file():
                self._process_file(item)
            elif item.is_dir() and item.name in self.EXTENSION_MAP:
                # Skip already organized folders
                continue
        
        self._print_summary()
        return self.stats

    def _process_file(self, file_path):
        """Move single file to appropriate folder"""
        ext = file_path.suffix.lower()
        folder_name = self._get_folder_name(ext)
        
        if not folder_name:
            print(f"⏭️  Skipped: {file_path.name} (unknown extension)")
            self.stats['skipped'] += 1
            return

        target_folder = self.directory / folder_name
        target_folder.mkdir(exist_ok=True)
        
        target_path = target_folder / file_path.name
        
        # Handle duplicates
        if target_path.exists():
            timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
            new_name = f"{file_path.stem}_{timestamp}{file_path.suffix}"
            target_path = target_folder / new_name
        
        try:
            if not self.dry_run:
                shutil.move(str(file_path), str(target_path))
            self.stats['moved'] += 1
            print(f"✅ {file_path.name} → {folder_name}/")
        except Exception as e:
            self.stats['errors'] += 1
            print(f"❌ Error moving {file_path.name}: {e}")

    def _get_folder_name(self, ext):
        """Determine folder based on file extension"""
        for folder, extensions in self.EXTENSION_MAP.items():
            if ext in extensions:
                return folder
        return None

    def _print_summary(self):
        """Display execution summary"""
        print("\n" + "="*40)
        print("📊 SUMMARY")
        print(f"   Moved: {self.stats['moved']} files")
        print(f"   Skipped: {self.stats['skipped']} files")
        print(f"   Errors: {self.stats['errors']} files")
        print(f"   Total: {sum(self.stats.values())} files")
        if self.dry_run:
            print("⚠️  DRY RUN - No files were actually moved")

def main():
    parser = argparse.ArgumentParser(description="Organize files by extension")
    parser.add_argument('directory', help="Directory to organize")
    parser.add_argument('--dry-run', '-d', action='store_true', 
                       help="Simulate without moving files")
    parser.add_argument('--add-ext', '-a', nargs=2, metavar=('FOLDER', 'EXT'),
                       help="Add custom folder-extension mapping (e.g., 'PDFs .pdf')")
    
    args = parser.parse_args()
    
    if args.add_ext:
        folder, ext = args.add_ext
        ext = ext if ext.startswith('.') else f'.{ext}'
        FileOrganizer.EXTENSION_MAP.setdefault(folder, []).append(ext)
        print(f"➕ Added: {folder} → {ext}")
        return
    
    organizer = FileOrganizer(args.directory, dry_run=args.dry_run)
    organizer.organize()

if __name__ == "__main__":
    main()
