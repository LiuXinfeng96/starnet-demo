package routers

import (
	"starnet-demo/src/handlers"
	"starnet-demo/src/services"
)

const ROUTERS_HEADER = "/satellitebc"

const ROUTERS_USER = ROUTERS_HEADER + "/user"

const ROUTERS_CONTROL = ROUTERS_HEADER + "/control"

const ROUTERS_EXEC = ROUTERS_HEADER + "/exec"

func LoadUserRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_USER)
	{
		routerGroup.POST("/register")
		routerGroup.POST("/login")
		routerGroup.POST("/logout")
		routerGroup.POST("/getuserinfo")
	}
}

func LoadControlRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_CONTROL)
	{
		routerGroup.POST("/adddebris", handlers.AddDebris(s))
		routerGroup.POST("/getdebrislist", handlers.GetDebrisList(s))
		routerGroup.POST("/getdebrishistory")

		routerGroup.POST("/addinstruction")
		routerGroup.POST("/getinstructionlist")
		routerGroup.POST("/getinstructionhistory")

		routerGroup.POST("/getexecresultlist")
		routerGroup.POST("/getexecresult")

		routerGroup.POST("/addorbit")
		routerGroup.POST("/getorbitlist")

		routerGroup.POST("/addconstellation")
		routerGroup.POST("/getconstellationlist")
		routerGroup.POST("/getconstellationhistory")

		routerGroup.POST("/getoperationlist")
		routerGroup.POST("/getoperationhistory")

		routerGroup.POST("/getloginloglist")

	}
}

func LoadExecRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_EXEC)
	{
		routerGroup.POST("/adddebris")
		routerGroup.POST("/getdebrislist")

		routerGroup.POST("/getinstructionlist")
		routerGroup.POST("/getinstruction")

		routerGroup.POST("/getexecresultlist")
		routerGroup.POST("/getexecresult")

		routerGroup.POST("/addsatellitestate")
		routerGroup.POST("/getsatellitestatelist")
		routerGroup.POST("/getsatellitestatehistory")

		routerGroup.POST("/addcontrols")
		routerGroup.POST("/getcontrolslist")
		routerGroup.POST("/getcontrolshistory")

		routerGroup.POST("/addfault")
		routerGroup.POST("/getfaultlist")
		routerGroup.POST("/getfaulthistory")

		routerGroup.POST("/addnetstate")
		routerGroup.POST("/getnetstatelist")
		routerGroup.POST("/getnetstatehistory")

		routerGroup.POST("/addcommstate")
		routerGroup.POST("/getcommstatelist")
		routerGroup.POST("/getcommstatehistory")

		routerGroup.POST("/getoperationlist")

	}
}
