package vectors

import (
	"errors"
	"math"
)

// CartesianToPolar converts a 2D vector from Cartesian to polar coordinates.
//
// Polar coordinates represent a point by its distance from the origin (r)
// and the angle (θ) from the positive x-axis in the counterclockwise direction.
//
// Parameters:
//   - v: Input 2D vector in Cartesian coordinates (x,y)
//
// Returns:
//   - r: Radius (distance from origin)
//   - theta: Angle in radians (-π to π), measured counterclockwise from positive x-axis
//   - err: Error if input vector is not 2D
//
// The angle θ is normalized to the range [-π, π] by the Atan2 function.
// For the zero vector (0,0), the angle is technically undefined but returns 0.
func CartesianToPolar[T int | float64](v Vector[T]) (r, theta float64, err error) {
	if len(v) != 2 {
		return 0, 0, errors.New("vector must be 2D")
	}

	x, y := float64(v[0]), float64(v[1])
	r = math.Sqrt(x*x + y*y)
	theta = math.Atan2(y, x)

	return r, theta, nil
}

// PolarToCartesian converts polar coordinates to a 2D Cartesian vector.
//
// This function performs the inverse of CartesianToPolar, creating a
// Cartesian vector from a radius and angle.
//
// Parameters:
//   - r: Radius (distance from origin), must be non-negative
//   - theta: Angle in radians, measured counterclockwise from positive x-axis
//
// Returns:
//   - Vector[float64]: 2D vector in Cartesian coordinates (x,y)
//   - error: Error if radius is negative
//
// conversion formulas:
//
//	x = r * cos(θ)
//
//	y = r * sin(θ)
func PolarToCartesian(r, theta float64) (Vector[float64], error) {
	if r < 0 {
		return nil, errors.New("radius must be non-negative")
	}

	x := r * math.Cos(theta)
	y := r * math.Sin(theta)

	return Vector[float64]{x, y}, nil
}

// CartesianToSpherical converts a 3D vector from Cartesian to spherical coordinates.
//
// Spherical coordinates represent a point by:
//   - rho: Distance from origin
//   - theta: Azimuthal angle in x-y plane from x-axis (-π to π), counterclockwise
//   - phi: Polar/zenith angle from positive z-axis (0 to π)
//
// Parameters:
//   - v: Input 3D vector in Cartesian coordinates (x,y,z)
//
// Returns:
//   - rho: Radius/distance from origin
//   - theta: Azimuthal angle in radians (-π to π)
//   - phi: Polar/zenith angle in radians (0 to π)
//   - err: Error if input vector is not 3D
func CartesianToSpherical[T int | float64](v Vector[T]) (rho, theta, phi float64, err error) {
	if len(v) != 3 {
		return 0, 0, 0, errors.New("vector must be 3D")
	}

	x, y, z := float64(v[0]), float64(v[1]), float64(v[2])

	// Distance from origin
	rho = math.Sqrt(x*x + y*y + z*z)

	// Handle zero vector case
	if rho < 1e-10 {
		return 0, 0, 0, nil
	}

	// Azimuthal angle in x-y plane (same as polar coordinates)
	theta = math.Atan2(y, x)

	// Polar/zenith angle from z-axis
	phi = math.Acos(z / rho)

	return rho, theta, phi, nil
}

// SphericalToCartesian converts spherical coordinates to a 3D Cartesian vector.
//
// Parameters:
//   - rho: Radius/distance from origin, must be non-negative
//   - theta: Azimuthal angle in radians, measured counterclockwise from positive x-axis
//   - phi: Polar/zenith angle in radians from positive z-axis
//
// Returns:
//   - Vector[float64]: 3D vector in Cartesian coordinates (x,y,z)
//   - error: Error if radius is negative
//
// Conversion formulas:
//
// x = rho * sin(phi) * cos(theta)
//
// y = rho * sin(phi) * sin(theta)
//
// z = rho * cos(phi)
func SphericalToCartesian(rho, theta, phi float64) (Vector[float64], error) {
	if rho < 0 {
		return nil, errors.New("radius must be non-negative")
	}

	sinPhi := math.Sin(phi)
	x := rho * sinPhi * math.Cos(theta)
	y := rho * sinPhi * math.Sin(theta)
	z := rho * math.Cos(phi)

	return Vector[float64]{x, y, z}, nil
}

// CartesianToCylindrical converts a 3D vector from Cartesian to cylindrical coordinates.
//
// Cylindrical coordinates represent a point by:
//   - r: Radial distance from z-axis
//   - theta: Azimuthal angle in x-y plane from x-axis (-π to π), counterclockwise
//   - z: Height along the z-axis
//
// Parameters:
//   - v: Input 3D vector in Cartesian coordinates (x,y,z)
//
// Returns:
//   - r: Radial distance from z-axis
//   - theta: Azimuthal angle in radians (-π to π)
//   - z: Height along the z-axis
//   - err: Error if input vector is not 3D
func CartesianToCylindrical[T int | float64](v Vector[T]) (r, theta, z float64, err error) {
	if len(v) != 3 {
		return 0, 0, 0, errors.New("vector must be 3D")
	}

	x, y := float64(v[0]), float64(v[1])
	z = float64(v[2])

	// Radial distance from z-axis (same as polar r)
	r = math.Sqrt(x*x + y*y)

	// Azimuthal angle in x-y plane (same as polar theta)
	theta = math.Atan2(y, x)

	return r, theta, z, nil
}

// CylindricalToCartesian converts cylindrical coordinates to a 3D Cartesian vector.
//
// Parameters:
//   - r: Radial distance from z-axis, must be non-negative
//   - theta: Azimuthal angle in radians, measured counterclockwise from positive x-axis
//   - z: Height along the z-axis
//
// Returns:
//   - Vector[float64]: 3D vector in Cartesian coordinates (x,y,z)
//   - error: Error if radial distance is negative
//
// Conversion formulas:
//
// x = r * cos(theta)
//
// y = r * sin(theta)
//
// z = z
func CylindricalToCartesian(r, theta, z float64) (Vector[float64], error) {
	if r < 0 {
		return nil, errors.New("radial distance must be non-negative")
	}

	x := r * math.Cos(theta)
	y := r * math.Sin(theta)

	return Vector[float64]{x, y, z}, nil
}
