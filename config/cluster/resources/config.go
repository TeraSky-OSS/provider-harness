package resources

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// identifierExtractor extracts the "identifier" param from a referenced
// resource's spec.forProvider. Used when the Terraform id is a composite path
// (e.g. org_id/project_id) but the referencing field only needs the leaf
// identifier.
const identifierExtractor = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("identifier",false)`

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// ── Common reference definitions ──────────────────────────────────────────

	// orgRef: Organization's TF id == its identifier, so no extractor needed.
	orgRef := config.Reference{
		TerraformName: "harness_platform_organization",
	}
	// projectRef: Project's TF id is "org_id/project_id", so we extract only
	// the "identifier" attribute.
	projectRef := config.Reference{
		TerraformName: "harness_platform_project",
		Extractor:     identifierExtractor,
	}
	envRef := config.Reference{
		TerraformName: "harness_platform_environment",
		Extractor:     identifierExtractor,
	}
	pipelineRef := config.Reference{
		TerraformName: "harness_platform_pipeline",
		Extractor:     identifierExtractor,
	}
	serviceRef := config.Reference{
		TerraformName: "harness_platform_service",
		Extractor:     identifierExtractor,
	}
	infraRef := config.Reference{
		TerraformName: "harness_platform_infrastructure",
		Extractor:     identifierExtractor,
	}
	agentRef := config.Reference{
		TerraformName: "harness_platform_gitops_agent",
		Extractor:     identifierExtractor,
	}
	apikeyRef := config.Reference{
		TerraformName: "harness_platform_apikey",
		Extractor:     identifierExtractor,
	}
	dbSchemaRef := config.Reference{
		TerraformName: "harness_platform_db_schema",
		Extractor:     identifierExtractor,
	}

	// ── Legacy / CG resources ─────────────────────────────────────────────────

	p.AddResourceConfigurator("harness_add_user_to_group", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "AddUserToGroup"
	})
	p.AddResourceConfigurator("harness_application", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "Application"
	})
	p.AddResourceConfigurator("harness_application_gitsync", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "ApplicationGitsync"
	})
	p.AddResourceConfigurator("harness_delegate_approval", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "DelegateApproval"
	})
	p.AddResourceConfigurator("harness_encrypted_text", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "EncryptedText"
	})
	p.AddResourceConfigurator("harness_environment", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "Environment"
	})
	p.AddResourceConfigurator("harness_git_connector", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "GitConnector"
	})
	p.AddResourceConfigurator("harness_infrastructure_definition", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "InfrastructureDefinition"
	})
	p.AddResourceConfigurator("harness_ssh_credential", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "SshCredential"
	})
	p.AddResourceConfigurator("harness_user", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "User"
	})
	p.AddResourceConfigurator("harness_user_group", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "UserGroup"
	})
	p.AddResourceConfigurator("harness_user_group_permissions", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "UserGroupPermissions"
	})
	p.AddResourceConfigurator("harness_yaml_config", func(r *config.Resource) {
		r.ShortGroup = "harness"
		r.Kind = "YamlConfig"
	})

	// ── AutoStopping resources ─────────────────────────────────────────────────

	p.AddResourceConfigurator("harness_autostopping_alert", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "Alert"
	})
	p.AddResourceConfigurator("harness_autostopping_aws_alb", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "AwsAlb"
	})
	p.AddResourceConfigurator("harness_autostopping_aws_proxy", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "AwsProxy"
	})
	p.AddResourceConfigurator("harness_autostopping_azure_gateway", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "AzureGateway"
	})
	p.AddResourceConfigurator("harness_autostopping_azure_proxy", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "AzureProxy"
	})
	p.AddResourceConfigurator("harness_autostopping_gcp_proxy", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "GcpProxy"
	})
	p.AddResourceConfigurator("harness_autostopping_rule_ecs", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "RuleEcs"
	})
	p.AddResourceConfigurator("harness_autostopping_rule_k8s", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "RuleK8s"
	})
	p.AddResourceConfigurator("harness_autostopping_rule_rds", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "RuleRds"
	})
	p.AddResourceConfigurator("harness_autostopping_rule_scale_group", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "RuleScaleGroup"
	})
	p.AddResourceConfigurator("harness_autostopping_rule_vm", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "RuleVm"
	})
	p.AddResourceConfigurator("harness_autostopping_schedule", func(r *config.Resource) {
		r.ShortGroup = "autostopping"
		r.Kind = "Schedule"
	})

	// ── Chaos resources ───────────────────────────────────────────────────────

	p.AddResourceConfigurator("harness_chaos_action_template", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "ActionTemplate"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_experiment", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "Experiment"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_experiment_template", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "ExperimentTemplate"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_fault_template", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "FaultTemplate"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_hub", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "Hub"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_hub_sync", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "HubSync"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_hub_v2", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "HubV2"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_image_registry", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "ImageRegistry"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_infrastructure", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "Infrastructure"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["environment_id"] = envRef
	})
	p.AddResourceConfigurator("harness_chaos_infrastructure_v2", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "InfrastructureV2"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["environment_id"] = envRef
	})
	p.AddResourceConfigurator("harness_chaos_probe_template", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "ProbeTemplate"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_security_governance_condition", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "SecurityGovernanceCondition"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_chaos_security_governance_rule", func(r *config.Resource) {
		r.ShortGroup = "chaos"
		r.Kind = "SecurityGovernanceRule"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})

	// ── Cloud Provider (legacy CG) resources ──────────────────────────────────

	p.AddResourceConfigurator("harness_cloudprovider_aws", func(r *config.Resource) {
		r.ShortGroup = "cloudprovider"
		r.Kind = "Aws"
	})
	p.AddResourceConfigurator("harness_cloudprovider_azure", func(r *config.Resource) {
		r.ShortGroup = "cloudprovider"
		r.Kind = "Azure"
	})
	p.AddResourceConfigurator("harness_cloudprovider_datacenter", func(r *config.Resource) {
		r.ShortGroup = "cloudprovider"
		r.Kind = "Datacenter"
	})
	p.AddResourceConfigurator("harness_cloudprovider_gcp", func(r *config.Resource) {
		r.ShortGroup = "cloudprovider"
		r.Kind = "Gcp"
	})
	p.AddResourceConfigurator("harness_cloudprovider_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "cloudprovider"
		r.Kind = "Kubernetes"
	})
	p.AddResourceConfigurator("harness_cloudprovider_spot", func(r *config.Resource) {
		r.ShortGroup = "cloudprovider"
		r.Kind = "Spot"
	})
	p.AddResourceConfigurator("harness_cloudprovider_tanzu", func(r *config.Resource) {
		r.ShortGroup = "cloudprovider"
		r.Kind = "Tanzu"
	})

	// ── Cluster Orchestrator resources ────────────────────────────────────────

	p.AddResourceConfigurator("harness_cluster_orchestrator", func(r *config.Resource) {
		r.ShortGroup = "cluster"
		r.Kind = "Orchestrator"
	})
	p.AddResourceConfigurator("harness_cluster_orchestrator_config", func(r *config.Resource) {
		r.ShortGroup = "cluster"
		r.Kind = "OrchestratorConfig"
	})

	// ── Governance resources ──────────────────────────────────────────────────

	p.AddResourceConfigurator("harness_governance_rule", func(r *config.Resource) {
		r.ShortGroup = "governance"
		r.Kind = "Rule"
	})
	p.AddResourceConfigurator("harness_governance_rule_enforcement", func(r *config.Resource) {
		r.ShortGroup = "governance"
		r.Kind = "RuleEnforcement"
	})
	p.AddResourceConfigurator("harness_governance_rule_set", func(r *config.Resource) {
		r.ShortGroup = "governance"
		r.Kind = "RuleSet"
	})

	// ── Service (legacy NG) resources ─────────────────────────────────────────

	p.AddResourceConfigurator("harness_service_ami", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "Ami"
	})
	p.AddResourceConfigurator("harness_service_aws_codedeploy", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "AwsCodeDeploy"
	})
	p.AddResourceConfigurator("harness_service_aws_lambda", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "AwsLambda"
	})
	p.AddResourceConfigurator("harness_service_discovery_agent", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "DiscoveryAgent"
	})
	p.AddResourceConfigurator("harness_service_discovery_setting", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "DiscoverySetting"
	})
	p.AddResourceConfigurator("harness_service_ecs", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "Ecs"
	})
	p.AddResourceConfigurator("harness_service_helm", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "Helm"
	})
	p.AddResourceConfigurator("harness_service_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "Kubernetes"
	})
	p.AddResourceConfigurator("harness_service_ssh", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "Ssh"
	})
	p.AddResourceConfigurator("harness_service_tanzu", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "Tanzu"
	})
	p.AddResourceConfigurator("harness_service_winrm", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "Winrm"
	})

	// ── Platform: Organization / Project ──────────────────────────────────────

	p.AddResourceConfigurator("harness_platform_organization", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Organization"
	})
	p.AddResourceConfigurator("harness_platform_project", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Project"
		r.References["org_id"] = orgRef
	})

	// ── Platform: Connectors ──────────────────────────────────────────────────

	p.AddResourceConfigurator("harness_platform_connector_appdynamics", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "AppDynamics"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_artifactory", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Artifactory"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_aws", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Aws"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_aws_secret_manager", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "AwsSecretManager"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_awscc", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Awscc"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_awskms", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "AwsKms"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_azure_artifacts", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "AzureArtifacts"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_azure_cloud_cost", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "AzureCloudCost"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_azure_cloud_provider", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "AzureCloudProvider"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_azure_key_vault", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "AzureKeyVault"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_azure_repo", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "AzureRepo"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_bitbucket", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Bitbucket"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_custom_secret_manager", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "CustomSecretManager"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_customhealthsource", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "CustomHealthSource"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_datadog", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Datadog"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_docker", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Docker"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_dynatrace", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Dynatrace"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_elasticsearch", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Elasticsearch"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_gcp", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Gcp"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_gcp_cloud_cost", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "GcpCloudCost"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_gcp_kms", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "GcpKms"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_gcp_secret_manager", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "GcpSecretManager"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_git", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Git"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_github", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Github"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_gitlab", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Gitlab"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_helm", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Helm"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_jdbc", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Jdbc"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_jenkins", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Jenkins"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_jira", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Jira"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Kubernetes"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_kubernetes_cloud_cost", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "KubernetesCloudCost"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_newrelic", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "NewRelic"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_nexus", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Nexus"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_oci_helm", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "OciHelm"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_pagerduty", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Pagerduty"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_pdc", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Pdc"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_prometheus", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Prometheus"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_rancher", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Rancher"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_service_now", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "ServiceNow"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_splunk", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Splunk"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_spot", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Spot"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_sumologic", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Sumologic"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_tas", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Tas"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_terraform_cloud", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "TerraformCloud"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_connector_vault", func(r *config.Resource) {
		r.ShortGroup = "connector"
		r.Kind = "Vault"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})

	// ── Platform: Secrets ─────────────────────────────────────────────────────

	p.AddResourceConfigurator("harness_platform_secret_file", func(r *config.Resource) {
		r.ShortGroup = "secret"
		r.Kind = "File"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_secret_sshkey", func(r *config.Resource) {
		r.ShortGroup = "secret"
		r.Kind = "SshKey"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_secret_text", func(r *config.Resource) {
		r.ShortGroup = "secret"
		r.Kind = "Text"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_secret_winrm", func(r *config.Resource) {
		r.ShortGroup = "secret"
		r.Kind = "Winrm"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})

	// ── Platform: GitOps ──────────────────────────────────────────────────────

	p.AddResourceConfigurator("harness_platform_gitops_agent", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "Agent"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_app_project", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "AppProject"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["agent_id"] = agentRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_app_project_mapping", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "AppProjectMapping"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["agent_id"] = agentRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_applications", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "Applications"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["agent_id"] = agentRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_cluster", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "Cluster"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["agent_id"] = agentRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_filters", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "Filters"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_gnupg", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "GnuPG"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["agent_id"] = agentRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_repo_cert", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "RepoCert"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["agent_id"] = agentRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_repo_cred", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "RepoCred"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["agent_id"] = agentRef
	})
	p.AddResourceConfigurator("harness_platform_gitops_repository", func(r *config.Resource) {
		r.ShortGroup = "gitops"
		r.Kind = "Repository"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["agent_id"] = agentRef
	})

	// ── Platform: IDP ─────────────────────────────────────────────────────────

	p.AddResourceConfigurator("harness_platform_idp_catalog_entity", func(r *config.Resource) {
		r.ShortGroup = "idp"
		r.Kind = "CatalogEntity"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_idp_environment", func(r *config.Resource) {
		r.ShortGroup = "idp"
		r.Kind = "Environment"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_idp_environment_blueprint", func(r *config.Resource) {
		r.ShortGroup = "idp"
		r.Kind = "EnvironmentBlueprint"
	})

	// ── Platform: Infra (IaCM / Module Registry) ──────────────────────────────

	p.AddResourceConfigurator("harness_platform_infra_module", func(r *config.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "Module"
	})
	p.AddResourceConfigurator("harness_platform_infra_module_testing", func(r *config.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "ModuleTesting"
	})
	p.AddResourceConfigurator("harness_platform_infra_provider", func(r *config.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "Provider"
	})
	p.AddResourceConfigurator("harness_platform_infra_provider_signing_key", func(r *config.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "ProviderSigningKey"
	})
	p.AddResourceConfigurator("harness_platform_infra_provider_version", func(r *config.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "ProviderVersion"
	})
	p.AddResourceConfigurator("harness_platform_infra_provider_version_file", func(r *config.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "ProviderVersionFile"
	})
	p.AddResourceConfigurator("harness_platform_infra_provider_version_publish", func(r *config.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "ProviderVersionPublish"
	})
	p.AddResourceConfigurator("harness_platform_infra_variable_set", func(r *config.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "VariableSet"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})

	// ── Platform: Core resources ──────────────────────────────────────────────

	p.AddResourceConfigurator("harness_platform_apikey", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "ApiKey"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_ccm_filters", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "CcmFilters"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_central_notification_channel", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "CentralNotificationChannel"
	})
	p.AddResourceConfigurator("harness_platform_central_notification_rule", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "CentralNotificationRule"
	})
	p.AddResourceConfigurator("harness_platform_dashboard_folders", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "DashboardFolders"
	})
	p.AddResourceConfigurator("harness_platform_dashboards", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Dashboards"
	})
	p.AddResourceConfigurator("harness_platform_db_instance", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "DbInstance"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["schema"] = dbSchemaRef
	})
	p.AddResourceConfigurator("harness_platform_db_schema", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "DbSchema"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_default_images", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "DefaultImages"
	})
	p.AddResourceConfigurator("harness_platform_default_notification_template_set", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "DefaultNotificationTemplateSet"
	})
	p.AddResourceConfigurator("harness_platform_delegatetoken", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "DelegateToken"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_environment", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Environment"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_environment_clusters_mapping", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "EnvironmentClustersMapping"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["env_id"] = envRef
	})
	p.AddResourceConfigurator("harness_platform_environment_group", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "EnvironmentGroup"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_environment_service_overrides", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "EnvironmentServiceOverrides"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["env_id"] = envRef
		r.References["service_id"] = serviceRef
	})
	p.AddResourceConfigurator("harness_platform_feature_flag", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "FeatureFlag"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_feature_flag_target", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "FeatureFlagTarget"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["environment"] = envRef
	})
	p.AddResourceConfigurator("harness_platform_feature_flag_target_group", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "FeatureFlagTargetGroup"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["environment"] = envRef
	})
	p.AddResourceConfigurator("harness_platform_ff_api_key", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "FfApiKey"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["env_id"] = envRef
	})
	p.AddResourceConfigurator("harness_platform_file_store_file", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "FileStoreFile"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_file_store_folder", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "FileStoreFolder"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_filters", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Filters"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_gitx_webhook", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "GitxWebhook"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_har_registry", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "HarRegistry"
	})
	p.AddResourceConfigurator("harness_platform_iacm_default_pipeline", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "IacmDefaultPipeline"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_infrastructure", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Infrastructure"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["env_id"] = envRef
	})
	p.AddResourceConfigurator("harness_platform_input_set", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "InputSet"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["pipeline_id"] = pipelineRef
	})
	p.AddResourceConfigurator("harness_platform_ip_allowlist", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "IpAllowlist"
	})
	p.AddResourceConfigurator("harness_platform_manual_freeze", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "ManualFreeze"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_monitored_service", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "MonitoredService"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_notification_rule", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "NotificationRule"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_overrides", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Overrides"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["env_id"] = envRef
		r.References["service_id"] = serviceRef
		r.References["infra_id"] = infraRef
	})
	p.AddResourceConfigurator("harness_platform_pipeline", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Pipeline"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_pipeline_central_notification_rule", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "PipelineCentralNotificationRule"
	})
	p.AddResourceConfigurator("harness_platform_pipeline_filters", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "PipelineFilters"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_policy", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Policy"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_policyset", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "PolicySet"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_provider", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Provider"
	})
	p.AddResourceConfigurator("harness_platform_repo", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Repo"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_repo_rule_branch", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "RepoRuleBranch"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_repo_webhook", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "RepoWebhook"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_resource_group", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "ResourceGroup"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_role_assignments", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "RoleAssignments"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_roles", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Roles"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_service", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Service"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_service_account", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "ServiceAccount"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_service_overrides_v2", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "ServiceOverridesV2"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["env_id"] = envRef
		r.References["service_id"] = serviceRef
		r.References["infra_id"] = infraRef
	})
	p.AddResourceConfigurator("harness_platform_slo", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Slo"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_template", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Template"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_template_filters", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "TemplateFilters"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_token", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Token"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["apikey_id"] = apikeyRef
	})
	p.AddResourceConfigurator("harness_platform_triggers", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Triggers"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
		r.References["target_id"] = pipelineRef
	})
	p.AddResourceConfigurator("harness_platform_user", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "User"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_usergroup", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "UserGroup"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_variables", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Variables"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
	p.AddResourceConfigurator("harness_platform_workspace", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Workspace"
		r.References["org_id"] = orgRef
		r.References["project_id"] = projectRef
	})
}
