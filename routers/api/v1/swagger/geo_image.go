// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package swagger

import (
	api "code.gitea.io/gitea/modules/structs"
)

// GeoImage
// swagger:response GeoImage
type swaggerResponseGeoImage struct {
	// in:body
	Body api.GeoImage `json:"body"`
}

// GeoImageList
// swagger:response GeoImageList
type swaggerResponseGeoImageList struct {
	// in:body
	Body []api.GeoImage `json:"body"`
}
