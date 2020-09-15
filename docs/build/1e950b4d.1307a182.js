(window.webpackJsonp=window.webpackJsonp||[]).push([[22],{89:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return r})),n.d(t,"metadata",(function(){return s})),n.d(t,"rightToc",(function(){return c})),n.d(t,"default",(function(){return l}));var a=n(2),i=n(6),o=(n(0),n(236)),r={title:"Symbol",sidebar_label:"Symbol",description:"Description of the symbol data type. This QuestDB specific type is used to store repetitive strings in order to enable certain optimizations."},s={unversionedId:"__uneeded/concept/symbol",id:"__uneeded/concept/symbol",isDocsHomePage:!1,title:"Symbol",description:"Description of the symbol data type. This QuestDB specific type is used to store repetitive strings in order to enable certain optimizations.",source:"@site/docs/__uneeded/concept/symbol.md",slug:"/__uneeded/concept/symbol",permalink:"/docs/__uneeded/concept/symbol",version:"current",sidebar_label:"Symbol"},c=[{value:"Advantages",id:"advantages",children:[]},{value:"Usage",id:"usage",children:[]},{value:"Properties",id:"properties",children:[]}],b={rightToc:c};function l(e){var t=e.components,n=Object(i.a)(e,["components"]);return Object(o.b)("wrapper",Object(a.a)({},b,n,{components:t,mdxType:"MDXLayout"}),Object(o.b)("p",null,"QuestDB introduces a specific data type called ",Object(o.b)("inlineCode",{parentName:"p"},"symbol"),". It is a data structure\nused to store repetitive strings as a table of integers and corresponding string\nvalues."),Object(o.b)("h2",{id:"advantages"},"Advantages"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"reduced complexity of database schemas by removing the need for explicit extra\ntables and joins."),Object(o.b)("li",{parentName:"ul"},"transparent to the user: exact same behaviour as if the table was storing\nstring values, without the burden of actually doing so."),Object(o.b)("li",{parentName:"ul"},"greatly improved query performance (comparing and writing ",Object(o.b)("inlineCode",{parentName:"li"},"int")," instead of\n",Object(o.b)("inlineCode",{parentName:"li"},"string"),")"),Object(o.b)("li",{parentName:"ul"},"greatly improved storage efficiency (storing ",Object(o.b)("inlineCode",{parentName:"li"},"int")," instead of ",Object(o.b)("inlineCode",{parentName:"li"},"string"),")")),Object(o.b)("div",{className:"admonition admonition-note alert alert--secondary"},Object(o.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-heading"}),Object(o.b)("h5",{parentName:"div"},Object(o.b)("span",Object(a.a)({parentName:"h5"},{className:"admonition-icon"}),Object(o.b)("svg",Object(a.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(o.b)("path",Object(a.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"})))),"note")),Object(o.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-content"}),Object(o.b)("p",{parentName:"div"},Object(o.b)("inlineCode",{parentName:"p"},"symbol")," comparison across tables is not directly supported."))),Object(o.b)("h2",{id:"usage"},"Usage"),Object(o.b)("p",null,"To declare a column as ",Object(o.b)("inlineCode",{parentName:"p"},"SYMBOL")," please refer to the\n",Object(o.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/reference/sql/create-table/"}),"CREATE TABLE")," section. To create an ",Object(o.b)("inlineCode",{parentName:"p"},"INDEX"),"\non ",Object(o.b)("inlineCode",{parentName:"p"},"SYMBOL"),", please refer to the ",Object(o.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/concept/indexes/"}),"index")," section."),Object(o.b)("h2",{id:"properties"},"Properties"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"Symbol tables are stored separately from column data."),Object(o.b)("li",{parentName:"ul"},"Q conversion from ",Object(o.b)("inlineCode",{parentName:"li"},"string")," to ",Object(o.b)("inlineCode",{parentName:"li"},"int")," and vice-versa when reading or writing\ndata."),Object(o.b)("li",{parentName:"ul"},Object(o.b)("inlineCode",{parentName:"li"},"symbol")," supports indexes."),Object(o.b)("li",{parentName:"ul"},"For greater speed, ",Object(o.b)("inlineCode",{parentName:"li"},"symbol")," can be stored in the heap.")))}l.isMDXComponent=!0}}]);