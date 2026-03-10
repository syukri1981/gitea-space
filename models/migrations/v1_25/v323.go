// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_25

import (
	"xorm.io/xorm"
)

func AddGeoImageTable(x *xorm.Engine) error {
	type GeoImage struct {
		ID          int64   `xorm:"pk autoincr"`
		Name        string  `xorm:"VARCHAR(255) NOT NULL"`
		ImageBase64 string  `xorm:"LONGTEXT NOT NULL"`
		Lat         float64 `xorm:"DOUBLE NOT NULL"`
		Lon         float64 `xorm:"DOUBLE NOT NULL"`
		ZoomLevel   int     `xorm:"INTEGER NOT NULL"`
		Status      string  `xorm:"VARCHAR(50) NOT NULL"`
		Description string  `xorm:"TEXT"`
		CreatedUnix int64   `xorm:"INDEX 'created_dtm' created"`
		UpdatedUnix int64   `xorm:"INDEX 'update_dtm' updated"`
	}
	return x.Sync(new(GeoImage))
}
