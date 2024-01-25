# Go Gin + SvelteKit - JWT Authentication Example using Cookies

An example of how to use JWT authentication with Go Gin and SvelteKit using cookies. The server returns the jwt token in response body and inside SvelteKit we store it in a cookie. The cookie is sent with every request to the server(SvelteKit) when using SSR.

## Setup

### Install dependencies

```bash
# Installing the node dependencies for the frontend
make install_frontend

# Installing the go dependencies for the server
make install_server
```

### Start the server

```bash
make server
```

### Start the client

```bash
make frontend
```

## Routes

| Route           | Method | Request Body         | Description                                                                   |
| --------------- | ------ | -------------------- | ----------------------------------------------------------------------------- |
| /auth/login     | `POST` | `email` + `password` | Login route, returns a jwt `token` + `message`                                |
| /auth/signup    | `POST` | `email` + `password` | Signup route, returns a jwt `token` + `message`                               |
| /user/get-my-id | `GET`  | `none`               | Returns the `id` of the user, needs `Authorization` header with the jwt token |
