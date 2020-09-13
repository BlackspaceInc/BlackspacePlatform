---
title: Distributed Systems Learning Resources
---

Most links tend to be readings on architecture itself rather than code itself.

## Bootcamp
Read things here before you start.
- [CAP Theorem](http://en.wikipedia.org/wiki/CAP_theorem), Also [plain english](http://ksat.me/a-plain-english-introduction-to-cap-theorem) explanation
- [Fallacies of Distributed Computing](http://en.wikipedia.org/wiki/Fallacies_of_distributed_computing), expect things to break, *everything*
- [Distributed systems theory for the distributed engineer](http://the-paper-trail.org/blog/distributed-systems-theory-for-the-distributed-systems-engineer/), most of the papers/books in the blog might reappear in this list again. Still a good BFS approach to distributed systems.
- [FLP Impossibility Result (paper)](https://groups.csail.mit.edu/tds/papers/Lynch/jacm85.pdf), an easier [blog post](http://the-paper-trail.org/blog/a-brief-tour-of-flp-impossibility/) to follow along
- [An Introduction to Distributed Systems](https://github.com/aphyr/distsys-class) @aphyr's excellent introduction to distributed systems 

## Books
- [Distributed Systems for fun and profit](http://book.mixu.net/distsys/single-page.html) [Free]
- [Distributed Systems Principles and Paradigms, Andrew Tanenbaum](https://www.distributed-systems.net/index.php/books/distributed-systems-3rd-edition-2017/ds3-sneak-preview/) [Free with registration]
- [Scalable Web Architecture and Distributed Systems](http://www.aosabook.org/en/distsys.html) [Free]
- [Principles of Distributed Systems](http://dcg.ethz.ch/lectures/podc_allstars/lecture/podc.pdf) [Free] [ETH Zurich University]
- [Making reliable distributed systems in the presence of software errors](http://www.erlang.org/download/armstrong_thesis_2003.pdf), [Free] Joe Amstrong's (Author of Erlang) PhD thesis 
- [Designing Data Intensive Applications](https://www.amazon.com/Designing-Data-Intensive-Applications-Reliable-Maintainable/dp/1449373321) [Amazon Link]
- [Distributed Computing, By Hagit Attiya and Jennifer Welch](http://hagit.net.technion.ac.il/publications/dc/)
- [Distributed Algorithms, Nancy Lynch](https://www.amazon.com/Distributed-Algorithms-Kaufmann-Management-Systems/dp/1558603484) [Amazon Link]
- [Impossibility Results for Distributed Computing](http://www.morganclaypool.com/doi/abs/10.2200/S00551ED1V01Y201311DCT012) [paywall]
- [Designing Distributed Systems, Brandon Burns](https://azure.microsoft.com/en-us/resources/designing-distributed-systems/) [Free with registration]
- [Making Sense of Edge Computing](https://www.manning.com/books/making-sense-of-edge-computing)

## Papers
Must read papers on distributed systems. While nearly *all* of Lamport's work should feature here, just adding a few that *must* be read.
- [Times, Clocks and Ordering of Events in Distributed Systems](http://research.microsoft.com/en-us/um/people/lamport/pubs/time-clocks.pdf) Lamport's paper, the Quintessential distributed systems primer
- [Session Guarantees for Weakly Consistent Replicated Data](http://www.cs.utexas.edu/~dahlin/Classes/GradOS/papers/SessionGuaranteesPDIS.pdf) a '94 paper that talks about various recommendations for session guarantees for eventually consistent systems, many of this would be standard vocabulary in reading other dist. sys papers, like monotonic reads, read your writes etc.

### Storage & Databases
- [Dynamo: Amazon's Highly Available Key Value Store](http://bnrg.eecs.berkeley.edu/~randy/Courses/CS294.F07/Dynamo.pdf)
Paraphrasing @fogus from their [blog](http://blog.fogus.me/2011/09/08/10-technical-papers-every-programmer-should-read-at-least-twice/), it is very rare for a paper describing an active production system to influence the state of active research in any industry; this is one of those seminal distributed systems paper that solves the problem of a highly available and fault tolerant database in an elegant way, later paving the way for systems like Cassandra, and many other AP systems using a consistent hashing.
- [Bigtable: A Distributed Storage System for Structured Data](http://static.googleusercontent.com/media/research.google.com/en//archive/bigtable-osdi06.pdf)
- [The Google File System](http://static.googleusercontent.com/external_content/untrusted_dlcp/research.google.com/en/us/archive/gfs-sosp2003.pdf)
- [Cassandra: A Decentralized Structured Storage System](http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.161.6751&rep=rep1&type=pdf) Inspired heavily by Dynamo, an now an open source 
- [CRUSH: Controlled, Scalable, Decentralized Placement of Replicated Data](http://www.ssrc.ucsc.edu/Papers/weil-sc06.pdf), the algorithm for the basis of Ceph distributed storage system, for the architecture itself read [RADOS](http://ceph.com/papers/weil-rados-pdsw07.pdf)

### Messaging systems
- [The Log: What every software engineer should know about real-time data's unifying abstraction](http://engineering.linkedin.com/distributed-systems/log-what-every-software-engineer-should-know-about-real-time-datas-unifying), a somewhat long read, but covers brilliantly on logs, which are at the heart of most distributed systems
- [Kafka: a Distributed Messaging System for Log Processing](http://notes.stephenholiday.com/Kafka.pdf)

### Distributed Consensus and Fault-Tolerance
- [Practical Byzantine Fault Tolerance](http://pmg.csail.mit.edu/papers/osdi99.pdf)
- [The Byzantine Generals Problem](http://bnrg.cs.berkeley.edu/~adj/cs16x/hand-outs/Original_Byzantine.pdf)
- [Impossibility of Distributed Consensus with One Faulty Process](http://macs.citadel.edu/rudolphg/csci604/ImpossibilityofConsensus.pdf)
- [The Part Time Parliament](http://research.microsoft.com/en-us/um/people/lamport/pubs/lamport-paxos.pdf) Paxos, Lamport's original Paxos paper, a bit difficult to understand, may require multiple passes
- [Paxos Made Simple](http://research.microsoft.com/en-us/um/people/lamport/pubs/paxos-simple.pdf), a more terse readable Paxos paper by Lamport himself. Shorter and more easier compared to the original.
- [The Chubby Lock Service for loosely coupled distributed systems](http://static.googleusercontent.com/media/research.google.com/en//archive/chubby-osdi06.pdf) Google's lock service used for loosely coupled distributed systems. Sort of Paxos as a Service for building other distributed systems. Primary inspiration behind other Service Discovery & Coordination tools like Zookeeper, etcd, Consul etc.
- [Paxos made live - An engineering perspective](http://research.google.com/archive/paxos_made_live.html) Google's learning while implementing systems atop of Paxos. Demonstrates various practical issues encountered while implementing a theoretical concept.
- [Raft Consensus Algorithm](https://raftconsensus.github.io/) An alternative to Paxos for distributed consensus, that is much simpler to understand. Do checkout an [interesting visualization of raft](http://thesecretlivesofdata.com/raft/)
- [Conflict-free Replicated Data Types](https://pages.lip6.fr/Marc.Shapiro/papers/RR-7687.pdf) presents an approach for Strong Eventual Consistency which as been applied in projects such as [Riak](http://basho.com/products/riak-kv/), [Redis](https://redis.io/) and [Akka](https://akka.io/). A great talk on the subject by Martin Kleppmann can be found [here](https://www.youtube.com/watch?v=B5NULPSiOGw) 

### Testing, monitoring and tracing
While designing distributed systems are hard enough, testing them is even harder. 
- [Dapper](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/36356.pdf), Google's large scale distributed-systems tracing infrastructure, this was also the basis for the design of open source projects such as [Zipkin](http://zipkin.io/), [Apache SkyWalking](https://github.com/apache/incubator-skywalking), [Pinpoint](https://github.com/naver/pinpoint) and [HTrace](http://htrace.incubator.apache.org/).

### Programming Models
- [Distributed Programming Model](http://web.cs.ucdavis.edu/~pandey/Research/Papers/icdcs01.pdf)
- [PSync: a partially synchronous language for fault-tolerant distributed algorithms](http://www.di.ens.fr/~cezarad/popl16.pdf) Video: [Conference Video](https://www.youtube.com/watch?v=jxfq9_L9T1U&t=51s)
- [Programming Models for Distributed Computing](http://heather.miller.am/teaching/cs7680/)
- [Logic and Lattices for Distributed Programming](http://db.cs.berkeley.edu/papers/UCB-lattice-tr.pdf)

### Verification of Distributed Systems
- [Jepsen](https://github.com/jepsen-io/jepsen) A framework for distributed systems verification, with fault injection
  @aphyr has featured enough times in this list already, but Jepsen and the blog posts that go with are a quintessntial addition to any distributed systems reading list.
- [Verdi](http://verdi.uwplse.org/) A Framework for Implementing and Formally Verifying Distributed Systems [Paper](http://verdi.uwplse.org/verdi.pdf)

## Videos
- [Distributed Deep Dive](https://www.ably.io/blog/introducing-distributed-deep-dive-interview-series-by-ably-realtime/) interview series by [Ably Relatime](https://ably.io).
- [Distributed Systems in One Lesson](https://www.youtube.com/watch?v=Y6Ev8GIlbxc&t=17s) Distributed Systems in One Lesson by Tim Berglund

## Courses
- [Reliable Distributed Algorithms, Part 1](https://www.edx.org/course/reliable-distributed-algorithms-part-1-kthx-id2203-1x-0), KTH Sweden
- [Reliable Distributed Algorithms, Part 2](https://www.edx.org/course/reliable-distributed-algorithms-part-2-kthx-id2203-2x), KTH Sweden
- [Cloud Computing Concepts](https://class.coursera.org/cloudcomputing-001), University of Illinois
- [CMU: Distributed Systems](http://www.cs.cmu.edu/~dga/15-440/F12/syllabus.html) in Go Programming Language
- [Software Defined Networking](https://www.coursera.org/course/sdn) , Georgia Tech.
- [ETH Zurich: Distributed Systems](http://dcg.ethz.ch/lectures/podc_allstars/)
- [ETH Zurich: Distributed Systems Part 2](http://dcg.ethz.ch/lectures/distsys), covers  Distributed control algorithms, communication models, fault-tolerance among other things. In particular fault tolerance issues (models, consensus, agreement) and replication issues (2PC,3PC, Paxos), which are critical in understanding distributed systems are explained in great detail.
- [Distributed Systems Course](http://www.distributedsystemscourse.com/), A beginner course on distributed system by Chris Colohan, A google employee who contributed to  SUIF, MapReduce, TCMalloc, Percolator, Caffeine, Borg, Omega, and Piper.


## Blogs and other reading links
- [Amazon Builder's Library](https://aws.amazon.com/builders-library/), a collection of Amazon's learnings on distributed systems
- [How we implemented consistent hashing efficiently](https://blog.ably.io/how-to-implement-consistent-hashing-efficiently-fe038d59fff2)
- [Notes on Distributed Systems for Young Bloods](http://www.somethingsimilar.com/2013/01/14/notes-on-distributed-systems-for-young-bloods/)
- [High Scalability](http://highscalability.com/) Several architectures of huge internet services, for eg [twitter](http://highscalability.com/blog/2013/7/8/the-architecture-twitter-uses-to-deal-with-150m-active-users.html), [whatsapp](http://highscalability.com/blog/2014/2/26/the-whatsapp-architecture-facebook-bought-for-19-billion.html)
- [There is No Now](http://queue.acm.org/detail.cfm?id=2745385), Problems with simultaneity in distributed systems
- [Turing Lecture: The Computer Science of Concurrency: The Early Years](http://cacm.acm.org/magazines/2015/6/187316-turing-lecture-the-computer-science-of-concurrency/fulltext), An article by Leslie Lamport on concurrency
- [The Paper Trail](http://the-paper-trail.org/blog/tag/distributed-systems/) blog, a very readable blog covering various aspects of distributed systems
- [aphyr](https://aphyr.com/tags/Distributed-Systems), Posts on [jepsen](https://github.com/aphyr/jepsen) series are pretty awesome
- [All Things Distributed](http://www.allthingsdistributed.com/) - Wernel Vogel's (Amazon CTO) blog on distributed systems 
- [Distributed Systems: Take Responsibility for Failover](http://ivolo.me/distributed-systems-take-responsibility-for-failover/)
- [The C10K problem](http://www.kegel.com/c10k.html)
- [On Designing and Deploying Internet-Scale Services](http://static.usenix.org/event/lisa07/tech/full_papers/hamilton/hamilton_html/)
- [Files are hard](http://danluu.com/file-consistency/) A blog post on filesystem consistency, pretty important to read if you are into distributed storage or databases.
- [Distributed Systems Testing: The Lost World](http://tagide.com/blog/research/distributed-systems-testing-the-lost-world/) Testing distributed systems are hard enough, a well researched blog post which again covers a lot of links to various approaches and other papers
- [SWIM Protocol explained](https://asafdav2.github.io/2017/swim-protocol/) A blog post on popular SWIM failure detector


## Meta Lists
Other lists like this one
- [Readings in distributed systems](http://christophermeiklejohn.com/distributed/systems/2013/07/12/readings-in-distributed-systems.html)
- [Distributed Systems meta list](https://gist.github.com/macintux/6227368)
- [List of required readings for Distributed Systems](http://www.andrew.cmu.edu/course/15-749/READINGS/required/) Part of CMU's Engineering Distributed Systems course
- [The Distributed Reader](http://reiddraper.github.io/distreader/)
- [A Distributed Systems Reading List](https://dancres.github.io/Pages/), A collection of material, mostly papers on Distributed Systems Theory as well as seminal industry papers 
- [Distributed Systems Readings](https://henryr.github.io/distributed-systems-readings/), A comprehensive list of online courses related to distributed systems 
- [Awesome Distributed Consensus](https://github.com/dgryski/awesome-consensus), Another list of materials on distributed consensus protocols
- [Distributed Systems Journey](https://github.com/ericdouglas/distributed-systems-journey) - 📚 Roadmap to becoming a distributed systems specialist 🎓

## Organization Design / Team Dynamics

- [How Do Committees Invent?](http://www.melconway.com/Home/pdf/committees.pdf) :small_orange_diamond:<sup>PDF</sup> - Melvin E. Conway, Datamation magazine 1968. The original article defining Conway's Law.
- [Service per Team](https://microservices.io/patterns/decomposition/service-per-team.html) - Each team is responsible for one or more business functions (e.g. business capabilities). A team owns a code base consisting of one or more modules. Its code base is sized so as to not exceed the cognitive capacity of team. The team deploys its code as one or more services. A team should have exactly one service unless there is a proven need to have multiple services.
- [Start with Team Cognitive Load - Team Topologies](https://www.youtube.com/watch?v=haejb5rzKsM) :small_red_triangle:<sup>YT</sup> - DOES19 London. The "monoliths vs microservices" debate often focuses on technological aspects, ignoring strategy and team dynamics. Instead of technology, smart-thinking organizations are beginning with team cognitive load as the guiding principle for modern software. In this talk, we explain how and why, illustrated by real case studies.

## Real Life Stories

- [Clean microservice architecture](http://blog.cleancoder.com/uncle-bob/2014/10/01/CleanMicroserviceArchitecture.html)
- [Failing at microservices](https://rclayton.silvrback.com/failing-at-microservices)
- [How to talk to your friends about microservices](https://blog.pivotal.io/labs/labs/how-to-talk-to-your-friends-about-microservices)
- [How we build microservices at Karma](https://blog.yourkarma.com/building-microservices-at-karma)
- [How we ended up with microservices at SoundCloud](http://philcalcado.com/2015/09/08/how_we_ended_up_with_microservices.html)
- [Microservices: lessons from the frontline](https://www.thoughtworks.com/insights/blog/microservices-lessons-frontline)
- [Monolith first](http://martinfowler.com/bliki/MonolithFirst.html)
- [Scaling microservices at Gilt with Scala, Docker and AWS](http://www.infoq.com/news/2015/04/scaling-microservices-gilt)
- [From a Monolithic Big Data System to a Microservices Event-Driven Architecture: Challenges and Lessons Learned](https://www.researchgate.net/publication/341822188_From_a_Monolithic_Big_Data_System_to_a_Microservices_Event-Driven_Architecture)

## Enterprise & Verticals

- [Commercetools ![c]](https://commercetools.com/) - Headless commerce platform.
- [Flamingo](https://www.flamingo.me/) - Framework to build flexible and modern e-commerce applications.
- [Interact ![c]](https://interact.io/) - CRM microservices for rapid delivery of tailored solutions.
- [Moltin ![c]](https://moltin.com/) - E-commerce API for developers.
- [Predix ![c]](https://predix.io/) - Industrial microservices platform.
- [Skava ![c]](https://www.skava.com/) - Provides microservices for all the functions of your store, and the glue to hold them together.

## Theory

### Articles & Papers

- [A sidecar for your service mesh](https://www.abhishek-tiwari.com/a-sidecar-for-your-service-mesh/) - A short service mesh introduction.
- [AKF Scale Cube](http://akfpartners.com/techblog/2008/05/08/splitting-applications-or-services-for-scale/) - Model depicting the dimensions to scale a service.
- [Benchmark Requirements for Microservices Architecture Research](https://goo.gl/14ElnV) :small_orange_diamond:<sup>PDF</sup> - Set of requirements that may be useful in selecting a community-owned architecture benchmark to support repeatable microservices research.
- [Building Microservices? Here is What You Should Know](https://cloudncode.blog/2016/07/22/msa-getting-started/) - A practical overview, based on real-world experience, of what one would need to know in order to build microservices.
- [CALM](http://db.cs.berkeley.edu/papers/cidr11-bloom.pdf) :small_orange_diamond:<sup>PDF</sup> - Consistency as logical monotonicity.
- [Canary Release](http://martinfowler.com/bliki/CanaryRelease.html) - Technique to reduce the risk of introducing a new software version in production by slowly rolling out the change to a small subset of users before rolling it out to the entire infrastructure and making it available to everybody.
- [CAP Theorem](http://blog.thislongrun.com/2015/03/the-cap-theorem-series.html) -  States that it is impossible for a distributed computer system to simultaneously provide all three of the following guarantees: Consistency, Availability and Partition tolerance.
- [Formal Foundations of Serverless Computing](https://arxiv.org/pdf/1902.05870.pdf) :small_orange_diamond:<sup>PDF</sup> - The serverless computing abstraction exposes several low-level operational details that make it hard for programmers to write and reason about their code. This paper sheds light on this problem by presenting λ, an operational semantics of the essence of serverless computing.
- [Introducing Domain-Oriented Microservice Architecture](https://eng.uber.com/microservice-architecture/) - Introduction to Uber Engineering generalized approach to microservice architectures, named “Domain-Oriented Microservice Architecture” (DOMA).
- [Java Microservices: A Practical Guide](https://www.marcobehler.com/guides/java-microservices-a-practical-guide) - You can use this guide to understand what Java microservices are, how you architect and build them. Also: A look at Java microservice libraries & common questions.
- [Microservice Architecture](http://martinfowler.com/articles/microservices.html) - Particular way of designing software applications as suites of independently deployable services.
- [Microservices – Please, don’t](http://basho.com/posts/technical/microservices-please-dont/) - Critical advice about some problems regarding a microservices approach.
- [Microservices RefCard](https://dzone.com/refcardz/getting-started-with-microservices) - Getting started with microservices.
- [Microservices Trade-Offs](http://martinfowler.com/articles/microservice-trade-offs.html) - Guide to ponder costs and benefits of the mircoservices architectural style.
- [Reactive Manifesto](http://www.reactivemanifesto.org/) - Reactive systems definition.
- [Reactive Streams](http://www.reactive-streams.org/) - Initiative to provide a standard for asynchronous stream processing with non-blocking back pressure.
- [ROCAS](http://resources.1060research.com/docs/2015/Resource-Oriented-Computing-Adaptive-Systems-ROCAS-1.2.pdf) :small_orange_diamond:<sup>PDF</sup> - Resource Oriented Computing for Adaptive Systems.
- [SECO](http://ceur-ws.org/Vol-746/IWSECO2011-6-DengYu.pdf) :small_orange_diamond:<sup>PDF</sup> - Understanding software ecosystems: a strategic modeling approach.
- [Service Discovery in a Microservice Architecture](https://www.nginx.com/blog/service-discovery-in-a-microservices-architecture/) - Overview of discovery and registration patterns.
- [Testing Strategies in a Microservice Architecture](http://martinfowler.com/articles/microservice-testing/) - Approaches for managing the additional testing complexity of multiple independently deployable components.
- [Your Server as a Function](http://monkey.org/~marius/funsrv.pdf) :small_orange_diamond:<sup>PDF</sup> - Describes three abstractions which combine to present a powerful programming model for building safe, modular, and efficient server software: Composable futures, services and filters.
- [Microservices - The Journey So Far and Challenges Ahead](https://ieeexplore.ieee.org/stamp/stamp.jsp?tp=&arnumber=8354433) :small_orange_diamond:<sup>PDF</sup> - Overview of the state of microservices in both industrial and academia.

### Talks

- [10 Tips For Failing Badly at Microservices](https://www.youtube.com/watch?v=X0tjziAQfNQ) :small_red_triangle:<sup>YT</sup> - A presentation at Voxxed Days by David Schmitz.
- [Bla Bla Microservices Bla Bla](http://jonasboner.com/bla-bla-microservices-bla-bla/) - A talk at the O’Reilly Software Architecture Conference, April 2016.
- [Challenges in Implementing MicroServices](https://www.youtube.com/watch?v=yPf5MfOZPY0) :small_red_triangle:<sup>YT</sup> - A presentation at GOTO 2015 by Fred George.
- [Mastering Chaos - A Netflix Guide to Microservices](https://www.youtube.com/watch?v=CZ3wIuvmHeM) :small_red_triangle:<sup>YT</sup> - A presentation at QCon 2016 by Josh Evans.
- [Microservices](https://www.youtube.com/watch?v=wgdBVIX9ifA) :small_red_triangle:<sup>YT</sup> - A presentation at GOTO Berlin 2014 by Martin Fowler.
- [Principles Of Microservices](https://www.youtube.com/watch?v=PFQnNFe27kU) :small_red_triangle:<sup>YT</sup> - A presentation at Devoxx Belgium by Sam Newman.

### Tutorials

- [Developing a RESTful Microservice in Python](http://www.skybert.net/python/developing-a-restful-micro-service-in-python/) - A story of how an aging Java project was replaced with a microservice built with Python and Flask.
- [Developing and Testing Microservices With Docker](http://mherman.org/blog/2017/04/18/developing-and-testing-microservices-with-docker) - An example of the processes involved in creating a simple Docker-packaged Node microservice.
- [Microservices without the Servers](https://aws.amazon.com/blogs/compute/microservices-without-the-servers/) - Step by step demo-driven talk about serverless architecture.
- Microservices in C#: [Part 1](http://insidethecpu.com/2015/07/17/microservices-in-c-part-1-building-and-testing/), [Part 2](http://insidethecpu.com/2015/07/31/microservices-in-c-part-2-consistent-message-delivery/), [Part 3](http://insidethecpu.com/2015/08/14/microservices-in-c-part-3-queue-pool-sizing/), [Part 4](http://insidethecpu.com/2015/08/28/microservices-in-c-part-4-scaling-out/), [Part 5](http://insidethecpu.com/2015/09/11/microservices-in-c-part-5-autoscaling/).
- [Microservices with Python, RabbitMQ and Nameko](http://brunorocha.org/python/microservices-with-python-rabbitmq-and-nameko.html)
- [Reactive Microservices](https://github.com/theiterators/reactive-microservices) - Project showcasing different microservice communication styles using Scala, Akka, Play and other tools from Scala ecosystem.
- [Using Packer and Ansible to build immutable infrastructure](https://blog.codeship.com/packer-ansible/)

### Books

- [Building Microservices](https://www.nginx.com/wp-content/uploads/2015/01/Building_Microservices_Nginx.pdf) :small_orange_diamond:<sup>PDF</sup> - Building Microservices: Designing Fine-grained Systems. Sam Newman. Preview Edition.
- [Istio in Action](https://www.manning.com/books/istio-in-action) - Teaches you how to implement a full-featured Istio-based service mesh to manage a microservices application.
- [Microservice Architecture: Aligning Principles, Practices, and Culture](http://shop.oreilly.com/product/0636920050308.do) - Practical advice for the strategy and design of Microservices.
- [Microservices in Action](https://www.manning.com/books/microservices-in-action) - A practical book about building and deploying microservice-based applications.
- [Microservices in .NET Core](https://www.manning.com/books/microservices-in-net-core-second-edition) - A comprehensive guide to building microservice systems using the .NET stack.
- [Microservice Patterns](https://www.manning.com/books/microservice-patterns) - Teaches how to build applications with the microservice architecture and how to refactor a monolithic application to a microservices.
- [Microservices from Theory to Practice](http://www.redbooks.ibm.com/abstracts/sg248275.html?Open) - Microservices from Theory to Practice: Creating Applications in IBM Bluemix Using the Microservices Approach. IBM Redbooks publication.
- [Migrating to Cloud Native Application Architectures](http://pivotal.io/platform/migrating-to-cloud-native-application-architectures-ebook) - This O’Reilly report defines the unique characteristics of cloud native application architectures such as microservices and twelve-factor applications.
- [Pulsar in Action](https://www.manning.com/books/pulsar-in-action) - A practical book about developing microservice-based applications using Apache Pulsar by Manning Press.
- [Testing Microservices with Mountebank](https://www.manning.com/books/testing-microservices-with-mountebank) - Provides a testing strategy using mountebank for service virtualization, promoting independent releases of Microservices
- [The Art of Scalability](http://theartofscalability.com/) - The Art of Scalability: Scalable Web Architecture, Processes, and Organizations for the Modern Enterprise. Martin L. Abbott, Michael T. Fisher.
- [The New Stack eBook Series](http://thenewstack.io/ebookseries/) - A Comprehensive Overview of the Docker and Container Ecosystem.
  + Book 1: The Docker Container Ecosystem.
  + Book 2: Applications & Microservices with Docker & Containers.
  + Book 3: Automation & Orchestration with Docker & Containers.
  + Book 4: Network, Security & Storage with Docker & Containers.
  + Book 5: Monitoring & Management with Docker & Containers.
- [The Tao of Microservices](https://www.manning.com/books/the-tao-of-microservices) - Teaches the path to understanding how to apply microservices architecture with your own real-world projects.
- [Micro Frontends in Action](https://www.manning.com/books/micro-frontends-in-action) - A practical guide that teaches how to develop large software projects with multiple independent teams.

### Sites & Organizations

- [Cloud Native Computing Foundation](https://www.cncf.io/) - The Cloud Native Computing Foundation builds sustainable ecosystems and fosters a community around a constellation of high-quality projects that orchestrate containers as part of a microservices architecture.
- [CNCF Cloud Native Interactive Landscape](https://landscape.cncf.io/) - Interactive landscape of cloud native technologies.
- [Microservices Resource Guide](http://martinfowler.com/microservices/) - Martin Fowler's choice of articles, videos, books, and podcasts that can teach you more about the microservices architectural style.
- [Microservice Patterns](http://microservices.io/) - Microservice architecture patterns and best practices.
- [Microservice Antipatterns and Pitfalls](https://www.oreilly.com/ideas/microservices-antipatterns-and-pitfalls) - Microservice mostly known antipatterns and pitfalls.

## Emerging Technologies

- [Enso](https://enso.org/) - Visual and textual functional programming language with a focus on productivity, collaboration and development ergonomics.
- [Holochain](https://holochain.org/) - A framework for distributed applications, allowing you to build apps without any network constraints. This means every user controls their own data, and it can't be sold or exposed to third parties.
- [Ops](https://ops.city/) - Free open source tool to build, run and deploy existing linux applications as unikernels.
- [SAFE Network](https://safenetwork.tech/) - Powered by the spare capacity of everyday computers, SAFE replaces the vulnerable structure of the existing Web with a decentralised, autonomous network. One that is secure, and accessible to everyone.
- [Solid](https://solid.inrupt.com/) - Empowers users and organizations to separate their data from the applications that use it. It allows people to look at the same data with different apps at the same time. It opens brand new avenues for creativity, problem-solving, and commerce.