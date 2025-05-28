# Go Crypto Demo

This repository demonstrates the use of various cryptographic functions from `golang.org/x/crypto` package and includes Renovate configuration for automated security updates.

## Features

- Password hashing using multiple methods:
  - Argon2id (winner of the 2015 Password Hashing Competition)
  - Bcrypt (industry standard)
- Secure salt generation
- SSH key pair handling
- Password verification
- Automated security updates via Renovate

## Prerequisites

- Go 1.21 or later
- Renovate bot access to your repository

## Running the Program

```bash
go run main.go
```

The program will demonstrate:
- Password hashing with Argon2 and Bcrypt
- Password verification
- SSH key generation (if SSH_PRIVATE_KEY is set)

## Renovate Configuration

The repository includes a Renovate configuration (`renovate.json`) that:

- Monitors Go module dependencies for security vulnerabilities
- Creates PRs only for security-related updates
- Uses CODEOWNERS for automatic assignment
- Applies security and dependency labels
- Includes clear commit messages with version information
- Runs `go mod tidy` after updates

### Configuration Details

```json
{
  "extends": ["config:base"],
  "packageRules": [
    {
      "matchManagers": ["gomod"],
      "labels": ["area/security", "dependencies"],
      "automerge": false,
      "commitMessageTopic": "{{depName}}",
      "commitMessageExtra": "to {{newVersion}}"
    }
  ],
  "assignAutomerge": true,
  "assigneesFromCodeOwners": true,
  "postUpdateOptions": ["gomodTidy"],
  "baseBranches": ["main"]
}
```

## Security

This demo uses secure cryptographic practices:
- Argon2id with recommended parameters
- Bcrypt with default cost
- Cryptographically secure random number generation
- Proper salt generation and storage
- Secure password verification

## Dependencies

- `golang.org/x/crypto` v0.16.0
  - Used for bcrypt, argon2, and ssh functionality
  - Includes security fixes and improvements 