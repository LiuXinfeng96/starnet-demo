package routers

import (
	"starnet-demo/src/handlers"
	"starnet-demo/src/services"
)

const ROUTERS_HEADER = "/satellitebc"

const ROUTERS_USER = ROUTERS_HEADER + "/user"

const ROUTERS_CONTROL = ROUTERS_HEADER + "/control"

const ROUTERS_EXEC = ROUTERS_HEADER + "/exec"

const ROUTERS_TRACE = ROUTERS_HEADER + "/trace"

const ROUTERS_MONITOR = ROUTERS_HEADER + "/monitor"

func LoadNoTokenRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_USER)
	{
		routerGroup.POST("/register", handlers.Register(s))
		routerGroup.POST("/login", handlers.Login(s))
	}
}

func LoadUserRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_USER)
	{
		routerGroup.POST("/getuserinfo", handlers.GetUserInfo(s))
	}
}

func LoadControlRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_CONTROL)
	{
		routerGroup.POST("/adddebris", handlers.ControlAddDebris(s))
		routerGroup.POST("/getdebrislist", handlers.ControlGetDebrisList(s))

		routerGroup.POST("/addinstruction")
		routerGroup.POST("/getinstructionlist")

		routerGroup.POST("/getexecresultlist")
		routerGroup.POST("/getexecresult")

		routerGroup.POST("/addorbit")
		routerGroup.POST("/getorbitlist")

		routerGroup.POST("/addconstellation")
		routerGroup.POST("/getconstellationlist")

		routerGroup.POST("/getoperationlist")

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

		routerGroup.POST("/addcontrols")
		routerGroup.POST("/getcontrolslist")

		routerGroup.POST("/addfault")
		routerGroup.POST("/getfaultlist")

		routerGroup.POST("/addnetstate")
		routerGroup.POST("/getnetstatelist")

		routerGroup.POST("/addcommstate")
		routerGroup.POST("/getcommstatelist")

		routerGroup.POST("/getoperationlist")

	}
}

func LoadTraceRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_TRACE)
	{
		routerGroup.POST("/getdebrishistory")
		routerGroup.POST("/getinstructionhistory")
		routerGroup.POST("/getconstellationhistory")
		routerGroup.POST("/getoperationhistory")
		routerGroup.POST("/getsatellitestatehistory")
		routerGroup.POST("/getcontrolshistory")
		routerGroup.POST("/getfaulthistory")
		routerGroup.POST("/getnetstatehistory")
		routerGroup.POST("/getcommstatehistory")
	}
}

// func LoadMonitorRouter(s *services.Server) {
// 	routerGroup := s.GetGinEngine().Group(ROUTERS_MONITOR)
// 	{

// 	}
// }
