"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[3628],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>k});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=a.createContext({}),d=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},p=function(e){var t=d(e.components);return a.createElement(s.Provider,{value:t},e.children)},m="mdxType",c={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},u=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,s=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),m=d(n),u=r,k=m["".concat(s,".").concat(u)]||m[u]||c[u]||i;return n?a.createElement(k,l(l({ref:t},p),{},{components:n})):a.createElement(k,l({ref:t},p))}));function k(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,l=new Array(i);l[0]=u;var o={};for(var s in t)hasOwnProperty.call(t,s)&&(o[s]=t[s]);o.originalType=e,o[m]="string"==typeof e?e:r,l[1]=o;for(var d=2;d<i;d++)l[d]=n[d];return a.createElement.apply(null,l)}return a.createElement.apply(null,n)}u.displayName="MDXCreateElement"},4284:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>l,default:()=>c,frontMatter:()=>i,metadata:()=>o,toc:()=>d});var a=n(7462),r=(n(7294),n(3905));const i={slug:"/",sidebar_position:1},l="Getting started",o={unversionedId:"getting-started",id:"version-v0.15.0/getting-started",title:"Getting started",description:"Installation",source:"@site/versioned_docs/version-v0.15.0/getting-started.md",sourceDirName:".",slug:"/",permalink:"/ack-ram-tool/v0.15.0/",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.15.0/getting-started.md",tags:[],version:"v0.15.0",sidebarPosition:1,frontMatter:{slug:"/",sidebar_position:1},sidebar:"tutorialSidebar",next:{title:"\u65b0\u624b\u6307\u5357",permalink:"/ack-ram-tool/v0.15.0/zh-CN/getting-started"}},s={},d=[{value:"Installation",id:"installation",level:2},{value:"Configuration",id:"configuration",level:2},{value:"Credentials",id:"credentials",level:3},{value:"Permissions",id:"permissions",level:2}],p={toc:d},m="wrapper";function c(e){let{components:t,...n}=e;return(0,r.kt)(m,(0,a.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"getting-started"},"Getting started"),(0,r.kt)("h2",{id:"installation"},"Installation"),(0,r.kt)("p",null,"lease go to the ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/AliyunContainerService/ack-ram-tool/releases"},"Releases")," page\nto download the latest version of the ack-ram-tool."),(0,r.kt)("h2",{id:"configuration"},"Configuration"),(0,r.kt)("h3",{id:"credentials"},"Credentials"),(0,r.kt)("p",null,"ack-ram-tool will search for credential information in the system in the following order\uff1a"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Automatically use credential information that exists in the environment variables\uff08\nNote: This tool also supports the credential-related environment variables supported by ",(0,r.kt)("a",{parentName:"li",href:"https://github.com/aliyun/aliyun-cli#support-for-environment-variables"},"aliyun cli")," \uff09:")),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"environment variables"),(0,r.kt)("th",{parentName:"tr",align:null},"description"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABA_CLOUD_ACCESS_KEY_ID"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALICLOUD_ACCESS_KEY"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABACLOUD_ACCESS_KEY_ID"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALICLOUD_ACCESS_KEY_ID"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABACLOUD_ACCESS_KEY_ID"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ACCESS_KEY_ID")),(0,r.kt)("td",{parentName:"tr",align:null},"access key id")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABA_CLOUD_ACCESS_KEY_SECRET"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALICLOUD_SECRET_KEY"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABACLOUD_ACCESS_KEY_SECRET"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALICLOUD_ACCESS_KEY_SECRET"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ACCESS_KEY_SECRET")),(0,r.kt)("td",{parentName:"tr",align:null},"access key secret")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABA_CLOUD_SECURITY_TOKEN"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALICLOUD_ACCESS_KEY_STS_TOKEN"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABACLOUD_SECURITY_TOKEN"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALICLOUD_SECURITY_TOKEN"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABACLOUD_SECURITY_TOKEN"),"\u3001",(0,r.kt)("inlineCode",{parentName:"td"},"SECURITY_TOKEN")),(0,r.kt)("td",{parentName:"tr",align:null},"sts token")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABA_CLOUD_CREDENTIALS_URI")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("a",{parentName:"td",href:"https://github.com/aliyun/aliyun-cli#use-credentials-uri"},"credentials URI"))),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABA_CLOUD_ROLE_ARN")),(0,r.kt)("td",{parentName:"tr",align:null},"RAM Role ARN")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABA_CLOUD_OIDC_PROVIDER_ARN")),(0,r.kt)("td",{parentName:"tr",align:null},"OIDC Provider ARN")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("inlineCode",{parentName:"td"},"ALIBABA_CLOUD_OIDC_TOKEN_FILE")),(0,r.kt)("td",{parentName:"tr",align:null},"OIDC Token File")))),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},"When credential information does not exist in the environment variables, if there is an aliyun cli configuration file\n",(0,r.kt)("inlineCode",{parentName:"p"},"~/.aliyun/config.json")," (For details on the aliyun cli configuration file,\nplease refer to the ",(0,r.kt)("a",{parentName:"p",href:"https://www.alibabacloud.com/help/doc-detail/110341.htm"},"official documentation")," ) ,\nthe program will automatically use that configuration file.")),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},"When the aliyun cli configuration file does not exist, the program will attempt to use the credential information\nconfigured in the ",(0,r.kt)("inlineCode",{parentName:"p"},"~/.alibabacloud/credentials")," file (which can be specified by the ",(0,r.kt)("inlineCode",{parentName:"p"},"--profile-file")," flags):"))),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"$ cat ~/.alibabacloud/credentials\n\n[default]\ntype = access_key\naccess_key_id = foo\naccess_key_secret = bar\n")),(0,r.kt)("h2",{id:"permissions"},"Permissions"),(0,r.kt)("p",null,"In order to use ack-ram-tool normally, you need to grant the necessary RAM permissions and RBAC permissions\nfor the Alibaba Cloud RAM user or RAM role that uses this tool.\nFor the minimum permission information required for each subcommand, please refer to ",(0,r.kt)("a",{parentName:"p",href:"permissions"},"Permissions"),"."))}c.isMDXComponent=!0}}]);