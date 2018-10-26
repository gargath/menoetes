package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/gargath/menoetes/pkg/server/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDisco(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Disco Suite")
}

var _ = Describe("Discovery", func() {
	var server *httptest.Server
	var httpClient *http.Client
	log.SetOutput(ioutil.Discard)
	BeforeEach(func() {
		r := mux.NewRouter()
		r.HandleFunc("/.well-known/{disco_path}", m.Use(discoHandler))
		server = httptest.NewServer(r)
		httpClient = &http.Client{}
	})

	AfterEach(func() {
		server.Close()
	})

	It("Response to service discovery", func() {
		req, _ := http.NewRequest("GET", server.URL+"/.well-known/terraform.json", nil)
		resp, err := httpClient.Do(req)
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(string(bodyBytes)).To(Equal(fmt.Sprintf("{\n\"modules.v1\": \"https://%s/v1/modules\"\n}", req.Host)))
		Expect(err).NotTo(HaveOccurred())
		Expect(err2).NotTo(HaveOccurred())
	})

	It("Response 404 to unknown paths", func() {
		req, _ := http.NewRequest("GET", server.URL+"/.well-known/foo.json", nil)
		resp, err := httpClient.Do(req)
		Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		Expect(err).NotTo(HaveOccurred())
	})

})
