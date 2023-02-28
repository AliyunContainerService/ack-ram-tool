package ecs

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

// PhysicalConnectionType is a nested struct in ecs response
type PhysicalConnectionType struct {
	AdLocation                    string `json:"AdLocation" xml:"AdLocation"`
	CreationTime                  string `json:"CreationTime" xml:"CreationTime"`
	Status                        string `json:"Status" xml:"Status"`
	Type                          string `json:"Type" xml:"Type"`
	PortNumber                    string `json:"PortNumber" xml:"PortNumber"`
	CircuitCode                   string `json:"CircuitCode" xml:"CircuitCode"`
	Spec                          string `json:"Spec" xml:"Spec"`
	Bandwidth                     int64  `json:"Bandwidth" xml:"Bandwidth"`
	Description                   string `json:"Description" xml:"Description"`
	PortType                      string `json:"PortType" xml:"PortType"`
	EnabledTime                   string `json:"EnabledTime" xml:"EnabledTime"`
	BusinessStatus                string `json:"BusinessStatus" xml:"BusinessStatus"`
	LineOperator                  string `json:"LineOperator" xml:"LineOperator"`
	Name                          string `json:"Name" xml:"Name"`
	RedundantPhysicalConnectionId string `json:"RedundantPhysicalConnectionId" xml:"RedundantPhysicalConnectionId"`
	PeerLocation                  string `json:"PeerLocation" xml:"PeerLocation"`
	AccessPointId                 string `json:"AccessPointId" xml:"AccessPointId"`
	PhysicalConnectionId          string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
}
