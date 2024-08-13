package scanaddon

import (
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WorkloadType string

var (
	WorkloadTypeDeployment WorkloadType = "Deployment"
	WorkloadTypeDaemonSet  WorkloadType = "DaemonSet"
)

type Workload struct {
	Type         WorkloadType
	Namespace    string
	Name         string
	Images       []string
	ImageNames   []string
	MountedNames []string
	Hardened     []string
	CreateTime   metav1.Time
	ReadyTime    metav1.Time
}

type Addon struct {
	Name             string
	MinVersion       string
	ImageNames       []string
	DefaultRoleNames []string
}

type InstalledAddon struct {
	Name           string
	CurrentVersion string
}

type RolePolicy struct {
	Name    string
	Actions []string
}

type NeedUpdateAddon struct {
	Addon     Addon
	Workloads []Workload
	RoleNames []string
}

var CheckAddons = []Addon{
	{
		Name:       "metrics-server",
		MinVersion: "v0.3.9.4-ff225cd-aliyun",
		ImageNames: []string{
			"metrics-server",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCmsRole",
		},
	},
	{
		Name:       "alicloud-monitor-controller",
		MinVersion: "v1.5.5",
		ImageNames: []string{
			"alicloud-monitor-controller",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCmsRole",
			"AliyunCSManagedLogRole",
			"AliyunCSManagedArmsRole",
		},
	},
	{
		Name:       "logtail-ds",
		MinVersion: "v1.0.29.1-0550501-aliyun",
		ImageNames: []string{
			"log-controller",
			"logtail",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedLogRole",
		},
	},
	{
		Name:       "terway",
		MinVersion: "v1.0.10.333-gfd2b7b8-aliyun",
		ImageNames: []string{
			"terway",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedNetworkRole",
		},
	},
	{
		Name:       "terway-eni",
		MinVersion: "v1.0.10.333-gfd2b7b8-aliyun",
		ImageNames: []string{
			"terway",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedNetworkRole",
		},
	},
	{
		Name:       "terway-eniip",
		MinVersion: "v1.0.10.333-gfd2b7b8-aliyun",
		ImageNames: []string{
			"terway",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedNetworkRole",
		},
	},
	{
		Name:       "terway-controlplane",
		MinVersion: "v1.2.1",
		ImageNames: []string{
			"terway-controlplane",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedNetworkRole",
		},
	},
	{
		Name:       "flexvolume",
		MinVersion: "v1.14.8.109-649dc5a-aliyun",
		ImageNames: []string{
			"flexvolume",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCsiRole",
		},
	},
	{
		Name:       "csi-provisioner",
		MinVersion: "v1.18.8.45-1c5d2cd1-aliyun",
		ImageNames: []string{
			"csi-provisioner",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCsiRole",
		},
	},
	{
		Name:       "csi-plugin",
		MinVersion: "v1.18.8.45-1c5d2cd1-aliyun",
		ImageNames: []string{
			"csi-plugin",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCsiRole",
		},
	},
	{
		Name:       "storage-operator",
		MinVersion: "v1.18.8.55-e398ce5-aliyun",
		ImageNames: []string{
			"storage-auto-expander",
			"storage-cnfs",
			"storage-controller",
			"storage-monitor",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCsiRole",
		},
	},
	{
		Name:       "alicloud-disk-controller",
		MinVersion: "v1.14.8.51-842f0a81-aliyun",
		ImageNames: []string{
			"alicloud-disk-controller",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCsiRole",
		},
	},
	{
		Name:       "ack-node-problem-detector",
		MinVersion: "1.2.16",
		ImageNames: []string{
			"kube-eventer",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedLogRole",
		},
	},
	{
		Name:       "aliyun-acr-credential-helper",
		MinVersion: "v23.02.06.2-74e2172-aliyun",
		ImageNames: []string{
			"aliyun-acr-credential-helper",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedAcrRole",
		},
	},
	{
		Name:       "ack-cost-exporter",
		MinVersion: "1.0.10",
		ImageNames: []string{
			"alibaba-cloud-price-exporter",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCostRole",
		},
	},
	{
		Name:       "mse-ingress-controller",
		MinVersion: "1.1.5",
		ImageNames: []string{
			"mse-ingress-controller",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedMseRole",
		},
	},
	{
		Name:       "arms-prometheus",
		MinVersion: "1.1.11",
		ImageNames: []string{
			"arms-prometheus-agent",
			"o11y-addon-controller",
		},
		DefaultRoleNames: []string{
			//"AliyunCSManagedArmsRole",
		},
	},
	{
		Name:       "arms-cmonitor",
		MinVersion: "4.0.0",
		ImageNames: []string{
			"cmonitor-agent",
			"alicollector",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedArmsRole",
			//"AliyunCSManagedMseRole",
		},
	},
	{
		Name:       "ack-onepilot",
		MinVersion: "3.0.11",
		ImageNames: []string{
			"ack-onepilot",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedArmsRole",
			"AliyunCSManagedMseRole",
		},
	},
	{
		Name:       "cluster-autoscaler",
		MinVersion: "v1.3.1-bcf13de9-aliyun",
		ImageNames: []string{
			"autoscaler",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedAutoScalerRole",
		},
	},
	{
		Name:       "ack-goatscaler",
		MinVersion: "v1.3.1-bcf13de9-aliyun",
		ImageNames: []string{
			"goatscaler",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedAutoScalerRole",
		},
	},
	{
		Name:       "migrate-controller",
		MinVersion: "v1.8.1-187f707-aliyun",
		ImageNames: []string{
			"velero-installer",
			"velero-plugin-alibabacloud",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedBackupRestoreRole",
		},
	},
	{
		Name:       "ack-alibaba-cloud-metrics-adapter",
		MinVersion: "v0.2.7-f1ee5c3-aliyun",
		ImageNames: []string{
			"alibaba-cloud-metrics-adapter-amd64",
		},
		DefaultRoleNames: []string{
			"AliyunCSManagedCmsRole",
		},
	},
	{
		Name:       "kube-ai",
		MinVersion: "v0.1.0",
		ImageNames: []string{
			"ai-dashboard",
			"kubeai-dev-console",
			"commit-agent",
		},
		DefaultRoleNames: []string{
			"kube-ai",
		},
	},
}

func getAddonByWorkload(wl Workload, installedAddons map[string]InstalledAddon) (Addon, bool) {
	var addons []Addon
	for _, ad := range CheckAddons {
		for _, name := range wl.ImageNames {
			if utils.StringSliceInclude(ad.ImageNames, name) {
				addons = append(addons, ad)
			}
		}
	}
	if len(addons) == 0 {
		return Addon{}, false
	}
	for _, ad := range addons {
		if _, ok := installedAddons[ad.Name]; ok {
			return ad, true
		}
	}

	return addons[0], true
}

var defaultPoliciesByRoleMap = map[string]RolePolicy{
	"AliyunCSManagedCmsRole": {
		Name: "AliyunCSManagedCmsRolePolicy",
		Actions: []string{
			"cms:DescribeMonitorGroups",
			"cms:DescribeMonitorGroupInstances",
			"cms:CreateMonitorGroup",
			"cms:DeleteMonitorGroup",
			"cms:ModifyMonitorGroupInstances",
			"cms:CreateMonitorGroupInstances",
			"cms:DeleteMonitorGroupInstances",
			"cms:TaskConfigCreate",
			"cms:TaskConfigList",
			"cms:DescribeMetricList",
			"cms:QueryMetricList",
			"cms:CreateDynamicTagGroup",
			"cms:PutGroupMetricRule",
			"cms:DescribeMetricRuleList",
			"cms:DeleteMetricRules",
			"cs:DescribeMonitorToken",
			"ahas:GetSentinelAppSumMetric",
			"log:GetLogStoreLogs",
			"slb:DescribeMetricList",
			"sls:GetLogs",
			"sls:PutLogs",
		},
	},
	"AliyunCSManagedLogRole": {
		Name: "AliyunCSManagedLogRolePolicy",
		Actions: []string{
			"log:CreateProject",
			"log:GetProject",
			"log:DeleteProject",
			"log:CreateLogStore",
			"log:GetLogStore",
			"log:UpdateLogStore",
			"log:DeleteLogStore",
			"log:CreateConfig",
			"log:UpdateConfig",
			"log:GetConfig",
			"log:DeleteConfig",
			"log:CreateMachineGroup",
			"log:UpdateMachineGroup",
			"log:GetMachineGroup",
			"log:DeleteMachineGroup",
			"log:ApplyConfigToGroup",
			"log:GetAppliedMachineGroups",
			"log:GetAppliedConfigs",
			"log:RemoveConfigFromMachineGroup",
			"log:RemoveConfigFromGroup",
			"log:CreateIndex",
			"log:GetIndex",
			"log:UpdateIndex",
			"log:DeleteIndex",
			"log:CreateSavedSearch",
			"log:GetSavedSearch",
			"log:UpdateSavedSearch",
			"log:DeleteSavedSearch",
			"log:CreateDashboard",
			"log:GetDashboard",
			"log:UpdateDashboard",
			"log:DeleteDashboard",
			"log:CreateJob",
			"log:GetJob",
			"log:DeleteJob",
			"log:UpdateJob",
			"log:PostLogStoreLogs",
			"log:CreateSortedSubStore",
			"log:GetSortedSubStore",
			"log:ListSortedSubStore",
			"log:UpdateSortedSubStore",
			"log:DeleteSortedSubStore",
			"log:CreateApp",
			"log:UpdateApp",
			"log:GetApp",
			"log:DeleteApp",
			"log:GetLogStoreLogs",
			"log:TagResources",
			"log:ListJobs",
			"log:ListTagResources",
			"log:UntagResources",
			"log:CreateResourceRecord",
			"log:UpdateResourceRecord",
			"log:UpsertResourceRecord",
			"log:GetResourceRecord",
			"log:DeleteResourceRecord",
			"log:ListResourceRecords",
			"log:ListResources",
			"log:GetResource",
			"log:PutLogs",
			"log:UpdateLogStoreMeteringMode",
			"log:GetLogStoreMeteringMode",
			"log:CreateLogtailPipelineConfig",
			"log:DeleteLogtailPipelineConfig",
			"log:GetLogtailPipelineConfig",
			"log:UpdateLogtailPipelineConfig",
			"log:ListLogtailPipelineConfig",
			"cs:UpdateContactGroup",
			"cs:DescribeTemplates",
			"cs:DescribeTemplateAttribute",
			"eventbridge:PutEvents",
		},
	},
	"AliyunCSManagedNetworkRole": {
		Name: "AliyunCSManagedNetworkRolePolicy",
		Actions: []string{
			"ecs:CreateNetworkInterface",
			"ecs:DescribeNetworkInterfaces",
			"ecs:AttachNetworkInterface",
			"ecs:DetachNetworkInterface",
			"ecs:DeleteNetworkInterface",
			"ecs:DescribeInstanceAttribute",
			"ecs:AssignPrivateIpAddresses",
			"ecs:UnassignPrivateIpAddresses",
			"ecs:DescribeInstances",
			"ecs:AssignIpv6Addresses",
			"ecs:UnassignIpv6Addresses",
			"ecs:DescribeInstanceTypes",
			"ecs:ModifyNetworkInterfaceAttribute",
			"vpc:DescribeVSwitches",
		},
	},
	"AliyunCSManagedCsiRole": {
		Name: "AliyunCSManagedCsiRolePolicy",
		Actions: []string{
			"ecs:AttachDisk",
			"ecs:DetachDisk",
			"ecs:DescribeDisks",
			"ecs:CreateDisk",
			"ecs:ResizeDisk",
			"ecs:CreateSnapshot",
			"ecs:DeleteSnapshot",
			"ecs:CreateAutoSnapshotPolicy",
			"ecs:ApplyAutoSnapshotPolicy",
			"ecs:CancelAutoSnapshotPolicy",
			"ecs:DeleteAutoSnapshotPolicy",
			"ecs:DescribeAutoSnapshotPolicyEX",
			"ecs:ModifyAutoSnapshotPolicyEx",
			"ecs:AddTags",
			"ecs:RemoveTags",
			"ecs:DescribeTags",
			"ecs:DescribeSnapshots",
			"ecs:ListTagResources",
			"ecs:TagResources",
			"ecs:UntagResources",
			"ecs:ModifyDiskSpec",
			"ecs:CreateSnapshot",
			"ecs:DescribeSnapshotGroups",
			"ecs:CreateSnapshotGroup",
			"ecs:DeleteSnapshotGroup",
			"ecs:CopySnapshot",
			"ecs:DeleteDisk",
			"ecs:DescribeInstanceAttribute",
			"ecs:DescribeInstanceHistoryEvents",
			"ecs:DescribeTaskAttribute",
			"ecs:DescribeInstances",
			"nas:DescribeFileSystems",
			"nas:DescribeMountTargets",
			"nas:AddTags",
			"nas:DescribeTags",
			"nas:RemoveTags",
			"nas:CreateFileSystem",
			"nas:DeleteFileSystem",
			"nas:ModifyFileSystem",
			"nas:CreateMountTarget",
			"nas:DeleteMountTarget",
			"nas:ModifyMountTarget",
			"nas:TagResources",
			"nas:SetDirQuota",
			"nas:EnableRecycleBin",
			"nas:GetRecycleBinAttribute",
			"nas:DescribeProtocolMountTarget",
			"nas:CancelDirQuota",
			"nas:CreateDir",
			"nas:DescribeDirQuotas",
			"cs:CreateResourcesSystemTags",
			"cs:DescribeTemplateAttribute",
			"cs:DescribeTemplates",
			"oss:PutBucket",
			"oss:GetObjectTagging",
			"oss:ListBuckets",
			"oss:PutBucketTags",
			"oss:GetBucketTags",
			"oss:PutBucketEncryption",
			"oss:GetBucketStat",
			"oss:PutBucketVersioning",
			"oss:GetBucketInfo",
			"ens:DescribeInstances",
			"ens:DescribeDisks",
			"ens:ModifyDiskAttribute",
			"ens:CreateDisk",
			"ens:DetachDisk",
			"ens:AttachDisk",
			"ens:DeleteDisk",
			"kms:ListAliases",
			"hbr:CreateVault",
			"hbr:CreateBackupJob",
			"hbr:DescribeVaults",
			"hbr:DescribeBackupJobs2",
			"hbr:DescribeRestoreJobs",
			"hbr:SearchHistoricalSnapshots",
			"hbr:CreateRestoreJob",
			"hbr:AddContainerCluster",
			"hbr:DescribeContainerCluster",
			"hbr:DescribeRestoreJobs2",
			"oss:PutObject",
			"oss:IsObjectExist",
			"oss:ListObjects",
			"oss:GetObject",
			"oss:DeleteObject",
			"oss:GetBucket",
		},
	},
	"AliyunCSManagedAcrRole": {
		Name: "AliyunCSManagedAcrRolePolicy",
		Actions: []string{
			"cr:GetAuthorizationToken",
			"cr:ListInstanceEndpoint",
			"cr:PullRepository",
		},
	},
	"AliyunCSManagedCostRole": {
		Name: "AliyunCSManagedCostRolePolicy",
		Actions: []string{
			"bssapi:QueryInstanceBill",
			"bssapi:DescribeInstanceBill",
			"ecs:DescribeDisks",
			"ecs:DescribeSpotPriceHistory",
			"ecs:DescribeInstances",
			"ecs:DescribePrice",
			"eci:DescribeContainerGroups",
			"eci:DescribeContainerGroupPrice",
		},
	},
	"AliyunCSManagedMseRole": {
		Name: "AliyunCSManagedMseRolePolicy",
		Actions: []string{
			"mse:AddBlackWhiteList",
			"mse:AddGateway",
			"mse:AddServiceSource",
			"mse:CreateApplication",
			"mse:DeleteGateway",
			"mse:GetBlackWhiteList",
			"mse:GetGateway",
			"mse:GetGatewayDetail",
			"mse:GetGatewayOption",
			"mse:ListServiceSource",
			"mse:ListTagResources",
			"mse:ModifyLosslessRule",
			"mse:TagResources",
			"mse:UntagResources",
			"mse:UpdateBlackWhiteList",
			"mse:UpdateGatewayOption",
			"mse:UpdateServiceSource",
			"mse:GetLicenseKey",
			"log:CloseProductDataCollection",
			"log:OpenProductDataCollection",
			"log:GetProductDataCollection",
			"ram:CreateServiceLinkedRole",
		},
	},
	"AliyunCSManagedArmsRole": {
		Name: "AliyunCSManagedArmsRolePolicy",
		Actions: []string{
			"arms:CMonitorCloudInstances",
			"arms:CMonitorRegister",
			"arms:ConfigAgentLabel",
			"arms:CreateAlertRules",
			"arms:CreateAlertTemplate",
			"arms:CreateApp",
			"arms:CreateContact",
			"arms:CreateContactGroup",
			"arms:CreateDispatchRule",
			"arms:CreateOrUpdateIMRobot",
			"arms:CreateOrUpdateWebhookContact",
			"arms:CreateProm",
			"arms:CreatePrometheusAlertRule",
			"arms:DeleteAlert",
			"arms:DeleteAlertContact",
			"arms:DeleteAlertContactGroup",
			"arms:DeleteAlertRules",
			"arms:DeleteAlertTemplate",
			"arms:DeleteApp",
			"arms:DeleteContact",
			"arms:DeleteContactGroup",
			"arms:DeleteContactLink",
			"arms:DeleteContactMember",
			"arms:DeleteDispatchRule",
			"arms:DeleteIMRobot",
			"arms:DeletePrometheusAlertRule",
			"arms:DeleteWebhookContact",
			"arms:DescribeDispatchRule",
			"arms:DescribeIMRobots",
			"arms:DescribePrometheusAlertRule",
			"arms:DescribeWebhookContacts",
			"arms:DisableAlertTemplate",
			"arms:EnableAlertTemplate",
			"arms:GetAlarmHistories",
			"arms:GetAlert",
			"arms:GetAlertEvents",
			"arms:GetAlertRules",
			"arms:GetAlertRulesByPage",
			"arms:GetAssumeRoleCredentials",
			"arms:GetCommercialStatus",
			"arms:InstallEventer",
			"arms:InstallManagedPrometheus",
			"arms:ListActivatedAlerts",
			"arms:ListAlertTemplates",
			"arms:ListDashboards",
			"arms:ListDispatchRule",
			"arms:ListEscalationPolicies",
			"arms:ListOnCallSchedules",
			"arms:ListPrometheusAlertRules",
			"arms:ListPrometheusAlertTemplates",
			"arms:QueryAlarmHistory",
			"arms:QueryAlarmName",
			"arms:SaveAlert",
			"arms:SaveContactGroup",
			"arms:SaveContactMember",
			"arms:SaveTraceAppConfig",
			"arms:SearchAlarmHistories",
			"arms:SearchAlertRules",
			"arms:SearchContact",
			"arms:SearchContactGroup",
			"arms:SearchEvents",
			"arms:SendTTSVerifyLink",
			"arms:StartAlert",
			"arms:StartAlertRule",
			"arms:StopAlert",
			"arms:StopAlertRule",
			"arms:UninstallManagedPrometheus",
			"arms:UpdateAlertRules",
			"arms:UpdateAlertTemplate",
			"arms:UpdateContact",
			"arms:UpdateContactGroup",
			"arms:UpdateContactMember",
			"arms:UpdateDispatchRule",
			"arms:UpdatePrometheusAlertRule",
			"arms:CheckServiceStatus",
			"arms:GetClusterAllUrl",
			"arms:GetClusterInfoForArms",
			"arms:GetExploreUrl",
			"arms:GetIntegrationState",
			"arms:GetManagedPrometheusStatus",
			"arms:ListAlertEvents",
			"arms:QueryMetric",
			"arms:QueryPromInstallStatus",
			"arms:SearchAlertContactGroup",
			"arms:SearchAlertHistories",
			"arms:CreateAlertContact",
			"arms:CreateAlertContactGroup",
			"arms:ImportCustomAlertRules",
			"arms:SearchAlertContact",
			"arms:UpdateAlertContact",
			"arms:UpdateAlertContactGroup",
			"arms:UpdateAlertRule",
			"arms:UpdateWebhook",
			"arms:InnerFetchContactGroupByArmsContactGroupId",
			"xtrace:GetToken",
			"arms:ListEnvironments",
			"arms:DescribeAddonRelease",
			"arms:InstallAddon",
			"arms:DeleteAddonRelease",
			"arms:ListEnvironmentDashboards",
			"arms:ListAddonReleases",
			"arms:CreateEnvironment",
			"arms:InitEnvironment",
			"arms:DescribeEnvironment",
			"arms:InstallEnvironmentFeature",
			"arms:ListEnvironmentFeatures",
			"arms:UpdateEnvironment",
			"mse:GetLicenseKey",
		},
	},
	"AliyunCSManagedAutoScalerRole": {
		Name: "AliyunCSManagedAutoScalerRolePolicy",
		Actions: []string{
			"ess:DescribeScalingGroups",
			"ess:DescribeScalingInstances",
			"ess:DescribeScalingActivities",
			"ess:DescribeScalingConfigurations",
			"ess:DescribeScalingRules",
			"ess:DescribeScheduledTasks",
			"ess:DescribeLifecycleHooks",
			"ess:DescribeNotificationConfigurations",
			"ess:DescribeNotificationTypes",
			"ess:DescribeRegions",
			"ess:CreateScalingRule",
			"ess:ModifyScalingGroup",
			"ess:RemoveInstances",
			"ess:ExecuteScalingRule",
			"ess:ModifyScalingRule",
			"ess:DeleteScalingRule",
			"ecs:DescribeInstanceTypes",
			"ess:DetachInstances",
			"ess:CompleteLifecycleAction",
			"ess:ScaleWithAdjustment",
			"vpc:DescribeVSwitches",
			"cs:DeleteClusterNodes",
			"cs:DescribeClusterNodes",
			"cs:DescribeClusterNodePools",
			"cs:DescribeClusterNodePoolDetail",
			"cs:DescribeTaskInfo",
			"cs:ScaleClusterNodePool",
			"cs:RemoveNodePoolNodes",
			"ecs:DescribeAvailableResource",
			"ecs:DescribeInstanceTypeFamilies",
			"ecs:DescribeInstances",
			"ecs:DescribeImages",
		},
	},
	"AliyunCSManagedBackupRestoreRole": {
		Name: "AliyunCSManagedBackupRestoreRolePolicy",
		Actions: []string{
			"hbr:CreateVault",
			"hbr:CreateBackupJob",
			"hbr:DescribeVaults",
			"hbr:DescribeBackupJobs2",
			"hbr:DescribeRestoreJobs",
			"hbr:SearchHistoricalSnapshots",
			"hbr:CreateRestoreJob",
			"hbr:AddContainerCluster",
			"hbr:DescribeContainerCluster",
			"hbr:DescribeRestoreJobs2",
			"ecs:CreateSnapshot",
			"ecs:DeleteSnapshot",
			"ecs:DescribeSnapshotGroups",
			"ecs:CreateAutoSnapshotPolicy",
			"ecs:ApplyAutoSnapshotPolicy",
			"ecs:CancelAutoSnapshotPolicy",
			"ecs:DeleteAutoSnapshotPolicy",
			"ecs:DescribeAutoSnapshotPolicyEX",
			"ecs:ModifyAutoSnapshotPolicyEx",
			"ecs:DescribeSnapshots",
			"ecs:DescribeInstances",
			"ecs:CopySnapshot",
			"ecs:CreateSnapshotGroup",
			"ecs:DeleteSnapshotGroup",
			"oss:PutObject",
			"oss:GetObject",
			"oss:DeleteObject",
			"oss:GetBucket",
			"oss:ListObjects",
			"oss:ListBuckets",
			"oss:GetBucketStat",
		},
	},
	"AliyunCSManagedWebhookInjectorRole": {
		Name: "AliyunCSManagedWebhookInjectorRolePolicy",
		Actions: []string{
			"ecs:DescribeSecurityGroupAttribute",
			"ecs:AutherizeSecurityGroup",
			"ecs:RevokeSecurityGroup",
			"rds:ModifySecurityIps",
			"rds:DescribeDBInstanceIPArrayList",
		},
	},
	"kube-ai": {
		Name: "kube-ai-policy",
		Actions: []string{
			"log:GetProject",
			"log:GetLogStore",
			"log:GetConfig",
			"log:GetMachineGroup",
			"log:GetAppliedMachineGroups",
			"log:GetAppliedConfigs",
			"log:GetIndex",
			"log:GetSavedSearch",
			"log:GetDashboard",
			"log:GetJob",
			"ecs:DescribeInstances",
			"ecs:DescribeSpotPriceHistory",
			"ecs:DescribePrice",
			"eci:DescribeContainerGroups",
			"eci:DescribeContainerGroupPrice",
			"log:GetLogStoreLogs",
			"ims:CreateApplication",
			"ims:UpdateApplication",
			"ims:GetApplication",
			"ims:ListApplications",
			"ims:DeleteApplication",
			"ims:CreateAppSecret",
			"ims:GetAppSecret",
			"ims:ListAppSecretIds",
			"ims:ListUsers",
		},
	},
}

func (a NeedUpdateAddon) Policies() []RolePolicy {
	var polices []RolePolicy
	for _, role := range a.RoleNames {
		policy := defaultPoliciesByRoleMap[role]
		polices = append(polices, policy)
	}

	return polices
}

func (a NeedUpdateAddon) PolicyNames() []string {
	var names []string
	polices := a.Policies()
	for _, p := range polices {
		if utils.StringSliceInclude(names, p.Name) {
			continue
		}
		names = append(names, p.Name)
	}
	return names
}

func (a NeedUpdateAddon) PolicyActions() []string {
	var actions []string
	polices := a.Policies()
	for _, p := range polices {
		for _, as := range p.Actions {
			if utils.StringSliceInclude(actions, as) {
				continue
			}
			actions = append(actions, as)
		}
	}
	return actions
}

func (w Workload) String() string {
	return fmt.Sprintf("%s.%s.%s", w.Name, w.Type, w.Namespace)
}
