package ddoscoo

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

// DescribeWebCCRules invokes the ddoscoo.DescribeWebCCRules API synchronously
func (client *Client) DescribeWebCCRules(request *DescribeWebCCRulesRequest) (response *DescribeWebCCRulesResponse, err error) {
	response = CreateDescribeWebCCRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebCCRulesWithChan invokes the ddoscoo.DescribeWebCCRules API asynchronously
func (client *Client) DescribeWebCCRulesWithChan(request *DescribeWebCCRulesRequest) (<-chan *DescribeWebCCRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeWebCCRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebCCRules(request)
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

// DescribeWebCCRulesWithCallback invokes the ddoscoo.DescribeWebCCRules API asynchronously
func (client *Client) DescribeWebCCRulesWithCallback(request *DescribeWebCCRulesRequest, callback func(response *DescribeWebCCRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebCCRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebCCRules(request)
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

// DescribeWebCCRulesRequest is the request struct for api DescribeWebCCRules
type DescribeWebCCRulesRequest struct {
	*requests.RpcRequest
	IsOffset        requests.Boolean `position:"Query" name:"IsOffset"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        string           `position:"Query" name:"PageSize"`
	Offset          requests.Integer `position:"Query" name:"Offset"`
	Domain          string           `position:"Query" name:"Domain"`
}

// DescribeWebCCRulesResponse is the response struct for api DescribeWebCCRules
type DescribeWebCCRulesResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	TotalCount int64       `json:"TotalCount" xml:"TotalCount"`
	WebCCRules []WebCCRule `json:"WebCCRules" xml:"WebCCRules"`
}

// CreateDescribeWebCCRulesRequest creates a request to invoke DescribeWebCCRules API
func CreateDescribeWebCCRulesRequest() (request *DescribeWebCCRulesRequest) {
	request = &DescribeWebCCRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeWebCCRules", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeWebCCRulesResponse creates a response to parse from DescribeWebCCRules response
func CreateDescribeWebCCRulesResponse() (response *DescribeWebCCRulesResponse) {
	response = &DescribeWebCCRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
