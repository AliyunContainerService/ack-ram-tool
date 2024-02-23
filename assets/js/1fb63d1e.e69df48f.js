"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[3757],{5159:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>o,contentTitle:()=>a,default:()=>u,frontMatter:()=>l,metadata:()=>s,toc:()=>c});var t=r(4848),i=r(8453);const l={slug:"/zh-CN/export-credentials/export-credentials",title:"export-credentials\uff08\u4e2d\u6587\uff09"},a="export-credentials",s={id:"export-credentials/export-credentials.zh-CN",title:"export-credentials\uff08\u4e2d\u6587\uff09",description:"\u5bfc\u51fa\u83b7\u53d6\u5230\u7684\u51ed\u8bc1\u4fe1\u606f\u6216\u8005\u4f7f\u7528\u8be5\u51ed\u8bc1\u6267\u884c\u5916\u90e8\u7a0b\u5e8f\u3002",source:"@site/docs/export-credentials/export-credentials.zh-CN.md",sourceDirName:"export-credentials",slug:"/zh-CN/export-credentials/export-credentials",permalink:"/ack-ram-tool/next/zh-CN/export-credentials/export-credentials",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/docs/export-credentials/export-credentials.zh-CN.md",tags:[],version:"current",frontMatter:{slug:"/zh-CN/export-credentials/export-credentials",title:"export-credentials\uff08\u4e2d\u6587\uff09"},sidebar:"tutorialSidebar",previous:{title:"export-credentials",permalink:"/ack-ram-tool/next/export-credentials/export-credentials"},next:{title:"Global Flags",permalink:"/ack-ram-tool/next/global-flags"}},o={},c=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u9ed8\u8ba4\u8f93\u51fa",id:"\u9ed8\u8ba4\u8f93\u51fa",level:3},{value:"--format aliyun-cli-uri-json",id:"--format-aliyun-cli-uri-json",level:3},{value:"--format ecs-metadata-json",id:"--format-ecs-metadata-json",level:3},{value:"--format credential-file-ini",id:"--format-credential-file-ini",level:3},{value:"--format environment-variables",id:"--format-environment-variables",level:3},{value:"--format aliyun-cli-uri-json --serve ADDR",id:"--format-aliyun-cli-uri-json---serve-addr",level:3},{value:"--format aliyun-cli-uri-json -- COMMAND [ARGS]",id:"--format-aliyun-cli-uri-json----command-args",level:3},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}];function d(e){const n={code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,i.R)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(n.h1,{id:"export-credentials",children:"export-credentials"}),"\n",(0,t.jsx)(n.p,{children:"\u5bfc\u51fa\u83b7\u53d6\u5230\u7684\u51ed\u8bc1\u4fe1\u606f\u6216\u8005\u4f7f\u7528\u8be5\u51ed\u8bc1\u6267\u884c\u5916\u90e8\u7a0b\u5e8f\u3002"}),"\n",(0,t.jsx)(n.h2,{id:"\u4f7f\u7528\u793a\u4f8b",children:"\u4f7f\u7528\u793a\u4f8b"}),"\n",(0,t.jsx)(n.h3,{id:"\u9ed8\u8ba4\u8f93\u51fa",children:"\u9ed8\u8ba4\u8f93\u51fa"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials\n{\n  "mode": "AK",\n  "access_key_id": "LT***",\n  "access_key_secret": "vHLE***"\n }\n'})}),"\n",(0,t.jsx)(n.h3,{id:"--format-aliyun-cli-uri-json",children:"--format aliyun-cli-uri-json"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials --format aliyun-cli-uri-json\n{\n  "Code": "Success",\n  "AccessKeyId": "LT***",\n  "AccessKeySecret": "vHLE***",\n  "SecurityToken": "",\n  "Expiration": "2023-04-20T12:09:37Z"\n }\n'})}),"\n",(0,t.jsx)(n.h3,{id:"--format-ecs-metadata-json",children:"--format ecs-metadata-json"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials --format ecs-metadata-json\n{\n  "Code": "Success",\n  "AccessKeyId": "LT***",\n  "AccessKeySecret": "vHLE***",\n  "SecurityToken": "",\n  "Expiration": "2023-04-20T12:11:04Z"\n }\n'})}),"\n",(0,t.jsx)(n.h3,{id:"--format-credential-file-ini",children:"--format credential-file-ini"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:"$ ack-ram-tool export-credentials --format credential-file-ini\n[default]\nenable = true\ntype = access_key\naccess_key_id = LT***\naccess_key_secret = vHLE***\n"})}),"\n",(0,t.jsx)(n.h3,{id:"--format-environment-variables",children:"--format environment-variables"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:"$ ack-ram-tool export-credentials --format environment-variables\n\nfor aliyun cli:\n\nexport ALIBABACLOUD_ACCESS_KEY_ID=LT***\nexport ALIBABACLOUD_ACCESS_KEY_SECRET=vHLE***\n\nfor terraform:\n\nexport ALICLOUD_ACCESS_KEY=LT***\nexport ALICLOUD_SECRET_KEY=vHLE***\n\nfor other tools:\n\nexport ALIBABA_CLOUD_ACCESS_KEY_ID=LT***\nexport ALICLOUD_ACCESS_KEY=LT***\nexport ALIBABACLOUD_ACCESS_KEY_ID=LT***\nexport ALICLOUD_SECRET_KEY=LT***\nexport ALIBABA_CLOUD_ACCESS_KEY_SECRET=vHLE***\nexport ALIBABACLOUD_ACCESS_KEY_SECRET=vHLE***\n"})}),"\n",(0,t.jsx)(n.h3,{id:"--format-aliyun-cli-uri-json---serve-addr",children:"--format aliyun-cli-uri-json --serve ADDR"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials --format aliyun-cli-uri-json --serve 127.0.0.1:1234\n2023-04-20T20:05:40+08:00 WARN Serving HTTP on 127.0.0.1:1234\n$ curl http://127.0.0.1:1234\n{\n  "Code": "Success",\n  "AccessKeyId": "LT***",\n  "AccessKeySecret": "vHLE***",\n  "SecurityToken": "",\n  "Expiration": "2023-04-20T12:14:15Z"\n }\n'})}),"\n",(0,t.jsx)(n.h3,{id:"--format-aliyun-cli-uri-json----command-args",children:"--format aliyun-cli-uri-json -- COMMAND [ARGS]"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials --format environment-variables -- aliyun sts GetCallerIdentity\n{\n\t"AccountId": "113***",\n\t"Arn": "acs:ram::113***:user/***",\n\t"IdentityType": "RAMUser",\n\t"PrincipalId": "272***",\n\t"RequestId": "28B93***",\n\t"UserId": "272***"\n}\n'})}),"\n",(0,t.jsx)(n.h2,{id:"\u547d\u4ee4\u884c\u53c2\u6570",children:"\u547d\u4ee4\u884c\u53c2\u6570"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:'Usage:\n  ack-ram-tool export-credentials [flags]\n\nFlags:\n  -f, --format string   The output format to display credentials (aliyun-cli-config-json, aliyun-cli-uri-json, ecs-metadata-json, credential-file-ini, environment-variables) (default "aliyun-cli-config-json")\n  -h, --help            help for export-credentials\n  -s, --serve string    start a server to export credentials\n  --role-arn string     Assume an RAM Role ARN when send request or sign token\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,t.jsx)(n.p,{children:"\u53c2\u6570\u8bf4\u660e\uff1a"}),"\n",(0,t.jsxs)(n.table,{children:[(0,t.jsx)(n.thead,{children:(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.th,{children:"\u53c2\u6570\u540d\u79f0"}),(0,t.jsx)(n.th,{children:"\u9ed8\u8ba4\u503c"}),(0,t.jsx)(n.th,{children:"\u5fc5\u9700\u53c2\u6570"}),(0,t.jsx)(n.th,{children:"\u8bf4\u660e"})]})}),(0,t.jsxs)(n.tbody,{children:[(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"-f, --format"}),(0,t.jsx)(n.td,{children:"aliyun-cli-config-json"}),(0,t.jsx)(n.td,{children:"\u5426"}),(0,t.jsx)(n.td,{children:"\u6307\u5b9a\u8f93\u51fa\u683c\u5f0f\uff0c\u8be6\u89c1\u4f7f\u7528\u793a\u4f8b"})]}),(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"-s, --serve"}),(0,t.jsx)(n.td,{children:"\u65e0"}),(0,t.jsx)(n.td,{children:"\u5426"}),(0,t.jsx)(n.td,{children:"\u542f\u52a8\u4e00\u4e2a\u76d1\u542c\u6307\u5b9a\u5730\u5740\u7684 HTTP Server\uff0c\u8bbf\u95ee\u8be5\u670d\u52a1\u5c06\u8fd4\u56de\u51ed\u8bc1\u4fe1\u606f\uff0c\u8be6\u89c1\u4f7f\u7528\u793a\u4f8b"})]}),(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"--role-arn"}),(0,t.jsx)(n.td,{children:"\u65e0"}),(0,t.jsx)(n.td,{children:"\u5426"}),(0,t.jsx)(n.td,{children:"\u626e\u6f14\u8be5\u89d2\u8272"})]})]})]})]})}function u(e={}){const{wrapper:n}={...(0,i.R)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(d,{...e})}):d(e)}},8453:(e,n,r)=>{r.d(n,{R:()=>a,x:()=>s});var t=r(6540);const i={},l=t.createContext(i);function a(e){const n=t.useContext(l);return t.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function s(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:a(e.components),t.createElement(l.Provider,{value:n},e.children)}}}]);