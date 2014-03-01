//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

type Polygon interface {
	Vertices() []Vector
}

func Borders(p Polygon) (borders []Vector) {

	vertices := p.Vertices()

	vCount := len(vertices)

	borders = make([]Vector, vCount)

	for i, _ := range vertices {

		borders[i] = vertices[(i+1)%vCount].Minus(vertices[i])
	}

	return borders
}

func Collide(a, b Polygon) bool {

	for _, axis := range normalAxes(a) {

		if separateShadows(a, b, axis) {

			return false
		}
	}

	for _, axis := range normalAxes(b) {

		if separateShadows(a, b, axis) {

			return false
		}
	}

	return true
}

func normalAxes(p Polygon) (axes []Line) {

	borders := Borders(p)

	axes = make([]Line, len(borders))

	for i, border := range borders {

		next := borders[(i+1)%len(borders)]

		axes[i].Origin = NewZeroVector()
		axes[i].Direction = next.Rejection(border).Unit().Negate()
	}

	return axes
}

func separateShadows(a, b Polygon, axis Line) bool {

	aMin, aMax := shadow(a, axis)
	bMin, bMax := shadow(b, axis)

	return aMax < bMin || bMax < aMin
}

func shadow(p Polygon, axis Line) (min, max float64) {

	for i, vertex := range p.Vertices() {

		image := axis.PointImage(vertex).Norm()

		if i == 0 {

			min, max = image, image

		} else {

			if image < min {
				min = image
			}

			if image > max {
				max = image
			}
		}
	}

	return min, max
}
