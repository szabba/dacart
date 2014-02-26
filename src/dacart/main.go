//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
)

func main() {

	t := NewTriangle(
		NewVector(0, 0, 0), NewVector(1, 0, 0), NewVector(1, 1, 0),
	)

	d := Translation{By: NewVector(1, 2, 3)}

	fmt.Println("Triangle:", t)
	fmt.Println("Translating by:", d.By)

	t = t.Transform(d)

	fmt.Println("Translated triangle:", t)
}
