package gamePart

import (
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	//LOAD GENERATED DATA TO ARRRAYS
	for _, val := range MOD_GENERATED_DATA.blocks {
		//colors[val.ID] = val.col
		blocks[val.ID] = BLOCK{ID: val.ID, col: val.col}
	}

	mod := L.SetFuncs(L.NewTable(), exports)
	//L.SetField(mod, "name", lua.LString("value"))

	// returns the module
	L.Push(mod)

	return 1
}

func WorldEventLoader(L *lua.LState) map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"isWorldAlreadyOpened": func(L *lua.LState) int {
			L.Push(lua.LBool(api.WorldEvent.isWorldAlreadyOpened))
			return 1
		},
	}
}

func BlockEventLoader(L *lua.LState) map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"isBlockSet": func(L *lua.LState) int {
			L.Push(lua.LBool(api.isBlockSet.isActive))
			return 1
		},
		"getBlockCords": func(L *lua.LState) int {
			mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{})
			L.SetField(mod, "x", lua.LNumber(int(api.isBlockSet.cords.x)))
			L.SetField(mod, "y", lua.LNumber(int(api.isBlockSet.cords.y)))

			L.Push(mod)
			return 1
		},
		"getBlockType": func(L *lua.LState) int {
			L.Push(lua.LNumber(api.isBlockSet.blockType))
			return 1
		},
	}
}

var exports = map[string]lua.LGFunction{
	"set_block": set_block,
	"get_block": func(L *lua.LState) int {
		X := L.CheckNumber(1)
		Y := L.CheckNumber(2)
		L.Push(lua.LNumber(MAP[int(X)][int(Y)].ID))
		return 1
	},
	"getBlockEvent": func(L *lua.LState) int {
		mod := L.SetFuncs(L.NewTable(), BlockEventLoader(L))
		L.Push(mod)
		return 1
	},
	"getWorldEvent": func(L *lua.LState) int {
		mod := L.SetFuncs(L.NewTable(), WorldEventLoader(L))
		L.Push(mod)
		return 1
	},
}

func set_block(L *lua.LState) int {
	X := L.ToInt(1)
	Y := L.ToInt(2)
	TYPE := L.ToInt(3)
	MAP[X][Y].ID = rune(TYPE)
	return 0
}
