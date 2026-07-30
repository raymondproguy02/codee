import subprocess
import platform
import time
import sys

def ping(host, count=4):
    param = '-n' if platform.system().lower() == 'windows' else '-c'
    cmd = ['ping', param, str(count), host]
    
    try:
        result = subprocess.run(cmd, capture_output=True, text=True, timeout=10)
        return result.returncode == 0
    except:
        return False

def monitor(host, interval=2):
    print(f"🌐 Monitoring {host} (Ctrl+C to stop)")
    
    success = 0
    fail = 0
    
    try:
        while True:
            if ping(host):
                success += 1
                print(f"✅ OK - {host} [{success}/{success+fail}]")
            else:
                fail += 1
                print(f"❌ FAIL - {host} [{success}/{success+fail}]")
            time.sleep(interval)
    except KeyboardInterrupt:
        print(f"\n📊 Summary: {success} up, {fail} down")

if __name__ == "__main__":
    host = sys.argv[1] if len(sys.argv) > 1 else "google.com"
    monitor(host)
