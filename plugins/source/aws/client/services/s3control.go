// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

//go:generate mockgen -package=mocks -destination=../mocks/s3control.go -source=s3control.go S3controlClient
type S3controlClient interface {
	DescribeJob(context.Context, *s3control.DescribeJobInput, ...func(*s3control.Options)) (*s3control.DescribeJobOutput, error)
	DescribeMultiRegionAccessPointOperation(context.Context, *s3control.DescribeMultiRegionAccessPointOperationInput, ...func(*s3control.Options)) (*s3control.DescribeMultiRegionAccessPointOperationOutput, error)
	GetAccessGrant(context.Context, *s3control.GetAccessGrantInput, ...func(*s3control.Options)) (*s3control.GetAccessGrantOutput, error)
	GetAccessGrantsInstance(context.Context, *s3control.GetAccessGrantsInstanceInput, ...func(*s3control.Options)) (*s3control.GetAccessGrantsInstanceOutput, error)
	GetAccessGrantsInstanceForPrefix(context.Context, *s3control.GetAccessGrantsInstanceForPrefixInput, ...func(*s3control.Options)) (*s3control.GetAccessGrantsInstanceForPrefixOutput, error)
	GetAccessGrantsInstanceResourcePolicy(context.Context, *s3control.GetAccessGrantsInstanceResourcePolicyInput, ...func(*s3control.Options)) (*s3control.GetAccessGrantsInstanceResourcePolicyOutput, error)
	GetAccessGrantsLocation(context.Context, *s3control.GetAccessGrantsLocationInput, ...func(*s3control.Options)) (*s3control.GetAccessGrantsLocationOutput, error)
	GetAccessPoint(context.Context, *s3control.GetAccessPointInput, ...func(*s3control.Options)) (*s3control.GetAccessPointOutput, error)
	GetAccessPointConfigurationForObjectLambda(context.Context, *s3control.GetAccessPointConfigurationForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error)
	GetAccessPointForObjectLambda(context.Context, *s3control.GetAccessPointForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.GetAccessPointForObjectLambdaOutput, error)
	GetAccessPointPolicy(context.Context, *s3control.GetAccessPointPolicyInput, ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyOutput, error)
	GetAccessPointPolicyForObjectLambda(context.Context, *s3control.GetAccessPointPolicyForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error)
	GetAccessPointPolicyStatus(context.Context, *s3control.GetAccessPointPolicyStatusInput, ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusOutput, error)
	GetAccessPointPolicyStatusForObjectLambda(context.Context, *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error)
	GetBucket(context.Context, *s3control.GetBucketInput, ...func(*s3control.Options)) (*s3control.GetBucketOutput, error)
	GetBucketLifecycleConfiguration(context.Context, *s3control.GetBucketLifecycleConfigurationInput, ...func(*s3control.Options)) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	GetBucketPolicy(context.Context, *s3control.GetBucketPolicyInput, ...func(*s3control.Options)) (*s3control.GetBucketPolicyOutput, error)
	GetBucketReplication(context.Context, *s3control.GetBucketReplicationInput, ...func(*s3control.Options)) (*s3control.GetBucketReplicationOutput, error)
	GetBucketTagging(context.Context, *s3control.GetBucketTaggingInput, ...func(*s3control.Options)) (*s3control.GetBucketTaggingOutput, error)
	GetBucketVersioning(context.Context, *s3control.GetBucketVersioningInput, ...func(*s3control.Options)) (*s3control.GetBucketVersioningOutput, error)
	GetDataAccess(context.Context, *s3control.GetDataAccessInput, ...func(*s3control.Options)) (*s3control.GetDataAccessOutput, error)
	GetJobTagging(context.Context, *s3control.GetJobTaggingInput, ...func(*s3control.Options)) (*s3control.GetJobTaggingOutput, error)
	GetMultiRegionAccessPoint(context.Context, *s3control.GetMultiRegionAccessPointInput, ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointOutput, error)
	GetMultiRegionAccessPointPolicy(context.Context, *s3control.GetMultiRegionAccessPointPolicyInput, ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointPolicyOutput, error)
	GetMultiRegionAccessPointPolicyStatus(context.Context, *s3control.GetMultiRegionAccessPointPolicyStatusInput, ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointPolicyStatusOutput, error)
	GetMultiRegionAccessPointRoutes(context.Context, *s3control.GetMultiRegionAccessPointRoutesInput, ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointRoutesOutput, error)
	GetPublicAccessBlock(context.Context, *s3control.GetPublicAccessBlockInput, ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error)
	GetStorageLensConfiguration(context.Context, *s3control.GetStorageLensConfigurationInput, ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationOutput, error)
	GetStorageLensConfigurationTagging(context.Context, *s3control.GetStorageLensConfigurationTaggingInput, ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationTaggingOutput, error)
	GetStorageLensGroup(context.Context, *s3control.GetStorageLensGroupInput, ...func(*s3control.Options)) (*s3control.GetStorageLensGroupOutput, error)
	ListAccessGrants(context.Context, *s3control.ListAccessGrantsInput, ...func(*s3control.Options)) (*s3control.ListAccessGrantsOutput, error)
	ListAccessGrantsInstances(context.Context, *s3control.ListAccessGrantsInstancesInput, ...func(*s3control.Options)) (*s3control.ListAccessGrantsInstancesOutput, error)
	ListAccessGrantsLocations(context.Context, *s3control.ListAccessGrantsLocationsInput, ...func(*s3control.Options)) (*s3control.ListAccessGrantsLocationsOutput, error)
	ListAccessPoints(context.Context, *s3control.ListAccessPointsInput, ...func(*s3control.Options)) (*s3control.ListAccessPointsOutput, error)
	ListAccessPointsForObjectLambda(context.Context, *s3control.ListAccessPointsForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.ListAccessPointsForObjectLambdaOutput, error)
	ListJobs(context.Context, *s3control.ListJobsInput, ...func(*s3control.Options)) (*s3control.ListJobsOutput, error)
	ListMultiRegionAccessPoints(context.Context, *s3control.ListMultiRegionAccessPointsInput, ...func(*s3control.Options)) (*s3control.ListMultiRegionAccessPointsOutput, error)
	ListRegionalBuckets(context.Context, *s3control.ListRegionalBucketsInput, ...func(*s3control.Options)) (*s3control.ListRegionalBucketsOutput, error)
	ListStorageLensConfigurations(context.Context, *s3control.ListStorageLensConfigurationsInput, ...func(*s3control.Options)) (*s3control.ListStorageLensConfigurationsOutput, error)
	ListStorageLensGroups(context.Context, *s3control.ListStorageLensGroupsInput, ...func(*s3control.Options)) (*s3control.ListStorageLensGroupsOutput, error)
	ListTagsForResource(context.Context, *s3control.ListTagsForResourceInput, ...func(*s3control.Options)) (*s3control.ListTagsForResourceOutput, error)
}
