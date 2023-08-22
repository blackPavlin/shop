// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package rest

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for OrderStatus.
const (
	OrderStatusACCEPTED OrderStatus = "ACCEPTED"

	OrderStatusCANCELED OrderStatus = "CANCELED"

	OrderStatusCREATED OrderStatus = "CREATED"
)

// Defines values for UserRole.
const (
	UserRoleADMIN UserRole = "ADMIN"

	UserRoleUSER UserRole = "USER"
)

// Address defines model for Address.
type Address struct {
	City     string  `json:"city"`
	Country  string  `json:"country"`
	Flat     *int    `json:"flat,omitempty"`
	House    int     `json:"house"`
	Id       int64   `json:"id"`
	Letter   *string `json:"letter,omitempty"`
	Postcode int     `json:"postcode"`
	Street   string  `json:"street"`
	UserId   int64   `json:"userId"`
}

// AddressList defines model for AddressList.
type AddressList []Address

// Cart defines model for Cart.
type Cart struct {
	Amount  int64   `json:"amount"`
	Id      int64   `json:"id"`
	Product Product `json:"product"`
	UserId  int64   `json:"userId"`
}

// CartList defines model for CartList.
type CartList []Cart

// CartProduct defines model for CartProduct.
type CartProduct struct {
	Amount    int   `json:"amount"`
	ProductId int64 `json:"productId"`
}

// Category defines model for Category.
type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// CategoryList defines model for CategoryList.
type CategoryList []Category

// CreateAddressRequest defines model for CreateAddressRequest.
type CreateAddressRequest struct {
	City     string  `json:"city"`
	Country  string  `json:"country"`
	Flat     *int    `json:"flat,omitempty"`
	House    int     `json:"house"`
	Letter   *string `json:"letter,omitempty"`
	Postcode int     `json:"postcode"`
	Street   string  `json:"street"`
}

// CreateCategoryRequest defines model for CreateCategoryRequest.
type CreateCategoryRequest struct {
	Name string `json:"name"`
}

// CreateProductRequest defines model for CreateProductRequest.
type CreateProductRequest struct {
	Amount      int64  `json:"amount"`
	CategoryId  int    `json:"categoryId"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Image defines model for Image.
type Image struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	ProductId int64  `json:"productId"`
}

// ImageList defines model for ImageList.
type ImageList []Image

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	Token string `json:"token"`
}

// Order defines model for Order.
type Order struct {
	Id     int64       `json:"id"`
	Status OrderStatus `json:"status"`
	UserId int64       `json:"userId"`
}

// OrderStatus defines model for Order.Status.
type OrderStatus string

// OrderDetailed defines model for OrderDetailed.
type OrderDetailed struct {
	Id       int64            `json:"id"`
	Products OrderProductList `json:"products"`
	UserId   int64            `json:"userId"`
}

// OrderList defines model for OrderList.
type OrderList []Order

// OrderProduct defines model for OrderProduct.
type OrderProduct struct {
	Amount    int   `json:"amount"`
	ProductId int64 `json:"productId"`
}

// OrderProductList defines model for OrderProductList.
type OrderProductList []OrderProduct

// PaginationResponse defines model for PaginationResponse.
type PaginationResponse struct {
	Quantity int `json:"quantity"`
}

// Product defines model for Product.
type Product struct {
	Amount      int64     `json:"amount"`
	CategoryId  int64     `json:"categoryId"`
	Description *string   `json:"description,omitempty"`
	Id          int64     `json:"id"`
	Images      ImageList `json:"images"`
	Name        string    `json:"name"`
	Price       int64     `json:"price"`
}

// ProductList defines model for ProductList.
type ProductList struct {
	// Embedded struct due to allOf(#/components/schemas/PaginationResponse)
	PaginationResponse `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Items []Product `json:"items"`
}

// SignupRequest defines model for SignupRequest.
type SignupRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// UploadImagesRequest defines model for UploadImagesRequest.
type UploadImagesRequest struct {
	Files []string `json:"files"`
}

// User defines model for User.
type User struct {
	Email string   `json:"email"`
	Id    int64    `json:"id"`
	Name  string   `json:"name"`
	Phone string   `json:"phone"`
	Role  UserRole `json:"role"`
}

// UserRole defines model for User.Role.
type UserRole string

// BadRequest defines model for BadRequest.
type BadRequest Error

// Conflict defines model for Conflict.
type Conflict Error

// InternalError defines model for InternalError.
type InternalError Error

// NotFound defines model for NotFound.
type NotFound Error

// Unauthorized defines model for Unauthorized.
type Unauthorized Error

// PostAddressJSONBody defines parameters for PostAddress.
type PostAddressJSONBody CreateAddressRequest

// PostAuthLoginJSONBody defines parameters for PostAuthLogin.
type PostAuthLoginJSONBody LoginRequest

// PostAuthSignupJSONBody defines parameters for PostAuthSignup.
type PostAuthSignupJSONBody SignupRequest

// PatchCartJSONBody defines parameters for PatchCart.
type PatchCartJSONBody CartProduct

// PostCartJSONBody defines parameters for PostCart.
type PostCartJSONBody CartProduct

// PatchCategoryJSONBody defines parameters for PatchCategory.
type PatchCategoryJSONBody Category

// PostCategoryJSONBody defines parameters for PostCategory.
type PostCategoryJSONBody CreateCategoryRequest

// PostProductJSONBody defines parameters for PostProduct.
type PostProductJSONBody CreateProductRequest

// PatchProductProductIdJSONBody defines parameters for PatchProductProductId.
type PatchProductProductIdJSONBody CreateProductRequest

// PostAddressJSONRequestBody defines body for PostAddress for application/json ContentType.
type PostAddressJSONRequestBody PostAddressJSONBody

// PostAuthLoginJSONRequestBody defines body for PostAuthLogin for application/json ContentType.
type PostAuthLoginJSONRequestBody PostAuthLoginJSONBody

// PostAuthSignupJSONRequestBody defines body for PostAuthSignup for application/json ContentType.
type PostAuthSignupJSONRequestBody PostAuthSignupJSONBody

// PatchCartJSONRequestBody defines body for PatchCart for application/json ContentType.
type PatchCartJSONRequestBody PatchCartJSONBody

// PostCartJSONRequestBody defines body for PostCart for application/json ContentType.
type PostCartJSONRequestBody PostCartJSONBody

// PatchCategoryJSONRequestBody defines body for PatchCategory for application/json ContentType.
type PatchCategoryJSONRequestBody PatchCategoryJSONBody

// PostCategoryJSONRequestBody defines body for PostCategory for application/json ContentType.
type PostCategoryJSONRequestBody PostCategoryJSONBody

// PostProductJSONRequestBody defines body for PostProduct for application/json ContentType.
type PostProductJSONRequestBody PostProductJSONBody

// PatchProductProductIdJSONRequestBody defines body for PatchProductProductId for application/json ContentType.
type PatchProductProductIdJSONRequestBody PatchProductProductIdJSONBody
