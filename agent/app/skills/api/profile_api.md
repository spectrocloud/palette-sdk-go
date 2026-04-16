# profile_api

Create, list, update, delete, publish, import, and export cluster profiles

**list_profiles** — List all cluster profiles (unfiltered)
`{"mode": "list_profiles"}`

**get_profile** — Get full cluster profile details
`{"mode": "get_profile", "uid": "profile-uid-here"}`

**get_profile_summary** — Get cluster profile summary
`{"mode": "get_profile_summary", "uid": "profile-uid-here"}`

**get_profile_uid** — Get profile UID by name and optional version
`{"mode": "get_profile_uid", "name": "eks-default", "version": "1.0.0"}`

**create_profile** — Create a new cluster profile
`{"mode": "create_profile", "profile": {"metadata": {"name": "my-profile"}, "spec": {"type": "cluster"}}}`

**update_profile** — Update an existing cluster profile
`{"mode": "update_profile", "profile": {"metadata": {"uid": "profile-uid", "name": "my-profile"}, "spec": {}}}`

**delete_profile** — Delete a cluster profile
`{"mode": "delete_profile", "uid": "profile-uid-here"}`

**import_profile** — Import a cluster profile from content string
`{"mode": "import_profile", "content": "...exported profile data..."}`

**search_profiles** — Search cluster profiles with filters (by cloud type, profile type, name, scope)
`{"mode": "search_profiles", "filter": {"filter": {"environment": ["aws"]}}}`
*To list AWS profiles: {"mode": "search_profiles", "filter": {"filter": {"environment": ["aws"]}}}*
*To list infra (cluster type) profiles: {"mode": "search_profiles", "filter": {"filter": {"profileType": ["infra"]}}}*
*To list add-on profiles: {"mode": "search_profiles", "filter": {"filter": {"profileType": ["add-on"]}}}*
*To list AWS infra profiles: {"mode": "search_profiles", "filter": {"filter": {"environment": ["aws"], "profileType": ["infra"]}}}*
*Profile types: infra, add-on, cluster (full stack), system*
*Environment values: aws, eks, azure, aks, gcp, gke, edge-native, vsphere, maas, all*
*Use this instead of list_profiles when filtering is needed*

**publish_profile** — Publish a draft cluster profile
`{"mode": "publish_profile", "uid": "profile-uid-here"}`

**get_profile_variables** — Get profile variables
`{"mode": "get_profile_variables", "uid": "profile-uid-here"}`

**get_pack** — Get pack details including values.yaml
`{"mode": "get_pack", "uid": "pack-uid-here"}`

**search_packs** — Search available packs with filters. ALL filter values MUST be arrays.
`{"mode": "search_packs", "filter": {"type": ["spectro"], "layer": ["os"], "environment": ["edge-native"]}}`
*Filter fields (all values are arrays):*
*- type: ["spectro"], ["helm"], ["manifest"], ["oci"] — pack source type*
*- layer: ["os"], ["k8s"], ["cni"], ["csi"], ["addon"] — pack layer*
*- environment: ["aws"], ["eks"], ["azure"], ["edge-native"], etc.*
*- addOnType: ["monitoring"], ["logging"], ["security"], ["authentication"], ["ingress"], ["load balancer"]*
*- addOnSubType: ["db"], ["cache"], etc.*
*Monitoring packs:* `{"mode": "search_packs", "filter": {"addOnType": ["monitoring"]}}`
*Logging packs:* `{"mode": "search_packs", "filter": {"addOnType": ["logging"]}}`
*Security packs:* `{"mode": "search_packs", "filter": {"addOnType": ["security"]}}`
*Edge OS packs:* `{"mode": "search_packs", "filter": {"type": ["spectro"], "layer": ["os"], "environment": ["edge-native"]}}`
*Search by name:* `{"mode": "search_packs", "filter": {"name": {"contains": "prometheus"}}}`
*Exact name match:* `{"mode": "search_packs", "filter": {"name": {"eq": "prometheus"}}}`
*Name filter uses object (not array): {"eq": "value"}, {"contains": "value"}, {"beginsWith": "value"}*

