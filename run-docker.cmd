@echo off
start "" "http://localhost:3000"
start cmd /c "docker-compose up --build"
