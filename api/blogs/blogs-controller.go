package blogs

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateBlog(c *fiber.Ctx) error {

	response := fiber.Map{
		"code":      201,
		"info":      "Created",
		"message":   "Blog created successfully.",
		"status":    true,
		"timestamp": time.Now(),
	}

	// blog := new(models.Blog)

	// if err := c.BodyParser(&blog); err != nil {
	// 	return err
	// }

	// fmt.Println(blog)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func FetchBlogs(c *fiber.Ctx) error {

	id := c.Params("id")

	if id != "" {
		return c.SendString(id)
	} else {
		return c.SendString("Id not passed")
	}

}

func UpdateBlog(c *fiber.Ctx) error {
	return c.SendString("UpdateBlog")
}

func DeleteBlog(c *fiber.Ctx) error {
	return c.SendString("DeleteBlog")
}
