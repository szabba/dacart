//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
)

func testCollision(t1, t2 Triangle) {

	fmt.Printf("The triangles are %v and %v.\n", t1, t2)

	doThey := "do not"
	if Collide(t1, t2) {

		doThey = "do"
	}

	fmt.Printf("The triangles %s collide.\n", doThey)
}

func main() {

	t := NewTriangle(
		NewVector(0, 0, 0), NewVector(1, 0, 0), NewVector(1, 1, 0),
	)

	t2 := NewTriangle(
		NewVector(1, 1, 0), NewVector(1, 0, 0), NewVector(0, 1, 0),
	)

	t3 := NewTriangle(
		NewVector(2, 0, 0), NewVector(3, 0, 0), NewVector(2, 1, 0),
	)

	testCollision(t, t2)
	testCollision(t, t3)
	testCollision(t2, t3)
}
