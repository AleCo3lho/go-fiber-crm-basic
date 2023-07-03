package lead

import (
	"go-fiber-crm-basic/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	err := db.Find(&leads).Error
	c.JSON(leads)
	return err
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	err := db.Find(&lead, id).Error
	c.JSON(lead)
	return err
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(500).SendString(err.Error())
		return err
	}
	err := db.Create(&lead).Error
	c.JSON(lead)
	return err
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		err := c.Status(500).SendString("No lead found with ID")
		return err
	}
	err := db.Delete(&lead).Error
	c.JSON(lead)
	return err
}
