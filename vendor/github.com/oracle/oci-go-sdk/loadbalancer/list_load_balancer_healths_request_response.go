// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLoadBalancerHealthsRequest wrapper for the ListLoadBalancerHealths operation
type ListLoadBalancerHealthsRequest struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancers to return health status information for.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`
}

func (request ListLoadBalancerHealthsRequest) String() string {
	return common.PointerString(request)
}

// ListLoadBalancerHealthsResponse wrapper for the ListLoadBalancerHealths operation
type ListLoadBalancerHealthsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []LoadBalancerHealthSummary instance
	Items []LoadBalancerHealthSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLoadBalancerHealthsResponse) String() string {
	return common.PointerString(response)
}
