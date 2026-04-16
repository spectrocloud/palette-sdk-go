package clusters

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) InitiateLogDownload(uid string, req *models.V1ClusterLogFetcherRequest) (*string, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.InitiateDownloadOfClusterLogs(uid, req)
}

func (s *service) GetLogFetcherStatus(uid string, logFetcherUID *string) (*models.V1ClusterLogFetcher, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetLogFetcherStatus(uid, logFetcherUID)
}
