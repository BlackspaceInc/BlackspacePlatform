(window.webpackJsonp=window.webpackJsonp||[]).push([[133],{200:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return r})),n.d(t,"metadata",(function(){return s})),n.d(t,"rightToc",(function(){return l})),n.d(t,"default",(function(){return h}));var o=n(2),i=n(6),a=(n(0),n(236)),r={title:"Content hierarchy"},s={unversionedId:"__guidelines/content-hierarchy",id:"__guidelines/content-hierarchy",isDocsHomePage:!1,title:"Content hierarchy",description:"Documentation should follow a hierarchy, this is true both for the content and",source:"@site/docs/__guidelines/content-hierarchy.md",slug:"/__guidelines/content-hierarchy",permalink:"/docs/__guidelines/content-hierarchy",version:"current",sidebar:"docs",previous:{title:"Naming convention",permalink:"/docs/__guidelines/naming-convention"},next:{title:"Lexicon",permalink:"/docs/__guidelines/lexicon"}},l=[{value:"Content",id:"content",children:[]},{value:"Titles",id:"titles",children:[]},{value:"Bad practices",id:"bad-practices",children:[]}],c={rightToc:l};function h(e){var t=e.components,n=Object(i.a)(e,["components"]);return Object(a.b)("wrapper",Object(o.a)({},c,n,{components:t,mdxType:"MDXLayout"}),Object(a.b)("p",null,"Documentation should follow a hierarchy, this is true both for the content and\nhow titles are organized. In most cases, you can refer to a template and reuse\nthe hierarchy exposed there. When you write a page that does not derive from a\ntemplate, please follow the guidelines exposed here."),Object(a.b)("h2",{id:"content"},"Content"),Object(a.b)("p",null,"When you need to show a command, please show it at the top of your page as much\nas possible. This will ensure that users can easily copy/paste code without\nhaving to scan the whole page."),Object(a.b)("p",null,"It is okay to be very descriptive and thorough, however not every detail should\nhave the same weight. If you are documenting a function with many arguments,\nplease start with the most common ones first, gradually defining the ones that\npeople are less likely to use."),Object(a.b)("h2",{id:"titles"},"Titles"),Object(a.b)("p",null,"Pages need to start with text, not a title. Titles should always follow the\nfollowing hierarchy:"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-shell"}),"h1 (title) > h2 (##) > h3 (###) > h4 (####)\n")),Object(a.b)("p",null,"This will improve readability and SEO. Example:"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-markdown"}),"## The first title should be H2\n\nThen there should be some text\n\n### Then further titles should be H3\n\nThen ideally some text. Subsequent titles can then be used\n\n#### For example H4\n\nor even\n\n##### For example H5\n\netc\n")),Object(a.b)("h2",{id:"bad-practices"},"Bad practices"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"Repetitive subtitles"),Object(a.b)("li",{parentName:"ul"},"Too many Tip/Info/Warning. Please use maximum 1-2 per page.")))}h.isMDXComponent=!0}}]);