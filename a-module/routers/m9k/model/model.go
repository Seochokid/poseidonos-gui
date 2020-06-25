package model

type Request struct {
	Command string      `json:"command"`
	Rid     string      `json:"rid"`
	Param   interface{} `json:"param,omitempty"`
}

type Response struct {
	Rid             string `json:"rid"`
	LastSuccessTime int64  `json:"lastSuccessTime"`
	Result          Result `json:"result"`
	Info            Info   `json:"info"`
}

type Result struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type Status struct {
	Module      string `json:"module"`
	Code        int    `json:"code"`
	Level       string `json:'level,omitempty"` 
	Description string `json:"description"`
	Problem     string `json:"problem,omitempty"`
	Solution    string `json:"solution,omitempty"`
}

type StatusList struct {
	StatusList []Status `json:"StatusList"`
}

type Info struct {
	State              string `json:"state"`
	Situation          string `json:"situation"`
	RebuildingProgress uint32 `json:"rebuliding_progress"`
	Capacity           uint64 `json:"capacity"`
	Used               uint64 `json:"used"`
}

type Device struct {
	DeviceName string `json:"deviceName"`
}

type SystemParam struct {
	Level string `json:"level,omitempty"`
	Name  string `json:"testname,omitempty"`
	Argc  int    `json:"argc"`
	Argv  string `json:"argv"`
}

type ArrayParam struct {
	FtType int      `json:"fttype,omitempty"`
	Buffer []Device `json:"buffer,omitempty"`
	Data   []Device `json:"data,omitempty"`
	Spare  []Device `json:"spare,omitempty"`
}
type MAgentParam struct {
	Time  string
	Level string
}
type DeviceParam struct {
	Name  string `json:"name,omitempty"`
	Spare string `json:"spare,omitempty"`
}
type VolumeParam struct {
	Name        string `json:"name,omitempty"`
	NewName     string `json:"newname,omitempty"`
	Size        uint64 `json:"size,omitempty"`
	Maxiops     uint64 `json:"maxiops,omitempty"`
	Maxbw       uint64 `json:"maxbw,omitempty"`
	NameSuffix  uint64 `json:"namesuffix,omitempty"`
	TotalCount  uint64 `json:"totalcount,omitempty"`
	StopOnError bool   `json:"stoponerror,omitempty"`
	MountAll    bool   `json:"mountall,omitempty"`
}

type CallbackMultiVol struct {
	TotalCount    int
	Pass          int
	Fail          int
	MultiVolArray []Response
}

type WBTParam struct {
	TestName string  `json:"testname,omitempty"`
	Argv     WBTArgv `json:"argv"`
}

type WBTArgv struct {
	Name      string `json:"name,omitempty"`
	Input     string `json:"input,omitempty"`
	Output    string `json:"output,omitempty"`
	Integrity string `json:"integrity,omitempty"`
	Access    string `json:"access,omitempty"`
	Operation string `json:"operation,omitempty"`
	Rba       string `json:"rba,omitempty"`
	Lba       string `json:"lba,omitempty"`
	Vsid      string `json:"vsid,omitempty"`
	Lsid      string `json:"lsid,omitempty"`
	Offset    string `json:"offset,omitempty"`
	Size      string `json:"size,omitempty"`
	Count     string `json:"count,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	Loc       string `json:"loc,omitempty"`
	Fd        string `json:"fd,omitempty"`
	Dev       string `json:"dev,omitempty"`
	Normal    string `json:"normal,omitempty"`
	Urgent    string `json:"urgent,omitempty"`
}

type BuildInfo struct {
	GitHash   string `json:"githash"`
	BuildTime string `json:"build_time"`
}

//type SMART struct {
//	AvailableSpare           string `json:"available_spare,omitempty"`
//	AvailableSpareSpace      string `json:"available_spare_space,omitempty"`
//	AvailableSpareThreshold  string `json:"available_spare_threshold,omitempty"`
//	ContollerBusyTime        string `json:"contoller_busy_time,omitempty"`
//	CriticalTemperatureTime  string `json:"critical_temperature_time,omitempty"`
//	CurrentTemperature       string `json:"current_temperature,omitempty"`
//	DataUnitsRead            string `json:"data_units_read,omitempty"`
//	DataUnitsWritten         string `json:"data_units_written,omitempty"`
//	DeviceReliability        string `json:"device_reliability,omitempty"`
//	HostReadCommands         string `json:"host_read_commands,omitempty"`
//	HostWriteCommands        string `json:"host_write_commands,omitempty"`
//	LifPercentageUsed        string `json:"life_percentage_used,omitempty"`
//	LifetimeErrorLogEntries  string `json:"lifetime_error_log_entries,omitempty"`
//	PowerCycles              string `json:"power_cycles,omitempty"`
//	PowerOnHours             string `json:"power_on_hours,omitempty"`
//	ReadOnly                 string `json:"read_only,omitempty"`
//	Temperature              string `json:"temperature,omitempty"`
//	UnrecoverableMediaErrors string `json:"unrecoverable_media_errors,omitempty"`
//	UnsafeShutdowns          string `json:"unsafe_shutdowns,omitempty"`
//	VolatileMemoryBackup     string `json:"volatile_memory_backup,omitempty"`
//	WarningTemperatureTime   string `json:"warning_temperature_time,omitempty"`
//}
