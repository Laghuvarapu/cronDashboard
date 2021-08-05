package Models

import "gorm.io/gorm"

type Executed struct {
	gorm.Model
	CronId          string      `json:"cron_id"`
	NamespaceName   string      `json:"namespace_name"`
	TimeOfExecution string      `json:"time_of_execution"`
	Status          string      `json:"status"`
	StatusCode      uint        `json:"status_code"`
	Response        interface{} `json:"response"`
	ExecutorId      uint        `json:"executor_id"`
	Cronjob     Cronjob `gorm:"foreignKey:executorId"`
}

func (b *Executed) TableName() string {
	return "executed"
}
