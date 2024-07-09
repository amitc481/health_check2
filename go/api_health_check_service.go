package openapi

/*
 * DKAM Service API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0-dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

// package openapi

import (
	"context"
	"fmt"
	"log"

	"github.com/GIT_USER_ID/GIT_REPO_ID/dbpackage"
	// "github.com/amitc481/health_check/dbpackage"

	"go.mongodb.org/mongo-driver/bson"
)

// HealthCheckAPIService is a service that implements the logic for the HealthCheckAPIServicer
// This service should implement the business logic for every endpoint for the HealthCheckAPI API.
// Include any external packages or services that will be required by this service.
type HealthCheckAPIService struct {
}

// NewHealthCheckAPIService creates a default api service
func NewHealthCheckAPIService() HealthCheckAPIServicer {
	return &HealthCheckAPIService{}
}

// GetHealthCheck - Retrieves all health_check
func (s *HealthCheckAPIService) GetHealthCheck(ctx context.Context) (ImplResponse, error) {

	fmt.Println("jjaaja")

	var DB = dbpackage.DB
	const table = "health_check_table"

	filter := bson.D{{}}

	var result []map[string]interface{}
	err := dbpackage.GetCollectionData(dbpackage.MongoCtx, dbpackage.MongoClient, DB, table, filter, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("kkkk", result)

	// TODO: Uncomment the next line to return response Response(200, []HealthCheckResource{}) or use other options such as http.Ok ...
	// return Response(200, []HealthCheckResource{}), nil
	if len(result) > 0 {
		return Response(200, result), nil
	}

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	return Response(404, result), nil

	// TODO - update GetHealthCheck with the required logic for this service method.
	// Add api_health_check_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// r := table + " Table does exist in the " + DB + " database.."

	// return Response(http.StatusNotImplemented, nil), errors.New(r)
}
