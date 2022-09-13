/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package glogger defines glog-based logging for grpc.
// Importing this package will install glog as the logger used by grpclog.
package glog

import (
	"fmt"
	"google.golang.org/grpc/grpclog"
)

const d = 2

func initGrpcLog() {
	grpclog.SetLoggerV2(&glogger{})
}

type glogger struct{}

func (g *glogger) Info(args ...interface{}) {
	InfoDepth(d, args...)
}

func (g *glogger) Infoln(args ...interface{}) {
	InfoDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Infof(format string, args ...interface{}) {
	InfoDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) InfoDepth(depth int, args ...interface{}) {
	InfoDepth(depth+d, args...)
}

func (g *glogger) Warning(args ...interface{}) {
	WarningDepth(d, args...)
}

func (g *glogger) Warningln(args ...interface{}) {
	WarningDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Warningf(format string, args ...interface{}) {
	WarningDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) WarningDepth(depth int, args ...interface{}) {
	WarningDepth(depth+d, args...)
}

func (g *glogger) Error(args ...interface{}) {
	ErrorDepth(d, args...)
}

func (g *glogger) Errorln(args ...interface{}) {
	ErrorDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Errorf(format string, args ...interface{}) {
	ErrorDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) ErrorDepth(depth int, args ...interface{}) {
	ErrorDepth(depth+d, args...)
}

func (g *glogger) Fatal(args ...interface{}) {
	FatalDepth(d, args...)
}

func (g *glogger) Fatalln(args ...interface{}) {
	FatalDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Fatalf(format string, args ...interface{}) {
	FatalDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) FatalDepth(depth int, args ...interface{}) {
	FatalDepth(depth+d, args...)
}

func (g *glogger) V(l int) bool {
	return bool(V(Level(l)))
}
