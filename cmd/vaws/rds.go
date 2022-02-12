package vaws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

// rdsCmd represents the rds command
var rdsCmd = &cobra.Command{
	Use:   "rds",
	Short: "Show RDS instances.",
	Long:  `Show RDS instances.`,
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
		instanceFlg, err := cmd.Flags().GetBool("instance")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if instanceFlg {
			output, err := getRdsInstances(newAwsConfig())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = showRdsInstances(output, tablewriter.NewWriter(os.Stdout), sortPosition)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			output, err := getRdsClusters(newAwsConfig())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = showRdsClusters(output, tablewriter.NewWriter(os.Stdout), sortPosition)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(rdsCmd)
	rdsCmd.Flags().BoolP("instance", "i", false, "Show instances")
}

func getRdsClusters(cfg aws.Config) (*rds.DescribeDBClustersOutput, error) {
	client := rds.NewFromConfig(cfg)
	output, err := client.DescribeDBClusters(context.TODO(), &rds.DescribeDBClustersInput{
		DBClusterIdentifier: nil,
		Filters:             nil,
		IncludeShared:       false,
		Marker:              nil,
		MaxRecords:          nil,
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}

func getRdsInstances(cfg aws.Config) (*rds.DescribeDBInstancesOutput, error) {
	client := rds.NewFromConfig(cfg)
	output, err := client.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: nil,
		Filters:              nil,
		Marker:               nil,
		MaxRecords:           nil,
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}

func showRdsClusters(instances *rds.DescribeDBClustersOutput, table *tablewriter.Table, sortPosition int) error {
	header := []string{"CLUSTER", "STATUS", "INSTANCES", "WRITE-ENDPOINT", "READ-ENDPOINT"}
	if sortPosition > len(header) || 1 > sortPosition {
		return fmt.Errorf("out of sort range number when using --sort option")
	}
	recordIndex := sortPosition - 1
	table.SetHeader(header)
	var records [][]string
	for _, object := range instances.DBClusters {
		cluster := *object.DBClusterIdentifier
		status := *object.Status
		wEndpoint := *object.Endpoint
		rEndpoint := *object.ReaderEndpoint
		instanceId := ""
		for i, v := range object.DBClusterMembers {
			if i == 0 {
				instanceId += *v.DBInstanceIdentifier
			} else {
				instanceId += fmt.Sprintf(", %s", *v.DBInstanceIdentifier)
			}
		}
		records = append(records, []string{cluster, status, instanceId, wEndpoint, rEndpoint})
	}
	sort.Slice(records, func(i, j int) bool { return records[i][recordIndex] < records[j][recordIndex] })
	table.AppendBulk(records)
	table.Render()
	return nil
}

func showRdsInstances(instances *rds.DescribeDBInstancesOutput, table *tablewriter.Table, sortPosition int) error {
	header := []string{"CLUSTER", "INSTANCE", "TYPE", "ENGINE", "STATUS", "ENDPOINT(INSTANCE)"}
	if sortPosition > len(header) || 1 > sortPosition {
		return fmt.Errorf("out of sort range number when using --sort option")
	}
	recordIndex := sortPosition - 1
	table.SetHeader(header)
	var records [][]string
	for _, object := range instances.DBInstances {
		cluster := *object.DBClusterIdentifier
		instanceName := *object.DBInstanceIdentifier
		class := *object.DBInstanceClass
		engine := *object.EngineVersion
		status := *object.DBInstanceStatus
		endpoint := *object.Endpoint.Address
		records = append(records, []string{cluster, instanceName, class, engine, status, endpoint})
	}
	sort.Slice(records, func(i, j int) bool { return records[i][recordIndex] < records[j][recordIndex] })
	table.AppendBulk(records)
	table.Render()
	return nil
}
