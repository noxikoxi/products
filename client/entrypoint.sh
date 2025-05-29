#!/bin/sh

# Generuj plik config.js na podstawie zmiennych środowiskowych

mkdir -p /usr/share/nginx/html
cat <<EOF > /usr/share/nginx/html/config.js
window._env_ = {
  VITE_API_URL: "${VITE_API_URL:-http://localhost:1323}",
};
EOF

# Wypisz zawartość pliku dla celów debugowania
echo "Generated config.js:"
cat /usr/share/nginx/html/config.js
