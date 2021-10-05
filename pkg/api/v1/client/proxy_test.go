package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kubeshop/kubtest/pkg/api/v1/kubtest"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type Rest struct {
}

func (r *Rest) GetRateLimiter() flowcontrol.RateLimiter {
	return flowcontrol.NewFakeAlwaysRateLimiter()
}
func (r *Rest) Verb(verb string) *rest.Request {
	return &rest.Request{}
}
func (r *Rest) Post() *rest.Request {
	return &rest.Request{}

}
func (r *Rest) Put() *rest.Request {
	return &rest.Request{}

}
func (r *Rest) Patch(pt types.PatchType) *rest.Request {
	return &rest.Request{}

}
func (r *Rest) Get() *rest.Request {
	return &rest.Request{}

}
func (r *Rest) Delete() *rest.Request {
	return &rest.Request{}

}
func (r *Rest) APIVersion() schema.GroupVersion {
	return schema.GroupVersion{Group: "api", Version: "v1"}
}

func TestDefaultDirectScriptsAPI(t *testing.T) {
	t.Skip("Implement working test")

	k8sClient := fake.NewSimpleClientset()
	// can't override REST client to change requested URI
	// k8sClient.CoreV1().RESTCli nt()
	config := NewProxyConfig("default")
	client := NewProxyScriptsAPI(k8sClient, config)

	t.Run("Execute script with given ID", func(t *testing.T) {
		// given
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, `{"id":"1", "executionResult":{"status": "success", "output":"execution completed"}}`)
		}))
		defer srv.Close()

		// when
		execution, err := client.ExecuteScript("test", "default", "some name", map[string]string{})

		// then
		assert.Equal(t, "1", execution.Id)
		assert.Equal(t, kubtest.SUCCESS_ExecutionStatus, *execution.ExecutionResult.Status)
		assert.Equal(t, "execution completed", execution.ExecutionResult.Output)
		assert.NoError(t, err)
	})

}
