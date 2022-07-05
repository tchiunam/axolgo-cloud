/*
Copyright © 2022 tchiunam

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

// ListInstancesAPI defines the interface for the ListInstances function.
// We use this interface to test the function using a mocked service.
type ListInstancesAPI interface {
	List(
		ctx context.Context,
		req *computepb.ListInstancesRequest,
		opts ...gax.CallOption) *compute.InstanceIterator
}

// ListInstances retrieves an aggregated list of instances
// Inputs:
//     c is the context of the method call.
//     api is the interface that defines the method call.
//     req defines the request arguments to the service call.
//     opts is the options for the service call.
// Output:
//     Remote call client.
//     A InstanceIterator object containing the result of the service call.
func ListInstances(c context.Context, api ListInstancesAPI, req *computepb.ListInstancesRequest, opts ...gax.CallOption) *compute.InstanceIterator {
	return api.List(c, req)
}

// Retrieves an aggregated list of instances
func RunListInstances(req *computepb.ListInstancesRequest, opts ...option.ClientOption) (*compute.InstancesClient, *compute.InstanceIterator, error) {
	ctx := context.Background()
	c, err := compute.NewInstancesRESTClient(ctx, opts...)
	if err != nil {
		return c, nil, fmt.Errorf("Encountered error when listing instances: %v", err)
	}

	return c, c.List(ctx, req), nil
}
