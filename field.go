package main

import "fmt"

type Field struct {
	size   Vec3u
	maxOft uint
	planes []*Plane
}

type Plane struct {
	Normal Vec3f
	Dist   float32
}

func NewField(size Vec3u) *Field {
	res := &Field{size: size, maxOft: size.Mul3()}
	res.planes = make([]*Plane, res.maxOft)
	for i := uint(0); i < res.maxOft; i++ {
		res.planes[i] = &Plane{}
	}
	return res
}

func (f *Field) Size() Vec3u {
	return f.size
}

func (f *Field) Cell(pos *Vec3u) *Plane {
	oft := pos[2]*f.size[2]*f.size[1] + pos[1]*f.size[1] + pos[0]
	if oft > f.maxOft {
		panic(fmt.Sprintf("invalid position in field: %v (size: %v)", pos, f.size))
	}
	return f.planes[pos[2]*f.size[2]*f.size[1]+pos[1]*f.size[1]+pos[0]]
}

func (f *Field) CloseField() {
	// Ensure periphery has dist>0
	oa := [][]int{{1, 2}, {0, 2}, {0, 1}}
	fs := f.Size()

	for a2, a := range oa {
		a0 := a[0]
		a1 := a[1]
		sideSize := Vec2u{fs[a0], fs[a1]}

		sideSize.Foreach(func(p2 *Vec2u) {
			var p3 Vec3u
			p3[a0] = p2[0]
			p3[a1] = p2[1]
			p3[a2] = 0

			cell := f.Cell(&p3)
			if cell.Dist <= 0 {
				cell.Dist = 0.5
				//cell.Normal = -Vec3::Axes[a]
			}

			p3[a2] = fs[a2] - 1

			cell = f.Cell(&p3)
			if cell.Dist <= 0 {
				cell.Dist = 0.5
				//cell.Normal = +Vec3::Axes[a]
			}
		})
	}
}
