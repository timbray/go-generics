package gg

import "fmt"

type container[S comparable] struct {
	label string
	list  []S
}

func (c *container[S]) size() int {
	return len(c.list)
}

type fish struct {
	species string
}

type fishTank container[fish]

func (t *fishTank) fishCount() string {
	// following doesn't compile:
	// 1. Unresolved reference 'size'
	// 2. t.size undefined (type *fishTank has no field or method size)
	return fmt.Sprintf("How many fishies? %d!", t.size())
}

type dataRichFish struct {
	species string
	datum   [8]float64
}

type stingyFish struct {
	species string
}

type flexibleFish interface {
	getSpecies() string
	getDatum() [8]float64
}

func (f *dataRichFish) getSpecies() string {
	return f.species
}
func (f *dataRichFish) getDatum() [8]float64 {
	return f.datum
}

var dataNotProvided [8]float64

func (f *stingyFish) getSpecies() string {
	return f.species
}
func (f *stingyFish) getDatum() [8]float64 {
	return dataNotProvided
}

// won't compile because "flexibleFish does not implement comparable"
type flexibleTank container[flexibleFish]

func (t *flexibleTank) contains(target flexibleFish) bool {
	for _, fish := range t.list {
		if fish == target {
			return true
		}
	}
	return false
}

type comparableFlexibleFish interface {
	comparable
	getSpecies() string
	getDatum() [8]float64
}

// won't compile for two reasons:
// 1. "interface includes constraint elements, can only be used in type parameters
// 2. "interface is (or embeds) comparable"
type comparableFlexibleFishTank container[comparableFlexibleFish]
