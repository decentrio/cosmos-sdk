package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
)

const chainUpgradeGuide = "https://github.com/cosmos/cosmos-sdk/blob/main/UPGRADING.md"

// ValidateGenesisCmd takes a genesis file, and makes sure that it is valid.
func ValidateGenesisCmd(mm *module.Manager) *cobra.Command {
	return &cobra.Command{
		Use:     "validate [file]",
		Aliases: []string{"validate-genesis"},
		Args:    cobra.RangeArgs(0, 1),
		Short:   "Validates the genesis file at the default location or at the location passed as an arg",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			serverCtx := server.GetServerContextFromCmd(cmd)
			cfg, ok := serverCtx.GetConfig().(server.CometConfig)
			if !ok {
				return fmt.Errorf("Can not convert cometbft config")
			}

			// Load default if passed no args, otherwise load passed file
			var genesis string
			if len(args) == 0 {
				genesis = cfg.GenesisFile()
			} else {
				genesis = args[0]
			}

			appGenesis, err := types.AppGenesisFromFile(genesis)
			if err != nil {
				return err
			}

			if err := appGenesis.ValidateAndComplete(); err != nil {
				return fmt.Errorf("make sure that you have correctly migrated all CometBFT consensus params. Refer the UPGRADING.md (%s): %w", chainUpgradeGuide, err)
			}

			var genState map[string]json.RawMessage
			if err = json.Unmarshal(appGenesis.AppState, &genState); err != nil {
				return fmt.Errorf("error unmarshalling genesis doc %s: %w", genesis, err)
			}

			if mm != nil {
				if err = mm.ValidateGenesis(genState); err != nil {
					return fmt.Errorf("error validating genesis file %s: %w", genesis, err)
				}
			}

			fmt.Fprintf(cmd.OutOrStdout(), "File at %s is a valid genesis file\n", genesis)
			return nil
		},
	}
}
