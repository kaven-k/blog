package es_service

import (
	"server/models"
)

type Option struct {
	models.PageInfo
	Fields []string `form:"fields" json:"fields"`
	Tag    string   `form:"tag" json:"tag"`
}

// GetForm 生效于原值，这里就要用指针
func (o *Option) GetForm() int {
	if o.Page <= 0 {
		o.Page = 1
	}
	if o.Limit <= 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit

}
