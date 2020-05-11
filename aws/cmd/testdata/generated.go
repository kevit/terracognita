package reader

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Code generated by github.com/cycloidio/terracognita/aws/cmd; DO NOT EDIT

// Reader is the interface defining all methods that need to be implemented
//
// The next behavior commented in the below paragraph, applies to every method
// which clearly match what's explained, for the sake of not repeating the same,
// over and over.
// The most of the methods defined by this interface, return their results in a
// map. Those maps, have as keys, the AWS region which have been requested and
// the values are the items returned by AWS for such region.
// Because the methods may make calls to different regions, in case that there
// is an error on a region, the returned map won't have any entry for such
// region and such errors will be reported by the returned error, nonetheless
// the items, got from the successful requests to other regions, will be
// returned, with the meaning that the methods will return partial results, in
// case of errors.
// For avoiding by the callers the problem of if the returned map may be nil,
// the function will always return a map instance, which will be of length 0
// in case that there is not any successful request.
type Reader interface {
	// GetAccountID returns the current ID for the account used
	GetAccountID() string

	// GetRegion returns the currently used region for the Connector
	GetRegion() string

	// GetInstances returns all EC2 instances based on the input given.
	// Returned values are commented in the interface doc comment block.
	GetInstances(ctx context.Context, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)

	// DownloadObject downloads an object in a bucket based on the input given
	DownloadObject(ctx context.Context, w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (int64, error)
}

func (c *connector) GetInstances(ctx context.Context, input *ec2.DescribeInstancesInput) ([]*ec2.Instance, error) {
	if c.svc.ec2 == nil {
		c.svc.ec2 = ec2.New(c.svc.session)
	}

	opt := make([]*ec2.Instance, 0)

	hasNextToken := true
	for hasNextToken != nil {
		o, err := c.svc.ec2.DescribeInstancesWithContext(ctx, input)
		if err != nil {
			return nil, err
		}
		input.NextToken = o.NextToken
		hasNextToken = o.NextToken != nil

		opt = append(opt, o.Instances...)
	}

	return opt, nil
}
