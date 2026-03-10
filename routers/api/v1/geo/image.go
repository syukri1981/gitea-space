// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package geo

import (
	"net/http"

	"code.gitea.io/gitea/models/db"
	geo_model "code.gitea.io/gitea/models/geo"
	"code.gitea.io/gitea/services/context"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/web"
	geo_service "code.gitea.io/gitea/services/geo"
)

// ToGeoImage converts a geo_model.GeoImage to an api.GeoImage
func ToGeoImage(img *geo_model.GeoImage) *api.GeoImage {
	return &api.GeoImage{
		ID:          img.ID,
		Name:        img.Name,
		ImageBase64: img.ImageBase64,
		Lat:         img.Lat,
		Lon:         img.Lon,
		ZoomLevel:   img.ZoomLevel,
		Status:      img.Status,
		Description: img.Description,
		Created:     img.CreatedUnix.AsTime(),
		Updated:     img.UpdatedUnix.AsTime(),
	}
}

// ListImages list all geolocation images
func ListImages(ctx *context.APIContext) {
	// swagger:operation GET /geo/images geo listGeoImages
	// ---
	// summary: List all geolocation images
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/GeoImageList"

	imgs, count, err := geo_service.ListGeoImages(ctx, db.ListOptions{
		Page:     ctx.FormInt("page"),
		PageSize: ctx.FormInt("limit"),
	})
	if err != nil {
		ctx.APIErrorInternal(err)
		return
	}

	apiImgs := make([]*api.GeoImage, len(imgs))
	for i, img := range imgs {
		apiImgs[i] = ToGeoImage(img)
	}

	ctx.SetTotalCountHeader(count)
	ctx.JSON(http.StatusOK, apiImgs)
}

// CreateImage create a geolocation image
func CreateImage(ctx *context.APIContext) {
	// swagger:operation POST /geo/images geo createGeoImage
	// ---
	// summary: Create a geolocation image
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/CreateGeoImageOption"
	// responses:
	//   "201":
	//     "$ref": "#/responses/GeoImage"

	form := web.GetForm(ctx).(*api.CreateGeoImageOption)
	img, err := geo_service.CreateGeoImage(ctx, *form)
	if err != nil {
		ctx.APIErrorInternal(err)
		return
	}

	ctx.JSON(http.StatusCreated, ToGeoImage(img))
}

// GetImage get a geolocation image
func GetImage(ctx *context.APIContext) {
	// swagger:operation GET /geo/images/{id} geo getGeoImage
	// ---
	// summary: Get a geolocation image
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the image to get
	//   type: integer
	//   format: int64
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/GeoImage"
	//   "44":
	//     "$ref": "#/responses/notFound"

	id := ctx.PathParamInt64("id")
	img, err := geo_service.GetGeoImageByID(ctx, id)
	if err != nil {
		ctx.APIErrorInternal(err)
		return
	}
	if img == nil {
		ctx.APIErrorNotFound()
		return
	}

	ctx.JSON(http.StatusOK, ToGeoImage(img))
}

// EditImage edit a geolocation image
func EditImage(ctx *context.APIContext) {
	// swagger:operation PATCH /geo/images/{id} geo editGeoImage
	// ---
	// summary: Edit a geolocation image
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the image to edit
	//   type: integer
	//   format: int64
	//   required: true
	// - name: body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/EditGeoImageOption"
	// responses:
	//   "200":
	//     "$ref": "#/responses/GeoImage"
	//   "44":
	//     "$ref": "#/responses/notFound"

	id := ctx.PathParamInt64("id")
	form := web.GetForm(ctx).(*api.EditGeoImageOption)
	img, err := geo_service.UpdateGeoImage(ctx, id, *form)
	if err != nil {
		ctx.APIErrorInternal(err)
		return
	}
	if img == nil {
		ctx.APIErrorNotFound()
		return
	}

	ctx.JSON(http.StatusOK, ToGeoImage(img))
}

// DeleteImage delete a geolocation image
func DeleteImage(ctx *context.APIContext) {
	// swagger:operation DELETE /geo/images/{id} geo deleteGeoImage
	// ---
	// summary: Delete a geolocation image
	// parameters:
	// - name: id
	//   in: path
	//   description: id of the image to delete
	//   type: integer
	//   format: int64
	//   required: true
	// responses:
	//   "24":
	//     "$ref": "#/responses/empty"
	//   "44":
	//     "$ref": "#/responses/notFound"

	id := ctx.PathParamInt64("id")
	img, err := geo_service.GetGeoImageByID(ctx, id)
	if err != nil {
		ctx.APIErrorInternal(err)
		return
	}
	if img == nil {
		ctx.APIErrorNotFound()
		return
	}

	if err := geo_service.DeleteGeoImage(ctx, id); err != nil {
		ctx.APIErrorInternal(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
