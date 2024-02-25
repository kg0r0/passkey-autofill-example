# fido2-go-example

Experimental implementation of Passkey Autofill in Go.

## Usage

Run the server with the following command:

```
$ go run .
```

Acces http://localhost:8080 and register a passkey.

<img width="1054" alt="Screenshot 2024-02-25 at 22 59 45" src="https://github.com/kg0r0/fido2-go-example/assets/33596117/104678f4-f34f-405e-9b86-f772532c85e4">

Access http://localhost:8080/login and authenticate.

<img width="532" alt="Screenshot 2024-02-25 at 23 00 04" src="https://github.com/kg0r0/fido2-go-example/assets/33596117/30e251a2-21c4-4f77-a273-2f740128addd">

## References
- https://github.com/go-webauthn/webauthn
- https://github.com/duo-labs/webauthn