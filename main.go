package main

import "fmt"
import "flag"
import "io/ioutil"
import "strings"
import "time"

const TEMPLATE = `
# name

## date


### Attendees
@mention yourself and add others


### Agenda
- Stuff to talk about


### Discussion
- Stuff we actually talked about


### Action items
[ ] Let’s get this done @someone
`

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func processTemplateData(meetingName, attendees string) string {
	templateData := strings.Replace(TEMPLATE, "name", meetingName, 1)
	templateData = strings.Replace(templateData, "date", time.Now().Format("2006-01-02 15:04:05"), 1)
	return templateData
}

func writeTemplate(name, data string) {
	err := ioutil.WriteFile(name, []byte(data), 0644)
	checkErr(err)
}

func main() {
	meetingName := flag.String("name", "Meeting Name", "Name of the meeting")
	attendees := flag.String("attendees", "@me", "Attendes list (comma separated)")
	flag.Parse()
	fmt.Printf("Creating meeting notes for meeting: %s, with attendees: %s \n", *meetingName, *attendees)
	processedTemplate := processTemplateData(*meetingName, *attendees)
	writeTemplate(*meetingName, processedTemplate)
}
