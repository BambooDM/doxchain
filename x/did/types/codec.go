package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDidRequest{}, "did/CreateDid", nil)
	cdc.RegisterConcrete(&MsgUpdateDidRequest{}, "did/UpdateDid", nil)
	cdc.RegisterConcrete(&MsgDeleteDidRequest{}, "did/DeleteDid", nil)
	cdc.RegisterConcrete(&MsgCreateDidDocumentRequest{}, "did/CreateDidDocument", nil)
	cdc.RegisterConcrete(&MsgUpdateDidDocumentRequest{}, "did/UpdateDidDocument", nil)
	cdc.RegisterConcrete(&MsgDeleteDidDocumentRequest{}, "did/DeleteDidDocument", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDidRequest{},
		&MsgUpdateDidRequest{},
		&MsgDeleteDidRequest{},
		&MsgCreateDidDocumentRequest{},
		&MsgUpdateDidDocumentRequest{},
		&MsgDeleteDidDocumentRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
