# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.27+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

### GitHub Codespaces

Open the repository in a Codespace to get Go 1.27.1 and the Go editor extension configured automatically. The Codespace creates the `.env` file and forwards port 8080 for the app.

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!
