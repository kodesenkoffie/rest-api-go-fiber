package routes

import (
	"fmt"
	"time"

	"github/kodesenkoffie/server/api/blogs"
	"github/kodesenkoffie/server/api/tenants"
	"github/kodesenkoffie/server/internal/config/env"

	"github.com/gofiber/fiber/v2"
)

func endPoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", env.API_PREFIX, endpoint)
}

func SetupRoutes(app *fiber.App) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"code":      200,
			"status":    true,
			"info":      "OK",
			"timestamp": time.Now(),
			"message":   "Server is running",
		})
	})

	// Blogs
	app.Post(endPoint("blogs"), blogs.CreateBlog)
	app.Get(endPoint("blogs/:id?"), blogs.FetchBlogs)
	app.Put(endPoint("blogs/:id"), blogs.UpdateBlog)
	app.Delete(endPoint("blogs/:id"), blogs.DeleteBlog)

	// Tenants
	app.Post(endPoint("tenants"), tenants.CreateTenant)
	app.Get(endPoint("tenants/:id?"), tenants.FetchTenants)
	app.Put(endPoint("tenants/:id"), tenants.UpdateTenant)
	app.Delete(endPoint("tenants/:id"), tenants.DeleteTenant)

}
