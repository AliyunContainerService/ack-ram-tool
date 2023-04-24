"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[2266],{3905:(e,t,r)=>{r.d(t,{Zo:()=>p,kt:()=>h});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function l(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?l(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):l(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},l=Object.keys(e);for(n=0;n<l.length;n++)r=l[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(n=0;n<l.length;n++)r=l[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var s=n.createContext({}),d=function(e){var t=n.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},p=function(e){var t=d(e.components);return n.createElement(s.Provider,{value:t},e.children)},c="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},m=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,l=e.originalType,s=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),c=d(r),m=a,h=c["".concat(s,".").concat(m)]||c[m]||u[m]||l;return r?n.createElement(h,o(o({ref:t},p),{},{components:r})):n.createElement(h,o({ref:t},p))}));function h(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=r.length,o=new Array(l);o[0]=m;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i[c]="string"==typeof e?e:a,o[1]=i;for(var d=2;d<l;d++)o[d]=r[d];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}m.displayName="MDXCreateElement"},4436:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>s,contentTitle:()=>o,default:()=>u,frontMatter:()=>l,metadata:()=>i,toc:()=>d});var n=r(7462),a=(r(7294),r(3905));const l={slug:"/zh-CN/rrsa/install-helper-addon",title:"install-helper-addon\uff08\u4e2d\u6587\uff09",sidebar_position:3},o="install-helper-addon",i={unversionedId:"rrsa/install-helper-addon.zh-CN",id:"version-v0.13.0/rrsa/install-helper-addon.zh-CN",title:"install-helper-addon\uff08\u4e2d\u6587\uff09",description:"\u5728\u96c6\u7fa4\u5185\u5b89\u88c5 RRSA \u8f85\u52a9\u7ec4\u4ef6 ack-pod-identity-webhook\u3002",source:"@site/versioned_docs/version-v0.13.0/rrsa/install-helper-addon.zh-CN.md",sourceDirName:"rrsa",slug:"/zh-CN/rrsa/install-helper-addon",permalink:"/ack-ram-tool/zh-CN/rrsa/install-helper-addon",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.13.0/rrsa/install-helper-addon.zh-CN.md",tags:[],version:"v0.13.0",sidebarPosition:3,frontMatter:{slug:"/zh-CN/rrsa/install-helper-addon",title:"install-helper-addon\uff08\u4e2d\u6587\uff09",sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"install-helper-addon",permalink:"/ack-ram-tool/rrsa/install-helper-addon"},next:{title:"status",permalink:"/ack-ram-tool/rrsa/status"}},s={},d=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}],p={toc:d},c="wrapper";function u(e){let{components:t,...r}=e;return(0,a.kt)(c,(0,n.Z)({},p,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"install-helper-addon"},"install-helper-addon"),(0,a.kt)("p",null,"\u5728\u96c6\u7fa4\u5185\u5b89\u88c5 RRSA \u8f85\u52a9\u7ec4\u4ef6 ",(0,a.kt)("a",{parentName:"p",href:"https://www.alibabacloud.com/help/doc-detail/600451.html"},"ack-pod-identity-webhook"),"\u3002"),(0,a.kt)("h2",{id:"\u4f7f\u7528\u793a\u4f8b"},"\u4f7f\u7528\u793a\u4f8b"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"$ ack-ram-tool rrsa install-helper-addon --cluster-id <clusterId>\n\n? Are you sure you want to install ack-pod-identity-webhook? Yes\n2023-04-20T15:39:41+08:00 INFO Start to install ack-pod-identity-webhook\n2023-04-20T15:40:49+08:00 INFO Install ack-pod-identity-webhook for cluster c12d3*** successfully\n")),(0,a.kt)("h2",{id:"\u547d\u4ee4\u884c\u53c2\u6570"},"\u547d\u4ee4\u884c\u53c2\u6570"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'Usage:\n  ack-ram-tool rrsa install-helper-addon [flags]\n\nFlags:\n  -c, --cluster-id string   The cluster id to use\n  -h, --help                help for install-helper-addon\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n')),(0,a.kt)("p",null,"\u53c2\u6570\u8bf4\u660e\uff1a"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"\u53c2\u6570\u540d\u79f0"),(0,a.kt)("th",{parentName:"tr",align:null},"\u9ed8\u8ba4\u503c"),(0,a.kt)("th",{parentName:"tr",align:null},"\u5fc5\u9700\u53c2\u6570"),(0,a.kt)("th",{parentName:"tr",align:null},"\u8bf4\u660e"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-c, --cluster-id"),(0,a.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,a.kt)("td",{parentName:"tr",align:null},"\u96c6\u7fa4 ID")))))}u.isMDXComponent=!0}}]);