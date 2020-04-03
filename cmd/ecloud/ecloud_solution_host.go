package ecloud

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/factory"
	"github.com/ukfast/cli/internal/pkg/helper"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
)

func ecloudSolutionHostRootCmd(f factory.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "host",
		Short: "sub-commands relating to solution hosts",
	}

	// Child commands
	cmd.AddCommand(ecloudSolutionHostListCmd(f))

	return cmd
}

func ecloudSolutionHostListCmd(f factory.ClientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "Lists solution hosts",
		Long:    "This command lists solution hosts",
		Example: "ukfast ecloud solution host list 123",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing solution")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ecloudSolutionHostList(f.NewClient().ECloudService(), cmd, args)
		},
	}
}

func ecloudSolutionHostList(service ecloud.ECloudService, cmd *cobra.Command, args []string) error {
	solutionID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid solution ID [%s]", args[0])
	}

	params, err := helper.GetAPIRequestParametersFromFlags(cmd)
	if err != nil {
		return err
	}

	hosts, err := service.GetSolutionHosts(solutionID, params)
	if err != nil {
		return fmt.Errorf("Error retrieving solution hosts: %s", err)
	}

	return output.CommandOutput(cmd, OutputECloudHostsProvider(hosts))
}