package storage

type peer struct {
	id              int
	maxCapacity     int
	currentCapacity int
	children        []*peer
}
