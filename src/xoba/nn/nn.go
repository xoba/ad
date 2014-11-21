// a simple neural network example using automatic differentiation
package nn

// without hidden units: 27096644. loss = 0.415957; acc = 77.23%; mcc = 30.10%; beta = [-1.637  1.429  0.715]

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
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
		case 0:
			for i := 0; i < 5; i++ {
				f := fmt.Sprintf("0 * %s * (1 / (1 + exp2(- (%s + %s * x1 + %s * x2))))", p(), p(), p(), p())
				layer = append(layer, f)
			}
			fmt.Fprintf(f, "f := log2( 1 + exp2(-z * (%s +  %s)))\n", p(), strings.Join(layer, " + "))
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
	eta := 0.1

	beta := randSlice(21)

	f, err := os.Create(fmt.Sprintf("loss_%d.csv", hidden))
	check(err)
	defer f.Close()
	fmt.Fprintln(f, "t,risk")

	lastTime := time.Now()
	var totalLoss, last float64
	var iterations int
	for iterations < 10000000 {
		iterations++
		x1 := rand.NormFloat64()
		y1 := rand.NormFloat64()
		var z float64
		if math.Sqrt(math.Pow(x0-x1, 2)+math.Pow(y0-y1, 2)) < r {
			z = +1
		} else {
			z = -1
		}
		v, g := ComputeAD(beta[0], beta[1], beta[2], beta[3], beta[4], beta[5], beta[6], beta[7], beta[8], beta[9], beta[10], beta[11], beta[12], beta[13], beta[14], beta[15], beta[16], beta[17], beta[18], beta[19], beta[20], x1, y1, z)
		score := log2(exp2(v-1)) / (-z)
		pt.Update(score, z == +1)
		for i := 0; i < len(beta); i++ {
			beta[i] = beta[i] - eta*g[fmt.Sprintf("b%02d", i)]
		}
		totalLoss += v
		if time.Now().Sub(lastTime) > 100*time.Millisecond {
			lastTime = time.Now()
			meanLoss := totalLoss / float64(iterations)
			var msg string
			if meanLoss > last {
				msg = "*"
				eta *= 0.99
			}
			fmt.Fprintf(f, "%d,%f\n", iterations, meanLoss)
			fmt.Printf("%1s%10d. eta=%f; risk = %f; acc=%.2f%%; mcc =%.2f%%; beta = %6.3f\n",
				msg,
				iterations,
				eta,
				meanLoss,
				100*pt.Accuracy(),
				100*pt.Mcc(),
				beta,
			)
			last = meanLoss

		}
	}
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
