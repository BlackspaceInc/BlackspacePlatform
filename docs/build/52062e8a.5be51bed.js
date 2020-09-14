(window.webpackJsonp=window.webpackJsonp||[]).push([[31],{150:function(e,t,n){"use strict";n.d(t,"a",(function(){return p})),n.d(t,"b",(function(){return d}));var r=n(0),i=n.n(r);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function c(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function u(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var s=i.a.createContext({}),l=function(e){var t=i.a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):c(c({},t),e)),n},p=function(e){var t=l(e.components);return i.a.createElement(s.Provider,{value:t},e.children)},b={inlineCode:"code",wrapper:function(e){var t=e.children;return i.a.createElement(i.a.Fragment,{},t)}},f=i.a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,a=e.originalType,o=e.parentName,s=u(e,["components","mdxType","originalType","parentName"]),p=l(n),f=r,d=p["".concat(o,".").concat(f)]||p[f]||b[f]||a;return n?i.a.createElement(d,c(c({ref:t},s),{},{components:n})):i.a.createElement(d,c({ref:t},s))}));function d(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var a=n.length,o=new Array(a);o[0]=f;var c={};for(var u in t)hasOwnProperty.call(t,u)&&(c[u]=t[u]);c.originalType=e,c.mdxType="string"==typeof e?e:r,o[1]=c;for(var s=2;s<a;s++)o[s]=n[s];return i.a.createElement.apply(null,o)}return i.a.createElement.apply(null,n)}f.displayName="MDXCreateElement"},151:function(e,t,n){"use strict";var r=n(0),i=n(19);t.a=function(){var e=Object(r.useContext)(i.a);if(null===e)throw new Error("Docusaurus context not provided");return e}},152:function(e,t,n){"use strict";n.d(t,"b",(function(){return a})),n.d(t,"a",(function(){return o}));var r=n(151),i=n(153);function a(){var e=Object(r.a)().siteConfig,t=(e=void 0===e?{}:e).baseUrl,n=void 0===t?"/":t,a=e.url;return{withBaseUrl:function(e,t){return function(e,t,n,r){var a=void 0===r?{}:r,o=a.forcePrependBaseUrl,c=void 0!==o&&o,u=a.absolute,s=void 0!==u&&u;if(!n)return n;if(n.startsWith("#"))return n;if(Object(i.b)(n))return n;if(c)return t+n;var l=!n.startsWith(t)?t+n.replace(/^\//,""):n;return s?e+l:l}(a,n,e,t)}}}function o(e,t){return void 0===t&&(t={}),(0,a().withBaseUrl)(e,t)}},153:function(e,t,n){"use strict";function r(e){return!0===/^(\w*:|\/\/)/.test(e)}function i(e){return void 0!==e&&!r(e)}n.d(t,"b",(function(){return r})),n.d(t,"a",(function(){return i}))},86:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return c})),n.d(t,"metadata",(function(){return u})),n.d(t,"rightToc",(function(){return s})),n.d(t,"default",(function(){return p}));var r=n(2),i=n(6),a=(n(0),n(150)),o=n(152),c={id:"creating-and-viewing-outputs-files",sidebar_label:"Outputs",title:"Creating and viewing outputs files"},u={unversionedId:"creating-and-viewing-outputs-files",id:"creating-and-viewing-outputs-files",isDocsHomePage:!1,title:"Creating and viewing outputs files",description:"1. Write anything in context.output_dir: text, images, logs, pointclouds...",source:"@site/docs/creating-and-viewing-outputs-files.md",permalink:"/docs/creating-and-viewing-outputs-files",editUrl:"https://github.com/BlackspaceInc/BlackspacePlatform/edit/master/website/docs/creating-and-viewing-outputs-files.md",sidebar_label:"Outputs"},s=[{value:"Accessing output files",id:"accessing-output-files",children:[]}],l={rightToc:s};function p(e){var t=e.components,n=Object(i.a)(e,["components"]);return Object(a.b)("wrapper",Object(r.a)({},l,n,{components:t,mdxType:"MDXLayout"}),Object(a.b)("ol",null,Object(a.b)("li",{parentName:"ol"},"Write anything in ",Object(a.b)("inlineCode",{parentName:"li"},"context.output_dir"),": text, images, logs, pointclouds..."),Object(a.b)("li",{parentName:"ol"},'View in the web interface a list of all those files in the "Output Files" tab!'),Object(a.b)("li",{parentName:"ol"},"Click on a file to open it:")),Object(a.b)("img",{alt:"https://qa/tof/swip_tof/commit/42778afb1fea31e19c00291a2a52bf490e3acc2c?reference=a451dda9cfdd586702ead95f436e41c5b074ebfa&selected_views=bit_accuracy",src:Object(o.a)("img/output-files.png")}),Object(a.b)("p",null,"QA-Board will try to guess the right file viewer depending on the extension. Many are available, read the ",Object(a.b)("a",Object(r.a)({parentName:"p"},{href:"visualizations"}),"Read the visualizations guide")," to learn more."),Object(a.b)("blockquote",null,Object(a.b)("p",{parentName:"blockquote"},Object(a.b)("strong",{parentName:"p"},'"Visualizations"')," can help you declare pre-sets of relevant files. ",Object(a.b)("a",Object(r.a)({parentName:"p"},{href:"visualizations"}),"Read the docs")," to learn more!")),Object(a.b)("h2",{id:"accessing-output-files"},"Accessing output files"),Object(a.b)("p",null,"All the outputs are saved as files. To get them out and QA-Board provides multiple ways to get them out."),Object(a.b)("ol",null,Object(a.b)("li",{parentName:"ol"},Object(a.b)("strong",{parentName:"li"},"Next to each output"),", there is always a button to copy-to-clipboard the path to the files it created.")),Object(a.b)("img",{alt:"Export batch outputs",src:Object(o.a)("img/copy-windows-output-dir.png")}),Object(a.b)("ol",{start:2},Object(a.b)("li",{parentName:"ol"},Object(a.b)("strong",{parentName:"li"},"From the Navigation bar"),", you can copy-to-clipboard the windows-ish path where each commit saves its results:",Object(a.b)("img",{alt:"Export batch outputs",src:Object(o.a)("img/export-commit-folder.png")}))))}p.isMDXComponent=!0}}]);