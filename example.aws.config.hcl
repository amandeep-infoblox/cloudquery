// Configuration AutoGenerated by CloudQuery CLI
cloudquery {
  provider "aws" {
    source = "cloudquery/cq-provider-aws"
    version = "latest"
  }
  // Can be configured via CLI variables
  connection {
    dsn = "host=localhost user=postgres password=pass database=postgres port=5432"
  }
}
// All Provider Configurations
provider "aws" {
  configuration {
    // Optional. if you want to assume role to multiple account and fetch data from them
    //accounts "<YOUR ID>"{
    // Optional. Role ARN we want to assume when accessing this account
    // role_arn = <YOUR_ROLE_ARN>
    // }
    // Optional. by default assumes all regions
    //    regions = [
    //      "us-east-1",
    //      "us-west-2",
    //    ]
    // Optional. Enable AWS SDK debug logging.
    aws_debug = false
    // The maximum number of times that a request will be retried for failures. Defaults to 5 retry attempts.
    max_retries = 7
    // The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 60 seconds.
    max_backoff = 60
  }
  resources = [
    "autoscaling.launch_configurations",
    "cloudfront.distributions",
    "cloudfront.cache_policies",
    "cloudtrail.trails",
    "cloudwatch.alarms",
    "cloudwatchlogs.filters",
    "s3.buckets",
    "directconnect.gateways",
    "directconnect.virtual_gateways",
    "directconnect.virtual_interfaces",
    "ec2.byoip_cidrs",
    "ec2.customer_gateways",
    "ec2.flow_logs",
    "ec2.images",
    "ec2.internet_gateways",
    "ec2.nat_gateways",
    "ec2.network_acls",
    "ec2.route_tables",
    "ec2.subnets",
    "ec2.transit_gateways",
    "ec2.vpc_peering_connections",
    "ec2.vpc_endpoints",
    "ec2.vpcs",
    "ec2.instances",
    "ec2.security_groups",
    "ec2.ebs_volumes",
    "ecr.repositories",
    "efs.filesystems",
    "eks.clusters",
    "ecs.clusters",
    "elasticbeanstalk.environments",
    "elbv2.target_groups",
    "elbv2.load_balancers",
    "emr.clusters",
    "fsx.backups",
    "iam.groups",
    "iam.policies",
    "iam.password_policies",
    "iam.roles",
    "iam.users",
    "iam.virtual_mfa_devices",
    "iam.openid_connect_identity_providers",
    "iam.saml_identity_providers",
    "kms.keys",
    "organizations.accounts",
    "sns.topics",
    "sns.subscriptions",
    "rds.certificates",
    "rds.clusters",
    "rds.db_subnet_groups",
    "rds.instances",
    "redshift.clusters",
    "redshift.subnet_groups",
    "route53.reusable_delegation_sets",
    "route53.health_checks",
    "route53.hosted_zones",
    "route53.traffic_policies",
  ]
}