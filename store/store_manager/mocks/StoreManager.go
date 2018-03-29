// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import pb "github.com/maichain/eth-indexer/indexer/pb"

// StoreManager is an autogenerated mock type for the StoreManager type
type StoreManager struct {
	mock.Mock
}

// GetLatestHeader provides a mock function with given fields:
func (_m *StoreManager) GetLatestHeader() (*pb.BlockHeader, error) {
	ret := _m.Called()

	var r0 *pb.BlockHeader
	if rf, ok := ret.Get(0).(func() *pb.BlockHeader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.BlockHeader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: header, transaction, receipts
func (_m *StoreManager) Upsert(header *pb.BlockHeader, transaction []*pb.Transaction, receipts []*pb.TransactionReceipt) error {
	ret := _m.Called(header, transaction, receipts)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pb.BlockHeader, []*pb.Transaction, []*pb.TransactionReceipt) error); ok {
		r0 = rf(header, transaction, receipts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
