package container

import (
	getStatusHandler "uala-challenge/internal/domain/get_status/handler"

	"github.com/gin-gonic/gin"
)

type container struct {
	GetStatusHandler gin.HandlerFunc
}

func LoadContainer() *container {

	getStatusHandler := getStatusHandler.NewGetStatusHandler()

	return &container{
		GetStatusHandler: getStatusHandler.Handle,
	}
}
