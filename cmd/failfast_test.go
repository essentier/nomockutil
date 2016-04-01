package cmd_test

import (
	"testing"

	"github.com/essentier/nomockutil/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FailFastCmdTestSuite struct {
	suite.Suite
	cmdRunner cmd.CmdRunner
}

func (suite *FailFastCmdTestSuite) SetupTest() {
	suite.cmdRunner = cmd.CreateFailFastCmdRunner(cmd.CreateCmdConsole())
}

func (suite *FailFastCmdTestSuite) TestValidCmd() {
	suite.cmdRunner.RunCmd("ls")
	assert.False(suite.T(), suite.cmdRunner.HasError(), "command ls should not fail.")
}

func (suite *FailFastCmdTestSuite) TestNonExistentCmd() {
	suite.cmdRunner.RunCmd(nonexistent_cmd)
	assert.True(suite.T(), suite.cmdRunner.HasError(), "nonexistent command should fail.")
}

func (suite *FailFastCmdTestSuite) TestValidCmdAfterInvalidCmd() {
	suite.cmdRunner.RunCmd(nonexistent_cmd)
	suite.cmdRunner.RunCmd("ls")
	assert.True(suite.T(), suite.cmdRunner.HasError(), "command ls should fail after an invalid command.")
}

func (suite *FailFastCmdTestSuite) TestNewRunner() {
	assert := assert.New(suite.T())
	assert.False(suite.cmdRunner.HasError(), "Runner should not fail before any command is executed.")

	newRunner := suite.cmdRunner.NewRunner()
	assert.False(suite.cmdRunner.HasError(), "New runner should not fail before any command is executed.")

	newRunner.RunCmd(nonexistent_cmd)
	assert.True(newRunner.HasError(), "command executed by new runner should fail.")
	assert.False(suite.cmdRunner.HasError(), "Command in new runner should not affect the original runner.")
}

func TestFailFastCmdTestSuite(t *testing.T) {
	suite.Run(t, new(FailFastCmdTestSuite))
}
