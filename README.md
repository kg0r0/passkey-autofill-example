# related-origins-example
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)  
🔑 Experimental implementation of Related Origin Requests in Go.  
This simple implementation provides for developers to try out [Related Origin Requests](https://passkeys.dev/docs/advanced/related-origins/) in their environment. This reproduces the scenario described in [passkeys.dev](https://passkeys.dev/docs/advanced/related-origins/#example). In this scenario, a user who created a passkey in the ``UK (https://shopping.co.uk)`` travels to the ``US (https://shopping.com)``.  
Note that this implementation uses the following two libraries:

- https://github.com/go-webauthn/webauthn
- https://github.com/MasterKale/SimpleWebAuthn

The endpoint design of this implementation is based on [[Web Authentication API Flow](https://www.w3.org/TR/webauthn-3/#sctn-api)] and [[Server Requirements and Transport Binding Profile](https://fidoalliance.org/specs/fido-v2.0-rd-20180702/fido-server-v2.0-rd-20180702.html)].
If you want to check various behaviors according to your requirements, you should read the specifications on which Passkey is based, such as FIDO2 (Web Authentication + CTAP2).

## Usage

To test this implementation, you need to make several preparations. First, add the following definitions to ``/etc/hosts`` to allow access to the local RP on shopping.com and shopping.co.uk.

```bash
127.0.0.1	localhost shopping.com shopping.co.uk
```

Additionally, if you are using an origin other than localhost when calling the Web Authentication API, you need to ensure it is HTTPS and that no certificate errors occur. You can start an HTTPS server locally without certificate errors by following the steps described in ["Use HTTPS for local development"](https://web.dev/articles/how-to-use-local-https). For example, in this case, executing the following command will generate ``shopping.com+1-key.pem`` and ``shopping.com+1.pem``, allowing access to ``shopping.com`` and ``shopping.co.uk`` without certificate errors.

```bash
$ mkcert  shopping.com  shopping.co.uk
```

Run the server with the following command:

```bash
$ git clone -b reloated-origins https://github.com/kg0r0/passkey-autofill-example.git
$ cd passkey-autofill-example
$ cp /path-to-your-key-file/shopping.com+1-key.pem certs/shopping.com+1-key.pem
$ cp /path-to-your-cert-file/shopping.com+1.pem certs/shopping.com+1.pem
$ ls certs
shopping.com+1-key.pem shopping.com+1.pem
$  go run .
```

Access to https://shopping.co.uk and register a passkey.

![](https://github.com/kg0r0/zenn-docs/blob/main/images/15f64a2dc54200/registration_uk.png?raw=true)

Access to https://shopping.co.uk/login and authenticate.

![](https://github.com/kg0r0/zenn-docs/blob/main/images/15f64a2dc54200/authentication_uk.png?raw=true)

Next, access to ``https://shopping.com``. At this time, specify ``shopping.co.uk`` as the RP ID and call the Web Authentication API.

![](https://github.com/kg0r0/zenn-docs/blob/main/images/15f64a2dc54200/authentication_com_1.png?raw=true)

It can be confirmed that a request was triggered to ``https://[RP ID]/.well-known/webauthn`` during the call of the Web Authentication API due to the mismatch between the accessing origin and the RP ID as follows.

```
$ go run .
2024/12/22 18:29:13 INFO shopping.co.uk/.well-known/webauthn
```

The sign-in is successful as well.

![](https://github.com/kg0r0/zenn-docs/blob/main/images/15f64a2dc54200/authentication_com_2.png?raw=true)

If ``https://shopping.com`` is removed from the allowed origins included in the response from ``https://shopping.co.uk/.well-known/webauthn``, the following error will occur, resulting in a failed sign-in.

```bash
SecurityError: The RP ID "shopping.co.uk" is invalid for this domain
    at error (index.umd.min.js:2:3961)
    at e.startAuthentication (index.umd.min.js:2:4289)
    at async login:47:22Caused by: SecurityError: The relying party ID is not a registrable domain suffix of, nor equal to the current domain. Subsequently, fetching the .well-known/webauthn resource of the claimed RP ID was successful, but no listed origin matched the caller.
```