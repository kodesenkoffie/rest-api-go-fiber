package blogs

import (
	"github/kodesenkoffie/server/pkg/crs"
	"github/kodesenkoffie/server/pkg/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateBlog(c *fiber.Ctx) error {
	blog := new(models.Blog)

	if err := c.BodyParser(&blog); err != nil {
		badRequest := crs.BadRequest("Failed to create blog")
		return c.Status(badRequest.Code).JSON(badRequest)
	}

	if blogData, err := createBlog(*blog); err != nil {
		log.Println("Failed to create blog:", err)
		internalServer := crs.InternalServerError("Failed to create blog")
		return c.Status(internalServer.Code).JSON(internalServer)
	} else {
		created := crs.Created("Blog created successfully.", blogData)
		return c.Status(created.Code).JSON(created)
	}
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
