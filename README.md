# 🚀 grpc-demo
> A small demo for gRPC for IF4031 Distributed Application Development

## 🚀 What?

As the subtitle says, it is a small demo made to demonstrate some key features of gRPC.
This repository includes:
- Go-based gRPC server
- Some basic Protobuf schema
- Several gRPC clients (implemented)

## 🚀 Pitch Deck

The slides can be seen in the slide folder by using `yarn dev`.
Preferably, see the live version on https://grpc-demo.mkamadeus.dev.

## 🚀 How?

Here's a quick setup guide to run this demo.

First, generate the stubs for the gRPC server and the client:
```bash
make proto
```
You can generate them individually per feature, but to test the server we need to compile all of them.

Next, run the server:
```bash
make server
```

Finally, you can test with your client of choice. Just run the client and you're good to go.

Or, alternatively, you can use `evans`, a gRPC client (https://github.com/ktr0731/evans).
```
evans -r repl -p 1337
```