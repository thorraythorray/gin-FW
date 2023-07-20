package op

import "gorm.io/gorm"

type (
	Struct interface{}
)

type BasicApier interface {
	Create(obj Struct) (Struct, error)
	Update(ins, opts Struct) (Struct, error)
	DeleteByID(ins Struct) error
}

type GormApi struct {
	db *gorm.DB
}

func (r *GormApi) Create(obj Struct) (Struct, error) {
	return obj, r.db.Create(&obj).Error
}

func (r *GormApi) Update(ins, opts Struct) (Struct, error) {
	return ins, r.db.Model(&ins).Updates(opts).Error
}

func (r *GormApi) Delete(ins Struct) error {
	return r.db.Delete(&ins).Error
}
