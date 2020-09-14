package orderedmap

type pair struct {
	Key   string
	Value int
}

// OrderedMap is a map like type with string kays and int values.
// It implement sorting methods by values.
type OrderedMap []pair

// Len return length of a given OrderedMap
func (p OrderedMap) Len() int { return len(p) }

// Less return if a element of an OrderedMap is smaller than another
func (p OrderedMap) Less(i, j int) bool { return p[i].Value < p[j].Value }

// Swap two members of an OrderedMap
func (p OrderedMap) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
