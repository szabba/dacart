//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

// A straight line represented as a position vector on the line and a direction
// vector
type Line struct {
	Origin, Direction Vector
}

// Transform a line using T. Should work ok with skews, translations, rotations
// and reflections.
func (l Line) Transform(T Transform) (k Line) {

	k.Origin = T.Transform(l.Origin)
	k.Direction = T.Transform(l.Direction)

	return k
}

// Image a given point projects on the line
func (l Line) PointImage(v Vector) (u Vector) {

	unitDir := l.Direction.Unit()

	_, oRej := l.Origin.ProjectionRejection(unitDir)

	vPrj, _ := v.ProjectionRejection(unitDir)

	return oRej.Plus(vPrj)
}
