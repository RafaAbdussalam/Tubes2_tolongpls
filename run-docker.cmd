@echo off
start "" "http://localhost:3000"
start cmd /k "docker-compose up --build"
