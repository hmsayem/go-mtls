## About
Implementation of mTLS (mutual TLS) using Golang and OpenSSL

### TLS
TLS (Transport Layer Security) provides the necessary encryption for applications when communicating over a network. HTTPS (Hypertext Transfer Protocol Secure) is an extension of HTTP that leverages TLS for security. The TLS technique requires a CA (Certificate Authority) to issue a X.509 digital certificate to a service, which is then handed over to the consumer of the service for it to validate it with the CA itself. mTLS extends the same idea to applications, for example, microservices wherein both the provider and the consumer require to produce their own certificates to the other party. These certificates are validated by both parties with their respective CAs. Once validated, the communication between the server/client or provider/consumer happens securely.

##### Reference: 
[A step by step guide to mTLS in Go](https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go).