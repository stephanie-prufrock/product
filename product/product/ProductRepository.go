package product

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/pmoule/go2hal/hal"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (self *Product) Get(c echo.Context) error {
	repository := ProductRepository()
	entities := repository.GetList()

	// Make hateoas Resource
	rootHref := fmt.Sprintf("%s%s", c.Request().Host, c.Path())

	root := hal.NewResourceObject()
	link := &hal.LinkObject{Href: rootHref}

	selfRel := hal.NewSelfLinkRelation()
	selfRel.SetLink(link)

	root.AddLink(selfRel)

	var embeddedActors []hal.Resource
	for _, entity := range entities {
		// entity link
		href := fmt.Sprintf("%s/%d", rootHref, entity.Id)
		selfLink, _ := hal.NewLinkObject(href)

		entitySelf, _ := hal.NewLinkRelation("self")
		entitySelf.SetLink(selfLink)

		embeddedActor := hal.NewResourceObject()
		embeddedActor.AddLink(entitySelf)
		embeddedActor.AddData(entity)
		embeddedActors = append(embeddedActors, embeddedActor)
	}
	embeddedRel, _ := hal.NewResourceRelation("products")
	embeddedRel.SetResources(embeddedActors)
	root.AddResource(embeddedRel)

	encoder := hal.NewEncoder()
	jsonBytes, _ := encoder.ToJSON(root)

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (self *Product) FindById(c echo.Context) error{
	repository := ProductRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self, err := repository.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, err)
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Make hateoas Resource
	href := fmt.Sprintf("%s%s", c.Request().Host, c.Request().URL)
	selfLink, _ := hal.NewLinkObject(href)

	entitySelf, _ := hal.NewLinkRelation("self")
	entitySelf.SetLink(selfLink)

	embeddedActor := hal.NewResourceObject()
	embeddedActor.AddLink(entitySelf)
	embeddedActor.AddData(self)

	encoder := hal.NewEncoder()
	jsonBytes, _ := encoder.ToJSON(embeddedActor)

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (self *Product) Persist(c echo.Context) error{
	repository := ProductRepository()
	params := make(map[string]string)
	var err error
	c.Bind(&params)
	err = ObjectMapping(self, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	self.onPrePersist()
	err = repository.save(self)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	self.onPostPersist()

	rootHref := fmt.Sprintf("%s%s", c.Request().Host, c.Path())
	href := fmt.Sprintf("%s/%d", rootHref, self.Id)
	selfLink, _ := hal.NewLinkObject(href)

	entitySelf, _ := hal.NewLinkRelation("self")
	entitySelf.SetLink(selfLink)

	embeddedActor := hal.NewResourceObject()
	embeddedActor.AddLink(entitySelf)
	embeddedActor.AddData(self)

	encoder := hal.NewEncoder()
	jsonBytes, _ := encoder.ToJSON(embeddedActor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSONBlob(http.StatusOK, jsonBytes)
	}
}

func (self *Product) Put(c echo.Context) error{
	repository := ProductRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string]string)

	c.Bind(&params)
	self.onPreUpdate()
	self, err := repository.Update(id, params)
	self.onPostUpdate()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		href := fmt.Sprintf("%s%s", c.Request().Host, c.Request().URL)
		selfLink, _ := hal.NewLinkObject(href)

		entitySelf, _ := hal.NewLinkRelation("self")
		entitySelf.SetLink(selfLink)

		embeddedActor := hal.NewResourceObject()
		embeddedActor.AddLink(entitySelf)
		embeddedActor.AddData(self)

		encoder := hal.NewEncoder()
		jsonBytes, _ := encoder.ToJSON(embeddedActor)

		return c.JSONBlob(http.StatusOK, jsonBytes)
	}
}

func (self *Product) Remove(c echo.Context) error{
	repository := ProductRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self, err := repository.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, err)
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	self.onPreRemove()
	err = repository.Delete(self)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	self.onPostRemove()
	return c.JSON(http.StatusOK, err)
}