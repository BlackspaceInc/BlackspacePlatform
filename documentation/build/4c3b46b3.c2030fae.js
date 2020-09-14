(window.webpackJsonp=window.webpackJsonp||[]).push([[46],{113:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return c})),n.d(t,"metadata",(function(){return i})),n.d(t,"rightToc",(function(){return l})),n.d(t,"default",(function(){return s}));var a=n(2),b=n(6),r=(n(0),n(236)),c={title:"Row generator",sidebar_label:"Row generator",description:"Row generator function reference documentation."},i={unversionedId:"__uneeded/function/row-generator",id:"__uneeded/function/row-generator",isDocsHomePage:!1,title:"Row generator",description:"Row generator function reference documentation.",source:"@site/docs/__uneeded/function/row-generator.md",slug:"/__uneeded/function/row-generator",permalink:"/docs/__uneeded/function/row-generator",version:"current",sidebar_label:"Row generator"},l=[{value:"long_sequence",id:"long_sequence",children:[{value:"Arguments",id:"arguments",children:[]},{value:"Description",id:"description",children:[]},{value:"Examples",id:"examples",children:[]}]}],o={rightToc:l};function s(e){var t=e.components,n=Object(b.a)(e,["components"]);return Object(r.b)("wrapper",Object(a.a)({},o,n,{components:t,mdxType:"MDXLayout"}),Object(r.b)("h2",{id:"long_sequence"},"long_sequence"),Object(r.b)("ul",null,Object(r.b)("li",{parentName:"ul"},Object(r.b)("inlineCode",{parentName:"li"},"long_sequence(iterations)")," - generates rows"),Object(r.b)("li",{parentName:"ul"},Object(r.b)("inlineCode",{parentName:"li"},"long_sequence(iterations, seed1, seed2)")," - generates rows deterministically")),Object(r.b)("h3",{id:"arguments"},"Arguments"),Object(r.b)("p",null,"-",Object(r.b)("inlineCode",{parentName:"p"},"iterations"),": is a ",Object(r.b)("inlineCode",{parentName:"p"},"long")," representing the number of rows to generate. -",Object(r.b)("inlineCode",{parentName:"p"},"seed1"),"\nand ",Object(r.b)("inlineCode",{parentName:"p"},"seed2")," are ",Object(r.b)("inlineCode",{parentName:"p"},"long64")," representing both parts of a ",Object(r.b)("inlineCode",{parentName:"p"},"long128")," seed."),Object(r.b)("h3",{id:"description"},"Description"),Object(r.b)("h4",{id:"row-generation"},"Row generation"),Object(r.b)("p",null,Object(r.b)("inlineCode",{parentName:"p"},"long_sequence(iterations)")," is used to:"),Object(r.b)("ul",null,Object(r.b)("li",{parentName:"ul"},"Generate a number of rows defined by ",Object(r.b)("inlineCode",{parentName:"li"},"iterations"),"."),Object(r.b)("li",{parentName:"ul"},"Generate a column ",Object(r.b)("inlineCode",{parentName:"li"},"x:long")," of monotonically increasing long integers starting\nfrom 1, which can be accessed for queries.")),Object(r.b)("div",{className:"admonition admonition-tip alert alert--success"},Object(r.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-heading"}),Object(r.b)("h5",{parentName:"div"},Object(r.b)("span",Object(a.a)({parentName:"h5"},{className:"admonition-icon"}),Object(r.b)("svg",Object(a.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"}),Object(r.b)("path",Object(a.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"})))),"tip")),Object(r.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-content"}),Object(r.b)("p",{parentName:"div"},"You can use this to generate very large datasets for your testing e.g billions\nof rows or more if your disk allows."))),Object(r.b)("h4",{id:"random-number-seed"},"Random number seed"),Object(r.b)("p",null,"When ",Object(r.b)("inlineCode",{parentName:"p"},"long_sequence")," is used conjointly with\n",Object(r.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/reference/function/random-value-generator"}),"random generators"),", these\nvalues are usually generated at random. The function supports a seed to be\npassed in order to produce deterministic results."),Object(r.b)("div",{className:"admonition admonition-tip alert alert--success"},Object(r.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-heading"}),Object(r.b)("h5",{parentName:"div"},Object(r.b)("span",Object(a.a)({parentName:"h5"},{className:"admonition-icon"}),Object(r.b)("svg",Object(a.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"}),Object(r.b)("path",Object(a.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"})))),"tip")),Object(r.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-content"}),Object(r.b)("p",{parentName:"div"},"Deterministic procedural generation makes it easy to test on vasts amounts of\ndata without actually moving large files around across machines. Using the same\nseed on any machine at any time will consistently produce the same results for\nall random functions."))),Object(r.b)("h3",{id:"examples"},"Examples"),Object(r.b)("pre",null,Object(r.b)("code",Object(a.a)({parentName:"pre"},{className:"language-questdb-sql",metastring:'title="Generating multiple rows"',title:'"Generating',multiple:!0,'rows"':!0}),"SELECT x, rnd_double()\nFROM long_sequence(5);\n")),Object(r.b)("table",null,Object(r.b)("thead",{parentName:"table"},Object(r.b)("tr",{parentName:"thead"},Object(r.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"x"),Object(r.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"rnd_double"))),Object(r.b)("tbody",{parentName:"table"},Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.3279246687")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"2"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.8341038236")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"3"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.1023834675")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"4"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.9130602021")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"5"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.718276777")))),Object(r.b)("pre",null,Object(r.b)("code",Object(a.a)({parentName:"pre"},{className:"language-questdb-sql",metastring:'title="Accessing row_number using the x column"',title:'"Accessing',row_number:!0,using:!0,the:!0,x:!0,'column"':!0}),"SELECT x, x*x\nFROM long_sequence(5);\n")),Object(r.b)("table",null,Object(r.b)("thead",{parentName:"table"},Object(r.b)("tr",{parentName:"thead"},Object(r.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"x"),Object(r.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"x","*","x"))),Object(r.b)("tbody",{parentName:"table"},Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"2"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"4")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"3"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"9")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"4"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"16")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"5"),Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"25")))),Object(r.b)("pre",null,Object(r.b)("code",Object(a.a)({parentName:"pre"},{className:"language-questdb-sql",metastring:'title="Using with a seed"',title:'"Using',with:!0,a:!0,'seed"':!0}),"SELECT rnd_double()\nFROM long_sequence(2,128349234,4327897);\n")),Object(r.b)("div",{className:"admonition admonition-note alert alert--secondary"},Object(r.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-heading"}),Object(r.b)("h5",{parentName:"div"},Object(r.b)("span",Object(a.a)({parentName:"h5"},{className:"admonition-icon"}),Object(r.b)("svg",Object(a.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(r.b)("path",Object(a.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"})))),"note")),Object(r.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-content"}),Object(r.b)("p",{parentName:"div"},"The results below will be the same on any machine at any time as long as they\nuse the same seed in long_sequence."))),Object(r.b)("table",null,Object(r.b)("thead",{parentName:"table"},Object(r.b)("tr",{parentName:"thead"},Object(r.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"rnd_double"))),Object(r.b)("tbody",{parentName:"table"},Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.8251337821991485")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.2714941145110299")))))}s.isMDXComponent=!0}}]);