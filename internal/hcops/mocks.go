package hcops

import (
	"context"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/stretchr/testify/mock"
	"github.com/syself/hetzner-cloud-controller-manager/internal/mocks"
	v1 "k8s.io/api/core/v1"
)

type MockLoadBalancerOps struct {
	mock.Mock
}

func (m *MockLoadBalancerOps) GetByName(ctx context.Context, name string) (*hcloud.LoadBalancer, error) {
	args := m.Called(ctx, name)
	return mocks.GetLoadBalancerPtr(args, 0), args.Error(1)
}

func (m *MockLoadBalancerOps) GetByID(ctx context.Context, id int) (*hcloud.LoadBalancer, error) {
	args := m.Called(ctx, id)
	return mocks.GetLoadBalancerPtr(args, 0), args.Error(1)
}

func (m *MockLoadBalancerOps) Create(
	ctx context.Context, lbName string, service *v1.Service,
) (*hcloud.LoadBalancer, error) {
	args := m.Called(ctx, lbName, service)
	return mocks.GetLoadBalancerPtr(args, 0), args.Error(1)
}

func (m *MockLoadBalancerOps) Delete(ctx context.Context, lb *hcloud.LoadBalancer) error {
	args := m.Called(ctx, lb)
	return args.Error(0)
}

func (m *MockLoadBalancerOps) ReconcileHCLB(
	ctx context.Context, lb *hcloud.LoadBalancer, svc *v1.Service,
) (bool, error) {
	args := m.Called(ctx, lb, svc)
	return args.Bool(0), args.Error(1)
}

func (m *MockLoadBalancerOps) ReconcileHCLBTargets(
	ctx context.Context, lb *hcloud.LoadBalancer, svc *v1.Service, nodes []*v1.Node, disableIPV6 bool,
) (bool, error) {
	args := m.Called(ctx, lb, svc, nodes, disableIPV6)
	return args.Bool(0), args.Error(1)
}

func (m *MockLoadBalancerOps) ReconcileHCLBServices(
	ctx context.Context, lb *hcloud.LoadBalancer, svc *v1.Service,
) (bool, error) {
	args := m.Called(ctx, lb, svc)
	return args.Bool(0), args.Error(1)
}

func (m *MockLoadBalancerOps) GetByK8SServiceUID(ctx context.Context, svc *v1.Service) (*hcloud.LoadBalancer, error) {
	args := m.Called(ctx, svc)
	return mocks.GetLoadBalancerPtr(args, 0), args.Error(1)
}
