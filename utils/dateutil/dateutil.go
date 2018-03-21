package dateutil

import "time"
import "strings"
import "strconv"

func GetDateObject(dateStr string) time.Time {
  dateStr2 := strings.Replace(dateStr,"-","",-1)
  dateStr3 := strings.Replace(dateStr2,"/","",-1)
  yearStr  := dateStr3[0:4]
  monthStr := dateStr3[4:6]
  dayStr   := dateStr3[6:8]

  yearInt,_  := strconv.Atoi(yearStr)
  monthInt,_ := strconv.Atoi(monthStr)
  dayInt,  _ := strconv.Atoi(dayStr)
  dateObj := time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0,time.Local)
  return dateObj
}
