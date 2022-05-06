package apiserver_test

import (
	"gomod.pekora.dev/tabasco/internal/eventadapter/interfaces/apiserver"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("API Server", func() {
	app := apiserver.MakeServer()

	Context("Call GET root/", func() {
		request := httptest.NewRequest("GET", "/", nil)
		It("Return 200", func() {
			response, err := app.Test(request)
			Expect(response.StatusCode).Should(Equal(http.StatusOK))
			Expect(err).Should(Succeed())
		})
	})
})
