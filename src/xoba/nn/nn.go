// a simple neural network example using automatic differentiation
package nn

// without hidden units: 27096644. loss = 0.415957; acc = 77.23%; mcc = 30.10%; beta = [-1.637  1.429  0.715]

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

//go:generate run nn -gen
//go:generate run compile -formula=nn.txt -output nn_ad.go -templates "../ad/parser/templates" -package nn -main=false -time=false
func Run(args []string) {
	var gen bool
	var hidden int
	flags := flag.NewFlagSet("parse", flag.ExitOnError)
	flags.BoolVar(&gen, "gen", false, "whether to generate formula")
	flags.IntVar(&hidden, "hidden", 5, "number of hidden units")
	flags.Parse(args)
	switch hidden {
	case 0:
	case 5:
	default:
		panic("unsupported hidden units")
	}
	if gen {
		// 2-d input, and one output; classifies as true/false (like logistic regression)
		f, err := os.Create("nn.txt")
		check(err)
		defer f.Close()
		var layer []string
		var parms []string
		var betas []string
		p := func() string {
			i := len(parms)
			x := fmt.Sprintf("b%02d", i)
			parms = append(parms, x)
			betas = append(betas, fmt.Sprintf("beta[%d]", i))
			return x
		}
		switch hidden {
		case 0: // zero out the hidden=5 stuff...
			var a, b string
			for i := 0; i < 5; i++ {
				a = p()
				b = p()
				f := fmt.Sprintf("0 * %s * (1 / (1 + exp2(- (%s + %s * x1 + %s * x2))))", a, b, p(), p())
				layer = append(layer, f)
			}
			fmt.Fprintf(f, "f := log2( 1 + exp2(-z * (%s + %s + %s * x1 + %s * x2)))\n", p(), strings.Join(layer, " + "), a, b)
		case 5:
			for i := 0; i < 5; i++ {
				f := fmt.Sprintf("%s * (1 / (1 + exp2(- (%s + %s * x1 + %s * x2))))", p(), p(), p(), p())
				layer = append(layer, f)
			}
			fmt.Fprintf(f, "f := log2( 1 + exp2(-z * (%s +  %s)))\n", p(), strings.Join(layer, " + "))
		default:
			panic("unsupported hidden units")
		}
		return
	}

	pt := &PerformanceTracker{}

	// a circle in quadrant I:
	x0, y0 := 1.0, 0.5
	r := 1.0
	var eta float64
	if hidden == 5 {
		eta = 0.01
	} else {
		eta = 0.0001
	}

	beta := randSlice(21)

	f, err := os.Create(fmt.Sprintf("loss_%d.csv", hidden))
	check(err)
	defer f.Close()
	fmt.Fprintln(f, "t,risk,acc")

	w, h := 1000, 1000

	trial := func() (x float64, y float64, z float64) {
		x1 := rand.NormFloat64()
		y1 := rand.NormFloat64()
		if math.Sqrt(math.Pow(x0-x1, 2)+math.Pow(y0-y1, 2)) < r {
			z = +1
		} else {
			z = -1
		}
		return x1, y1, z
	}

	run := func(x1, y1, z float64) (float64, map[string]float64, float64) {
		v, g := ComputeAD(beta[0], beta[1], beta[2], beta[3], beta[4], beta[5], beta[6], beta[7], beta[8], beta[9], beta[10], beta[11], beta[12], beta[13], beta[14], beta[15], beta[16], beta[17], beta[18], beta[19], beta[20], x1, y1, z)
		score := log2(exp2(v)-1) / (-z)
		return v, g, score
	}

	plot := func(name string) {
		img := image.NewRGBA(image.Rect(0, 0, w, h))
		draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 255}}, image.ZP, draw.Src)
		for n := 0; n < 100000; n++ {
			x, y, z := trial()
			_, _, score := run(x, y, z)
			j := int(-y*float64(h)/10) + w/2
			i := int(x*float64(w)/10) + w/2
			if i < 0 || i >= w {
				continue
			}
			if j < 0 || j >= h {
				continue
			}
			var c color.RGBA
			if score > 0 {
				if z == +1 {
					c = color.RGBA{200, 200, 200, 255}
				} else {
					c = color.RGBA{200, 0, 0, 255}
				}
			} else {
				if z == -1 {
					c = color.RGBA{100, 100, 100, 255}
				} else {
					c = color.RGBA{100, 0, 0, 255}
				}
			}
			img.SetRGBA(i, j, c)
		}
		save(name, img)
	}

	skip := 10.0
	var recentTotal, totalLoss, last float64
	var iterations, recentIterations, frames int
	for iterations < 2000000 {

		// run a trial
		x1, y1, z := trial()
		v, g, score := run(x1, y1, z)
		totalLoss += v
		recentTotal += v
		pt.Update(score, z == +1)

		// gradient descent
		for i := 0; i < len(beta); i++ {
			beta[i] = beta[i] - eta*g[fmt.Sprintf("b%02d", i)]
		}

		// logging
		if recentIterations%int(skip) == 0 {
			skip *= 1.05
			risk := totalLoss / float64(iterations)
			var msg string
			if risk > last {
				msg = "*"
			}
			fmt.Fprintf(f, "%d,%f,%f\n", iterations, risk, 100*pt.Accuracy())
			fmt.Printf("%1s%10d. eta=%f; risk = %f / %f; %v; beta = %6.3f\n",
				msg,
				iterations,
				eta,
				risk,
				recentTotal/float64(recentIterations),
				pt,
				beta,
			)
			last = risk
			recentTotal = 0
			recentIterations = 0
			plot(fmt.Sprintf("img_%d_%05d.jpg", hidden, frames))
			frames++
		}

		iterations++
		recentIterations++
	}

	cmd := exec.Command("ffmpeg", "-y", "-r", "30", "-b", "1800", "-i", fmt.Sprintf("img_%d%%05d.jpg", hidden), fmt.Sprintf("test_%d.mp4", hidden))
	fmt.Println(cmd.Args)
	check(cmd.Run())
}

func save(name string, img image.Image) {
	f, err := os.Create(name)
	check(err)
	jpeg.Encode(f, img, &jpeg.Options{100})
	f.Close()
}

func randSlice(n int) (out []float64) {
	for i := 0; i < n; i++ {
		out = append(out, rand.NormFloat64())
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
