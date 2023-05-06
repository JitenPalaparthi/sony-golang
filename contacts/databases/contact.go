package database

import (
	"contacts/models"

	"gorm.io/gorm"
)

type Contact struct {
	DB *gorm.DB
}

type IContact interface {
	Create(contact *models.Contact) error
	Get(id string) (contact *models.Contact, error error)
}

func (c *Contact) Create(contact *models.Contact) error {
	c.DB.AutoMigrate(&models.Contact{})
	tx := c.DB.Create(contact)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (c *Contact) Get(id string) (contact *models.Contact, error error) {
	contact = new(models.Contact)
	tx := c.DB.First(&contact, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}
