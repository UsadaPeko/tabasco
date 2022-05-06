package apiserver_test

import (
	"gomod.pekora.dev/tabasco/internal/eventadapter/interfaces/apiserver"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"bytes"
	"encoding/json"
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

	Context("Call POST root/service/{ServiceKey}/event", func() {
		When("Without user id", func() {
			jsonObject := map[string]interface{}{
				"event_id":   "c5d7c94f-0fcf-4791-bc45-06d8fecbd356",
				"event_name": "login",
				"event_time": "2022-01-01T00:00:00Z",
			}
			requestBody, err := json.Marshal(jsonObject)
			Expect(err).Should(Succeed())

			request := httptest.NewRequest("POST", "/service/c5d7c94f/event", bytes.NewBuffer(requestBody))
			It("Return 400", func() {
				response, err := app.Test(request)
				Expect(response.StatusCode).Should(Equal(http.StatusBadRequest))
				Expect(err).Should(Succeed())
			})
		})
		When("With login event", func() {
			jsonObject := map[string]interface{}{
				"event_id":   "c5d7c94f-0fcf-4791-bc45-06d8fecbd356",
				"event_name": "login",
				"event_time": "2022-01-01T00:00:00Z",

				"user_id": "usada pekora",
			}
			requestBody, err := json.Marshal(jsonObject)
			Expect(err).Should(Succeed())

			request := httptest.NewRequest("POST", "/service/c5d7c94f/event", bytes.NewBuffer(requestBody))
			It("Return 200", func() {
				response, err := app.Test(request)
				Expect(response.StatusCode).Should(Equal(http.StatusOK))
				Expect(err).Should(Succeed())
			})
		})
	})

	Context("Call POST root/service/{ServiceKey}/event with Invalid Service Key", func() {
		invalidServiceKey := "InvalidServiceKey"
		jsonObject := map[string]interface{}{
			"event_id":   "c5d7c94f-0fcf-4791-bc45-06d8fecbd356",
			"event_name": "login",
			"event_time": "2022-01-01T00:00:00Z",

			"user_id": "usada pekora",
		}
		requestBody, err := json.Marshal(jsonObject)
		Expect(err).Should(Succeed())

		request := httptest.NewRequest("POST", "/service/"+invalidServiceKey+"/event", bytes.NewBuffer(requestBody))
		It("Return 401", func() {
			response, err := app.Test(request)
			Expect(response.StatusCode).Should(Equal(http.StatusUnauthorized))
			Expect(err).Should(Succeed())
		})
	})
})
