# Linalg

[![Linear Algebra Library](https://img.shields.io/badge/linalg-Library-blue)](https://pkg.go.dev/github.com/rickykimani/linalg)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/licenses/MIT)
[![CI Status](https://github.com/rickykimani/linalg/workflows/CI/badge.svg)](https://github.com/rickykimani/linalg/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickykimani/linalg)](https://goreportcard.com/report/github.com/rickykimani/linalg)

---

## Overview

**Linalg** is a comprehensive linear algebra library. It provides a set of tools for matrix and vector operations. Currently, `linalg` focuses on foundational matrix and vector functionalities, with plans to expand into more advanced numerical computations.


### Core Components

* **Matrix Operations**: Functionality for creating, manipulating, and decomposing matrices.
* **Vector Operations**: A suite of tools for vector arithmetic, analysis, and transformations.

### Project Status

`linalg` is under active development. While the core functionalities are stabilizing, the API is subject to change as new features and optimizations are introduced.

---

## Features

`linalg` currently supports this set of operations for matrices and vectors:

### Matrix Functions

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

### Vector Functions

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

* **Solvers for Linear Systems of Equations**:
  * Gaussian elimination
  * Jacobi method
  * Gauss-Seidel method
  * And more advanced iterative solvers.
* **Expanded Matrix Decompositions**:
  * Singular Value Decomposition (SVD)
  * Cholesky Decomposition
---

## Installation

To start using `linalg` in your Go project, you can install the necessary packages using `go get`:

```bash

go get github.com/rickykimani/linalg/matrix

```

```bash

go get github.com/rickykimani/linalg/vectors

```

---

## Usage Example

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

## Documentation

Comprehensive documentation for `linalg` is available on [GoDoc](https://pkg.go.dev/github.com/rickykimani/linalg). This includes detailed info on all functions and types.

---

## License

This project is licensed under the MIT License - see the LICENSE file for details.
