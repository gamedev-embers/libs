package blocks

import "fmt"

type Server struct {
	Host  string         `json:"host" toml:"host"`
	Ports map[string]int `json:"ports" toml:"ports"`
}

func (s *Server) Check() error {
	if s.Host == "" {
		return fmt.Errorf("host is required")
	}
	if len(s.Ports) == 0 {
		return fmt.Errorf("at least one port is required")
	}
	return nil
}

func (s *Server) HasPort(portName string) bool {
	_, ok := s.Ports[portName]
	return ok
}

func (s *Server) GetPort(portName string) int {
	port, ok := s.Ports[portName]
	if !ok {
		panic(fmt.Errorf("port %s not found", portName))
	}
	return port
}

func (s *Server) LocalAddr(portName string) string {
	return fmt.Sprintf("127.0.0.1:%d", s.Ports[portName])
}

func (s *Server) PrivateAddr(portName string) string {
	return fmt.Sprintf("%s:%d", s.Host, s.Ports[portName])
}

func (s *Server) PublicAddr(portName string) string {
	return fmt.Sprintf("0.0.0.0:%d", s.Ports[portName])
}
