package product

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"product/external"
)

type Product struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	Name string `json:"name"`
	Stock int `json:"stock"`

}

func (self *Product) onPostPersist() (err error){
	productChanged := NewProductChanged()
	model.Copy(productChanged, self)

	Publish(productChanged)

	return nil
}
func (self *Product) onPrePersist() (err error){ return nil }
func (self *Product) onPreUpdate() (err error){ return nil }
func (self *Product) onPostUpdate() (err error){ return nil }
func (self *Product) onPreRemove() (err error){ return nil }
func (self *Product) onPostRemove() (err error){ return nil }


