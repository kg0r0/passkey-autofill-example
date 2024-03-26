# passkey-autofill-example
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)  
🔑 Experimental implementation of Passkey Autofill in Go.  
This simple implementation provides for developers to try out [Passkey Autofill](https://passkeys.dev/docs/reference/terms/#autofill-ui) in their environment. Note that this implementation uses the following two libraries:

- https://github.com/go-webauthn/webauthn
- https://github.com/MasterKale/SimpleWebAuthn

The endpoint design of this implementation is based on [[Web Authentication API Flow](https://www.w3.org/TR/webauthn-3/#sctn-api)] and [[Server Requirements and Transport Binding Profile](https://fidoalliance.org/specs/fido-v2.0-rd-20180702/fido-server-v2.0-rd-20180702.html)].
If you want to check various behaviors according to your requirements, you should read the specifications on which Passkey is based, such as FIDO2 (Web Authentication + CTAP2).

## Usage

Run the server with the following command:

```
$ go run .
```

Access to http://localhost:8080 and register a passkey.

![](https://github.com/kg0r0/passkey-autofill-example/blob/assets/registration.png?raw=true)

Access to http://localhost:8080/login and authenticate.

![](https://github.com/kg0r0/passkey-autofill-example/blob/assets/authentication.png?raw=true)

## Registration

Basically you can see the implementation for registration from the following library documentation description:

- https://pkg.go.dev/github.com/go-webauthn/webauthn#readme-registering-an-account
- https://simplewebauthn.dev/docs/packages/browser#startregistration

Note that you will need to adjust some arguments based on the descriptions in [passkey.dev](https://passkeys.dev/docs/use-cases/bootstrapping/#opting-the-user-into-passkeys) and others. Also make sure that there are [parameters](https://www.w3.org/TR/webauthn-3/#sctn-authenticator-data) that should be properly verified on the server side.

The registration-related implementations in this repository can be seen in [attestation.go](https://github.com/kg0r0/passkey-autofill-example/blob/main/attestation.go) and [templates/index.html](https://github.com/kg0r0/passkey-autofill-example/blob/main/templates/index.html).


## Authentication

Basically you can see the implementation for authentication from the following library documentation description:

- https://pkg.go.dev/github.com/go-webauthn/webauthn#readme-logging-into-an-account
- https://simplewebauthn.dev/docs/packages/browser#browser-autofill-aka-conditional-ui

Note that you will need to adjust some arguments based on the descriptions in [passkey.dev](https://passkeys.dev/docs/use-cases/bootstrapping/) and others. Also make sure that there are [parameters](https://www.w3.org/TR/webauthn-3/#sctn-authenticator-data) that should be properly verified on the server side.

The complete authentication-related implementations in this repository can be seen in [assertion.go](https://github.com/kg0r0/passkey-autofill-example/blob/main/assertion.go) and [templates/login.html](https://github.com/kg0r0/passkey-autofill-example/blob/main/templates/login.html).