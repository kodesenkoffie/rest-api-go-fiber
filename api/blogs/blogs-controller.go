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

	response := fiber.Map{
		"code":      200,
		"info":      "Ok",
		"status":    true,
		"timestamp": time.Now(),
	}

	id := c.Params("id")

	if id != "" {
		data := fetchBlogById(id)
		response["message"] = "Blog fetched successfully."
		response["data"] = data
		return c.Status(fiber.StatusOK).JSON(response)
	} else {
		data := fetchAllBlogs()
		response["message"] = "Blogs fetched successfully."
		response["data"] = data
		return c.Status(fiber.StatusOK).JSON(response)
	}

}

func UpdateBlog(c *fiber.Ctx) error {
	return c.SendString("UpdateBlog")
}

func DeleteBlog(c *fiber.Ctx) error {
	return c.SendString("DeleteBlog")
}
