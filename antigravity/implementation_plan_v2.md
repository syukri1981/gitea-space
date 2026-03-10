# Implementation Plan - Continue Geolocation Image Feature

This plan covers the final steps to complete the Geolocation Image Feature, ensuring the API routes are correct and verified with integration tests.

## Proposed Changes

### API Layer

#### [MODIFY] [api.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/routers/api/v1/api.go)
- Fix the bug in route registration for `POST /geo/images`.
- Line 1174: Change `geo.ListImages` to `geo.CreateImage`.

### Testing

#### [NEW] [api_geo_image_test.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/tests/integration/api_geo_image_test.go)
- Implement integration tests for the Geolocation Image API:
    - `TestAPIGeoImageCreate`: Verify creation of a geo image.
    - `TestAPIGeoImageList`: Verify listing of geo images.
    - `TestAPIGeoImageGet`: Verify retrieval of a single geo image.
    - `TestAPIGeoImageEdit`: Verify updating of a geo image.
    - `TestAPIGeoImageDelete`: Verify deletion of a geo image.

## Verification Plan

### Automated Tests
The integration tests will be run using the standard Gitea test suite:
```bash
go test -v tests/integration/api_geo_image_test.go tests/integration/integration_test.go
```
> [!NOTE]
> Integration tests require [tests/integration/integration_test.go](file:///Users/macbook/Documents/radiusdata/project001/rukerai/sas/gitea-space/tests/integration/integration_test.go) to be included for proper setup of the test environment.

### Manual Verification
1. Start Gitea: `go run main.go`
2. Verify endpoints via Swagger UI at `/api/v1/swagger`.
3. Perform CRUD operations using a REST client or `curl`.
