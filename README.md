# Poly Contracts

**Poly Contracts** is a language-agnostic contract definition system that uses Protocol Buffers and Buf to generate strongly-typed data transfer objects (DTOs) for multiple programming languages. It eliminates duplication, ensures consistency across services, and accelerates development by providing a single source of truth for inter-service communication.

## Table of Contents

- [Why Poly Contracts?](#why-poly-contracts)
- [Supported Languages](#supported-languages)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Generating Code](#generating-code)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Why Poly Contracts?

Microservices often need to exchange structured data. Without a shared contract, each service must maintain its own DTOs, leading to:

- **Duplication** – identical definitions in multiple languages.
- **Inconsistency** – mismatched fields or types across services.
- **Slow onboarding** – new developers must learn multiple codebases.

Poly Contracts solves these problems by:

1. **Defining contracts once** in Protocol Buffers (`.proto` files).
2. **Generating code** for all supported languages with a single command.
3. **Enforcing type safety** and versioning across services.

## Supported Languages

| Language | Generated Package |
|----------|-------------------|
| Go       | `gen/go`          |
| Python   | `gen/python`      |
| Java     | `gen/java`        |
| PHP      | `gen/php`         |
