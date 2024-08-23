package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

var (
	nsteps = 5
	f      = "alpha(black, %.2f) %s,\n"
	min    = 0.0
	max    = 1.0
	start  = pointFlag{0.0, "%"}
	end    = pointFlag{100.0, "%"}
)

func main() {
	flag.StringVar(&f, "f", f, "printf string")
	flag.IntVar(&nsteps, "steps", nsteps, "number of steps to iterate")
	flag.Float64Var(&min, "min", min, "minimum")
	flag.Float64Var(&max, "max", max, "maximum")
	flag.Var(&start, "start", "start percentage")
	flag.Var(&end, "end", "end percentage")
	flag.Parse()

	if start.u != end.u {
		log.Fatalf("start and end units must be the same")
	}

	var rev bool
	if min > max {
		rev = true
		min, max = max, min
	}

	vs := make([]float64, nsteps)
	ps := make([]point, nsteps)
	for i := 0; i < nsteps; i++ {
		s := float64(i) / float64(nsteps-1)
		n := easeInOutCubic(s)
		vs[i] = (max-min)*n + min
		ps[i] = point{v: (end.v-start.v)*s + start.v, u: end.u}
	}

	if rev {
		slices.Reverse(vs)
	}

	for i := range vs {
		ff := f
		if i == nsteps-1 {
			ff = strings.TrimSuffix(ff, ",\n") + "\n"
		}
		fmt.Printf(ff, vs[i], ps[i])
	}
}

type point struct {
	v float64
	u string // unit
}

type pointFlag point

var _ flag.Value = (*pointFlag)(nil)

func (p *pointFlag) Set(s string) error {
	for _, u := range []string{"%", "px", "em"} {
		if !strings.HasSuffix(s, u) {
			continue
		}

		s = strings.TrimSuffix(s, u)

		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}

		*p = pointFlag{v: n, u: u}
		return nil
	}

	return fmt.Errorf("invalid unit for %q", s)
}

func (p *pointFlag) String() string {
	return (*point)(p).String()
}

func (p point) String() string {
	return fmt.Sprintf("%.0f%s", p.v, p.u)
}

// https://easings.net/#easeInOutCubic
func easeInOutCubic(x float64) float64 {
	if x < 0.5 {
		return 4 * x * x * x
	}
	return 1 - math.Pow((-2*x)+2, 3)/2
}
