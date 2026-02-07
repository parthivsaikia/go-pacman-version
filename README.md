# go-pacman-version

A lightweight, pure Go library for parsing and comparing Arch Linux (Pacman) package versions.

This library implements the exact version comparison logic used by `libalpm` (the core library of Pacman), ensuring that your Go applications resolve version ordering exactly as Pacman does.

## Features

- **Pure Go:** No CGO or external dependencies.
- **Spec Compliant:** Implements the `rpmvercmp` algorithm used by Arch Linux.
- **Robust:** Handles complex version strings including Epochs (`1:1.0`), versions (`1.0`), and releases (`-1`).
- **Permissive:** Like Pacman, it handles malformed strings gracefully without panicking.

## Installation

```bash
go get [github.com/parthivsaikia/go-pacman-version](https://github.com/parthivsaikia/go-pacman-version)
```

## Usage

```go
package main

import (
    "fmt"
    pacman "[github.com/yourusername/go-pacman-version](https://github.com/yourusername/go-pacman-version)"
)

func main() {
    v1 := "1.0.0-1"
    v2 := "1.0.0-2"

    // Check if a version string is valid (not empty)
    if pacman.IsValid(v1) {
        fmt.Println("Version is valid")
    }

    // Compare versions
    if pacman.LessThan(v1, v2) {
        fmt.Printf("%s is older than %s\n", v1, v2)
    }

    // Direct comparison (returns -1, 0, or 1)
    // -1 = v1 < v2
    //  0 = v1 == v2
    //  1 = v1 > v2
    cmp := pacman.Compare("1:2.0-1", "2.0-1")
    fmt.Println(cmp) // Output: 1 (because Epoch 1 > Epoch 0)
}
```

## Version Format

The library expects strings in the standard Arch Linux format:

```
[epoch:]version[-release]
```
