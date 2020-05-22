package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePost{}, "dither/CreatePost", nil)
	cdc.RegisterConcrete(MsgLikePost{}, "dither/LikePost", nil)
	cdc.RegisterConcrete(MsgCreateChannel{}, "dither/CreateChannel", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
