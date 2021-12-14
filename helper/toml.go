package helper

import (
	"bytes"
	"text/template"

	"github.com/spf13/viper"
	cmn "github.com/tendermint/tendermint/libs/common"
)

// Note: any changes to the comments/variables/mapstructure
// must be reflected in the appropriate struct in helper/config.go
const defaultConfigTemplate = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### RPC and REST configs #####

# RPC endpoint for ethereum chain
eth_rpc_url = "{{ .EthRPCUrl }}"

# RPC endpoint for bsc chain
bsc_rpc_url = "{{ .BscRPCUrl }}"

# RPC endpoint for pano chain
pano_rpc_url = "{{ .PanoRPCUrl }}"

# RPC endpoint for tendermint
tendermint_rpc_url = "{{ .TendermintRPCUrl }}"

# Delivery REST server endpoint
delivery_rest_server = "{{ .DeliveryServerURL }}"

# RPC endpoint for tron
tron_rpc_url = "{{ .TronRPCUrl }}"
tron_grid_url = "{{ .TronGridUrl }}"
tron_grid_api_key = "{{ .TronGridApiKey }}"

#### Bridge configs ####

# AMQP endpoint
amqp_url = "{{ .AmqpURL }}"

## Poll intervals
checkpoint_poll_interval = "{{ .CheckpointerPollInterval }}"
eth_syncer_poll_interval = "{{ .EthSyncerPollInterval }}"
bsc_syncer_poll_interval = "{{ .BscSyncerPollInterval }}"
tron_syncer_poll_interval = "{{ .TronSyncerPollInterval }}"
noack_poll_interval = "{{ .NoACKPollInterval }}"
clerk_poll_interval = "{{ .ClerkPollInterval }}"
span_poll_interval = "{{ .SpanPollInterval }}"
staking_poll_interval = "{{ .StakingPollInterval }}"

#### gas limits ####
main_chain_gas_limit = "{{ .MainchainGasLimit }}"
tron_chain_fee_limit = "{{ .TronchainFeeLimit }}"

#### busy limits ####
eth_unconfirmed_txs_busy_limit = "{{ .EthUnconfirmedTxsBusyLimit }}"
bsc_unconfirmed_txs_busy_limit = "{{ .BscUnconfirmedTxsBusyLimit }}"
tron_unconfirmed_txs_busy_limit = "{{ .TronUnconfirmedTxsBusyLimit }}"

eth_max_query_blocks = "{{ .EthMaxQueryBlocks }}"
bsc_max_query_blocks = "{{ .BscMaxQueryBlocks }}"
tron_max_query_blocks = "{{ .TronMaxQueryBlocks }}"

##### Timeout Config #####
no_ack_wait_time = "{{ .NoACKWaitTime }}"

`

var configTemplate *template.Template

func init() {
	var err error
	tmpl := template.New("appConfigFileTemplate")
	if configTemplate, err = tmpl.Parse(defaultConfigTemplate); err != nil {
		panic(err)
	}
}

// ParseConfig retrieves the default environment configuration for the
// application.
func ParseConfig() (*Configuration, error) {
	conf := GetDefaultHeimdallConfig()
	err := viper.Unmarshal(conf)
	return &conf, err
}

// WriteConfigFile renders config using the template and writes it to
// configFilePath.
func WriteConfigFile(configFilePath string, config *Configuration) {
	var buffer bytes.Buffer

	if err := configTemplate.Execute(&buffer, config); err != nil {
		panic(err)
	}

	cmn.MustWriteFile(configFilePath, buffer.Bytes(), 0644)
}
