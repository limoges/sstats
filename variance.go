package sstats

import "fmt"

type Variance struct {
	n                      int
	oldM, newM, oldS, newS float64
	sum                    float64
}

func (vr *Variance) Push(values ...float64) {
	for _, v := range values {
		vr.n++
		if vr.n == 1 {
			vr.oldM = v
			vr.newM = v
			vr.oldS = 0.0
		} else {
			vr.newM = vr.oldM + (v-vr.oldM)/float64(vr.n)
			vr.newS = vr.oldS + (v-vr.oldM)*(v-vr.newM)

			vr.oldM = vr.newM
			vr.oldS = vr.newS
		}
	}
}

func (v Variance) Value() (float64, bool) {
	if v.n < 1 {
		return 0.0, false
	}
	if v.n == 1 {
		return 0.0, true
	}
	return v.newS / float64(v.n-1), true
}

func (v Variance) String() string {
	variance, valid := v.Value()
	if valid {
		return fmt.Sprintf("Var:%v", variance)
	}
	return fmt.Sprintf("Var:-")
}
