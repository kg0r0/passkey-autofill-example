# passkey-autofill-example
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)  
🔑 Experimental implementation of Passkey Autofill in Go.  
This simple implementation provides for developers to try out [Passkey Autofill](https://passkeys.dev/docs/reference/terms/#autofill-ui) in their environment. Note that this implementation uses the following two libraries:

- https://github.com/go-webauthn/webauthn
- https://github.com/MasterKale/SimpleWebAuthn

The endpoint design of this implementation is based on [[Web Authentication API Flow](https://www.w3.org/TR/webauthn-3/#sctn-api)] and [[Server Requirements and Transport Binding Profile](https://fidoalliance.org/specs/fido-v2.0-rd-20180702/fido-server-v2.0-rd-20180702.html)].
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
