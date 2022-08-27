//go:build example
// +build example

package examples

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/cute"
)

func TestExample_TwoSteps_Example_1(t *testing.T) {
	cute.NewTestBuilder().
		Title("TestExample_TwoSteps_Example_1").
		Tags("TestExample_TwoSteps_Example_1", "some_tag").
		Parallel().
		CreateStep("Creat entry /posts/1").

		// CreateWithStep first step

		RequestBuilder(
			cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusCreated).
		NextTest().
		CreateStep("Delete entry").

		// CreateWithStep second step for delete
		RequestBuilder(
			cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
			cute.WithMethod(http.MethodDelete),
			cute.WithHeaders(map[string][]string{
				"some_auth_token": []string{fmt.Sprint(11111)},
			}),
		).
		ExecuteTest(context.Background(), t)
}

func TestExample_TwoSteps_Example_2(t *testing.T) {
	runner.Run(t, "Test with two steps", func(t provider.T) {
		test := cute.NewTestBuilder().
			Title("Two steps").
			Description("some_description").
			CreateStep("Request 1").
			RequestBuilder(
				cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
				cute.WithMethod(http.MethodGet),
			).
			ExpectStatus(http.StatusOK).
			ExecuteTest(context.Background(), t)

		bodyBytes, err := io.ReadAll(test[0].GetHTTPResponse().Body)
		if err != nil {
			log.Fatal(err)
		}
		// process body
		_ = string(bodyBytes)

		cute.NewTestBuilder().
			CreateStep("Request 2").
			RequestBuilder(
				cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
				cute.WithMethod(http.MethodGet),
			).
			ExpectExecuteTimeout(10*time.Second).
			ExpectStatus(http.StatusOK).
			ExecuteTest(context.Background(), t)
	})
}

func TestExample_TwoSteps_Example_3(t *testing.T) {
	responseCode := 0

	cute.NewTestBuilder().
		Create().
		RequestBuilder(
			cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectStatus(http.StatusOK).
		NextTest().
		AfterTestExecute(func(response *http.Response, errors []error) error {
			responseCode = response.StatusCode

			return nil
		}).
		Create().
		RequestBuilder(
			cute.WithURI("https://jsonplaceholder.typicode.com/posts/2/comments"),
			cute.WithMethod(http.MethodDelete),
		).
		ExecuteTest(context.Background(), t)

	fmt.Println("Response code from first request", responseCode)
}
