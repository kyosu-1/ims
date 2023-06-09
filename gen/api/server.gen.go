// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
)

const (
	SessionAuthScopes = "SessionAuth.Scopes"
)

// Category defines model for Category.
type Category struct {
	Id   openapi_types.UUID `json:"id"`
	Name string             `json:"name"`
}

// Item defines model for Item.
type Item struct {
	AcquisitionDate openapi_types.Date `json:"acquisitionDate"`
	BorrowerName    *string            `json:"borrowerName,omitempty"`
	Categories      []Category         `json:"categories"`
	Description     *string            `json:"description,omitempty"`
	Id              openapi_types.UUID `json:"id"`
	ItemName        string             `json:"itemName"`
	Location        string             `json:"location"`
}

// ItemPostBody defines model for ItemPostBody.
type ItemPostBody struct {
	AcquisitionDate openapi_types.Date   `json:"acquisitionDate"`
	CategoryIDs     []openapi_types.UUID `json:"categoryIDs"`
	Description     *string              `json:"description,omitempty"`
	Id              openapi_types.UUID   `json:"id"`
	ItemName        string               `json:"itemName"`
	Location        string               `json:"location"`
}

// User defines model for User.
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// UserSignin defines model for UserSignin.
type UserSignin struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

// CategoryID defines model for categoryID.
type CategoryID = openapi_types.UUID

// ItemID defines model for itemID.
type ItemID = openapi_types.UUID

// PostCategoriesJSONBody defines parameters for PostCategories.
type PostCategoriesJSONBody = Category

// PostItemsJSONBody defines parameters for PostItems.
type PostItemsJSONBody = ItemPostBody

// PostSigninJSONBody defines parameters for PostSignin.
type PostSigninJSONBody = UserSignin

// PostCategoriesJSONRequestBody defines body for PostCategories for application/json ContentType.
type PostCategoriesJSONRequestBody = PostCategoriesJSONBody

// PostItemsJSONRequestBody defines body for PostItems for application/json ContentType.
type PostItemsJSONRequestBody = PostItemsJSONBody

// PostSigninJSONRequestBody defines body for PostSignin for application/json ContentType.
type PostSigninJSONRequestBody = PostSigninJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all categories
	// (GET /categories)
	GetCategories(ctx echo.Context) error
	// Create or update a category
	// (POST /categories)
	PostCategories(ctx echo.Context) error
	// Delete a category
	// (DELETE /categories/{categoryID})
	DeleteCategoriesCategoryID(ctx echo.Context, categoryID CategoryID) error
	// Get a category
	// (GET /categories/{categoryID})
	GetCategoriesCategoryID(ctx echo.Context, categoryID CategoryID) error
	// Health check
	// (GET /health)
	GetHealth(ctx echo.Context) error
	// Get all items
	// (GET /items)
	GetItems(ctx echo.Context) error
	// Create or update an item
	// (POST /items)
	PostItems(ctx echo.Context) error
	// Delete an item
	// (DELETE /items/{itemID})
	DeleteItemsItemID(ctx echo.Context, itemID ItemID) error
	// Get an item
	// (GET /items/{itemID})
	GetItemsItemID(ctx echo.Context, itemID ItemID) error
	// Borrow an item
	// (POST /items/{itemID}/borrow)
	PostItemsItemIDBorrow(ctx echo.Context, itemID ItemID) error
	// Return an item
	// (POST /items/{itemID}/return)
	PostItemsItemIDReturn(ctx echo.Context, itemID ItemID) error
	// Get logged in user's information
	// (GET /me)
	GetMe(ctx echo.Context) error
	// User sign in
	// (POST /signin)
	PostSignin(ctx echo.Context) error
	// User sign out
	// (POST /signout)
	PostSignout(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCategories converts echo context to params.
func (w *ServerInterfaceWrapper) GetCategories(ctx echo.Context) error {
	var err error

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCategories(ctx)
	return err
}

// PostCategories converts echo context to params.
func (w *ServerInterfaceWrapper) PostCategories(ctx echo.Context) error {
	var err error

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostCategories(ctx)
	return err
}

// DeleteCategoriesCategoryID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCategoriesCategoryID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "categoryID" -------------
	var categoryID CategoryID

	err = runtime.BindStyledParameterWithLocation("simple", false, "categoryID", runtime.ParamLocationPath, ctx.Param("categoryID"), &categoryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter categoryID: %s", err))
	}

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteCategoriesCategoryID(ctx, categoryID)
	return err
}

// GetCategoriesCategoryID converts echo context to params.
func (w *ServerInterfaceWrapper) GetCategoriesCategoryID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "categoryID" -------------
	var categoryID CategoryID

	err = runtime.BindStyledParameterWithLocation("simple", false, "categoryID", runtime.ParamLocationPath, ctx.Param("categoryID"), &categoryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter categoryID: %s", err))
	}

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCategoriesCategoryID(ctx, categoryID)
	return err
}

// GetHealth converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetHealth(ctx)
	return err
}

// GetItems converts echo context to params.
func (w *ServerInterfaceWrapper) GetItems(ctx echo.Context) error {
	var err error

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetItems(ctx)
	return err
}

// PostItems converts echo context to params.
func (w *ServerInterfaceWrapper) PostItems(ctx echo.Context) error {
	var err error

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostItems(ctx)
	return err
}

// DeleteItemsItemID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteItemsItemID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "itemID" -------------
	var itemID ItemID

	err = runtime.BindStyledParameterWithLocation("simple", false, "itemID", runtime.ParamLocationPath, ctx.Param("itemID"), &itemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemID: %s", err))
	}

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteItemsItemID(ctx, itemID)
	return err
}

// GetItemsItemID converts echo context to params.
func (w *ServerInterfaceWrapper) GetItemsItemID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "itemID" -------------
	var itemID ItemID

	err = runtime.BindStyledParameterWithLocation("simple", false, "itemID", runtime.ParamLocationPath, ctx.Param("itemID"), &itemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemID: %s", err))
	}

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetItemsItemID(ctx, itemID)
	return err
}

// PostItemsItemIDBorrow converts echo context to params.
func (w *ServerInterfaceWrapper) PostItemsItemIDBorrow(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "itemID" -------------
	var itemID ItemID

	err = runtime.BindStyledParameterWithLocation("simple", false, "itemID", runtime.ParamLocationPath, ctx.Param("itemID"), &itemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemID: %s", err))
	}

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostItemsItemIDBorrow(ctx, itemID)
	return err
}

// PostItemsItemIDReturn converts echo context to params.
func (w *ServerInterfaceWrapper) PostItemsItemIDReturn(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "itemID" -------------
	var itemID ItemID

	err = runtime.BindStyledParameterWithLocation("simple", false, "itemID", runtime.ParamLocationPath, ctx.Param("itemID"), &itemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemID: %s", err))
	}

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostItemsItemIDReturn(ctx, itemID)
	return err
}

// GetMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetMe(ctx echo.Context) error {
	var err error

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMe(ctx)
	return err
}

// PostSignin converts echo context to params.
func (w *ServerInterfaceWrapper) PostSignin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostSignin(ctx)
	return err
}

// PostSignout converts echo context to params.
func (w *ServerInterfaceWrapper) PostSignout(ctx echo.Context) error {
	var err error

	ctx.Set(SessionAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostSignout(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/categories", wrapper.GetCategories)
	router.POST(baseURL+"/categories", wrapper.PostCategories)
	router.DELETE(baseURL+"/categories/:categoryID", wrapper.DeleteCategoriesCategoryID)
	router.GET(baseURL+"/categories/:categoryID", wrapper.GetCategoriesCategoryID)
	router.GET(baseURL+"/health", wrapper.GetHealth)
	router.GET(baseURL+"/items", wrapper.GetItems)
	router.POST(baseURL+"/items", wrapper.PostItems)
	router.DELETE(baseURL+"/items/:itemID", wrapper.DeleteItemsItemID)
	router.GET(baseURL+"/items/:itemID", wrapper.GetItemsItemID)
	router.POST(baseURL+"/items/:itemID/borrow", wrapper.PostItemsItemIDBorrow)
	router.POST(baseURL+"/items/:itemID/return", wrapper.PostItemsItemIDReturn)
	router.GET(baseURL+"/me", wrapper.GetMe)
	router.POST(baseURL+"/signin", wrapper.PostSignin)
	router.POST(baseURL+"/signout", wrapper.PostSignout)

}
