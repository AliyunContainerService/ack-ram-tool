module github.com/AliyunContainerService/ack-ram-tool

go 1.16

require (
	github.com/AlecAivazis/survey/v2 v2.3.6
	github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper v0.0.0
	github.com/alibabacloud-go/cs-20151215/v3 v3.0.9
	github.com/alibabacloud-go/darabonba-openapi v0.1.15
	github.com/alibabacloud-go/ram-20150501 v1.0.1
	github.com/alibabacloud-go/sts-20150401 v1.1.0
	github.com/alibabacloud-go/tea v1.1.19
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1518
	github.com/aliyun/credentials-go v1.2.4
	github.com/briandowns/spinner v1.19.0
	github.com/fatih/color v1.13.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/sergi/go-diff v1.2.0
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.8.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper v0.0.0 => ./pkg/credentials/alibabacloudsdkgo/helper
