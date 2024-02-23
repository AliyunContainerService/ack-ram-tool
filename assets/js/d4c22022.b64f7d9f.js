"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[8530],{776:(e,r,n)=>{n.r(r),n.d(r,{assets:()=>a,contentTitle:()=>i,default:()=>h,frontMatter:()=>o,metadata:()=>l,toc:()=>d});var s=n(4848),t=n(8453);const o={slug:"/zh-CN/rrsa/assume-role",title:"assume-role\uff08\u4e2d\u6587\uff09",sidebar_position:5},i="assume-role",l={id:"rrsa/assume-role.zh-CN",title:"assume-role\uff08\u4e2d\u6587\uff09",description:"\u6d4b\u8bd5\u4f7f\u7528 oidc token \u626e\u6f14\u7279\u5b9a RAM \u89d2\u8272\u3002",source:"@site/versioned_docs/version-v0.15.0/rrsa/assume-role.zh-CN.md",sourceDirName:"rrsa",slug:"/zh-CN/rrsa/assume-role",permalink:"/ack-ram-tool/v0.15.0/zh-CN/rrsa/assume-role",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.15.0/rrsa/assume-role.zh-CN.md",tags:[],version:"v0.15.0",sidebarPosition:5,frontMatter:{slug:"/zh-CN/rrsa/assume-role",title:"assume-role\uff08\u4e2d\u6587\uff09",sidebar_position:5},sidebar:"tutorialSidebar",previous:{title:"assume-role",permalink:"/ack-ram-tool/v0.15.0/rrsa/assume-role"},next:{title:"examples",permalink:"/ack-ram-tool/v0.15.0/rrsa/examples"}},a={},d=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}];function c(e){const r={code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,t.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(r.h1,{id:"assume-role",children:"assume-role"}),"\n",(0,s.jsx)(r.p,{children:"\u6d4b\u8bd5\u4f7f\u7528 oidc token \u626e\u6f14\u7279\u5b9a RAM \u89d2\u8272\u3002"}),"\n",(0,s.jsx)(r.h2,{id:"\u4f7f\u7528\u793a\u4f8b",children:"\u4f7f\u7528\u793a\u4f8b"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-shell",children:"$ ack-ram-tool rrsa assume-role --oidc-provider-arn <oidcProviderArn> \\\n  --role-arn <roleArn> --oidc-token-file <pathToTokenFile>\n\n    Retrieved a STS token:\n    AccessKeyId:       STS.***\n    AccessKeySecret:   7UVy***\n    SecurityToken:     CAIS***\n    Expiration:        2021-12-03T05:51:37Z\n\n"})}),"\n",(0,s.jsx)(r.h2,{id:"\u547d\u4ee4\u884c\u53c2\u6570",children:"\u547d\u4ee4\u884c\u53c2\u6570"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{children:'Usage:\n  ack-ram-tool rrsa assume-role [flags]\n\nFlags:\n  -h, --help                       help for assume-role\n  -p, --oidc-provider-arn string   The arn of OIDC provider\n  -t, --oidc-token-file string     Path to OIDC token file. If value is \'-\', will read token from stdin\n  -r, --role-arn string            The arn of RAM role\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,s.jsx)(r.p,{children:"\u53c2\u6570\u8bf4\u660e\uff1a"}),"\n",(0,s.jsxs)(r.table,{children:[(0,s.jsx)(r.thead,{children:(0,s.jsxs)(r.tr,{children:[(0,s.jsx)(r.th,{children:"\u53c2\u6570\u540d\u79f0"}),(0,s.jsx)(r.th,{children:"\u9ed8\u8ba4\u503c"}),(0,s.jsx)(r.th,{children:"\u5fc5\u9700\u53c2\u6570"}),(0,s.jsx)(r.th,{children:"\u8bf4\u660e"})]})}),(0,s.jsxs)(r.tbody,{children:[(0,s.jsxs)(r.tr,{children:[(0,s.jsx)(r.td,{children:"-p, --oidc-provider-arn"}),(0,s.jsx)(r.td,{children:"\u65e0"}),(0,s.jsx)(r.td,{children:"\u662f"}),(0,s.jsx)(r.td,{children:"\u4e3a\u96c6\u7fa4\u6ce8\u518c\u7684 RAM \u89d2\u8272 SSO \u4f9b\u5e94\u5546 ARN"})]}),(0,s.jsxs)(r.tr,{children:[(0,s.jsx)(r.td,{children:"-r, --role-arn"}),(0,s.jsx)(r.td,{children:"\u65e0"}),(0,s.jsx)(r.td,{children:"\u662f"}),(0,s.jsx)(r.td,{children:"\u88ab\u626e\u6f14\u7684 RAM \u89d2\u8272\u7684 ARN"})]}),(0,s.jsxs)(r.tr,{children:[(0,s.jsx)(r.td,{children:"-t, --oidc-token-file"}),(0,s.jsx)(r.td,{children:"\u65e0"}),(0,s.jsx)(r.td,{children:"\u662f"}),(0,s.jsxs)(r.td,{children:["oidc token \u6587\u4ef6\u7684\u8def\u5f84\u3002\u5f53\u503c\u4e3a ",(0,s.jsx)(r.code,{children:"-"})," \u65f6\u652f\u6301\u4ece\u6807\u51c6\u8f93\u5165\u4ece\u8bfb\u53d6 token\uff08\u6bd4\u5982\u901a\u8fc7\u7ba1\u9053\u4f20\u9012 token\uff09"]})]})]})]})]})}function h(e={}){const{wrapper:r}={...(0,t.R)(),...e.components};return r?(0,s.jsx)(r,{...e,children:(0,s.jsx)(c,{...e})}):c(e)}},8453:(e,r,n)=>{n.d(r,{R:()=>i,x:()=>l});var s=n(6540);const t={},o=s.createContext(t);function i(e){const r=s.useContext(o);return s.useMemo((function(){return"function"==typeof e?e(r):{...r,...e}}),[r,e])}function l(e){let r;return r=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:i(e.components),s.createElement(o.Provider,{value:r},e.children)}}}]);