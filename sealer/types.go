package sealer

type Entry struct {
	DST         string   `xml:"DST,attr"`
	Severity    string   `xml:"Severity,attr"`
	Secs        float32  `xml:"Secs,attr"`
	ThreadID    int      `xml:"ThreadId,attr"`
	Audience    string   `xml:"Audience,attr"`
	SysDateTime string   `xml:"SysDateTime,attr"`
	CategoryStr string   `xml:",chardata"`
	Category    Category `xml:"-"`
}

type EventLog struct {
	Entries               []Entry `xml:"Entry"`
	ComputerName          string  `xml:"ComputerName,attr"`
	SoftwareVersion       string  `xml:"SoftwareVersion,attr"`
	LogVersion            string  `xml:"LogVersion,attr"`
	ComponentName         string  `xml:"ComponentName,attr"`
	Process               string  `xml:"Process,attr"`
	ProcessID             string  `xml:"ProcessId,attr"`
	CreationDateTime      string  `xml:"CreationDateTime,attr"`
	TimeZoneOffsetMinutes string  `xml:"TimeZoneOffsetMinutes,attr"`
}
