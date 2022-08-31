# secure-todo-api

Fullstack TODO App with Angular as Frontend and Go with ECHO as Backend and MongoDB as database

## Run the application

1. Open a terminal in the root of the repository where the `docker-compose.yml` file is located
2. Run the docker compose command with the needed arguments for the database user and password:

```bash
DB_USER=<your-username> DB_PASSWORD=<your-password> docker-compose up -d --build
```

3. Navigate to [http://localhost:3000](http://localhost:3000)
4. Register a user, after that you are redirected to the TODOs view
5. Add Todo's to your liking! 😋
