package client

type GSBucketClient struct {
	baseUrl string
}

func CreateGSBucketClient(baseUrl string) *GSBucketClient {
	return &GSBucketClient{
		baseUrl: baseUrl,
	}
}

// func (c *GSBucketClient) PostData(fileName string, data []string, ttl time.Duration) (responses.PostResponse, error) {

// }

// func (c *GSBucketClient) GetFiles() ([]responses.PostResponse, error) {

// }

// func (c *GSBucketClient) Delete() error{

// }
