"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[7871],{6972:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>a,contentTitle:()=>o,default:()=>h,frontMatter:()=>s,metadata:()=>l,toc:()=>c});var i=n(4848),r=n(8453);const s={slug:"/zh-CN/credential-plugin/get-token",title:"get-token\uff08\u4e2d\u6587\uff09",sidebar_position:3},o="get-token",l={id:"credential-plugin/get-token.zh-CN",title:"get-token\uff08\u4e2d\u6587\uff09",description:"\u96c6\u6210 ack-ram-authenticator\uff0c\u83b7\u53d6\u7528\u4e8e\u8bbf\u95ee api server \u7684 ExecCredential token\u3002",source:"@site/versioned_docs/version-v0.13.2/credential-plugin/get-token.zh-CN.md",sourceDirName:"credential-plugin",slug:"/zh-CN/credential-plugin/get-token",permalink:"/ack-ram-tool/v0.13.2/zh-CN/credential-plugin/get-token",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.13.2/credential-plugin/get-token.zh-CN.md",tags:[],version:"v0.13.2",sidebarPosition:3,frontMatter:{slug:"/zh-CN/credential-plugin/get-token",title:"get-token\uff08\u4e2d\u6587\uff09",sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"get-token",permalink:"/ack-ram-tool/v0.13.2/credential-plugin/get-token"},next:{title:"export-credentials",permalink:"/ack-ram-tool/v0.13.2/category/export-credentials"}},a={},c=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}];function d(e){const t={a:"a",code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(t.h1,{id:"get-token",children:"get-token"}),"\n",(0,i.jsxs)(t.p,{children:["\u96c6\u6210 ack-ram-authenticator\uff0c\u83b7\u53d6\u7528\u4e8e\u8bbf\u95ee api server \u7684 ",(0,i.jsx)(t.a,{href:"https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins",children:"ExecCredential"})," token\u3002"]}),"\n",(0,i.jsx)(t.h2,{id:"\u4f7f\u7528\u793a\u4f8b",children:"\u4f7f\u7528\u793a\u4f8b"}),"\n",(0,i.jsx)(t.pre,{children:(0,i.jsx)(t.code,{className:"language-shell",children:'$ ack-ram-tool credential-plugin get-token --cluster-id <clusterId>\n\n{\n "kind": "ExecCredential",\n "apiVersion": "client.authentication.k8s.io/v1beta1",\n "spec": {\n  "interactive": false\n },\n "status": {\n  "token": "k8s-ack-v1.aHR0cHM6Ly9zd***"\n }\n}\n'})}),"\n",(0,i.jsx)(t.h2,{id:"\u547d\u4ee4\u884c\u53c2\u6570",children:"\u547d\u4ee4\u884c\u53c2\u6570"}),"\n",(0,i.jsx)(t.pre,{children:(0,i.jsx)(t.code,{children:'Usage:\n  ack-ram-tool credential-plugin get-token [flags]\n\nFlags:\n      --api-version string   v1 or v1beta1 (default "v1beta1")\n  -c, --cluster-id string    The cluster id to use\n  -h, --help                 help for get-token\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,i.jsx)(t.p,{children:"\u53c2\u6570\u8bf4\u660e\uff1a"}),"\n",(0,i.jsxs)(t.table,{children:[(0,i.jsx)(t.thead,{children:(0,i.jsxs)(t.tr,{children:[(0,i.jsx)(t.th,{children:"\u53c2\u6570\u540d\u79f0"}),(0,i.jsx)(t.th,{children:"\u9ed8\u8ba4\u503c"}),(0,i.jsx)(t.th,{children:"\u5fc5\u9700\u53c2\u6570"}),(0,i.jsx)(t.th,{children:"\u8bf4\u660e"})]})}),(0,i.jsxs)(t.tbody,{children:[(0,i.jsxs)(t.tr,{children:[(0,i.jsx)(t.td,{children:"-c, --cluster-id"}),(0,i.jsx)(t.td,{children:"\u65e0"}),(0,i.jsx)(t.td,{children:"\u662f"}),(0,i.jsx)(t.td,{children:"\u96c6\u7fa4 ID"})]}),(0,i.jsxs)(t.tr,{children:[(0,i.jsx)(t.td,{children:"--api-version"}),(0,i.jsx)(t.td,{children:"v1beta1"}),(0,i.jsx)(t.td,{children:"\u5426"}),(0,i.jsxs)(t.td,{children:["\u6307\u5b9a\u8fd4\u56de\u7684\u6570\u636e\u4e2d\u4f7f\u7528\u54ea\u4e2a\u7248\u672c\u7684 apiVersion\u3002v1beta1 \u8868\u793a ",(0,i.jsx)(t.code,{children:"client.authentication.k8s.io/v1beta1"}),"\uff0cv1 \u8868\u793a ",(0,i.jsx)(t.code,{children:"client.authentication.k8s.io/v1beta1"})]})]})]})]})]})}function h(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},8453:(e,t,n)=>{n.d(t,{R:()=>o,x:()=>l});var i=n(6540);const r={},s=i.createContext(r);function o(e){const t=i.useContext(s);return i.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function l(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:o(e.components),i.createElement(s.Provider,{value:t},e.children)}}}]);