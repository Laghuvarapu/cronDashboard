package Routes

import (
	"github.com/gin-gonic/gin"
	"recursion/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	cron := r.Group("/cronjobs")
	{
		cron.POST(":id/createCron ", Controllers.CreateCronjob)
		cron.GET(":id/getAllCronJobs", Controllers.GetAllCronJobs)
		cron.PATCH(":id/updateCronJob/:namespaceName/:name", Controllers.UpdateCronjob)
		cron.GET(":id/getByName/:namespaceName/:name", Controllers.GetCronByUrlName)
		cron.DELETE(":id/deleteCronJob/:namespaceName/:name", Controllers.DeleteCronjob)
	}
	exec := r.Group("/executed")
	{
		exec.GET(":id/:namespaceName/:name", Controllers.GetPastNExecutions)
	}

	return r
}
