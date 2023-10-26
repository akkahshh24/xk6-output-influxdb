// Package influxdb registers the xk6-influxdb extension
package influxdb

import (
	"github.com/akkahshh24/xk6-output-influxdb/pkg/influxdb"
	"github.com/akkahshh24/xk6-output-influxdb/xmlsigner"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/output"
)

func init() {
	modules.Register("k6/x/xmlsigner", new(xmlsigner.XmlSigner))
	output.RegisterExtension("xk6-influxdb", func(p output.Params) (output.Output, error) {
		return influxdb.New(p)
	})
}
