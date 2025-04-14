# Password Hashing with bcrypt

This project demonstrates secure password hashing using the bcrypt algorithm in Go.

## Overview

Hashing is a cryptographic technique that converts any form of data into a fixed-size string of characters. When used for passwords:

- It creates a unique, fixed-length representation of a password
- It's a one-way function (cannot be reversed to obtain the original password)
- It's essential for secure password storage in databases

## Why bcrypt?

bcrypt is specifically designed for password hashing and offers several advantages:

1. **Salt Integration**: Automatically generates and incorporates random salt to protect against rainbow table attacks
2. **Adjustable Work Factor**: Allows setting the computational cost to slow down brute force attacks
3. **Future-Proofing**: Can be adjusted as hardware becomes more powerful

## Dependencies

-[bcrypt](golang.org/x/crypto/bcrypt)

## Usage

This implementation provides two main functions:

### Hashing a Password

```go
hash, err := hashSecret("yourpassword")
```

The hashSecret function:

- Takes a plaintext password string
- Returns a hashed string and any potential errors
- Uses a cost factor of 14 (higher values increase security but require more processing time)

### Verifying a Password

```go
isCorrect := isSecretHashCorrect("attemptedPassword", storedHash)
```

The `isSecretHashCorrect` function:

- Takes a plaintext password and a previously generated hash
- Returns a boolean indicating whether the password matches the hash

## Security Considerations
- Never store plaintext passwords
- Never transmit plaintext passwords over networks
- The cost factor (14 in this example) should be adjusted based on your security requirements and hardware capabilities
- Hashed passwords should still be treated as sensitive data