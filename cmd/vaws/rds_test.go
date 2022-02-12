package vaws

import (
	"bytes"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/olekukonko/tablewriter"
	"testing"
)

func Test_showRdsClusters(t *testing.T) {
	type args struct {
		instances    *rds.DescribeDBClustersOutput
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
				instances: &rds.DescribeDBClustersOutput{
					DBClusters: []types.DBCluster{
						{
							DBClusterIdentifier: aws.String("test-cluster-01"),
							DBClusterMembers: []types.DBClusterMember{
								{
									DBInstanceIdentifier: aws.String("test-cluster-instance-01"),
								},
							},
							Endpoint:       aws.String("test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com"),
							ReaderEndpoint: aws.String("test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com"),
							Status:         aws.String("available"),
						},
						{
							DBClusterIdentifier: aws.String("test-cluster-01"),
							DBClusterMembers: []types.DBClusterMember{
								{
									DBInstanceIdentifier: aws.String("test-cluster-instance-02"),
								},
							},
							Endpoint:       aws.String("test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com"),
							ReaderEndpoint: aws.String("test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com"),
							Status:         aws.String("available"),
						},
					},
				},
				sortPosition: 1,
			},
			want: `+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
|     CLUSTER     |  STATUS   |        INSTANCES         |                            WRITE-ENDPOINT                             |                              READ-ENDPOINT                               |
+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
| test-cluster-01 | available | test-cluster-instance-01 | test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com | test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com |
| test-cluster-01 | available | test-cluster-instance-02 | test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com | test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com |
+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
`,
		},
		{
			name: "sort request",
			args: args{
				instances: &rds.DescribeDBClustersOutput{
					DBClusters: []types.DBCluster{
						{
							DBClusterIdentifier: aws.String("test-cluster-01"),
							DBClusterMembers: []types.DBClusterMember{
								{
									DBInstanceIdentifier: aws.String("test-cluster-instance-02"),
								},
							},
							Endpoint:       aws.String("test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com"),
							ReaderEndpoint: aws.String("test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com"),
							Status:         aws.String("available"),
						},
						{
							DBClusterIdentifier: aws.String("test-cluster-01"),
							DBClusterMembers: []types.DBClusterMember{
								{
									DBInstanceIdentifier: aws.String("test-cluster-instance-01"),
								},
							},
							Endpoint:       aws.String("test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com"),
							ReaderEndpoint: aws.String("test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com"),
							Status:         aws.String("available"),
						},
					},
				},
				sortPosition: 3,
			},
			want: `+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
|     CLUSTER     |  STATUS   |        INSTANCES         |                            WRITE-ENDPOINT                             |                              READ-ENDPOINT                               |
+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
| test-cluster-01 | available | test-cluster-instance-01 | test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com | test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com |
| test-cluster-01 | available | test-cluster-instance-02 | test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com | test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com |
+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
`,
		},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		t.Run(tt.name, func(t *testing.T) {
			showRdsClusters(tt.args.instances, tablewriter.NewWriter(&buf), tt.args.sortPosition)
			if buf.String() != tt.want {
				t.Errorf("failed to test: %s\n", tt.name)
				t.Errorf("\nwant:\n%s\n", tt.want)
				t.Errorf("\ninput:\n%s\n", buf.String())
			}
		})
	}
}
