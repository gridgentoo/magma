module fbc/cwf/radius

replace (
	fbc/lib/go/machine => ../lib/go/machine
	fbc/lib/go/radius => ../lib/go/radius
)

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	fbc/lib/go/machine v0.0.0-00010101000000-000000000000
	fbc/lib/go/radius v0.0.0-00010101000000-000000000000
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/donovanhide/eventsource v0.0.0-20171031113327-3ed64d21fb0b
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/golang/protobuf v1.3.1
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/uuid v1.1.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3
	github.com/stretchr/testify v1.3.0
	github.com/yuin/gopher-lua v0.0.0-20190514113301-1cd887cd7036 // indirect
	go.opencensus.io v0.21.0
	go.uber.org/atomic v1.4.0
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/sys v0.0.0-20191002091554-b397fe3ad8ed // indirect
	google.golang.org/grpc v1.21.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

go 1.13
