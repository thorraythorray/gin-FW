package gorm

import "gorm.io/gorm"

type Dict map[string]interface{}

type RestApi interface {
	Create(newOpts *Dict) error
	Update(checkOpts, upOpts *Dict) error
	List(filOpts *Dict) error
	Delete(checkOpts *Dict) error
}

type RestApiImpl struct {
	db *gorm.DB
}

func NewRestApiImpl(db *gorm.DB) *RestApiImpl {
	return &RestApiImpl{db: db}
}
