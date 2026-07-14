#!/bin/bash

# ============================================
# CREATE DIRECTORY STRUCTURE
# ============================================
mkdir -p css js/utils

echo "📁 Creating files..."

# ============================================
# index.html
# ============================================
cat > index.html << 'EOF'
[PASTE YOUR index.html CONTENT HERE]
EOF
echo "✅ index.html created"

# ============================================
# css/style.css
# ============================================
cat > css/style.css << 'EOF'
[PASTE YOUR css/style.css CONTENT HERE]
EOF
echo "✅ css/style.css created"

# ============================================
# js/config.js
# ============================================
cat > js/config.js << 'EOF'
[PASTE YOUR js/config.js CONTENT HERE]
EOF
echo "✅ js/config.js created"

# ============================================
# js/app.js
# ============================================
cat > js/app.js << 'EOF'
[PASTE YOUR js/app.js CONTENT HERE]
EOF
echo "✅ js/app.js created"

# ============================================
# js/pages.js
# ============================================
cat > js/pages.js << 'EOF'
[PASTE YOUR js/pages.js CONTENT HERE]
EOF
echo "✅ js/pages.js created"

# ============================================
# js/lessons.js
# ============================================
cat > js/lessons.js << 'EOF'
[PASTE YOUR js/lessons.js CONTENT HERE]
EOF
echo "✅ js/lessons.js created"

# ============================================
# js/verses.js
# ============================================
cat > js/verses.js << 'EOF'
[PASTE YOUR js/verses.js CONTENT HERE]
EOF
echo "✅ js/verses.js created"

# ============================================
# js/bible-api.js
# ============================================
cat > js/bible-api.js << 'EOF'
[PASTE YOUR js/bible-api.js CONTENT HERE]
EOF
echo "✅ js/bible-api.js created"

# ============================================
# js/utils/storage.js
# ============================================
cat > js/utils/storage.js << 'EOF'
[PASTE YOUR js/utils/storage.js CONTENT HERE]
EOF
echo "✅ js/utils/storage.js created"

echo ""
echo "🎉 All files created successfully!"
echo ""
echo "📁 Project structure:"
echo "├── index.html"
echo "├── css/"
echo "│   └── style.css"
echo "└── js/"
echo "    ├── app.js"
echo "    ├── config.js"
echo "    ├── pages.js"
echo "    ├── lessons.js"
echo "    ├── verses.js"
echo "    ├── bible-api.js"
echo "    └── utils/"
echo "        └── storage.js"
echo ""
echo "▶️ Run: python3 -m http.server 3000"
echo "🌐 Open: http://localhost:3000"
