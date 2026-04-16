# diagnostics_api

Compliance scans, alerts, events, and tag filters

**get_alert** — Get an alert by UID
`{"mode": "get_alert", "project_uid": "proj-uid", "component": "cluster", "alert_uid": "alert-uid"}`

**delete_alert** — Delete an alert
`{"mode": "delete_alert", "project_uid": "proj-uid", "component": "cluster", "alert_uid": "alert-uid"}`

**list_tag_filters** — List all tag filters
`{"mode": "list_tag_filters"}`

**get_scan_config** — Get cluster compliance scan configuration
`{"mode": "get_scan_config", "uid": "cluster-uid-here"}`

**create_scan_config** — Create cluster compliance scan configuration
`{"mode": "create_scan_config", "uid": "cluster-uid", "config": {"schedule": {"scheduledRunTime": "0 3 * * 0"}, "driverSpec": {"kubeBench": {"enabled": true}}}}`

**update_scan_config** — Update cluster compliance scan configuration
`{"mode": "update_scan_config", "uid": "cluster-uid", "config": {"schedule": {"scheduledRunTime": "0 3 * * 0"}}}`

**create_alert** — Create an alert channel
`{"mode": "create_alert", "alert": {"metadata": {"name": "cpu-alert"}}, "project_uid": "proj-uid", "component": "cluster"}`

