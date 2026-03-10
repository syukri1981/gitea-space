# Implementation Plan - Geolocation Image Feature

This plan outlines the steps to add a new table for storing base64 images and geolocation data, along with a corresponding REST API integrated into Gitea.

## Proposed Changes

### Database Layer

#### [NEW] [image.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/models/geo/image.go)
Create a new model for geolocation images using XORM.
- **Fields**: [ID](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/models/user/user.go#986-997), [Name](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/models/user/user.go#469-484), `ImageBase64`, `Lat`, `Lon`, `ZoomLevel`, `Status`, `Description`, `CreatedUnix`, `UpdatedUnix`.
- **Registration**: Register the model in the [init()](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/main.go#34-39) function with `db.RegisterModel`.

#### [NEW] [v323.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/models/migrations/v1_25/v323.go)
Add a migration script to create the `geo_image` table.

#### [MODIFY] [migrations.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/models/migrations/migrations.go)
Register the new migration `v323`.

---

### Service Layer

#### [NEW] [image.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/services/geo/image.go)
Implement business logic for CRUD operations on geolocation images.
- `CreateImage`: Validates and saves a new image.
- `GetImageByID`: Retrieves an image by its ID.
- `UpdateImage`: Updates an existing image.
- `DeleteImage`: Deletes an image.
- `ListImages`: Lists images with optional filtering.

---

### Module Layer (API Structs)

#### [NEW] [geo_image.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/modules/structs/geo_image.go)
Define the API request and response structures with Swagger annotations.
- `GeoImage`: Represents the image object in the API.
- `CreateGeoImageOption`: Parameters for creating an image.
- `EditGeoImageOption`: Parameters for updating an image.

---

### API Layer

#### [NEW] [image.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/routers/api/v1/geo/image.go)
Implement REST API handlers for:
- `GET /geo/images`: List images.
- `POST /geo/images`: Create a new image.
- `GET /geo/images/{id}`: Get image details.
- `PATCH /geo/images/{id}`: Update image details.
- `DELETE /geo/images/{id}`: Delete an image.

#### [MODIFY] [api.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/routers/api/v1/api.go)
Register the new `/geo` routes.

#### [NEW] [geo_image.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/routers/api/v1/swagger/geo_image.go)
Add Swagger response definitions for the new API.

---

## Verification Plan

### Automated Tests
I will add a new test file `tests/integration/api_geo_image_test.go` to verify the CRUD operations.
- **Command**: `go test -v tests/integration/api_geo_image_test.go` (Note: Need to check exact integration test command in this repo).

### Manual Verification
1. Run Gitea in development mode (`make watch` or `go run main.go`).
2. Access the Swagger UI (`/api/v1/swagger`) to verify the new endpoints are listed.
3. Use `curl` or a REST client to test the CRUD endpoints:
   - Create an image with base64 data and geolocation.
   - List images and verify the new entry exists.
   - Update image details and verify changes.
   - Delete image and verify it is gone.
