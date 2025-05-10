@echo off

cd /d "%~dp0"

echo [*] Menjalankan scraper...
go run cmd/scraper/main.go