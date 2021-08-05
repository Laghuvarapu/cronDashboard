package Models

import "gorm.io/gorm"

type Cronjob struct {
	gorm.Model
	JobName       string      `json:"job_name"`
	Url           string      `json:"url"`
	Method        string      `json:"quantity"`
	Payload       interface{} `json:"payload"`
	NoOfRetries   uint        `json:"no_of_retries"`
	TimeOut       uint        `json:"time_out"`
	SuspendedFor  string      `json:"suspended_for"`
	CronPattern   string      `json:"cron_pattern"`
	NamespaceName string      `json:"namespace_name"`
	Enabled       bool        `json:"enabled"`
	UserId        string      `json:"user_id"`
}

func (b *Cronjob) TableName() string {
	return "cronjobs"
}
