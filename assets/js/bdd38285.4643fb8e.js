"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[8648],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>m});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),u=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},p=function(e){var t=u(e.components);return r.createElement(c.Provider,{value:t},e.children)},d="mdxType",s={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},g=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,c=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),d=u(n),g=a,m=d["".concat(c,".").concat(g)]||d[g]||s[g]||i;return n?r.createElement(m,l(l({ref:t},p),{},{components:n})):r.createElement(m,l({ref:t},p))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,l=new Array(i);l[0]=g;var o={};for(var c in t)hasOwnProperty.call(t,c)&&(o[c]=t[c]);o.originalType=e,o[d]="string"==typeof e?e:a,l[1]=o;for(var u=2;u<i;u++)l[u]=n[u];return r.createElement.apply(null,l)}return r.createElement.apply(null,n)}g.displayName="MDXCreateElement"},4500:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>s,frontMatter:()=>i,metadata:()=>o,toc:()=>u});var r=n(7462),a=(n(7294),n(3905));const i={slug:"/zh-CN/credential-plugin/kubeconfig",title:"get-kubeconfig\uff08\u4e2d\u6587\uff09",sidebar_position:1},l="get-kubeconfig",o={unversionedId:"credential-plugin/get-kubeconfig.zh-CN",id:"credential-plugin/get-kubeconfig.zh-CN",title:"get-kubeconfig\uff08\u4e2d\u6587\uff09",description:"\u83b7\u53d6\u4f7f\u7528 ack-ram-tool \u4f5c\u4e3a credential plugin \u7684 kubeconfig\u3002",source:"@site/docs/credential-plugin/get-kubeconfig.zh-CN.md",sourceDirName:"credential-plugin",slug:"/zh-CN/credential-plugin/kubeconfig",permalink:"/ack-ram-tool/next/zh-CN/credential-plugin/kubeconfig",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/docs/credential-plugin/get-kubeconfig.zh-CN.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{slug:"/zh-CN/credential-plugin/kubeconfig",title:"get-kubeconfig\uff08\u4e2d\u6587\uff09",sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"get-kubeconfig",permalink:"/ack-ram-tool/next/credential-plugin/get-kubeconfig"},next:{title:"get-credential",permalink:"/ack-ram-tool/next/credential-plugin/get-credential"}},c={},u=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}],p={toc:u},d="wrapper";function s(e){let{components:t,...n}=e;return(0,a.kt)(d,(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"get-kubeconfig"},"get-kubeconfig"),(0,a.kt)("p",null,"\u83b7\u53d6\u4f7f\u7528 ack-ram-tool \u4f5c\u4e3a ",(0,a.kt)("a",{parentName:"p",href:"https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins"},"credential plugin")," \u7684 kubeconfig\u3002"),(0,a.kt)("p",null,"\u5305\u542b\u5982\u4e0b\u7279\u6027\uff1a"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"\u8bc1\u4e66\u8fc7\u671f\u524d\u5c06\u81ea\u52a8\u83b7\u53d6\u65b0\u7684\u8bc1\u4e66"),(0,a.kt)("li",{parentName:"ul"},"\u652f\u6301\u4f7f\u7528\u4e34\u65f6\u8bc1\u4e66"),(0,a.kt)("li",{parentName:"ul"},"\u96c6\u6210 ack-ram-authenticator")),(0,a.kt)("h2",{id:"\u4f7f\u7528\u793a\u4f8b"},"\u4f7f\u7528\u793a\u4f8b"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'$ ack-ram-tool credential-plugin get-kubeconfig --cluster-id c5e***\n\nkind: Config\napiVersion: v1\nclusters:\n    - name: kubernetes\n      cluster:\n        server: https://106.*.*.*:6443\n        certificate-authority-data: LS0tL***\ncontexts:\n    - name: 272***-c5e***\n      context:\n        cluster: kubernetes\n        user: "272***"\ncurrent-context: 272***-c5e***\nusers:\n    - name: "272***"\n      user:\n        exec:\n            command: ack-ram-tool\n            args:\n                - credential-plugin\n                - get-credential\n                - --cluster-id\n                - c5e***\n                - --api-version\n                - v1beta1\n                - --expiration\n                - 3h\n                - --log-level\n                - error\n            apiVersion: client.authentication.k8s.io/v1beta1\n            provideClusterInfo: false\n            interactiveMode: Never\npreferences: {}\n\n$ ack-ram-tool credential-plugin get-kubeconfig --cluster-id c5e*** > kubeconfig\n$ proxy_ack kubectl --kubeconfig kubeconfig get ns\nNAME                         STATUS   AGE\ndefault                      Active   6d3h\nkube-node-lease              Active   6d3h\nkube-public                  Active   6d3h\nkube-system                  Active   6d3h\n\n### --mode ram-authenticator-token\n\n$ ack-ram-tool credential-plugin get-kubeconfig --mode ram-authenticator-token --cluster-id c5e***\n\nkind: Config\napiVersion: v1\nclusters:\n    - name: kubernetes\n      cluster:\n        server: https://106.*.*.*:6443\n        certificate-authority-data: LS0t***\ncontexts:\n    - name: 272***-c5e***\n      context:\n        cluster: kubernetes\n        user: "272***"\ncurrent-context: 272***-c5e***\nusers:\n    - name: "272***"\n      user:\n        exec:\n            command: ack-ram-tool\n            args:\n                - credential-plugin\n                - get-token\n                - --cluster-id\n                - c5e***\n                - --api-version\n                - v1beta1\n                - --log-level\n                - error\n            apiVersion: client.authentication.k8s.io/v1beta1\n            provideClusterInfo: false\n            interactiveMode: Never\npreferences: {}\n\n')),(0,a.kt)("h2",{id:"\u547d\u4ee4\u884c\u53c2\u6570"},"\u547d\u4ee4\u884c\u53c2\u6570"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'Usage:\n  ack-ram-tool credential-plugin get-kubeconfig [flags]\n\nFlags:\n      --api-version string            v1 or v1beta1 (default "v1beta1")\n  -c, --cluster-id string             The cluster id to use\n      --credential-cache-dir string   Directory to cache certificate (default "~/.kube/cache/ack-ram-tool/credential-plugin")\n      --expiration duration           The certificate expiration (default 3h0m0s)\n  -h, --help                          help for get-kubeconfig\n  -m, --mode string                   credential mode: certificate or ram-authenticator-token (default "certificate")\n      --private-address               Use private ip as api-server address\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n')),(0,a.kt)("p",null,"\u53c2\u6570\u8bf4\u660e\uff1a"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"\u53c2\u6570\u540d\u79f0"),(0,a.kt)("th",{parentName:"tr",align:null},"\u9ed8\u8ba4\u503c"),(0,a.kt)("th",{parentName:"tr",align:null},"\u5fc5\u9700\u53c2\u6570"),(0,a.kt)("th",{parentName:"tr",align:null},"\u8bf4\u660e"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-c, --cluster-id"),(0,a.kt)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.kt)("td",{parentName:"tr",align:null},"\u662f"),(0,a.kt)("td",{parentName:"tr",align:null},"\u96c6\u7fa4 ID")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-m, --mode"),(0,a.kt)("td",{parentName:"tr",align:null},"certificate"),(0,a.kt)("td",{parentName:"tr",align:null},"\u5426"),(0,a.kt)("td",{parentName:"tr",align:null},"kubeconfig \u4e2d\u7684\u8ba4\u8bc1\u65b9\u6cd5\uff1a ",(0,a.kt)("inlineCode",{parentName:"td"},"certificate")," \u8868\u793a\u8bc1\u4e66\u8ba4\u8bc1\uff0c",(0,a.kt)("inlineCode",{parentName:"td"},"ram-authenticator-token")," \u8868\u793a\u57fa\u4e8e ack-ram-authenticator \u7684 token \u8ba4\u8bc1")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--expiration"),(0,a.kt)("td",{parentName:"tr",align:null},"3h"),(0,a.kt)("td",{parentName:"tr",align:null},"\u5426"),(0,a.kt)("td",{parentName:"tr",align:null},"--mode \u88ab\u8bbe\u7f6e\u4e3a ",(0,a.kt)("inlineCode",{parentName:"td"},"certificate")," \u65f6\uff0c\u901a\u8fc7\u8fd9\u4e2a\u53c2\u6570\u8bbe\u7f6e\u8bc1\u4e66\u8fc7\u671f\u65f6\u95f4\u3002\u4e3a 0 \u65f6\u8868\u793a\u4e0d\u4f7f\u7528\u4e34\u65f6\u8bc1\u4e66\u800c\u662f\u4f7f\u7528\u6709\u6548\u671f\u66f4\u957f\u7684\u8bc1\u4e66\uff08\u8fc7\u671f\u65f6\u95f4\u7531\u670d\u52a1\u7aef\u81ea\u52a8\u786e\u5b9a\uff09")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--private-address"),(0,a.kt)("td",{parentName:"tr",align:null},"false"),(0,a.kt)("td",{parentName:"tr",align:null},"\u5426"),(0,a.kt)("td",{parentName:"tr",align:null},"\u662f\u5426\u4f7f\u7528\u5185\u7f51 api server \u5730\u5740")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--api-version"),(0,a.kt)("td",{parentName:"tr",align:null},"v1beta1"),(0,a.kt)("td",{parentName:"tr",align:null},"\u5426"),(0,a.kt)("td",{parentName:"tr",align:null},"\u6307\u5b9a\u8fd4\u56de\u7684\u6570\u636e\u4e2d\u4f7f\u7528\u54ea\u4e2a\u7248\u672c\u7684 apiVersion\u3002v1beta1 \u8868\u793a ",(0,a.kt)("inlineCode",{parentName:"td"},"client.authentication.k8s.io/v1beta1"),"\uff0cv1 \u8868\u793a ",(0,a.kt)("inlineCode",{parentName:"td"},"client.authentication.k8s.io/v1beta1"))),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--credential-cache-dir"),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"~/.kube/cache/ack-ram-tool/credential-plugin")),(0,a.kt)("td",{parentName:"tr",align:null},"\u5426"),(0,a.kt)("td",{parentName:"tr",align:null},"\u7528\u4e8e\u7f13\u5b58\u8bc1\u4e66\u7684\u76ee\u5f55\uff0c\u53ea\u5728 ",(0,a.kt)("inlineCode",{parentName:"td"},"--mode")," \u88ab\u8bbe\u7f6e\u4e3a ",(0,a.kt)("inlineCode",{parentName:"td"},"certificate")," \u65f6\u6709\u6548")))))}s.isMDXComponent=!0}}]);