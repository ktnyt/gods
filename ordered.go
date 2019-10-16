package gods

import "fmt"

// Entry is a key-value pair.
type Entry struct {
	Key   string
	Value string
}

// Ordered is an ordered dictionary type.
type Ordered []Entry

// NewOrdered returns a new ordered dictionary with the given entries.
func NewOrdered(entries ...Entry) *Ordered {
	dict := new([]Entry)
	*dict = entries
	return (*Ordered)(dict)
}

// Len returns the length of the dictionary.
func (dict Ordered) Len() int {
	return len(dict)
}

// Iter returns the underlying slice.
func (dict Ordered) Iter() []Entry {
	return []Entry(dict)
}

// Indices lists the indices for entries matching the filter.
func (dict Ordered) Indices(filter func(Entry) bool) []int {
	is := make([]int, 0)
	for i, entry := range dict {
		if filter(entry) {
			is = append(is, i)
		}
	}
	return is
}

// Get returns the first entry with the given key.
func (dict Ordered) Get(key string) string {
	is := dict.Indices(func(entry Entry) bool { return entry.Key == key })
	if len(is) == 0 {
		panic(fmt.Errorf("Ordered does not have key: %s", key))
	}
	return dict[is[0]].Value
}

// All returns all entries with the given key.
func (dict Ordered) All(key string) []string {
	is := dict.Indices(func(entry Entry) bool { return entry.Key == key })
	ret := make([]string, len(is))
	for j, i := range is {
		ret[j] = dict[i].Value
	}
	return ret
}

// Add an entry with the given key and value.
func (dict *Ordered) Add(key, value string) {
	*dict = append(*dict, Entry{Key: key, Value: value})
}

// Delete all entries with the given key.
func (dict *Ordered) Delete(key string) {
	is := dict.Indices(func(entry Entry) bool { return entry.Key != key })
	entries := make([]Entry, len(is))
	for j, i := range is {
		entries[j] = (*dict)[i]
	}
	*dict = Ordered(entries)
}

// Set will delete all entries with the given key and insert a new entry.
func (dict *Ordered) Set(key, value string) {
	dict.Delete(key)
	dict.Add(key, value)
}
