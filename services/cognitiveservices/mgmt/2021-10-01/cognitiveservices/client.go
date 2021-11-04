// Package cognitiveservices implements the Azure ARM Cognitiveservices service API version 2021-10-01.
//
// Cognitive Services Management Client
package cognitiveservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Cognitiveservices
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Cognitiveservices.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckDomainAvailability check whether a domain is available.
// Parameters:
// parameters - check Domain Availability parameter.
func (client BaseClient) CheckDomainAvailability(ctx context.Context, parameters CheckDomainAvailabilityParameter) (result DomainAvailability, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckDomainAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SubdomainName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.BaseClient", "CheckDomainAvailability", err.Error())
	}

	req, err := client.CheckDomainAvailabilityPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckDomainAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckDomainAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckDomainAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckDomainAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckDomainAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckDomainAvailabilityPreparer prepares the CheckDomainAvailability request.
func (client BaseClient) CheckDomainAvailabilityPreparer(ctx context.Context, parameters CheckDomainAvailabilityParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/checkDomainAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckDomainAvailabilitySender sends the CheckDomainAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckDomainAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckDomainAvailabilityResponder handles the response to the CheckDomainAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckDomainAvailabilityResponder(resp *http.Response) (result DomainAvailability, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CheckSkuAvailability check available SKUs.
// Parameters:
// location - resource location.
// parameters - check SKU Availability POST body.
func (client BaseClient) CheckSkuAvailability(ctx context.Context, location string, parameters CheckSkuAvailabilityParameter) (result SkuAvailabilityListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckSkuAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Skus", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Kind", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.BaseClient", "CheckSkuAvailability", err.Error())
	}

	req, err := client.CheckSkuAvailabilityPreparer(ctx, location, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckSkuAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckSkuAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckSkuAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckSkuAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckSkuAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckSkuAvailabilityPreparer prepares the CheckSkuAvailability request.
func (client BaseClient) CheckSkuAvailabilityPreparer(ctx context.Context, location string, parameters CheckSkuAvailabilityParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/locations/{location}/checkSkuAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckSkuAvailabilitySender sends the CheckSkuAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckSkuAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckSkuAvailabilityResponder handles the response to the CheckSkuAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckSkuAvailabilityResponder(resp *http.Response) (result SkuAvailabilityListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
