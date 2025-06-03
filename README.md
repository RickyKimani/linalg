# 🔢 Linalg

[![Linear Algebra Library](https://img.shields.io/badge/linalg-Library-blue)](https://pkg.go.dev/github.com/rickykimani/linalg)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/licenses/MIT)
[![CI Status](https://github.com/rickykimani/linalg/workflows/CI/badge.svg)](https://github.com/rickykimani/linalg/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickykimani/linalg)](https://goreportcard.com/report/github.com/rickykimani/linalg)

---

## 📋 Overview

**Linalg** is a comprehensive linear algebra library for the Go programming language. It provides a robust set of tools for matrix and vector operations, designed with an emphasis on performance, ease of use, and idiomatic Go practices. Currently, `linalg` focuses on foundational matrix and vector functionalities, with plans to expand into more advanced numerical computations.

### Why Linalg?

* **Pure Go Implementation**: Crafted in 100% Go, `linalg` has no external dependencies (CGO-free), ensuring simple builds and excellent portability across all Go-supported platforms.
* **Type Safety**: Leverages Go's strong type system to minimize runtime errors and enhance developer confidence.
* **Memory Efficiency**: Designed with memory-conscious patterns, particularly beneficial for operations on large matrices and vectors.
* **Developer-Friendly API**: Offers an intuitive and easy-to-understand API, making complex linear algebra operations accessible.
* **Modular Design**: Allows users to import and utilize only the necessary components, keeping applications lean.

### Core Components

* **Matrix Operations**: Functionality for creating, manipulating, and decomposing matrices.
* **Vector Operations**: A suite of tools for vector arithmetic, analysis, and transformations.

### Project Status

`linalg` is under active development. While the core functionalities are stabilizing, the API is subject to change as we introduce new features and optimizations. Community feedback and contributions are highly encouraged.

---

## ✨ Features

`linalg` currently supports a rich set of operations for matrices and vectors:

### 📊 Matrix Functions

* **Arithmetic Operations**:
  * Addition & Subtraction
  * Multiplication (Matrix-Matrix & Scalar)
  * Powers of Matrices
* **Decomposition**:
  * QR Decomposition
  * LU Decomposition
* **Properties & Transformations**:
  * Rank
  * Trace
  * Determinant
  * Transpose
  * Inverse
  * Eigenvalues
* **Utilities**:
  * Identity Matrix Generator

### 📐 Vector Functions

* **Basic Arithmetic**:
  * Addition & Subtraction
* **Geometric Operations**:
  * Angle Between Vectors
  * Cross Product
  * Dot Product
* **Distance Metrics**:
  * Euclidean Distance
  * Manhattan Distance
  * Chebyshev Distance
* **Properties & Transformations**:
  * Magnitude (Norm)
  * Normalization
  * Scalar Projection & Vector Projection
  * Scaling (Scalar Multiplication)
  * Transformations:
    * Rotate 2D
    * Rotate 3D
    * Reflection
* **Direction & Coordinates**:
  * Direction Cosines
  * Coordinate System Conversions (Cartesian, Polar, Cylindrical, Spherical)

---

## 🔮 Future Enhancements

We are continuously working to expand the capabilities of `linalg`. Key areas for future development include:

* **Solvers for Linear Equations**:
  * Gaussian elimination
  * Jacobi method
  * Gauss-Seidel method
  * And more advanced iterative solvers.
* **Expanded Matrix Decompositions**:
  * Singular Value Decomposition (SVD)
  * Cholesky Decomposition
* **Performance Optimizations**: Further profiling and optimization for critical computation paths.

---

## 🚀 Installation

To start using `linalg` in your Go project, you can install the necessary packages using `go get`:

```bash

go get github.com/rickykimani/linalg/matrix

```

```bash

go get github.com/rickykimani/linalg/vectors

```

---

## 🔍 Usage Example

``` go
package main
import (
  "fmt"

  "github.com/rickykimani/linalg"
)

func main(){
  data := [][]int {
  {2, 4, 7},
  {3, 8, 2},
  {5, 9, 6},
  }

  mat, err := matrix.NewMatrix(data)
  
  if err != nil{
    panic(err)
  }

  transpose := matrix.Transpose(mat)
  det, err := matrix.Det(mat)
  
  if err != nil{
    panic(err)
  }

  fmt.Printf("Transpose:\n %v", transpose)
  fmt.Printf("Determinant: %.6f", det)
}

```

---

---

## 📖 Documentation

Comprehensive documentation for `linalg` is available on [GoDoc](https://pkg.go.dev/github.com/rickykimani/linalg). This includes detailed info on all functions and types.

---

## 🤝 Contributing

We welcome contributions to the `linalg` project! Here's how you can help:

### Getting Started

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Add YourFeature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a Pull Request

### Guidelines

* **Code Style**: Follow Go's official [style guide](https://golang.org/doc/effective_go)
* **Documentation**: Add comments to functions and update documentation as needed
* **Testing**: Include tests for new functionality with good coverage
* **Commit Messages**: Write clear, concise commit messages explaining the changes

### Reporting Issues

* Use the GitHub issue tracker to report bugs
* Include detailed steps to reproduce the issue
* Mention your environment (Go version, OS, etc.)

### Development Workflow

* Check existing issues and PRs before starting work
* For significant changes, open an issue for discussion first
* Ask for code review from maintainers when ready

We appreciate your contributions to making `linalg` better!

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository

2. Create a feature branch (`git checkout -b feature/AmazingFeature`)

3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)

4. Push to the branch (`git push origin feature/AmazingFeature`)

5. Open a Pull Request

---

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.
