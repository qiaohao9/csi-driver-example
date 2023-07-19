package mycsi

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IdentityServer struct {
	Driver *driver
}

var _ csi.IdentityServer = &IdentityServer{}

func NewIdentityServer(d *driver) *IdentityServer {
	return &IdentityServer{Driver: d}
}

func (ids *IdentityServer) GetPluginInfo(ctx context.Context, request *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	if ids.Driver.name == "" {
		return &csi.GetPluginInfoResponse{}, status.Errorf(codes.Unavailable, "driver name is missing")
	}
	if ids.Driver.version == "" {
		return &csi.GetPluginInfoResponse{}, status.Error(codes.Unavailable, "driver version is missing")
	}

	return &csi.GetPluginInfoResponse{
		Name:          ids.Driver.name,
		VendorVersion: ids.Driver.version,
	}, nil
}

func (ids *IdentityServer) GetPluginCapabilities(ctx context.Context, request *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	return &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{Type: &csi.PluginCapability_Service_{Service: &csi.PluginCapability_Service{Type: csi.PluginCapability_Service_CONTROLLER_SERVICE}}},
		},
	}, nil
}

func (ids *IdentityServer) Probe(ctx context.Context, request *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	return &csi.ProbeResponse{}, nil
}
