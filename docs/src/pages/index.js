import React from 'react';

import classnames from 'classnames';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import useBaseUrl from '@docusaurus/useBaseUrl';
import {Feature} from '../components/feature';
import styles from './styles.module.css';


const publish_github_samsung_public  = process.env.PUBLISH === 'github_samsung_public' 
const is_for_webapp = !publish_github_samsung_public


const features = [
  {
    title: <>Who Are We?</>,
    imageUrl: 'img/undraw/undraw_forming_ideas_0pav.svg',
    description: (
      <>
        Blackspace connects consumers to minority owned business and further empowers businesses do more.
      </>
    ),
  },
  {
    title: <>What does it do?</>,
    imageUrl: 'img/undraw/undraw_ideation_2a64.svg',
    description: (
      <>
        We provide a platform through which consumers can engage with minority owned business from both a social and ecommerce context. Additionally, we provide businesses with a suite of B2B offerings ultimately empowering companies to better serve the customers
      </>
    ),
  },
  {
    title: <>How do I use it?</>,
    imageUrl: 'img/undraw/undraw_factory_dy0a.svg',
    description: (
      <>
        Go To Blackspace.com and get started.
      </>
    ),
  },
];

function Home() {
  const context = useDocusaurusContext();
  const { siteConfig = {} } = context;

  return (
    <Layout
      description={siteConfig.tagline.replace("<br/>", " ")}
    >
      <header className={classnames('hero hero--primary', styles.heroBanner)}>
        <div className="container">
          <h1 className="hero__title">{siteConfig.title}</h1>
          <p className="hero__subtitle" dangerouslySetInnerHTML={{__html: siteConfig.tagline}}></p>
          <div className={styles.buttons}>
            <Link
              className={classnames(
                'button button--secondary button--lg',
                styles.getStarted,
              )}
              to={useBaseUrl('docs/introduction')}
            >
              Get Started
            </Link>
            <span className={styles.indexCtasGitHubButtonWrapper}>
              <iframe
                className={styles.indexCtasGitHubButton}
                src="https://ghbtns.com/github-btn.html?user=Samsung&amp;repo=qaboard&amp;type=star&amp;count=true&amp;size=large"
                width={160}
                height={30}
                title="GitHub Stars"
              />
            </span>
          </div>
        </div>
      </header>
      <main>
        {features && features.length && (
          <section className={styles.features}>
            <div className="container">
              <div className="row">
                {features.map(({ imageUrl, title, description }, idx) => (
                  <div
                    key={idx}
                    className={classnames('col col--4', styles.feature)}
                  >
                    {imageUrl && (
                      <div className="text--center">
                        <img
                          className={styles.featureImage}
                          src={useBaseUrl(imageUrl)}
                          alt={title}
                        />
                      </div>
                    )}
                    <h3>{title}</h3>
                    <p>{description}</p>
                  </div>
                ))}
              </div>
            </div>
          </section>
        )}

      <div className="container">
      <Feature
        img={<img src={useBaseUrl('img/custom/market.gif')} alt="Annotations" loading="lazy" />}
        title="Blackspace Marketplace"
        text={
          <>
            <p>
              Fill This In Later  <strong>no </strong>, yes<strong> maybe </strong>.
            </p>
            <p>
              Fill This In Later ....
            </p>
          </>
        }
      />
      <Feature
        img={<img src={useBaseUrl('img/custom/business.gif')} alt="Always Compare" loading="lazy" />}
        title="Blackspace Business"
        reversed
        text={
          <>
            <p>
            Fill This In Later <a href="https://samsung.github.io/qaboard/docs/references-and-milestones">FIll</a>.
            </p>
            <p>
              Fill This In Later
            </p>
          </>
        }
      />
      <Feature
        img={<img src={useBaseUrl('img/custom/analytics.gif')} alt="Aggregation and rich KPIs" loading="lazy" />}
        title="Blackspace Analytics"
        text={
          <>
            <p>
            Fill This In Later...
            </p>
            <p>
              Fill This <a href="https://samsung.github.io/qaboard/docs/computing-quantitative-metrics">in </a> later.
            </p>
          </>
        }
      />
      <Feature
        img={<img src={useBaseUrl('img/custom/creator.gif')} alt="File-based Visualizations" loading="lazy" />}
        title="Blackspace Makers"
        reversed
        text={
          <>
            <p>
              Fill This In Later
            </p>
            <p>
              Fill <a href="https://samsung.github.io/qaboard/docs/visualizations">This</a> In Later....
            </p>
          </>
        }
      />
      <Feature
        img={<img src={useBaseUrl('img/custom/integration.gif')} alt="Advanced Image Viewer" loading="lazy" />}
        title="Blackspace 3rd Party Integrations"
        text={
          <>
            <p>
              Fill This In <a href="https://openseadragon.github.io/">Later</a>. 
            </p>
            <p>
            Fill This In Later
            </p>
          </>
        }
      />
      <Feature
        img={<img src={useBaseUrl('img/custom/ads.gif')} alt="Tuning & Optimization" loading="lazy" />}
        title="Blackspace Ads"
        reversed
        text={
          <>
            <p>
              <a href="https://samsung.github.io/qaboard/docs/batches-running-on-multiple-inputs">Fill This In</a> Later 
            </p>
            <p>
              Fill <strong>This</strong> (via <a href="https://scikit-optimize.github.io/">In</a>), Later <a href="https://github.com/Samsung/qaboard/wiki/Adding-new-runners"> yes</a>.
            </p>
          </>
        }
      />
      <Feature
        img={<img src={useBaseUrl('img/custom/AI.gif')} alt="Integrations" loading="lazy" />}
        title="Blackspace AI & Research"
        text={
          <>
            <p>
              Fill This In Later
            </p>
          </>
        }
      />
      <Feature
        img={<img src={useBaseUrl('img/custom/money.gif')} alt="Regression Explorer" loading="lazy" />}
        title="Blackspace VC (Fund Raising)"
        reversed
        text={
          <>
            <p>
              Fill This In Later
            </p>
            <p>
              Fill This In Later
            </p>
          </>
        }
      />
      <Feature
        title="More features..."
        text={
          <>
            <p>
           
            </p>
          </>
        }
      />
    </div>


      {/* <div className="container">
        <div className="row">
          <div className="col col--6 col--offset-3 padding-vert--lg">
            <h2>Introduction Video</h2>
            <iframe
              width="100%"
              height="315"
              src="https://www.youtube.com/embed/nYkdrAPrdcw"
              frameborder="0"
              allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture"
              allowfullscreen
            />
            <div className="text--center padding-vert--lg">
              <Link
                className="button button--primary button--lg"
                to={useBaseUrl('docs/introduction')}
              >
                Learn more about QA-Board!
              </Link>
            </div>
          </div>
        </div>
      </div> */}
      </main>
    </Layout>
  );
}

export default Home;
