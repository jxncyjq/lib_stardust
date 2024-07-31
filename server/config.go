/**
    @author: cloudy
    @date: 2023-01-02
    @note:
**/
package server

// Server struct
type ServerConfig struct {
	ID        string            `json:"id" yaml:"id"`
	Type      string            `json:"type" yaml:"type"`
	Metadata  map[string]string `json:"metadata" yaml:"metadata"`
	Frontend  bool              `json:"frontend" yaml:"frontend"`
	Hostname  string            `json:"host_name" yaml:"hostname"`
	NameSpace string            `json:"name_space" yaml:"name_space"`
	SubSystem string            `json:"sub_system" yaml:"sub_system"`
}
