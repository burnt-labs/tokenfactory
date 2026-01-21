# Contributing to Token Factory

Thank you for your interest in contributing to the Token Factory project! This document provides guidelines and information to help you get started.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Running Tests](#running-tests)
- [Code Style and Linting](#code-style-and-linting)
- [Submitting Changes](#submitting-changes)
- [Reporting Issues](#reporting-issues)
- [Code of Conduct](#code-of-conduct)

## Getting Started

This repository is a fork of [strangelove-ventures/tokenfactory](https://github.com/strangelove-ventures/tokenfactory). Before contributing, please:

1. Check the [README.md](README.md) for project overview and setup instructions.
2. Review open issues in both this fork and the upstream repository.
3. Ensure your contribution aligns with the project's goals.

## Development Setup

### Prerequisites

- Go 1.21 or later
- Make
- Docker (for integration tests)
- jq (for test scripts)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/burnt-labs/tokenfactory.git
   cd tokenfactory
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the application:
   ```bash
   make install
   ```

4. Verify installation:
   ```bash
   tokend version
   ```

## Running Tests

### Unit Tests

Run all unit tests:
```bash
make test
```

### Integration Tests

Run end-to-end tests using Interchain Test:
```bash
make local-image
make ictest-tokenfactory
```

### Coverage

Generate test coverage report:
```bash
make local-image
make coverage
```

### Simulation Tests

Run simulation tests:
```bash
make sim-full-app
```

## Code Style and Linting

This project uses [golangci-lint](https://golangci-lint.run/) for code quality checks.

### Running Linter

```bash
make lint
```

### Fixing Lint Issues

```bash
make lint-fix
```

### Code Style Guidelines

- Follow standard Go conventions
- Use `gofmt` for formatting
- Keep functions and methods focused and concise
- Add comments for exported functions and complex logic
- Use meaningful variable and function names

## Submitting Changes

### Pull Request Process

1. Fork the repository and create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following the code style guidelines.

3. Run tests and linting:
   ```bash
   make test
   make lint
   ```

4. Commit your changes with a clear, descriptive message:
   ```bash
   git commit -m "feat: add new feature description"
   ```

5. Push to your fork and create a pull request:
   ```bash
   git push origin feature/your-feature-name
   ```

6. In your PR description:
   - Describe the problem you're solving
   - Explain your solution
   - Reference any related issues
   - Include screenshots or examples if applicable

### Commit Message Format

Use conventional commit format:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for test additions
- `refactor:` for code refactoring
- `chore:` for maintenance tasks

## Reporting Issues

### Bug Reports

When reporting bugs, please include:

- A clear title and description
- Steps to reproduce
- Expected vs. actual behavior
- Environment details (Go version, OS, etc.)
- Relevant logs or error messages

### Feature Requests

For feature requests:

- Describe the feature and its use case
- Explain why it's needed
- Provide examples if possible

## Code of Conduct

This project follows a code of conduct to ensure a welcoming environment for all contributors. By participating, you agree to:

- Be respectful and inclusive
- Focus on constructive feedback
- Accept responsibility for mistakes
- Show empathy towards other contributors

## Additional Resources

- [Cosmos SDK Documentation](https://docs.cosmos.network/)
- [Interchain Test Framework](https://github.com/strangelove-ventures/interchaintest)
- [Token Factory Module Documentation](https://docs.cosmos.network/main/modules/bank#token-factory)

Thank you for contributing to Token Factory!</content>
<parameter name="filePath">/home/thee1/tokenfactory/CONTRIBUTING.md