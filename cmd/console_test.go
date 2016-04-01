package cmd_test

import (
	"testing"

	"github.com/essentier/nomockutil/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	nonexistent_cmd string = "nonexistent_cmd"
)

type ConsoleCmdTestSuite struct {
	suite.Suite
	cmdRunner cmd.CmdRunner
}

func (suite *ConsoleCmdTestSuite) SetupTest() {
	suite.cmdRunner = cmd.CreateCmdConsole()
}

func (suite *ConsoleCmdTestSuite) TestValidCmd() {
	suite.cmdRunner.RunCmd("ls")
	assert.False(suite.T(), suite.cmdRunner.HasError(), "command ls should not fail.")
}

func (suite *ConsoleCmdTestSuite) TestNonExistentCmd() {
	suite.cmdRunner.RunCmd(nonexistent_cmd)
	assert.True(suite.T(), suite.cmdRunner.HasError(), "nonexistent command should fail.")
}

func (suite *ConsoleCmdTestSuite) TestValidCmdAfterInvalidCmd() {
	suite.cmdRunner.RunCmd(nonexistent_cmd)
	suite.cmdRunner.RunCmd("ls")
	assert.False(suite.T(), suite.cmdRunner.HasError(), "command ls should not fail.")
}

func (suite *ConsoleCmdTestSuite) TestNewRunner() {
	assert := assert.New(suite.T())
	assert.False(suite.cmdRunner.HasError(), "Runner should not fail before any command is executed.")

	newRunner := suite.cmdRunner.NewRunner()
	assert.False(suite.cmdRunner.HasError(), "New runner should not fail before any command is executed.")

	newRunner.RunCmd(nonexistent_cmd)
	assert.True(newRunner.HasError(), "command executed by new runner should fail.")
	assert.False(suite.cmdRunner.HasError(), "Command in new runner should not affect the original runner.")
}

func TestConsoleCmdTestSuite(t *testing.T) {
	suite.Run(t, new(ConsoleCmdTestSuite))
}
