"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[9014],{8453:(e,s,t)=>{t.d(s,{R:()=>o,x:()=>i});var n=t(6540);const r={},c=n.createContext(r);function o(e){const s=n.useContext(c);return n.useMemo((function(){return"function"==typeof e?e(s):{...s,...e}}),[s,e])}function i(e){let s;return s=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:o(e.components),n.createElement(c.Provider,{value:s},e.children)}},8900:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>a,contentTitle:()=>o,default:()=>h,frontMatter:()=>c,metadata:()=>i,toc:()=>l});var n=t(4848),r=t(8453);const c={slug:"/zh-CN/rrsa/associate-role",title:"associate-role\uff08\u4e2d\u6587\uff09",sidebar_position:2},o="associate-role",i={id:"rrsa/associate-role.zh-CN",title:"associate-role\uff08\u4e2d\u6587\uff09",description:"\u914d\u7f6e RAM \u89d2\u8272\uff0c\u5141\u8bb8\u4f7f\u7528\u8868\u793a\u7279\u5b9a service account \u8eab\u4efd\u7684 oidc token \u626e\u6f14\u8be5 RAM \u89d2\u8272\u3002",source:"@site/versioned_docs/version-v0.15.0/rrsa/associate-role.zh-CN.md",sourceDirName:"rrsa",slug:"/zh-CN/rrsa/associate-role",permalink:"/ack-ram-tool/v0.15.0/zh-CN/rrsa/associate-role",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.15.0/rrsa/associate-role.zh-CN.md",tags:[],version:"v0.15.0",sidebarPosition:2,frontMatter:{slug:"/zh-CN/rrsa/associate-role",title:"associate-role\uff08\u4e2d\u6587\uff09",sidebar_position:2},sidebar:"tutorialSidebar",previous:{title:"associate-role",permalink:"/ack-ram-tool/v0.15.0/rrsa/associate-role"},next:{title:"install-helper-addon",permalink:"/ack-ram-tool/v0.15.0/rrsa/install-helper-addon"}},a={},l=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}];function d(e){const s={code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,r.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(s.h1,{id:"associate-role",children:"associate-role"}),"\n",(0,n.jsx)(s.p,{children:"\u914d\u7f6e RAM \u89d2\u8272\uff0c\u5141\u8bb8\u4f7f\u7528\u8868\u793a\u7279\u5b9a service account \u8eab\u4efd\u7684 oidc token \u626e\u6f14\u8be5 RAM \u89d2\u8272\u3002"}),"\n",(0,n.jsx)(s.h2,{id:"\u4f7f\u7528\u793a\u4f8b",children:"\u4f7f\u7528\u793a\u4f8b"}),"\n",(0,n.jsx)(s.pre,{children:(0,n.jsx)(s.code,{className:"language-shell",children:'$ ack-ram-tool rrsa associate-role --cluster-id <clusterId> \\\n  --namespace <namespce> --service-account <serviceAccountName> \\\n  --role-name <roleName>\n\n? Are you sure you want to associate RAM Role "<roleName>" to service account "<serviceAccountName>" (namespace: "<namespce>")? Yes\n2023-04-20T14:30:02+08:00 INFO will change the AssumeRole Policy of RAM Role "<roleName>" with blow content:\n{\n  "Statement": [\n   {\n    "Action": "sts:AssumeRole",\n    "Condition": {\n     "StringEquals": {\n      "oidc:aud": "sts.aliyuncs.com",\n      "oidc:iss": "https://oidc-ack-***.aliyuncs.com/c132c***",\n      "oidc:sub": "system:serviceaccount:<namespce>:<serviceAccountName>"\n     }\n    },\n    "Effect": "Allow",\n    "Principal": {\n     "Federated": [\n      "acs:ram::113***:oidc-provider/ack-rrsa-c132c***"\n     ]\n    }\n   }\n  ],\n  "Version": "1"\n }\n\n? Are you sure you want to associate RAM Role "test" to service account "sa" (namespace: "test")? Yes\n2023-04-20T14:30:04+08:00 INFO Associate RAM Role "test" to service account "sa" (namespace: "test") successfully\n'})}),"\n",(0,n.jsx)(s.h2,{id:"\u547d\u4ee4\u884c\u53c2\u6570",children:"\u547d\u4ee4\u884c\u53c2\u6570"}),"\n",(0,n.jsx)(s.pre,{children:(0,n.jsx)(s.code,{children:'Usage:\n  ack-ram-tool rrsa associate-role [flags]\n\nFlags:\n      --attach-custom-policy string   Attach this custom policy to the RAM Role\n      --attach-system-policy string   Attach this system policy to the RAM Role\n  -c, --cluster-id string             The cluster id to use\n      --create-role-if-not-exist      Create the RAM Role if it does not exist\n  -h, --help                          help for associate-role\n  -n, --namespace string              The Kubernetes namespace to use\n  -r, --role-name string              The RAM Role name to use\n  -s, --service-account string        The Kubernetes service account to use\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,n.jsx)(s.p,{children:"\u53c2\u6570\u8bf4\u660e\uff1a"}),"\n",(0,n.jsxs)(s.table,{children:[(0,n.jsx)(s.thead,{children:(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.th,{children:"\u53c2\u6570\u540d\u79f0"}),(0,n.jsx)(s.th,{children:"\u9ed8\u8ba4\u503c"}),(0,n.jsx)(s.th,{children:"\u5fc5\u9700\u53c2\u6570"}),(0,n.jsx)(s.th,{children:"\u8bf4\u660e"})]})}),(0,n.jsxs)(s.tbody,{children:[(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"-c, --cluster-id"}),(0,n.jsx)(s.td,{children:"\u65e0"}),(0,n.jsx)(s.td,{children:"\u662f"}),(0,n.jsx)(s.td,{children:"\u96c6\u7fa4 ID"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"-n, --namespace"}),(0,n.jsx)(s.td,{children:"\u65e0"}),(0,n.jsx)(s.td,{children:"\u662f"}),(0,n.jsxs)(s.td,{children:["\u547d\u540d\u7a7a\u95f4\uff0c\u53ef\u4ee5\u4f7f\u7528 ",(0,n.jsx)(s.code,{children:"*"})," \u8868\u793a\u6240\u6709\u547d\u540d\u7a7a\u95f4\uff1a",(0,n.jsx)(s.code,{children:"--namespace '*'"})]})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"-s, --service-account"}),(0,n.jsx)(s.td,{children:"\u65e0"}),(0,n.jsx)(s.td,{children:"\u662f"}),(0,n.jsx)(s.td,{children:"service account"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"-r, --role-name"}),(0,n.jsx)(s.td,{children:"\u65e0"}),(0,n.jsx)(s.td,{children:"\u662f"}),(0,n.jsx)(s.td,{children:"RAM \u89d2\u8272"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"--create-role-if-not-exist"}),(0,n.jsx)(s.td,{children:"\u65e0"}),(0,n.jsx)(s.td,{children:"\u5426"}),(0,n.jsx)(s.td,{children:"\u5982\u679c\u8be5 RAM \u89d2\u8272\u4e0d\u5b58\u5728\uff0c\u90a3\u4e48\u81ea\u52a8\u521b\u5efa\u4e00\u4e2a\u540c\u540d\u7684 RAM \u89d2\u8272"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"--attach-system-policy"}),(0,n.jsx)(s.td,{children:"\u65e0"}),(0,n.jsx)(s.td,{children:"\u5426"}),(0,n.jsx)(s.td,{children:"\u4e3a\u8be5\u89d2\u8272\u6388\u4e88\u6307\u5b9a\u7684\u7cfb\u7edf\u6743\u9650\u7b56\u7565"})]}),(0,n.jsxs)(s.tr,{children:[(0,n.jsx)(s.td,{children:"--attach-custom-policy"}),(0,n.jsx)(s.td,{children:"\u65e0"}),(0,n.jsx)(s.td,{children:"\u5426"}),(0,n.jsx)(s.td,{children:"\u4e3a\u8be5\u89d2\u8272\u6388\u4e88\u6307\u5b9a\u7684\u81ea\u5b9a\u4e49\u6743\u9650\u7b56\u7565"})]})]})]})]})}function h(e={}){const{wrapper:s}={...(0,r.R)(),...e.components};return s?(0,n.jsx)(s,{...e,children:(0,n.jsx)(d,{...e})}):d(e)}}}]);