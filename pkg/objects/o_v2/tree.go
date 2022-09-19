package o_v2

import (
	"fmt"
	"mockers/pkg/utils"
	"strconv"
	"strings"
	"sync"
)

const PathSeparator = "/"

// Tree is node or root of a tree, it save subTree and all this node value.
// The save path of '/a','/a/','a' is same path /a，and node name is 'a'.
// If the path is '/a/b', root node name is 'a', and one of sub node is named 'b'.
// The character '/' see as split.
type Tree struct {
	// TreeName is the name of this node
	TreeName string

	// SubTree is the sub node of this node. Key is path name, the name can't be empty, if empty is mean this node.
	SubTree map[string]*Tree

	// Save all datas here.
	Datas []NodeObject

	// inner locker for write and read tree info
	locker sync.RWMutex

	idFunction func(interface{}, int) string
}

// Prune will delete all empty none-data node.
func (t *Tree) Prune() error {
	t.locker.Lock()
	defer t.locker.Unlock()

	utils.Logger.Warnf("Do prune for node: %s", t.TreeName)

	// remove: sube tree
	for _, item := range t.SubTree {
		if err := item.Prune(); err != nil {
			return err
		}
	}

	// remove: data is emtpy, subtree is nil
	for key, item := range t.SubTree {
		if len(item.Datas) == 0 && len(item.SubTree) == 0 {
			delete(t.SubTree, key)
		}
	}

	return nil
}

// GetNode return the path of node.
func (t *Tree) GetNode(path string) (*Tree, error) {
	paths := strings.Split(path, PathSeparator)
	var newPath []string
	// remove empty path
	for _, item := range paths {
		if len(item) != 0 {
			newPath = append(newPath, item)
		}
	}

	node := t.completeNode(newPath...)
	if node == nil {
		err := NotFoundObjectErr.Param("Node in path: " + path)
		utils.Logger.Errorf("%v", err.Error())
		return nil, err
	}
	return node, nil
}

// completeNode will complete the path node, and return last node.
func (t *Tree) completeNode(path ...string) *Tree {
	t.locker.Lock()
	defer t.locker.Unlock()

	if len(path) == 0 {
		return t
	}

	currentName := path[0]

	if treeNode, ok := t.SubTree[currentName]; ok {
		if len(path) == 1 {
			return treeNode
		}
		return treeNode.completeNode(path[1:]...)
	}

	newTreeNode := &Tree{
		TreeName:   currentName,
		idFunction: DefaultIdFunctionGenerator(),
	}

	utils.Logger.Infof("Create new node: %s", currentName)
	if t.SubTree == nil {
		t.SubTree = map[string]*Tree{}
	}
	t.SubTree[currentName] = newTreeNode
	if len(path) == 1 {
		return newTreeNode
	}
	return newTreeNode.completeNode(path[1:]...)
}

// ---------------------------------------------------------------------------------------------------------------------
// node operator

//--------------------------------------------------
// add node data

// AddNodeData will add an obj to node, in tree, the obj will be package to the other type of data (NodeObject).
// If id same, it will cover old value.
func (t *Tree) AddNodeData(obj interface{}) (NodeMetadata, error) {

	// create new node object
	nd, err := NewNodeObject(t.idFunction, obj, len(t.Datas))
	if err != nil {
		return NodeMetadata{}, err
	}

	return t.AddNodeObject(nd)
}

// AddNodeDataToPath will add and object to specify node, the object will be package to other type of data(NodeObject).
// If id
func (t *Tree) AddNodeDataToPath(obj interface{}, path string) (NodeMetadata, error) {
	node, err := t.GetNode(path)
	if err != nil {
		return NodeMetadata{}, err
	}
	return node.AddNodeData(obj)
}

// AddNodeObject will append an object to tree, if specify object's id is already in datas, old data will be cover.
func (t *Tree) AddNodeObject(object NodeObject) (NodeMetadata, error) {

	t.locker.Lock()
	defer t.locker.Unlock()

	for index, item := range t.Datas {
		if item.Metadata.Id == object.Metadata.Id {
			// if found same id, cover it and return nil directly.
			t.Datas[index].Metadata.Id = object.Metadata.Id
			return object.Metadata, nil
		}
	}

	// not found, append directly
	t.Datas = append(t.Datas, object)
	return object.Metadata, nil
}

// AddNodeObjectToPath will append a NodeObject to specify path.
func (t *Tree) AddNodeObjectToPath(object NodeObject, path string) (NodeMetadata, error) {
	node, err := t.GetNode(path)
	if err != nil {
		return NodeMetadata{}, err
	}
	return node.AddNodeObject(object)
}

// ListNodeObjects return all data
func (t *Tree) ListNodeObjects() []NodeObject {
	t.locker.RLocker()
	defer t.locker.RUnlock()

	value := make([]NodeObject, 0, len(t.Datas))
	copy(value, t.Datas)

	return value
}

//--------------------------------------------------
// list node data

// ListNodeObjectsFromPath return the objects from specify path.
func (t *Tree) ListNodeObjectsFromPath(path string) ([]NodeObject, error) {
	node, err := t.GetNode(path)
	if err != nil {
		return nil, err
	}

	return node.ListNodeObjects(), nil
}

// GetNodeData can read object by id and write to obj, so, the obj must a pointer.
func (t *Tree) GetNodeData(id string, obj interface{}) (NodeMetadata, error) {
	value, err := t.GetNodeObject(id)
	if err != nil {
		return NodeMetadata{}, err
	}
	return value.Metadata, value.Data.To(obj)
}

func (t *Tree) GetNodeDataFromPath(id string, obj interface{}, path string) (NodeMetadata, error) {
	node, err := t.GetNode(path)
	if err != nil {
		return NodeMetadata{}, err
	}

	return node.GetNodeData(id, obj)
}

func (t *Tree) GetNodeObject(id string) (NodeObject, error) {
	t.locker.RLock()
	defer t.locker.RUnlock()

	for _, item := range t.Datas {
		if item.Metadata.Id == id {
			return item, nil
		}
	}
	// not found record
	return NodeObject{}, fmt.Errorf("not found data by id: %s", id)
}

func (t *Tree) GetNodeObjectFromPath(id string, path string) (NodeObject, error) {
	node, err := t.GetNode(path)
	if err != nil {
		return NodeObject{}, err
	}

	return node.GetNodeObject(id)
}

//--------------------------------------------------
// delete node data

func (t *Tree) DeleteNodeData(id ...string) (int, error) {
	deleteCache := map[string]struct{}{}
	deleteCounter := 0

	var res []NodeObject

	for _, item := range id {
		deleteCache[item] = struct{}{}
	}

	for _, item := range t.Datas {
		if _, ok := deleteCache[item.Metadata.Id]; ok {
			deleteCounter++
			continue
		}
		res = append(res, item)
	}

	t.Datas = res
	return deleteCounter, nil
}

func (t *Tree) DeleteNodeDataFromPath(path string, id ...string) (int, error) {
	n, err := t.GetNode(path)
	if err != nil {
		return 0, err
	}
	return n.DeleteNodeData(id...)
}

// DefaultIdFunctionGenerator return a function automatically add station of id.
func DefaultIdFunctionGenerator() func(interface{}, int) string {
	counter := 0
	locker := sync.Mutex{}
	return func(interface{}, int) string {
		locker.Lock()
		defer locker.Unlock()
		counter++
		return strconv.Itoa(counter)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// options & funcs

//--------------------------------------------------
// for node object

type NodeObjectOption func(object *NodeObject)

// NodeObjectCreateBy return a NodeObjectOption to set object create by info.
func NodeObjectCreateBy(createBy string) NodeObjectOption {
	return func(object *NodeObject) {
		if object != nil {
			object.Metadata.CreateBy = createBy
		}
	}
}

// NodeObjectUpdateBy return a NodeObjectOption to set object update by info.
func NodeObjectUpdateBy(updateBy string) NodeObjectOption {
	return func(object *NodeObject) {
		if object != nil {
			object.Metadata.UpdateBy = updateBy
		}
	}
}

// NodeObjectDeleteBy return a NodeObjectOption to set object delete by info.
func NodeObjectDeleteBy(deleteBy string) NodeObjectOption {
	return func(object *NodeObject) {
		if object != nil {
			object.Metadata.DeleteBy = deleteBy
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// public function for quick create tree or node

// NewNodeObject will create new node object, used NodeObjectOption to set some info. But the option can't change id
// info, the id will be cover by DefaultIdFunctionGenerator().
func NewNodeObject(idFunction func(obj interface{}, count int) string, obj interface{}, length int, options ...NodeObjectOption) (NodeObject, error) {
	if idFunction == nil {
		idFunction = DefaultIdFunctionGenerator()
	}

	no := NodeObject{}
	for _, option := range options {
		option(&no)
	}

	no.Metadata.Id = idFunction(obj, length)

	data, err := NewNodeData(obj)
	if err != nil {
		return NodeObject{}, err
	}
	no.Data = data
	return no, nil
}
