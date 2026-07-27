import re
import math
import argparse
from typing import Dict, List, Tuple

class PasswordChecker:
    """Advanced password strength analyzer with entropy calculation"""
    
    COMMON_PASSWORDS = {
        'password', '123456', 'qwerty', 'admin', 'letmein', 
        'welcome', 'monkey', 'dragon', 'master', 'hello'
    }
    
    def __init__(self, password: str):
        self.password = password
        self.length = len(password)
        self.score = 0
        self.feedback = []
        self.entropy = 0.0
        
    def check(self) -> Dict:
        """Run all checks and return results"""
        self._check_length()
        self._check_complexity()
        self._check_common()
        self._check_patterns()
        self._calculate_entropy()
        
        return {
            'password': '*' * self.length,
            'length': self.length,
            'score': self.score,
            'max_score': 10,
            'strength': self._get_strength_label(),
            'entropy_bits': round(self.entropy, 2),
            'feedback': self.feedback,
            'is_secure': self.score >= 7
        }
    
    def _check_length(self):
        """Check password length"""
        if self.length >= 12:
            self.score += 3
            self.feedback.append("✅ Excellent length (12+ chars)")
        elif self.length >= 8:
            self.score += 2
            self.feedback.append("✅ Good length (8-11 chars)")
        elif self.length >= 6:
            self.score += 1
            self.feedback.append("⚠️ Minimum length (6 chars)")
        else:
            self.feedback.append("❌ Too short (minimum 6 chars)")
    
    def _check_complexity(self):
        """Check character variety"""
        checks = {
            'lowercase': bool(re.search(r'[a-z]', self.password)),
            'uppercase': bool(re.search(r'[A-Z]', self.password)),
            'digits': bool(re.search(r'\d', self.password)),
            'special': bool(re.search(r'[!@#$%^&*()_+\-=\[\]{};:\'",.<>?/\\|`~]', self.password))
        }
        
        count = sum(checks.values())
        self.score += count
        
        if count == 4:
            self.feedback.append("✅ All character types used")
        elif count == 3:
            self.feedback.append("✅ Good variety (3 types)")
        elif count == 2:
            self.feedback.append("⚠️ Limited variety (2 types)")
        elif count <= 1:
            self.feedback.append("❌ Poor variety (1 type only)")
        
        # Add specific feedback
        if not checks['uppercase']:
            self.feedback.append("⚠️ Add uppercase letters")
        if not checks['lowercase']:
            self.feedback.append("⚠️ Add lowercase letters")
        if not checks['digits']:
            self.feedback.append("⚠️ Add numbers")
        if not checks['special']:
            self.feedback.append("⚠️ Add special characters")
    
    def _check_common(self):
        """Check against common passwords and patterns"""
        if self.password.lower() in self.COMMON_PASSWORDS:
            self.score -= 3
            self.feedback.append("❌ Common password detected")
        
        # Check for keyboard patterns
        keyboard_patterns = ['qwerty', 'asdfgh', 'zxcvbn', '123456']
        for pattern in keyboard_patterns:
            if pattern in self.password.lower():
                self.score -= 2
                self.feedback.append(f"❌ Keyboard pattern detected: {pattern}")
                break
        
        # Check for repeated characters
        if re.search(r'(.)\1{3,}', self.password):
            self.score -= 1
            self.feedback.append("⚠️ Repeated characters detected")
    
    def _check_patterns(self):
        """Check for common patterns"""
        # Sequential numbers
        if re.search(r'012|123|234|345|456|567|678|789|890', self.password):
            self.score -= 1
            self.feedback.append("⚠️ Sequential numbers detected")
        
        # Year patterns
        if re.search(r'(19|20)\d{2}', self.password):
            self.score -= 1
            self.feedback.append("⚠️ Year pattern detected")
    
    def _calculate_entropy(self):
        """Calculate password entropy in bits"""
        charset_size = 0
        if re.search(r'[a-z]', self.password):
            charset_size += 26
        if re.search(r'[A-Z]', self.password):
            charset_size += 26
        if re.search(r'\d', self.password):
            charset_size += 10
        if re.search(r'[!@#$%^&*()_+\-=\[\]{};:\'",.<>?/\\|`~]', self.password):
            charset_size += 32
        
        if charset_size > 0:
            self.entropy = self.length * math.log2(charset_size)
    
    def _get_strength_label(self) -> str:
        """Get human-readable strength label"""
        if self.score >= 8:
            return "🏆 EXCELLENT"
        elif self.score >= 6:
            return "✅ GOOD"
        elif self.score >= 4:
            return "⚠️ WEAK"
        else:
            return "❌ VERY WEAK"

def main():
    parser = argparse.ArgumentParser(description="Check password strength")
    parser.add_argument('password', nargs='?', help="Password to check")
    parser.add_argument('--interactive', '-i', action='store_true', 
                       help="Interactive mode with hidden input")
    
    args = parser.parse_args()
    
    if args.interactive:
        import getpass
        password = getpass.getpass("Enter password to check: ")
    elif args.password:
        password = args.password
    else:
        print("Usage: python password_checker.py <password>")
        print("       python password_checker.py --interactive")
        return
    
    checker = PasswordChecker(password)
    result = checker.check()
    
    print("\n" + "="*50)
    print(f"🔐 PASSWORD STRENGTH REPORT")
    print("="*50)
    print(f"Password: {result['password']}")
    print(f"Length: {result['length']} chars")
    print(f"Entropy: {result['entropy_bits']} bits")
    print(f"Score: {result['score']}/10")
    print(f"Strength: {result['strength']}")
    print("\nFeedback:")
    for msg in result['feedback']:
        print(f"  {msg}")
    print("="*50)
    
    if result['is_secure']:
        print("✅ Password is secure!")
    else:
        print("❌ Password needs improvement!")

if __name__ == "__main__":
    main()
