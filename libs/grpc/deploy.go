package grpc

import (
	"context"
	"time"

	pbDeploy "github.com/es-hs/erpc/deploy"
	esUtils "github.com/es-hs/es-helper/utils"
)

func DeployServiceDeployShop(shopDomain, region, sourceBucket, sourceVersion, customerDomain string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := GRPCConn.DeployService.Client.DeployShop(ctx, &pbDeploy.DeployShopRequest{
		ShopDomain:    shopDomain,
		Region:        esUtils.DefaultValueString(region, "ap-southeast-1"),
		SourceBucket:  esUtils.DefaultValueString(sourceBucket, "storefront-code-ap-southeast-1"),
		SourceVersion: sourceVersion,
		CustomDomain:  customerDomain,
	})
	return err
}
