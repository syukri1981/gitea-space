// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package geo

import (
	"context"

	"code.gitea.io/gitea/models/db"
	geo_model "code.gitea.io/gitea/models/geo"
	api "code.gitea.io/gitea/modules/structs"
)

// CreateGeoImage creates a new geolocation image.
func CreateGeoImage(ctx context.Context, opts api.CreateGeoImageOption) (*geo_model.GeoImage, error) {
	img := &geo_model.GeoImage{
		Name:        opts.Name,
		ImageBase64: opts.ImageBase64,
		Lat:         opts.Lat,
		Lon:         opts.Lon,
		ZoomLevel:   opts.ZoomLevel,
		Status:      opts.Status,
		Description: opts.Description,
	}
	if err := geo_model.InsertGeoImage(ctx, img); err != nil {
		return nil, err
	}
	return img, nil
}

// GetGeoImageByID gets a geolocation image by ID.
func GetGeoImageByID(ctx context.Context, id int64) (*geo_model.GeoImage, error) {
	return geo_model.GetGeoImageByID(ctx, id)
}

// UpdateGeoImage updates an existing geolocation image.
func UpdateGeoImage(ctx context.Context, id int64, opts api.EditGeoImageOption) (*geo_model.GeoImage, error) {
	img, err := GetGeoImageByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if img == nil {
		return nil, nil
	}

	if opts.Name != nil {
		img.Name = *opts.Name
	}
	if opts.ImageBase64 != nil {
		img.ImageBase64 = *opts.ImageBase64
	}
	if opts.Lat != nil {
		img.Lat = *opts.Lat
	}
	if opts.Lon != nil {
		img.Lon = *opts.Lon
	}
	if opts.ZoomLevel != nil {
		img.ZoomLevel = *opts.ZoomLevel
	}
	if opts.Status != nil {
		img.Status = *opts.Status
	}
	if opts.Description != nil {
		img.Description = *opts.Description
	}

	if _, err := db.GetEngine(ctx).ID(img.ID).AllCols().Update(img); err != nil {
		return nil, err
	}
	return img, nil
}

// DeleteGeoImage deletes a geolocation image.
func DeleteGeoImage(ctx context.Context, id int64) error {
	_, err := db.GetEngine(ctx).ID(id).Delete(new(geo_model.GeoImage))
	return err
}

// ListGeoImages lists geolocation images.
func ListGeoImages(ctx context.Context, listOptions db.ListOptions) ([]*geo_model.GeoImage, int64, error) {
	sess := db.GetEngine(ctx)
	if listOptions.Page > 0 {
		sess = db.SetSessionPagination(sess, &listOptions)
	}
	imgs := make([]*geo_model.GeoImage, 0, listOptions.PageSize)
	count, err := sess.Count(new(geo_model.GeoImage))
	if err != nil {
		return nil, 0, err
	}
	err = sess.Find(&imgs)
	return imgs, count, err
}
