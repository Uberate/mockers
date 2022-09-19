package objects

import "github.com/mitchellh/mapstructure"

// NodeData save the data of node in a map[string]interface{}
type NodeData map[string]interface{}

// To parse the node data, and convert to obj, obj must a point or map.
func (n NodeData) To(obj interface{}) error {
	return mapstructure.Decode(n, obj)
}

// NewNodeData return a data from specify data.
func NewNodeData(obj interface{}) (NodeData, error) {
	r := NodeData{}
	err := mapstructure.Decode(obj, &r)
	return r, err
}

// NodeObject save the object to tree.
type NodeObject struct {
	Metadata NodeMetadata `json:"_metadata"`

	// Data contain al data of object
	Data NodeData `json:"data"`
}

// GetData will set value to input res. So res should a pointer.
func (no NodeObject) GetData(res interface{}) error {
	return no.Data.To(res)
}

// NodeMetadata is inner metadata info of the tree. The business' model should implement id info by itself.
// In common, user of the NodeObject can ignore the NodeMetadata. But, you also can rewrite metadata info. Mocker
// support write NodeObject directly(default logic only need NodeData).
type NodeMetadata struct {
	// the id of the node
	Id string `json:"_id" mapstructure:"_id"`

	CreateAt string `json:"_create_at" mapstructure:"_create_at"`
	CreateBy string `json:"_create_by" mapstructure:"_create_by"`
	UpdateAt string `json:"_update_at" mapstructure:"_update_at"`
	UpdateBy string `json:"_update_by" mapstructure:"_Update_by"`
	DeleteAt string `json:"_delete_at" mapstructure:"_delete_at"`
	DeleteBy string `json:"_delete_by" mapstructure:"_delete_by"`
}
