package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"recursion/Config"
)

func CreateCronjob(cronjob *Cronjob) (err error) {
	if err = Config.DB.Create(cronjob).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCronJobs(cronjob *[]Cronjob, nsArray []string) (err error) {

	if err = Config.DB.Where("namespace_name IN ?", nsArray).Find(&cronjob).Error; err != nil {
		return err
	}

	return nil
}

func UpdateCronjob(cronjob *Cronjob) (err error) {

	Config.DB.Save(cronjob)

	return nil
}

func GetCronjobByID(cronjob *Cronjob, jName string, namespaceName string) (err error) {
	if err = Config.DB.Where("job_name = ? AND namespace_name =?", jName, namespaceName).First(&cronjob).Error; err != nil {
		return err
	}
	return nil
}

func CheckName(cronjob *Cronjob, name string, namespace_name string) (err error) {
	if err = Config.DB.Where("name = ? AND namespace_name = ?", name, namespace_name).First(&cronjob).Error; err != nil {
		return err
	}
	return nil
}

func GetCronByUrlName(cronjob *Cronjob, name string, namespaceName string) (err error) {
	if err = Config.DB.Where("url = ? OR cron_name =?  AND namespace_name=?", name, namespaceName).Or("cron_name = ? AND namespace_name= ?", name, namespaceName).First(&cronjob).Error; err != nil {
		return err
	}
	return nil
}

func CheckNamespace(roleAccess *[]RolesAccess, roleName string) (err error) {
	if err := Config.DB.Where("role_name", roleName).First(&roleAccess).Error; err != nil {
		return err
	}
	return nil

}

func DeleteCronjob(cronjob *Cronjob, name string, nsName string) (err error) {
	Config.DB.Where("job_name = ? AND namespace_name =? ", name, nsName).Delete(cronjob)
	return nil
}

func CheckRoles(roles *Roles, Id string) (err error) {
	if err := Config.DB.Where("user_id", Id).First(&roles).Error; err != nil {
		return err
	}
	return nil

}
