(window.webpackJsonp=window.webpackJsonp||[]).push([[29],{150:function(e,t,a){"use strict";a.d(t,"a",(function(){return d})),a.d(t,"b",(function(){return u}));var i=a(0),c=a.n(i);function n(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function l(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,i)}return a}function b(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?l(Object(a),!0).forEach((function(t){n(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):l(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function s(e,t){if(null==e)return{};var a,i,c=function(e,t){if(null==e)return{};var a,i,c={},n=Object.keys(e);for(i=0;i<n.length;i++)a=n[i],t.indexOf(a)>=0||(c[a]=e[a]);return c}(e,t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);for(i=0;i<n.length;i++)a=n[i],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(c[a]=e[a])}return c}var r=c.a.createContext({}),o=function(e){var t=c.a.useContext(r),a=t;return e&&(a="function"==typeof e?e(t):b(b({},t),e)),a},d=function(e){var t=o(e.components);return c.a.createElement(r.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return c.a.createElement(c.a.Fragment,{},t)}},m=c.a.forwardRef((function(e,t){var a=e.components,i=e.mdxType,n=e.originalType,l=e.parentName,r=s(e,["components","mdxType","originalType","parentName"]),d=o(a),m=i,u=d["".concat(l,".").concat(m)]||d[m]||p[m]||n;return a?c.a.createElement(u,b(b({ref:t},r),{},{components:a})):c.a.createElement(u,b({ref:t},r))}));function u(e,t){var a=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var n=a.length,l=new Array(n);l[0]=m;var b={};for(var s in t)hasOwnProperty.call(t,s)&&(b[s]=t[s]);b.originalType=e,b.mdxType="string"==typeof e?e:i,l[1]=b;for(var r=2;r<n;r++)l[r]=a[r];return c.a.createElement.apply(null,l)}return c.a.createElement.apply(null,a)}m.displayName="MDXCreateElement"},84:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return l})),a.d(t,"metadata",(function(){return b})),a.d(t,"rightToc",(function(){return s})),a.d(t,"default",(function(){return o}));var i=a(2),c=a(6),n=(a(0),a(150)),l={id:"production-security-checklist",sidebar_label:"Backend Security Checklist",title:"Backend Production Ready Security Checklist"},b={unversionedId:"technology/checklists/backend/backend-security-checklist/production-security-checklist",id:"technology/checklists/backend/backend-security-checklist/production-security-checklist",isDocsHomePage:!1,title:"Backend Production Ready Security Checklist",description:"API Security Checklist",source:"@site/docs/technology/checklists/backend/backend-security-checklist/README.md",permalink:"/docs/technology/checklists/backend/backend-security-checklist/production-security-checklist",editUrl:"https://github.com/BlackspaceInc/BlackspacePlatform/edit/master/website/docs/technology/checklists/backend/backend-security-checklist/README.md",sidebar_label:"Backend Security Checklist",sidebar:"docs",previous:{title:"Backend Production Ready Checklist",permalink:"/docs/technology/checklists/backend/backend-production-ready/production-ready-checklist"},next:{title:"Frontend Production Ready Performance Checklist",permalink:"/docs/technology/checklists/frontend/frontend-performance-checklist/frontend-performance-checklist"}},s=[{value:"Authentication",id:"authentication",children:[{value:"JWT (JSON Web Token)",id:"jwt-json-web-token",children:[]},{value:"OAuth",id:"oauth",children:[]}]},{value:"Access",id:"access",children:[]},{value:"Input",id:"input",children:[]},{value:"Processing",id:"processing",children:[]},{value:"Output",id:"output",children:[]},{value:"CI &amp; CD",id:"ci--cd",children:[]},{value:"See also:",id:"see-also",children:[]}],r={rightToc:s};function o(e){var t=e.components,a=Object(c.a)(e,["components"]);return Object(n.b)("wrapper",Object(i.a)({},r,a,{components:t,mdxType:"MDXLayout"}),Object(n.b)("h1",{id:"api-security-checklist"},"API Security Checklist"),Object(n.b)("p",null,"Checklist of the most important security countermeasures when designing, testing, and releasing your API."),Object(n.b)("hr",null),Object(n.b)("h2",{id:"authentication"},"Authentication"),Object(n.b)("ul",{className:"contains-task-list"},Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Don't use ",Object(n.b)("inlineCode",{parentName:"li"},"Basic Auth"),". Use standard authentication instead (e.g. ",Object(n.b)("a",Object(i.a)({parentName:"li"},{href:"https://jwt.io/"}),"JWT"),", ",Object(n.b)("a",Object(i.a)({parentName:"li"},{href:"https://oauth.net/"}),"OAuth"),")."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Don't reinvent the wheel in ",Object(n.b)("inlineCode",{parentName:"li"},"Authentication"),", ",Object(n.b)("inlineCode",{parentName:"li"},"token generation"),", ",Object(n.b)("inlineCode",{parentName:"li"},"password storage"),". Use the standards."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use ",Object(n.b)("inlineCode",{parentName:"li"},"Max Retry")," and jail features in Login."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use encryption on all sensitive data.")),Object(n.b)("h3",{id:"jwt-json-web-token"},"JWT (JSON Web Token)"),Object(n.b)("ul",{className:"contains-task-list"},Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use a random complicated key (",Object(n.b)("inlineCode",{parentName:"li"},"JWT Secret"),") to make brute forcing the token very hard."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Don't extract the algorithm from the header. Force the algorithm in the backend (",Object(n.b)("inlineCode",{parentName:"li"},"HS256")," or ",Object(n.b)("inlineCode",{parentName:"li"},"RS256"),")."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Make token expiration (",Object(n.b)("inlineCode",{parentName:"li"},"TTL"),", ",Object(n.b)("inlineCode",{parentName:"li"},"RTTL"),") as short as possible."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Don't store sensitive data in the JWT payload, it can be decoded ",Object(n.b)("a",Object(i.a)({parentName:"li"},{href:"https://jwt.io/#debugger-io"}),"easily"),".")),Object(n.b)("h3",{id:"oauth"},"OAuth"),Object(n.b)("ul",{className:"contains-task-list"},Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Always validate ",Object(n.b)("inlineCode",{parentName:"li"},"redirect_uri")," server-side to allow only whitelisted URLs."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Always try to exchange for code and not tokens (don't allow ",Object(n.b)("inlineCode",{parentName:"li"},"response_type=token"),")."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use ",Object(n.b)("inlineCode",{parentName:"li"},"state")," parameter with a random hash to prevent CSRF on the OAuth authentication process."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Define the default scope, and validate scope parameters for each application.")),Object(n.b)("h2",{id:"access"},"Access"),Object(n.b)("ul",{className:"contains-task-list"},Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Limit requests (Throttling) to avoid DDoS / brute-force attacks."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use HTTPS on server side to avoid MITM (Man in the Middle Attack)."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use ",Object(n.b)("inlineCode",{parentName:"li"},"HSTS")," header with SSL to avoid SSL Strip attack.")),Object(n.b)("h2",{id:"input"},"Input"),Object(n.b)("ul",{className:"contains-task-list"},Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use the proper HTTP method according to the operation: ",Object(n.b)("inlineCode",{parentName:"li"},"GET (read)"),", ",Object(n.b)("inlineCode",{parentName:"li"},"POST (create)"),", ",Object(n.b)("inlineCode",{parentName:"li"},"PUT/PATCH (replace/update)"),", and ",Object(n.b)("inlineCode",{parentName:"li"},"DELETE (to delete a record)"),", and respond with ",Object(n.b)("inlineCode",{parentName:"li"},"405 Method Not Allowed")," if the requested method isn't appropriate for the requested resource."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Validate ",Object(n.b)("inlineCode",{parentName:"li"},"content-type")," on request Accept header (Content Negotiation) to allow only your supported format (e.g. ",Object(n.b)("inlineCode",{parentName:"li"},"application/xml"),", ",Object(n.b)("inlineCode",{parentName:"li"},"application/json"),", etc.) and respond with ",Object(n.b)("inlineCode",{parentName:"li"},"406 Not Acceptable")," response if not matched."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Validate ",Object(n.b)("inlineCode",{parentName:"li"},"content-type")," of posted data as you accept (e.g. ",Object(n.b)("inlineCode",{parentName:"li"},"application/x-www-form-urlencoded"),", ",Object(n.b)("inlineCode",{parentName:"li"},"multipart/form-data"),", ",Object(n.b)("inlineCode",{parentName:"li"},"application/json"),", etc.)."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Validate user input to avoid common vulnerabilities (e.g. ",Object(n.b)("inlineCode",{parentName:"li"},"XSS"),", ",Object(n.b)("inlineCode",{parentName:"li"},"SQL-Injection"),", ",Object(n.b)("inlineCode",{parentName:"li"},"Remote Code Execution"),", etc.)."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Don't use any sensitive data (",Object(n.b)("inlineCode",{parentName:"li"},"credentials"),", ",Object(n.b)("inlineCode",{parentName:"li"},"Passwords"),", ",Object(n.b)("inlineCode",{parentName:"li"},"security tokens"),", or ",Object(n.b)("inlineCode",{parentName:"li"},"API keys"),") in the URL, but use standard Authorization header."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use an API Gateway service to enable caching, Rate Limit policies (e.g. ",Object(n.b)("inlineCode",{parentName:"li"},"Quota"),", ",Object(n.b)("inlineCode",{parentName:"li"},"Spike Arrest"),", or ",Object(n.b)("inlineCode",{parentName:"li"},"Concurrent Rate Limit"),") and deploy APIs resources dynamically.")),Object(n.b)("h2",{id:"processing"},"Processing"),Object(n.b)("ul",{className:"contains-task-list"},Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Check if all the endpoints are protected behind authentication to avoid broken authentication process."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","User own resource ID should be avoided. Use ",Object(n.b)("inlineCode",{parentName:"li"},"/me/orders")," instead of ",Object(n.b)("inlineCode",{parentName:"li"},"/user/654321/orders"),"."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Don't auto-increment IDs. Use ",Object(n.b)("inlineCode",{parentName:"li"},"UUID")," instead."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","If you are parsing XML files, make sure entity parsing is not enabled to avoid ",Object(n.b)("inlineCode",{parentName:"li"},"XXE")," (XML external entity attack)."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","If you are parsing XML files, make sure entity expansion is not enabled to avoid ",Object(n.b)("inlineCode",{parentName:"li"},"Billion Laughs/XML bomb")," via exponential entity expansion attack."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use a CDN for file uploads."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","If you are dealing with huge amount of data, use Workers and Queues to process as much as possible in background and return response fast to avoid HTTP Blocking."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Do not forget to turn the DEBUG mode OFF.")),Object(n.b)("h2",{id:"output"},"Output"),Object(n.b)("ul",{className:"contains-task-list"},Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Send ",Object(n.b)("inlineCode",{parentName:"li"},"X-Content-Type-Options: nosniff")," header."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Send ",Object(n.b)("inlineCode",{parentName:"li"},"X-Frame-Options: deny")," header."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Send ",Object(n.b)("inlineCode",{parentName:"li"},"Content-Security-Policy: default-src 'none'")," header."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Remove fingerprinting headers - ",Object(n.b)("inlineCode",{parentName:"li"},"X-Powered-By"),", ",Object(n.b)("inlineCode",{parentName:"li"},"Server"),", ",Object(n.b)("inlineCode",{parentName:"li"},"X-AspNet-Version"),", etc."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Force ",Object(n.b)("inlineCode",{parentName:"li"},"content-type")," for your response. If you return ",Object(n.b)("inlineCode",{parentName:"li"},"application/json"),", then your ",Object(n.b)("inlineCode",{parentName:"li"},"content-type")," response is ",Object(n.b)("inlineCode",{parentName:"li"},"application/json"),"."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Don't return sensitive data like ",Object(n.b)("inlineCode",{parentName:"li"},"credentials"),", ",Object(n.b)("inlineCode",{parentName:"li"},"Passwords"),", or ",Object(n.b)("inlineCode",{parentName:"li"},"security tokens"),"."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Return the proper status code according to the operation completed. (e.g. ",Object(n.b)("inlineCode",{parentName:"li"},"200 OK"),", ",Object(n.b)("inlineCode",{parentName:"li"},"400 Bad Request"),", ",Object(n.b)("inlineCode",{parentName:"li"},"401 Unauthorized"),", ",Object(n.b)("inlineCode",{parentName:"li"},"405 Method Not Allowed"),", etc.).")),Object(n.b)("h2",{id:"ci--cd"},"CI & CD"),Object(n.b)("ul",{className:"contains-task-list"},Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Audit your design and implementation with unit/integration tests coverage."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Use a code review process and disregard self-approval."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Ensure that all components of your services are statically scanned by AV software before pushing to production, including vendor libraries and other dependencies."),Object(n.b)("li",Object(i.a)({parentName:"ul"},{className:"task-list-item"}),Object(n.b)("input",Object(i.a)({parentName:"li"},{type:"checkbox",checked:!1,disabled:!0}))," ","Design a rollback solution for deployments.")),Object(n.b)("hr",null),Object(n.b)("h2",{id:"see-also"},"See also:"),Object(n.b)("ul",null,Object(n.b)("li",{parentName:"ul"},Object(n.b)("a",Object(i.a)({parentName:"li"},{href:"https://github.com/yosriady/api-development-tools"}),"yosriady/api-development-tools")," - A collection of useful resources for building RESTful HTTP+JSON APIs.")),Object(n.b)("hr",null))}o.isMDXComponent=!0}}]);