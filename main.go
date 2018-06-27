package main

import "log"

var (
	FieldSize uint = 256
)

func testAddCylinder(field *Field) {
	rad := float32(field.Size()[0]) / 10
	center := field.Size().XY().Div2()

	field.Size().Foreach(func(upos *Vec3u) {
		cell := field.Cell(upos)
		vec := upos.XY().Sub(center)
		d := vec.LenF() - rad

		if d < cell.Dist {
			cell.Dist = d
			cell.Normal = vec.Normalize().AddZ(0)
		}
	})
}

func testMakeField() *Field {
	field := NewField(Vec3u{FieldSize, FieldSize, FieldSize})

	testAddCylinder(field)

	return field
}

func main() {
	log.Printf("Generating %d^3 field...", FieldSize)

	field := testMakeField()

	_ = field
}
