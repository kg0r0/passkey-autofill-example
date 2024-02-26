# passkey-autofill-example

🔑 Experimental implementation of Passkey Autofill in Go.

## Usage

Run the server with the following command:

```
$ go run .
```

Acces http://localhost:8080 and register a passkey.

![](/assets/registration.png)

Access http://localhost:8080/login and authenticate.

![](/assets/authentication.png)

## References
- FIDO2 & Passkeys
  - https://www.w3.org/TR/webauthn-2/
  - https://passkeys.dev/docs/intro/what-are-passkeys/
  - https://fidoalliance.org/specs/fido-v2.0-rd-20180702/fido-server-v2.0-rd-20180702.html
- Backend
  - https://github.com/go-webauthn/webauthn
  - https://github.com/duo-labs/webauthn
- Frontend
  - https://simplewebauthn.dev/docs/
- Other example implementation
  - https://github.com/go-webauthn/example/tree/master
  - https://github.com/NHAS/webauthn-example