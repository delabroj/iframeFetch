Demonstration: Using an iframe to access a cross-domain authenticated api.

## Api

The api is served from `localhost:8080`.

In order to receive the message from the `localhost:8080/securemessage` endpoint, the user must be logged in (a 'session' cookie must be set).

#### Login / Logout

Press the 'login' link on either of the apps below to set to set the session cookie and 'logout' to expire it.

## Authorized app

The app at http://localhost:8081 is whitelisted to make api requests via the iframe. If the api session cookie is set, it will receive the secure message.

## Unauthorized app

The identical app at http://localhost:8082 is not whitelisted. It will only receive an error indicating that the app is not authenticated.

## Try locally

`go run *`
