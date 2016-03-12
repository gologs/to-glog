/*
Copyright 2016 James DeFelice

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package toglog

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/gologs/log/levels"
)

var Default = LeveledLogger(2)

type adapter struct {
	callerDepth int
}

func LeveledLogger(callerDepth int) levels.Interface {
	return &adapter{callerDepth}
}

// Debugf invokes glog.V(1).Info
func (x *adapter) Debugf(m string, a ...interface{}) {
	if glog.V(1) {
		if m == "" {
			glog.InfoDepth(x.callerDepth, a...)
		} else {
			glog.InfoDepth(x.callerDepth, fmt.Sprintf(m, a...))
		}
	}
}

// Infof invokes glog.Info
func (x *adapter) Infof(m string, a ...interface{}) {
	if m == "" {
		glog.InfoDepth(x.callerDepth, a...)
	} else {
		glog.InfoDepth(x.callerDepth, fmt.Sprintf(m, a...))
	}
}

// Warnf invokes glog.Warning
func (x *adapter) Warnf(m string, a ...interface{}) {
	if m == "" {
		glog.WarningDepth(x.callerDepth, a...)
	} else {
		glog.WarningDepth(x.callerDepth, fmt.Sprintf(m, a...))
	}
}

// Errorf invokes glog.Error
func (x *adapter) Errorf(m string, a ...interface{}) {
	if m == "" {
		glog.ErrorDepth(x.callerDepth, a...)
	} else {
		glog.ErrorDepth(x.callerDepth, fmt.Sprintf(m, a...))
	}
}

// Fatalf invokes glog.Exit
func (x *adapter) Fatalf(m string, a ...interface{}) {
	if m == "" {
		glog.ExitDepth(x.callerDepth, a...)
	} else {
		glog.ExitDepth(x.callerDepth, fmt.Sprintf(m, a...))
	}
}

// Panicf deviates from the intent of the interface: it doesn't panic but instead
// invokes glog.Fatal
func (x *adapter) Panicf(m string, a ...interface{}) {
	if m == "" {
		glog.FatalDepth(x.callerDepth, a...)
	} else {
		glog.FatalDepth(x.callerDepth, fmt.Sprintf(m, a...))
	}
}
