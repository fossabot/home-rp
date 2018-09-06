package model

import (
	"github.com/sirupsen/logrus"
	"strings"
)

type Route struct {
	Host   string
	RawUrl string
}

func ReadRules(str string) []Route {

	var results []Route

	if str == "" {
		return results
	}

	var routes []string

	//Handle multiple comma delimited items
	if strings.Contains(str, ",") {
		routes = strings.Split(str, ",")
	} else {
		routes = append(routes, str)
	}

	//Handle equals
	for _, ro := range routes {
		sl := strings.Split(ro, "=")
		if len(sl) != 2 {
			//Malformed Route!
			logrus.Fatalln("Malformed Route! Must include single equals. Instead found:", ro)
			continue
		}
		results = append(results, Route{Host: sl[0], RawUrl: sl[1]})

	}

	return results
}
