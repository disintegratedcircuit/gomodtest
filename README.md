# Known Changes

Due to this error:

        /go/pkg/mod/github.com/prometheus/prometheus@v2.5.0+incompatible/storage/tsdb/tsdb.go:144:3: unknown field 'WALFlushInterval' in struct literal of type tsdb.Options

Made this update:

        github.com/prometheus/prometheus v2.5.0+incompatible -> github.com/prometheus/prometheus v2.7.0+incompatible
