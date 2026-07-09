```markdown
# open-code-review Development Patterns

> Auto-generated skill from repository analysis

## Overview
This skill teaches the core development patterns and conventions used in the `open-code-review` Go codebase. It covers file naming, import/export styles, commit message conventions, and testing patterns. While no automated workflows were detected, this guide includes suggested commands and step-by-step instructions for common development tasks.

## Coding Conventions

### File Naming
- Use **snake_case** for all file names.
  - Example: `review_handler.go`, `utils_test.go`

### Import Style
- Use **relative imports** for local packages.
  - Example:
    ```go
    import "./helpers"
    ```

### Export Style
- Use **named exports** for functions, types, and variables.
  - Example:
    ```go
    // In review_handler.go
    package review

    func StartReview() {
        // ...
    }
    ```

### Commit Messages
- Use **conventional commits** with the `feat` prefix for new features.
  - Example:
    ```
    feat: add support for multi-file reviews
    ```

## Workflows

### Creating a New Feature
**Trigger:** When adding a new feature to the codebase  
**Command:** `/new-feature`

1. Create a new file using snake_case naming.
2. Implement your feature using named exports.
3. Use relative imports for any local dependencies.
4. Write or update corresponding test files (`*_test.go`).
5. Commit your changes using the `feat` prefix in the commit message.
    - Example: `feat: implement batch comment posting`
6. Push your branch and open a pull request.

### Running Tests
**Trigger:** When verifying code correctness  
**Command:** `/run-tests`

1. Locate or create test files matching the `*_test.go` pattern.
2. Run tests using the Go testing tool:
    ```sh
    go test ./...
    ```
3. Review test output and fix any failing tests.

## Testing Patterns

- Test files use the `*_test.go` naming convention.
- The testing framework is not explicitly specified; use Go's built-in testing.
- Example test file:
    ```go
    // review_handler_test.go
    package review

    import "testing"

    func TestStartReview(t *testing.T) {
        // test logic here
    }
    ```

## Commands
| Command        | Purpose                                    |
|----------------|--------------------------------------------|
| /new-feature   | Start the process for adding a new feature |
| /run-tests     | Run all tests in the codebase              |
```
