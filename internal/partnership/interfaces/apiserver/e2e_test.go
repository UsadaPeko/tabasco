package apiserver_test

import (
	"github.com/UsadaPeko/jsn"
	"github.com/google/uuid"
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
			response, err := app.Test(request)
			It("Return 201", func() {
				Expect(response.StatusCode).Should(Equal(http.StatusCreated))
				Expect(err).Should(Succeed())
			})
			It("Return partnership's id", func() {
				body, err := ioutil.ReadAll(response.Body)
				Expect(err).Should(Succeed())

				responseBody, err := jsn.New(string(body))
				Expect(err).Should(Succeed())

				_, ok := responseBody.StringVal("id")
				Expect(ok).Should(BeTrue())
			})
		})

		When("Already made partnership", func() {
			jsonObject := map[string]interface{}{
				"name": "Tabasco",
			}
			requestBody, err := json.Marshal(jsonObject)
			Expect(err).Should(Succeed())

			request := httptest.NewRequest("POST", "/partnership", bytes.NewBuffer(requestBody))
			response, err := app.Test(request)
			Expect(response.StatusCode).Should(Equal(http.StatusCreated))
			Expect(err).Should(Succeed())

			body, err := ioutil.ReadAll(response.Body)
			Expect(err).Should(Succeed())

			responseBody, err := jsn.New(string(body))
			Expect(err).Should(Succeed())

			id, ok := responseBody.StringVal("id")
			Expect(ok).Should(BeTrue())

			Context("Call GET root/partnership/{id}", func() {
				request = httptest.NewRequest("GET", "/partnership/"+id, nil)
				response, err = app.Test(request)
				It("Return 200", func() {
					Expect(response.StatusCode).Should(Equal(http.StatusOK))
					Expect(err).Should(Succeed())
				})
				It("Can found information", func() {
					body, err := ioutil.ReadAll(response.Body)
					Expect(err).Should(Succeed())

					responseBody, err := jsn.New(string(body))
					Expect(err).Should(Succeed())

					foundID, ok := responseBody.StringVal("id")
					Expect(ok).Should(BeTrue())
					Expect(foundID).Should(Equal(id))

					name, ok := responseBody.StringVal("name")
					Expect(ok).Should(BeTrue())
					Expect(name).Should(Equal("Tabasco"))
				})
			})
			Context("Call POST root/partnership/{id}/integrations", func() {
				jsonObject := map[string]interface{}{
					"type": "custom_event",
				}
				requestBody, err := json.Marshal(jsonObject)
				Expect(err).Should(Succeed())

				request = httptest.NewRequest("POST", "/partnership/"+id+"/integrations", bytes.NewBuffer(requestBody))
				response, err = app.Test(request)
				It("Return 201", func() {
					Expect(response.StatusCode).Should(Equal(http.StatusCreated))
					Expect(err).Should(Succeed())
				})
				It("Can found integration key", func() {
					body, err := ioutil.ReadAll(response.Body)
					Expect(err).Should(Succeed())

					responseBody, err := jsn.New(string(body))
					Expect(err).Should(Succeed())

					key, ok := responseBody.StringVal("key")
					Expect(ok).Should(BeTrue())
					Expect(key).Should(Not(BeEmpty()))
				})
			})
		})

		Context("Call GET root/partnership/{id}", func() {
			When("With not exist id", func() {
				request := httptest.NewRequest("GET", "/partnership/"+uuid.NewString(), nil)
				response, err := app.Test(request)
				It("Return 404", func() {
					Expect(response.StatusCode).Should(Equal(http.StatusNotFound))
					Expect(err).Should(Succeed())
				})
			})
		})

	})
})
