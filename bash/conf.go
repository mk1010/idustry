package bash

var logConf = `
level: "debug"
development: true
disableCaller: false
disableStacktrace: false
sampling:
encoding: "console"

# encoder
encoderConfig:
  messageKey: "message"
  levelKey: "level"
  timeKey: "time"
  nameKey: "logger"
  callerKey: "caller"
  stacktraceKey: "stacktrace"
  lineEnding: ""
  levelEncoder: "capital"
  timeEncoder: "iso8601"
  durationEncoder: "millis"
  callerEncoder: "short"
  nameEncoder: ""

outputPaths:
  - "stderr"
errorOutputPaths:
  - "stderr"
initialFields:
`
