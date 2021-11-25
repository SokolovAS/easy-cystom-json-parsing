package main

type ByAge []Person
type ByCity []Place

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByCity) Len() int           { return len(a) }
func (a ByCity) Less(i, j int) bool { return len(a[i].City) < len(a[j].City) }
func (a ByCity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
