let guidelines

if (process.env.NODE_ENV === "development" || process.env.NODE_ENV === "production" || process.env.NODE_ENV === "prod") {
  guidelines = {
    label: "Guidelines (DEV ONLY)",
    type: "category",
    items: [
      {
        type: "category",
        label: "Templates",
        items: [
          "__guidelines/template/guide",
          "__guidelines/template/function",
          "__guidelines/template/sql",
        ],
      },
      "__guidelines/naming-convention",
      "__guidelines/content-hierarchy",
      "__guidelines/lexicon",
      "__guidelines/markdown",
      "__guidelines/influences",
    ],
  }
}

module.exports = {
  docs: [
    guidelines,
    {
      id: "introduction",
      type: "doc",
    },
    {
        label: "Learning",
        type: "category",
        items: [
          {
            type: "category",
            label: "Distributed Systems",
            items: [
              "learning/distributed-systems/big-data",
              "learning/distributed-systems/distributed-systems",
              "learning/distributed-systems/patterns",
              "learning/scalability/scalability",
            ],
          },
        ],
    },
    {
        label: "Checklists",
        type: "category",
        items: [
          {
            type: "category",
            label: "Backend",
            items: [
              "checklists/backend/backend-kubernetes/application-development",
              "checklists/backend/backend-kubernetes/best-practices",
              "checklists/backend/backend-kubernetes/cluster-configuration",
              "checklists/backend/backend-kubernetes/common",
              "checklists/backend/backend-kubernetes/governance",
              "checklists/backend/backend-kubernetes/patterns",
              "checklists/backend/backend-production-ready/production-ready-checklist",
              "checklists/backend/backend-security-checklist/production-security-checklist"
            ],
          },
          {
            type: "category",
            label: "Frontend",
            items: [
              "checklists/frontend/frontend-design-checklist/frontend-design-checklist",
              "checklists/frontend/frontend-performance-checklist/frontend-performance-checklist",
              "checklists/frontend/frontend-prod-checklist/frontend-prod-checklist",
            ],
          },
        ],
    },
    {
        label: "Technology",
        type: "category",
        items: [
          {
            type: "category",
            label: "Architecture",
            items: [
              "technology/architecture/instrumentation/instrumentation",
              "technology/architecture/metrics/metrics",
              "technology/architecture/service-architecture/service-architecture",
              "technology/architecture/service-mesh/service-mesh",
            ],
          },
          {
            type: "category",
            label: "Backend Services",
            items: [
              "technology/backend-services/authentication-service/authentication-service",
              "technology/backend-services/company-service/company-service",
              "technology/backend-services/email-service/email-service",
              "technology/backend-services/user-service/user-service"
            ],
          },
          {
            type: "category",
            label: "Deployments",
            items: [
              "technology/deployments/deployments",
            ],
          },
          {
            type: "category",
            label: "System Benchmarks",
            items: [
              "technology/system-benchmarks/system-benchmark",
            ],
          },
          {
            type: "category",
            label: "Third Party API",
            items: [
              "technology/third-party-apis/third-party-apis",
            ],
          },
        ],
    },
    {
      label: "Guides",
      type: "category",
      items: [
        {
          type: "category",
          label: "Setting up Blackspace",
          items: ["guide/docker", "guide/kubernetes"],
        },
      ],
    },
    {
      label: "Roadmap",
      type: "category",
      items: [ 
        {
        type: "category",
        label: "Product Roadmaps",
        items: [
          "roadmap/blackspace-ecommerce/ecommerce-roadmap",
        ],
      },
    ],
    },
    {
      label: "Products",
      type: "category",
      items: [
        {
          type: "category",
          label: "Features",
          items: [
            "products/ads/ads",
            "products/analytics/analytics",
            "products/business/business",
            "products/integration/integration",
            "products/makers/makers",
            "products/marketplace/marketplace",
            "products/research/research",
            "products/vc/vc",
            "products/social/social"
          ],
        },
      ],
  },
  ].filter(Boolean),
}
