# passkey-autofill-example

🔑 Experimental implementation of Passkey Autofill in Go.

This simple implementation provides for developers to try out Passkey Autofill in their environment. Note that this implementation uses the following two libraries:

- https://github.com/go-webauthn/webauthn
- https://github.com/MasterKale/SimpleWebAuthn

You can read additional explanations of this implementation in [this document](https://kg0r0.medium.com/experimental-implementation-of-passkey-autofill-in-go-b10c5c5d98b4).

## Usage

Run the server with the following command:

```
$ go run .
```

Access to http://localhost:8080 and register a passkey.

![](https://github.com/kg0r0/passkey-autofill-example/blob/assets/registration.png?raw=true)

Access to http://localhost:8080/login and authenticate.

![](https://github.com/kg0r0/passkey-autofill-example/blob/assets/authentication.png?raw=true)

## References
- FIDO2 & Passkeys
  - https://www.w3.org/TR/webauthn-3/
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