package cms

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

// ResourceInDescribeAlertingMetricRuleResources is a nested struct in cms response
type ResourceInDescribeAlertingMetricRuleResources struct {
	RuleId         string `json:"RuleId" xml:"RuleId"`
	RuleName       string `json:"RuleName" xml:"RuleName"`
	GroupId        string `json:"GroupId" xml:"GroupId"`
	Resource       string `json:"Resource" xml:"Resource"`
	MetricName     string `json:"MetricName" xml:"MetricName"`
	Namespace      string `json:"Namespace" xml:"Namespace"`
	Enable         string `json:"Enable" xml:"Enable"`
	LastAlertTime  string `json:"LastAlertTime" xml:"LastAlertTime"`
	LastModifyTime string `json:"LastModifyTime" xml:"LastModifyTime"`
	StartTime      string `json:"StartTime" xml:"StartTime"`
	MetricValues   string `json:"MetricValues" xml:"MetricValues"`
	RetryTimes     string `json:"RetryTimes" xml:"RetryTimes"`
	Statistics     string `json:"Statistics" xml:"Statistics"`
	Threshold      string `json:"Threshold" xml:"Threshold"`
}
