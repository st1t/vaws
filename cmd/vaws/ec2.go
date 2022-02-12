package vaws

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const ec2MaxResults = 1000

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Show EC2 instances.",
	Long:  `Show EC2 instances.`,
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
		outputs, err := getEc2Instances(newAwsConfig())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sortPosition, err := cmd.Flags().GetInt("sort-position")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = showEc2Instances(outputs, tablewriter.NewWriter(os.Stdout), sortPosition)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(ec2Cmd)
}

func getEc2Instances(cfg aws.Config) ([]*ec2.DescribeInstancesOutput, error) {
	var outputs []*ec2.DescribeInstancesOutput
	var err error
	client := ec2.NewFromConfig(cfg)
	output := &ec2.DescribeInstancesOutput{
		NextToken: aws.String(""),
	}
	for output.NextToken != nil {
		output, err = client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
			MaxResults: aws.Int32(ec2MaxResults),
			NextToken:  output.NextToken,
		})
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func showEc2Instances(outputs []*ec2.DescribeInstancesOutput, table *tablewriter.Table, sortPosition int) error {
	header := []string{"NAME", "ID", "TYPE", "PRIVATE_IP", "PUBLIC_IP", "STATE", "SECURITY_GROUP"}
	if sortPosition > len(header) || 1 > sortPosition {
		return fmt.Errorf("out of sort range number when using --sort option")
	}
	recordIndex := sortPosition - 1
	table.SetHeader(header)
	var records [][]string
	for _, o := range outputs {
		for _, r := range o.Reservations {
			var privateIp, publicIp, name string
			for _, instance := range r.Instances {
				securityGroups := ""
				for _, sg := range instance.SecurityGroups {
					securityGroups = securityGroups + fmt.Sprintf("%s(%s)", *sg.GroupName, *sg.GroupId)
				}
				if instance.PrivateIpAddress != nil {
					privateIp = *instance.PrivateIpAddress
				}
				if instance.PublicIpAddress != nil {
					publicIp = *instance.PublicIpAddress
				}
				for _, t := range instance.Tags {
					if *t.Key == "Name" {
						name = *t.Value
					}
				}
				records = append(records, []string{
					name,
					*instance.InstanceId,
					string(instance.InstanceType),
					privateIp,
					publicIp,
					string(instance.State.Name),
					securityGroups,
				})
			}
		}
	}
	sort.Slice(records, func(i, j int) bool { return records[i][recordIndex] < records[j][recordIndex] })
	table.AppendBulk(records)
	table.Render()
	return nil
}
