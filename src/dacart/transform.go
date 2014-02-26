//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

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
