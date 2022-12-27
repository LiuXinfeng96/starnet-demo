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
		routerGroup.GET("/getuserinfo", handlers.GetUserInfo(s))
	}
}

func LoadControlRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_CONTROL)
	{
		routerGroup.POST("/adddebris", handlers.ControlAddDebris(s))
		routerGroup.GET("/getdebrislist", handlers.ControlGetDebrisList(s))

		routerGroup.POST("/addinstruction", handlers.ControlAddInstruction(s))
		routerGroup.GET("/getinstructionlist", handlers.ControlGetInstructionList(s))

		routerGroup.GET("/getexecresultlist", handlers.ControlGetExecResultList(s))
		routerGroup.GET("/getexecresult", handlers.ControlGetExecResult(s))

		routerGroup.POST("/addorbit", handlers.ControlAddOrbit(s))
		routerGroup.GET("/getorbitlist", handlers.ControlGetOrbitList(s))

		routerGroup.POST("/addconstellation", handlers.ControlAddConstellation(s))
		routerGroup.GET("/getconstellationlist", handlers.ControlGetConstellationList(s))

		routerGroup.GET("/getoperationlist", handlers.ControlGetOperationList(s))

		routerGroup.GET("/getloginloglist", handlers.ControlGetLoginLogList(s))

	}
}

func LoadExecRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_EXEC)
	{
		routerGroup.POST("/adddebris", handlers.ExecAddDebris(s))
		routerGroup.GET("/getdebrislist", handlers.ExecGetDebrisList(s))

		routerGroup.GET("/getinstructionlist", handlers.ExecGetInstructionList(s))
		routerGroup.GET("/getinstruction", handlers.ExecGetInstruction(s))

		routerGroup.GET("/getexecresultlist", handlers.ExecGetExecResultList(s))
		routerGroup.GET("/getexecresult", handlers.ExecGetExecResult(s))

		routerGroup.POST("/addsatellitestate", handlers.ExecAddSatelliteState(s))
		routerGroup.GET("/getsatellitestatelist", handlers.ExecGetSatelliteStateList(s))

		routerGroup.POST("/addcontrols", handlers.ExecAddControl(s))
		routerGroup.GET("/getcontrolslist", handlers.ExecGetControlList(s))

		routerGroup.POST("/addfault", handlers.ExecAddFault(s))
		routerGroup.GET("/getfaultlist", handlers.ExecGetFaultList(s))

		routerGroup.POST("/addnetstate", handlers.ExecAddNetState(s))
		routerGroup.GET("/getnetstatelist", handlers.ExecGetNetStateList(s))

		routerGroup.POST("/addcommstate", handlers.ExecAddCommState(s))
		routerGroup.GET("/getcommstatelist", handlers.ExecGetCommStateList(s))

		routerGroup.GET("/getoperationlist", handlers.ExecGetOperationList(s))

	}
}

func LoadTraceRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_TRACE)
	{
		routerGroup.GET("/getdebrishistory", handlers.TraceGetDebris(s))
		routerGroup.GET("/tracedebrislist", handlers.TraceGetDebrisList(s))

		routerGroup.GET("/getinstructionhistory", handlers.TraceGetInstruction(s))
		routerGroup.GET("/traceinstructionlist", handlers.TraceGetInstructionList(s))

		routerGroup.GET("/getconstellationhistory", handlers.TraceGetConstellation(s))
		routerGroup.GET("/traceconstellationlist", handlers.TraceGetConstellationList(s))

		routerGroup.GET("/getoperationhistory", handlers.TraceGetOperation(s))
		routerGroup.GET("/traceoperationlist", handlers.TraceGetOperationList(s))

		routerGroup.GET("/getsatellitestatehistory", handlers.TraceGetSatelliteState(s))
		routerGroup.GET("/tracesatellitestatelist", handlers.TraceGetSatelliteStateList(s))

		routerGroup.GET("/getcontrolshistory", handlers.TraceGetControl(s))
		routerGroup.GET("/tracecontrolslist", handlers.TraceGetControlList(s))

		routerGroup.GET("/getfaulthistory", handlers.TraceGetFault(s))
		routerGroup.GET("/tracefaultlist", handlers.TraceGetFaultList(s))

		routerGroup.GET("/getnetstatehistory", handlers.TraceGetNetState(s))
		routerGroup.GET("/tracenetstatelist", handlers.TraceGetNetStateList(s))

		routerGroup.GET("/getcommstatehistory", handlers.TraceGetCommetState(s))
		routerGroup.GET("/tracecommstatelist", handlers.TraceGetCommStateList(s))
	}
}

func LoadMonitorRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_MONITOR)
	{
		routerGroup.GET("/getstate", handlers.MonitorGetAllState(s))
		routerGroup.GET("/getchaindatanum", handlers.MonitorGetTableCount(s))
		routerGroup.GET("/getfaultinfo", handlers.MonitorGetFaultInfo(s))
		routerGroup.GET("/getearlywarning", handlers.MonitorGetEarlWarning(s))
		routerGroup.GET("/getchaininfo", handlers.MonitorGetChainInfo(s))
		routerGroup.GET("/getlatestthreat", handlers.MonitorGetLatestThreat(s))
		routerGroup.GET("/getlatestinstruction", handlers.MonitorGetLatestInstruction(s))
		routerGroup.GET("/getserverlasttime", handlers.MonitorGetLastTime(s))
	}
}
