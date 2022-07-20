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

package ec2

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/stretchr/testify/assert"
	"github.com/tchiunam/axolgo-cloud/aws/util"
)

// MockAWSConfigOptions is a mock implementation of AWSConfigOptions
// that can be used for testing error.
func MockWithAWSConfigError(v string) util.AWSConfigOptionsFunc {
	return func(o *util.AWSConfigOptions) error {
		o.Region = v
		return fmt.Errorf("mock error")
	}
}

// TestRunDescribeInstancesInvalid calls RunDescribeInstances
// and expects error since the AWS credentials are not set.
func TestRunDescribeInstancesInvalid(t *testing.T) {
	cases := map[string]struct {
		input               *ec2.DescribeInstancesInput
		optFns              func(*util.AWSConfigOptions) error
		expectStringInError string
	}{
		"nil input": {
			input:               nil,
			optFns:              nil,
			expectStringInError: "failed to retrieve credentials",
		},
		"input with no filter": {
			input:               &ec2.DescribeInstancesInput{Filters: nil},
			optFns:              nil,
			expectStringInError: "failed to retrieve credentials",
		},
		"nil input with error AWS config": {
			input:               nil,
			optFns:              MockWithAWSConfigError("us-east-1"),
			expectStringInError: "failed to retrieve credentials",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			_, _, err := RunDescribeInstances(c.input)
			assert.ErrorContains(
				t,
				err,
				c.expectStringInError,
				"Expected error containing %q, got %s", c.expectStringInError, err)
		})
	}
}
