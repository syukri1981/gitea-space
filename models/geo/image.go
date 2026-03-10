// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package geo

import (
	"context"

	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/modules/timeutil"
)

// GeoImage represents a geolocation image entity.
type GeoImage struct {
	ID          int64              `xorm:"pk autoincr"`
	Name        string             `xorm:"VARCHAR(255) NOT NULL"`
	ImageBase64 string             `xorm:"LONGTEXT NOT NULL"`
	Lat         float64            `xorm:"DOUBLE NOT NULL"`
	Lon         float64            `xorm:"DOUBLE NOT NULL"`
	ZoomLevel   int                `xorm:"INTEGER NOT NULL"`
	Status      string             `xorm:"VARCHAR(50) NOT NULL"`
	Description string             `xorm:"TEXT"`
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX 'created_dtm' created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"INDEX 'update_dtm' updated"`
}

func init() {
	db.RegisterModel(new(GeoImage))
}

// InsertGeoImage inserts a new geolocation image into the database.
func InsertGeoImage(ctx context.Context, img *GeoImage) error {
	return db.Insert(ctx, img)
}

// GetGeoImageByID retrieves a geolocation image by its ID.
func GetGeoImageByID(ctx context.Context, id int64) (*GeoImage, error) {
	img := new(GeoImage)
	has, err := db.GetEngine(ctx).ID(id).Get(img)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil // Or return a specific error if preferred
	}
	return img, nil
}
