// Code generated by goctl. DO NOT EDIT.
// Source: address.proto

package server

import (
	"context"

	"hanye/service/address/rpc/internal/logic"
	"hanye/service/address/rpc/internal/svc"
	"hanye/service/address/rpc/pb"
)

type AddressServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedAddressServer
}

func NewAddressServer(svcCtx *svc.ServiceContext) *AddressServer {
	return &AddressServer{
		svcCtx: svcCtx,
	}
}

func (s *AddressServer) AddAddress(ctx context.Context, in *pb.AddAddressReq) (*pb.AddAddressResp, error) {
	l := logic.NewAddAddressLogic(ctx, s.svcCtx)
	return l.AddAddress(in)
}

func (s *AddressServer) GetAddressList(ctx context.Context, in *pb.GetAddressListReq) (*pb.GetAddressListResp, error) {
	l := logic.NewGetAddressListLogic(ctx, s.svcCtx)
	return l.GetAddressList(in)
}

func (s *AddressServer) GetAddress(ctx context.Context, in *pb.GetAddressReq) (*pb.GetAddressResp, error) {
	l := logic.NewGetAddressLogic(ctx, s.svcCtx)
	return l.GetAddress(in)
}

func (s *AddressServer) GetDefaultAddress(ctx context.Context, in *pb.GetDefaultAddressReq) (*pb.GetDefaultAddressResp, error) {
	l := logic.NewGetDefaultAddressLogic(ctx, s.svcCtx)
	return l.GetDefaultAddress(in)
}

func (s *AddressServer) UpdateAddress(ctx context.Context, in *pb.UpdateAddressReq) (*pb.UpdateAddressResp, error) {
	l := logic.NewUpdateAddressLogic(ctx, s.svcCtx)
	return l.UpdateAddress(in)
}

func (s *AddressServer) SetDefaultAddress(ctx context.Context, in *pb.SetDefaultAddressReq) (*pb.SetDefaultAddressResp, error) {
	l := logic.NewSetDefaultAddressLogic(ctx, s.svcCtx)
	return l.SetDefaultAddress(in)
}

func (s *AddressServer) DeleteAddress(ctx context.Context, in *pb.DeleteAddressReq) (*pb.DeleteAddressResp, error) {
	l := logic.NewDeleteAddressLogic(ctx, s.svcCtx)
	return l.DeleteAddress(in)
}
