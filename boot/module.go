package boot

import (
	"novelcenter/module/home"
	"novelcenter/module/public"
	"novelcenter/module/sys"
)

func InitModules() {
	public.InitModule()
	home.InitModule()
	sys.InitModule()
}
