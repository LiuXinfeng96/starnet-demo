package services

func (s *Server) GetMasterContractName() string {
	return s.config.MaterC.Name
}

func (s *Server) GetExecContractName() string {
	return s.config.ExecC.Name
}

func (s *Server) GetMasterChainId() string {
	return s.config.MaterC.ChainId
}

func (s *Server) GetExecChainId() string {
	return s.config.ExecC.ChainId
}
