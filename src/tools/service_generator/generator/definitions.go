package generator


const blackspaceRepo string = "github.com/BlackspaceInc/BlackspacePlatform"
const templateRepoUrl string = "https://github.com/stefanprodan/authentication_handler_service.git"
const repoName string = "authentication_handler_service"
const EMPTY string = ""

var podInfoStrToRemove = []string{
	"authentication_handler_service", "Podinfo","PODINFO",
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

