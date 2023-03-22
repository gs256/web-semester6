package application

import (
	"lab5/product_repository"
	"lab5/products"
	"mime/multipart"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const imageRoute string = "/images/products/"
const productImageDirectory string = "resources/products/"

type AdminController struct {
	repo *product_repository.Repository
}

func (controller *AdminController) Initialize(engine *gin.Engine, repo *product_repository.Repository) {
	controller.repo = repo
	engine.GET("/admin/products", controller.adminRoute)
	engine.GET("/admin/products/delete/:id", controller.deleteRoute)
	engine.GET("/admin/products/create", controller.createFormRoute)
	engine.POST("/admin/products/create", controller.createProductRoute)
	engine.GET("/admin/products/edit/:id", controller.editFormRoute)
	engine.POST("/admin/products/edit/:id", controller.editRoute)
}

func (controller *AdminController) adminRoute(c *gin.Context) {
	products, err := controller.repo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get products"})
	} else {
		c.HTML(http.StatusOK, "admin/admin.tmpl", gin.H{"products": products})
	}
}

func (controller *AdminController) deleteRoute(c *gin.Context) {
	id := c.Param("id")

	err := controller.repo.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete item with such id"})
	} else {
		c.Redirect(http.StatusPermanentRedirect, "/admin/products")
	}
}

func (controller *AdminController) editFormRoute(c *gin.Context) {
	id := c.Param("id")

	product, err := controller.repo.GetById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get item with such id"})
	} else {
		c.HTML(http.StatusOK, "admin/edit.tmpl", gin.H{"product": product})
	}
}

func (controller *AdminController) createFormRoute(c *gin.Context) {
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

	controller.repo.Create(&product)
}

func (controller *AdminController) editRoute(c *gin.Context) {
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

	err = controller.repo.Update(&product)

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
