// a simple neural network example using automatic differentiation
package nn

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

//go:generate run nn -gen
//go:generate run compile -formula=nn.txt -output nn_ad.go -templates "../ad/parser/templates" -package nn -main=false -time=false
func Run(args []string) {
	var gen bool
	flags := flag.NewFlagSet("parse", flag.ExitOnError)
	flags.BoolVar(&gen, "gen", false, "whether to generate formula")
	flags.Parse(args)
	if gen {
		// 2-d input, and one output; classifies as true/false (like logistic regression)
		fmt.Println("generating formula")
		f, err := os.Create("nn.txt")
		check(err)
		defer f.Close()
		fmt.Fprintln(f, "f := log( 1 + exp(-z * (b0 +  b1 * x1 + b2 * y1)))")
		return
	}

	// a circcle in quadrant I:
	x0, y0 := 1.0, 0.5
	r := 1.0
	eta := 0.00001

	beta := make([]float64, 3)

	var tp, tn, fn, fp int
	var lastTime time.Time
	var totalLoss, last float64
	var iterations int
	for {
		iterations++
		x1 := rand.NormFloat64()
		y1 := rand.NormFloat64()
		var z float64
		if math.Sqrt(math.Pow(x0-x1, 2)+math.Pow(y0-y1, 2)) < r {
			z = +1
		} else {
			z = -1
		}
		v, g := ComputeAD(beta[0], beta[1], beta[2], x1, y1, z)
		totalLoss += v
		{
			f := beta[0] + beta[1]*x1 + beta[2]*y1
			if f > 0 {
				if z > 0 {
					tp++
				} else {
					fp++
				}
			} else {
				if z > 0 {
					fn++
				} else {
					tn++
				}
			}
		}
		for i := 0; i < len(beta); i++ {
			beta[i] = beta[i] - eta*g[fmt.Sprintf("b%d", i)]
		}
		if time.Now().Sub(lastTime) > 100*time.Millisecond {
			lastTime = time.Now()
			meanLoss := totalLoss / float64(iterations)
			var msg string
			if meanLoss > last {
				msg = "*"
			}
			p := tp + fn
			n := fp + tn
			acc := 100 * float64(tp+tn) / float64(p+n)
			mcc := 100 * float64(tp*tn-fp*fn) / (math.Sqrt(float64(tp+fp)) * math.Sqrt(float64(tp+fn)) * math.Sqrt(float64(tn+fp)) * math.Sqrt(float64(tn+fn)))
			fmt.Printf("%1s%10d. loss = %f; acc = %.2f%%; mcc = %.2f%%; beta = %6.3f\n",
				msg,
				iterations,
				meanLoss,
				acc,
				mcc,
				beta,
			)
			last = meanLoss

		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
