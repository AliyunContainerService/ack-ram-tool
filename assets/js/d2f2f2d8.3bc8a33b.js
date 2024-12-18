"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[9269],{8734:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>u,frontMatter:()=>s,metadata:()=>o,toc:()=>c});var i=n(4848),r=n(8453);const s={slug:"get-token",sidebar_position:3},a="get-token",o={id:"credential-plugin/get-token",title:"get-token",description:"Integrate ack-ram-authenticator to obtain the ExecCredential token used to access the API server.",source:"@site/versioned_docs/version-v0.13.2/credential-plugin/get-token.md",sourceDirName:"credential-plugin",slug:"/credential-plugin/get-token",permalink:"/ack-ram-tool/v0.13.2/credential-plugin/get-token",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.13.2/credential-plugin/get-token.md",tags:[],version:"v0.13.2",sidebarPosition:3,frontMatter:{slug:"get-token",sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"get-credential\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.13.2/zh-CN/credential-plugin/get-credential"},next:{title:"get-token\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.13.2/zh-CN/credential-plugin/get-token"}},l={},c=[{value:"Usage",id:"usage",level:2},{value:"Flags",id:"flags",level:2}];function d(e){const t={a:"a",code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(t.h1,{id:"get-token",children:"get-token"}),"\n",(0,i.jsxs)(t.p,{children:["Integrate ack-ram-authenticator to obtain the ",(0,i.jsx)(t.a,{href:"https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins",children:"ExecCredential"})," token used to access the API server."]}),"\n",(0,i.jsx)(t.h2,{id:"usage",children:"Usage"}),"\n",(0,i.jsx)(t.pre,{children:(0,i.jsx)(t.code,{className:"language-shell",children:'$ ack-ram-tool credential-plugin get-token --cluster-id <clusterId>\n\n{\n "kind": "ExecCredential",\n "apiVersion": "client.authentication.k8s.io/v1beta1",\n "spec": {\n  "interactive": false\n },\n "status": {\n  "token": "k8s-ack-v1.aHR0cHM6Ly9zd***"\n }\n}\n'})}),"\n",(0,i.jsx)(t.h2,{id:"flags",children:"Flags"}),"\n",(0,i.jsx)(t.pre,{children:(0,i.jsx)(t.code,{children:'Usage:\n  ack-ram-tool credential-plugin get-token [flags]\n\nFlags:\n      --api-version string   v1 or v1beta1 (default "v1beta1")\n  -c, --cluster-id string    The cluster id to use\n  -h, --help                 help for get-token\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,i.jsx)(t.p,{children:"Descriptions\uff1a"}),"\n",(0,i.jsxs)(t.table,{children:[(0,i.jsx)(t.thead,{children:(0,i.jsxs)(t.tr,{children:[(0,i.jsx)(t.th,{children:"Flag"}),(0,i.jsx)(t.th,{children:"Default"}),(0,i.jsx)(t.th,{children:"Required"}),(0,i.jsx)(t.th,{children:"Description"})]})}),(0,i.jsxs)(t.tbody,{children:[(0,i.jsxs)(t.tr,{children:[(0,i.jsx)(t.td,{children:"-c, --cluster-id"}),(0,i.jsx)(t.td,{}),(0,i.jsx)(t.td,{children:"Yes"}),(0,i.jsx)(t.td,{children:"Cluster ID"})]}),(0,i.jsxs)(t.tr,{children:[(0,i.jsx)(t.td,{children:"--api-version"}),(0,i.jsx)(t.td,{children:"v1beta1"}),(0,i.jsx)(t.td,{}),(0,i.jsxs)(t.td,{children:["Specify which version of apiVersion to use in the returned data. ",(0,i.jsx)(t.code,{children:"v1beta1"})," represents ",(0,i.jsx)(t.code,{children:"client.authentication.k8s.io/v1beta1"}),", and ",(0,i.jsx)(t.code,{children:"v1"})," represents ",(0,i.jsx)(t.code,{children:"client.authentication.k8s.io/v1"}),"."]})]})]})]})]})}function u(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},8453:(e,t,n)=>{n.d(t,{R:()=>a,x:()=>o});var i=n(6540);const r={},s=i.createContext(r);function a(e){const t=i.useContext(s);return i.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function o(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:a(e.components),i.createElement(s.Provider,{value:t},e.children)}}}]);