@echo off

cd /d "%~dp0"

echo [*] Menjalankan API...
go run cmd/api/main.go

echo.
pause
