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

package rds

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tchiunam/axolgo-cloud/aws/util"
	"github.com/tchiunam/axolgo-lib/types"
)

// MockAWSConfigOptions is a mock implementation of AWSConfigOptions
// that can be used for testing error.
func MockWithAWSConfigErrorMCDBPG(v string) util.AWSConfigOptionsFunc {
	return func(o *util.AWSConfigOptions) error {
		o.Region = v
		return fmt.Errorf("mock error")
	}
}

// TestRunModifyDBClusterParameterGroupInvaoid calls RunModifyDBClusterParameterGroup
// and expects error since the AWS credentials are not set.
func TestRunModifyDBClusterParameterGroupInvalid(t *testing.T) {
	paramName := "some-parameter-name"
	paramValue := "some-parameter-value"

	cases := map[string]struct {
		parameterGroupName  string
		staticParameters    []types.Parameter
		dynamicParameters   []types.Parameter
		optFns              func(*util.AWSConfigOptions) error
		expectStringInError string
	}{
		"static params": {
			parameterGroupName: "dev-db-parameter-group",
			staticParameters: []types.Parameter{{
				Name:  &paramName,
				Value: &paramValue,
			}},
			dynamicParameters: []types.Parameter{{
				Name:  &paramName,
				Value: &paramValue,
			}},
			optFns:              util.WithRegion("us-east-1"),
			expectStringInError: "failed to retrieve credentials",
		},
		"nil input with error AWS config": {
			parameterGroupName: "dev-db-parameter-group",
			staticParameters: []types.Parameter{{
				Name:  &paramName,
				Value: &paramValue,
			}},
			dynamicParameters: []types.Parameter{{
				Name:  &paramName,
				Value: &paramValue,
			}},
			optFns:              MockWithAWSConfigErrorMCDBPG("us-east-1"),
			expectStringInError: "failed to retrieve credentials",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := RunModifyDBClusterParameterGroup(c.parameterGroupName, c.staticParameters, c.dynamicParameters)
			if !strings.Contains(err.Error(), c.expectStringInError) {
				t.Errorf("Expected error containing %q, got %s", c.expectStringInError, err)
			}
		})
	}
}
