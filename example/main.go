package main

import (
	"log"
	"math/rand"

	"github.com/rdoorn/kalmanfilter"
)

func main() {

	staticData := staticArray(5, 20)
	noiseData := noiseArray(5, 5, 20)

	log.Printf("static: %v\n", staticData)
	log.Printf("noise: %v\n", noiseData)

	kf := kalmanfilter.New(0.01, 4)
	for _, i := range noiseData {
		log.Printf("%f", kf.Filter(float64(i)))
	}
}

func staticArray(val, count int) (r []int) {
	for i := 0; i < count; i++ {
		r = append(r, val)
	}
	return
}

func noiseArray(val, noise, count int) (r []int) {
	for i := 0; i < count; i++ {
		n := rand.Intn(noise)
		p := rand.Intn(2)
		if p == 0 {
			r = append(r, val+n)
		} else {
			r = append(r, val-n)
		}
	}
	return
}
