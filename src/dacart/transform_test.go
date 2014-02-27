//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"math"
	"testing"
)

func doesntRotateAroundItself(a Vector, t *testing.T) {

	rot := NewAxisAngleRotation(a, math.Pi/2)

	if !rot.Transform(a).Equal(a) {

		t.Fatalf(
			"Rotating %v around itself by %f should have no effect on it "+
				"(returns %v).",
			a, rot.Angle(), rot.Transform(a),
		)
	}
}

func TestAngleAxisRotation(t *testing.T) {

	eX, eY, eZ := UnitVectors()

	doesntRotateAroundItself(eX, t)
	doesntRotateAroundItself(eY, t)
	doesntRotateAroundItself(eZ, t)
}
