package index

import "mockers/pkg/hash"

type Indexable interface {
	// Equal return true when 'a' equal 'b'. Else return false.
	Equal(a, b interface{}) bool
	// Less return true when 'a' < 'b'. Else return false.
	Less(a, b interface{}) bool
	// Greater return true when 'a' > 'b'. Else return false.
	Greater(a, b interface{}) bool

	hash.Hashable

	// Name is the indexable function name. The name with version is unique in the application. We will find the
	// function by name and version in function center. And when all the Indexable can serializable and deserializable.
	// In other worlds, the Indexable should implement the Hashable interface. About more info of the Name, see GetFunc.
	Name() (name, version string)
}
