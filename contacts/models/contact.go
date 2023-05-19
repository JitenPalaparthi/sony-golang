package models

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Status       string `json:"status" gorm:"default:active"`
	LastModified int64  `json:"lastModified" gorm:"autoCreateTime;column:lastModified"` //last_modified
}

func (c *Contact) Validate() error {
	if c.Name == "" {

		//go:cover ignoredefault
		return fmt.Errorf("filed Name cannot be empty")
	}
	if c.Email == "" {

		//go:cover ignoredefault
		return fmt.Errorf("filed Email cannot be empty")
	}

	//go:cover ignoredefault
	return nil
}

func (c *Contact) ToJson() (string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	//go:cover ignore
	return string(bytes), nil
}
