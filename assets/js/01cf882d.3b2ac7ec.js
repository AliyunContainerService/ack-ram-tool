"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[9062],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>m});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),s=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},u=function(e){var t=s(e.components);return r.createElement(c.Provider,{value:t},e.children)},d="mdxType",p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},g=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,c=e.parentName,u=o(e,["components","mdxType","originalType","parentName"]),d=s(n),g=a,m=d["".concat(c,".").concat(g)]||d[g]||p[g]||i;return n?r.createElement(m,l(l({ref:t},u),{},{components:n})):r.createElement(m,l({ref:t},u))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,l=new Array(i);l[0]=g;var o={};for(var c in t)hasOwnProperty.call(t,c)&&(o[c]=t[c]);o.originalType=e,o[d]="string"==typeof e?e:a,l[1]=o;for(var s=2;s<i;s++)l[s]=n[s];return r.createElement.apply(null,l)}return r.createElement.apply(null,n)}g.displayName="MDXCreateElement"},8289:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>p,frontMatter:()=>i,metadata:()=>o,toc:()=>s});var r=n(7462),a=(n(7294),n(3905));const i={slug:"get-credential",sidebar_position:2},l="get-credential",o={unversionedId:"credential-plugin/get-credential",id:"version-v0.15.0/credential-plugin/get-credential",title:"get-credential",description:"Get the ExecCredential certificate data used to access the API server.",source:"@site/versioned_docs/version-v0.15.0/credential-plugin/get-credential.md",sourceDirName:"credential-plugin",slug:"/credential-plugin/get-credential",permalink:"/ack-ram-tool/v0.15.0/credential-plugin/get-credential",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.15.0/credential-plugin/get-credential.md",tags:[],version:"v0.15.0",sidebarPosition:2,frontMatter:{slug:"get-credential",sidebar_position:2},sidebar:"tutorialSidebar",previous:{title:"get-kubeconfig\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.15.0/zh-CN/credential-plugin/kubeconfig"},next:{title:"get-credential\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.15.0/zh-CN/credential-plugin/get-credential"}},c={},s=[{value:"Usage",id:"usage",level:2},{value:"Flags",id:"flags",level:2}],u={toc:s},d="wrapper";function p(e){let{components:t,...n}=e;return(0,a.kt)(d,(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"get-credential"},"get-credential"),(0,a.kt)("p",null,"Get the ExecCredential certificate data used to access the API server."),(0,a.kt)("p",null,"It has the following features:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Automatically obtains a new certificate before the certificate expires."),(0,a.kt)("li",{parentName:"ul"},"Supports using temporary certificate.")),(0,a.kt)("h2",{id:"usage"},"Usage"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'$ ack-ram-tool credential-plugin get-credential --cluster-id <clusterId>\n\n{\n "kind": "ExecCredential",\n "apiVersion": "client.authentication.k8s.io/v1beta1",\n "spec": {\n  "interactive": false\n },\n "status": {\n  "expirationTimestamp": "2023-04-20T09:29:06Z",\n  "clientCertificateData": "-----BEGIN CERTIFICATE-----\\nMIID***\\n-----END CERTIFICATE-----\\n",\n  "clientKeyData": "-----BEGIN RSA PRIVATE KEY-----\\nMIIE***\\n-----END RSA PRIVATE KEY-----\\n"\n }\n}\n')),(0,a.kt)("h2",{id:"flags"},"Flags"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'Usage:\n  ack-ram-tool credential-plugin get-credential [flags]\n\nFlags:\n      --api-version string            v1 or v1beta1 (default "v1beta1")\n  -c, --cluster-id string             The cluster id to use\n      --credential-cache-dir string   Directory to cache credential (default "~/.kube/cache/ack-ram-tool/credential-plugin")\n      --expiration duration           The credential expiration (default 3h0m0s)\n  -h, --help                          help for get-credential\n      --role-arn string               Assume an RAM Role ARN when send request or sign token\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n')),(0,a.kt)("p",null,"Descriptions\uff1a"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"Flag"),(0,a.kt)("th",{parentName:"tr",align:null},"Default"),(0,a.kt)("th",{parentName:"tr",align:null},"Required"),(0,a.kt)("th",{parentName:"tr",align:null},"Description"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-c, --cluster-id"),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null},"Yes"),(0,a.kt)("td",{parentName:"tr",align:null},"Cluster ID")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--api-version"),(0,a.kt)("td",{parentName:"tr",align:null},"v1beta1"),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null},"Specify which version of apiVersion to use in the returned data. ",(0,a.kt)("inlineCode",{parentName:"td"},"v1beta1")," represents ",(0,a.kt)("inlineCode",{parentName:"td"},"client.authentication.k8s.io/v1beta1"),", and ",(0,a.kt)("inlineCode",{parentName:"td"},"v1")," represents ",(0,a.kt)("inlineCode",{parentName:"td"},"client.authentication.k8s.io/v1"),".")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--expiration"),(0,a.kt)("td",{parentName:"tr",align:null},"3h0m0s"),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null},"Specify the certificate expiration time. When it is 0, it means not to use a temporary certificate but to use a longer valid certificate (the expiration time is automatically determined by the server).")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--credential-cache-dir"),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"~/.kube/cache/ack-ram-tool/credential-plugin")),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null},"Directory used to cache the certificate")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--role-arn"),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null},"Assume an RAM Role ARN when send request or sign token")))))}p.isMDXComponent=!0}}]);