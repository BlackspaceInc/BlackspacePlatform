(window.webpackJsonp=window.webpackJsonp||[]).push([[95],{164:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return l})),n.d(t,"metadata",(function(){return s})),n.d(t,"rightToc",(function(){return c})),n.d(t,"default",(function(){return o}));var r=n(2),a=n(6),i=(n(0),n(236)),l={title:"LIMIT keyword",sidebar_label:"LIMIT",description:"LIMIT SQL keyword reference documentation."},s={unversionedId:"__uneeded/reference/sql/limit",id:"__uneeded/reference/sql/limit",isDocsHomePage:!1,title:"LIMIT keyword",description:"LIMIT SQL keyword reference documentation.",source:"@site/docs/__uneeded/reference/sql/limit.md",slug:"/__uneeded/reference/sql/limit",permalink:"/docs/__uneeded/reference/sql/limit",version:"current",sidebar_label:"LIMIT"},c=[{value:"Syntax",id:"syntax",children:[]},{value:"Examples",id:"examples",children:[]}],b={rightToc:c};function o(e){var t=e.components,l=Object(a.a)(e,["components"]);return Object(i.b)("wrapper",Object(r.a)({},b,l,{components:t,mdxType:"MDXLayout"}),Object(i.b)("p",null,"Specify the number and position of records returned by a\n",Object(i.b)("a",Object(r.a)({parentName:"p"},{href:"/docs/reference/sql/select/"}),"SELECT statement"),"."),Object(i.b)("p",null,"In other implementations of SQL, this is sometimes replaced by statements such\nas ",Object(i.b)("inlineCode",{parentName:"p"},"OFFSET")," or ",Object(i.b)("inlineCode",{parentName:"p"},"ROWNUM")," Our implementation of ",Object(i.b)("inlineCode",{parentName:"p"},"LIMIT")," encompasses both in one\nstatement."),Object(i.b)("h2",{id:"syntax"},"Syntax"),Object(i.b)("p",null,Object(i.b)("img",{alt:"Flow chart showing the syntax of the LIMIT keyword",src:n(272).default})),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},Object(i.b)("inlineCode",{parentName:"li"},"numberOfRecords")," is the number of records to return."),Object(i.b)("li",{parentName:"ul"},Object(i.b)("inlineCode",{parentName:"li"},"upperBound")," and ",Object(i.b)("inlineCode",{parentName:"li"},"lowerBound")," is the return range. ",Object(i.b)("inlineCode",{parentName:"li"},"lowerBound")," is\n",Object(i.b)("strong",{parentName:"li"},"exclusive")," and ",Object(i.b)("inlineCode",{parentName:"li"},"upperBound")," is ",Object(i.b)("strong",{parentName:"li"},"inclusive"),".")),Object(i.b)("p",null,"A ",Object(i.b)("inlineCode",{parentName:"p"},"positive")," number will return the ",Object(i.b)("inlineCode",{parentName:"p"},"first")," n records. A ",Object(i.b)("inlineCode",{parentName:"p"},"negative")," number will\nreturn the ",Object(i.b)("inlineCode",{parentName:"p"},"last")," n records."),Object(i.b)("h2",{id:"examples"},"Examples"),Object(i.b)("pre",null,Object(i.b)("code",Object(r.a)({parentName:"pre"},{5:!0,className:"language-questdb-sql",metastring:'title="First 5 results"',title:'"First','results"':!0}),"SELECT * FROM ratings LIMIT 5;\n")),Object(i.b)("pre",null,Object(i.b)("code",Object(r.a)({parentName:"pre"},{5:!0,className:"language-questdb-sql",metastring:'title="Last 5 results"',title:'"Last','results"':!0}),"SELECT * FROM ratings LIMIT -5;\n")),Object(i.b)("pre",null,Object(i.b)("code",Object(r.a)({parentName:"pre"},{4:!0,className:"language-questdb-sql",metastring:'title="Range results - this will return records 3, 4 and 5"',title:'"Range',results:!0,"-":!0,this:!0,will:!0,return:!0,records:!0,"3,":!0,and:!0,'5"':!0}),"SELECT * FROM ratings LIMIT 2,5;\n")),Object(i.b)("p",null,Object(i.b)("inlineCode",{parentName:"p"},"negative")," range parameters will return results from the bottom of the table.\nAssuming a table with ",Object(i.b)("inlineCode",{parentName:"p"},"n")," records, the following will return records between n-7\n(exclusive) and n-3 (inclusive), i.e {n-6, n-5, n-4, n-3}"),Object(i.b)("pre",null,Object(i.b)("code",Object(r.a)({parentName:"pre"},{className:"language-questdb-sql",metastring:'title="Range results (negative)"',title:'"Range',results:!0,'(negative)"':!0}),"SELECT * FROM ratings LIMIT -7, -3;\n")))}o.isMDXComponent=!0}}]);