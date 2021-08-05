package Models

import (
	"recursion/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetPastNExecutions(executed *[]Executed, jName string, nsName string, n int) (err error) {
	if err := Config.DB.Order("time_of_execution desc").Limit(n).Where("job_name = ? AND namespace_name =? " ,jName, nsName).Find(&executed).Error; err != nil {
		return err
	}
	return nil
}
