# Go Gin + SvelteKit - JWT Authentication Example using Cookies

An example of how to use JWT authentication with Go Gin and SvelteKit using cookies. The server returns the jwt token in response body and inside SvelteKit we store it in a cookie. The cookie is sent with every request to the server(SvelteKit) when using SSR.

## Setup

### Start the server

```bash
make server
```

### Start the client

```bash
make frontend
```