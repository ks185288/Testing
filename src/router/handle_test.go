package router

import (
	"Testing/src/router"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	handler router
	contex  gin.Context
}

func (suite *TestSuite) SetupTest() {
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	handler := Handler{}
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestHandlerMessage() {
	suite.contex.Request.Method = "POST" // or PUT

	//suite.contex.Request.Header.Set("Content-Type", "application/json")

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	suite.context.Request.Body = io.NopCloser(bytes.NewBuffer("message"))
	suite.handler.HandleMessage(suite.contex)
	suite.NotNil(suite.handler.Response)

}
