// Copyright 2025 Amazon Ads API Go SDK Authors.
//
// This file is part of Amazon Ads API Go SDK.
//
// Amazon Ads API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license. All rights reserved.

package main

import (
	"fmt"
	"log"

	"github.com/vanling1111/amazon-ads-api-go-sdk/internal/models"
	"github.com/vanling1111/amazon-ads-api-go-sdk/pkg/adsapi"
)

func main() {
	// 使用 Functional Options 模式创建客户端
	client, err := adsapi.NewClient(
		adsapi.WithRegion(models.RegionNorthAmerica),
		adsapi.WithCredentials(
			"your-client-id",
			"your-client-secret",
			"your-refresh-token",
		),
		adsapi.WithProfileID(123456789),
		adsapi.WithDebug(),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Amazon Ads API Client created successfully!")
	fmt.Printf("Region: %s\n", client.GetConfig().Region.Name)
	fmt.Printf("Profile ID: %d\n", client.GetConfig().ProfileID)
}
