FRONTEND_DIR = ./src/frontend
BACKEND_DIR = ./src/backend

# RUN DEV MODE
run-scraper:
	cd $(BACKEND_DIR) && go run cmd/scraper/main.go
run-backend:
	cd $(BACKEND_DIR) && go run cmd/api/main.go
run-frontend:
	cd $(FRONTEND_DIR) && npm start

# DOCKER
up:
	docker-compose up --build
down:
	docker-compose down

# ADDITIONAL
directory-tree:
	tree -F --prune -I node_modules .