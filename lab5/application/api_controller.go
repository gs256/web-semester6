package application

import (
	"lab5/orders"
	"lab5/products"
	"lab5/users"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	productService *products.ProductService
	userService    *users.UserService
	orderService   *orders.OrderService
}

func (controller *ApiController) Initialize(engine *gin.Engine, productService *products.ProductService, userService *users.UserService, orderService *orders.OrderService) {
	controller.productService = productService
	controller.userService = userService
	controller.orderService = orderService
	engine.GET("/api/products", controller.productsRoute)
	engine.DELETE("/api/products/:id", controller.productsDeleteRoute)
	engine.GET("/api/users", controller.usersRoute)
	engine.POST("/api/users/create", controller.userCreateRoute)
	engine.POST("/api/users/edit", controller.userEditRoute)
	engine.DELETE("/api/users/:id", controller.userRemoveRoute)
	engine.GET("/api/orders", controller.ordersRoute)
	engine.DELETE("/api/orders/all", controller.orderClearRoute)
	engine.POST("/api/orders/create", controller.orderCreateRoute)
	engine.POST("/api/orders/filter", controller.orderFilterRoute)
	engine.POST("/api/analytics/order-history", controller.orderHistoryRoute)
}

type ProductDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}

func (controller *ApiController) productsRoute(c *gin.Context) {
	products, err := controller.productService.GetAllProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get products"})
	}

	productDtos := ToProductDtos(products)
	c.JSON(http.StatusOK, productDtos)
}

func (controller *ApiController) productsDeleteRoute(c *gin.Context) {
	id := c.Param("id")
	err := controller.productService.RemoveProductById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete product"})
	}
}

func ToProductDtos(products []products.Product) []ProductDto {
	productDtos := make([]ProductDto, len(products))
	for i := 0; i < len(products); i++ {
		productDtos[len(products)-i-1] = ToProductDto(products[i])
	}
	return productDtos
}

func ToProductDto(product products.Product) ProductDto {
	return ProductDto{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageUrl:    product.ImageUrl,
	}
}

func (controller *ApiController) usersRoute(c *gin.Context) {
	users, err := controller.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get users"})
	}

	userDtos := ToDtoArray(users)
	c.JSON(http.StatusOK, userDtos)
}

type UserDto struct {
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type CreateUserDto struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

func (controller *ApiController) userCreateRoute(c *gin.Context) {
	var createUserDto CreateUserDto
	err := c.BindJSON(&createUserDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	user := CreateDtoToUser(&createUserDto)
	id, err := controller.userService.CreateUser(user)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	user.Id = id
	userDto := ToDto(&user)
	c.JSON(http.StatusOK, userDto)
}

func (controller *ApiController) userEditRoute(c *gin.Context) {
	var userDto UserDto
	err := c.BindJSON(&userDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	user := DtoToUser(userDto)
	err = controller.userService.UpdateUser(user)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func DtoToUser(userDto UserDto) users.User {
	return users.User{
		Id:    userDto.Id,
		Name:  userDto.Name,
		Phone: userDto.Phone,
	}
}

func (controller *ApiController) userRemoveRoute(c *gin.Context) {
	id := c.Param("id")
	err := controller.userService.RemoveUser(id)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
	}
}

func CreateDtoToUser(createUserDto *CreateUserDto) users.User {
	return users.User{
		Name:  createUserDto.Name,
		Phone: createUserDto.Phone,
	}
}

func ToDto(user *users.User) UserDto {
	return UserDto{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
	}
}

func ToDtoArray(users []users.User) []UserDto {
	userDtos := make([]UserDto, len(users))
	for i := 0; i < len(userDtos); i++ {
		userDtos[len(userDtos)-i-1] = ToDto(&users[i])
	}
	return userDtos
}

type CreateOrderDto struct {
	UserId     string   `json:"userId" binding:"required"`
	ProductIds []string `json:"productIds" binding:"required"`
}

type CreateOrderResponseDto struct {
	OrderId string `json:"orderId" binding:"required"`
}

type OrderFilterDto struct {
	MinDate time.Time `json:"minDate"`
	MaxDate time.Time `json:"maxDate"`
}

func (controller *ApiController) orderCreateRoute(c *gin.Context) {
	var createOrderDto CreateOrderDto
	err := c.BindJSON(&createOrderDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := controller.orderService.CreateOrder(createOrderDto.UserId, createOrderDto.ProductIds)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	response := CreateOrderResponseDto{OrderId: id}
	c.JSON(http.StatusOK, response)
}

func (controller *ApiController) orderFilterRoute(c *gin.Context) {
	var orderFilterDto OrderFilterDto
	err := c.BindJSON(&orderFilterDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := controller.orderService.GetWithinTimespan(orderFilterDto.MinDate, orderFilterDto.MaxDate)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	orderDtos := ToOrderDtos(orders)
	c.JSON(http.StatusOK, orderDtos)
}

type OrderHistoryRequestDto struct {
	UserId string `json:"userId" binding:"required"`
}

func (controller *ApiController) orderHistoryRoute(c *gin.Context) {
	var orderHistoryDto OrderHistoryRequestDto
	err := c.BindJSON(&orderHistoryDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	products, err := controller.orderService.GetUserOrderHistory(orderHistoryDto.UserId)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	productDtos := ToProductDtos(products)
	c.JSON(http.StatusOK, productDtos)
}

func (controller *ApiController) ordersRoute(c *gin.Context) {
	orders, err := controller.orderService.GetAllOrders()

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	orderDtos := ToOrderDtos(orders)
	c.JSON(http.StatusOK, orderDtos)
}

func (controller *ApiController) orderClearRoute(c *gin.Context) {
	err := controller.orderService.Clear()

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
}

type OrderDto struct {
	Id        string       `json:"id"`
	User      UserDto      `json:"user"`
	Products  []ProductDto `json:"products"`
	CreatedAt time.Time    `json:"createdAt"`
}

func ToOrderDto(order orders.Order) OrderDto {
	return OrderDto{
		Id: order.Id,
		User: UserDto{
			Id:    order.User.Id,
			Name:  order.User.Name,
			Phone: order.User.Phone,
		},
		Products:  ToProductDtos(order.Products),
		CreatedAt: order.CreatedAt,
	}
}

func ToOrderDtos(orders []orders.Order) []OrderDto {
	orderDtos := make([]OrderDto, len(orders))
	for i := 0; i < len(orderDtos); i++ {
		orderDtos[i] = ToOrderDto(orders[i])
	}
	return orderDtos
}
