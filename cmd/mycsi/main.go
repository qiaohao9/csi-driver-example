package main

import (
	"csi-driver-example/pkg/mycsi"
	"flag"
	"github.com/container-storage-interface/spec/lib/go/csi"
)

var (
	endpoint                = flag.String("endpoint", "unix:///csi/csi.sock", "CSI endpoint")
	nodeId                  = flag.String("nodeid", "", "node id")
	enableControllerService = flag.Bool("enable-controllerserver", false, "enable controller service")
	enableNodeService       = flag.Bool("enable-nodeserver", false, "enable node service")
)

func main() {
	flag.Parse()

	d := mycsi.NewDriver(*nodeId, *endpoint)

	if *enableControllerService {
		csi.RegisterControllerServer(nil, mycsi.NewControllerServer(d))
	}

	if *enableNodeService {
		csi.RegisterNodeServer(nil, mycsi.NewNodeServer(d))
	}
}
