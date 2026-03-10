// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"testing"

	auth_model "code.gitea.io/gitea/models/auth"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/tests"

	"github.com/stretchr/testify/assert"
)

func TestAPIGeoImage(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	session := loginUser(t, "user2")
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeAll)

	var createdImg api.GeoImage

	t.Run("Create", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()
		opt := api.CreateGeoImageOption{
			Name:        "Test Image",
			ImageBase64: "dGVzdCBpZ21hZ2UgZGF0YQ==", // "test image data"
			Lat:         1.23,
			Lon:         4.56,
			ZoomLevel:   10,
			Status:      "active",
			Description: "A test geo image",
		}
		req := NewRequestWithJSON(t, "POST", "/api/v1/geo/images", &opt).
			AddTokenAuth(token)
		resp := MakeRequest(t, req, http.StatusCreated)

		DecodeJSON(t, resp, &createdImg)
		assert.Equal(t, opt.Name, createdImg.Name)
		assert.Equal(t, opt.Lat, createdImg.Lat)
		assert.Equal(t, opt.Lon, createdImg.Lon)
	})

	t.Run("List", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()
		req := NewRequest(t, "GET", "/api/v1/geo/images").
			AddTokenAuth(token)
		resp := MakeRequest(t, req, http.StatusOK)

		var imgs []api.GeoImage
		DecodeJSON(t, resp, &imgs)
		assert.NotEmpty(t, imgs)
		
		found := false
		for _, img := range imgs {
			if img.ID == createdImg.ID {
				found = true
				break
			}
		}
		assert.True(t, found)
	})

	t.Run("Get", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()
		req := NewRequestf(t, "GET", "/api/v1/geo/images/%d", createdImg.ID).
			AddTokenAuth(token)
		resp := MakeRequest(t, req, http.StatusOK)

		var img api.GeoImage
		DecodeJSON(t, resp, &img)
		assert.Equal(t, createdImg.ID, img.ID)
		assert.Equal(t, createdImg.Name, img.Name)
	})

	t.Run("Edit", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()
		newName := "Updated Test Image"
		opt := api.EditGeoImageOption{
			Name: &newName,
		}
		req := NewRequestWithJSON(t, "PATCH", NewRequestf(t, "PATCH", "/api/v1/geo/images/%d", createdImg.ID).URL.Path, &opt).
			AddTokenAuth(token)
		resp := MakeRequest(t, req, http.StatusOK)

		var img api.GeoImage
		DecodeJSON(t, resp, &img)
		assert.Equal(t, newName, img.Name)
	})

	t.Run("Delete", func(t *testing.T) {
		defer tests.PrintCurrentTest(t)()
		req := NewRequestf(t, "DELETE", "/api/v1/geo/images/%d", createdImg.ID).
			AddTokenAuth(token)
		MakeRequest(t, req, http.StatusNoContent)

		req = NewRequestf(t, "GET", "/api/v1/geo/images/%d", createdImg.ID).
			AddTokenAuth(token)
		MakeRequest(t, req, http.StatusNotFound)
	})
}
