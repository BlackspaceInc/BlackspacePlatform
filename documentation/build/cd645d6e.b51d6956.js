(window.webpackJsonp=window.webpackJsonp||[]).push([[128],{195:function(e,I,t){"use strict";t.r(I),t.d(I,"frontMatter",(function(){return n})),t.d(I,"metadata",(function(){return l})),t.d(I,"rightToc",(function(){return b})),t.d(I,"default",(function(){return C}));var a=t(2),i=t(6),g=(t(0),t(236)),n={title:"InfluxDB line protocol",description:"Tutorial showing how to use InfluxDB line protocol. QuestDB can listen for packets both over TCP and UDP."},l={unversionedId:"__uneeded/reference/influxdb-line-protocol",id:"__uneeded/reference/influxdb-line-protocol",isDocsHomePage:!1,title:"InfluxDB line protocol",description:"Tutorial showing how to use InfluxDB line protocol. QuestDB can listen for packets both over TCP and UDP.",source:"@site/docs/__uneeded/reference/influxdb-line-protocol.md",slug:"/__uneeded/reference/influxdb-line-protocol",permalink:"/docs/__uneeded/reference/influxdb-line-protocol",version:"current"},b=[{value:"Using line protocol",id:"using-line-protocol",children:[{value:"Syntax",id:"syntax",children:[]},{value:"Behaviour",id:"behaviour",children:[]},{value:"Examples",id:"examples",children:[]},{value:"Dealing with irregularly-structured data",id:"dealing-with-irregularly-structured-data",children:[]}]},{value:"TCP receiver",id:"tcp-receiver",children:[{value:"Overview",id:"overview",children:[]},{value:"Load balancing",id:"load-balancing",children:[]},{value:"Commit strategy",id:"commit-strategy",children:[]},{value:"Configuration",id:"configuration",children:[]}]},{value:"UDP receiver",id:"udp-receiver",children:[{value:"Overview",id:"overview-1",children:[]},{value:"Commit strategy",id:"commit-strategy-1",children:[]},{value:"Configuration",id:"configuration-1",children:[]}]}],c={rightToc:b};function C(e){var I=e.components,n=Object(i.a)(e,["components"]);return Object(g.b)("wrapper",Object(a.a)({},c,n,{components:I,mdxType:"MDXLayout"}),Object(g.b)("p",null,"InfluxDB line protocol ingestion makes it easy for existing InfluxDB users to\ntry QuestDB by only changing the address they send data to."),Object(g.b)("p",null,"It is not necessary to create a table schema beforehand: the table will be\ncreated on the fly. If new columns are added, the table is automatically updated\nto reflect the new structure."),Object(g.b)("p",null,"QuestDB can listen for line protocol packets both over ",Object(g.b)("a",Object(a.a)({parentName:"p"},{href:"#tcp-receiver"}),"TCP")," and\n",Object(g.b)("a",Object(a.a)({parentName:"p"},{href:"#udp-receiver"}),"UDP"),"."),Object(g.b)("h2",{id:"using-line-protocol"},"Using line protocol"),Object(g.b)("h3",{id:"syntax"},"Syntax"),Object(g.b)("pre",null,Object(g.b)("code",Object(a.a)({parentName:"pre"},{className:"language-shell",metastring:'title="ILP syntax"',title:'"ILP','syntax"':!0}),"table_name,tagset valueset timestamp\n")),Object(g.b)("table",null,Object(g.b)("thead",{parentName:"table"},Object(g.b)("tr",{parentName:"thead"},Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"Element"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"Definition"))),Object(g.b)("tbody",{parentName:"table"},Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),Object(g.b)("inlineCode",{parentName:"td"},"table_name")),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Name of the table where QuestDB will write data.")),Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),Object(g.b)("inlineCode",{parentName:"td"},"tagset")),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Array of string key-value pairs separated by commas that represent the reading's associated metadata")),Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),Object(g.b)("inlineCode",{parentName:"td"},"values")),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Array of key-value pairs separated by commas that represent the readings. The keys are string, values can be numeric or boolean")),Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),Object(g.b)("inlineCode",{parentName:"td"},"timestamp")),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"UNIX timestamp. By default in microseconds. Can be changed in the configuration")))),Object(g.b)("h3",{id:"behaviour"},"Behaviour"),Object(g.b)("ul",null,Object(g.b)("li",{parentName:"ul"},"When the ",Object(g.b)("inlineCode",{parentName:"li"},"table_name")," does not correspond to an existing table, QuestDB will\ncreate the table on the fly using the name provided. Column types will be\nautomatically recognized and assigned based on the data."),Object(g.b)("li",{parentName:"ul"},"The ",Object(g.b)("inlineCode",{parentName:"li"},"timestamp")," column is automatically created as\n",Object(g.b)("a",Object(a.a)({parentName:"li"},{href:"/docs/concept/designated-timestamp/"}),"designated timestamp")," with the\n",Object(g.b)("a",Object(a.a)({parentName:"li"},{href:"/docs/concept/partitions/"}),"partition strategy")," set to ",Object(g.b)("inlineCode",{parentName:"li"},"NONE"),". If you would\nlike to define a partition strategy, you should\n",Object(g.b)("a",Object(a.a)({parentName:"li"},{href:"/docs/reference/sql/create-table/"}),"CREATE")," the table beforehand."),Object(g.b)("li",{parentName:"ul"},"When the timestamp is empty, QuestDB will use the server timestamp.")),Object(g.b)("h3",{id:"examples"},"Examples"),Object(g.b)("p",null,"Let's assume the following data:"),Object(g.b)("table",null,Object(g.b)("thead",{parentName:"table"},Object(g.b)("tr",{parentName:"thead"},Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"timestamp"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"city"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"temperature"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"humidity"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"make"))),Object(g.b)("tbody",{parentName:"table"},Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1465839830100400000"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"London"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"23.5"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.343"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Omron")),Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1465839830100600000"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Bristol"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"23.2"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.443"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Honeywell")),Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1465839830100700000"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"London"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"23.6"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.358"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Omron")))),Object(g.b)("p",null,"Line protocol to insert this data in the ",Object(g.b)("inlineCode",{parentName:"p"},"readings")," table would look like this:"),Object(g.b)("pre",null,Object(g.b)("code",Object(a.a)({parentName:"pre"},{className:"language-shell"}),"readings,city=London,make=Omron temperature=23.5,humidity=0.343 1465839830100400000\nreadings,city=Bristol,make=Honeywell temperature=23.2,humidity=0.443 1465839830100600000\nreadings,city=London,make=Omron temperature=23.6,humidity=0.348 1465839830100700000\n")),Object(g.b)("div",{className:"admonition admonition-note alert alert--secondary"},Object(g.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-heading"}),Object(g.b)("h5",{parentName:"div"},Object(g.b)("span",Object(a.a)({parentName:"h5"},{className:"admonition-icon"}),Object(g.b)("svg",Object(a.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(g.b)("path",Object(a.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"})))),"note")),Object(g.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-content"}),Object(g.b)("p",{parentName:"div"},"There are only 2 spaces in each line. First between the ",Object(g.b)("inlineCode",{parentName:"p"},"tagset")," and ",Object(g.b)("inlineCode",{parentName:"p"},"values"),".\nSecond between ",Object(g.b)("inlineCode",{parentName:"p"},"values")," and ",Object(g.b)("inlineCode",{parentName:"p"},"timestamp"),"."))),Object(g.b)("h3",{id:"dealing-with-irregularly-structured-data"},"Dealing with irregularly-structured data"),Object(g.b)("div",{className:"admonition admonition-info alert alert--info"},Object(g.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-heading"}),Object(g.b)("h5",{parentName:"div"},Object(g.b)("span",Object(a.a)({parentName:"h5"},{className:"admonition-icon"}),Object(g.b)("svg",Object(a.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(g.b)("path",Object(a.a)({parentName:"svg"},{fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"})))),"info")),Object(g.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-content"}),Object(g.b)("p",{parentName:"div"},"QuestDB can support on-the-fly data structure changes with minimal overhead.\nShould users decide to send varying quantities of readings or metadata tags for\ndifferent entries, QuestDB will adapt on the fly."))),Object(g.b)("p",null,"InfluxDB line protocol makes it possible to send data under different shapes.\nEach new entry may contain certain metadata tags or readings, and others not.\nWhilst the example just above highlights structured data, it is possible for\nInfluxDB line protocol users to send data as follows."),Object(g.b)("pre",null,Object(g.b)("code",Object(a.a)({parentName:"pre"},{className:"language-shell"}),"readings,city=London temperature=23.2 1465839830100400000\nreadings,city=London temperature=23.6 1465839830100700000\nreadings,make=Honeywell temperature=23.2,humidity=0.443 1465839830100800000\n")),Object(g.b)("p",null,"Note that on the third line,"),Object(g.b)("ul",null,Object(g.b)("li",{parentName:"ul"},"a new ",Object(g.b)("inlineCode",{parentName:"li"},"tag"),' is added: "make"'),Object(g.b)("li",{parentName:"ul"},"a new ",Object(g.b)("inlineCode",{parentName:"li"},"field"),' is added: "humidity"')),Object(g.b)("p",null,"After writing two entries, the data would look like this"),Object(g.b)("table",null,Object(g.b)("thead",{parentName:"table"},Object(g.b)("tr",{parentName:"thead"},Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"timestamp"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"city"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"temperature"))),Object(g.b)("tbody",{parentName:"table"},Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1465839830100400000"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"London"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"23.5")),Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1465839830100700000"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"London"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"23.6")))),Object(g.b)("p",null,"The third entry would result in the following table"),Object(g.b)("table",null,Object(g.b)("thead",{parentName:"table"},Object(g.b)("tr",{parentName:"thead"},Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"timestamp"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"city"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"temperature"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"humidity"),Object(g.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"make"))),Object(g.b)("tbody",{parentName:"table"},Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1465839830100400000"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"London"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"23.5"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"NULL"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"NULL")),Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1465839830100700000"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"London"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"23.6"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"NULL"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"NULL")),Object(g.b)("tr",{parentName:"tbody"},Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"1465839830100800000"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"NULL"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"23.2"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"0.358"),Object(g.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Honeywell")))),Object(g.b)("div",{className:"admonition admonition-tip alert alert--success"},Object(g.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-heading"}),Object(g.b)("h5",{parentName:"div"},Object(g.b)("span",Object(a.a)({parentName:"h5"},{className:"admonition-icon"}),Object(g.b)("svg",Object(a.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"}),Object(g.b)("path",Object(a.a)({parentName:"svg"},{fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"})))),"tip")),Object(g.b)("div",Object(a.a)({parentName:"div"},{className:"admonition-content"}),Object(g.b)("p",{parentName:"div"},"Adding columns on the fly is no issue for QuestDB. New columns will be created\nin the affected partitions, and only populated if they contain values. Whilst we\noffer this function for flexibility. However, we recommend that users try to\nminimise structural changes to maintain operational simplicity."))),Object(g.b)("h2",{id:"tcp-receiver"},"TCP receiver"),Object(g.b)("p",null,"The TCP receiver can handle both single and multi-row write requests. It is\nfully multi-threaded and customizable. It can work from the common worker pool\nor out of dedicated threads. A load balancing mechanism dynamically assigns work\nbetween the threads."),Object(g.b)("h3",{id:"overview"},"Overview"),Object(g.b)("p",null,"By default, QuestDB listens to line protocol packets over TCP on ",Object(g.b)("inlineCode",{parentName:"p"},"0.0.0.0:9009"),".\nIf you are running QuestDB with Docker, you will need to map port 9009 using\n",Object(g.b)("inlineCode",{parentName:"p"},"-p 9009:9009 --net=host"),". This port can be customised."),Object(g.b)("p",null,"The TCP receiver uses two types of threads."),Object(g.b)("ul",null,Object(g.b)("li",{parentName:"ul"},Object(g.b)("strong",{parentName:"li"},"Worker threads")," - write data to the different tables. Each worker is\nwriting to designated tables. The worker-table repartition is modified over\ntime by the load balancing jobs."),Object(g.b)("li",{parentName:"ul"},Object(g.b)("strong",{parentName:"li"},"Network IO thread")," - receives messages from the network and adds them in a\nqueue for the writers. The network IO thread does not have a dedicated thread.\nInstead, it shares a common thread with the least busy worker.")),Object(g.b)("p",null,"The workflow is as follows."),Object(g.b)("p",null,Object(g.b)("img",{alt:"InfluxDB line protocol structure diagram",src:t(378).default})),Object(g.b)("p",null,"The network IO thread receives write requests and sets up a queue for the\nworkers. Workers pick up write requests for their assigned tables and insert the\ndata."),Object(g.b)("h3",{id:"load-balancing"},"Load balancing"),Object(g.b)("p",null,"A load balancing job reassigns work between threads in order to relieve the\nbusiest threads and maintain high ingestion speed. It can be triggered in two\nways."),Object(g.b)("ul",null,Object(g.b)("li",{parentName:"ul"},"After a certain number of updates per table"),Object(g.b)("li",{parentName:"ul"},"After a certain amount of time has passed")),Object(g.b)("p",null,"Once either is met, QuestDB will calculate a load ratio as the number of writes\nby the busiest thread divided by the number of writes in the least busy thread.\nIf this ratio is above the threshold, the table with the least writes in the\nbusiest worker thread will be reassigned to the least busy worker thread."),Object(g.b)("p",null,Object(g.b)("img",{alt:"InfluxDB line protocol load balancing diagram",src:t(379).default})),Object(g.b)("h3",{id:"commit-strategy"},"Commit strategy"),Object(g.b)("p",null,"Uncommitted rows are committed either:"),Object(g.b)("ul",null,Object(g.b)("li",{parentName:"ul"},"after ",Object(g.b)("inlineCode",{parentName:"li"},"line.tcp.maintenance.job.hysterisis.in.ms")," milliseconds have passed"),Object(g.b)("li",{parentName:"ul"},"once reaching ",Object(g.b)("inlineCode",{parentName:"li"},"line.tcp.max.uncommitted.rows")," uncommitted rows.")),Object(g.b)("h3",{id:"configuration"},"Configuration"),Object(g.b)("p",null,"The TCP receiver configuration can be completely customised using\n",Object(g.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/reference/configuration/server/#influxdb-line-protocol-tcp"}),"configuration keys"),".\nYou can use this to configure the tread pool, buffer and queue sizes, receiver\nIP address and port, load balancing etc."),Object(g.b)("h2",{id:"udp-receiver"},"UDP receiver"),Object(g.b)("p",null,"The UDP receiver can handle both single and multi row write requests. It is\ncurrently single-threaded, and performs both network IO and write jobs out of\none thread. The UDP worker thread can work either on its own thread or use the\ncommon thread pool. It supports both multicast and unicast."),Object(g.b)("p",null,"Find an example of how to use this in the\n",Object(g.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/api/java/#influxdb-sender-library/"}),"InfluxDB sender library section"),"."),Object(g.b)("h3",{id:"overview-1"},"Overview"),Object(g.b)("p",null,"By default, QuestDB listens for ",Object(g.b)("inlineCode",{parentName:"p"},"multicast")," line protocol packets over UDP on\n",Object(g.b)("inlineCode",{parentName:"p"},"232.1.2.3:9009"),". If you are running QuestDB with Docker, you will need to map\nport 9009 using ",Object(g.b)("inlineCode",{parentName:"p"},"-p 9009:9009 --net=host")," and publish multicast packets with TTL\nof at least 2. This port can be customised, and you can also configure QuestDB\nto listen for ",Object(g.b)("inlineCode",{parentName:"p"},"unicast"),"."),Object(g.b)("h3",{id:"commit-strategy-1"},"Commit strategy"),Object(g.b)("p",null,"Uncommitted rows are committed either:"),Object(g.b)("ul",null,Object(g.b)("li",{parentName:"ul"},"after receiving a number of continuous messages equal to\n",Object(g.b)("inlineCode",{parentName:"li"},"line.udp.commit.rate")),Object(g.b)("li",{parentName:"ul"},"when messages are no longer being received")),Object(g.b)("h3",{id:"configuration-1"},"Configuration"),Object(g.b)("p",null,"The UDP receiver configuration can be completely customised using\n",Object(g.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/reference/configuration/server/#influxdb-line-protocol-udp"}),"configuration keys"),".\nYou can use this to configure the IP address and port the receiver binds to,\ncommit rates, buffer size, whether it should run on a separate thread etc."))}C.isMDXComponent=!0},378:function(e,I,t){"use strict";t.r(I),I.default="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI5MjMiIGhlaWdodD0iMjU3Ij4KICAgIDxkZWZzPgogICAgICAgIDxzdHlsZSB0eXBlPSJ0ZXh0L2NzcyI+CiAgICAgICAgICAgIEBuYW1lc3BhY2UgImh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIjsKICAgICAgICAgICAgICAgIC5saW5lICAgICAgICAgICAgICAgICB7ZmlsbDogbm9uZTsgc3Ryb2tlOiAjNjM2MjczO30KICAgICAgICAgICAgICAgIC5ib2xkLWxpbmUgICAgICAgICAgICB7c3Ryb2tlOiAjNjM2MjczOyBzaGFwZS1yZW5kZXJpbmc6IGNyaXNwRWRnZXM7IHN0cm9rZS13aWR0aDogMjsgfQogICAgICAgICAgICAgICAgLnRoaW4tbGluZSAgICAgICAgICAgIHtzdHJva2U6ICM2MzYyNzM7IHNoYXBlLXJlbmRlcmluZzogY3Jpc3BFZGdlc30KICAgICAgICAgICAgICAgIC5maWxsZWQgICAgICAgICAgICAgICB7ZmlsbDogIzYzNjI3Mzsgc3Ryb2tlOiBub25lO30KICAgICAgICAgICAgICAgIHRleHQudGVybWluYWwgICAgICAgICB7Zm9udC1mYW1pbHk6IC1hcHBsZS1zeXN0ZW0sIEJsaW5rTWFjU3lzdGVtRm9udCwgIlNlZ29lIFVJIiwgUm9ib3RvLCBVYnVudHUsIENhbnRhcmVsbCwgSGVsdmV0aWNhLCBzYW5zLXNlcmlmOwogICAgICAgICAgICAgICAgZm9udC1zaXplOiAxMnB4OwogICAgICAgICAgICAgICAgZmlsbDogI2ZmZmZmZjsKICAgICAgICAgICAgICAgIGZvbnQtd2VpZ2h0OiBib2xkOwogICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgdGV4dC5ub250ZXJtaW5hbCAgICAgIHtmb250LWZhbWlseTogLWFwcGxlLXN5c3RlbSwgQmxpbmtNYWNTeXN0ZW1Gb250LCAiU2Vnb2UgVUkiLCBSb2JvdG8sIFVidW50dSwgQ2FudGFyZWxsLCBIZWx2ZXRpY2EsIHNhbnMtc2VyaWY7CiAgICAgICAgICAgICAgICBmb250LXNpemU6IDEycHg7CiAgICAgICAgICAgICAgICBmaWxsOiAjZTI4OWE0OwogICAgICAgICAgICAgICAgZm9udC13ZWlnaHQ6IG5vcm1hbDsKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIHRleHQucmVnZXhwICAgICAgICAgICB7Zm9udC1mYW1pbHk6IC1hcHBsZS1zeXN0ZW0sIEJsaW5rTWFjU3lzdGVtRm9udCwgIlNlZ29lIFVJIiwgUm9ib3RvLCBVYnVudHUsIENhbnRhcmVsbCwgSGVsdmV0aWNhLCBzYW5zLXNlcmlmOwogICAgICAgICAgICAgICAgZm9udC1zaXplOiAxMnB4OwogICAgICAgICAgICAgICAgZmlsbDogIzAwMTQxRjsKICAgICAgICAgICAgICAgIGZvbnQtd2VpZ2h0OiBub3JtYWw7CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICByZWN0LCBjaXJjbGUsIHBvbHlnb24ge2ZpbGw6IG5vbmU7IHN0cm9rZTogbm9uZTt9CiAgICAgICAgICAgICAgICByZWN0LnRlcm1pbmFsICAgICAgICAge2ZpbGw6IG5vbmU7IHN0cm9rZTogI2JlMmY1Yjt9CiAgICAgICAgICAgICAgICByZWN0Lm5vbnRlcm1pbmFsICAgICAge2ZpbGw6IHJnYmEoMjU1LDI1NSwyNTUsMC4xKTsgc3Ryb2tlOiBub25lO30KICAgICAgICAgICAgICAgIHJlY3QudGV4dCAgICAgICAgICAgICB7ZmlsbDogbm9uZTsgc3Ryb2tlOiBub25lO30KICAgICAgICAgICAgICAgIHBvbHlnb24ucmVnZXhwICAgICAgICB7ZmlsbDogI0M3RUNGRjsgc3Ryb2tlOiAjMDM4Y2JjO30KICAgICAgICA8L3N0eWxlPgogICAgPC9kZWZzPgogICAgPHBvbHlnb24gcG9pbnRzPSI5IDE3IDEgMTMgMSAyMSIvPgogICAgPHBvbHlnb24gcG9pbnRzPSIxNyAxNyA5IDEzIDkgMjEiLz48YSB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeGxpbms6aHJlZj0iI3VzZXJfMSIgeGxpbms6dGl0bGU9InVzZXJfMSI+CiAgICA8cmVjdCB4PSI1MSIgeT0iMyIgd2lkdGg9IjY0IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSI0OSIgeT0iMSIgd2lkdGg9IjY0IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iNTkiIHk9IjIxIj51c2VyXzE8L3RleHQ+PC9hPjxhIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4bGluazpocmVmPSIjdXNlcl8yIiB4bGluazp0aXRsZT0idXNlcl8yIj4KICAgIDxyZWN0IHg9IjUxIiB5PSI0NyIgd2lkdGg9IjY0IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSI0OSIgeT0iNDUiIHdpZHRoPSI2NCIgaGVpZ2h0PSIzMiIgY2xhc3M9Im5vbnRlcm1pbmFsIi8+CiAgICA8dGV4dCBjbGFzcz0ibm9udGVybWluYWwiIHg9IjU5IiB5PSI2NSI+dXNlcl8yPC90ZXh0PjwvYT48cmVjdCB4PSI1MSIgeT0iOTEiIHdpZHRoPSIzMiIgaGVpZ2h0PSIzMiIgcng9IjEwIi8+CiAgICA8cmVjdCB4PSI0OSIgeT0iODkiIHdpZHRoPSIzMiIgaGVpZ2h0PSIzMiIgY2xhc3M9InRlcm1pbmFsIiByeD0iMTAiLz4KICAgIDx0ZXh0IGNsYXNzPSJ0ZXJtaW5hbCIgeD0iNTkiIHk9IjEwOSI+Li4uPC90ZXh0PgogICAgPHJlY3QgeD0iMTU1IiB5PSIzIiB3aWR0aD0iMTA4IiBoZWlnaHQ9IjMyIiByeD0iMTAiLz4KICAgIDxyZWN0IHg9IjE1MyIgeT0iMSIgd2lkdGg9IjEwOCIgaGVpZ2h0PSIzMiIgY2xhc3M9InRlcm1pbmFsIiByeD0iMTAiLz4KICAgIDx0ZXh0IGNsYXNzPSJ0ZXJtaW5hbCIgeD0iMTYzIiB5PSIyMSI+VENQIHJlcXVlc3RzPC90ZXh0PjxhIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4bGluazpocmVmPSIjbmV0d29ya19JT190aHJlYWQiIHhsaW5rOnRpdGxlPSJuZXR3b3JrX0lPX3RocmVhZCI+CiAgICA8cmVjdCB4PSIyODMiIHk9IjMiIHdpZHRoPSIxNDQiIGhlaWdodD0iMzIiLz4KICAgIDxyZWN0IHg9IjI4MSIgeT0iMSIgd2lkdGg9IjE0NCIgaGVpZ2h0PSIzMiIgY2xhc3M9Im5vbnRlcm1pbmFsIi8+CiAgICA8dGV4dCBjbGFzcz0ibm9udGVybWluYWwiIHg9IjI5MSIgeT0iMjEiPm5ldHdvcmtfSU9fdGhyZWFkPC90ZXh0PjwvYT48cmVjdCB4PSI0NDciIHk9IjMiIHdpZHRoPSIxMDYiIGhlaWdodD0iMzIiIHJ4PSIxMCIvPgogICAgPHJlY3QgeD0iNDQ1IiB5PSIxIiB3aWR0aD0iMTA2IiBoZWlnaHQ9IjMyIiBjbGFzcz0idGVybWluYWwiIHJ4PSIxMCIvPgogICAgPHRleHQgY2xhc3M9InRlcm1pbmFsIiB4PSI0NTUiIHk9IjIxIj53cml0ZV9xdWV1ZTwvdGV4dD48YSB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeGxpbms6aHJlZj0iI3dvcmtlcl8xIiB4bGluazp0aXRsZT0id29ya2VyXzEiPgogICAgPHJlY3QgeD0iNTkzIiB5PSIzIiB3aWR0aD0iNzgiIGhlaWdodD0iMzIiLz4KICAgIDxyZWN0IHg9IjU5MSIgeT0iMSIgd2lkdGg9Ijc4IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iNjAxIiB5PSIyMSI+d29ya2VyXzE8L3RleHQ+PC9hPjxyZWN0IHg9IjY5MSIgeT0iMyIgd2lkdGg9IjU2IiBoZWlnaHQ9IjMyIiByeD0iMTAiLz4KICAgIDxyZWN0IHg9IjY4OSIgeT0iMSIgd2lkdGg9IjU2IiBoZWlnaHQ9IjMyIiBjbGFzcz0idGVybWluYWwiIHJ4PSIxMCIvPgogICAgPHRleHQgY2xhc3M9InRlcm1pbmFsIiB4PSI2OTkiIHk9IjIxIj53cml0ZTwvdGV4dD48YSB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeGxpbms6aHJlZj0iI3RhYmxlXzEiIHhsaW5rOnRpdGxlPSJ0YWJsZV8xIj4KICAgIDxyZWN0IHg9Ijc4NyIgeT0iMyIgd2lkdGg9IjY4IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSI3ODUiIHk9IjEiIHdpZHRoPSI2OCIgaGVpZ2h0PSIzMiIgY2xhc3M9Im5vbnRlcm1pbmFsIi8+CiAgICA8dGV4dCBjbGFzcz0ibm9udGVybWluYWwiIHg9Ijc5NSIgeT0iMjEiPnRhYmxlXzE8L3RleHQ+PC9hPjxhIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4bGluazpocmVmPSIjdGFibGVfMiIgeGxpbms6dGl0bGU9InRhYmxlXzIiPgogICAgPHJlY3QgeD0iNzg3IiB5PSI0NyIgd2lkdGg9IjY4IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSI3ODUiIHk9IjQ1IiB3aWR0aD0iNjgiIGhlaWdodD0iMzIiIGNsYXNzPSJub250ZXJtaW5hbCIvPgogICAgPHRleHQgY2xhc3M9Im5vbnRlcm1pbmFsIiB4PSI3OTUiIHk9IjY1Ij50YWJsZV8yPC90ZXh0PjwvYT48cmVjdCB4PSI3ODciIHk9IjkxIiB3aWR0aD0iMzIiIGhlaWdodD0iMzIiIHJ4PSIxMCIvPgogICAgPHJlY3QgeD0iNzg1IiB5PSI4OSIgd2lkdGg9IjMyIiBoZWlnaHQ9IjMyIiBjbGFzcz0idGVybWluYWwiIHJ4PSIxMCIvPgogICAgPHRleHQgY2xhc3M9InRlcm1pbmFsIiB4PSI3OTUiIHk9IjEwOSI+Li4uPC90ZXh0PjxhIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4bGluazpocmVmPSIjd29ya2VyXzIiIHhsaW5rOnRpdGxlPSJ3b3JrZXJfMiI+CiAgICA8cmVjdCB4PSI1OTMiIHk9IjEzNSIgd2lkdGg9Ijc4IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSI1OTEiIHk9IjEzMyIgd2lkdGg9Ijc4IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iNjAxIiB5PSIxNTMiPndvcmtlcl8yPC90ZXh0PjwvYT48cmVjdCB4PSI2OTEiIHk9IjEzNSIgd2lkdGg9IjU2IiBoZWlnaHQ9IjMyIiByeD0iMTAiLz4KICAgIDxyZWN0IHg9IjY4OSIgeT0iMTMzIiB3aWR0aD0iNTYiIGhlaWdodD0iMzIiIGNsYXNzPSJ0ZXJtaW5hbCIgcng9IjEwIi8+CiAgICA8dGV4dCBjbGFzcz0idGVybWluYWwiIHg9IjY5OSIgeT0iMTUzIj53cml0ZTwvdGV4dD48YSB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeGxpbms6aHJlZj0iI3RhYmxlXzMiIHhsaW5rOnRpdGxlPSJ0YWJsZV8zIj4KICAgIDxyZWN0IHg9Ijc4NyIgeT0iMTM1IiB3aWR0aD0iNjgiIGhlaWdodD0iMzIiLz4KICAgIDxyZWN0IHg9Ijc4NSIgeT0iMTMzIiB3aWR0aD0iNjgiIGhlaWdodD0iMzIiIGNsYXNzPSJub250ZXJtaW5hbCIvPgogICAgPHRleHQgY2xhc3M9Im5vbnRlcm1pbmFsIiB4PSI3OTUiIHk9IjE1MyI+dGFibGVfMzwvdGV4dD48L2E+PGEgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHhsaW5rOmhyZWY9IiN0YWJsZV80IiB4bGluazp0aXRsZT0idGFibGVfNCI+CiAgICA8cmVjdCB4PSI3ODciIHk9IjE3OSIgd2lkdGg9IjY4IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSI3ODUiIHk9IjE3NyIgd2lkdGg9IjY4IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iNzk1IiB5PSIxOTciPnRhYmxlXzQ8L3RleHQ+PC9hPjxyZWN0IHg9Ijc4NyIgeT0iMjIzIiB3aWR0aD0iMzIiIGhlaWdodD0iMzIiIHJ4PSIxMCIvPgogICAgPHJlY3QgeD0iNzg1IiB5PSIyMjEiIHdpZHRoPSIzMiIgaGVpZ2h0PSIzMiIgY2xhc3M9InRlcm1pbmFsIiByeD0iMTAiLz4KICAgIDx0ZXh0IGNsYXNzPSJ0ZXJtaW5hbCIgeD0iNzk1IiB5PSIyNDEiPi4uLjwvdGV4dD4KICAgIDxwYXRoIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGNsYXNzPSJsaW5lIiBkPSJtMTcgMTcgaDIgbTIwIDAgaDEwIG02NCAwIGgxMCBtLTEwNCAwIGgyMCBtODQgMCBoMjAgbS0xMjQgMCBxMTAgMCAxMCAxMCBtMTA0IDAgcTAgLTEwIDEwIC0xMCBtLTExNCAxMCB2MjQgbTEwNCAwIHYtMjQgbS0xMDQgMjQgcTAgMTAgMTAgMTAgbTg0IDAgcTEwIDAgMTAgLTEwIG0tOTQgMTAgaDEwIG02NCAwIGgxMCBtLTk0IC0xMCB2MjAgbTEwNCAwIHYtMjAgbS0xMDQgMjAgdjI0IG0xMDQgMCB2LTI0IG0tMTA0IDI0IHEwIDEwIDEwIDEwIG04NCAwIHExMCAwIDEwIC0xMCBtLTk0IDEwIGgxMCBtMzIgMCBoMTAgbTAgMCBoMzIgbTIwIC04OCBoMTAgbTEwOCAwIGgxMCBtMCAwIGgxMCBtMTQ0IDAgaDEwIG0wIDAgaDEwIG0xMDYgMCBoMTAgbTIwIDAgaDEwIG03OCAwIGgxMCBtMCAwIGgxMCBtNTYgMCBoMTAgbTIwIDAgaDEwIG02OCAwIGgxMCBtLTEwOCAwIGgyMCBtODggMCBoMjAgbS0xMjggMCBxMTAgMCAxMCAxMCBtMTA4IDAgcTAgLTEwIDEwIC0xMCBtLTExOCAxMCB2MjQgbTEwOCAwIHYtMjQgbS0xMDggMjQgcTAgMTAgMTAgMTAgbTg4IDAgcTEwIDAgMTAgLTEwIG0tOTggMTAgaDEwIG02OCAwIGgxMCBtLTk4IC0xMCB2MjAgbTEwOCAwIHYtMjAgbS0xMDggMjAgdjI0IG0xMDggMCB2LTI0IG0tMTA4IDI0IHEwIDEwIDEwIDEwIG04OCAwIHExMCAwIDEwIC0xMCBtLTk4IDEwIGgxMCBtMzIgMCBoMTAgbTAgMCBoMzYgbS0zMDIgLTg4IGgyMCBtMzAyIDAgaDIwIG0tMzQyIDAgcTEwIDAgMTAgMTAgbTMyMiAwIHEwIC0xMCAxMCAtMTAgbS0zMzIgMTAgdjExMiBtMzIyIDAgdi0xMTIgbS0zMjIgMTEyIHEwIDEwIDEwIDEwIG0zMDIgMCBxMTAgMCAxMCAtMTAgbS0zMTIgMTAgaDEwIG03OCAwIGgxMCBtMCAwIGgxMCBtNTYgMCBoMTAgbTIwIDAgaDEwIG02OCAwIGgxMCBtLTEwOCAwIGgyMCBtODggMCBoMjAgbS0xMjggMCBxMTAgMCAxMCAxMCBtMTA4IDAgcTAgLTEwIDEwIC0xMCBtLTExOCAxMCB2MjQgbTEwOCAwIHYtMjQgbS0xMDggMjQgcTAgMTAgMTAgMTAgbTg4IDAgcTEwIDAgMTAgLTEwIG0tOTggMTAgaDEwIG02OCAwIGgxMCBtLTk4IC0xMCB2MjAgbTEwOCAwIHYtMjAgbS0xMDggMjAgdjI0IG0xMDggMCB2LTI0IG0tMTA4IDI0IHEwIDEwIDEwIDEwIG04OCAwIHExMCAwIDEwIC0xMCBtLTk4IDEwIGgxMCBtMzIgMCBoMTAgbTAgMCBoMzYgbTQzIC0yMjAgaC0zIi8+CiAgICA8cG9seWdvbiBwb2ludHM9IjkxMyAxNyA5MjEgMTMgOTIxIDIxIi8+CiAgICA8cG9seWdvbiBwb2ludHM9IjkxMyAxNyA5MDUgMTMgOTA1IDIxIi8+PC9zdmc+Cg=="},379:function(e,I,t){"use strict";t.r(I),I.default="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI4MzEiIGhlaWdodD0iMTU3Ij4KICAgIDxkZWZzPgogICAgICAgIDxzdHlsZSB0eXBlPSJ0ZXh0L2NzcyI+CiAgICAgICAgICAgIEBuYW1lc3BhY2UgImh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIjsKICAgICAgICAgICAgICAgIC5saW5lICAgICAgICAgICAgICAgICB7ZmlsbDogbm9uZTsgc3Ryb2tlOiAjNjM2MjczO30KICAgICAgICAgICAgICAgIC5ib2xkLWxpbmUgICAgICAgICAgICB7c3Ryb2tlOiAjNjM2MjczOyBzaGFwZS1yZW5kZXJpbmc6IGNyaXNwRWRnZXM7IHN0cm9rZS13aWR0aDogMjsgfQogICAgICAgICAgICAgICAgLnRoaW4tbGluZSAgICAgICAgICAgIHtzdHJva2U6ICM2MzYyNzM7IHNoYXBlLXJlbmRlcmluZzogY3Jpc3BFZGdlc30KICAgICAgICAgICAgICAgIC5maWxsZWQgICAgICAgICAgICAgICB7ZmlsbDogIzYzNjI3Mzsgc3Ryb2tlOiBub25lO30KICAgICAgICAgICAgICAgIHRleHQudGVybWluYWwgICAgICAgICB7Zm9udC1mYW1pbHk6IC1hcHBsZS1zeXN0ZW0sIEJsaW5rTWFjU3lzdGVtRm9udCwgIlNlZ29lIFVJIiwgUm9ib3RvLCBVYnVudHUsIENhbnRhcmVsbCwgSGVsdmV0aWNhLCBzYW5zLXNlcmlmOwogICAgICAgICAgICAgICAgZm9udC1zaXplOiAxMnB4OwogICAgICAgICAgICAgICAgZmlsbDogI2ZmZmZmZjsKICAgICAgICAgICAgICAgIGZvbnQtd2VpZ2h0OiBib2xkOwogICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgdGV4dC5ub250ZXJtaW5hbCAgICAgIHtmb250LWZhbWlseTogLWFwcGxlLXN5c3RlbSwgQmxpbmtNYWNTeXN0ZW1Gb250LCAiU2Vnb2UgVUkiLCBSb2JvdG8sIFVidW50dSwgQ2FudGFyZWxsLCBIZWx2ZXRpY2EsIHNhbnMtc2VyaWY7CiAgICAgICAgICAgICAgICBmb250LXNpemU6IDEycHg7CiAgICAgICAgICAgICAgICBmaWxsOiAjZTI4OWE0OwogICAgICAgICAgICAgICAgZm9udC13ZWlnaHQ6IG5vcm1hbDsKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIHRleHQucmVnZXhwICAgICAgICAgICB7Zm9udC1mYW1pbHk6IC1hcHBsZS1zeXN0ZW0sIEJsaW5rTWFjU3lzdGVtRm9udCwgIlNlZ29lIFVJIiwgUm9ib3RvLCBVYnVudHUsIENhbnRhcmVsbCwgSGVsdmV0aWNhLCBzYW5zLXNlcmlmOwogICAgICAgICAgICAgICAgZm9udC1zaXplOiAxMnB4OwogICAgICAgICAgICAgICAgZmlsbDogIzAwMTQxRjsKICAgICAgICAgICAgICAgIGZvbnQtd2VpZ2h0OiBub3JtYWw7CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICByZWN0LCBjaXJjbGUsIHBvbHlnb24ge2ZpbGw6IG5vbmU7IHN0cm9rZTogbm9uZTt9CiAgICAgICAgICAgICAgICByZWN0LnRlcm1pbmFsICAgICAgICAge2ZpbGw6IG5vbmU7IHN0cm9rZTogI2JlMmY1Yjt9CiAgICAgICAgICAgICAgICByZWN0Lm5vbnRlcm1pbmFsICAgICAge2ZpbGw6IHJnYmEoMjU1LDI1NSwyNTUsMC4xKTsgc3Ryb2tlOiBub25lO30KICAgICAgICAgICAgICAgIHJlY3QudGV4dCAgICAgICAgICAgICB7ZmlsbDogbm9uZTsgc3Ryb2tlOiBub25lO30KICAgICAgICAgICAgICAgIHBvbHlnb24ucmVnZXhwICAgICAgICB7ZmlsbDogI0M3RUNGRjsgc3Ryb2tlOiAjMDM4Y2JjO30KICAgICAgICA8L3N0eWxlPgogICAgPC9kZWZzPgogICAgPHBvbHlnb24gcG9pbnRzPSI5IDEwNSAxIDEwMSAxIDEwOSIvPgogICAgPHBvbHlnb24gcG9pbnRzPSIxNyAxMDUgOSAxMDEgOSAxMDkiLz48YSB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeGxpbms6aHJlZj0iI0xhdW5jaF9Mb2FkX0JhbGFuY2luZ19Kb2IiIHhsaW5rOnRpdGxlPSJMYXVuY2hfTG9hZF9CYWxhbmNpbmdfSm9iIj4KICAgIDxyZWN0IHg9IjUxIiB5PSI5MSIgd2lkdGg9IjIwMiIgaGVpZ2h0PSIzMiIvPgogICAgPHJlY3QgeD0iNDkiIHk9Ijg5IiB3aWR0aD0iMjAyIiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iNTkiIHk9IjEwOSI+TGF1bmNoX0xvYWRfQmFsYW5jaW5nX0pvYjwvdGV4dD48L2E+PHJlY3QgeD0iMjkzIiB5PSIxMjMiIHdpZHRoPSIzNCIgaGVpZ2h0PSIzMiIgcng9IjEwIi8+CiAgICA8cmVjdCB4PSIyOTEiIHk9IjEyMSIgd2lkdGg9IjM0IiBoZWlnaHQ9IjMyIiBjbGFzcz0idGVybWluYWwiIHJ4PSIxMCIvPgogICAgPHRleHQgY2xhc3M9InRlcm1pbmFsIiB4PSIzMDEiIHk9IjE0MSI+SUY8L3RleHQ+PGEgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHhsaW5rOmhyZWY9IiNsb2FkX3JhdGlvIiB4bGluazp0aXRsZT0ibG9hZF9yYXRpbyI+CiAgICA8cmVjdCB4PSIzNDciIHk9IjEyMyIgd2lkdGg9Ijg0IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSIzNDUiIHk9IjEyMSIgd2lkdGg9Ijg0IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iMzU1IiB5PSIxNDEiPmxvYWRfcmF0aW88L3RleHQ+PC9hPjxyZWN0IHg9IjQ1MSIgeT0iMTIzIiB3aWR0aD0iMzAiIGhlaWdodD0iMzIiIHJ4PSIxMCIvPgogICAgPHJlY3QgeD0iNDQ5IiB5PSIxMjEiIHdpZHRoPSIzMCIgaGVpZ2h0PSIzMiIgY2xhc3M9InRlcm1pbmFsIiByeD0iMTAiLz4KICAgIDx0ZXh0IGNsYXNzPSJ0ZXJtaW5hbCIgeD0iNDU5IiB5PSIxNDEiPiZndDs8L3RleHQ+CiAgICA8cmVjdCB4PSI1MDEiIHk9IjEyMyIgd2lkdGg9IjExOCIgaGVpZ2h0PSIzMiIgcng9IjEwIi8+CiAgICA8cmVjdCB4PSI0OTkiIHk9IjEyMSIgd2lkdGg9IjExOCIgaGVpZ2h0PSIzMiIgY2xhc3M9InRlcm1pbmFsIiByeD0iMTAiLz4KICAgIDx0ZXh0IGNsYXNzPSJ0ZXJtaW5hbCIgeD0iNTA5IiB5PSIxNDEiPm1heC5sb2FkLnJhdGlvPC90ZXh0PjxhIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4bGluazpocmVmPSIjUmViYWxhbmNlX0xvYWQiIHhsaW5rOnRpdGxlPSJSZWJhbGFuY2VfTG9hZCI+CiAgICA8cmVjdCB4PSI2MzkiIHk9IjEyMyIgd2lkdGg9IjEyNCIgaGVpZ2h0PSIzMiIvPgogICAgPHJlY3QgeD0iNjM3IiB5PSIxMjEiIHdpZHRoPSIxMjQiIGhlaWdodD0iMzIiIGNsYXNzPSJub250ZXJtaW5hbCIvPgogICAgPHRleHQgY2xhc3M9Im5vbnRlcm1pbmFsIiB4PSI2NDciIHk9IjE0MSI+UmViYWxhbmNlX0xvYWQ8L3RleHQ+PC9hPjxyZWN0IHg9IjcxIiB5PSIzIiB3aWR0aD0iMjAyIiBoZWlnaHQ9IjMyIiByeD0iMTAiLz4KICAgIDxyZWN0IHg9IjY5IiB5PSIxIiB3aWR0aD0iMjAyIiBoZWlnaHQ9IjMyIiBjbGFzcz0idGVybWluYWwiIHJ4PSIxMCIvPgogICAgPHRleHQgY2xhc3M9InRlcm1pbmFsIiB4PSI3OSIgeT0iMjEiPm4udXBkYXRlcy5wZXIubG9hZC5iYWxhbmNlPC90ZXh0PgogICAgPHJlY3QgeD0iMjkzIiB5PSIzIiB3aWR0aD0iMzAiIGhlaWdodD0iMzIiIHJ4PSIxMCIvPgogICAgPHJlY3QgeD0iMjkxIiB5PSIxIiB3aWR0aD0iMzAiIGhlaWdodD0iMzIiIGNsYXNzPSJ0ZXJtaW5hbCIgcng9IjEwIi8+CiAgICA8dGV4dCBjbGFzcz0idGVybWluYWwiIHg9IjMwMSIgeT0iMjEiPiZsdDs8L3RleHQ+PGEgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHhsaW5rOmhyZWY9IiN1cGRhdGVzX3NpbmNlX2xhc3RfcmViYWxhbmNlIiB4bGluazp0aXRsZT0idXBkYXRlc19zaW5jZV9sYXN0X3JlYmFsYW5jZSI+CiAgICA8cmVjdCB4PSIzNDMiIHk9IjMiIHdpZHRoPSIyMTQiIGhlaWdodD0iMzIiLz4KICAgIDxyZWN0IHg9IjM0MSIgeT0iMSIgd2lkdGg9IjIxNCIgaGVpZ2h0PSIzMiIgY2xhc3M9Im5vbnRlcm1pbmFsIi8+CiAgICA8dGV4dCBjbGFzcz0ibm9udGVybWluYWwiIHg9IjM1MSIgeT0iMjEiPnVwZGF0ZXNfc2luY2VfbGFzdF9yZWJhbGFuY2U8L3RleHQ+PC9hPjxyZWN0IHg9IjcxIiB5PSI0NyIgd2lkdGg9IjI0NiIgaGVpZ2h0PSIzMiIgcng9IjEwIi8+CiAgICA8cmVjdCB4PSI2OSIgeT0iNDUiIHdpZHRoPSIyNDYiIGhlaWdodD0iMzIiIGNsYXNzPSJ0ZXJtaW5hbCIgcng9IjEwIi8+CiAgICA8dGV4dCBjbGFzcz0idGVybWluYWwiIHg9Ijc5IiB5PSI2NSI+bWFpbnRlbmFuY2Uuam9iLmh5c3RlcmVzaXMuaW4ubXM8L3RleHQ+CiAgICA8cmVjdCB4PSIzMzciIHk9IjQ3IiB3aWR0aD0iMzAiIGhlaWdodD0iMzIiIHJ4PSIxMCIvPgogICAgPHJlY3QgeD0iMzM1IiB5PSI0NSIgd2lkdGg9IjMwIiBoZWlnaHQ9IjMyIiBjbGFzcz0idGVybWluYWwiIHJ4PSIxMCIvPgogICAgPHRleHQgY2xhc3M9InRlcm1pbmFsIiB4PSIzNDUiIHk9IjY1Ij4mbHQ7PC90ZXh0PjxhIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4bGluazpocmVmPSIjdGltZV9zaW5jZV9sYXN0X3JlYmFsYW5jZSIgeGxpbms6dGl0bGU9InRpbWVfc2luY2VfbGFzdF9yZWJhbGFuY2UiPgogICAgPHJlY3QgeD0iMzg3IiB5PSI0NyIgd2lkdGg9IjE5MCIgaGVpZ2h0PSIzMiIvPgogICAgPHJlY3QgeD0iMzg1IiB5PSI0NSIgd2lkdGg9IjE5MCIgaGVpZ2h0PSIzMiIgY2xhc3M9Im5vbnRlcm1pbmFsIi8+CiAgICA8dGV4dCBjbGFzcz0ibm9udGVybWluYWwiIHg9IjM5NSIgeT0iNjUiPnRpbWVfc2luY2VfbGFzdF9yZWJhbGFuY2U8L3RleHQ+PC9hPjxyZWN0IHg9IjYxNyIgeT0iMyIgd2lkdGg9IjM0IiBoZWlnaHQ9IjMyIiByeD0iMTAiLz4KICAgIDxyZWN0IHg9IjYxNSIgeT0iMSIgd2lkdGg9IjM0IiBoZWlnaHQ9IjMyIiBjbGFzcz0idGVybWluYWwiIHJ4PSIxMCIvPgogICAgPHRleHQgY2xhc3M9InRlcm1pbmFsIiB4PSI2MjUiIHk9IjIxIj5JRjwvdGV4dD4KICAgIDxwYXRoIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGNsYXNzPSJsaW5lIiBkPSJtMTcgMTA1IGgyIG0yMCAwIGgxMCBtMjAyIDAgaDEwIG0yMCAwIGgxMCBtMCAwIGg0ODAgbS01MTAgMCBoMjAgbTQ5MCAwIGgyMCBtLTUzMCAwIHExMCAwIDEwIDEwIG01MTAgMCBxMCAtMTAgMTAgLTEwIG0tNTIwIDEwIHYxMiBtNTEwIDAgdi0xMiBtLTUxMCAxMiBxMCAxMCAxMCAxMCBtNDkwIDAgcTEwIDAgMTAgLTEwIG0tNTAwIDEwIGgxMCBtMzQgMCBoMTAgbTAgMCBoMTAgbTg0IDAgaDEwIG0wIDAgaDEwIG0zMCAwIGgxMCBtMCAwIGgxMCBtMTE4IDAgaDEwIG0wIDAgaDEwIG0xMjQgMCBoMTAgbS03NTIgLTMyIGwyMCAwIG0tMSAwIHEtOSAwIC05IC0xMCBsMCAtNjggcTAgLTEwIDEwIC0xMCBtNzUyIDg4IGwyMCAwIG0tMjAgMCBxMTAgMCAxMCAtMTAgbDAgLTY4IHEwIC0xMCAtMTAgLTEwIG0tNzMyIDAgaDEwIG0yMDIgMCBoMTAgbTAgMCBoMTAgbTMwIDAgaDEwIG0wIDAgaDEwIG0yMTQgMCBoMTAgbTAgMCBoMjAgbS01NDYgMCBoMjAgbTUyNiAwIGgyMCBtLTU2NiAwIHExMCAwIDEwIDEwIG01NDYgMCBxMCAtMTAgMTAgLTEwIG0tNTU2IDEwIHYyNCBtNTQ2IDAgdi0yNCBtLTU0NiAyNCBxMCAxMCAxMCAxMCBtNTI2IDAgcTEwIDAgMTAgLTEwIG0tNTM2IDEwIGgxMCBtMjQ2IDAgaDEwIG0wIDAgaDEwIG0zMCAwIGgxMCBtMCAwIGgxMCBtMTkwIDAgaDEwIG0yMCAtNDQgaDEwIG0zNCAwIGgxMCBtMCAwIGgxMzIgbTIzIDg4IGgtMyIvPgogICAgPHBvbHlnb24gcG9pbnRzPSI4MjEgMTA1IDgyOSAxMDEgODI5IDEwOSIvPgogICAgPHBvbHlnb24gcG9pbnRzPSI4MjEgMTA1IDgxMyAxMDEgODEzIDEwOSIvPjwvc3ZnPgo="}}]);