// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package structs

import (
	"time"
)

// GeoImage represents a geolocation image for the API
// swagger:model
type GeoImage struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	ImageBase64 string    `json:"image_base64"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
	ZoomLevel   int       `json:"zoom_level"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	// swagger:strfmt date-time
	Created     time.Time `json:"created"`
	// swagger:strfmt date-time
	Updated     time.Time `json:"updated"`
}

// CreateGeoImageOption options for creating a geolocation image
// swagger:model
type CreateGeoImageOption struct {
	// required: true
	Name        string  `json:"name" binding:"Required;MaxSize(255)"`
	// required: true
	ImageBase64 string  `json:"image_base64" binding:"Required"`
	// required: true
	Lat         float64 `json:"lat" binding:"Required"`
	// required: true
	Lon         float64 `json:"lon" binding:"Required"`
	// required: true
	ZoomLevel   int     `json:"zoom_level" binding:"Required"`
	Status      string  `json:"status" binding:"MaxSize(50)"`
	Description string  `json:"description"`
}

// EditGeoImageOption options for editing a geolocation image
// swagger:model
type EditGeoImageOption struct {
	Name        *string  `json:"name" binding:"MaxSize(255)"`
	ImageBase64 *string  `json:"image_base64"`
	Lat         *float64 `json:"lat"`
	Lon         *float64 `json:"lon"`
	ZoomLevel   *int     `json:"zoom_level"`
	Status      *string  `json:"status" binding:"MaxSize(50)"`
	Description *string  `json:"description"`
}
