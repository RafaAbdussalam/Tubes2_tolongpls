@REM Menjalankan Frontend (dev mode)
@echo off
start cmd /c "cd ./src/frontend && npm start && exit"

@REM Menjalankan Backend - Scraper
@echo off
start cmd /c "cd ./src/backend && CALL ./run-scraper.cmd && exit"

@REM Menjalankan Backend - API
@echo off
start cmd /c "cd ./src/backend && CALL ./run-api.cmd && exit"