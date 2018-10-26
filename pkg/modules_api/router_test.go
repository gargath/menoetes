package modules_api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRouter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Router Suite")
}

var server *httptest.Server
var httpClient *http.Client

func makeAuthorizedRequest(method string, url string) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization", "Bearer reallylongstringthatstotallygoingtostandoutinthelistofheaders")
	return httpClient.Do(req)
}

var _ = Describe("Modules Router", func() {
	log.SetOutput(ioutil.Discard)
	BeforeEach(func() {
		r := mux.NewRouter()
		RegisterModulesAPI(r)
		server = httptest.NewServer(r)
		httpClient = &http.Client{}
	})

	AfterEach(func() {
		server.Close()
	})

	It("Response to modules URL", func() {
		resp, err := makeAuthorizedRequest("GET", server.URL+"/v1/modules")
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(err).NotTo(HaveOccurred())
		resp, err = makeAuthorizedRequest("GET", server.URL+"/v1/modules/somenamespace")
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(err).NotTo(HaveOccurred())
	})

	It("Response 404 to unknown paths", func() {
		req, _ := http.NewRequest("GET", server.URL+"/foo", nil)
		req.Header.Set("Authorization", "Bearer reallylongstringthatstotallygoingtostandoutinthelistofheaders")
		resp, err := httpClient.Do(req)
		Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		Expect(err).NotTo(HaveOccurred())
	})

})
