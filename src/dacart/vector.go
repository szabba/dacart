//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"math"
)

// A three component vector
type Vector []float64

// Create new zero vector
func NewZeroVector() Vector {

	return Vector(make([]float64, 3))
}

// Create new vector
func NewVector(x, y, z float64) Vector {

	v := NewZeroVector()

	v[0], v[1], v[2] = x, y, z

	return v
}

// Add two vectors
func (a Vector) Plus(b Vector) (c Vector) {

	c = NewZeroVector()

	for i, _ := range a {

		c[i] = a[i] + b[i]
	}

	return c
}

// Invert a vector's direction
func (a Vector) Negate() (minusA Vector) {

	minusA = NewZeroVector()

	for i, aComponent := range a {

		minusA[i] = -aComponent
	}

	return minusA
}

// Subtract two vectors
func (a Vector) Minus(b Vector) Vector {

	return a.Plus(b.Negate())
}

// Take the dot product of two vectors
func (a Vector) Dot(b Vector) float64 {

	sum := 0.0

	for i, _ := range a {

		sum += a[i] * b[i]
	}

	return sum
}

// Calculate the norm of a vector
func (a Vector) Norm() float64 {

	return math.Sqrt(a.Dot(a))
}