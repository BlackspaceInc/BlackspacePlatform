// https://github.com/Samsung/qaboard/tree/master/website
module.exports = {
    docs: {
        "Getting Started" : [
            "introduction",
        ],
        "Learning Resources" : [
            {
                type: 'category',
                label: 'Backend Resources',
                items: [
                    "learning/distributed-systems/distributed-systems",
                    "learning/scalability/scalability",
                    "learning/distributed-systems/patterns",
                    "learning/distributed-systems/big-data"
                ],
            },
            
        ],
        "Technology": [
            {
                type: 'category',
                label: 'Architecture',
                items: [
                    "technology/architecture/service-mesh/service-mesh",
                    "technology/architecture/instrumentation/instrumentation",
                    "technology/architecture/backend-services/services",
                    "technology/architecture/metrics/metrics"
                ],
            },
            {
                type: 'category',
                label: 'Backend Services',
                items: [
                    "technology/backend-services/authentication-service/auth-service",
                    "technology/backend-services/email-service/email-service",
                    "technology/backend-services/company-service/company-service",
                    "technology/backend-services/user-service/user-service"
                ],
            },
            {
                type: 'category',
                label: 'Kubernetes',
                items: [
                    "technology/checklists/backend/backend-kubernetes/kubernetes-development",
                    "technology/checklists/backend/backend-kubernetes/kubernetes-standards",
                    "technology/checklists/backend/backend-kubernetes/kubernetes-cluster-config",
                    "technology/checklists/backend/backend-kubernetes/kubernetes-common",
                    "technology/checklists/backend/backend-kubernetes/kubernetes-governance",
                    "technology/checklists/backend/backend-kubernetes/kubernetes-patterns",
                ],
            },
            {
                type: 'category',
                label: 'Production Checklist',
                items: [
                    "technology/checklists/backend/backend-production-ready/production-ready-checklist",
                    "technology/checklists/backend/backend-security-checklist/production-security-checklist",
                ],
            },
            {
                type: 'category',
                label: 'Frontend Checklist',
                items: [
                    "technology/checklists/frontend/frontend-performance-checklist/frontend-performance-checklist",
                    "technology/checklists/frontend/frontend-design-checklist/frontend-design-checklist",
                    "technology/checklists/frontend/frontend-prod-checklist/frontend-prod-checklist"
                ],
            },
            "technology/deployments/deployment",
            "technology/system-benchmarks/system-benchmarks",
            {
                type: 'category',
                label: 'Product Roadmap',
                items: [
                    "technology/roadmap/blackspace-analytics/blackspace-analytics-roadmap",
                    "technology/roadmap/blackspace-ecommerce/blackspace-ecommerce-roadmap",
                    "technology/roadmap/blackspace-social/blackspace-social-roadmap",
                    "technology/roadmap/blackspace-roadmap",
                ],
            },
            "technology/third-party-apis/third-party-apis"
        ],
        /*
        "Getting Started": [
            "introduction",
            "deploy",
            "installation",
            "project-init",
            "inputs",
            "running-your-code",
            "creating-and-viewing-outputs-files",
            "computing-quantitative-metrics",
            "specifying-configurations",
        ],
        "Guides": [
            "visualizations",
            "batches-running-on-multiple-inputs",
            "using-the-qa-cli",
            "references-and-milestones",
            "tuning-from-the-webapp",
            {
                type: 'category',
                label: 'Distributed Task Queues',
                items: [
                    "celery-integration",
                    "lsf-integration",
                ],
            },
            {
                type: 'category',
                label: 'Storage & Artifacts',
                items: [
                    "storage/where-is-the-data",
                    "storage/artifacts",
                    "storage/deleting-old-data",
                ],
            },
            "triggering-third-party-tools",
            "ci-integration",
            "debugging-runs-with-an-IDE",
            "metadata-integration-external-databases",
            "apis",
            "tuning-workflows",
            "dag-pipelines",
            "bit-accuracy",
            "faq",
            // "history"
            // "monorepos-subprojects",
            // "docker-integration",
            // "remote-platforms",
        ],
        // "Parameter Tuning": [
        // 	"Tuning Workflows",
        // 	"Enabling Tuning from QA-Board", // Save artifacts..
        // 	"Tuning runners", // setup LSF and != LSF///
        //   "Auto-Tuning"
        // ],
        //      ""
        //   "Admin Guides": [
        // 	  "starting-server",
        // 	  "server-maintenance"
        //   ]
        // "alternatives",
        "Backend Admin": [
            "backend-admin/troubleshooting",
        
            "backend-admin/host-upgrades",
        ]
        */
    }
}

