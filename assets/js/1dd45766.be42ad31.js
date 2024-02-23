"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[7980],{8077:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>o,contentTitle:()=>s,default:()=>u,frontMatter:()=>i,metadata:()=>l,toc:()=>c});var t=r(4848),a=r(8453);const i={slug:"export-credentials"},s="export-credentials",l={id:"export-credentials/export-credentials",title:"export-credentials",description:"Export the obtained credential information or use the credential to execute an external program.",source:"@site/docs/export-credentials/export-credentials.md",sourceDirName:"export-credentials",slug:"/export-credentials/export-credentials",permalink:"/ack-ram-tool/next/export-credentials/export-credentials",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/docs/export-credentials/export-credentials.md",tags:[],version:"current",frontMatter:{slug:"export-credentials"},sidebar:"tutorialSidebar",previous:{title:"export-credentials",permalink:"/ack-ram-tool/next/category/export-credentials"},next:{title:"export-credentials\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/next/zh-CN/export-credentials/export-credentials"}},o={},c=[{value:"Usage",id:"usage",level:2},{value:"default",id:"default",level:3},{value:"--format aliyun-cli-uri-json",id:"--format-aliyun-cli-uri-json",level:3},{value:"--format ecs-metadata-json",id:"--format-ecs-metadata-json",level:3},{value:"--format credential-file-ini",id:"--format-credential-file-ini",level:3},{value:"--format environment-variables",id:"--format-environment-variables",level:3},{value:"--format aliyun-cli-uri-json --serve ADDR",id:"--format-aliyun-cli-uri-json---serve-addr",level:3},{value:"--format aliyun-cli-uri-json -- COMMAND [ARGS]",id:"--format-aliyun-cli-uri-json----command-args",level:3},{value:"Flags",id:"flags",level:2}];function d(e){const n={code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,a.R)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(n.h1,{id:"export-credentials",children:"export-credentials"}),"\n",(0,t.jsx)(n.p,{children:"Export the obtained credential information or use the credential to execute an external program."}),"\n",(0,t.jsx)(n.h2,{id:"usage",children:"Usage"}),"\n",(0,t.jsx)(n.h3,{id:"default",children:"default"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials\n{\n  "mode": "AK",\n  "access_key_id": "LT***",\n  "access_key_secret": "vHLE***"\n }\n'})}),"\n",(0,t.jsx)(n.h3,{id:"--format-aliyun-cli-uri-json",children:"--format aliyun-cli-uri-json"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials --format aliyun-cli-uri-json\n{\n  "Code": "Success",\n  "AccessKeyId": "LT***",\n  "AccessKeySecret": "vHLE***",\n  "SecurityToken": "",\n  "Expiration": "2023-04-20T12:09:37Z"\n }\n'})}),"\n",(0,t.jsx)(n.h3,{id:"--format-ecs-metadata-json",children:"--format ecs-metadata-json"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials --format ecs-metadata-json\n{\n  "Code": "Success",\n  "AccessKeyId": "LT***",\n  "AccessKeySecret": "vHLE***",\n  "SecurityToken": "",\n  "Expiration": "2023-04-20T12:11:04Z"\n }\n'})}),"\n",(0,t.jsx)(n.h3,{id:"--format-credential-file-ini",children:"--format credential-file-ini"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:"$ ack-ram-tool export-credentials --format credential-file-ini\n[default]\nenable = true\ntype = access_key\naccess_key_id = LT***\naccess_key_secret = vHLE***\n"})}),"\n",(0,t.jsx)(n.h3,{id:"--format-environment-variables",children:"--format environment-variables"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:"$ ack-ram-tool export-credentials --format environment-variables\n\nfor aliyun cli:\n\nexport ALIBABACLOUD_ACCESS_KEY_ID=LT***\nexport ALIBABACLOUD_ACCESS_KEY_SECRET=vHLE***\n\nfor terraform:\n\nexport ALICLOUD_ACCESS_KEY=LT***\nexport ALICLOUD_SECRET_KEY=vHLE***\n\nfor other tools:\n\nexport ALIBABA_CLOUD_ACCESS_KEY_ID=LT***\nexport ALICLOUD_ACCESS_KEY=LT***\nexport ALIBABACLOUD_ACCESS_KEY_ID=LT***\nexport ALICLOUD_SECRET_KEY=LT***\nexport ALIBABA_CLOUD_ACCESS_KEY_SECRET=vHLE***\nexport ALIBABACLOUD_ACCESS_KEY_SECRET=vHLE***\n"})}),"\n",(0,t.jsx)(n.h3,{id:"--format-aliyun-cli-uri-json---serve-addr",children:"--format aliyun-cli-uri-json --serve ADDR"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials --format aliyun-cli-uri-json --serve 127.0.0.1:1234\n2023-04-20T20:05:40+08:00 WARN Serving HTTP on 127.0.0.1:1234\n$ curl http://127.0.0.1:1234\n{\n  "Code": "Success",\n  "AccessKeyId": "LT***",\n  "AccessKeySecret": "vHLE***",\n  "SecurityToken": "",\n  "Expiration": "2023-04-20T12:14:15Z"\n }\n'})}),"\n",(0,t.jsx)(n.h3,{id:"--format-aliyun-cli-uri-json----command-args",children:"--format aliyun-cli-uri-json -- COMMAND [ARGS]"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool export-credentials --format environment-variables -- aliyun sts GetCallerIdentity\n{\n\t"AccountId": "113***",\n\t"Arn": "acs:ram::113***:user/***",\n\t"IdentityType": "RAMUser",\n\t"PrincipalId": "272***",\n\t"RequestId": "28B93***",\n\t"UserId": "272***"\n}\n'})}),"\n",(0,t.jsx)(n.h2,{id:"flags",children:"Flags"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:'Usage:\n  ack-ram-tool export-credentials [flags]\n\nFlags:\n  -f, --format string   The output format to display credentials (aliyun-cli-config-json, aliyun-cli-uri-json, ecs-metadata-json, credential-file-ini, environment-variables) (default "aliyun-cli-config-json")\n  -h, --help            help for export-credentials\n  -s, --serve string    start a server to export credentials\n  --role-arn string     Assume an RAM Role ARN when send request or sign token\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,t.jsx)(n.p,{children:"Descriptions\uff1a"}),"\n",(0,t.jsxs)(n.table,{children:[(0,t.jsx)(n.thead,{children:(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.th,{children:"Flag"}),(0,t.jsx)(n.th,{children:"Default"}),(0,t.jsx)(n.th,{children:"Required"}),(0,t.jsx)(n.th,{children:"Description"})]})}),(0,t.jsxs)(n.tbody,{children:[(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"-f, --format"}),(0,t.jsx)(n.td,{children:"aliyun-cli-config-json"}),(0,t.jsx)(n.td,{}),(0,t.jsx)(n.td,{children:"Specify the output format, see usage examples for details."})]}),(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"-s, --serve"}),(0,t.jsx)(n.td,{}),(0,t.jsx)(n.td,{}),(0,t.jsx)(n.td,{children:"Start an HTTP server listening on a specified address, accessing the service will return credential information. See usage examples for details."})]}),(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"-role-arn"}),(0,t.jsx)(n.td,{}),(0,t.jsx)(n.td,{}),(0,t.jsx)(n.td,{children:"Assume an RAM Role ARN when send request or sign token"})]})]})]})]})}function u(e={}){const{wrapper:n}={...(0,a.R)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(d,{...e})}):d(e)}},8453:(e,n,r)=>{r.d(n,{R:()=>s,x:()=>l});var t=r(6540);const a={},i=t.createContext(a);function s(e){const n=t.useContext(i);return t.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function l(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(a):e.components||a:s(e.components),t.createElement(i.Provider,{value:n},e.children)}}}]);