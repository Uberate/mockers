package web

import (
	"encoding/json"
	"io/ioutil"
	"mockers/internal/web/server"
	"mockers/pkg/objects/o_v2"
	"net/http"

	"github.com/gin-gonic/gin"
)

var objectInstance = o_v2.Tree{}

func Handlers() []server.GinHandler {
	res := []server.GinHandler{
		{
			HandlerPath: "/objects/node/*path",
			Method:      []string{http.MethodGet},
			Handler: []gin.HandlerFunc{
				GetNode,
			},
		},
		//--------------------------------------------------
		//datas
		{
			HandlerPath: "/objects/datas/*path",
			Method:      []string{http.MethodPost},
			Handler: []gin.HandlerFunc{
				PostData,
			},
		},
		{
			HandlerPath: "/objects/data/id/:id/*path",
			Method:      []string{http.MethodGet},
			Handler: []gin.HandlerFunc{
				GetData,
			},
		},
	}

	return res
}

func GetNode(ctx *gin.Context) {
	path := ctx.Param("path")
	obj, err := objectInstance.GetNode(path)
	if err != nil {
		ctx.JSONP(http.StatusInternalServerError, err)
	}
	ctx.JSONP(http.StatusOK, obj)
}

func ListNode(ctx *gin.Context) {

}

func GetData(ctx *gin.Context) {
	path := ctx.Param("path")
	id := ctx.Param("id")

	var res map[string]interface{}
	node, err := objectInstance.GetNode(path)
	if err != nil {
		server.AutoRes(ctx, 0, nil, err)
		return
	}

	_, err = node.GetNodeData(id, &res)
	server.AutoRes(ctx, 0, res, err)
}

func PostData(ctx *gin.Context) {
	path := ctx.Param("path")

	body := map[string]interface{}{}
	var value []byte
	var err error
	if value, err = ioutil.ReadAll(ctx.Request.Body); err != nil {
		ctx.JSONP(http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(value, &body); err != nil {
		ctx.JSONP(http.StatusInternalServerError, err)
		return
	}

	var metadata o_v2.NodeMetadata
	if metadata, err = objectInstance.AddNodeDataToPath(body, path); err != nil {
		ctx.JSONP(http.StatusInternalServerError, err)
	} else {
		ctx.JSONP(200, metadata)
	}
}

//func UpdateData(ctx *gin.Context) {
//	path := ctx.Param("path")
//	id := ctx.Param("id")
//    obj, err := objectInstance.GetNode(path)
//	if err != nil {
//		errors.
//	}
//}
