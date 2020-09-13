(window.webpackJsonp=window.webpackJsonp||[]).push([[63],{134:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return i})),n.d(t,"metadata",(function(){return s})),n.d(t,"rightToc",(function(){return c})),n.d(t,"default",(function(){return l}));var a=n(2),o=n(6),r=(n(0),n(243)),i={title:"Things we learned about sums",author:"Tancrede Collard",author_title:"QuestDB Team",author_url:"https://github.com/TheTanc",author_image_url:"https://avatars.githubusercontent.com/TheTanc",tags:["performance","deep-dive"],description:"What we learned implementing Kahan and Neumaier compensated sum algorithms, benchmark and comparison with Clickhouse."},s={permalink:"/blog/2020/05/12/interesting-things-we-learned-about-sums",source:"@site/blog/2020-05-12-interesting-things-we-learned-about-sums.md",description:"What we learned implementing Kahan and Neumaier compensated sum algorithms, benchmark and comparison with Clickhouse.",date:"2020-05-12T00:00:00.000Z",tags:[{label:"performance",permalink:"/blog/tags/performance"},{label:"deep-dive",permalink:"/blog/tags/deep-dive"}],title:"Things we learned about sums",readingTime:8.47,truncated:!0,prevItem:{title:"IoT on QuestDB",permalink:"/blog/2020/06/05/iot-on-questdb"},nextItem:{title:"Aggregating billions of rows per sec with SIMD",permalink:"/blog/2020/04/02/using-simd-to-aggregate-billions-of-rows-per-second"}},c=[],m={rightToc:c};function l(e){var t=e.components,n=Object(o.a)(e,["components"]);return Object(r.b)("wrapper",Object(a.a)({},m,n,{components:t,mdxType:"MDXLayout"}),Object(r.b)("img",{alt:"Wile E. Coyote and the Road Runner cartoon",className:"banner",src:"/img/blog/2020-05-12/banner.png"}),Object(r.b)("p",null,"In the world of databases, benchmarking performance has always been the hottest\ntopic. Who is faster for data ingestion and queries? About a month ago we\nannounced a new release with SIMD aggregations on\n",Object(r.b)("a",Object(a.a)({parentName:"p"},{href:"https://news.ycombinator.com/item?id=22803504"}),"HackerNews")," and\n",Object(r.b)("a",Object(a.a)({parentName:"p"},{href:"https://www.reddit.com/r/programming/comments/fwlk0k/questdb_using_simd_to_aggregate_billions_of/"}),"Reddit"),".\nFast. But were those results numerically accurate?"))}l.isMDXComponent=!0}}]);