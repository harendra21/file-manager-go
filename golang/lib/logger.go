package lib

import (
	"encoding/json"
	"os"
	"runtime"
	"sync"

	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
)

type Logger struct{}

var logrusLogger *logrus.Logger
var loggerSetup sync.Once

func SetupLogger() {
	loggerSetup.Do(func() {
		logrusLogger = logrus.New()
	})

	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	var err error
	logrusLogger.Level, err = logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logrusLogger.Level = logrus.TraceLevel
	}
}

func init() {
	SetupLogger()
}

func (l Logger) LogStruct(data interface{}) {

	x, _ := json.Marshal(data)
	logrusLogger.Debug(string(x))
}

func (l Logger) Log(data interface{}) {

	if logrusLogger == nil {
		SetupLogger()
	}

	logrusLogger.Debug(data)
}

func (l Logger) Info(msg interface{}, extraFields map[string]interface{}) {

	if logrusLogger == nil {
		SetupLogger()
	}

	logrusLogger.WithFields(l.GetWithFields(extraFields, false)).Info(msg)
}

func (l Logger) Warning(msg interface{}, extraFields map[string]interface{}) {

	if logrusLogger == nil {
		SetupLogger()
	}

	logrusLogger.WithFields(l.GetWithFields(extraFields, false)).Warn(msg)
}

func (l Logger) Debug(msg interface{}, extraFields map[string]interface{}) {

	if logrusLogger == nil {
		SetupLogger()
	}

	logrusLogger.WithFields(l.GetWithFields(extraFields, true)).Debug(msg)
}

func (l Logger) Panic(msg interface{}, extraFields map[string]interface{}) {

	if logrusLogger == nil {
		SetupLogger()
	}

	logrusLogger.WithFields(l.GetWithFields(extraFields, true)).Panic(msg)
}

func (l Logger) Error(err error) {

	if logrusLogger == nil {
		SetupLogger()
	}

	err = errors.Wrap(err, 1)
	errStack := err.(*errors.Error).ErrorStack()
	extraFields := map[string]interface{}{
		"trace": errStack,
	}
	logrusLogger.WithFields(l.GetWithFields(extraFields, true)).Error(err.Error())
}

func (l Logger) GetWithFields(extraFields map[string]interface{}, tracing bool) logrus.Fields {
	var fields = logrus.Fields{}

	if tracing {
		pc, file, line, ok := runtime.Caller(2)
		if ok {
			f := runtime.FuncForPC(pc)
			fields = logrus.Fields{
				"file": file,
				"func": f.Name(),
				"line": line,
			}
		}
	}

	if extraFields != nil {
		fields["meta_data"] = extraFields
	}

	return fields
}
