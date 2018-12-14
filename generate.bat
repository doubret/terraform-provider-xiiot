go fmt github.com/doubret/terraform-provider-xiiot/generator
go build github.com/doubret/terraform-provider-xiiot/generator

generator.exe -i %GOPATH%/src/github.com/doubret/terraform-provider-xiiot/spec

REM go fmt github.com/doubret/terraform-provider-netscaler/netscaler/utils
go fmt github.com/doubret/terraform-provider-xiiot/xiiot/resources
REM go fmt github.com/doubret/terraform-provider-netscaler/netscaler/bindings
go fmt github.com/doubret/terraform-provider-xiiot/xiiot

REM go fmt