//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"math"
)

// A Vector -> Vector transformation
type Transform interface {
	Transform(Vector) Vector
}

// A translation transform
type Translation struct {
	By Vector
}

// Translates a Vector by t.By.
func (t Translation) Transform(v Vector) Vector {

	return t.By.Plus(v)
}

// A rotation transform
type AxisAngleRotation struct {
	axis  Vector
	angle float64
}

// Create an AxisAngleRotation
func NewAxisAngleRotation(axis Vector, angle float64) *AxisAngleRotation {

	aar := new(AxisAngleRotation)

	aar.axis = axis
	aar.angle = angle

	return aar
}

// Angle of rotation
func (aar *AxisAngleRotation) Angle() float64 {

	return aar.angle
}

// Axis of rotation
func (aar *AxisAngleRotation) Axis() Vector {

	return aar.axis.Copy()
}

// Rotates a vector by aar.Angle around an axis parallel to aar.Axis going
// through the origin point
func (aar AxisAngleRotation) Transform(v Vector) (vRot Vector) {

	cos, sin := math.Cos(aar.angle), math.Sin(aar.angle)
	unit := aar.axis.Unit()

	vRot = NewZeroVector()

	vRot = vRot.Plus(v.Scale(cos))
	vRot = vRot.Plus(unit.Cross(v).Scale(sin))
	vRot = vRot.Plus(unit.Scale(unit.Dot(v) * (1 - cos)))

	return vRot
}
