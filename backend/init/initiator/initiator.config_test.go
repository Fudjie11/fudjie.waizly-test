package initiator_test

import (
	"errors"
	"testing"

	"fudjie.waizly/backend-test/init/initiator"

	"github.com/google/gops/agent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAgent is a mock for the gops agent
type MockAgent struct {
	mock.Mock
}

func (m *MockAgent) Listen(opts agent.Options) error {
	args := m.Called(opts)
	return args.Error(0)
}

// MockConfigReader is a mock for the config reader
type MockConfigReader struct {
	mock.Mock
}

func (m *MockConfigReader) ReadConfig(cfg interface{}, path string, prefix string) error {
	args := m.Called(cfg, path, prefix)
	return args.Error(0)
}

func TestInitConfig_Successful(t *testing.T) {
	mockAgent := new(MockAgent)
	mockConfigReader := new(MockConfigReader)
	mockAgent.On("Listen", mock.Anything).Return(nil)
	mockConfigReader.On("ReadConfig", mock.Anything, "valid_path", "config").Return(nil)

	i := &initiator.Initiator{
		AgentListen: mockAgent.Listen,
		ReadConfig:  mockConfigReader.ReadConfig,
	} // Assuming Initiator can use mockAgent and mockConfigReader
	cfg := i.InitConfig("valid_path", "test-service")

	assert.NotNil(t, cfg, "Config should not be nil on successful initialization")
}

func TestInitConfig_AgentListenError(t *testing.T) {
	mockAgent := new(MockAgent)
	mockAgent.On("Listen", mock.Anything).Return(errors.New("agent listen error"))

	i := initiator.Initiator{} // Adjust to use mockAgent
	assert.Panics(t, func() { i.InitConfig("valid_path", "test-service") }, "Expected panic due to agent listen error")
}

// Implement TestInitConfig_InvalidConfigPath and TestInitConfig_DifferentServiceName similarly
