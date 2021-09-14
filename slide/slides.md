---
theme: default
layout: cover
title: grpc-demo
---

# gRPC

General Overview

<div class="text-sm">
  Made By:
  <ul>
    <li>Steve Bezalel Iman Gustaman 13518018</li>
    <li>Matthew Kevin Amadeus 13518035</li>
    <li>Hengky Surya Angkasa 13518048</li>
    <li>Yonatan Viody 13518120</li>
  </ul>
</div>

<div class="abs-br m-6 flex gap-2">
  <a href="https://github.com/mkamadeus/grpc-demo" target="_blank" alt="GitHub"
    class="text-xl icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

<!--
The last comment block of each slide will be treated as slide notes. It will be visible and editable in Presenter Mode along with the slide. [Read more in the docs](https://sli.dev/guide/syntax.html#notes)
-->

---
layout: cover
title: gRPC
---

# What is gRPC?
General overview of gRPC.

---
layout: quote
---
<div class="italic text-3xl opacity-70">
  "gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment."
</div>

---

# What is gRPC?
Taken from https://grpc.io/

Developed by Google in 2015, here's what they aim for in gRPC:

- 💡 **Simple service definition** - made with Protobuf (https://developers.google.com/protocol-buffers).
- ⚡ **Start quickly and scale** - easy to scale.
- 👨‍💻 **Works across languages and platforms** - generate code for multiple languages.
- ✌ **Bi-directional streaming and integrated auth** - and it uses HTTP/2.0 under the hood.

---

# Where is it used?
Scenarios where gRPC is used, taken from https://grpc.io/

- Interservice communication in a microservices architecture
- Connecting mobile devices and IoT to services
- Generating efficient client libraries

Companies that adopts them:

<div class="flex text-white text-5xl gap-4 items-center mb-4">
  <cib-netflix />
  <cib-cisco />
  <ri-square />
  <ri-coreos-fill />
  <!-- <img class="w-24 h-24" src="https://simpleicons.org/icons/square.svg" /> -->
  <div class="italic opacity-50 text-sm">...even Warung Pintar, Xendit, Tokopedia, Gojek</div>
</div>
<div class="italic opacity-50 text-sm">...and many more!</div>


---
layout: cover
title: gRPC
---

# How does gRPC Work?
Using HTTP 2.0 under the hood, let's see how it works.

---

# How does gRPC Work?

Taken from https://grpc.io/docs/what-is-grpc/introduction/

<div class="flex">
  <div class="w-full">
    <ul>
      <li>Basically, it's sending blob via HTTP/2 for requests and responses</li>
      <li>Similar to other RPCs, it can generate stub for each languages</li>
      <li>Marshaling and unmarshaling happens inside the stub</li>
    </ul>
  </div>
  <div class="w-full"><img src="https://grpc.io/img/landing-2.svg" /></div>
</div>


---
layout: cover
title: gRPC
---

# What's Unique About gRPC?
Some notable things that make it unique from others.

---
layout: cover
title: gRPC
---

# Pros and Cons of gRPC
Things to consider before adopting gRPC.

---

# Pros and Cons of gRPC

Taken from https://www.capitalone.com/tech/software-engineering/grpc-framework-for-microservices-communication/

## Pros
- ☝️ Reduced network latency <br>
- ✌️ Duplex streaming <br>
- 👌 Code generation
## Cons
- ☝️ Lack of consistent error handling <br>
- ✌️ Lack of support for additional content types <br>
- 👌 Most browsers don't support gRPC
---
layout: cover
title: gRPC
---

# Demo Time!
First, let's see how it's used in practice.

<div class="abs-br m-6 flex gap-2">
  <a href="https://github.com/mkamadeus/grpc-demo" target="_blank" alt="GitHub"
    class="flex items-center gap-4 text-xl icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
    <span class="text-xs italic">
    Check out our GitHub Demo.
    </span>
  </a>
</div>

---

# Some steps to follow
Quite similar with other RPCs.

1. **Make your service definition** - achieved by using Protobuf to define your methods and service.
2. **Generate the stub** - using `protoc` and some plugins, generate the code for the required language.
3. **Implement the stub** - implement the methods generated from the stub.

---

# Protobuf Definition Example
Short sample, taken from https://github.com/grpc/grpc-go/blob/master/examples/

```go {all|1|3-8|10-12|14-20}
syntax = "proto3";

option go_package = "google.golang.org/grpc/examples/helloworld/helloworld";
option java_multiple_files = true;
option java_package = "io.grpc.examples.helloworld";
option java_outer_classname = "HelloWorldProto";

package helloworld;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

---

# Generate Stub from Protobuf
Generating Go code, taken from https://github.com/grpc/grpc-go/blob/master/examples/

This command requires `protoc` and `protoc-gen-go` to be installed.
Refer to the documentation for other languages.

```bash {all|2-3|4-5|6}
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  helloworld/helloworld.proto
```

---
layout: end
---
