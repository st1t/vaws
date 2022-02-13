package vaws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

// subnetCmd represents the subnet command
var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Show subnet",
	Long:  `Show subnet`,
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
		outputs, err := getSubnets(newAwsConfig())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sortPosition, err := cmd.Flags().GetInt("sort-position")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		showSubnets(outputs, tablewriter.NewWriter(os.Stdout), sortPosition)
	},
}

func init() {
	rootCmd.AddCommand(subnetCmd)
}

func getSubnets(cfg aws.Config) ([]*ec2.DescribeSubnetsOutput, error) {
	var outputs []*ec2.DescribeSubnetsOutput
	var err error
	client := ec2.NewFromConfig(cfg)
	// The DescribeSubnets API executes the API once at the beginning because the NextToken "" is disallowed
	output, err := client.DescribeSubnets(context.TODO(), &ec2.DescribeSubnetsInput{})
	if err != nil {
		return nil, err
	}
	outputs = append(outputs, output)
	for output.NextToken != nil {
		output, err = client.DescribeSubnets(context.TODO(), &ec2.DescribeSubnetsInput{
			NextToken: output.NextToken,
		})
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func showSubnets(outputs []*ec2.DescribeSubnetsOutput, table *tablewriter.Table, sortPosition int) error {
	header := []string{"NAME", "SUBNET ID", "CIDR", "VPC", "AZ", "AZ ID", "MAP PUBLIC IP", "AVAILABLE IP COUNT"}
	if sortPosition > len(header) || 1 > sortPosition {
		return fmt.Errorf("out of sort range number when using --sort option")
	}
	recordIndex := sortPosition - 1
	table.SetHeader(header)
	var records [][]string
	for _, o := range outputs {
		for _, subnet := range o.Subnets {
			name := ""
			for _, tag := range subnet.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}
			id := *subnet.SubnetId
			cidr := *subnet.CidrBlock
			vpc := *subnet.VpcId
			az := *subnet.AvailabilityZone
			azId := *subnet.AvailabilityZoneId
			isPublicIp := strconv.FormatBool(*subnet.MapPublicIpOnLaunch)
			count := fmt.Sprintf("%d", *subnet.AvailableIpAddressCount)
			records = append(records, []string{name, id, cidr, vpc, az, azId, isPublicIp, count})
		}
	}
	sort.Slice(records, func(i, j int) bool { return records[i][recordIndex] < records[j][recordIndex] })
	table.AppendBulk(records)
	table.Render()
	return nil
}
