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
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/tchiunam/axolgo-lib/util"
)

// TestRunDescribeInstances calls RunDescribeInstances
// and expects error since the AWS credentials are not set.
func TestRunDescribeInstances(t *testing.T) {
	cases := map[string]struct {
		Input               *ec2.DescribeInstancesInput
		ExpectStringInError string
	}{
		"nil input": {
			Input:               nil,
			ExpectStringInError: "failed to retrieve credentials",
		},
		"input with no filter": {
			Input:               &ec2.DescribeInstancesInput{Filters: nil},
			ExpectStringInError: "failed to retrieve credentials",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			_, _, err := RunDescribeInstances(c.Input)
			if !util.ErrorContains(err, c.ExpectStringInError) {
				t.Errorf("Expected error containing %q, got %s", c.ExpectStringInError, err)
			}
		})
	}
}
