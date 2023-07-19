package mycsi

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
)

type NodeServer struct {
	Driver *driver
}

var _ csi.NodeServer = &NodeServer{}

func NewNodeServer(d *driver) *NodeServer {
	return &NodeServer{Driver: d}
}

func (ns *NodeServer) NodeStageVolume(ctx context.Context, request *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (ns *NodeServer) NodeUnstageVolume(ctx context.Context, request *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (ns *NodeServer) NodePublishVolume(ctx context.Context, request *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (ns *NodeServer) NodeUnpublishVolume(ctx context.Context, request *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (ns *NodeServer) NodeGetVolumeStats(ctx context.Context, request *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (ns *NodeServer) NodeExpandVolume(ctx context.Context, request *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (ns *NodeServer) NodeGetCapabilities(ctx context.Context, request *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{Capabilities: ns.Driver.nscap}, nil
}

func (ns *NodeServer) NodeGetInfo(ctx context.Context, request *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}
