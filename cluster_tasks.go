package goproxmoxapi

// structure representing Proxmox Task status
type TaskEntry struct {
  ExitStatus string  // "OK"
  Id         string  // "300"
  Node       string  // "pve"
  Pid        int     // 17899
  PStart     int     // 15768969
  Status     string  // "stopped"
  Type       string  // "vzcreate"
  UpId       string  // "UPID:pve:000045EB:00F09D89:5836FD57:vzcreate:300:root@pam:"
  User       string  // "root@pam"
  StartTime  float64 // 1479998807
}

// structure representing Proxmox Recent Task status
type RecentTask struct {
  ExitStatus string  // "OK"
  Id         string  // "300"
  Node       string  // "pve"
  Pid        int     // 17899
  PStart     int     // 15768969
  Status     string  // "stopped"
  Type       string  // "vzcreate"
  UpId       string  // "UPID:pve:000045EB:00F09D89:5836FD57:vzcreate:300:root@pam:"
  User       string  // "root@pam"
  Saved      string  // "1"
  StartTime  string  // 1479998807
}

// structure representing Task Log Entry
type TaskLogEntry struct {
  N int    // Task Log entry sequence number
  T string // Task log entry message
}

// Read task status.
func (tsks TaskEntry) GetTaskStatus(c *Client) (TaskEntry, error) {
  rts := TaskEntry{}
  // check definition of required fields
  if tsks.UpId == "" {
    return rts, &errorString{ "Task UpId is required to retrieve Task Status", }
  }
  if tsks.Node == "" {
    return rts, &errorString{ "Node name is required to retrieve Task Status", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/nodes/" + tsks.Node + "/tasks/" + tsks.UpId + "/status", nil )
  if err != nil {
	  return rts, err
  } else {
    err = dataUnmarshal( rbody, &rts )

    // Any Error parsing json ?
		return rts, err
  }
}

// Get all log entries for a given task
func (tsk TaskEntry) GetTaskLogEntries(c *Client) ([]TaskLogEntry, error) {
  tasklogs := []TaskLogEntry{}
  // check definition of required fields
  if tsk.UpId == "" {
    return tasklogs, &errorString{ "Task UpId is required to retrieve Task log entries", }
  }
  if tsk.Node == "" {
    return tasklogs, &errorString{ "Node name is required to retrieve Task log entries", }
  }

  _, rbody, err := c.NewRequest("GET", "/api2/json/nodes/" + tsk.Node + "/tasks/" + tsk.UpId + "/log", nil )
  if err != nil {
	  return tasklogs, err
  } else {
    err = dataUnmarshal( rbody, &tasklogs )

    // Any Error parsing json ?
		return tasklogs, err
  }
}

// Get Proxmox recent tasks (cluster wide).
func GetRecentTasks(c *Client) ([]RecentTask, error) {
  tasks := []RecentTask{}

  _, rbody, err := c.NewRequest("GET", "/api2/json/cluster/tasks", nil )
  if err != nil {
	  return tasks, err
  } else {
    err = dataUnmarshal( rbody, &tasks )

    // Any Error parsing json ?
		return tasks, err
  }
}
