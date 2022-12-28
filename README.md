# golang-grpc
This repo illustrates how to create unary, server streaming, client streaming and bi-directional APIs using gRPC

# gRPC
Free and open source framework developed by Google
Part of the CNCF (Cloud Native Computation Foundation)
At a high level, it allows you to define REQUEST and RESPONSE for Remote Procedure Calls(RPCs) and handles the rest for you.
Built on top of HTTP/2
Data format → protocol buffers (proto requests and proto responses)

# Types of APIs in gRPC
Unary
Classic request-response type
Basically the REST style, but you get all the advantages of protocol buffers and  HTTP/2
The server and the client can push multiple messages as part of one request
Unary calls are defined using Protocol buffers
For each RPC call, we have to define a “Request” message and a “Response” message 
Server streaming
Enabled through HTTP/2
Client makes one request and receives a stream of responses
Client streaming
Client sends a stream of requests to the server and expects a response at specific points
Bi directional streaming
Asynchronous sending of messages between the client and the server

# gRPC vs REST 

| gRPC 	     					        | REST        
| ----------------------------------------------------- | ------------------------------------------------------------ |
| Protocol buffers - smaller, faster, binary            | JSON - text based, slow, bigger                              |
| HTTP/2 - low latency  		                | HTTP 1.1 - high latency   				       |
| Bidirectional and Async                               | Client-server requests only                                  |
| Stream support                                        | Request - response support only                              |
| API oriented (no constraints - free design)           | CRUD Oriented - Resource oriented                            |
| RPC based → call functions directly on the server from the client | HTTP verb based                                  |
| Code generation through protocol buffers in any language | Code generation through OpenAPI / Swagger                 |


# Summary: why use gRPC
Easy code generation in over 11 languages
Uses a HTTP/2 → low latency, fast, secure, multiplexed
SSL security is built in
Support for streaming APIs for maximum performance


