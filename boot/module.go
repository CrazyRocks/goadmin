package boot

import (
	"goadmin/module/home"
	"goadmin/module/public"
	"goadmin/module/sys"
)

func InitModules() {
	public.InitModule()
	home.InitModule()
	sys.InitModule()
}
