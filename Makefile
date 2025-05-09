# Build and run dengan docker
up:
	docker-compose up --build

# Hentikan services
down:
	docker-compose down

# Jalankan backend secara lokal (tanpa docker)
run-backend:
	cd src/backend && go run cmd/api/main.go

# Install dependencies frontend
install-frontend:
	cd src/frontend && npm install

# Jalankan frontend secara lokal
run-frontend:
	cd src/frontend && npm start

# Jalankan scraper
scrape-data:
	cd src/backend && go run cmd/scraper/main.go

# Print struktur direktori
directory-tree:
	tree -F --prune -I node_modules .