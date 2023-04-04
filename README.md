# Music API server
## RESTful API server, written in Go with Gin framework

### Features:
- CRUD operations
- JWT token autorization
- MongoDB database

### Used technologies:
- Go
- Gin
- HTTP
- JSON
- JWT
- Docker
- MongoDB
- Clean Architecture
- Git
- Enviroment variables
- Config file
- Docker Compose

### Hot to run
Create .env file in root directory which contains login and password to MongoDB. According to env in docker-compose.yaml, for this project it should look like:
```
MONGODB_USERNAME = root
MONGODB_PASSWORD = funnycat
```

To run server, go to your directory, download required modules and run docker compose.
```BASH
cd /PATH/TO/Your/REPO/MusicApi
docker compose up --build
```

if you want to run MongoDB detached: `docker compose up server_container`

Then use your prefered HTTP Client. By default, URL is `http://localhost:8000/`

### Routers
 - `GET http://localhost:8000/api/album/:id` Get album data by id
 - `GET http://localhost:8000/api/album` Get all album data
 - `DELETE http://localhost:8000/api/album/:id` Delete albym by id
 - `POST http://localhost:8000/api/album` Add new album to database. Requires auth token 
 - `POST http://localhost:8000/api/album/array` Add array of albums to database. Requires auth token 
 - `POST http://localhost:8000/api/admin` Add new admin. Requires auth token
 - `POST http://localhost:8000/auth/register` Create new user
 - `POST http://localhost:8000/auth/login` Login to your account.
 
 **Examples of requests you can find in [requests.md](requests.md)**
