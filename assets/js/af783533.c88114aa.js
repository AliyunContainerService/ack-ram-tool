"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[4656],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>m});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),s=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},p=function(e){var t=s(e.components);return r.createElement(c.Provider,{value:t},e.children)},u="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},g=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,c=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),u=s(n),g=a,m=u["".concat(c,".").concat(g)]||u[g]||d[g]||i;return n?r.createElement(m,l(l({ref:t},p),{},{components:n})):r.createElement(m,l({ref:t},p))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,l=new Array(i);l[0]=g;var o={};for(var c in t)hasOwnProperty.call(t,c)&&(o[c]=t[c]);o.originalType=e,o[u]="string"==typeof e?e:a,l[1]=o;for(var s=2;s<i;s++)l[s]=n[s];return r.createElement.apply(null,l)}return r.createElement.apply(null,n)}g.displayName="MDXCreateElement"},6092:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>d,frontMatter:()=>i,metadata:()=>o,toc:()=>s});var r=n(7462),a=(n(7294),n(3905));const i={slug:"/zh-CN/credential-plugin/get-token",title:"get-token\uff08\u4e2d\u6587\uff09",sidebar_position:3},l="get-token",o={unversionedId:"credential-plugin/get-token.zh-CN",id:"credential-plugin/get-token.zh-CN",title:"get-token\uff08\u4e2d\u6587\uff09",description:"\u96c6\u6210 ack-ram-authenticator\uff0c\u83b7\u53d6\u7528\u4e8e\u8bbf\u95ee api server \u7684 ExecCredential token\u3002",source:"@site/docs/credential-plugin/get-token.zh-CN.md",sourceDirName:"credential-plugin",slug:"/zh-CN/credential-plugin/get-token",permalink:"/ack-ram-tool/next/zh-CN/credential-plugin/get-token",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/docs/credential-plugin/get-token.zh-CN.md",tags:[],version:"current",sidebarPosition:3,frontMatter:{slug:"/zh-CN/credential-plugin/get-token",title:"get-token\uff08\u4e2d\u6587\uff09",sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"get-token",permalink:"/ack-ram-tool/next/credential-plugin/get-token"},next:{title:"export-credentials",permalink:"/ack-ram-tool/next/category/export-credentials"}},c={},s=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}],p={toc:s},u="wrapper";function d(e){let{components:t,...n}=e;return(0,a.kt)(u,(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"get-token"},"get-token"),(0,a.kt)("p",null,"\u96c6\u6210 ack-ram-authenticator\uff0c\u83b7\u53d6\u7528\u4e8e\u8bbf\u95ee api server \u7684 ",(0,a.kt)("a",{parentName:"p",href:"https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins"},"ExecCredential")," token\u3002"),(0,a.kt)("h2",{id:"\u4f7f\u7528\u793a\u4f8b"},"\u4f7f\u7528\u793a\u4f8b"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'$ ack-ram-tool credential-plugin get-token --cluster-id <clusterId>\n\n{\n "kind": "ExecCredential",\n "apiVersion": "client.authentication.k8s.io/v1beta1",\n "spec": {\n  "interactive": false\n },\n "status": {\n  "token": "k8s-ack-v1.aHR0cHM6Ly9zd***"\n }\n}\n')),(0,a.kt)("h2",{id:"\u547d\u4ee4\u884c\u53c2\u6570"},"\u547d\u4ee4\u884c\u53c2\u6570"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'Usage:\n  ack-ram-tool credential-plugin get-token [flags]\n\nFlags:\n      --api-version string   v1 or v1beta1 (default "v1beta1")\n  -c, --cluster-id string    The cluster id to use\n  -h, --help                 help for get-token\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n')),(0,a.kt)("p",null,"\u53c2\u6570\u8bf4\u660e\uff1a"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"\u53c2\u6570\u540d\u79f0"),(0,a.kt)("th",{parentName:"tr",align:null},"\u9ed8\u8ba4\u503c"),(0,a.kt)("th",{parentName:"tr",align:null},"\u5fc5\u9700\u53c2\u6570"),(0,a.kt)("th",{parentName:"tr",align:null},"\u8bf4\u660e"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-c, --cluster-id"),(0,a.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,a.kt)("td",{parentName:"tr",align:null},"\u96c6\u7fa4 ID")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--api-version"),(0,a.kt)("td",{parentName:"tr",align:null},"v1beta1"),(0,a.kt)("td",{parentName:"tr",align:null},"\u5426"),(0,a.kt)("td",{parentName:"tr",align:null},"\u6307\u5b9a\u8fd4\u56de\u7684\u6570\u636e\u4e2d\u4f7f\u7528\u54ea\u4e2a\u7248\u672c\u7684 apiVersion\u3002v1beta1 \u8868\u793a ",(0,a.kt)("inlineCode",{parentName:"td"},"client.authentication.k8s.io/v1beta1"),"\uff0cv1 \u8868\u793a ",(0,a.kt)("inlineCode",{parentName:"td"},"client.authentication.k8s.io/v1beta1"))))))}d.isMDXComponent=!0}}]);