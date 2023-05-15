package application

import (
	"lab6/orders"
	"lab6/products"
	"lab6/users"
	"mime/multipart"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const imageRoute string = "/images/products/"
const productImageDirectory string = "resources/products/"

type AdminController struct {
	productRepo  *products.Repository
	orderService *orders.OrderService
	userService  *users.UserService
}

func (controller *AdminController) Initialize(engine *gin.Engine, productRepo *products.Repository, orderService *orders.OrderService, userService *users.UserService) {
	controller.productRepo = productRepo
	controller.orderService = orderService
	controller.userService = userService
	engine.GET("/admin/products", controller.indexRoute)
	engine.GET("/admin/products/delete/:id", controller.deleteProductRoute)
	engine.GET("/admin/products/create", controller.createProductFormRoute)
	engine.POST("/admin/products/create", controller.createProductRoute)
	engine.GET("/admin/products/edit/:id", controller.editProductFormRoute)
	engine.POST("/admin/products/edit/:id", controller.editProductRoute)
	engine.GET("/admin/orders", controller.ordersRoute)
	engine.GET("/admin/users", controller.usersRoute)
}

func (controller *AdminController) indexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func (controller *AdminController) productsRoute(c *gin.Context) {
	products, err := controller.productRepo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get products"})
	} else {
		c.HTML(http.StatusOK, "admin/admin-products.tmpl", gin.H{"products": products})
	}
}

func (controller *AdminController) deleteProductRoute(c *gin.Context) {
	id := c.Param("id")

	err := controller.productRepo.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete item with such id"})
	} else {
		c.Redirect(http.StatusPermanentRedirect, "/admin/products")
	}
}

func (controller *AdminController) editProductFormRoute(c *gin.Context) {
	id := c.Param("id")

	product, err := controller.productRepo.GetById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get item with such id"})
	} else {
		c.HTML(http.StatusOK, "admin/edit.tmpl", gin.H{"product": product})
	}
}

func (controller *AdminController) createProductFormRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/create.tmpl", gin.H{})
}

type ProductRequestDto struct {
	Name        string                `form:"name"`
	Price       int                   `form:"price"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
}

func (controller *AdminController) createProductRoute(c *gin.Context) {
	var productDto ProductRequestDto
	err := c.ShouldBind(&productDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, imageUrl, err := saveImageByRandomName(c, productDto.Image)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	product := products.Product{
		Id:          "",
		Name:        productDto.Name,
		Description: productDto.Description,
		Price:       productDto.Price,
		ImageUrl:    imageUrl,
	}

	controller.productRepo.Create(&product)
}

func (controller *AdminController) editProductRoute(c *gin.Context) {
	id := c.Param("id")
	var productDto ProductRequestDto

	err := c.ShouldBind(&productDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, imageUrl, err := saveImageByRandomName(c, productDto.Image)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	product := products.Product{
		Id:          id,
		Name:        productDto.Name,
		Description: productDto.Description,
		Price:       productDto.Price,
		ImageUrl:    imageUrl,
	}

	err = controller.productRepo.Update(&product)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func saveImageByRandomName(c *gin.Context, image *multipart.FileHeader) (string, string, error) {
	randomString := uuid.New().String()
	extension := path.Ext(image.Filename)
	imageName := randomString + extension
	err := c.SaveUploadedFile(image, productImageDirectory+imageName)
	imageUrl := imageRoute + imageName
	return imageName, imageUrl, err
}

func (controller *AdminController) ordersRoute(c *gin.Context) {
	orders, err := controller.orderService.GetAllOrders()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get products"})
	} else {
		c.HTML(http.StatusOK, "admin/admin-orders.tmpl", gin.H{"orders": orders})
	}
}

func (controller *AdminController) usersRoute(c *gin.Context) {
	users, err := controller.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get products"})
	} else {
		c.HTML(http.StatusOK, "admin/admin-users.tmpl", gin.H{"users": users})
	}
}
