package main

type Field struct {
	size   Vec3u
	planes []*Plane
}

type Plane struct {
	Normal Vec3f
	Dist   float32
}

func NewField(size Vec3u) *Field {
	res := &Field{size: size}
	res.planes = make([]*Plane, size.Mul3())
	return res
}

func (f *Field) Size() Vec3u {
	return f.size
}

func (f *Field) Cell(pos *Vec3u) *Plane {
	return f.planes[pos[2]*f.size[2]*f.size[1]+pos[1]*f.size[1]+pos[0]]
}
