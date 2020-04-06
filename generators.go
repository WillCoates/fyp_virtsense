package main

import (
	"math"
	"math/rand"
	"time"
)

// A Generator generates data points between min and max
type Generator func(min, max float64) func() float64

// Generators provides a mapping of generator config names to functions
var Generators = map[string]Generator{
	"random": RandomGenerator,
	"sine":   SineGenerator,
	"sin":    SineGenerator,
	"cosine": CosineGenerator,
	"cos":    CosineGenerator,
}

// RandomGenerator generates random data points between min and max
func RandomGenerator(min, max float64) func() float64 {
	return func() float64 {
		return rand.Float64()*(max-min) + min
	}
}

// SineGenerator generates data points matching a sine curve with a period of 15 minutes
func SineGenerator(min, max float64) func() float64 {
	start := time.Now()
	return func() float64 {
		elapsed := time.Now().Sub(start)
		location := elapsed.Minutes() * 2 * math.Pi / 15
		return (math.Sin(location)+1)*(max-min)/2 + min
	}
}

// CosineGenerator generates data points matching a cosine curve with a period of 15 minutes
func CosineGenerator(min, max float64) func() float64 {
	start := time.Now()
	return func() float64 {
		elapsed := time.Now().Sub(start)
		location := elapsed.Minutes() * 2 * math.Pi / 15
		return (math.Cos(location)+1)*(max-min)/2 + min
	}
}
