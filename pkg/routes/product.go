package routes

func ProductRoutes(app *fiber.App) {
	productGroup := app.Group("/products")
	productGroup.Get("/", product.List)
	productGroup.Get("/:id", product.Get)
	productGroup.Post("/", product.Create)
	productGroup.Post("/:id", product.Update)
	productGroup.Delete("/:id", product.Delete)
}
