package types

import "hanye/pkg/e"

func JsonSuccess() Default {
	Common := Default{
		Code: e.SUCCESS,
		Msg:  "",
	}
	return Common
}

func JsonError(errcode int) Default {
	Common := Default{
		Code: 1,
		Msg:  e.GetMsg(errcode),
	}
	return Common
}
