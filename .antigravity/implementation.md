I have analyzed the repository structure. Now I am planning the implementation of the new geolocation image table, API, and Swagger documentation.

Edited task.md
Listed directory db
Listed directory migrations
Listed directory v1
Listed directory user
Viewed migrations.go:1-530
Listed directory v1_23

I have researched the migration system and identified that new migrations should be added to `models/migrations/v1_25/` and registered in `models/migrations/migrations.go`. I am now analyzing how models are defined and how API routing is handled to ensure proper integration.

Viewed user.go:1-800
Viewed api.go:1-800
Listed directory swagger
Searched for files: swagger.v1.json
Searched for "swagger:route"
Searched for "swagger:parameters"
Listed directory structs

I have identified the key locations for database models (`models/`), migrations (`models/migrations/`), API structures (`modules/structs/`), and API handlers (`routers/api/v1/`). I am now ready to draft the implementation plan.

Viewed user.go:1-139
Viewed user.go:1-58

I have researched the repository structure and identified the necessary changes across models, migrations, services, and APIs. I am now creating a comprehensive implementation plan for the user to review.

Edited implementation_plan.md
Listed directory b88a1b13-37fb-4ca2-bc54-444f729f316e

I have researched the repository structure and identified the necessary changes across models, migrations, services, and APIs. I have created a comprehensive implementation plan and am waiting for user approval before proceeding to implementation.

Edited task.md

I have created an implementation plan for adding the geolocation image feature. This includes database changes, API integration, and Swagger updates. Please review the plan.

Implementation Plan - Geolocation Image Feature
This plan outlines the steps to add a new table for storing base64 images and geolocation data, along with a corresponding REST API integrated into Gitea.

Proposed Changes
Database Layer
[NEW] 
image.go
Create a new model for geolocation images using XORM.

Fields: 
ID
, 
Name
, ImageBase64, Lat, Lon, ZoomLevel, Status, Description, CreatedUnix, UpdatedUnix.
Registration: Register the model in the 
init()
 function with db.RegisterModel.
[NEW] 
v323.go
Add a migration script to create the geo_image table.

[MODIFY] 
migrations.go
Register the new migration v323.

Service Layer
[NEW] 
image.go
Implement business logic for CRUD operations on geolocation images.

CreateImage: Validates and saves a new image.
GetImageByID: Retrieves an image by its ID.
UpdateImage: Updates an existing image.
DeleteImage: Deletes an image.
ListImages: Lists images with optional filtering.
Module Layer (API Structs)
[NEW] 
geo_image.go
Define the API request and response structures with Swagger annotations.

GeoImage: Represents the image object in the API.
CreateGeoImageOption: Parameters for creating an image.
EditGeoImageOption: Parameters for updating an image.
API Layer
[NEW] 
image.go
Implement REST API handlers for:

GET /geo/images: List images.
POST /geo/images: Create a new image.
GET /geo/images/{id}: Get image details.
PATCH /geo/images/{id}: Update image details.
DELETE /geo/images/{id}: Delete an image.
[MODIFY] 
api.go
Register the new /geo routes.

[NEW] 
geo_image.go
Add Swagger response definitions for the new API.

Verification Plan
Automated Tests
I will add a new test file tests/integration/api_geo_image_test.go to verify the CRUD operations.

Command: go test -v tests/integration/api_geo_image_test.go (Note: Need to check exact integration test command in this repo).
Manual Verification
Run Gitea in development mode (make watch or go run main.go).
Access the Swagger UI (/api/v1/swagger) to verify the new endpoints are listed.
Use curl or a REST client to test the CRUD endpoints:
Create an image with base64 data and geolocation.
List images and verify the new entry exists.
Update image details and verify changes.
Delete image and verify it is gone.