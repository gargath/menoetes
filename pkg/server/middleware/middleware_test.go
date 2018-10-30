package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gargath/menoetes/pkg/store"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testResponder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\n\"test\": \"success\"\n}")
}

func TestMiddlware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Middleware Suite")
}

var _ = Describe("Middleware", func() {
	var server *httptest.Server
	var httpClient *http.Client
	log.SetOutput(ioutil.Discard)
	BeforeEach(func() {
		r := mux.NewRouter()
		s := store.NewMockStore()
		tm := &TokenManager{
			Store: s,
		}
		r.HandleFunc("/foo", Use(testResponder, tm.TokenAuth))
		server = httptest.NewServer(r)
		httpClient = &http.Client{}
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("TokenAuth", func() {
		It("Allows valid tokens", func() {
			req, _ := http.NewRequest("GET", server.URL+"/foo", nil)
			req.Header.Set("Authorization", "Bearer reallylongstringthatstotallygoingtostandoutinthelistofheaders")
			resp, err := httpClient.Do(req)
			bodyBytes, err2 := ioutil.ReadAll(resp.Body)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(string(bodyBytes)).To(Equal("{\n\"test\": \"success\"\n}"))
			Expect(err).NotTo(HaveOccurred())
			Expect(err2).NotTo(HaveOccurred())
		})

		It("Rejects anonymous requests", func() {
			resp, err := httpClient.Get(server.URL + "/foo")
			bodyBytes, err2 := ioutil.ReadAll(resp.Body)
			Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))
			Expect(string(bodyBytes)).To(Equal("anonymous access not allowed\n"))
			Expect(err).NotTo(HaveOccurred())
			Expect(err2).NotTo(HaveOccurred())
		})

		It("Rejects invalid tokens", func() {
			req, _ := http.NewRequest("GET", server.URL+"/foo", nil)
			req.Header.Set("Authorization", "Bearer illegal")
			resp, err := httpClient.Do(req)
			bodyBytes, err2 := ioutil.ReadAll(resp.Body)
			Expect(resp.StatusCode).To(Equal(http.StatusForbidden))
			Expect(string(bodyBytes)).To(Equal("access denied\n"))
			Expect(err).NotTo(HaveOccurred())
			Expect(err2).NotTo(HaveOccurred())
		})

	})

})
