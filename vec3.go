package main

import (
	"fmt"
	"math"
)

type Vec3u [3]uint

func (v Vec3u) String() string {
	return fmt.Sprintf("Vec3u[%u, %u, %u]", v[0], v[1], v[2])
}

func (v0 *Vec3u) Sub(v1 *Vec3u) Vec3u {
	return Vec3u{
		v0[0] - v1[0],
		v0[1] - v1[1],
		v0[2] - v1[2],
	}
}

func (v *Vec3u) Add(v0, v1 *Vec3u) {
	v[0] = v0[0] + v1[0]
	v[1] = v0[1] + v1[1]
	v[2] = v0[2] + v1[2]
}

func (v *Vec3u) Mul3() uint {
	return v[0] * v[1] * v[2]
}

func (v Vec3u) XY() Vec2u {
	return Vec2u{v[0], v[1]}
}

func (min Vec3u) ForeachRange(max Vec3u, cb func(*Vec3u)) {
	var p Vec3u

	for p[2] = min[2]; p[2] < max[2]; p[2]++ {
		for p[1] = min[1]; p[1] < max[1]; p[1]++ {
			for p[0] = min[0]; p[0] < max[0]; p[0]++ {
				cb(&p)
			}
		}
	}
}

func (max Vec3u) Foreach(cb func(*Vec3u)) {
	Vec3u{0, 0, 0}.ForeachRange(max, cb)
}

type Vec3f [3]float32

func (v Vec3f) String() string {
	return fmt.Sprintf("Vec3f[% 0.3f, % 0.3f, % 0.3f]", v[0], v[1], v[2])
}

func (v Vec3f) LenSq() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v Vec3f) Len() float32 {
	return float32(math.Sqrt(float64(v.LenSq())))

}

func (v Vec3f) Normalized() Vec3f {
	sq := v.LenSq()
	inv := 1 / math.Sqrt(float64(sq))
	return Vec3f{
		float32(float64(v[0]) * inv),
		float32(float64(v[1]) * inv),
		float32(float64(v[2]) * inv),
	}
}
