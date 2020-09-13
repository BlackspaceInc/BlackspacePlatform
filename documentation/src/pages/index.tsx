import clsx from "clsx"
import DocusaurusHead from "@docusaurus/Head"
import useDocusaurusContext from "@docusaurus/useDocusaurusContext"
import React, { useCallback, useEffect, useState } from "react"

import CodeBlock from "@theme/CodeBlock"
import Layout from "@theme/Layout"

import Button from "@theme/Button"
import { MetadataContextProvider } from "@theme/useMetadataContext"
import useWindowWidth from "@theme/useWindowWidth"

import cardsStyles from "../css/cards.module.css"
import consoleStyles from "../css/console.module.css"
import featureStyles from "../css/feature.module.css"
import jumbotronStyles from "../css/jumbotron.module.css"
import queryScrollerStyles from "../css/queryScroller.module.css"
import sectionStyles from "../css/section.module.css"
import seenOnStyles from "../css/seenOn.module.css"
import whyStyles from "../css/why.module.css"

const Why = () => {
  // eslint-disable-next-line prettier/prettier
  const [opened, setOpened] = useState<"digital" | "realtime" | "integration" | "business">(
    // eslint-disable-next-line prettier/prettier
    "digital",
  )
  const handleClickIs = useCallback(() => {
    setOpened("digital")
  }, [])
  const handleClickGoodFor = useCallback(() => {
    setOpened("realtime")
  }, [])
  const handleClickIsNot = useCallback(() => {
    setOpened("integration")
  }, [])
  const handleClickIsBusiness = useCallback(() => {
    setOpened("business")
  }, [])

  return (
    <section
      className={clsx(sectionStyles.section, sectionStyles["section--odd"])}
    >
      <div className={clsx(sectionStyles["section--inner"], whyStyles.why)}>
        <h2
          className={clsx(
            sectionStyles.section__title,
            whyStyles.why__title,
            "text--center",
          )}
        >
          Why Blackspace?
        </h2>

        <div className={whyStyles.why__footer}>
          <div className={whyStyles.why__menu}>
            <Button
              className={whyStyles["why__menu--item"]}
              onClick={handleClickIs}
              size="small"
              variant={opened === "digital" ? "primary" : "tertiary"}
            >
              Real-time insights & Analytics
            </Button>
            <Button
              className={whyStyles["why__menu--item"]}
              onClick={handleClickGoodFor}
              size="small"
              variant={opened === "realtime" ? "primary" : "tertiary"}
            >
              MarketPlace & E-Commerce
            </Button>
            <Button
              className={whyStyles["why__menu--item"]}
              onClick={handleClickIsNot}
              size="small"
              variant={opened === "integration" ? "primary" : "tertiary"}
            >
              Businesse & Enterprise integration
            </Button>
            <Button
              className={whyStyles["why__menu--item"]}
              onClick={handleClickIsBusiness}
              size="small"
              variant={opened === "business" ? "primary" : "tertiary"}
            >
              Extensions & Connected Services
            </Button>
          </div>

          <div className={whyStyles.why__content}>
            <div
              className={clsx(whyStyles.why__toggle, {
                [whyStyles["why__toggle--active"]]: opened === "digital",
              })}
            >
              <p className={whyStyles.why__item}>Site Analytics</p>
              <p className={whyStyles.why__item}>KPI Monitoring</p>
              <p className={whyStyles.why__item}>Traffic & Sales Overview</p>
              <p className={whyStyles.why__item}>
                Visitor Insights & Analytics
              </p>
              <p className={whyStyles.why__item}>Page Views</p>
              <p className={whyStyles.why__item}>Purchase Funnel</p>
            </div>

            <div
              className={clsx(whyStyles.why__toggle, {
                [whyStyles["why__toggle--active"]]: opened === "realtime",
              })}
            >
              <p className={whyStyles.why__item}>Multiple Payment Methods</p>
              <p className={whyStyles.why__item}>Product Merchandising</p>
              <p className={whyStyles.why__item}>Product Import</p>
              <p className={whyStyles.why__item}>Optimized Checkout</p>
              <p className={whyStyles.why__item}>Inventory Management</p>
              <p className={whyStyles.why__item}>Flexible Shipping Options</p>
            </div>

            <div
              className={clsx(whyStyles.why__toggle, {
                [whyStyles["why__toggle--active"]]: opened === "business",
              })}
            >
              <p className={whyStyles.why__item}>Business Product Analytics</p>
              <p className={whyStyles.why__item}>
                Powerful Customizable Infrastructure
              </p>
              <p className={whyStyles.why__item}>High availability</p>
              <p className={whyStyles.why__item}>Privacy Protected</p>
              <p className={whyStyles.why__item}>Enterprise security</p>
              <p className={whyStyles.why__item}>Marketing Campaigns & Ads </p>
            </div>

            <div
              className={clsx(whyStyles.why__toggle, {
                [whyStyles["why__toggle--active"]]: opened === "integration",
              })}
            >
              <p className={whyStyles.why__item}>Finance & Tax Extensions</p>
              <p className={whyStyles.why__item}>
                Marketing & Sales Extensions
              </p>
              <p className={whyStyles.why__item}>Shipping & Sales Extensions</p>
              <p className={whyStyles.why__item}>
                Inventory & Product Extensions
              </p>
              <p className={whyStyles.why__item}>Social Integrations</p>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}

const SeenOn = () => (
  <section
    className={clsx(
      sectionStyles.section,
      sectionStyles["section--inner"],
      seenOnStyles.section,
    )}
  />
)

const Jumbotron = () => {
  const { siteConfig } = useDocusaurusContext()

  return (
    <section
      className={clsx(
        sectionStyles["section--inner"],
        jumbotronStyles.jumbotron,
      )}
    >
      <div className={jumbotronStyles.jumbotron__left}>
        <h1 className={clsx("jumbotron", jumbotronStyles.jumbotron__title)}>
          Blackspace Platform
        </h1>
        <p className={clsx("jumbotron", jumbotronStyles.jumbotron__subtitle)}>
          {siteConfig.tagline}
        </p>
        <div className={jumbotronStyles.jumbotron__cta}>
          <Button
            className={jumbotronStyles.jumbotron__link}
            href={`http://try.${siteConfig.customFields.domain}:9000`}
          >
            Live Demo
          </Button>
          <Button
            className={clsx(
              jumbotronStyles.jumbotron__link,
              jumbotronStyles["jumbotron__cta--github"],
            )}
            href={siteConfig.customFields.githubUrl}
            icon={
              <img
                alt="GitHub logo"
                height="26"
                src="/img/github.svg"
                title="GitHub"
                width="26"
              />
            }
            variant="secondary"
          >
            GitHub
          </Button>
        </div>
        <p className={jumbotronStyles.jumbotron__description}>
          Access our live demo.
        </p>
      </div>

      <div className={jumbotronStyles.jumbotron__right}>
        <pre className={jumbotronStyles.jumbotron__docker}>
          <code>
            {`git clone blackspaceInc/BlackspacePlatform
cd BlackspacePlatform
make clean
make up`}
          </code>
          <a
            href={siteConfig.customFields.dockerUrl}
            rel="noopener noreferrer"
            target="_blank"
          >
            <img
              alt="Docker logo"
              className={jumbotronStyles["jumbotron__docker-icon"]}
              src="/img/pages/index/docker.svg"
              title="Docker"
              width="60"
            />
          </a>
        </pre>
      </div>
    </section>
  )
}

const Feature = () => (
  <section
    className={clsx(sectionStyles.section, sectionStyles["section--odd"])}
  >
    <div className={sectionStyles["section--inner"]}>
      <div className={featureStyles.feature}>
        <img
          alt="Speedometer"
          className={featureStyles.feature__illustration}
          src="/img/pages/index/rawPower.svg"
        />

        <h2 className={featureStyles.feature__title}>Built for scale</h2>

        <p className={featureStyles.feature__description}>
          Low Latency High Throughput System
        </p>
        <p className={featureStyles.feature__description}>
          Privacy Focused Platform
        </p>
        <p className={featureStyles.feature__description}>
          Enforced Atomicity & Consistency
        </p>
        <p className={featureStyles.feature__description}>
          Distributed Large Scale System
        </p>
        <p className={featureStyles.feature__description}>
          Built in Typescript & Golang
        </p>
      </div>

      <div className={featureStyles.feature}>
        <img
          alt="A code editor with a chart that shows the result of the query"
          className={featureStyles.feature__illustration}
          src="/img/pages/index/easyToUse.svg"
        />

        <h2 className={featureStyles.feature__title}>
          Optimized To Serve Businesses
        </h2>
        <p className={featureStyles.feature__description}>
          Rich Business Analytics
        </p>
        <p className={featureStyles.feature__description}>
          Granular Inference On Product Performance
        </p>
        <p className={featureStyles.feature__description}>
          Historical Analysis Of Buy Patterns
        </p>
        <p className={featureStyles.feature__description}>
          Immediate Consistency & Real Time Insights
        </p>
      </div>

      <div className={featureStyles.feature}>
        <img
          alt="A code editor containing a SQL statement"
          className={featureStyles.feature__illustration}
          src="/img/pages/index/featureRich.svg"
        />

        <h2 className={featureStyles.feature__title}>Social Commerce</h2>

        <p className={featureStyles.feature__description}>
          Community Based Shopping Experiences
        </p>
        <p className={featureStyles.feature__description}>
          Community Driven Reviews
        </p>
        <p className={featureStyles.feature__description}>
          Intimate Feedback Loop
        </p>
        <p className={featureStyles.feature__description}>
          Hyper Connected Customer Base
        </p>
      </div>
    </div>
  </section>
)

const Cards = () => (
  <section
    className={clsx(
      sectionStyles.section,
      sectionStyles["section--inner"],
      cardsStyles.cards,
    )}
  >
    <h3
      className={clsx(
        sectionStyles.section__title,
        cardsStyles.cards__title,
        "text--center",
      )}
    >
      Use Cases
    </h3>

    <div className={cardsStyles.cards__footer}>
      <div className={cardsStyles.cards__wrapper}>
        <h3 className={cardsStyles.cards__header}>Transaction monitoring</h3>
        <p className={cardsStyles.cards__content}>
          Collect metrics and events from your account (Sales, Cart Behaviors,
          Page Views, etc) and get real-time visibility into your entire product
          suite.
        </p>
      </div>

      <div className={cardsStyles.cards__wrapper}>
        <h3 className={cardsStyles.cards__header}>Transactional data</h3>
        <p className={cardsStyles.cards__content}>
          Store Transactional data to identify historical trends and
          correlations using statistical methods and generate better product
          marketing decisions.
        </p>
      </div>

      <div className={cardsStyles.cards__wrapper}>
        <h3 className={cardsStyles.cards__header}>Connected Channels</h3>
        <p className={cardsStyles.cards__content}>
          Capture, store and respond to data from various channels at any
          resolution.
        </p>
      </div>

      <div className={cardsStyles.cards__wrapper}>
        <h3 className={cardsStyles.cards__header}>Application metrics</h3>
        <p className={cardsStyles.cards__content}>
          Empower business owners to track and visualise logs and any
          application activity in real-time.
        </p>
      </div>

      <div className={cardsStyles.cards__wrapper}>
        <h3 className={cardsStyles.cards__header}>Flexibility </h3>
        <p className={cardsStyles.cards__content}>
          Allows easy changes to your account and product suite
        </p>
      </div>

      <div className={cardsStyles.cards__wrapper}>
        <h3 className={cardsStyles.cards__header}>Integrated data</h3>
        <p className={cardsStyles.cards__content}>
          Pull together all your account, consumers, and transactional
          datapoints for a complete, 360º view of all aspects of your business.
        </p>
      </div>
    </div>
  </section>
)

const Console = () => {
  const { siteConfig } = useDocusaurusContext()

  return (
    <section
      className={clsx(sectionStyles.section, sectionStyles["section--odd"])}
    >
      <div
        className={clsx(sectionStyles["section--inner"], consoleStyles.console)}
      >
        <div className={consoleStyles.console__footer}>
          <div
            className={clsx(
              consoleStyles.console__highlight,
              consoleStyles["console__highlight--primary"],
            )}
          >
            <img
              alt="Postgres logo"
              src="/img/pages/index/pgwire.svg"
              title="Postgres"
            />
            <h3 className={consoleStyles.console__label}>Subscribe</h3>
            <p className={consoleStyles.console__summary}>
              Try our live demo and get ready for our launch
            </p>
          </div>

          <div className={consoleStyles.console__highlight}>
            <img
              alt="Antenna"
              src="/img/pages/index/foss.svg"
              title="Open source"
            />
            <h3 className={consoleStyles.console__label}>Blackspace</h3>
            <p className={consoleStyles.console__summary}>
              Follow us on GitHub. Watch the repo to get notified of further
              releases and new features!
            </p>

            <div className={consoleStyles.console__actions}>
              <a
                className={consoleStyles.console__link}
                href={siteConfig.customFields.githubUrl}
                rel="noopener noreferrer"
                target="_blank"
              >
                Go to GitHub&nbsp;&nbsp;&gt;
              </a>
              <a
                className={consoleStyles.console__link}
                href={siteConfig.customFields.slackUrl}
              >
                Join Slack&nbsp;&nbsp;&gt;
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}

const S = [3, 1, 6, 10]
const M = [3, 0, 4, 8]
const L = [4, 0, 4, 8]

const getTopByIndex = (m: number[], index: 1 | 2 | 3 | 4): number => {
  const scale = {
    1: 25 * (m[0] || 0),
    2: -25 * (m[1] || 0),
    3: -25 * (m[2] || 0),
    4: -25 * (m[3] || 0),
  }

  return scale[index] || 0
}

const searchQuery = `Sign Up/Login 
Custom Auth. Or
Sign Up Using OAUTH (Google/Github/...etc);`

const sliceQuery = `Search Product
Gain Insights & Learn
Find Sellers Of Interest`

const navigateQuery = `Connect To Sellers
Engage In Transactions;`

const mergeQuery = `Review Products & Sellers
Follow Seller Profiles
Engage Community Discussions;`

const Chevron = () => (
  <svg
    fill="currentColor"
    focusable="false"
    role="img"
    viewBox="5.40 7.12 9.23 5.25"
    width="26"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path d="M6.582 12.141a.695.695 0 01-.978 0 .68.68 0 010-.969l3.908-3.83a.697.697 0 01.979 0l3.908 3.83a.68.68 0 010 .969.697.697 0 01-.979 0L10 9l-3.418 3.141z" />
  </svg>
)

type Index = 1 | 2 | 3 | 4

const QueryScroller = () => {
  const [top, setTop] = useState(S)
  const [index, setIndex] = useState<Index>(2)
  const windowWidth = useWindowWidth()
  const handleClick1 = useCallback(() => {
    setIndex(1)
  }, [])
  const handleClick2 = useCallback(() => {
    setIndex(2)
  }, [])
  const handleClick3 = useCallback(() => {
    setIndex(3)
  }, [])
  const handleClick4 = useCallback(() => {
    setIndex(4)
  }, [])
  const handleUpClick = useCallback(() => {
    setIndex(Math.max(index - 1, 1) as Index)
  }, [index])
  const handleDownClick = useCallback(() => {
    setIndex(Math.min(index + 1, 4) as Index)
  }, [index])

  useEffect(() => {
    if (windowWidth && windowWidth < 622) {
      setTop(S)
      return
    }

    if (windowWidth && windowWidth < 800) {
      setTop(M)
      return
    }

    setTop(L)
  }, [windowWidth])

  return (
    <section
      className={clsx(
        sectionStyles["section--inner"],
        queryScrollerStyles.queryScroller,
      )}
    >
      <h2
        className={clsx(
          sectionStyles.section__title,
          queryScrollerStyles.queryScroller__title,
          "text--center",
        )}
      >
        Blackspace Customer E2E Buy Experience
      </h2>

      <p
        className={clsx(
          sectionStyles.section__subtitle,
          queryScrollerStyles.queryScroller__subtitle,
          "text--center",
        )}
      >
        Blackspace is attempting to revolutionize the shopping experience of
        consumers by both enabling them to better find businesses of interested
        as well as allowing such consumers to share ideas and shop in
        communities
      </p>

      <div className={queryScrollerStyles.queryScroller__scroller}>
        <div className={queryScrollerStyles.queryScroller__inner}>
          <div
            className={clsx(queryScrollerStyles.queryScroller__chevron)}
            onClick={handleUpClick}
            style={{ visibility: index === 1 ? "hidden" : "visible" }}
          >
            <Chevron />
          </div>
          <div className={clsx(queryScrollerStyles.queryScroller__left)}>
            <div
              className={clsx(
                queryScrollerStyles.queryScroller__offset,
                queryScrollerStyles[`queryScroller__${index}`],
              )}
              style={{ top: getTopByIndex(top, index) }}
            >
              <CodeBlock>{`${searchQuery}`}</CodeBlock>
              <CodeBlock>
                {`-- Onboard
${searchQuery}`}
              </CodeBlock>
              <CodeBlock>{`${sliceQuery}`}</CodeBlock>
              <CodeBlock>
                {`-- Search Products & Businesses
${sliceQuery}`}
              </CodeBlock>
              <CodeBlock>{`${navigateQuery}`}</CodeBlock>
              <CodeBlock>
                {`-- Connect & Buy
${navigateQuery}`}
              </CodeBlock>
              <CodeBlock>{`${mergeQuery}`}</CodeBlock>
              <CodeBlock>
                {`-- Review & Follow
${mergeQuery}`}
              </CodeBlock>
            </div>
          </div>
          <div
            className={clsx(
              queryScrollerStyles.queryScroller__chevron,
              queryScrollerStyles["queryScroller__chevron--bottom"],
            )}
            onClick={handleDownClick}
            style={{ visibility: index === 4 ? "hidden" : "visible" }}
          >
            <Chevron />
          </div>
          <div className={queryScrollerStyles.queryScroller__right}>
            <div
              className={clsx(queryScrollerStyles.queryScroller__button, {
                [queryScrollerStyles["queryScroller__button--active"]]:
                  index === 1,
              })}
              onClick={handleClick1}
            >
              <h3 className={queryScrollerStyles.queryScroller__header}>
                <img
                  alt="Magnifying glass icon"
                  className={queryScrollerStyles.queryScroller__icon}
                  src="/img/pages/index/searchTime.svg"
                />
                Onboard
              </h3>
              <p className={queryScrollerStyles.queryScroller__description}>
                Create a user account on the blackspace platform
              </p>
            </div>

            <div
              className={clsx(queryScrollerStyles.queryScroller__button, {
                [queryScrollerStyles["queryScroller__button--active"]]:
                  index === 2,
              })}
              onClick={handleClick2}
            >
              <h3 className={queryScrollerStyles.queryScroller__header}>
                <img
                  alt="Knife icon"
                  className={queryScrollerStyles.queryScroller__icon}
                  src="/img/pages/index/sliceTime.svg"
                />
                Search Products & Businesses
              </h3>
              <p className={queryScrollerStyles.queryScroller__description}>
                Query for businesses or products of interest & search associated
                engaging in dialogue about the given product/business
              </p>
            </div>

            <div
              className={clsx(queryScrollerStyles.queryScroller__button, {
                [queryScrollerStyles["queryScroller__button--active"]]:
                  index === 3,
              })}
              onClick={handleClick3}
            >
              <h3 className={queryScrollerStyles.queryScroller__header}>
                <img
                  alt="Indication arrow icon"
                  className={queryScrollerStyles.queryScroller__icon}
                  src="/img/pages/index/navigateTime.svg"
                />
                Connect & Buy
              </h3>
              <p className={queryScrollerStyles.queryScroller__description}>
                Connect with the business selling the product of interest if
                interested. Engage in a transaction that fits your needs.
              </p>
            </div>
            <div
              className={clsx(queryScrollerStyles.queryScroller__button, {
                [queryScrollerStyles["queryScroller__button--active"]]:
                  index === 4,
              })}
              onClick={handleClick4}
            >
              <h3 className={queryScrollerStyles.queryScroller__header}>
                <img
                  alt="Two overlapping squares"
                  className={queryScrollerStyles.queryScroller__icon}
                  src="/img/pages/index/mergeTime.svg"
                />
                Review & Follow
              </h3>
              <p className={queryScrollerStyles.queryScroller__description}>
                Continue engagement with the business both directly and
                indirectly through the various blackspace communities discussing
                the product you bought.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}

const Home = () => {
  const { siteConfig } = useDocusaurusContext()
  const title = "Blackspace"

  return (
    <MetadataContextProvider>
      <Layout description={siteConfig.customFields.description} title={title}>
        <DocusaurusHead>
          <link rel="canonical" href={siteConfig.url} />
        </DocusaurusHead>
        <Jumbotron />
        <Feature />
        <QueryScroller />
        <Why />
        <Cards />
        <Console />
        <SeenOn />
      </Layout>
    </MetadataContextProvider>
  )
}

export default Home
