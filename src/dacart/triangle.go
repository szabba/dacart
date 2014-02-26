//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

// A triangle
type Triangle []Vector

// Create a new Triangle
func NewTriangle(a, b, c Vector) Triangle {

	t := Triangle(make([]Vector, 3))

	t[0], t[1], t[2] = a, b, c

	return t
}

// Copy a Triangle
func (t Triangle) Copy() Triangle {

	return NewTriangle(t[0], t[1], t[2])
}

// Transform a Triangle using a Transform
func (t Triangle) Transform(T Transform) Triangle {

	t2 := t.Copy()

	for i, v := range t {

		t2[i] = T.Transform(v)
	}

	return t2
}
