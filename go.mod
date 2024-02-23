module github.com/AliyunContainerService/ack-ram-tool

go 1.20

require (
	github.com/AlecAivazis/survey/v2 v2.3.7
	github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider v0.0.0
	github.com/alibabacloud-go/cs-20151215/v3 v3.0.35
	github.com/alibabacloud-go/darabonba-openapi v0.2.1
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.0.2
	github.com/alibabacloud-go/ram-20150501 v1.0.2
	github.com/alibabacloud-go/sts-20150401 v1.1.0
	github.com/alibabacloud-go/tea v1.2.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.648
	github.com/aliyun/credentials-go v1.3.0
	github.com/briandowns/spinner v1.23.0
	github.com/fatih/color v1.16.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/sergi/go-diff v1.3.1
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.8.4
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.29.0
	k8s.io/apimachinery v0.29.0
	k8s.io/client-go v0.29.0
)

require (
	github.com/alibabacloud-go/openapi-util v0.1.0
	github.com/alibabacloud-go/tea-utils v1.4.5
	github.com/alibabacloud-go/tea-utils/v2 v2.0.1
	github.com/go-logr/zapr v1.3.0
	github.com/olekukonko/tablewriter v0.0.6-0.20230925090304-df64c4bbad77
	go.uber.org/zap v1.26.0
	k8s.io/klog/v2 v2.110.1
)

require (
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.4 // indirect
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/alibabacloud-go/endpoint-util v1.1.0 // indirect
	github.com/alibabacloud-go/tea-xml v1.1.2 // indirect
	github.com/clbanning/mxj/v2 v2.5.6 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tjfoc/gmsm v1.3.2 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.10.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/kube-openapi v0.0.0-20231010175941-2dd684a91f00 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider => ./pkg/credentials/provider
