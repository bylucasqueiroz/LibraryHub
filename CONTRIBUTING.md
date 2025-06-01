# Contributing to Go Mono Repository

We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Development Process

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

1. Fork the repo and create your branch from `main`
2. If you've added code that should be tested, add tests
3. If you've changed APIs, update the documentation
4. Ensure the test suite passes and coverage is at least 90%
5. Make sure your code lints
6. Issue that pull request!

## Pull Request Process

1. Update the README.md with details of changes to the interface, if applicable
2. Update any relevant documentation
3. The PR may be merged once you have the sign-off of the maintainers and all tests pass

## Code Quality Requirements

- All new code must have associated tests
- Test coverage must remain at 90% or higher
- All code must pass `go fmt` and `go vet`
- Follow Go best practices and style guidelines

## Testing

Before submitting a PR, ensure all tests pass:

```bash
go test ./...
go test -cover ./...
```

## Commit Messages

We follow conventional commits specification:

- `feat:` - A new feature
- `fix:` - A bug fix
- `docs:` - Documentation only changes
- `style:` - Changes that do not affect the meaning of the code
- `refactor:` - A code change that neither fixes a bug nor adds a feature
- `test:` - Adding missing tests or correcting existing tests
- `chore:` - Changes to the build process or auxiliary tools

Example:
```
feat: add metadata support to context package
```

## License

By contributing, you agree that your contributions will be licensed under its MIT License. 