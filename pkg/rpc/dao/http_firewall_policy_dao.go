package dao

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ipconfigs"
)

var SharedHTTPFirewallPolicyDAO = new(HTTPFirewallPolicyDAO)

// WAF策略相关
type HTTPFirewallPolicyDAO struct {
	BaseDAO
}

// 查找WAF策略基本信息
func (this *HTTPFirewallPolicyDAO) FindEnabledHTTPFirewallPolicy(ctx context.Context, policyId int64) (*pb.HTTPFirewallPolicy, error) {
	resp, err := this.RPC().HTTPFirewallPolicyRPC().FindEnabledHTTPFirewallPolicy(ctx, &pb.FindEnabledHTTPFirewallPolicyRequest{HttpFirewallPolicyId: policyId})
	if err != nil {
		return nil, err
	}
	return resp.HttpFirewallPolicy, nil
}

// 查找WAF策略配置
func (this *HTTPFirewallPolicyDAO) FindEnabledHTTPFirewallPolicyConfig(ctx context.Context, policyId int64) (*firewallconfigs.HTTPFirewallPolicy, error) {
	resp, err := this.RPC().HTTPFirewallPolicyRPC().FindEnabledHTTPFirewallPolicyConfig(ctx, &pb.FindEnabledHTTPFirewallPolicyConfigRequest{HttpFirewallPolicyId: policyId})
	if err != nil {
		return nil, err
	}
	if len(resp.HttpFirewallPolicyJSON) == 0 {
		return nil, nil
	}
	firewallPolicy := &firewallconfigs.HTTPFirewallPolicy{}
	err = json.Unmarshal(resp.HttpFirewallPolicyJSON, firewallPolicy)
	if err != nil {
		return nil, err
	}
	return firewallPolicy, nil
}

// 查找WAF的Inbound
func (this *HTTPFirewallPolicyDAO) FindEnabledHTTPFirewallPolicyInboundConfig(ctx context.Context, policyId int64) (*firewallconfigs.HTTPFirewallInboundConfig, error) {
	config, err := this.FindEnabledHTTPFirewallPolicyConfig(ctx, policyId)
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, errors.New("not found")
	}
	return config.Inbound, nil
}

// 根据类型查找WAF的IP名单
func (this *HTTPFirewallPolicyDAO) FindEnabledPolicyIPListIdWithType(ctx context.Context, policyId int64, listType ipconfigs.IPListType) (int64, error) {
	switch listType {
	case ipconfigs.IPListTypeWhite:
		return this.FindEnabledPolicyWhiteIPListId(ctx, policyId)
	case ipconfigs.IPListTypeBlack:
		return this.FindEnabledPolicyBlackIPListId(ctx, policyId)
	default:
		return 0, errors.New("invalid ip list type '" + listType + "'")
	}
}

// 查找WAF的白名单
func (this *HTTPFirewallPolicyDAO) FindEnabledPolicyWhiteIPListId(ctx context.Context, policyId int64) (int64, error) {
	config, err := this.FindEnabledHTTPFirewallPolicyConfig(ctx, policyId)
	if err != nil {
		return 0, err
	}
	if config == nil {
		return 0, errors.New("not found")
	}
	if config.Inbound == nil {
		config.Inbound = &firewallconfigs.HTTPFirewallInboundConfig{IsOn: true}
	}
	if config.Inbound.AllowListRef == nil || config.Inbound.AllowListRef.ListId == 0 {
		createResp, err := this.RPC().IPListRPC().CreateIPList(ctx, &pb.CreateIPListRequest{
			Type:        "white",
			Name:        "白名单",
			Code:        "white",
			TimeoutJSON: nil,
		})
		if err != nil {
			return 0, err
		}
		listId := createResp.IpListId
		config.Inbound.AllowListRef = &ipconfigs.IPListRef{
			IsOn:   true,
			ListId: listId,
		}
		inboundJSON, err := json.Marshal(config.Inbound)
		if err != nil {
			return 0, err
		}
		_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallInboundConfig(ctx, &pb.UpdateHTTPFirewallInboundConfigRequest{
			HttpFirewallPolicyId: policyId,
			InboundJSON:          inboundJSON,
		})
		if err != nil {
			return 0, err
		}
		return listId, nil
	}

	return config.Inbound.AllowListRef.ListId, nil
}

// 查找WAF的黑名单
func (this *HTTPFirewallPolicyDAO) FindEnabledPolicyBlackIPListId(ctx context.Context, policyId int64) (int64, error) {
	config, err := this.FindEnabledHTTPFirewallPolicyConfig(ctx, policyId)
	if err != nil {
		return 0, err
	}
	if config == nil {
		return 0, errors.New("not found")
	}
	if config.Inbound == nil {
		config.Inbound = &firewallconfigs.HTTPFirewallInboundConfig{IsOn: true}
	}
	if config.Inbound.DenyListRef == nil || config.Inbound.DenyListRef.ListId == 0 {
		createResp, err := this.RPC().IPListRPC().CreateIPList(ctx, &pb.CreateIPListRequest{
			Type:        "black",
			Name:        "黑名单",
			Code:        "black",
			TimeoutJSON: nil,
		})
		if err != nil {
			return 0, err
		}
		listId := createResp.IpListId
		config.Inbound.DenyListRef = &ipconfigs.IPListRef{
			IsOn:   true,
			ListId: listId,
		}
		inboundJSON, err := json.Marshal(config.Inbound)
		if err != nil {
			return 0, err
		}
		_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallInboundConfig(ctx, &pb.UpdateHTTPFirewallInboundConfigRequest{
			HttpFirewallPolicyId: policyId,
			InboundJSON:          inboundJSON,
		})
		if err != nil {
			return 0, err
		}
		return listId, nil
	}

	return config.Inbound.DenyListRef.ListId, nil
}

// 根据服务Id查找WAF策略
func (this *HTTPFirewallPolicyDAO) FindEnabledHTTPFirewallPolicyWithServerId(ctx context.Context, serverId int64) (*pb.HTTPFirewallPolicy, error) {
	serverResp, err := this.RPC().ServerRPC().FindEnabledServer(ctx, &pb.FindEnabledServerRequest{ServerId: serverId})
	if err != nil {
		return nil, err
	}
	server := serverResp.Server
	if server == nil {
		return nil, nil
	}
	if server.NodeCluster == nil {
		return nil, nil
	}
	clusterId := server.NodeCluster.Id
	cluster, err := SharedNodeClusterDAO.FindEnabledNodeCluster(ctx, clusterId)
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, nil
	}
	if cluster.HttpFirewallPolicyId == 0 {
		return nil, nil
	}
	return SharedHTTPFirewallPolicyDAO.FindEnabledHTTPFirewallPolicy(ctx, cluster.HttpFirewallPolicyId)
}
