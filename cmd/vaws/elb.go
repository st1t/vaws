package vaws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"sort"
)

// elbCmd represents the elb command
var elbCmd = &cobra.Command{
	Use:   "elb",
	Short: "Show ELB.",
	Long:  `Show ELB.`,
	Run: func(cmd *cobra.Command, args []string) {
		profile, err := cmd.Flags().GetString("aws-profile")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if profile != "" {
			err := os.Setenv("AWS_PROFILE", profile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		sortPosition, err := cmd.Flags().GetInt("sort-position")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output, err := getElb(newAwsConfig())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		showElb(output, tablewriter.NewWriter(os.Stdout), sortPosition)
	},
}

func init() {
	rootCmd.AddCommand(elbCmd)
}

func getElb(cfg aws.Config) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error) {
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	output, err := client.DescribeLoadBalancers(context.TODO(), &elasticloadbalancingv2.DescribeLoadBalancersInput{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return output, nil
}

func showElb(output *elasticloadbalancingv2.DescribeLoadBalancersOutput, table *tablewriter.Table, sortPosition int) error {
	header := []string{"LB", "TYPE", "SCHEME", "VPC", "SUBNET", "SECURITY GROUP", "IP TYPE", "DNS NAME"}
	if sortPosition > len(header) || 1 > sortPosition {
		return fmt.Errorf("out of sort range number when using --sort option")
	}
	recordIndex := sortPosition - 1
	table.SetHeader(header)
	var records [][]string
	for _, lb := range output.LoadBalancers {
		name := *lb.LoadBalancerName
		lbType := lb.Type
		scheme := lb.Scheme
		vpc := lb.VpcId
		subnet := ""
		for i, zone := range lb.AvailabilityZones {
			if i == 0 {
				subnet = *zone.SubnetId
			} else {
				subnet = fmt.Sprintf("%s,%s", subnet, *zone.SubnetId)
			}
		}
		securityGroup := ""
		if lb.SecurityGroups != nil {
			for i, sg := range lb.SecurityGroups {
				if i == 0 {
					securityGroup = sg
				} else {
					securityGroup = fmt.Sprintf("%s,%s", securityGroup, sg)
				}
			}
		} else { // for classic EC2
			securityGroup = "none"
		}
		ipType := lb.IpAddressType
		dnsName := *lb.DNSName
		records = append(records, []string{name, string(lbType), string(scheme), *vpc, subnet, securityGroup, string(ipType), dnsName})
	}
	sort.Slice(records, func(i, j int) bool { return records[i][recordIndex] < records[j][recordIndex] })
	table.AppendBulk(records)
	table.Render()
	return nil
}
