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

type DatabaseVars struct {
	Host     string
	Name     string
	Password string
	Port     string
	User     string
}

func (dv *DatabaseVars) Setup(props map[helpers.ConnectionVar]string) {
	dv.Host = props[helpers.Host]
	dv.Name = props[helpers.Name]
	dv.Password = props[helpers.Password]
	dv.Port = props[helpers.Port]
	dv.User = props[helpers.User]
}

func (dv *DatabaseVars) ConnectionString() string {
	cs := ""
	for _, p := range []connProp{user, password, port, name, host} {
		prop := ""
		switch p {
		case host:
			prop = dv.Host
		case name:
			prop = dv.Name
		case password:
			prop = dv.Password
		case port:
			prop = dv.Port
		case user:
			prop = dv.User
		default:
			// TODO: implement proper error cases
			logger.Error("`DatabaseVars::ConnectionString` unhandled connection property:", p)
			panic("UNHANDLED_CONNECTION_PROPERTY")
		}

		cs += dv.strComp(p, cs, prop)
	}

	return cs
}

func (dv *DatabaseVars) strComp(p connProp, cs string, pv string) string {
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
