module github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service

go 1.15

require (
	github.com/BlackspaceInc/Backend/user-management-service v0.0.0-20200610031650-f616806382e1
	github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core v0.0.0-20201129002943-72f59f9f86de
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/color v1.10.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-openapi/runtime v0.19.24
	github.com/gomodule/redigo v1.8.3
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/keratin/authn-go v1.1.0
	github.com/prometheus/client_golang v1.8.0
	github.com/sony/gobreaker v0.4.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/swaggo/http-swagger v0.0.0-20200308142732-58ac5e232fba
	github.com/swaggo/swag v1.6.9
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gopkg.in/square/go-jose.v2 v2.5.1
	k8s.io/klog/v2 v2.4.0
)
