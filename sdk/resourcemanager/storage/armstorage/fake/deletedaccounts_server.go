//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"net/url"
	"regexp"
)

// DeletedAccountsServer is a fake server for instances of the armstorage.DeletedAccountsClient type.
type DeletedAccountsServer struct {
	// Get is the fake for method DeletedAccountsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deletedAccountName string, location string, options *armstorage.DeletedAccountsClientGetOptions) (resp azfake.Responder[armstorage.DeletedAccountsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DeletedAccountsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armstorage.DeletedAccountsClientListOptions) (resp azfake.PagerResponder[armstorage.DeletedAccountsClientListResponse])
}

// NewDeletedAccountsServerTransport creates a new instance of DeletedAccountsServerTransport with the provided implementation.
// The returned DeletedAccountsServerTransport instance is connected to an instance of armstorage.DeletedAccountsClient by way of the
// undefined.Transporter field.
func NewDeletedAccountsServerTransport(srv *DeletedAccountsServer) *DeletedAccountsServerTransport {
	return &DeletedAccountsServerTransport{srv: srv}
}

// DeletedAccountsServerTransport connects instances of armstorage.DeletedAccountsClient to instances of DeletedAccountsServer.
// Don't use this type directly, use NewDeletedAccountsServerTransport instead.
type DeletedAccountsServerTransport struct {
	srv          *DeletedAccountsServer
	newListPager *azfake.PagerResponder[armstorage.DeletedAccountsClientListResponse]
}

// Do implements the policy.Transporter interface for DeletedAccountsServerTransport.
func (d *DeletedAccountsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DeletedAccountsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DeletedAccountsClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DeletedAccountsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedAccounts/(?P<deletedAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deletedAccountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("deletedAccountName")])
	if err != nil {
		return nil, err
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), deletedAccountNameUnescaped, locationUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedAccount, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeletedAccountsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if d.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/deletedAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListPager(nil)
		d.newListPager = &resp
		server.PagerResponderInjectNextLinks(d.newListPager, req, func(page *armstorage.DeletedAccountsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListPager) {
		d.newListPager = nil
	}
	return resp, nil
}
