package coremodels

import(
  "strconv"
)

type Serverconfig interface {
  // Name of the server instance
  // Default value: OGC Api Features Server
  Title() string

  // Description of the server instance
  // Default value: This server provides an API to geospatial data
  Description() string

  // Protocol of the server (eg: http or https).
  // Default value: http
  Protocol() string

  // Hostname on which clients can access the server. This value is used to
  // build the links. This will be send as output. Therefore, if a network
  // component is placed between the server and the client, the hostname for
  // the client can be different (eg: if a proxy server host this component
  // and publishes under a specific hostname, than this hostname must be used)
  // Default value: localhost
  Host() string

  // Port on which clients can access the server
  // Default value: 8080
  Port() int

  // Mounting path on which clients can access the server.
  // Example: /endpoint
  Mountingpath() string

  // Returns the full host, including protocol and port, eg:
  // https://organisation.com:8080
  FullHost() string

  // Returns the full path of the mounted application, eg:
  // https://organisation.com:8080/endpoint
  FullUri() string

  SetTitle(string)
  SetDescription(string)
  SetProtocol(string)
  SetHost(string)
  SetPort(int)
  SetMountingpath(string)
}

type serverconfig struct {
  title string
  description string
  protocol string
  host string
  port int
  mountingpath string
}

func NewServerConfig() Serverconfig {
  return &serverconfig{
    title: "OGC Api Features Server",
    description: "This server provides an API to geospatial data",
    protocol: "http",
    host: "localhost",
    port: 8080,
  }
}

func (config *serverconfig) Title() string {
  return config.title
}

func (config *serverconfig) Description() string {
  return config.description
}

func (config *serverconfig) Protocol() string {
  return config.protocol
}

func (config *serverconfig) Host() string {
  return config.host
}

func (config *serverconfig) Port() int {
  return config.port
}

func (config *serverconfig) Mountingpath() string {
  return config.mountingpath
}

func (config *serverconfig) FullHost() string {
  result := config.Protocol() + "://" + config.Host()
  result += ":" + strconv.Itoa(config.Port())
  return result
}

func (config *serverconfig) FullUri() string {
  result := config.FullHost()
  result += config.Mountingpath()
  return result
}

func (config *serverconfig) SetTitle(title string) {
  config.title = title
}

func (config *serverconfig) SetDescription(description string) {
  config.description = description
}

func (config *serverconfig) SetProtocol(protocol string) {
  config.protocol = protocol
}

func (config *serverconfig) SetHost(host string) {
  config.host = host
}

func (config *serverconfig) SetPort(port int) {
  config.port = port
}

func (config *serverconfig) SetMountingpath(mountingpath string) {
  config.mountingpath = mountingpath
}
