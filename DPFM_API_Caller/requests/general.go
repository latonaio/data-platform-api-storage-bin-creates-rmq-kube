package requests

type General struct {
	BusinessPartner       int      `json:"BusinessPartner"`
	Plant                 string   `json:"Plant"`
	StorageLocation       string   `json:"StorageLocation"`
	StorageBin            string   `json:"StorageBin"`
	StorageBinDescription *string  `json:"StorageType"`
	XCoordinates          *float32 `json:"XCoordinates"`
	YCoordinates          *float32 `json:"YCoordinates"`
	ZCoordinates          *float32 `json:"ZCoordinates"`
	Capacity              *float32 `json:"Capacity"`
	CapacityUnit          *string  `json:"CapacityUnit"`
	CapacityWeight        *float32 `json:"CapacityWeight"`
	CapacityWeightUnit    *string  `json:"CapacityWeightUnit"`
	CreationDate          string   `json:"CreationDate"`
	LastChangeDate        string   `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool    `json:"IsMarkedForDeletion"`
}
