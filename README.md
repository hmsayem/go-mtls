## About
Implementation of mTLS (mutual TLS) using Golang and OpenSSL

### TLS
TLS (Transport Layer Security) provides the necessary encryption for applications when communicating over a network. HTTPS (Hypertext Transfer Protocol Secure) is an extension of HTTP that leverages TLS for security. The TLS technique requires a CA (Certificate Authority) to issue a X.509 digital certificate to a service, which is then handed over to the consumer of the service for it to validate it with the CA itself. mTLS extends the same idea to applications, for example, microservices wherein both the provider and the consumer require to produce their own certificates to the other party. These certificates are validated by both parties with their respective CAs. Once validated, the communication between the server/client or provider/consumer happens securely.
#### Certificate Generation with OpenSSL
```
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out cert.pem  -keyout key.pem -subj "/C=US/ST=Nevada/L=LasVegas/O=AppsCode/OU=Backend/CN=Server" -addext "subjectAltName = DNS:localhost"

```
### How to run
Copy all third-party dependencies to a vendor folder in your project root.
```
go mod vendor
```
Open an instance of the terminal and run the server.
```
go run server.go
```
Open another instance of the terminal and run the client.
```
go run server.go
```
You should see the following output from the client.
```
Hello world!
```
##### Reference: 
[A step by step guide to mTLS in Go](https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go).