package Controllers

import (
	"fmt"
	"net/http"

	"recursion/Models"

	"github.com/gin-gonic/gin"
)

func CreateCronjob(c *gin.Context) {
	var cronjob Models.Cronjob
	var cron Models.Cronjob
	c.BindJSON(&cronjob)
	//for checking role
	var roles Models.Roles
	var roleAccess []Models.RolesAccess
	err := Models.CheckRoles(&roles, cronjob.UserId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)

	}
	//checking access
	if roles.RolesID == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Access not there to create the cronjob"})
		return
	}
	//getting namespaces
	err = Models.Checknamespace(&roleAccess, roles.RoleName)
	flag := 1
	for _, namespa := range roleAccess {
		if cronjob.NamespaceName == namespa.NamespaceName {
			flag = 0
			break
		}
	}
	if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Provided Namespace is not matching with the registered one ",
		})
		return
	} else {
		if roles.RoleAccess == 0 {
			cronjob.Enabled = false
		}
		if roles.RoleAccess == 1 {
			cronjob.Enabled = true
		}

	}

	//for checking cron name
	err = Models.CheckName(&cron, cronjob.JobName, cronjob.NamespaceName)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Name already exists,try with a different name",
		})
		return
	}

	//creating cron job
	err = Models.CreateCronjob(&cronjob)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		c.JSON(http.StatusOK, cronjob)
	}

}

func GetAllCronJobs(c *gin.Context) {
	userid := c.Params.ByName("id")
	var cronjob []Models.Cronjob

	//checking in which namespace user is present
	var roles Models.Roles
	var roleAccess []Models.RolesAccess
	err := Models.CheckRoles(&roles, userid)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	err = Models.CheckNamespace(&roleAccess, roles.RoleName)
	//for creating an array of namespaces in which user is present
	var nsArray []string
	for _, namespace := range roleAccess {
		nsArray = append(nsArray, namespace.NamespaceName)
	}

	//getting all cron jobs
	err = Models.GetAllCronJobs(&cronjob, nsArray)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cronjob)
	}
}

func UpdateCronjob(c *gin.Context) {
	var cronjob Models.Cronjob
	userid := c.Params.ByName("id")
	nsName := c.Params.ByName("namespaceName")

	jName := c.PostForm("job_name")
	//checking role
	var roles Models.Roles
	var roleAccess []Models.RolesAccess
	err := Models.CheckRoles(&roles, userid)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	//getting namespaces
	err = Models.CheckNamespace(&roleAccess, roles.RoleName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	//Checking user have access or not
	if roles.RolesID == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Access not there to update the status"})
		return
	}
	//user is present in provided namespace or not
	flag := 1
	for _, namespa := range roleAccess {
		if cronjob.NamespaceName == namespa.NamespaceName {
			flag = 0
		}
	}
	if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Provided Namespace is not matching with the registered one ",
		})
		return
	}
	//Updating the cronjob
	err = Models.GetCronjobByID(&cronjob, jName, nsName)

	if err != nil {
		c.JSON(http.StatusNotFound, cronjob)
		return
	}
	c.BindJSON(&cronjob)

	err = Models.UpdateCronjob(&cronjob)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cronjob)
	}
}

func GetCronByUrlName(c *gin.Context) {
	var cronjob Models.Cronjob

	namespaceName := c.Params.ByName("namespaceName")
	name := c.Params.ByName("name")

	//fetching the data
	err := Models.GetCronByUrlName(&cronjob, name, namespaceName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cronjob)
	}

}

func DeleteCronjob(c *gin.Context) {

	var cronjob Models.Cronjob
	userId := c.Params.ByName("id")
	jName := c.Params.ByName("name")
	nsName := c.Params.ByName("namespaceName")
	//checking if the user have access to delete the cronjob or not
	var roles Models.Roles
	var roleAccess []Models.RolesAccess
	err := Models.CheckRoles(&roles, userId)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	//Checking user have access or not
	if roles.RolesID == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Access not there to update the status"})
	}
	//getting namespaces in which user is present
	err = Models.Checknamespace(&roleAccess, roles.RoleName)
	//checking namespaces in which user have access to provided namespace or not
	flag := 1
	for _, namespace := range roleAccess {
		if cronjob.NamespaceName == namespace.NamespaceName {
			flag = 0
		}
	}
	if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Provided Namespace is not matching with the registered one ",
		})
		return
	}
	//deleting the cronjob
	err = Models.DeleteCronjob(&cronjob, jName, nsName)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"name" + nsName: "is deleted"})
	}

}
