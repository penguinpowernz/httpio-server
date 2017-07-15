package remoteput

import (
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func outputFromContext(outputs Outputs, c *gin.Context) (*Output, bool) {
	idx, err := strconv.Atoi(c.Param("idx"))
	if err != nil {
		c.AbortWithError(400, err)
		return nil, false
	}

	o, ok := outputs[idx]

	if !ok {
		c.AbortWithStatus(404)
		return nil, false
	}

	return o, true
}

func NewAPI(outputs Outputs) *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	
	r.GET("/outputs", func(c *gin.Context) {
		c.JSON(200, outputs)
	})

	r.DELETE("/outputs", func(c *gin.Context) {
		outputs.AllOff()
		c.JSON(200, outputs)
	})

	r.PUT("/outputs", func(c *gin.Context) {
		outputs.AllOn()
		c.JSON(200, outputs)
	})

	r.GET("/outputs/:idx", func(c *gin.Context) {
		o, ok := outputFromContext(outputs, c)
		if !ok {
			return
		}

		c.JSON(200, o)
	})

	r.PUT("/outputs/:idx", func(c *gin.Context) {
		o, ok := outputFromContext(outputs, c)
		if !ok {
			return
		}

		o.SetPosition(1)
		c.JSON(200, o)
	})

	r.DELETE("/outputs/:idx", func(c *gin.Context) {
		o, ok := outputFromContext(outputs, c)
		if !ok {
			return
		}

		o.SetPosition(0)
		c.JSON(200, o)
	})

	return r
}
