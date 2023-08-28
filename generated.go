// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// GetR2MetricsResponse is returned by GetR2Metrics on success.
type GetR2MetricsResponse struct {
	Viewer GetR2MetricsViewer `json:"viewer"`
}

// GetViewer returns GetR2MetricsResponse.Viewer, and is useful for accessing the field via an interface.
func (v *GetR2MetricsResponse) GetViewer() GetR2MetricsViewer { return v.Viewer }

// GetR2MetricsStorageResponse is returned by GetR2MetricsStorage on success.
type GetR2MetricsStorageResponse struct {
	Viewer GetR2MetricsStorageViewer `json:"viewer"`
}

// GetViewer returns GetR2MetricsStorageResponse.Viewer, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageResponse) GetViewer() GetR2MetricsStorageViewer { return v.Viewer }

// GetR2MetricsStorageViewer includes the requested fields of the GraphQL type viewer.
type GetR2MetricsStorageViewer struct {
	Accounts []GetR2MetricsStorageViewerAccountsAccount `json:"accounts"`
}

// GetAccounts returns GetR2MetricsStorageViewer.Accounts, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewer) GetAccounts() []GetR2MetricsStorageViewerAccountsAccount {
	return v.Accounts
}

// GetR2MetricsStorageViewerAccountsAccount includes the requested fields of the GraphQL type account.
type GetR2MetricsStorageViewerAccountsAccount struct {
	// Beta. R2 storage with adaptive sampling
	Storage []GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroups `json:"storage"`
}

// GetStorage returns GetR2MetricsStorageViewerAccountsAccount.Storage, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewerAccountsAccount) GetStorage() []GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroups {
	return v.Storage
}

// GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroups includes the requested fields of the GraphQL type AccountR2StorageAdaptiveGroups.
// The GraphQL type's documentation follows.
//
// Beta. R2 storage with adaptive sampling
type GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroups struct {
	// List of dimensions to group by
	Dimensions GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsDimensions `json:"dimensions"`
	// The max of values for a metric per dimension
	Max GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax `json:"max"`
}

// GetDimensions returns GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroups.Dimensions, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroups) GetDimensions() GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsDimensions {
	return v.Dimensions
}

// GetMax returns GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroups.Max, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroups) GetMax() GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax {
	return v.Max
}

// GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsDimensions includes the requested fields of the GraphQL type AccountR2StorageAdaptiveGroupsDimensions.
type GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsDimensions struct {
	// The name of the R2 bucket, if applicable to this request
	BucketName string `json:"bucketName"`
}

// GetBucketName returns GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsDimensions.BucketName, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsDimensions) GetBucketName() string {
	return v.BucketName
}

// GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax includes the requested fields of the GraphQL type AccountR2StorageAdaptiveGroupsMax.
type GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax struct {
	// Max of metadata size
	MetadataSize uint64 `json:"metadataSize"`
	// Max of object count
	ObjectCount uint64 `json:"objectCount"`
	// Max of payload size
	PayloadSize uint64 `json:"payloadSize"`
	// Max of upload count
	UploadCount uint64 `json:"uploadCount"`
}

// GetMetadataSize returns GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax.MetadataSize, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax) GetMetadataSize() uint64 {
	return v.MetadataSize
}

// GetObjectCount returns GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax.ObjectCount, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax) GetObjectCount() uint64 {
	return v.ObjectCount
}

// GetPayloadSize returns GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax.PayloadSize, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax) GetPayloadSize() uint64 {
	return v.PayloadSize
}

// GetUploadCount returns GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax.UploadCount, and is useful for accessing the field via an interface.
func (v *GetR2MetricsStorageViewerAccountsAccountStorageAccountR2StorageAdaptiveGroupsMax) GetUploadCount() uint64 {
	return v.UploadCount
}

// GetR2MetricsViewer includes the requested fields of the GraphQL type viewer.
type GetR2MetricsViewer struct {
	Accounts []GetR2MetricsViewerAccountsAccount `json:"accounts"`
}

// GetAccounts returns GetR2MetricsViewer.Accounts, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewer) GetAccounts() []GetR2MetricsViewerAccountsAccount { return v.Accounts }

// GetR2MetricsViewerAccountsAccount includes the requested fields of the GraphQL type account.
type GetR2MetricsViewerAccountsAccount struct {
	// Beta. R2 operations with adaptive sampling
	ClassA []GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroups `json:"classA"`
	// Beta. R2 operations with adaptive sampling
	ClassB []GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroups `json:"classB"`
	// Beta. R2 operations with adaptive sampling
	ClassAperBucket []GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroups `json:"classAperBucket"`
	// Beta. R2 operations with adaptive sampling
	ClassBperBucket []GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroups `json:"classBperBucket"`
}

// GetClassA returns GetR2MetricsViewerAccountsAccount.ClassA, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccount) GetClassA() []GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroups {
	return v.ClassA
}

// GetClassB returns GetR2MetricsViewerAccountsAccount.ClassB, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccount) GetClassB() []GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroups {
	return v.ClassB
}

// GetClassAperBucket returns GetR2MetricsViewerAccountsAccount.ClassAperBucket, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccount) GetClassAperBucket() []GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroups {
	return v.ClassAperBucket
}

// GetClassBperBucket returns GetR2MetricsViewerAccountsAccount.ClassBperBucket, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccount) GetClassBperBucket() []GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroups {
	return v.ClassBperBucket
}

// GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroups includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroups.
// The GraphQL type's documentation follows.
//
// Beta. R2 operations with adaptive sampling
type GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroups struct {
	// List of dimensions to group by
	Dimensions GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsDimensions `json:"dimensions"`
	// The sum of values for a metric per dimension
	Sum GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsSum `json:"sum"`
}

// GetDimensions returns GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroups.Dimensions, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroups) GetDimensions() GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsDimensions {
	return v.Dimensions
}

// GetSum returns GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroups.Sum, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroups) GetSum() GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsSum {
	return v.Sum
}

// GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsDimensions includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroupsDimensions.
type GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsDimensions struct {
	// The name of the R2 operation
	ActionType string `json:"actionType"`
}

// GetActionType returns GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsDimensions.ActionType, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsDimensions) GetActionType() string {
	return v.ActionType
}

// GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsSum includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroupsSum.
type GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsSum struct {
	// Sum of Requests
	Requests uint64 `json:"requests"`
	// Sum of Response Object Sizes
	ResponseObjectSize uint64 `json:"responseObjectSize"`
}

// GetRequests returns GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsSum.Requests, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsSum) GetRequests() uint64 {
	return v.Requests
}

// GetResponseObjectSize returns GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsSum.ResponseObjectSize, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAAccountR2OperationsAdaptiveGroupsSum) GetResponseObjectSize() uint64 {
	return v.ResponseObjectSize
}

// GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroups includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroups.
// The GraphQL type's documentation follows.
//
// Beta. R2 operations with adaptive sampling
type GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroups struct {
	// List of dimensions to group by
	Dimensions GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsDimensions `json:"dimensions"`
	// The sum of values for a metric per dimension
	Sum GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsSum `json:"sum"`
}

// GetDimensions returns GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroups.Dimensions, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroups) GetDimensions() GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsDimensions {
	return v.Dimensions
}

// GetSum returns GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroups.Sum, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroups) GetSum() GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsSum {
	return v.Sum
}

// GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsDimensions includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroupsDimensions.
type GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsDimensions struct {
	// The name of the R2 bucket, if applicable to this request
	BucketName string `json:"bucketName"`
}

// GetBucketName returns GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsDimensions.BucketName, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsDimensions) GetBucketName() string {
	return v.BucketName
}

// GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsSum includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroupsSum.
type GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsSum struct {
	// Sum of Requests
	Requests uint64 `json:"requests"`
}

// GetRequests returns GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsSum.Requests, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassAperBucketAccountR2OperationsAdaptiveGroupsSum) GetRequests() uint64 {
	return v.Requests
}

// GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroups includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroups.
// The GraphQL type's documentation follows.
//
// Beta. R2 operations with adaptive sampling
type GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroups struct {
	// List of dimensions to group by
	Dimensions GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsDimensions `json:"dimensions"`
	// The sum of values for a metric per dimension
	Sum GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsSum `json:"sum"`
}

// GetDimensions returns GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroups.Dimensions, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroups) GetDimensions() GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsDimensions {
	return v.Dimensions
}

// GetSum returns GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroups.Sum, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroups) GetSum() GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsSum {
	return v.Sum
}

// GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsDimensions includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroupsDimensions.
type GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsDimensions struct {
	// The name of the R2 operation
	ActionType string `json:"actionType"`
}

// GetActionType returns GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsDimensions.ActionType, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsDimensions) GetActionType() string {
	return v.ActionType
}

// GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsSum includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroupsSum.
type GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsSum struct {
	// Sum of Requests
	Requests uint64 `json:"requests"`
	// Sum of Response Object Sizes
	ResponseObjectSize uint64 `json:"responseObjectSize"`
}

// GetRequests returns GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsSum.Requests, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsSum) GetRequests() uint64 {
	return v.Requests
}

// GetResponseObjectSize returns GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsSum.ResponseObjectSize, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBAccountR2OperationsAdaptiveGroupsSum) GetResponseObjectSize() uint64 {
	return v.ResponseObjectSize
}

// GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroups includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroups.
// The GraphQL type's documentation follows.
//
// Beta. R2 operations with adaptive sampling
type GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroups struct {
	// List of dimensions to group by
	Dimensions GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsDimensions `json:"dimensions"`
	// The sum of values for a metric per dimension
	Sum GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsSum `json:"sum"`
}

// GetDimensions returns GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroups.Dimensions, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroups) GetDimensions() GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsDimensions {
	return v.Dimensions
}

// GetSum returns GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroups.Sum, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroups) GetSum() GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsSum {
	return v.Sum
}

// GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsDimensions includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroupsDimensions.
type GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsDimensions struct {
	// The name of the R2 bucket, if applicable to this request
	BucketName string `json:"bucketName"`
}

// GetBucketName returns GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsDimensions.BucketName, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsDimensions) GetBucketName() string {
	return v.BucketName
}

// GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsSum includes the requested fields of the GraphQL type AccountR2OperationsAdaptiveGroupsSum.
type GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsSum struct {
	// Sum of Requests
	Requests uint64 `json:"requests"`
}

// GetRequests returns GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsSum.Requests, and is useful for accessing the field via an interface.
func (v *GetR2MetricsViewerAccountsAccountClassBperBucketAccountR2OperationsAdaptiveGroupsSum) GetRequests() uint64 {
	return v.Requests
}

// __GetR2MetricsInput is used internally by genqlient
type __GetR2MetricsInput struct {
	AccountTag string `json:"accountTag"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}

// GetAccountTag returns __GetR2MetricsInput.AccountTag, and is useful for accessing the field via an interface.
func (v *__GetR2MetricsInput) GetAccountTag() string { return v.AccountTag }

// GetStartDate returns __GetR2MetricsInput.StartDate, and is useful for accessing the field via an interface.
func (v *__GetR2MetricsInput) GetStartDate() string { return v.StartDate }

// GetEndDate returns __GetR2MetricsInput.EndDate, and is useful for accessing the field via an interface.
func (v *__GetR2MetricsInput) GetEndDate() string { return v.EndDate }

// __GetR2MetricsStorageInput is used internally by genqlient
type __GetR2MetricsStorageInput struct {
	AccountTag string `json:"accountTag"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}

// GetAccountTag returns __GetR2MetricsStorageInput.AccountTag, and is useful for accessing the field via an interface.
func (v *__GetR2MetricsStorageInput) GetAccountTag() string { return v.AccountTag }

// GetStartDate returns __GetR2MetricsStorageInput.StartDate, and is useful for accessing the field via an interface.
func (v *__GetR2MetricsStorageInput) GetStartDate() string { return v.StartDate }

// GetEndDate returns __GetR2MetricsStorageInput.EndDate, and is useful for accessing the field via an interface.
func (v *__GetR2MetricsStorageInput) GetEndDate() string { return v.EndDate }

// The query or mutation executed by GetR2Metrics.
const GetR2Metrics_Operation = `
query GetR2Metrics ($accountTag: string!, $startDate: Time!, $endDate: Time!) {
	viewer {
		accounts(filter: {accountTag:$accountTag}) {
			classA: r2OperationsAdaptiveGroups(limit: 1000, filter: {datetime_geq:$startDate,datetime_lt:$endDate,actionStatus:"success",actionType_in:["ListBuckets","PutBucket","ListObjects","PutObject","CopyObject","CompleteMultipartUpload","CreateMultipartUpload","UploadPart","UploadPartCopy","PutBucketEncryption","ListMultipartUploads"]}) {
				dimensions {
					actionType
				}
				sum {
					requests
					responseObjectSize
				}
			}
			classB: r2OperationsAdaptiveGroups(limit: 1000, filter: {datetime_geq:$startDate,datetime_lt:$endDate,actionStatus:"success",actionType_in:["HeadBucket","HeadObject","GetObject","ReportUsageSummary","GetBucketEncryption","GetBucketLocation"]}) {
				dimensions {
					actionType
				}
				sum {
					requests
					responseObjectSize
				}
			}
			classAperBucket: r2OperationsAdaptiveGroups(limit: 1000, filter: {datetime_geq:$startDate,datetime_lt:$endDate,actionStatus:"success",actionType_in:["ListBuckets","PutBucket","ListObjects","PutObject","CopyObject","CompleteMultipartUpload","CreateMultipartUpload","UploadPart","UploadPartCopy","PutBucketEncryption","ListMultipartUploads"]}) {
				dimensions {
					bucketName
				}
				sum {
					requests
				}
			}
			classBperBucket: r2OperationsAdaptiveGroups(limit: 1000, filter: {datetime_geq:$startDate,datetime_lt:$endDate,actionStatus:"success",actionType_in:["HeadBucket","HeadObject","GetObject","ReportUsageSummary","GetBucketEncryption","GetBucketLocation"]}) {
				dimensions {
					bucketName
				}
				sum {
					requests
				}
			}
		}
	}
}
`

func GetR2Metrics(
	ctx context.Context,
	client graphql.Client,
	accountTag string,
	startDate string,
	endDate string,
) (*GetR2MetricsResponse, error) {
	req := &graphql.Request{
		OpName: "GetR2Metrics",
		Query:  GetR2Metrics_Operation,
		Variables: &__GetR2MetricsInput{
			AccountTag: accountTag,
			StartDate:  startDate,
			EndDate:    endDate,
		},
	}
	var err error

	var data GetR2MetricsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by GetR2MetricsStorage.
const GetR2MetricsStorage_Operation = `
query GetR2MetricsStorage ($accountTag: string!, $startDate: Time!, $endDate: Time!) {
	viewer {
		accounts(filter: {accountTag:$accountTag}) {
			storage: r2StorageAdaptiveGroups(orderBy: [bucketName_ASC], limit: 1000, filter: {datetime_geq:$startDate,datetime_lt:$endDate}) {
				dimensions {
					bucketName
				}
				max {
					metadataSize
					objectCount
					payloadSize
					uploadCount
				}
			}
		}
	}
}
`

func GetR2MetricsStorage(
	ctx context.Context,
	client graphql.Client,
	accountTag string,
	startDate string,
	endDate string,
) (*GetR2MetricsStorageResponse, error) {
	req := &graphql.Request{
		OpName: "GetR2MetricsStorage",
		Query:  GetR2MetricsStorage_Operation,
		Variables: &__GetR2MetricsStorageInput{
			AccountTag: accountTag,
			StartDate:  startDate,
			EndDate:    endDate,
		},
	}
	var err error

	var data GetR2MetricsStorageResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
