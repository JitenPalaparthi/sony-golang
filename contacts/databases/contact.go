package database

import (
	"contacts/models"
	"fmt"

	"gorm.io/gorm"
)

type Contact struct {
	DB *gorm.DB
}

type IContact interface {
	Create(contact *models.Contact) (*models.Contact, error)
	Get(id string) (contact *models.Contact, error error)
}

func (c *Contact) Create(contact *models.Contact) (*models.Contact, error) {
	c.DB.AutoMigrate(&models.Contact{})
	tx := c.DB.Create(contact)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return c.Get(fmt.Sprint(contact.ID))
}

func (c *Contact) Get(id string) (contact *models.Contact, errerroror error) {
	contact = new(models.Contact)
	tx := c.DB.First(&contact, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}
