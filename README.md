# Go-Saas TaskForge v2.0

Go-Saas TaskForge is a modern SaaS task manager.

## Tech Stack
- Backend: Golang, Gin, MongoDB
- Frontend: React, Vite, TypeScript, Tailwind CSS

## Architecture
- Golang Gin REST API with JWT auth and MongoDB
- React Vite frontend with Tailwind CSS
- Auth endpoints: POST /auth/signup, POST /auth/login
- Task endpoints: GET/POST/PUT/DELETE /tasks (protected)
- Team endpoints: GET/POST /teams (protected)
