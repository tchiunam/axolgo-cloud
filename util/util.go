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
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"k8s.io/klog/v2"
)

// AWSConfigOptionsFunc is a type alias for AWSConfigOptions functional option
type AWSConfigOptionsFunc func(*AWSConfigOptions) error

// AWSConfigOptions are discrete set of options that are valid for loading the
// configuration that is supported Axolgo
type AWSConfigOptions struct {
	Region string
}

// Load AWS configuration
func LoadAWSConfig(optFns ...func(*AWSConfigOptions) error) (aws.Config, error) {
	var (
		options AWSConfigOptions
		cfg     = aws.Config{}
		err     error
	)
	for _, optFn := range optFns {
		if err = optFn(&options); err != nil {
			return cfg, fmt.Errorf("Fail to configure AWS config options: %v", err)
		}
	}

	if options.Region == "" {
		options.Region = getDefaultRegion()
		klog.V(1).InfoS("use default AWS region", "region", options.Region)
	}

	cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(options.Region))
	if err != nil {
		return cfg, fmt.Errorf("Fail to load AWS configuration: %v", err)
	}

	return cfg, nil
}

// WithRegion is a helper function to construct functional options
// that sets Region on config's AWSConfigOptions.
// If multiple WithRegion calls are made, the last call overrides
// the previous call values.
func WithRegion(v string) AWSConfigOptionsFunc {
	return func(o *AWSConfigOptions) error {
		o.Region = v
		return nil
	}
}

// A default region to perform actions
func getDefaultRegion() string {
	return "ap-east-1"
}
