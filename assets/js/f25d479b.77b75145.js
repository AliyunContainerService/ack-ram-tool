"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[4349],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>g});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},l=Object.keys(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=a.createContext({}),c=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},u=function(e){var t=c(e.components);return a.createElement(s.Provider,{value:t},e.children)},p="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},d=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,l=e.originalType,s=e.parentName,u=i(e,["components","mdxType","originalType","parentName"]),p=c(n),d=r,g=p["".concat(s,".").concat(d)]||p[d]||m[d]||l;return n?a.createElement(g,o(o({ref:t},u),{},{components:n})):a.createElement(g,o({ref:t},u))}));function g(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var l=n.length,o=new Array(l);o[0]=d;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[p]="string"==typeof e?e:r,o[1]=i;for(var c=2;c<l;c++)o[c]=n[c];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}d.displayName="MDXCreateElement"},8450:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>o,default:()=>m,frontMatter:()=>l,metadata:()=>i,toc:()=>c});var a=n(7462),r=(n(7294),n(3905));const l={slug:"/zh-CN/rrsa/associate-role",title:"associate-role\uff08\u4e2d\u6587\uff09",sidebar_position:2},o="associate-role",i={unversionedId:"rrsa/associate-role.zh-CN",id:"version-v0.13.2/rrsa/associate-role.zh-CN",title:"associate-role\uff08\u4e2d\u6587\uff09",description:"\u914d\u7f6e RAM \u89d2\u8272\uff0c\u5141\u8bb8\u4f7f\u7528\u8868\u793a\u7279\u5b9a service account \u8eab\u4efd\u7684 oidc token \u626e\u6f14\u8be5 RAM \u89d2\u8272\u3002",source:"@site/versioned_docs/version-v0.13.2/rrsa/associate-role.zh-CN.md",sourceDirName:"rrsa",slug:"/zh-CN/rrsa/associate-role",permalink:"/ack-ram-tool/v0.13.2/zh-CN/rrsa/associate-role",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.13.2/rrsa/associate-role.zh-CN.md",tags:[],version:"v0.13.2",sidebarPosition:2,frontMatter:{slug:"/zh-CN/rrsa/associate-role",title:"associate-role\uff08\u4e2d\u6587\uff09",sidebar_position:2},sidebar:"tutorialSidebar",previous:{title:"associate-role",permalink:"/ack-ram-tool/v0.13.2/rrsa/associate-role"},next:{title:"install-helper-addon",permalink:"/ack-ram-tool/v0.13.2/rrsa/install-helper-addon"}},s={},c=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}],u={toc:c},p="wrapper";function m(e){let{components:t,...n}=e;return(0,r.kt)(p,(0,a.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"associate-role"},"associate-role"),(0,r.kt)("p",null,"\u914d\u7f6e RAM \u89d2\u8272\uff0c\u5141\u8bb8\u4f7f\u7528\u8868\u793a\u7279\u5b9a service account \u8eab\u4efd\u7684 oidc token \u626e\u6f14\u8be5 RAM \u89d2\u8272\u3002"),(0,r.kt)("h2",{id:"\u4f7f\u7528\u793a\u4f8b"},"\u4f7f\u7528\u793a\u4f8b"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},'$ ack-ram-tool rrsa associate-role --cluster-id <clusterId> \\\n  --namespace <namespce> --service-account <serviceAccountName> \\\n  --role-name <roleName>\n\n? Are you sure you want to associate RAM Role "<roleName>" to service account "<serviceAccountName>" (namespace: "<namespce>")? Yes\n2023-04-20T14:30:02+08:00 INFO will change the AssumeRole Policy of RAM Role "<roleName>" with blow content:\n{\n  "Statement": [\n   {\n    "Action": "sts:AssumeRole",\n    "Condition": {\n     "StringEquals": {\n      "oidc:aud": "sts.aliyuncs.com",\n      "oidc:iss": "https://oidc-ack-***.aliyuncs.com/c132c***",\n      "oidc:sub": "system:serviceaccount:<namespce>:<serviceAccountName>"\n     }\n    },\n    "Effect": "Allow",\n    "Principal": {\n     "Federated": [\n      "acs:ram::113***:oidc-provider/ack-rrsa-c132c***"\n     ]\n    }\n   }\n  ],\n  "Version": "1"\n }\n\n? Are you sure you want to associate RAM Role "test" to service account "sa" (namespace: "test")? Yes\n2023-04-20T14:30:04+08:00 INFO Associate RAM Role "test" to service account "sa" (namespace: "test") successfully\n')),(0,r.kt)("h2",{id:"\u547d\u4ee4\u884c\u53c2\u6570"},"\u547d\u4ee4\u884c\u53c2\u6570"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},'Usage:\n  ack-ram-tool rrsa associate-role [flags]\n\nFlags:\n      --attach-custom-policy string   Attach this custom policy to the RAM Role\n      --attach-system-policy string   Attach this system policy to the RAM Role\n  -c, --cluster-id string             The cluster id to use\n      --create-role-if-not-exist      Create the RAM Role if it does not exist\n  -h, --help                          help for associate-role\n  -n, --namespace string              The Kubernetes namespace to use\n  -r, --role-name string              The RAM Role name to use\n  -s, --service-account string        The Kubernetes service account to use\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n')),(0,r.kt)("p",null,"\u53c2\u6570\u8bf4\u660e\uff1a"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"\u53c2\u6570\u540d\u79f0"),(0,r.kt)("th",{parentName:"tr",align:null},"\u9ed8\u8ba4\u503c"),(0,r.kt)("th",{parentName:"tr",align:null},"\u5fc5\u9700\u53c2\u6570"),(0,r.kt)("th",{parentName:"tr",align:null},"\u8bf4\u660e"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"-c, --cluster-id"),(0,r.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,r.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,r.kt)("td",{parentName:"tr",align:null},"\u96c6\u7fa4 ID")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"-n, --namespace"),(0,r.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,r.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,r.kt)("td",{parentName:"tr",align:null},"\u547d\u540d\u7a7a\u95f4\uff0c\u53ef\u4ee5\u4f7f\u7528 ",(0,r.kt)("inlineCode",{parentName:"td"},"*")," \u8868\u793a\u6240\u6709\u547d\u540d\u7a7a\u95f4\uff1a",(0,r.kt)("inlineCode",{parentName:"td"},"--namespace '*'"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"-s, --service-account"),(0,r.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,r.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,r.kt)("td",{parentName:"tr",align:null},"service account")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"-r, --role-name"),(0,r.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,r.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,r.kt)("td",{parentName:"tr",align:null},"RAM \u89d2\u8272")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"--create-role-if-not-exist"),(0,r.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,r.kt)("td",{parentName:"tr",align:null},"\u5426"),(0,r.kt)("td",{parentName:"tr",align:null},"\u5982\u679c\u8be5 RAM \u89d2\u8272\u4e0d\u5b58\u5728\uff0c\u90a3\u4e48\u81ea\u52a8\u521b\u5efa\u4e00\u4e2a\u540c\u540d\u7684 RAM \u89d2\u8272")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"--attach-system-policy"),(0,r.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,r.kt)("td",{parentName:"tr",align:null},"\u4e3a\u8be5\u89d2\u8272\u6388\u4e88\u6307\u5b9a\u7684\u7cfb\u7edf\u6743\u9650\u7b56\u7565"),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"--attach-custom-policy"),(0,r.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,r.kt)("td",{parentName:"tr",align:null},"\u4e3a\u8be5\u89d2\u8272\u6388\u4e88\u6307\u5b9a\u7684\u81ea\u5b9a\u4e49\u6743\u9650\u7b56\u7565"),(0,r.kt)("td",{parentName:"tr",align:null})))))}m.isMDXComponent=!0}}]);