package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"github.com/fjmendes1994/adrng/rng"
)

func main() {
	sizeFlag := flag.Int("size", 1000, "")
	flag.Parse()

	size := *sizeFlag

	rngResults := make(map[string]int64)
	goResults := make(map[string]int64)

	rng := rng.New(13)

	for i := 0; i < size; i++ {
		rngAN := rng.Float64()
		goAN := rand.Float64()

		rngResults[fmt.Sprintf("%.1f", rngAN)] = rngResults[fmt.Sprintf("%.1f", rngAN)] + 1
		goResults[fmt.Sprintf("%.1f", goAN)] = goResults[fmt.Sprintf("%.1f", goAN)] + 1
	}

	adrng := fmt.Sprintf("AD - RNG - %d", size)
	gorng := fmt.Sprintf("GOLANG - RNG %d", size)

	fmt.Println(adrng, " => ", rngResults)
	fmt.Println(gorng, " => ", goResults)

	err := plotResults(adrng, rngResults)
	if err != nil {
		fmt.Printf("[ERROR] %s \n", err.Error())
		os.Exit(1)
	}
	err = plotResults(gorng, goResults)
	if err != nil {
		fmt.Printf("[ERROR] %s \n", err.Error())
		os.Exit(1)
	}
}

func plotResults(generator string, data map[string]int64) error {
	values := plotter.Values{float64(data["0.0"]), float64(data["0.1"]), float64(data["0.2"]),
		float64(data["0.3"]), float64(data["0.4"]), float64(data["0.5"]), float64(data["0.6"]),
		float64(data["0.7"]), float64(data["0.8"]), float64(data["0.9"]), float64(data["1.0"]),
	}

	p, err := plot.New()
	if err != nil {
		return err
	}
	p.Title.Text = generator

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(values, w)
	if err != nil {
		return err
	}
	barsA.LineStyle.Width = vg.Length(2)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = 1

	p.Add(barsA)
	p.NominalX("0.0", "0.1", "0.2", "0.3", "0.4", "0.5", "0.6", "0.7", "0.8", "0.9", "1.0")

	err = os.Mkdir("results", 0755)
	if err != nil {
		if err.Error() != "mkdir results: file exists" {
			return err
		}
	}

	if err := p.Save(7*vg.Inch, 5*vg.Inch, fmt.Sprintf("./results/%s-chart.png", generator)); err != nil {
		return err
	}

	return nil
}
