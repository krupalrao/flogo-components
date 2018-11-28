package transformXmlToJson

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
	xj "github.com/basgys/goxml2json"
)

const (

	ivContent = "content"
	ivSpec    = "spec"
	ovResult  = "result"
)

var log = logger.GetLogger("activity-tibco-transform")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {


	input := context.GetInput(ivContent).(string)
	//spec := context.GetInput(ivSpec).(string)

	xml := strings.NewReader(input)
	output, err := xj.Convert(xml) 

	if err != nil {
		return false, err
	}

	log.Debugf("Result: %s", output)

	// Set the output value in the context
	context.SetOutput(ovResult, output.String())
	return true, nil
}
