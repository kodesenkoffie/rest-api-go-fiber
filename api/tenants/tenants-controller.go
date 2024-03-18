package tenants

import (
	"fmt"
	"github/kodesenkoffie/server/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTenant(c *fiber.Ctx) error {

	tenant := new(models.Tenant)

	if err := c.BodyParser(&tenant); err != nil {
		return err
	}

	fmt.Println(tenant)

	return c.Status(fiber.StatusCreated).JSON(tenant)
}

func FetchTenants(c *fiber.Ctx) error {

	id := c.Params("id")

	if id != "" {
		return c.SendString(id)
	} else {
		return c.SendString("Id not passed")
	}

}

func UpdateTenant(c *fiber.Ctx) error {
	return c.SendString("UpdateTenant")
}

func DeleteTenant(c *fiber.Ctx) error {
	// Logic for deleting a tenant
	return c.SendString("DeleteTenant")
}
