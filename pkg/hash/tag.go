package hash

import (
	"encoding/json"
	"strings"
	"sync"
)

// SyncTags is use like a map[string]string, but is different from Tags. The SyncTags is thread-safe. And use read-write
// lock.
type SyncTags struct {
	Tags Tags `json:"Tags"`

	// the locker
	lock sync.RWMutex
}

// FromMap will load SyncTags from a map[string]string.
func (st *SyncTags) FromMap(m map[string]string) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.Tags.FromMap(m)
}

// SyncForEach will invoke ForEach and do not block process as possible. But the value maybe be outdated.
func (st *SyncTags) SyncForEach(f func(key, value string)) {
	st.lock.RLock()
	newTagMap := st.Tags.Map()
	st.lock.RUnlock()

	nt := Tags{}
	nt.FromMap(newTagMap)
	nt.ForEach(f)
}

// ForEach receive a function, and will invoke it for all data one by one.
func (st *SyncTags) ForEach(f func(key, value string)) {
	st.lock.RLock()
	defer st.lock.RUnlock()

	st.Tags.ForEach(f)
}

// Map will convert Tags to a map[string]string.
func (st *SyncTags) Map() map[string]string {
	st.lock.RLock()
	defer st.lock.RUnlock()

	return st.Tags.Map()
}

func (st *SyncTags) Length() int {

	return st.Tags.Length()
}

// Put to set a value to the key, and if the key is exist, cover it.
func (st *SyncTags) Put(key, value string) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.Tags.Put(key, value)
}

// Get return the value by specify key
func (st *SyncTags) Get(key string) (string, bool) {
	st.lock.RLock()
	defer st.lock.RUnlock()

	return st.Tags.Get(key)
}

// Delete will delete a key from SyncTags
func (st *SyncTags) Delete(key string) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.Tags.Delete(key)
}

// Hash will return a hash value of Tags
func (st *SyncTags) Hash() string {
	st.lock.RLock()
	defer st.lock.RUnlock()

	return st.Tags.Hash()
}

func (st *SyncTags) IsHash(hashValue string) bool {
	st.lock.RLock()
	defer st.lock.RUnlock()

	return st.Tags.IsHash(hashValue)
}

// Serialization return a []byte by json.MarshalIndent from SyncTags.Tags
func (st *SyncTags) Serialization() []byte {
	st.lock.RLock()
	defer st.lock.RUnlock()

	return st.Tags.Serialization()
}

// Deserialization will parse bytes to Tag
func (st *SyncTags) Deserialization(bytes []byte) error {
	st.lock.Lock()
	defer st.lock.Unlock()
	return st.Tags.Deserialization(bytes)
}

// Tags is use like a map[string]string, but it is implemented by array of Tag. And Tags is Hashable. And the Tags is
// un-thread-safe.
type Tags struct {

	// TagArray save all tag info, 'why not use []string?' because for the Serialization and Deserialization. The mapper
	// will be ignored.
	TagArray []Tag `json:"tags"`

	mapper map[string]int
}

// FromMap will load Tags from a map[string]string.
func (t *Tags) FromMap(m map[string]string) {
	for key, value := range m {
		t.Put(key, value)
	}
}

// ForEach receive a function, and will invoke it for all data one by one.
func (t *Tags) ForEach(f func(key, value string)) {
	t.preDo()
	for key, index := range t.mapper {
		f(key, t.TagArray[index].Value)
	}
}

// Map will convert Tags to a map[string]string.
func (t *Tags) Map() map[string]string {
	m := make(map[string]string)
	t.ForEach(func(key, value string) {
		m[key] = value
	})
	return m
}

func (t Tags) Length() int {
	return len(t.TagArray)
}

func (t *Tags) preDo() {
	if t.mapper == nil || len(t.mapper) != len(t.TagArray) {
		t.mapper = map[string]int{}
	}
	if len(t.mapper) != len(t.TagArray) {
		for index, tagItem := range t.TagArray {
			t.mapper[tagItem.Key] = index + 1
		}
	}
}

// Put to set a value to the key, and if the key is exist, cover it.
func (t *Tags) Put(key, value string) {
	t.preDo()

	if tagIndex, ok := t.mapper[key]; ok {
		t.TagArray[tagIndex].Value = value
	} else {
		t.TagArray = append(t.TagArray, Tag{Key: key, Value: value})
		t.mapper[key] = len(t.TagArray) - 1
	}
}

// Get return the value by specify key
func (t *Tags) Get(key string) (string, bool) {
	t.preDo()

	if tagIndex, ok := t.mapper[key]; ok {
		return t.TagArray[tagIndex].Value, true
	} else {
		return "", false
	}
}

func (t *Tags) Delete(key string) {
	if tagIndex, ok := t.mapper[key]; ok {
		t.TagArray = append(t.TagArray[:tagIndex], t.TagArray[tagIndex+1:]...)
		delete(t.mapper, key)
		for mk, mv := range t.mapper {
			if mv >= tagIndex {
				t.mapper[mk] = tagIndex - 1
			}
		}
	} else {
		return
	}
}

// Hash will return a hash value of Tags
func (t Tags) Hash() string {
	hashValue := make([]string, 0, len(t.TagArray))
	for _, item := range t.TagArray {
		hashValue = append(hashValue, item.Hash())
	}

	return SHA256([]byte(strings.Join(hashValue, "-")))
}

func (t Tags) IsHash(hashValue string) bool {
	return IsHash(&t, hashValue)
}

// Serialization return a []byte by json.MarshalIndent
func (t Tags) Serialization() []byte {
	bytes, _ := json.MarshalIndent(t, "", "")
	return bytes
}

// Deserialization will parse bytes to Tag
func (t *Tags) Deserialization(bytes []byte) error {
	if err := json.Unmarshal(bytes, t); err != nil {
		return err
	}
	// rebuild the cache
	t.mapper = map[string]int{}
	for index, item := range t.TagArray {
		t.mapper[item.Key] = index
	}
	return nil
}

// Tag is a Hashable key-value struct. The key and value must a string.
type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Hash will return a hash value by SHA256
func (t Tag) Hash() string {
	bytes := []byte(t.Key + t.Value)
	return SHA256(bytes)
}

func (t Tag) IsHash(hashValue string) bool {
	return IsHash(&t, hashValue)
}

// Serialization return a []byte by json.MarshalIndent
func (t Tag) Serialization() []byte {
	bytes, _ := json.MarshalIndent(t, "", "")
	return bytes
}

// Deserialization will parse bytes to Tag
func (t *Tag) Deserialization(bytes []byte) error {
	return json.Unmarshal(bytes, t)
}
