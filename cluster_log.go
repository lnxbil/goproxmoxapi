package goproxmoxapi

// structure representing Proxmox Log Entry
type LogEntry struct {
  Msg  string // "successful auth for user 'root@pam'"
  Node string // "pve"
  PID  int    // 327
  Pri  int    // 6
  Tag  string // "pvedaemon"
  Time float64    // 1479459397
  UID  int    // 5327
  User string // "root@pam"
}

// Get Proxmox recent logs (cluster wide).
func GetRecentLogs(c *Client) ([]LogEntry, error) {
  logs := []LogEntry{}

  _, rbody, err := c.NewRequest("GET", "/api2/json/cluster/log", nil )
  if err != nil {
    return logs, err
  } else {
    err = dataUnmarshal( rbody, &logs )

    return logs, err
  }
}
