package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-storage-bin-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToGeneralUpdates(general dpfm_api_input_reader.General) *GeneralUpdates {
	data := general

	return &GeneralUpdates{
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
		StorageLocation: data.StorageLocation,
		StorageBin:      data.StorageBin,
	}
}
