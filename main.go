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
	log.Printf("Generating %d^3 field...", FieldSize)
	field := NewField(Vec3u{FieldSize, FieldSize, FieldSize})

	log.Printf("Adding cylinder...")
	testAddCylinder(field)

	// TODO here: removeSphere(), addSphere()

	log.Printf("Closing field...")
	field.CloseField()

	return field
}

func main() {
	field := testMakeField()

	_ = field
}
