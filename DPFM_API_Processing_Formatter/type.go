package dpfm_api_processing_formatter

type GeneralUpdates struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
	StorageBin      string `json:"StorageBin"`
}
