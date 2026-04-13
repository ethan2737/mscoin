package handler

import (
	"net/http"

	"swap-api/internal/midd"
	"swap-api/internal/svc"
)

func RegisterHandlers(r *Routers, svcCtx *svc.ServiceContext) {
	swapGroup := r.Group()
	swapGroup.Get("/swap/contract-coin/info/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewContractInfoHandler(svcCtx).Info(w, r)
	}))
	swapGroup.Get("/swap/symbol-thumb", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewContractInfoHandler(svcCtx).List(w, r)
	}))
	swapGroup.Get("/swap/exchange-plate-mini", exchangePlateMiniHandler(svcCtx))
	swapGroup.Get("/swap/exchange-plate-full", exchangePlateFullHandler(svcCtx))
	swapGroup.Get("/swap/latest-trade", latestTradeHandler(svcCtx))

	orderGroup := r.Group()
	orderGroup.Use(midd.Auth(svcCtx.Config.JWT.AccessSecret))
	orderGroup.Post("/swap/order-entrust/add", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svcCtx).Add(w, r)
	}))
	orderGroup.Post("/swap/order-entrust/cancel/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svcCtx).Cancel(w, r)
	}))
	orderGroup.Post("/swap/order-entrust/current", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svcCtx).Current(w, r)
	}))
	orderGroup.Post("/swap/order-entrust/history", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svcCtx).History(w, r)
	}))
	orderGroup.Post("/swap/order/position", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svcCtx).Position(w, r)
	}))
	orderGroup.Post("/swap/order-entrust/quick-close/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svcCtx).QuickClose(w, r)
	}))
	orderGroup.Post("/swap/contract-leverage", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svcCtx).Leverage(w, r)
	}))

	ucGroup := r.Group()
	ucGroup.Use(midd.Auth(svcCtx.Config.JWT.AccessSecret))
	ucGroup.Post("/uc/contract/current", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewUcContractHandler(svcCtx).Current(w, r)
	}))
	ucGroup.Post("/uc/contract/history", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewUcContractHandler(svcCtx).History(w, r)
	}))
	ucGroup.Post("/uc/contract/coin/thumb", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewUcContractHandler(svcCtx).Thumb(w, r)
	}))
	ucGroup.Post("/uc/contract/entrust/cancel/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewUcContractHandler(svcCtx).Cancel(w, r)
	}))

	walletGroup := r.Group()
	walletGroup.Use(midd.Auth(svcCtx.Config.JWT.AccessSecret))
	walletGroup.Post("/uc/contract-wallet", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewContractWalletHandler(svcCtx).Info(w, r)
	}))
	walletGroup.Post("/uc/contract-wallet/transfer", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewContractWalletHandler(svcCtx).Transfer(w, r)
	}))
	walletGroup.Post("/uc/asset/contract-transaction/all", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewContractWalletHandler(svcCtx).Transaction(w, r)
	}))
}
