package vaws

import (
	"bytes"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/olekukonko/tablewriter"
	"testing"
)

func Test_showElb(t *testing.T) {
	type args struct {
		output       *elasticloadbalancingv2.DescribeLoadBalancersOutput
		table        *tablewriter.Table
		sortPosition int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				output: &elasticloadbalancingv2.DescribeLoadBalancersOutput{
					LoadBalancers: []types.LoadBalancer{
						{
							AvailabilityZones: []types.AvailabilityZone{
								{
									SubnetId: aws.String("subnet-1234567e3xxxxxxxx"),
								},
								{
									SubnetId: aws.String("subnet-1234567e3zzzzzzzz"),
								},
							},
							DNSName:          aws.String("test-lb01.ap-northeast-1.elb.amazonaws.com"),
							IpAddressType:    "ipv4",
							LoadBalancerName: aws.String("test-lb01"),
							Scheme:           "internet-facing",
							SecurityGroups:   []string{"sg-084d3a6xxxxxxxxx"},
							Type:             "application",
							VpcId:            aws.String("vpc-xxxxxxxx"),
						},
						{
							AvailabilityZones: []types.AvailabilityZone{
								{
									SubnetId: aws.String("subnet-1234567e3xxxxxxxx"),
								},
								{
									SubnetId: aws.String("subnet-1234567e3zzzzzzzz"),
								},
							},
							DNSName:          aws.String("test-lb02.ap-northeast-1.elb.amazonaws.com"),
							IpAddressType:    "ipv4",
							LoadBalancerName: aws.String("test-lb02"),
							Scheme:           "internet-facing",
							SecurityGroups:   []string{"sg-084d3a6xxxxxxxxx"},
							Type:             "application",
							VpcId:            aws.String("vpc-xxxxxxxx"),
						},
						{
							AvailabilityZones: []types.AvailabilityZone{
								{
									SubnetId: aws.String("subnet-1234567e3xxxxxxxx"),
								},
								{
									SubnetId: aws.String("subnet-1234567e3zzzzzzzz"),
								},
							},
							DNSName:          aws.String("test-lb03.ap-northeast-1.elb.amazonaws.com"),
							IpAddressType:    "ipv4",
							LoadBalancerName: aws.String("test-lb03"),
							Scheme:           "internet-facing",
							Type:             "application",
							VpcId:            aws.String("vpc-xxxxxxxx"),
						},
					},
				},
				sortPosition: 1,
			},
			want: `+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
|    LB     |    TYPE     |     SCHEME      |     VPC      |                      SUBNET                       |   SECURITY GROUP    | IP TYPE |                  DNS NAME                  |
+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
| test-lb01 | application | internet-facing | vpc-xxxxxxxx | subnet-1234567e3xxxxxxxx,subnet-1234567e3zzzzzzzz | sg-084d3a6xxxxxxxxx | ipv4    | test-lb01.ap-northeast-1.elb.amazonaws.com |
| test-lb02 | application | internet-facing | vpc-xxxxxxxx | subnet-1234567e3xxxxxxxx,subnet-1234567e3zzzzzzzz | sg-084d3a6xxxxxxxxx | ipv4    | test-lb02.ap-northeast-1.elb.amazonaws.com |
| test-lb03 | application | internet-facing | vpc-xxxxxxxx | subnet-1234567e3xxxxxxxx,subnet-1234567e3zzzzzzzz | none                | ipv4    | test-lb03.ap-northeast-1.elb.amazonaws.com |
+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
`,
		},
		{
			name: "sort request",
			args: args{
				output: &elasticloadbalancingv2.DescribeLoadBalancersOutput{
					LoadBalancers: []types.LoadBalancer{
						{
							AvailabilityZones: []types.AvailabilityZone{
								{
									SubnetId: aws.String("subnet-1234567e3xxxxxxxx"),
								},
								{
									SubnetId: aws.String("subnet-1234567e3zzzzzzzz"),
								},
							},
							DNSName:          aws.String("test-lb01.ap-northeast-1.elb.amazonaws.com"),
							IpAddressType:    "ipv4",
							LoadBalancerName: aws.String("test-lb01"),
							Scheme:           "internet-facing",
							SecurityGroups:   []string{"sg-084d3a6xxxxxxxxx"},
							Type:             "application",
							VpcId:            aws.String("vpc-zzzzzzzz"),
						},
						{
							AvailabilityZones: []types.AvailabilityZone{
								{
									SubnetId: aws.String("subnet-1234567e3xxxxxxxx"),
								},
								{
									SubnetId: aws.String("subnet-1234567e3zzzzzzzz"),
								},
							},
							DNSName:          aws.String("test-lb02.ap-northeast-1.elb.amazonaws.com"),
							IpAddressType:    "ipv4",
							LoadBalancerName: aws.String("test-lb02"),
							Scheme:           "internet-facing",
							SecurityGroups:   []string{"sg-084d3a6xxxxxxxxx"},
							Type:             "application",
							VpcId:            aws.String("vpc-xxxxxxxx"),
						},
					},
				},
				sortPosition: 4,
			},
			want: `+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
|    LB     |    TYPE     |     SCHEME      |     VPC      |                      SUBNET                       |   SECURITY GROUP    | IP TYPE |                  DNS NAME                  |
+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
| test-lb02 | application | internet-facing | vpc-xxxxxxxx | subnet-1234567e3xxxxxxxx,subnet-1234567e3zzzzzzzz | sg-084d3a6xxxxxxxxx | ipv4    | test-lb02.ap-northeast-1.elb.amazonaws.com |
| test-lb01 | application | internet-facing | vpc-zzzzzzzz | subnet-1234567e3xxxxxxxx,subnet-1234567e3zzzzzzzz | sg-084d3a6xxxxxxxxx | ipv4    | test-lb01.ap-northeast-1.elb.amazonaws.com |
+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
`,
		},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		t.Run(tt.name, func(t *testing.T) {
			showElb(tt.args.output, tablewriter.NewWriter(&buf), tt.args.sortPosition)
			if buf.String() != tt.want {
				t.Errorf("failed to test: %s\n", tt.name)
				t.Errorf("\nwant:\n%s\n", tt.want)
				t.Errorf("\ninput:\n%s\n", buf.String())
			}
		})
	}
}
