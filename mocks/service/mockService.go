// Code generated by MockGen. DO NOT EDIT.
// Source: services/service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	big "math/big"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	models "github.com/gateway-fm/perpsv3-Go/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// FormatAccount mocks base method.
func (m *MockIService) FormatAccount(id *big.Int) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatAccount", id)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatAccount indicates an expected call of FormatAccount.
func (mr *MockIServiceMockRecorder) FormatAccount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatAccount", reflect.TypeOf((*MockIService)(nil).FormatAccount), id)
}

// FormatAccounts mocks base method.
func (m *MockIService) FormatAccounts() ([]*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatAccounts")
	ret0, _ := ret[0].([]*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatAccounts indicates an expected call of FormatAccounts.
func (mr *MockIServiceMockRecorder) FormatAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatAccounts", reflect.TypeOf((*MockIService)(nil).FormatAccounts))
}

// FormatAccountsLimit mocks base method.
func (m *MockIService) FormatAccountsLimit(limit uint64) ([]*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatAccountsLimit", limit)
	ret0, _ := ret[0].([]*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatAccountsLimit indicates an expected call of FormatAccountsLimit.
func (mr *MockIServiceMockRecorder) FormatAccountsLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatAccountsLimit", reflect.TypeOf((*MockIService)(nil).FormatAccountsLimit), limit)
}

// GetAccountLastInteraction mocks base method.
func (m *MockIService) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountLastInteraction", accountId)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountLastInteraction indicates an expected call of GetAccountLastInteraction.
func (mr *MockIServiceMockRecorder) GetAccountLastInteraction(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountLastInteraction", reflect.TypeOf((*MockIService)(nil).GetAccountLastInteraction), accountId)
}

// GetAccountOwner mocks base method.
func (m *MockIService) GetAccountOwner(accountId *big.Int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountOwner", accountId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountOwner indicates an expected call of GetAccountOwner.
func (mr *MockIServiceMockRecorder) GetAccountOwner(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountOwner", reflect.TypeOf((*MockIService)(nil).GetAccountOwner), accountId)
}

// GetAvailableMargin mocks base method.
func (m *MockIService) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableMargin", accountId)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableMargin indicates an expected call of GetAvailableMargin.
func (mr *MockIServiceMockRecorder) GetAvailableMargin(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableMargin", reflect.TypeOf((*MockIService)(nil).GetAvailableMargin), accountId)
}

// GetCollateralAmount mocks base method.
func (m *MockIService) GetCollateralAmount(accountId, marketId *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollateralAmount", accountId, marketId)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollateralAmount indicates an expected call of GetCollateralAmount.
func (mr *MockIServiceMockRecorder) GetCollateralAmount(accountId, marketId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollateralAmount", reflect.TypeOf((*MockIService)(nil).GetCollateralAmount), accountId, marketId)
}

// GetCollateralPrice mocks base method.
func (m *MockIService) GetCollateralPrice(blockNumber *big.Int, collateralType common.Address) (*models.CollateralPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollateralPrice", blockNumber, collateralType)
	ret0, _ := ret[0].(*models.CollateralPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollateralPrice indicates an expected call of GetCollateralPrice.
func (mr *MockIServiceMockRecorder) GetCollateralPrice(blockNumber, collateralType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollateralPrice", reflect.TypeOf((*MockIService)(nil).GetCollateralPrice), blockNumber, collateralType)
}

// GetFoundingRate mocks base method.
func (m *MockIService) GetFoundingRate(marketId *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoundingRate", marketId)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoundingRate indicates an expected call of GetFoundingRate.
func (mr *MockIServiceMockRecorder) GetFoundingRate(marketId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoundingRate", reflect.TypeOf((*MockIService)(nil).GetFoundingRate), marketId)
}

// GetFundingParameters mocks base method.
func (m *MockIService) GetFundingParameters(marketId *big.Int) (*models.FundingParameters, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFundingParameters", marketId)
	ret0, _ := ret[0].(*models.FundingParameters)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFundingParameters indicates an expected call of GetFundingParameters.
func (mr *MockIServiceMockRecorder) GetFundingParameters(marketId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFundingParameters", reflect.TypeOf((*MockIService)(nil).GetFundingParameters), marketId)
}

// GetLiquidationParameters mocks base method.
func (m *MockIService) GetLiquidationParameters(marketId *big.Int) (*models.LiquidationParameters, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiquidationParameters", marketId)
	ret0, _ := ret[0].(*models.LiquidationParameters)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLiquidationParameters indicates an expected call of GetLiquidationParameters.
func (mr *MockIServiceMockRecorder) GetLiquidationParameters(marketId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiquidationParameters", reflect.TypeOf((*MockIService)(nil).GetLiquidationParameters), marketId)
}

// GetMarketIDs mocks base method.
func (m *MockIService) GetMarketIDs() ([]*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketIDs")
	ret0, _ := ret[0].([]*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketIDs indicates an expected call of GetMarketIDs.
func (mr *MockIServiceMockRecorder) GetMarketIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketIDs", reflect.TypeOf((*MockIService)(nil).GetMarketIDs))
}

// GetMarketMetadata mocks base method.
func (m *MockIService) GetMarketMetadata(marketID *big.Int) (*models.MarketMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketMetadata", marketID)
	ret0, _ := ret[0].(*models.MarketMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketMetadata indicates an expected call of GetMarketMetadata.
func (mr *MockIServiceMockRecorder) GetMarketMetadata(marketID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketMetadata", reflect.TypeOf((*MockIService)(nil).GetMarketMetadata), marketID)
}

// GetMarketSummary mocks base method.
func (m *MockIService) GetMarketSummary(marketID *big.Int) (*models.MarketSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketSummary", marketID)
	ret0, _ := ret[0].(*models.MarketSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketSummary indicates an expected call of GetMarketSummary.
func (mr *MockIServiceMockRecorder) GetMarketSummary(marketID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketSummary", reflect.TypeOf((*MockIService)(nil).GetMarketSummary), marketID)
}

// GetPosition mocks base method.
func (m *MockIService) GetPosition(accountID, marketID *big.Int) (*models.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosition", accountID, marketID)
	ret0, _ := ret[0].(*models.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosition indicates an expected call of GetPosition.
func (mr *MockIServiceMockRecorder) GetPosition(accountID, marketID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosition", reflect.TypeOf((*MockIService)(nil).GetPosition), accountID, marketID)
}

// GetRequiredMaintenanceMargin mocks base method.
func (m *MockIService) GetRequiredMaintenanceMargin(accountId *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequiredMaintenanceMargin", accountId)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequiredMaintenanceMargin indicates an expected call of GetRequiredMaintenanceMargin.
func (mr *MockIServiceMockRecorder) GetRequiredMaintenanceMargin(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequiredMaintenanceMargin", reflect.TypeOf((*MockIService)(nil).GetRequiredMaintenanceMargin), accountId)
}

// RetrieveAccountLiquidationsLimit mocks base method.
func (m *MockIService) RetrieveAccountLiquidationsLimit(limit uint64) ([]*models.AccountLiquidated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveAccountLiquidationsLimit", limit)
	ret0, _ := ret[0].([]*models.AccountLiquidated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveAccountLiquidationsLimit indicates an expected call of RetrieveAccountLiquidationsLimit.
func (mr *MockIServiceMockRecorder) RetrieveAccountLiquidationsLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAccountLiquidationsLimit", reflect.TypeOf((*MockIService)(nil).RetrieveAccountLiquidationsLimit), limit)
}

// RetrieveCollateralDepositedLimit mocks base method.
func (m *MockIService) RetrieveCollateralDepositedLimit(limit uint64) ([]*models.CollateralDeposited, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveCollateralDepositedLimit", limit)
	ret0, _ := ret[0].([]*models.CollateralDeposited)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCollateralDepositedLimit indicates an expected call of RetrieveCollateralDepositedLimit.
func (mr *MockIServiceMockRecorder) RetrieveCollateralDepositedLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCollateralDepositedLimit", reflect.TypeOf((*MockIService)(nil).RetrieveCollateralDepositedLimit), limit)
}

// RetrieveCollateralWithdrawnLimit mocks base method.
func (m *MockIService) RetrieveCollateralWithdrawnLimit(limit uint64) ([]*models.CollateralWithdrawn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveCollateralWithdrawnLimit", limit)
	ret0, _ := ret[0].([]*models.CollateralWithdrawn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCollateralWithdrawnLimit indicates an expected call of RetrieveCollateralWithdrawnLimit.
func (mr *MockIServiceMockRecorder) RetrieveCollateralWithdrawnLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCollateralWithdrawnLimit", reflect.TypeOf((*MockIService)(nil).RetrieveCollateralWithdrawnLimit), limit)
}

// RetrieveDelegationUpdatedLimit mocks base method.
func (m *MockIService) RetrieveDelegationUpdatedLimit(limit uint64) ([]*models.DelegationUpdated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveDelegationUpdatedLimit", limit)
	ret0, _ := ret[0].([]*models.DelegationUpdated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveDelegationUpdatedLimit indicates an expected call of RetrieveDelegationUpdatedLimit.
func (mr *MockIServiceMockRecorder) RetrieveDelegationUpdatedLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveDelegationUpdatedLimit", reflect.TypeOf((*MockIService)(nil).RetrieveDelegationUpdatedLimit), limit)
}

// RetrieveLiquidations mocks base method.
func (m *MockIService) RetrieveLiquidations(fromBlock uint64, toBLock *uint64) ([]*models.Liquidation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveLiquidations", fromBlock, toBLock)
	ret0, _ := ret[0].([]*models.Liquidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveLiquidations indicates an expected call of RetrieveLiquidations.
func (mr *MockIServiceMockRecorder) RetrieveLiquidations(fromBlock, toBLock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveLiquidations", reflect.TypeOf((*MockIService)(nil).RetrieveLiquidations), fromBlock, toBLock)
}

// RetrieveLiquidationsLimit mocks base method.
func (m *MockIService) RetrieveLiquidationsLimit(limit uint64) ([]*models.Liquidation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveLiquidationsLimit", limit)
	ret0, _ := ret[0].([]*models.Liquidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveLiquidationsLimit indicates an expected call of RetrieveLiquidationsLimit.
func (mr *MockIServiceMockRecorder) RetrieveLiquidationsLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveLiquidationsLimit", reflect.TypeOf((*MockIService)(nil).RetrieveLiquidationsLimit), limit)
}

// RetrieveMarketUpdates mocks base method.
func (m *MockIService) RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveMarketUpdates", fromBlock, toBLock)
	ret0, _ := ret[0].([]*models.MarketUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMarketUpdates indicates an expected call of RetrieveMarketUpdates.
func (mr *MockIServiceMockRecorder) RetrieveMarketUpdates(fromBlock, toBLock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMarketUpdates", reflect.TypeOf((*MockIService)(nil).RetrieveMarketUpdates), fromBlock, toBLock)
}

// RetrieveMarketUpdatesBig mocks base method.
func (m *MockIService) RetrieveMarketUpdatesBig(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdateBig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveMarketUpdatesBig", fromBlock, toBLock)
	ret0, _ := ret[0].([]*models.MarketUpdateBig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMarketUpdatesBig indicates an expected call of RetrieveMarketUpdatesBig.
func (mr *MockIServiceMockRecorder) RetrieveMarketUpdatesBig(fromBlock, toBLock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMarketUpdatesBig", reflect.TypeOf((*MockIService)(nil).RetrieveMarketUpdatesBig), fromBlock, toBLock)
}

// RetrieveMarketUpdatesBigLimit mocks base method.
func (m *MockIService) RetrieveMarketUpdatesBigLimit(limit uint64) ([]*models.MarketUpdateBig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveMarketUpdatesBigLimit", limit)
	ret0, _ := ret[0].([]*models.MarketUpdateBig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMarketUpdatesBigLimit indicates an expected call of RetrieveMarketUpdatesBigLimit.
func (mr *MockIServiceMockRecorder) RetrieveMarketUpdatesBigLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMarketUpdatesBigLimit", reflect.TypeOf((*MockIService)(nil).RetrieveMarketUpdatesBigLimit), limit)
}

// RetrieveMarketUpdatesLimit mocks base method.
func (m *MockIService) RetrieveMarketUpdatesLimit(limit uint64) ([]*models.MarketUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveMarketUpdatesLimit", limit)
	ret0, _ := ret[0].([]*models.MarketUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMarketUpdatesLimit indicates an expected call of RetrieveMarketUpdatesLimit.
func (mr *MockIServiceMockRecorder) RetrieveMarketUpdatesLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMarketUpdatesLimit", reflect.TypeOf((*MockIService)(nil).RetrieveMarketUpdatesLimit), limit)
}

// RetrieveOrders mocks base method.
func (m *MockIService) RetrieveOrders(fromBlock uint64, toBLock *uint64) ([]*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveOrders", fromBlock, toBLock)
	ret0, _ := ret[0].([]*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveOrders indicates an expected call of RetrieveOrders.
func (mr *MockIServiceMockRecorder) RetrieveOrders(fromBlock, toBLock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveOrders", reflect.TypeOf((*MockIService)(nil).RetrieveOrders), fromBlock, toBLock)
}

// RetrieveOrdersLimit mocks base method.
func (m *MockIService) RetrieveOrdersLimit(limit uint64) ([]*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveOrdersLimit", limit)
	ret0, _ := ret[0].([]*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveOrdersLimit indicates an expected call of RetrieveOrdersLimit.
func (mr *MockIServiceMockRecorder) RetrieveOrdersLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveOrdersLimit", reflect.TypeOf((*MockIService)(nil).RetrieveOrdersLimit), limit)
}

// RetrieveRewardClaimedLimit mocks base method.
func (m *MockIService) RetrieveRewardClaimedLimit(limit uint64) ([]*models.RewardClaimed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveRewardClaimedLimit", limit)
	ret0, _ := ret[0].([]*models.RewardClaimed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveRewardClaimedLimit indicates an expected call of RetrieveRewardClaimedLimit.
func (mr *MockIServiceMockRecorder) RetrieveRewardClaimedLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveRewardClaimedLimit", reflect.TypeOf((*MockIService)(nil).RetrieveRewardClaimedLimit), limit)
}

// RetrieveRewardDistributedLimit mocks base method.
func (m *MockIService) RetrieveRewardDistributedLimit(limit uint64) ([]*models.RewardDistributed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveRewardDistributedLimit", limit)
	ret0, _ := ret[0].([]*models.RewardDistributed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveRewardDistributedLimit indicates an expected call of RetrieveRewardDistributedLimit.
func (mr *MockIServiceMockRecorder) RetrieveRewardDistributedLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveRewardDistributedLimit", reflect.TypeOf((*MockIService)(nil).RetrieveRewardDistributedLimit), limit)
}

// RetrieveTrades mocks base method.
func (m *MockIService) RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveTrades", fromBlock, toBLock)
	ret0, _ := ret[0].([]*models.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveTrades indicates an expected call of RetrieveTrades.
func (mr *MockIServiceMockRecorder) RetrieveTrades(fromBlock, toBLock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveTrades", reflect.TypeOf((*MockIService)(nil).RetrieveTrades), fromBlock, toBLock)
}

// RetrieveTradesLimit mocks base method.
func (m *MockIService) RetrieveTradesLimit(limit uint64) ([]*models.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveTradesLimit", limit)
	ret0, _ := ret[0].([]*models.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveTradesLimit indicates an expected call of RetrieveTradesLimit.
func (mr *MockIServiceMockRecorder) RetrieveTradesLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveTradesLimit", reflect.TypeOf((*MockIService)(nil).RetrieveTradesLimit), limit)
}

// RetrieveUSDBurnedLimit mocks base method.
func (m *MockIService) RetrieveUSDBurnedLimit(limit uint64) ([]*models.USDBurned, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveUSDBurnedLimit", limit)
	ret0, _ := ret[0].([]*models.USDBurned)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveUSDBurnedLimit indicates an expected call of RetrieveUSDBurnedLimit.
func (mr *MockIServiceMockRecorder) RetrieveUSDBurnedLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveUSDBurnedLimit", reflect.TypeOf((*MockIService)(nil).RetrieveUSDBurnedLimit), limit)
}

// RetrieveUSDMintedLimit mocks base method.
func (m *MockIService) RetrieveUSDMintedLimit(limit uint64) ([]*models.USDMinted, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveUSDMintedLimit", limit)
	ret0, _ := ret[0].([]*models.USDMinted)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveUSDMintedLimit indicates an expected call of RetrieveUSDMintedLimit.
func (mr *MockIServiceMockRecorder) RetrieveUSDMintedLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveUSDMintedLimit", reflect.TypeOf((*MockIService)(nil).RetrieveUSDMintedLimit), limit)
}
