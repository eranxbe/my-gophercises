# Notes
simple cmd app, read from csv, check user input, basic timer implementation and basic channel usage.

** flag package:
- flag.<Type> to get arg
- e.g `flag.String("name", value, "help message")`

** file operations:
- `file, err := os.Open(*fileName)`
- `reader = fileType.NewReader(file)`
- `reader.ReadAll()`

** channels:
- create 2 channels for timer and answer retrieval
- use select case to determine the value from channel
- answer:
`answerCh := make(chan string)`
`answerCh <- answer`
`case answer := <- answerCh`
- timer: 
`case <-timer.C`

** timer:
- initialize timer as close to use as possible: `timer := time.NewTimer(time.Duration(*timeValue * Duration))`
- use select case and create a channel for it: 
