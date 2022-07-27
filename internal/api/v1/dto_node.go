package v1

type NodeDto struct {
	PublicKey      string                 `json:"publicKey"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo"`
}
