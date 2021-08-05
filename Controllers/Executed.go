package Controllers

import (
	"net/http"
	"strconv"

	"recursion/Models"

	"github.com/gin-gonic/gin"
)

func GetPastNExecutions(c *gin.Context) {

	jName := c.Params.ByName("job_name")
	nsName := c.Params.ByName("namespace_name")
	n := c.Params.ByName("n")
	var executed []Models.Executed

	m, err := strconv.Atoi(n)
	err = Models.GetPastNExecutions(&executed, jName, nsName, m)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, executed)
	}
}
