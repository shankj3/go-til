package nsqpb

import (
	"fmt"
	"net/http"
)

//LookupTopics goes to nsqd and checks to see if topic is supported
func LookupTopic(nsqdLookupHostPort string, topic string) bool {
	nsqdLookupAddr := fmt.Sprintf("http://%s/lookup?topic=%s", nsqdLookupHostPort, topic)
	resp, err := http.Get(nsqdLookupAddr)
	if err != nil {
		return false
	}
	if resp.StatusCode == http.StatusNotFound {
		return false
	}
	// todo: check for other headers?
	return true
}
