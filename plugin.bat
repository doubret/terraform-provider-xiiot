go fmt github.com/doubret/terraform-provider-xiiot
go fmt github.com/doubret/terraform-provider-xiiot/client
go fmt github.com/doubret/terraform-provider-xiiot/xiiot
go install github.com/doubret/terraform-provider-xiiot
COPY %GOPATH%\bin\terraform-provider-xiiot.exe %APPDATA%\terraform.d\plugins\terraform-provider-xiiot.exe
