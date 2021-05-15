package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListClusterServiceComponentHealthInfo invokes the emr.ListClusterServiceComponentHealthInfo API synchronously
func (client *Client) ListClusterServiceComponentHealthInfo(request *ListClusterServiceComponentHealthInfoRequest) (response *ListClusterServiceComponentHealthInfoResponse, err error) {
	response = CreateListClusterServiceComponentHealthInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterServiceComponentHealthInfoWithChan invokes the emr.ListClusterServiceComponentHealthInfo API asynchronously
func (client *Client) ListClusterServiceComponentHealthInfoWithChan(request *ListClusterServiceComponentHealthInfoRequest) (<-chan *ListClusterServiceComponentHealthInfoResponse, <-chan error) {
	responseChan := make(chan *ListClusterServiceComponentHealthInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterServiceComponentHealthInfo(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListClusterServiceComponentHealthInfoWithCallback invokes the emr.ListClusterServiceComponentHealthInfo API asynchronously
func (client *Client) ListClusterServiceComponentHealthInfoWithCallback(request *ListClusterServiceComponentHealthInfoRequest, callback func(response *ListClusterServiceComponentHealthInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterServiceComponentHealthInfoResponse
		var err error
		defer close(result)
		response, err = client.ListClusterServiceComponentHealthInfo(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListClusterServiceComponentHealthInfoRequest is the request struct for api ListClusterServiceComponentHealthInfo
type ListClusterServiceComponentHealthInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	ServiceName     string           `position:"Query" name:"ServiceName"`
}

// ListClusterServiceComponentHealthInfoResponse is the response struct for api ListClusterServiceComponentHealthInfo
type ListClusterServiceComponentHealthInfoResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	ClusterId      string         `json:"ClusterId" xml:"ClusterId"`
	HealthInfoList HealthInfoList `json:"HealthInfoList" xml:"HealthInfoList"`
}

// CreateListClusterServiceComponentHealthInfoRequest creates a request to invoke ListClusterServiceComponentHealthInfo API
func CreateListClusterServiceComponentHealthInfoRequest() (request *ListClusterServiceComponentHealthInfoRequest) {
	request = &ListClusterServiceComponentHealthInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceComponentHealthInfo", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterServiceComponentHealthInfoResponse creates a response to parse from ListClusterServiceComponentHealthInfo response
func CreateListClusterServiceComponentHealthInfoResponse() (response *ListClusterServiceComponentHealthInfoResponse) {
	response = &ListClusterServiceComponentHealthInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
