package database

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"

	svcErrors "github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/errors"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/shopper_service/pkg/graphql_api/model"
)

// GetShopperAccountByQueryParam gets a shopper account by query parameter from the backend database
// This query is performed against a certain field present in the respective backend database table
func (db *Db) GetShopperAccountByQueryParam(ctx context.Context, queryField string, queryParam interface{}) *model.ShopperAccount {
	db.Logger.For(ctx).Info(fmt.Sprintf("get shopper account by %s", queryParam))
	ctx, span := db.startRootSpan(ctx, fmt.Sprintf("get_shopper_account_by_%s_op", queryParam))
	defer span.Finish()

	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		conn := db.Conn.Engine

		db.Logger.For(ctx).Info("starting db transactions")
		childSpan := db.TracingEngine.CreateChildSpan(ctx, fmt.Sprintf("get_shopper_account_by_%s_tx", queryParam))
		defer childSpan.Finish()

		var shopperAccountORM model.ShopperAccountORM

		// attempt to see if the record already exists
		recordNotFoundErr := conn.Where(map[string]interface{}{queryField: queryParam}).First(&shopperAccountORM).Error
		if recordNotFoundErr != nil {
			db.Logger.For(ctx).Error(svcErrors.ErrAccountDoesNotExist, "account does not exist")
			return nil, svcErrors.ErrAccountDoesNotExist
		}

		// transform orm type to account type
		account, err := shopperAccountORM.ToPB(ctx)
		if err != nil {
			db.Logger.For(ctx).Error(svcErrors.ErrFailedToConvertFromOrmType, err.Error())
			return nil, err
		}

		db.Logger.For(ctx).Info("successfully obtained business account", zap.Any(queryField, queryParam))
		return &account, nil
	}

	res, err := db.Conn.PerformComplexTransaction(ctx, tx)
	if err != nil {
		return nil
	}

	return res.(*model.ShopperAccount)
}
