// Code generated by "go generate gonum.org/v1/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Energy represents a quantity of energy in Joules.
type Energy float64

const (
	Yottajoule Energy = 1e24
	Zettajoule Energy = 1e21
	Exajoule   Energy = 1e18
	Petajoule  Energy = 1e15
	Terajoule  Energy = 1e12
	Gigajoule  Energy = 1e9
	Megajoule  Energy = 1e6
	Kilojoule  Energy = 1e3
	Hectojoule Energy = 1e2
	Decajoule  Energy = 1e1
	Joule      Energy = 1.0
	Decijoule  Energy = 1e-1
	Centijoule Energy = 1e-2
	Millijoule Energy = 1e-3
	Microjoule Energy = 1e-6
	Nanojoule  Energy = 1e-9
	Picojoule  Energy = 1e-12
	Femtojoule Energy = 1e-15
	Attojoule  Energy = 1e-18
	Zeptojoule Energy = 1e-21
	Yoctojoule Energy = 1e-24
)

// Unit converts the Energy to a *Unit
func (e Energy) Unit() *Unit {
	return New(float64(e), Dimensions{
		LengthDim: 2,
		MassDim:   1,
		TimeDim:   -2,
	})
}

// Energy allows Energy to implement a Energyer interface
func (e Energy) Energy() Energy {
	return e
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (e *Energy) From(u Uniter) error {
	if !DimensionsMatch(u, Joule) {
		*e = Energy(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*e = Energy(u.Unit().Value())
	return nil
}

func (e Energy) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", e, float64(e))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " J"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(e))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(e))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(e))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(e))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g J)", c, e, float64(e))
	}
}