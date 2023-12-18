package functions

import (
	"runtime"
	"strings"
	"strconv"
	"time"
)

func SplitCommand(fullCommand string) []string {

	//if runtime.GOOS == "windows" {
	//	return strings.Split(strings.TrimSuffix(fullCommand, "\r\n"), " ")
	//} else {
	//	return strings.Split(strings.TrimSuffix(fullCommand, "\n"), " ")
	//}

	return strings.Split(strings.TrimSuffix(fullCommand, "\n"), " ")

}

func TreatCommand(command string) string {

	if runtime.GOOS == "windows" {
		return strings.TrimSuffix(command, "\r\n")
	} else {
		return strings.TrimSuffix(command, "\n")
	}

}

func GetCurrentTime() int{

	dt := time.Now()
	timeCalcNow := validateTimeLength(strconv.Itoa(dt.Year())) +  validateTimeLength(strconv.Itoa(int(dt.Month()))) +  validateTimeLength(strconv.Itoa(dt.Day())) +  validateTimeLength(strconv.Itoa(dt.Hour())) +  validateTimeLength(strconv.Itoa(dt.Minute())) +  validateTimeLength(strconv.Itoa(dt.Second()))
	timeCalcNowInt, _ := strconv.Atoi(timeCalcNow)

	return timeCalcNowInt
}

func validateTimeLength(strTime string) string{

	

	if len(strTime) <= 1{
		strTime = "0" + strTime
	}

	return strTime

}