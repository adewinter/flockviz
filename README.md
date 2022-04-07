# flockviz

# setup

Clone this repo

## Go server

In a new terminal, navigate to this repo's root folder.
Run go tidy
You should now be able to run the server go program:

```
go run server/server.go
```

Keep this terminal open and running, launch a new terminal and navigate to the `<root>/webclient` folder.

## In `webclient/` folder

Run:

```
npm install
npm run start
```

This will launch the a server at http://localhost:3000. Open your browser to that URL (or whatever URL it tells you the server is running at in your console)

Keep this running.

# Usage

In your browser session that should now be open and pointing to something like http://localhost:3000, click the "go" button. You should see a lot of new numbers be written to the page. This means the browser successfully made a gRPC stream request to the go server backend. The data was streamed to the browser JS client, "parsed" and rendered as React components.

# Misc

There is a Go client in the `testclient/` folder to verify that the GRPC streaming works.
