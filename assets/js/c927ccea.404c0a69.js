"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[4485],{3905:(e,t,n)=>{n.d(t,{Zo:()=>s,kt:()=>m});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var c=a.createContext({}),u=function(e){var t=a.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},s=function(e){var t=u(e.components);return a.createElement(c.Provider,{value:t},e.children)},d="mdxType",p={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},g=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,c=e.parentName,s=o(e,["components","mdxType","originalType","parentName"]),d=u(n),g=r,m=d["".concat(c,".").concat(g)]||d[g]||p[g]||i;return n?a.createElement(m,l(l({ref:t},s),{},{components:n})):a.createElement(m,l({ref:t},s))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,l=new Array(i);l[0]=g;var o={};for(var c in t)hasOwnProperty.call(t,c)&&(o[c]=t[c]);o.originalType=e,o[d]="string"==typeof e?e:r,l[1]=o;for(var u=2;u<i;u++)l[u]=n[u];return a.createElement.apply(null,l)}return a.createElement.apply(null,n)}g.displayName="MDXCreateElement"},4055:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>p,frontMatter:()=>i,metadata:()=>o,toc:()=>u});var a=n(7462),r=(n(7294),n(3905));const i={slug:"get-kubeconfig",sidebar_position:1},l="get-kubeconfig",o={unversionedId:"credential-plugin/get-kubeconfig",id:"credential-plugin/get-kubeconfig",title:"get-kubeconfig",description:"Obtain a kubeconfig file that uses ack-ram-tool as the credential plugin.",source:"@site/docs/credential-plugin/get-kubeconfig.md",sourceDirName:"credential-plugin",slug:"/credential-plugin/get-kubeconfig",permalink:"/ack-ram-tool/next/credential-plugin/get-kubeconfig",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/docs/credential-plugin/get-kubeconfig.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{slug:"get-kubeconfig",sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"credential-plugin",permalink:"/ack-ram-tool/next/category/credential-plugin"},next:{title:"get-kubeconfig\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/next/zh-CN/credential-plugin/kubeconfig"}},c={},u=[{value:"Usage",id:"usage",level:2},{value:"--mode ram-authenticator-token",id:"--mode-ram-authenticator-token",level:3},{value:"Flags",id:"flags",level:2}],s={toc:u},d="wrapper";function p(e){let{components:t,...n}=e;return(0,r.kt)(d,(0,a.Z)({},s,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"get-kubeconfig"},"get-kubeconfig"),(0,r.kt)("p",null,"Obtain a kubeconfig file that uses ack-ram-tool as the ",(0,r.kt)("a",{parentName:"p",href:"https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins"},"credential plugin"),"."),(0,r.kt)("p",null,"It has the following features:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Automatically obtains a new certificate before the certificate expires."),(0,r.kt)("li",{parentName:"ul"},"Supports using temporary certificate."),(0,r.kt)("li",{parentName:"ul"},"Integrate ack-ram-authenticator.")),(0,r.kt)("h2",{id:"usage"},"Usage"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},'$ ack-ram-tool credential-plugin get-kubeconfig --cluster-id c5e***\n\nkind: Config\napiVersion: v1\nclusters:\n    - name: kubernetes\n      cluster:\n        server: https://106.*.*.*:6443\n        certificate-authority-data: LS0tL***\ncontexts:\n    - name: 272***-c5e***\n      context:\n        cluster: kubernetes\n        user: "272***"\ncurrent-context: 272***-c5e***\nusers:\n    - name: "272***"\n      user:\n        exec:\n            command: ack-ram-tool\n            args:\n                - credential-plugin\n                - get-credential\n                - --cluster-id\n                - c5e***\n                - --api-version\n                - v1beta1\n                - --expiration\n                - 3h\n                - --log-level\n                - error\n            apiVersion: client.authentication.k8s.io/v1beta1\n            provideClusterInfo: false\n            interactiveMode: Never\npreferences: {}\n\n$ ack-ram-tool credential-plugin get-kubeconfig --cluster-id c5e*** > kubeconfig\n$ kubectl --kubeconfig kubeconfig get ns\nNAME                         STATUS   AGE\ndefault                      Active   6d3h\nkube-node-lease              Active   6d3h\nkube-public                  Active   6d3h\nkube-system                  Active   6d3h\n')),(0,r.kt)("h3",{id:"--mode-ram-authenticator-token"},"--mode ram-authenticator-token"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},'$ ack-ram-tool credential-plugin get-kubeconfig --mode ram-authenticator-token --cluster-id c5e***\n\nkind: Config\napiVersion: v1\nclusters:\n    - name: kubernetes\n      cluster:\n        server: https://106.*.*.*:6443\n        certificate-authority-data: LS0t***\ncontexts:\n    - name: 272***-c5e***\n      context:\n        cluster: kubernetes\n        user: "272***"\ncurrent-context: 272***-c5e***\nusers:\n    - name: "272***"\n      user:\n        exec:\n            command: ack-ram-tool\n            args:\n                - credential-plugin\n                - get-token\n                - --cluster-id\n                - c5e***\n                - --api-version\n                - v1beta1\n                - --log-level\n                - error\n            apiVersion: client.authentication.k8s.io/v1beta1\n            provideClusterInfo: false\n            interactiveMode: Never\npreferences: {}\n\n')),(0,r.kt)("h2",{id:"flags"},"Flags"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},'Usage:\n  ack-ram-tool credential-plugin get-kubeconfig [flags]\n\nFlags:\n      --api-version string            v1 or v1beta1 (default "v1beta1")\n  -c, --cluster-id string             The cluster id to use\n      --credential-cache-dir string   Directory to cache certificate (default "~/.kube/cache/ack-ram-tool/credential-plugin")\n      --expiration duration           The certificate expiration (default 3h0m0s)\n  -h, --help                          help for get-kubeconfig\n  -m, --mode string                   credential mode: certificate or ram-authenticator-token (default "certificate")\n      --private-address               Use private ip as api-server address\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n')),(0,r.kt)("p",null,"Descriptions\uff1a"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Flag"),(0,r.kt)("th",{parentName:"tr",align:null},"Default"),(0,r.kt)("th",{parentName:"tr",align:null},"Required"),(0,r.kt)("th",{parentName:"tr",align:null},"Description"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"-c, --cluster-id"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"Yes"),(0,r.kt)("td",{parentName:"tr",align:null},"Cluster ID")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"-m, --mode"),(0,r.kt)("td",{parentName:"tr",align:null},"certificate"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"Authentication methods in kubeconfig: ",(0,r.kt)("inlineCode",{parentName:"td"},"certificate")," indicates certificate authentication, and ",(0,r.kt)("inlineCode",{parentName:"td"},"ram-authenticator-token")," indicates token authentication based on ack-ram-authenticator")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"--expiration"),(0,r.kt)("td",{parentName:"tr",align:null},"3h"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"When --mode is set to ",(0,r.kt)("inlineCode",{parentName:"td"},"certificate"),", set the certificate expiration time through this parameter. When it is 0, it means not to use a temporary certificate but to use a longer valid certificate (the expiration time is automatically determined by the server).")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"--private-address"),(0,r.kt)("td",{parentName:"tr",align:null},"false"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"Whether to use the intranet API server address?")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"--api-version"),(0,r.kt)("td",{parentName:"tr",align:null},"v1beta1"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"Specify which version of apiVersion to use in the returned data. ",(0,r.kt)("inlineCode",{parentName:"td"},"v1beta1")," represents ",(0,r.kt)("inlineCode",{parentName:"td"},"client.authentication.k8s.io/v1beta1"),", and ",(0,r.kt)("inlineCode",{parentName:"td"},"v1")," represents ",(0,r.kt)("inlineCode",{parentName:"td"},"client.authentication.k8s.io/v1"),".")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"--credential-cache-dir"),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"~/.kube/cache/ack-ram-tool/credential-plugin")),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"The directory used to cache the certificate is only valid when ",(0,r.kt)("inlineCode",{parentName:"td"},"--mode")," is set to ",(0,r.kt)("inlineCode",{parentName:"td"},"certificate"))))))}p.isMDXComponent=!0}}]);