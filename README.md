# Poly Contracts

Language independent contracts using Protocol Buffers.

The repository provides a single source of truth for communication between microservices.

Supported languages:

- Go
- Python
- Java
- PHP


## Why

Sharing DTOs between different languages creates duplication.

This project solves that problem using:

- Protocol Buffers
- Buf
- Code Generation



## Structure

.
├── buf.gen.yaml
├── buf.yaml
├── examples
 └── python
      └── main.py
├── gen
├── LICENSE
├── Makefile
├── proto
  ├── common
  │  └── v1
  │     └── common.proto
  └── wallet
│     └── v1
│         └── wallet.proto
├── README.md
└── scripts
    └── generate.sh
