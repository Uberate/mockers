package o_v2

import (
	"mockers/pkg/errors"
)

var (
	ObjectErrorNamespace = "object-error"
	ObjectError          = errors.NewErrMessage()

	NotFoundObjectErr = ObjectError.NewError("NotFoundObjectErr", ObjectErrorNamespace, errors.RegionZHCN, "Not fount specify object: [%s]")
)
