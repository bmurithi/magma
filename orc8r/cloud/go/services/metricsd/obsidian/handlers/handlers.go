/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package handlers

import (
	"fmt"

	"magma/orc8r/cloud/go/obsidian/handlers"
	"magma/orc8r/cloud/go/service/config"
	"magma/orc8r/cloud/go/services/metricsd/confignames"
	graphiteH "magma/orc8r/cloud/go/services/metricsd/graphite/handlers"
	graphiteAPI "magma/orc8r/cloud/go/services/metricsd/graphite/third_party/api"
	promH "magma/orc8r/cloud/go/services/metricsd/prometheus/handlers"

	"github.com/labstack/echo"
	promAPI "github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

// GetObsidianHandlers returns all obsidian handlers for metricsd
func GetObsidianHandlers(configMap *config.ConfigMap) []handlers.Handler {
	var ret []handlers.Handler
	client, err := promAPI.NewClient(promAPI.Config{Address: configMap.GetRequiredStringParam(confignames.PrometheusAddress)})
	if err != nil {
		ret = append(ret,
			handlers.Handler{Path: promH.QueryURL, Methods: handlers.GET, HandlerFunc: getInitErrorHandler(err)},
			handlers.Handler{Path: promH.QueryRangeURL, Methods: handlers.GET, HandlerFunc: getInitErrorHandler(err)},
		)
	} else {
		pAPI := v1.NewAPI(client)
		ret = append(ret,
			handlers.Handler{Path: promH.QueryURL, Methods: handlers.GET, HandlerFunc: promH.GetPrometheusQueryHandler(pAPI)},
			handlers.Handler{Path: promH.QueryRangeURL, Methods: handlers.GET, HandlerFunc: promH.GetPrometheusQueryRangeHandler(pAPI)},
		)
	}

	graphiteAddress := configMap.GetRequiredStringParam(confignames.GraphiteAddress)
	graphiteQueryPort := configMap.GetRequiredIntParam(confignames.GraphiteQueryPort)
	graphiteQueryAddress := fmt.Sprintf("%s://%s:%d", graphiteH.Protocol, graphiteAddress, graphiteQueryPort)
	graphiteClient, err := graphiteAPI.NewFromString(graphiteQueryAddress)
	if err != nil {
		ret = append(ret,
			handlers.Handler{Path: graphiteH.QueryURL, Methods: handlers.GET, HandlerFunc: getInitErrorHandler(err)},
		)
	} else {
		ret = append(ret,
			handlers.Handler{Path: graphiteH.QueryURL, Methods: handlers.GET, HandlerFunc: graphiteH.GetQueryHandler(graphiteClient)},
		)
	}
	return ret
}

func getInitErrorHandler(err error) func(c echo.Context) error {
	return func(c echo.Context) error {
		return handlers.HttpError(fmt.Errorf("initialization Error: %v", err), 500)
	}
}
