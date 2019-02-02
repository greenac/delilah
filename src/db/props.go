package db

import (
	"github.com/greenac/artemis/logger"
	"github.com/greenac/delilah/src/helpers"
)

type connProp string

const (
	host     connProp = "host"
	name     connProp = "dbname"
	password connProp = "password"
	port     connProp = "port"
	user     connProp = "user"
)

type Props struct {
	Host     string
	Name     string
	Password string
	Port     string
	User     string
}

func (pr *Props) Setup(props map[helpers.ConnectionVar]string) {
	pr.Host = props[helpers.Host]
	pr.Name = props[helpers.Name]
	pr.Password = props[helpers.Password]
	pr.Port = props[helpers.Port]
	pr.User = props[helpers.User]
}

func (pr *Props) ConnectionString() string {
	cs := ""
	for _, p := range []connProp{user, password, port, name, host} {
		prop := ""
		switch p {
		case host:
			prop = pr.Host
		case name:
			prop = pr.Name
		case password:
			prop = pr.Password
		case port:
			prop = pr.Port
		case user:
			prop = pr.User
		default:
			// TODO: implement proper error cases
			logger.Error("`DatabaseVars::ConnectionString` unhandled connection property:", p)
			panic("UNHANDLED_CONNECTION_PROPERTY")
		}

		cs += pr.strComp(p, cs, prop)
	}

	return cs
}

func (dv *Props) strComp(p connProp, cs string, pv string) string {
	v := ""
	if pv != "" {
		if cs == "" {
			v += string(p) + "=" + pv
		} else {
			v += " " + string(p) + "=" + pv
		}
	}

	return v
}
