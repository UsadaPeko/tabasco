package apiserver_test

import (
	"gomod.pekora.dev/tabasco/internal/partnership/interfaces/apiserver"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"

	"bytes"
	"encoding/json"
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

	Context("Call POST root/partnership", func() {
		When("With basic information", func() {
			jsonObject := map[string]interface{}{
				"name": "Tabasco",
			}
			requestBody, err := json.Marshal(jsonObject)
			Expect(err).Should(Succeed())

			request := httptest.NewRequest("POST", "/partnership", bytes.NewBuffer(requestBody))
			It("Return 201", func() {
				response, err := app.Test(request)
				Expect(response.StatusCode).Should(Equal(http.StatusCreated))
				Expect(err).Should(Succeed())
			})
		})

	})
})
