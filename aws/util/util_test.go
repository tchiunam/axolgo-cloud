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

package util

import (
	"testing"
)

// TestLoadAWSConfig calls LoadAWSConfig and expects nothing to happen
// since the AWS credentials are not set.
func TestLoadAWSConfig(t *testing.T) {
	cases := map[string]struct {
		optFns       func(*AWSConfigOptions) error
		expectString string
	}{
		"nil input": {
			optFns:       WithRegion("us-east-1"),
			expectString: "us-east-1",
		},
		"input with no filter": {
			optFns:       WithRegion("eu-west-1"),
			expectString: "eu-west-1",
		},
	}

	// region in the config file should equal to expectString
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			cfg, err := LoadAWSConfig(c.optFns)
			if err != nil {
				t.Errorf("LoadAWSConfig(%T) = %v, want %v", c.optFns, err, nil)
			}
			if cfg.Region != c.expectString {
				t.Errorf("AWSConfig.Region = %q, want %q", cfg.Region, c.expectString)
			}
		})
	}
}
