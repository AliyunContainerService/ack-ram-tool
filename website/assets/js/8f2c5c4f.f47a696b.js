"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[548],{3905:(e,t,r)=>{r.d(t,{Zo:()=>c,kt:()=>f});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function l(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?l(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):l(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},l=Object.keys(e);for(n=0;n<l.length;n++)r=l[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(n=0;n<l.length;n++)r=l[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var s=n.createContext({}),u=function(e){var t=n.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},c=function(e){var t=u(e.components);return n.createElement(s.Provider,{value:t},e.children)},p="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},m=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,l=e.originalType,s=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),p=u(r),m=a,f=p["".concat(s,".").concat(m)]||p[m]||d[m]||l;return r?n.createElement(f,o(o({ref:t},c),{},{components:r})):n.createElement(f,o({ref:t},c))}));function f(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=r.length,o=new Array(l);o[0]=m;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[p]="string"==typeof e?e:a,o[1]=i;for(var u=2;u<l;u++)o[u]=r[u];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}m.displayName="MDXCreateElement"},2182:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>s,contentTitle:()=>o,default:()=>d,frontMatter:()=>l,metadata:()=>i,toc:()=>u});var n=r(7462),a=(r(7294),r(3905));const l={slug:"/zh-CN/rrsa/assume-role",title:"assume-role\uff08\u4e2d\u6587\uff09",sidebar_position:5},o="assume-role",i={unversionedId:"rrsa/assume-role.zh-CN",id:"rrsa/assume-role.zh-CN",title:"assume-role\uff08\u4e2d\u6587\uff09",description:"\u6d4b\u8bd5\u4f7f\u7528 oidc token \u626e\u6f14\u7279\u5b9a RAM \u89d2\u8272\u3002",source:"@site/docs/rrsa/assume-role.zh-CN.md",sourceDirName:"rrsa",slug:"/zh-CN/rrsa/assume-role",permalink:"/zh-CN/rrsa/assume-role",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/docs/rrsa/assume-role.zh-CN.md",tags:[],version:"current",sidebarPosition:5,frontMatter:{slug:"/zh-CN/rrsa/assume-role",title:"assume-role\uff08\u4e2d\u6587\uff09",sidebar_position:5},sidebar:"tutorialSidebar",previous:{title:"assume-role",permalink:"/rrsa/assume-role"},next:{title:"credential-plugin",permalink:"/category/credential-plugin"}},s={},u=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}],c={toc:u},p="wrapper";function d(e){let{components:t,...r}=e;return(0,a.kt)(p,(0,n.Z)({},c,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"assume-role"},"assume-role"),(0,a.kt)("p",null,"\u6d4b\u8bd5\u4f7f\u7528 oidc token \u626e\u6f14\u7279\u5b9a RAM \u89d2\u8272\u3002"),(0,a.kt)("h2",{id:"\u4f7f\u7528\u793a\u4f8b"},"\u4f7f\u7528\u793a\u4f8b"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"$ ack-ram-tool rrsa assume-role --oidc-provider-arn <oidcProviderArn> \\\n  --role-arn <roleArn> --oidc-token-file <pathToTokenFile>\n\n    Retrieved a STS token:\n    AccessKeyId:       STS.***\n    AccessKeySecret:   7UVy***\n    SecurityToken:     CAIS***\n    Expiration:        2021-12-03T05:51:37Z\n\n")),(0,a.kt)("h2",{id:"\u547d\u4ee4\u884c\u53c2\u6570"},"\u547d\u4ee4\u884c\u53c2\u6570"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'Usage:\n  ack-ram-tool rrsa assume-role [flags]\n\nFlags:\n  -h, --help                       help for assume-role\n  -p, --oidc-provider-arn string   The arn of OIDC provider\n  -t, --oidc-token-file string     Path to OIDC token file. If value is \'-\', will read token from stdin\n  -r, --role-arn string            The arn of RAM role\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n')),(0,a.kt)("p",null,"\u53c2\u6570\u8bf4\u660e\uff1a"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"\u53c2\u6570\u540d\u79f0"),(0,a.kt)("th",{parentName:"tr",align:null},"\u9ed8\u8ba4\u503c"),(0,a.kt)("th",{parentName:"tr",align:null},"\u5fc5\u9700\u53c2\u6570"),(0,a.kt)("th",{parentName:"tr",align:null},"\u8bf4\u660e"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-p, --oidc-provider-arn"),(0,a.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,a.kt)("td",{parentName:"tr",align:null},"\u4e3a\u96c6\u7fa4\u6ce8\u518c\u7684 RAM \u89d2\u8272 SSO \u4f9b\u5e94\u5546 ARN")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-r, --role-arn"),(0,a.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,a.kt)("td",{parentName:"tr",align:null},"\u88ab\u626e\u6f14\u7684 RAM \u89d2\u8272\u7684 ARN")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-t, --oidc-token-file"),(0,a.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,a.kt)("td",{parentName:"tr",align:null},"oidc token \u6587\u4ef6\u7684\u8def\u5f84\u3002\u5f53\u503c\u4e3a ",(0,a.kt)("inlineCode",{parentName:"td"},"-")," \u65f6\u652f\u6301\u4ece\u6807\u51c6\u8f93\u5165\u4ece\u8bfb\u53d6 token\uff08\u6bd4\u5982\u901a\u8fc7\u7ba1\u9053\u4f20\u9012 token\uff09")))))}d.isMDXComponent=!0}}]);