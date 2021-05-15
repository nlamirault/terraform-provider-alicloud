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

// SubmitFlowJob invokes the emr.SubmitFlowJob API synchronously
func (client *Client) SubmitFlowJob(request *SubmitFlowJobRequest) (response *SubmitFlowJobResponse, err error) {
	response = CreateSubmitFlowJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitFlowJobWithChan invokes the emr.SubmitFlowJob API asynchronously
func (client *Client) SubmitFlowJobWithChan(request *SubmitFlowJobRequest) (<-chan *SubmitFlowJobResponse, <-chan error) {
	responseChan := make(chan *SubmitFlowJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitFlowJob(request)
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

// SubmitFlowJobWithCallback invokes the emr.SubmitFlowJob API asynchronously
func (client *Client) SubmitFlowJobWithCallback(request *SubmitFlowJobRequest, callback func(response *SubmitFlowJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitFlowJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitFlowJob(request)
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

// SubmitFlowJobRequest is the request struct for api SubmitFlowJob
type SubmitFlowJobRequest struct {
	*requests.RpcRequest
	Conf          string `position:"Query" name:"Conf"`
	ClusterId     string `position:"Query" name:"ClusterId"`
	JobId         string `position:"Query" name:"JobId"`
	HostName      string `position:"Query" name:"HostName"`
	Namespace     string `position:"Query" name:"Namespace"`
	JobInstanceId string `position:"Query" name:"JobInstanceId"`
	ProjectId     string `position:"Query" name:"ProjectId"`
}

// SubmitFlowJobResponse is the response struct for api SubmitFlowJob
type SubmitFlowJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateSubmitFlowJobRequest creates a request to invoke SubmitFlowJob API
func CreateSubmitFlowJobRequest() (request *SubmitFlowJobRequest) {
	request = &SubmitFlowJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "SubmitFlowJob", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitFlowJobResponse creates a response to parse from SubmitFlowJob response
func CreateSubmitFlowJobResponse() (response *SubmitFlowJobResponse) {
	response = &SubmitFlowJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
