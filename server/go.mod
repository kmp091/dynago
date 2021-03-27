module github.com/kmp091/dynago/server

go 1.16

replace github.com/kmp091/dynago/messages => ../messages

require (
	github.com/kmp091/dynago/messages v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
)
