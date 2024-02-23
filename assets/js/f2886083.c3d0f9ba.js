"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[6494],{1035:(e,r,n)=>{n.r(r),n.d(r,{assets:()=>l,contentTitle:()=>i,default:()=>u,frontMatter:()=>o,metadata:()=>a,toc:()=>d});var s=n(4848),t=n(8453);const o={slug:"assume-role",sidebar_position:5},i="assume-role",a={id:"rrsa/assume-role",title:"assume-role",description:"Test using an OIDC token to assume a specific RAM role.",source:"@site/versioned_docs/version-v0.14.0/rrsa/assume-role.md",sourceDirName:"rrsa",slug:"/rrsa/assume-role",permalink:"/ack-ram-tool/v0.14.0/rrsa/assume-role",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.14.0/rrsa/assume-role.md",tags:[],version:"v0.14.0",sidebarPosition:5,frontMatter:{slug:"assume-role",sidebar_position:5},sidebar:"tutorialSidebar",previous:{title:"status\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.14.0/zh-CN/rrsa/status"},next:{title:"assume-role\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.14.0/zh-CN/rrsa/assume-role"}},l={},d=[{value:"Usage",id:"usage",level:2},{value:"Flags",id:"flags",level:2}];function c(e){const r={code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,t.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(r.h1,{id:"assume-role",children:"assume-role"}),"\n",(0,s.jsx)(r.p,{children:"Test using an OIDC token to assume a specific RAM role."}),"\n",(0,s.jsx)(r.h2,{id:"usage",children:"Usage"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-shell",children:"$ ack-ram-tool rrsa assume-role --oidc-provider-arn <oidcProviderArn> \\\n  --role-arn <roleArn> --oidc-token-file <pathToTokenFile>\n\n    Retrieved a STS token:\n    AccessKeyId:       STS.***\n    AccessKeySecret:   7UVy***\n    SecurityToken:     CAIS***\n    Expiration:        2021-12-03T05:51:37Z\n\n"})}),"\n",(0,s.jsx)(r.h2,{id:"flags",children:"Flags"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{children:'Usage:\n  ack-ram-tool rrsa assume-role [flags]\n\nFlags:\n  -h, --help                       help for assume-role\n  -p, --oidc-provider-arn string   The arn of OIDC provider\n  -t, --oidc-token-file string     Path to OIDC token file. If value is \'-\', will read token from stdin\n  -r, --role-arn string            The arn of RAM role\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,s.jsx)(r.p,{children:"Descriptions\uff1a"}),"\n",(0,s.jsxs)(r.table,{children:[(0,s.jsx)(r.thead,{children:(0,s.jsxs)(r.tr,{children:[(0,s.jsx)(r.th,{children:"Flag"}),(0,s.jsx)(r.th,{children:"Default"}),(0,s.jsx)(r.th,{children:"Required"}),(0,s.jsx)(r.th,{children:"Description"})]})}),(0,s.jsxs)(r.tbody,{children:[(0,s.jsxs)(r.tr,{children:[(0,s.jsx)(r.td,{children:"-p, --oidc-provider-arn"}),(0,s.jsx)(r.td,{}),(0,s.jsx)(r.td,{children:"Yes"}),(0,s.jsx)(r.td,{children:"OIDC Provider ARN"})]}),(0,s.jsxs)(r.tr,{children:[(0,s.jsx)(r.td,{children:"-r, --role-arn"}),(0,s.jsx)(r.td,{}),(0,s.jsx)(r.td,{children:"Yes"}),(0,s.jsx)(r.td,{children:"Role ARN"})]}),(0,s.jsxs)(r.tr,{children:[(0,s.jsx)(r.td,{children:"-t, --oidc-token-file"}),(0,s.jsx)(r.td,{}),(0,s.jsx)(r.td,{children:"Yes"}),(0,s.jsx)(r.td,{children:'The path to the OIDC token file. If the value is "-", it mean that the token can be read from standard input(for example, by passing the token through a pipeline)'})]})]})]})]})}function u(e={}){const{wrapper:r}={...(0,t.R)(),...e.components};return r?(0,s.jsx)(r,{...e,children:(0,s.jsx)(c,{...e})}):c(e)}},8453:(e,r,n)=>{n.d(r,{R:()=>i,x:()=>a});var s=n(6540);const t={},o=s.createContext(t);function i(e){const r=s.useContext(o);return s.useMemo((function(){return"function"==typeof e?e(r):{...r,...e}}),[r,e])}function a(e){let r;return r=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:i(e.components),s.createElement(o.Provider,{value:r},e.children)}}}]);