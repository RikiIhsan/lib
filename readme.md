# asrs lib

[![Go Version](https://img.shields.io/badge/go-1.23-blue)](https://golang.org/)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)

## Introduction

This project is a Go application that provides password hashing and comparison using the Argon2 algorithm. It includes functionalities for generating secure hashes from passwords, comparing stored hashes with provided passwords, and handling validation errors in a web framework (Fiber).

## Features

- **Secure Password Hashing**: Utilizes Argon2id to securely hash passwords.
- **Password Comparison**: Compares provided passwords with stored hashes.
- **Validation Handling**: Handles validation errors using custom structs for easy error management.
- **Database Integration**: Includes support for GORM and Redis for database operations.

## Installation

To install the project, follow these steps:

1. get the module:
   ```sh
   go get github.com/RikiIhsan/lib.git
   ```

## Usage

### Comparing a Password with a Hash

To compare a provided password with a stored hash, use the following command:

```sh
match,err:=ComparePassAndHash(password, hash string)
```

This will output whether the passwords match or not.

### Handling Validation Errors

The project includes a validator package that can be used to handle validation errors in your web application. Here is an example of how to use it:

1. Define your validation rules and data structure.
2. Use the validator middleware to validate incoming requests.
3. Handle validation errors gracefully.

Example usage in Fiber:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    v "github.com/RikiIhsan/lib/validator"
)

func main() {
    app := fiber.New()

    // Define your validation rules and data structure
    type User struct {
        Name  string `json:"name" validate:"required"`
        Email string `json:"email" validate:"required,email"`
    }

    // Middleware to handle validation errors
    app.Use(func(c *fiber.Ctx) error {
        v := validator.New()
        if err := v.Validate(c, User{}); len(err) > 0 {
            return c.Status(400).JSON(err)
        }
        return c.Next()
    })

    app.Post("/user", func(c *fiber.Ctx) error {
        user := new(User)
        if err := c.BodyParser(user); err != nil {
            return err
        }
        // Your logic here
        return c.JSON(user)
    })

    app.Listen(":3000")
}
```

```

```
