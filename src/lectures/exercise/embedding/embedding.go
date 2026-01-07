//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) CalculateBandwitdhAvg() float32 {
	var sum Bytes = 0
	for i := 0; i < 5; i++ {
		sum += b.amount[i]
	}
	return float32(sum / 5)
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) CalculateCpuTempAvg() float32 {
	var sum Celcius = 0
	for i := 0; i < 5; i++ {
		sum += c.temp[i]
	}
	return float32(sum / 5)
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) CalculateMemoryAvg() float32 {
	var sum Bytes = 0
	for i := 0; i < 5; i++ {
		sum += m.amount[i]
	}
	return float32(sum / 5)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (d *Dashboard) CalculateAndPrint() {
	fmt.Println("Bandwitdh avg:", d.CalculateBandwitdhAvg())
	fmt.Println("Cpu temp avg:", d.CalculateCpuTempAvg())
	fmt.Println("Mememory avg:", d.CalculateMemoryAvg())
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}
	d := Dashboard{bandwidth, temp, memory}
	d.CalculateAndPrint()
}
