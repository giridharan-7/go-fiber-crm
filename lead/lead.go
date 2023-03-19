package lead

import (
	"github.com/giridharan-7/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)

	/* This line queries the database using the Find() method of the
	GORM database object db. The first argument to Find() is a pointer
	to the destination object &lead, which tells GORM to retrieve the data
	from the leads table and store it in the lead variable. The second argument
	id is the value of the id parameter retrieved earlier, and specifies the
	primary key value to search for in the leads table. */

	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	// lead := new(Lead) creates a new pointer to a Lead struct and
	// initializes it to the zero value of the Lead struct
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {

		/* The c.BodyParser() function in this context takes the request body,
		which is expected to be in JSON format, and unmarshals it into the lead variable. */

		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	/*The difference between the two is that lead := new(Lead) returns a pointer
	to the Lead struct, while var lead Lead returns the Lead struct itself. */

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with id")
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}
