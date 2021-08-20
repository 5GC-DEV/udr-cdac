// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package logger

import (
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"

)

var (
	log         *logrus.Logger
	AppLog      *logrus.Entry
	InitLog     *logrus.Entry
	CfgLog      *logrus.Entry
	HandlerLog  *logrus.Entry
	DataRepoLog *logrus.Entry
	UtilLog     *logrus.Entry
	HttpLog     *logrus.Entry
	ConsumerLog *logrus.Entry
	GinLog      *logrus.Entry
)

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}


	AppLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "App"})
	InitLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "Init"})
	CfgLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "CFG"})
	HandlerLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "HDLR"})
	DataRepoLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "DRepo"})
	UtilLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "Util"})
	HttpLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "HTTP"})
	ConsumerLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "Consumer"})
	GinLog = log.WithFields(logrus.Fields{"component": "UDR", "category": "GIN"})
}

func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}

func SetReportCaller(set bool) {
	log.SetReportCaller(set)
}
