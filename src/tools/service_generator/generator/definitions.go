package generator


const blackspaceRepo string = "github.com/blackspaceInc/BlackspacePlatform"
const templateRepoUrl string = "https://github.com/stefanprodan/podinfo.git"
const repoName string = "podinfo"
const EMPTY string = ""

var podInfoStrToRemove = []string{
	"podinfo", "Podinfo","PODINFO",
}

var fileMap = map[string]string{
	"./generator/templates/database": "pkg/database",
	"./generator/templates/logging": "pkg/logging",
	"./generator/templates/metrics": "pkg/metrics",
	"./generator/templates/middleware": "pkg/middleware",
	"./generator/templates/databaseModels": "pkg/database_models",
	"./generator/templates/models": "models",
	"./generator/templates/alerts" : "alerts",
	"./generator/templates/telemetry" : "monitoring",
	"./generator/templates/docker-compose.monitoring.yaml" : "docker-compose.monitoring.yaml",
	"./generator/templates/docker-compose.prod.yaml" : "docker-compose.prod.yaml",
	"./generator/templates/docker-compose.dev.yaml" : "docker-compose.dev.yaml",
	"./generator/templates/Dockerfile.dev" : "Dockerfile.dev",
	"./generator/templates/graphql" : "graphql",
}

