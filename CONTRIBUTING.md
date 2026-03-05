# Contributing to go-superstruct-port
## Introduction
The go-superstruct-port project welcomes contributions from the community. This guide provides an overview of the contribution process, including code style, testing requirements, and the pull request process.
## Code Style
* The project uses the standard Go formatting rules. To ensure consistency, please run `go fmt` before committing your code changes.
* The project uses Go modules for dependency management.
## Testing Requirements
* All code changes must include corresponding test cases to ensure the functionality works as expected.
* Test cases should cover success, error, and edge cases (e.g., nil, empty, boundary values).
* Use table-driven tests with `t.Run()` for better organization and readability.
## Pull Request Process
1. Fork the repository and create a new branch for your changes.
2. Run `go test` and `go vet` to ensure your changes pass the existing test suite and do not introduce any new issues.
3. Run `golangci-lint run` to check for any code style or best practice issues.
4. Create a pull request against the main branch, including a clear description of the changes and any relevant testing or benchmarking results.
5. Address any feedback or issues raised by the maintainers during the review process.
## Commit Messages
* Follow the standard Go commit message format: `<type>(<scope>): <subject>`
* Use the following types:
  + `fix` for bug fixes
  + `feat` for new features
  + `docs` for documentation changes
  + `style` for code style changes
  + `refactor` for code refactoring
  + `perf` for performance improvements
  + `test` for test additions or changes
  + `chore` for miscellaneous changes
## Code of Conduct
The go-superstruct-port project follows the standard Go code of conduct. By contributing to this project, you agree to abide by the terms outlined in the [Go Code of Conduct](https://go.dev/conduct).
