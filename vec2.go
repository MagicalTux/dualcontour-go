package main

import "math"

type Vec2u [2]uint

func (v Vec2u) MulAllF(val float32) Vec2u {
	return Vec2u{uint(float32(v[0]) * val), uint(float32(v[1]) * val)}
}

func (v Vec2u) Div2() Vec2u {
	return Vec2u{v[0] / 2, v[1] / 2}
}

func (v0 Vec2u) Sub(v1 Vec2u) Vec2u {
	return Vec2u{
		v0[0] - v1[0],
		v0[1] - v1[1],
	}
}

func (v Vec2u) LenSq() uint {
	return v[0]*v[0] + v[1]*v[1]
}

func (v Vec2u) Len() uint {
	return uint(math.Sqrt(float64(v.LenSq())))
}

func (v Vec2u) LenF() float32 {
	return float32(math.Sqrt(float64(v.LenSq())))
}

func (v Vec2u) Normalize() Vec2f {
	return Vec2f{float32(v[0]), float32(v[1])}.Normalize()
}

type Vec2f [2]float32

func (v Vec2f) LenSq() float32 {
	return v[0]*v[0] + v[1]*v[1]
}

func (v Vec2f) Len() float32 {
	return float32(math.Sqrt(float64(v.LenSq())))
}

func (v Vec2f) Normalize() Vec2f {
	sq := v.LenSq()
	inv := 1 / math.Sqrt(float64(sq))
	return Vec2f{
		float32(float64(v[0]) * inv),
		float32(float64(v[1]) * inv),
	}
}

func (v Vec2f) AddZ(z float32) Vec3f {
	return Vec3f{v[0], v[1], z}
}
