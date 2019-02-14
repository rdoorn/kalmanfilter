package main

import (
	"log"
	"testing"

	"github.com/rdoorn/kalmanfilter"
)

func Test1(t *testing.T) {
	kf := kalmanfilter.New(1, 1)

	log.Printf("test1: %f", kf.Filter(3))
	log.Printf("test1: %f", kf.Filter(2))
	log.Printf("test1: %f", kf.Filter(1))
}

func Test(t *testing.T) {
	main()
}
