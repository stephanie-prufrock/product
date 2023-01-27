package product
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type ProductDB struct{
	db *gorm.DB
}

var productrepository *ProductDB

func ProductDBInit() {
	var err error
	productrepository = &ProductDB{}
	productrepository.db, err = gorm.Open(sqlite.Open("Product_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	productrepository.db.AutoMigrate(&Product{})

}

func ProductRepository() *ProductDB {
	return productrepository
}

func (self *ProductDB)save(entity interface{}) error {
	
	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *ProductDB)GetList() []Product{
	
	entities := []Product{}
	self.db.Find(&entities)

	return entities
}

func (self *ProductDB)FindById(id int) (*Product, error){
	entity := &Product{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		return entity, nil
	}
}

func (self *ProductDB) Delete(entity *Product) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *ProductDB) Update(id int, params map[string]string) (*Product, error){
	entity := &Product{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &Product{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}