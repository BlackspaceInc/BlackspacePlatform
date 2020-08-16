(window.webpackJsonp=window.webpackJsonp||[]).push([[52],{104:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return o})),a.d(t,"metadata",(function(){return b})),a.d(t,"rightToc",(function(){return s})),a.d(t,"default",(function(){return u}));var c=a(2),n=a(6),r=(a(0),a(146)),i=a(148),o={id:"bit-accuracy",sidebar_label:"Bit Accuracy",title:"Bit accuracy tests"},b={unversionedId:"bit-accuracy",id:"bit-accuracy",isDocsHomePage:!1,title:"Bit accuracy tests",description:"QA-Board was started at Samsung, within a business unit focused on hardware digital design. Because of those root, QA-Board provides a number of ways to check that results are equal from commit to commit.",source:"@site/docs/bit-accuracy.md",permalink:"/bit-accuracy",editUrl:"https://github.com/BlackspaceInc/BlackspacePlatform/edit/master/website/docs/bit-accuracy.md",sidebar_label:"Bit Accuracy"},s=[{value:"&quot;Soft&quot; bit-accuracy checks from the UI",id:"soft-bit-accuracy-checks-from-the-ui",children:[{value:"&quot;Hard&quot; <code>qa check-bit-accuracy</code> on the CLI",id:"hard-qa-check-bit-accuracy-on-the-cli",children:[]},{value:"What files are checked?",id:"what-files-are-checked",children:[]}]},{value:"Sample CI for bit-accuracy checks",id:"sample-ci-for-bit-accuracy-checks",children:[]}],l={rightToc:s};function u(e){var t=e.components,a=Object(n.a)(e,["components"]);return Object(r.b)("wrapper",Object(c.a)({},l,a,{components:t,mdxType:"MDXLayout"}),Object(r.b)("p",null,"QA-Board was started at Samsung, within a business unit focused on hardware digital design. Because of those root, QA-Board provides a number of ways to check that results are equal from commit to commit."),Object(r.b)("h2",{id:"soft-bit-accuracy-checks-from-the-ui"},'"Soft" bit-accuracy checks from the UI'),Object(r.b)("p",null,"The web application lets you view and compare all files created by your algorithm's runs:"),Object(r.b)("ul",null,Object(r.b)("li",{parentName:"ul"},"Files are marked depending on their status (identical, different, added, removed...). Identical files are hidden by default."),Object(r.b)("li",{parentName:"ul"},"You can click on each file to open it with an appropriate file viewer:")),Object(r.b)("img",{alt:"bit accuracy viewer",src:Object(i.a)("img/bit-accuracy-viewer.jpg")}),Object(r.b)("div",{className:"admonition admonition-note alert alert--secondary"},Object(r.b)("div",Object(c.a)({parentName:"div"},{className:"admonition-heading"}),Object(r.b)("h5",{parentName:"div"},Object(r.b)("span",Object(c.a)({parentName:"h5"},{className:"admonition-icon"}),Object(r.b)("svg",Object(c.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(r.b)("path",Object(c.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"})))),"note")),Object(r.b)("div",Object(c.a)({parentName:"div"},{className:"admonition-content"}),Object(r.b)("p",{parentName:"div"},"The UI doesn't care about ",Object(r.b)("a",Object(c.a)({parentName:"p"},{href:"https://github.com/Samsung/qaboard/blob/master/qaboard/sample_project/qaboard.yaml#L93"}),Object(r.b)("em",{parentName:"a"},"qaboard.yaml")),"'s ",Object(r.b)("inlineCode",{parentName:"p"},"bit-accuracy.patterns")," ",Object(r.b)("em",{parentName:"p"},"(discussed later)"),"."))),Object(r.b)("h3",{id:"hard-qa-check-bit-accuracy-on-the-cli"},'"Hard" ',Object(r.b)("inlineCode",{parentName:"h3"},"qa check-bit-accuracy")," on the CLI"),Object(r.b)("p",null,Object(r.b)("inlineCode",{parentName:"p"},"qa check-bit-accuracy $batch")," compares the results of ",Object(r.b)("inlineCode",{parentName:"p"},"qa batch $batch")," to:"),Object(r.b)("ul",null,Object(r.b)("li",{parentName:"ul"},"The latest results on ",Object(r.b)("inlineCode",{parentName:"li"},"project.reference_branch")," from ",Object(r.b)("a",Object(c.a)({parentName:"li"},{href:"https://github.com/Samsung/qaboard/blob/master/qaboard/sample_project/qaboard.yaml"}),Object(r.b)("em",{parentName:"a"},"qaboard.yaml"))," (default: ",Object(r.b)("em",{parentName:"li"},"master"),")."),Object(r.b)("li",{parentName:"ul"},"...unlesss you're checking a merge made to that branch. In which case the commit's parents will act as references."),Object(r.b)("li",{parentName:"ul"},"You can ask to compare versus a specific git commit, branch or tag with ",Object(r.b)("inlineCode",{parentName:"li"},"qa check-bit-accuracy --reference $git-ref"),".")),Object(r.b)("blockquote",null,Object(r.b)("p",{parentName:"blockquote"},"If the commit you compare against has not finished its CI, ",Object(r.b)("inlineCode",{parentName:"p"},"qa")," will  wait.")),Object(r.b)("div",{className:"admonition admonition-note alert alert--secondary"},Object(r.b)("div",Object(c.a)({parentName:"div"},{className:"admonition-heading"}),Object(r.b)("h5",{parentName:"div"},Object(r.b)("span",Object(c.a)({parentName:"h5"},{className:"admonition-icon"}),Object(r.b)("svg",Object(c.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(r.b)("path",Object(c.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"})))),"Custom Needs")),Object(r.b)("div",Object(c.a)({parentName:"div"},{className:"admonition-content"}),Object(r.b)("p",{parentName:"div"},"You can opt-in to more complex behaviour in ",Object(r.b)("em",{parentName:"p"},Object(r.b)("a",Object(c.a)({parentName:"em"},{href:"https://github.com/Samsung/qaboard/blob/master/qaboard/sample_project/qaboard.yaml"}),"qaboard.yaml"))," with ",Object(r.b)("inlineCode",{parentName:"p"},"bit-accuracy.on_reference_failed_ci"),", in case there are not results in the reference commit. Maybe the build failed, in which case you want to compare against the previous commit... If you're interested open an issue we'll add more details to the docs."))),Object(r.b)("p",null,"If output files are different, ",Object(r.b)("inlineCode",{parentName:"p"},"qa")," prints a report and exists with a failure."),Object(r.b)("div",{className:"admonition admonition-note alert alert--secondary"},Object(r.b)("div",Object(c.a)({parentName:"div"},{className:"admonition-heading"}),Object(r.b)("h5",{parentName:"div"},Object(r.b)("span",Object(c.a)({parentName:"h5"},{className:"admonition-icon"}),Object(r.b)("svg",Object(c.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(r.b)("path",Object(c.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"})))),"note")),Object(r.b)("div",Object(c.a)({parentName:"div"},{className:"admonition-content"}),Object(r.b)("p",{parentName:"div"},"For specific use-case, there is also ",Object(r.b)("inlineCode",{parentName:"p"},"qa check-bit-accuracy-manifest")," which checks accuracy versus a manifest file stored in the database next to the input files. To generate those manifests, use ",Object(r.b)("inlineCode",{parentName:"p"},"qa batch --save-manifests-in-database"),"."))),Object(r.b)("h3",{id:"what-files-are-checked"},"What files are checked?"),Object(r.b)("p",null,"Files matching the patterns defined as ",Object(r.b)("inlineCode",{parentName:"p"},"bit-accuracy.patterns")," in ",Object(r.b)("a",Object(c.a)({parentName:"p"},{href:"https://github.com/Samsung/qaboard/blob/master/qaboard/sample_project/qaboard.yaml#L93"}),Object(r.b)("em",{parentName:"a"},"qaboard.yaml"))," will be checked."),Object(r.b)("div",{className:"admonition admonition-note alert alert--secondary"},Object(r.b)("div",Object(c.a)({parentName:"div"},{className:"admonition-heading"}),Object(r.b)("h5",{parentName:"div"},Object(r.b)("span",Object(c.a)({parentName:"h5"},{className:"admonition-icon"}),Object(r.b)("svg",Object(c.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(r.b)("path",Object(c.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"})))),"note")),Object(r.b)("div",Object(c.a)({parentName:"div"},{className:"admonition-content"}),Object(r.b)("p",{parentName:"div"},"If you work with text files on both Linux and Windows, EOL can make things tricky... You can decide what files are plaintext or binary using ",Object(r.b)("inlineCode",{parentName:"p"},"bit-accuracy.plaintext")," ",Object(r.b)("em",{parentName:"p"},"or")," ",Object(r.b)("inlineCode",{parentName:"p"},"bit-accuracy.binary"),"."))),Object(r.b)("h2",{id:"sample-ci-for-bit-accuracy-checks"},"Sample CI for bit-accuracy checks"),Object(r.b)("p",null,"You often want to know when your algorithm's results change, ",Object(r.b)("em",{parentName:"p"},"especially if another team is busy implementing them in hardware"),"!"),Object(r.b)("p",null,"Here is how you could get the CI to warn you with GitlabCI:"),Object(r.b)("pre",null,Object(r.b)("code",Object(c.a)({parentName:"pre"},{className:"language-yaml",metastring:'title="qaboard.yaml"',title:'"qaboard.yaml"'}),"stages:\n  - tests\n  - bit-accuracy\n\ntests-all:\n  stage: tests\n  script:\n  - qa batch all\n\nbit-accuracy-all:\n  stage: bit-accuracy\n  allowed_failure: true\n  script:\n  - qa check-bit-accuracy --batch all\n")))}u.isMDXComponent=!0},146:function(e,t,a){"use strict";a.d(t,"a",(function(){return u})),a.d(t,"b",(function(){return p}));var c=a(0),n=a.n(c);function r(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function i(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);t&&(c=c.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,c)}return a}function o(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?i(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):i(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function b(e,t){if(null==e)return{};var a,c,n=function(e,t){if(null==e)return{};var a,c,n={},r=Object.keys(e);for(c=0;c<r.length;c++)a=r[c],t.indexOf(a)>=0||(n[a]=e[a]);return n}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(c=0;c<r.length;c++)a=r[c],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(n[a]=e[a])}return n}var s=n.a.createContext({}),l=function(e){var t=n.a.useContext(s),a=t;return e&&(a="function"==typeof e?e(t):o(o({},t),e)),a},u=function(e){var t=l(e.components);return n.a.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.a.createElement(n.a.Fragment,{},t)}},m=n.a.forwardRef((function(e,t){var a=e.components,c=e.mdxType,r=e.originalType,i=e.parentName,s=b(e,["components","mdxType","originalType","parentName"]),u=l(a),m=c,p=u["".concat(i,".").concat(m)]||u[m]||d[m]||r;return a?n.a.createElement(p,o(o({ref:t},s),{},{components:a})):n.a.createElement(p,o({ref:t},s))}));function p(e,t){var a=arguments,c=t&&t.mdxType;if("string"==typeof e||c){var r=a.length,i=new Array(r);i[0]=m;var o={};for(var b in t)hasOwnProperty.call(t,b)&&(o[b]=t[b]);o.originalType=e,o.mdxType="string"==typeof e?e:c,i[1]=o;for(var s=2;s<r;s++)i[s]=a[s];return n.a.createElement.apply(null,i)}return n.a.createElement.apply(null,a)}m.displayName="MDXCreateElement"},147:function(e,t,a){"use strict";var c=a(0),n=a(19);t.a=function(){var e=Object(c.useContext)(n.a);if(null===e)throw new Error("Docusaurus context not provided");return e}},148:function(e,t,a){"use strict";a.d(t,"b",(function(){return r})),a.d(t,"a",(function(){return i}));var c=a(147),n=a(149);function r(){var e=Object(c.a)().siteConfig,t=(e=void 0===e?{}:e).baseUrl,a=void 0===t?"/":t,r=e.url;return{withBaseUrl:function(e,t){return function(e,t,a,c){var r=void 0===c?{}:c,i=r.forcePrependBaseUrl,o=void 0!==i&&i,b=r.absolute,s=void 0!==b&&b;if(!a)return a;if(a.startsWith("#"))return a;if(Object(n.b)(a))return a;if(o)return t+a;var l=!a.startsWith(t)?t+a.replace(/^\//,""):a;return s?e+l:l}(r,a,e,t)}}}function i(e,t){return void 0===t&&(t={}),(0,r().withBaseUrl)(e,t)}},149:function(e,t,a){"use strict";function c(e){return!0===/^(\w*:|\/\/)/.test(e)}function n(e){return void 0!==e&&!c(e)}a.d(t,"b",(function(){return c})),a.d(t,"a",(function(){return n}))}}]);