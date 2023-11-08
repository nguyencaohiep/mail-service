package cache

import "sync"

// New TagSet.
func NewTagSet() *TagSet {
	return &TagSet{
		Data: make(map[string]struct{}),
	}
}

// TagSet is our struct that acts as a set Data structure
// with string as membertagSet.
type TagSet struct {
	Mutex sync.RWMutex
	Data  map[string]struct{}
}

// Add method to add a member to the TagSet.
func (tagSet *TagSet) Add(member string) {
	tagSet.Mutex.Lock()
	defer tagSet.Mutex.Unlock()

	tagSet.Data[member] = struct{}{}
}

// Remove method to remove a member from the TagSet.
func (tagSet *TagSet) Remove(member string) {
	tagSet.Mutex.Lock()
	defer tagSet.Mutex.Unlock()

	delete(tagSet.Data, member)
}

// IsMember method to check if a member is present in the TagSet.
func (tagSet *TagSet) IsMember(member string) bool {
	tagSet.Mutex.RLock()
	defer tagSet.Mutex.RUnlock()

	_, found := tagSet.Data[member]
	return found
}

// Members method to retrieve all members of the TagSet.
func (tagSet *TagSet) Members() []string {
	tagSet.Mutex.RLock()
	defer tagSet.Mutex.RUnlock()

	keys := make([]string, 0)
	for k := range tagSet.Data {
		keys = append(keys, k)
	}
	return keys
}

// Size method to get the cardinality of the TagSet.
func (tagSet *TagSet) Size() int {
	tagSet.Mutex.RLock()
	defer tagSet.Mutex.RUnlock()

	return len(tagSet.Data)
}

// Clear method to remove all members from the TagSet.
func (tagSet *TagSet) Clear() {
	tagSet.Mutex.Lock()
	defer tagSet.Mutex.Unlock()

	tagSet.Data = make(map[string]struct{})
}
