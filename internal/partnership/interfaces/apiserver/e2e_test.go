package apiserver_test

import (
	"github.com/UsadaPeko/jsn"
	"gomod.pekora.dev/tabasco/internal/partnership/interfaces/apiserver"
	"io/ioutil"

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
			It("Return partnership's id", func() {
				response, err := app.Test(request)
				Expect(err).Should(Succeed())

				body, err := ioutil.ReadAll(response.Body)
				Expect(err).Should(Succeed())

				responseBody, err := jsn.New(string(body))
				Expect(err).Should(Succeed())

				_, ok := responseBody.StringVal("id")
				Expect(ok).Should(BeTrue())
			})
		})

	})
})
