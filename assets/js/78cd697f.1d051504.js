"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[3507],{7472:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>a,contentTitle:()=>c,default:()=>h,frontMatter:()=>o,metadata:()=>i,toc:()=>l});var n=t(4848),r=t(8453);const o={slug:"associate-role",sidebar_position:2},c="associate-role",i={id:"rrsa/associate-role",title:"associate-role",description:"Configure RAM roles to allow the use of OIDC tokens representing specific service account identities to assume the RAM roles.",source:"@site/versioned_docs/version-v0.13.0/rrsa/associate-role.md",sourceDirName:"rrsa",slug:"/rrsa/associate-role",permalink:"/ack-ram-tool/v0.13.0/rrsa/associate-role",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.13.0/rrsa/associate-role.md",tags:[],version:"v0.13.0",sidebarPosition:2,frontMatter:{slug:"associate-role",sidebar_position:2},sidebar:"tutorialSidebar",previous:{title:"enable\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.13.0/zh-CN/rrsa/enable"},next:{title:"associate-role\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.13.0/zh-CN/rrsa/associate-role"}},a={},l=[{value:"Usage",id:"usage",level:2},{value:"Flags",id:"flags",level:2}];function d(e){const s={code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,r.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(s.h1,{id:"associate-role",children:"associate-role"}),"\n",(0,n.jsx)(s.p,{children:"Configure RAM roles to allow the use of OIDC tokens representing specific service account identities to assume the RAM roles."}),"\n",(0,n.jsx)(s.h2,{id:"usage",children:"Usage"}),"\n",(0,n.jsx)(s.pre,{children:(0,n.jsx)(s.code,{className:"language-shell",children:'$ ack-ram-tool rrsa associate-role --cluster-id <clusterId> \\\n  --namespace <namespce> --service-account <serviceAccountName> \\\n  --role-name <roleName>\n\n? Are you sure you want to associate RAM Role "<roleName>" to service account "<serviceAccountName>" (namespace: "<namespce>")? Yes\n2023-04-20T14:30:02+08:00 INFO will change the AssumeRole Policy of RAM Role "<roleName>" with blow content:\n{\n  "Statement": [\n   {\n    "Action": "sts:AssumeRole",\n    "Condition": {\n     "StringEquals": {\n      "oidc:aud": "sts.aliyuncs.com",\n      "oidc:iss": "https://oidc-ack-***.aliyuncs.com/c132c***",\n      "oidc:sub": "system:serviceaccount:<namespce>:<serviceAccountName>"\n     }\n    },\n    "Effect": "Allow",\n    "Principal": {\n     "Federated": [\n      "acs:ram::113***:oidc-provider/ack-rrsa-c132c***"\n     ]\n    }\n   }\n  ],\n  "Version": "1"\n }\n\n? Are you sure you want to associate RAM Role "test" to service account "sa" (namespace: "test")? Yes\n2023-04-20T14:30:04+08:00 INFO Associate RAM Role "test" to service account "sa" (namespace: "test") successfully\n'})}),"\n",(0,n.jsx)(s.h2,{id:"flags",children:"Flags"}),"\n",(0,n.jsx)(s.pre,{children:(0,n.jsx)(s.code,{children:'Usage:\n  ack-ram-tool rrsa associate-role [flags]\n\nFlags:\n      --attach-custom-policy string   Attach this custom policy to the RAM Role\n      --attach-system-policy string   Attach this system policy to the RAM Role\n  -c, --cluster-id string             The cluster id to use\n      --create-role-if-not-exist      Create the RAM Role if it does not exist\n  -h, --help                          help for associate-role\n  -n, --namespace string              The Kubernetes namespace to use\n  -r, --role-name string              The RAM Role name to use\n  -s, --service-account string        The Kubernetes service account to use\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,n.jsx)(s.p,{children:"Descriptions\uff1a"}),"\n",(0,n.jsxs)(s.table,{children:[(0,n.jsx)(s.thead,{children:(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.th,{children:"Flag"}),(0,n.jsx)(s.th,{children:"Default"}),(0,n.jsx)(s.th,{children:"Required"}),(0,n.jsx)(s.th,{children:"Description"})]})}),(0,n.jsxs)(s.tbody,{children:[(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"-c, --cluster-id"}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{children:"Yes"}),(0,n.jsx)(s.td,{children:"Cluster ID"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"-n, --namespace"}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{children:"Yes"}),(0,n.jsxs)(s.td,{children:["namespace\uff0ccan use ",(0,n.jsx)(s.code,{children:"*"})," to represent all namespaces\uff1a",(0,n.jsx)(s.code,{children:"--namespace '*'"})]})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"-s, --service-account"}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{children:"Yes"}),(0,n.jsx)(s.td,{children:"service account"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"-r, --role-name"}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{children:"Yes"}),(0,n.jsx)(s.td,{children:"RAM Role"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"--create-role-if-not-exist"}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{children:"auto create an RAM Role if it does not exists"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"--attach-system-policy"}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{children:"attach a system policy to the role"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"--attach-custom-policy"}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{}),(0,n.jsx)(s.td,{children:"attach a custom policy to the role"})]})]})]})]})}function h(e={}){const{wrapper:s}={...(0,r.R)(),...e.components};return s?(0,n.jsx)(s,{...e,children:(0,n.jsx)(d,{...e})}):d(e)}},8453:(e,s,t)=>{t.d(s,{R:()=>c,x:()=>i});var n=t(6540);const r={},o=n.createContext(r);function c(e){const s=n.useContext(o);return n.useMemo((function(){return"function"==typeof e?e(s):{...s,...e}}),[s,e])}function i(e){let s;return s=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:c(e.components),n.createElement(o.Provider,{value:s},e.children)}}}]);