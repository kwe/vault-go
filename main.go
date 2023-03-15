package main

import (
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	keys, err := getKeysFromVault("http://127.0.0.1:8200", "vaulttokenhere", "cubbyhole/customer1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Keys: %+v\n", keys)
}

func getKeysFromVault(vaultAddress, vaultToken, path string) (map[string]interface{}, error) {
	// Initialize the Vault client
	config := &vault.Config{
		Address: vaultAddress,
	}

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create vault client: %v", err)
	}

	// Set the Vault token
	client.SetToken(vaultToken)

	// Read the data from the specified path
	secret, err := client.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from vault: %v", err)
	}

	if secret == nil {
		return nil, fmt.Errorf("no data found at path: %s", path)
	}

	// Return the downloaded keys
	return secret.Data, nil
}
