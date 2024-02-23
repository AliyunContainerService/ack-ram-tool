"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[6193],{4088:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>a,contentTitle:()=>o,default:()=>h,frontMatter:()=>s,metadata:()=>l,toc:()=>c});var i=t(4848),r=t(8453);const s={slug:"/zh-CN/credential-plugin/get-token",title:"get-token\uff08\u4e2d\u6587\uff09",sidebar_position:3},o="get-token",l={id:"credential-plugin/get-token.zh-CN",title:"get-token\uff08\u4e2d\u6587\uff09",description:"\u96c6\u6210 ack-ram-authenticator\uff0c\u83b7\u53d6\u7528\u4e8e\u8bbf\u95ee api server \u7684 ExecCredential token\u3002",source:"@site/versioned_docs/version-v0.15.0/credential-plugin/get-token.zh-CN.md",sourceDirName:"credential-plugin",slug:"/zh-CN/credential-plugin/get-token",permalink:"/ack-ram-tool/v0.15.0/zh-CN/credential-plugin/get-token",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.15.0/credential-plugin/get-token.zh-CN.md",tags:[],version:"v0.15.0",sidebarPosition:3,frontMatter:{slug:"/zh-CN/credential-plugin/get-token",title:"get-token\uff08\u4e2d\u6587\uff09",sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"get-token",permalink:"/ack-ram-tool/v0.15.0/credential-plugin/get-token"},next:{title:"export-credentials",permalink:"/ack-ram-tool/v0.15.0/category/export-credentials"}},a={},c=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}];function d(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.h1,{id:"get-token",children:"get-token"}),"\n",(0,i.jsxs)(n.p,{children:["\u96c6\u6210 ack-ram-authenticator\uff0c\u83b7\u53d6\u7528\u4e8e\u8bbf\u95ee api server \u7684 ",(0,i.jsx)(n.a,{href:"https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins",children:"ExecCredential"})," token\u3002"]}),"\n",(0,i.jsx)(n.h2,{id:"\u4f7f\u7528\u793a\u4f8b",children:"\u4f7f\u7528\u793a\u4f8b"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool credential-plugin get-token --cluster-id <clusterId>\n\n{\n "kind": "ExecCredential",\n "apiVersion": "client.authentication.k8s.io/v1beta1",\n "spec": {\n  "interactive": false\n },\n "status": {\n  "token": "k8s-ack-v1.aHR0cHM6Ly9zd***"\n }\n}\n'})}),"\n",(0,i.jsx)(n.h2,{id:"\u547d\u4ee4\u884c\u53c2\u6570",children:"\u547d\u4ee4\u884c\u53c2\u6570"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{children:'Usage:\n  ack-ram-tool credential-plugin get-token [flags]\n\nFlags:\n      --api-version string   v1 or v1beta1 (default "v1beta1")\n  -c, --cluster-id string    The cluster id to use\n  -h, --help                 help for get-token\n      --role-arn string      Assume an RAM Role ARN when send request or sign token\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,i.jsx)(n.p,{children:"\u53c2\u6570\u8bf4\u660e\uff1a"}),"\n",(0,i.jsxs)(n.table,{children:[(0,i.jsx)(n.thead,{children:(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.th,{children:"\u53c2\u6570\u540d\u79f0"}),(0,i.jsx)(n.th,{children:"\u9ed8\u8ba4\u503c"}),(0,i.jsx)(n.th,{children:"\u5fc5\u9700\u53c2\u6570"}),(0,i.jsx)(n.th,{children:"\u8bf4\u660e"})]})}),(0,i.jsxs)(n.tbody,{children:[(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"-c, --cluster-id"}),(0,i.jsx)(n.td,{children:"\u65e0"}),(0,i.jsx)(n.td,{children:"\u662f"}),(0,i.jsx)(n.td,{children:"\u96c6\u7fa4 ID"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"--api-version"}),(0,i.jsx)(n.td,{children:"v1beta1"}),(0,i.jsx)(n.td,{children:"\u5426"}),(0,i.jsxs)(n.td,{children:["\u6307\u5b9a\u8fd4\u56de\u7684\u6570\u636e\u4e2d\u4f7f\u7528\u54ea\u4e2a\u7248\u672c\u7684 apiVersion\u3002v1beta1 \u8868\u793a ",(0,i.jsx)(n.code,{children:"client.authentication.k8s.io/v1beta1"}),"\uff0cv1 \u8868\u793a ",(0,i.jsx)(n.code,{children:"client.authentication.k8s.io/v1beta1"})]})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"--role-arn"}),(0,i.jsx)(n.td,{}),(0,i.jsx)(n.td,{children:"\u5426"}),(0,i.jsx)(n.td,{children:"\u4f7f\u7528\u626e\u6f14\u8fd9\u4e2a\u89d2\u8272\u540e\u7684\u8eab\u4efd\u8bbf\u95ee\u963f\u91cc\u4e91API"})]})]})]})]})}function h(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},8453:(e,n,t)=>{t.d(n,{R:()=>o,x:()=>l});var i=t(6540);const r={},s=i.createContext(r);function o(e){const n=i.useContext(s);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function l(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:o(e.components),i.createElement(s.Provider,{value:n},e.children)}}}]);