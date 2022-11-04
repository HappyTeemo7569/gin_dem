package define

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var GinEngine *gin.Engine
var GroupServer errgroup.Group
