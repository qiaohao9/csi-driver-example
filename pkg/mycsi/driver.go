package mycsi

import (
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"os"
)

const (
	driverName    = "mycsi.csi.k8s.io"
	driverVersion = "v0.2.0"
)

type driver struct {
	nodeId  string
	name    string
	version string

	nscap []*csi.NodeServiceCapability
	cscap []*csi.ControllerServiceCapability
}

func NewDriver(nodeId, endpoint string) *driver {
	d := &driver{
		nodeId:  nodeId,
		name:    driverName,
		version: driverVersion,
	}

	if err := os.MkdirAll(fmt.Sprintf("/var/run/%s.sock", d.name), 0o755); err != nil {
		panic(err)
	}

	d.AddNodeServerServiceCapabilities([]csi.NodeServiceCapability_RPC_Type{
		csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
	})
	d.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_LIST_VOLUMES,
		csi.ControllerServiceCapability_RPC_GET_VOLUME,
	})

	return d
}

func (d *driver) Run() {
}

func (d *driver) AddNodeServerServiceCapabilities(nl []csi.NodeServiceCapability_RPC_Type) {
	var nsc []*csi.NodeServiceCapability

	for _, n := range nl {
		nsc = append(nsc, newNodeServiceCapability(n))
	}

	d.nscap = nsc
}

func (d *driver) AddControllerServiceCapabilities(cl []csi.ControllerServiceCapability_RPC_Type) {
	var csc []*csi.ControllerServiceCapability

	for _, c := range cl {
		csc = append(csc, newControllerServiceCapability(c))
	}

	d.cscap = csc
}

func newNodeServiceCapability(cap csi.NodeServiceCapability_RPC_Type) *csi.NodeServiceCapability {
	return &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: cap,
			},
		},
	}
}

func newControllerServiceCapability(cap csi.ControllerServiceCapability_RPC_Type) *csi.ControllerServiceCapability {
	return &csi.ControllerServiceCapability{
		Type: &csi.ControllerServiceCapability_Rpc{
			Rpc: &csi.ControllerServiceCapability_RPC{
				Type: cap,
			},
		},
	}
}
