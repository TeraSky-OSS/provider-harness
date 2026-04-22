package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// ── Legacy / CG resources ────────────────────────────────────────────────

	// terraform import harness_add_user_to_group.example <user_id>/<group_id>
	"harness_add_user_to_group": config.IdentifierFromProvider,
	// terraform import harness_application.myapp Xyz123
	"harness_application": config.IdentifierFromProvider,
	// terraform import harness_application_gitsync.myapp Xyz123
	"harness_application_gitsync": config.IdentifierFromProvider,
	// terraform import harness_delegate_approval.example <delegate_id>
	"harness_delegate_approval": config.IdentifierFromProvider,
	// terraform import harness_encrypted_text.example <secret_id>
	"harness_encrypted_text": config.IdentifierFromProvider,
	// terraform import harness_environment.dev <application_id>/<environment_id>
	"harness_environment": config.IdentifierFromProvider,
	// terraform import harness_git_connector.example <connector_id>
	"harness_git_connector": config.IdentifierFromProvider,
	// terraform import harness_infrastructure_definition.example <app_id>/<env_id>/<infradef_id>
	"harness_infrastructure_definition": config.IdentifierFromProvider,
	// terraform import harness_ssh_credential.example <credential_id>
	"harness_ssh_credential": config.IdentifierFromProvider,
	// terraform import harness_user.john_doe john.doe@example.com
	"harness_user": config.IdentifierFromProvider,
	// terraform import harness_user_group.example <USER_GROUP_ID>
	"harness_user_group": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_user_group_permissions": config.IdentifierFromProvider,
	// terraform import harness_yaml_config.k8s_cloudprovider "Setup/Cloud Providers/kubernetes.yaml"
	"harness_yaml_config": config.IdentifierFromProvider,

	// ── AutoStopping resources ────────────────────────────────────────────────

	// no import syntax documented
	"harness_autostopping_alert":            config.IdentifierFromProvider,
	"harness_autostopping_aws_alb":          config.IdentifierFromProvider,
	"harness_autostopping_aws_proxy":        config.IdentifierFromProvider,
	"harness_autostopping_azure_gateway":    config.IdentifierFromProvider,
	"harness_autostopping_azure_proxy":      config.IdentifierFromProvider,
	"harness_autostopping_gcp_proxy":        config.IdentifierFromProvider,
	"harness_autostopping_rule_ecs":         config.IdentifierFromProvider,
	"harness_autostopping_rule_k8s":         config.IdentifierFromProvider,
	"harness_autostopping_rule_rds":         config.IdentifierFromProvider,
	"harness_autostopping_rule_scale_group": config.IdentifierFromProvider,
	"harness_autostopping_rule_vm":          config.IdentifierFromProvider,
	"harness_autostopping_schedule":         config.IdentifierFromProvider,

	// ── Chaos resources ───────────────────────────────────────────────────────

	// terraform import harness_chaos_action_template.example <org_id>/<project_id>/<template_id>
	"harness_chaos_action_template": config.IdentifierFromProvider,
	// terraform import harness_chaos_experiment.basic my_org/my_project/my-experiment
	"harness_chaos_experiment": config.IdentifierFromProvider,
	// terraform import harness_chaos_experiment_template.example <org_id>/<project_id>/<template_id>
	"harness_chaos_experiment_template": config.IdentifierFromProvider,
	// terraform import harness_chaos_fault_template.example <org_id>/<project_id>/<template_id>
	"harness_chaos_fault_template": config.IdentifierFromProvider,
	// terraform import harness_chaos_hub.example <org_id>/<project_id>/<hub_id>
	"harness_chaos_hub": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_chaos_hub_sync": config.IdentifierFromProvider,
	// terraform import harness_chaos_hub_v2.project_level <org_id>/<project_id>/<hub_id>
	"harness_chaos_hub_v2": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_chaos_image_registry": config.IdentifierFromProvider,
	// terraform import harness_chaos_infrastructure.example <chaos_infra_id>
	"harness_chaos_infrastructure": config.IdentifierFromProvider,
	// terraform import harness_chaos_infrastructure_v2.example <org_id>/<project_id>/<environment_id>/<infra_id>
	"harness_chaos_infrastructure_v2": config.IdentifierFromProvider,
	// terraform import harness_chaos_probe_template.example <org_id>/<project_id>/<template_id>
	"harness_chaos_probe_template": config.IdentifierFromProvider,
	// terraform import harness_chaos_security_governance_condition.example org_id/project_id/condition_id
	"harness_chaos_security_governance_condition": config.IdentifierFromProvider,
	// terraform import harness_chaos_security_governance_rule.example org_id/project_id/rule_id
	"harness_chaos_security_governance_rule": config.IdentifierFromProvider,

	// ── Cloud Provider (legacy CG) resources ─────────────────────────────────

	// terraform import harness_cloudprovider_aws.example <provider_id>
	"harness_cloudprovider_aws": config.IdentifierFromProvider,
	// terraform import harness_cloudprovider_azure.example <provider_id>
	"harness_cloudprovider_azure": config.IdentifierFromProvider,
	// terraform import harness_cloudprovider_datacenter.example <provider_id>
	"harness_cloudprovider_datacenter": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_cloudprovider_gcp": config.IdentifierFromProvider,
	// terraform import harness_cloudprovider_kubernetes.example <provider_id>
	"harness_cloudprovider_kubernetes": config.IdentifierFromProvider,
	// terraform import harness_cloudprovider_spot.example <provider_id>
	"harness_cloudprovider_spot": config.IdentifierFromProvider,
	// terraform import harness_cloudprovider_tanzu.example <provider_id>
	"harness_cloudprovider_tanzu": config.IdentifierFromProvider,

	// ── Cluster Orchestrator resources ────────────────────────────────────────

	// no import syntax documented
	"harness_cluster_orchestrator":        config.IdentifierFromProvider,
	"harness_cluster_orchestrator_config": config.IdentifierFromProvider,

	// ── Governance resources ──────────────────────────────────────────────────

	// terraform import harness_governance_rule.example <rule_id>
	"harness_governance_rule": config.IdentifierFromProvider,
	// terraform import harness_governance_rule_enforcement.example <enforcement_id>
	"harness_governance_rule_enforcement": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_governance_rule_set": config.IdentifierFromProvider,

	// ── Legacy NG Service resources ───────────────────────────────────────────

	// terraform import harness_service_ami.example <app_id>/<svc_id>
	"harness_service_ami": config.IdentifierFromProvider,
	// terraform import harness_service_aws_codedeploy.example <app_id>/<svc_id>
	"harness_service_aws_codedeploy": config.IdentifierFromProvider,
	// terraform import harness_service_aws_lambda.example <app_id>/<svc_id>
	"harness_service_aws_lambda": config.IdentifierFromProvider,
	// terraform import harness_service_discovery_agent.example <org_id>/<project_id>/<env_id>/<infra_id>
	"harness_service_discovery_agent": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_service_discovery_setting": config.IdentifierFromProvider,
	// terraform import harness_service_ecs.example <app_id>/<svc_id>
	"harness_service_ecs": config.IdentifierFromProvider,
	// terraform import harness_service_helm.example <app_id>/<svc_id>
	"harness_service_helm": config.IdentifierFromProvider,
	// terraform import harness_service_kubernetes.example <app_id>/<svc_id>
	"harness_service_kubernetes": config.IdentifierFromProvider,
	// terraform import harness_service_ssh.example <app_id>/<svc_id>
	"harness_service_ssh": config.IdentifierFromProvider,
	// terraform import harness_service_tanzu.example <app_id>/<svc_id>
	"harness_service_tanzu": config.IdentifierFromProvider,
	// terraform import harness_service_winrm.example <app_id>/<svc_id>
	"harness_service_winrm": config.IdentifierFromProvider,

	// ── Platform: Organization / Project ─────────────────────────────────────

	// terraform import harness_platform_organization.example <organization_id>
	"harness_platform_organization": config.IdentifierFromProvider,
	// terraform import harness_platform_project.example <organization_id>/<project_id>
	"harness_platform_project": config.IdentifierFromProvider,

	// ── Platform: Connectors ──────────────────────────────────────────────────
	// All connectors support account / org / project import:
	//   <connector_id>
	//   <org_id>/<connector_id>
	//   <org_id>/<project_id>/<connector_id>

	"harness_platform_connector_appdynamics":        config.IdentifierFromProvider,
	"harness_platform_connector_artifactory":        config.IdentifierFromProvider,
	"harness_platform_connector_aws":                config.IdentifierFromProvider,
	"harness_platform_connector_aws_secret_manager": config.IdentifierFromProvider,
	"harness_platform_connector_awscc":              config.IdentifierFromProvider,
	"harness_platform_connector_awskms":             config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_connector_azure_artifacts":      config.IdentifierFromProvider,
	"harness_platform_connector_azure_cloud_cost":     config.IdentifierFromProvider,
	"harness_platform_connector_azure_cloud_provider": config.IdentifierFromProvider,
	"harness_platform_connector_azure_key_vault":      config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_connector_azure_repo":            config.IdentifierFromProvider,
	"harness_platform_connector_bitbucket":             config.IdentifierFromProvider,
	"harness_platform_connector_custom_secret_manager": config.IdentifierFromProvider,
	"harness_platform_connector_customhealthsource":    config.IdentifierFromProvider,
	"harness_platform_connector_datadog":               config.IdentifierFromProvider,
	"harness_platform_connector_docker":                config.IdentifierFromProvider,
	"harness_platform_connector_dynatrace":             config.IdentifierFromProvider,
	"harness_platform_connector_elasticsearch":         config.IdentifierFromProvider,
	"harness_platform_connector_gcp":                   config.IdentifierFromProvider,
	"harness_platform_connector_gcp_cloud_cost":        config.IdentifierFromProvider,
	"harness_platform_connector_gcp_kms":               config.IdentifierFromProvider,
	"harness_platform_connector_gcp_secret_manager":    config.IdentifierFromProvider,
	"harness_platform_connector_git":                   config.IdentifierFromProvider,
	"harness_platform_connector_github":                config.IdentifierFromProvider,
	"harness_platform_connector_gitlab":                config.IdentifierFromProvider,
	"harness_platform_connector_helm":                  config.IdentifierFromProvider,
	"harness_platform_connector_jdbc":                  config.IdentifierFromProvider,
	"harness_platform_connector_jenkins":               config.IdentifierFromProvider,
	"harness_platform_connector_jira":                  config.IdentifierFromProvider,
	"harness_platform_connector_kubernetes":            config.IdentifierFromProvider,
	"harness_platform_connector_kubernetes_cloud_cost": config.IdentifierFromProvider,
	"harness_platform_connector_newrelic":              config.IdentifierFromProvider,
	"harness_platform_connector_nexus":                 config.IdentifierFromProvider,
	"harness_platform_connector_oci_helm":              config.IdentifierFromProvider,
	"harness_platform_connector_pagerduty":             config.IdentifierFromProvider,
	"harness_platform_connector_pdc":                   config.IdentifierFromProvider,
	"harness_platform_connector_prometheus":            config.IdentifierFromProvider,
	"harness_platform_connector_rancher":               config.IdentifierFromProvider,
	"harness_platform_connector_service_now":           config.IdentifierFromProvider,
	"harness_platform_connector_splunk":                config.IdentifierFromProvider,
	"harness_platform_connector_spot":                  config.IdentifierFromProvider,
	"harness_platform_connector_sumologic":             config.IdentifierFromProvider,
	"harness_platform_connector_tas":                   config.IdentifierFromProvider,
	"harness_platform_connector_terraform_cloud":       config.IdentifierFromProvider,
	"harness_platform_connector_vault":                 config.IdentifierFromProvider,

	// ── Platform: Secrets ─────────────────────────────────────────────────────
	// All secrets support account / org / project import:
	//   <secret_id>
	//   <org_id>/<secret_id>
	//   <org_id>/<project_id>/<secret_id>

	"harness_platform_secret_file":   config.IdentifierFromProvider,
	"harness_platform_secret_sshkey": config.IdentifierFromProvider,
	"harness_platform_secret_text":   config.IdentifierFromProvider,
	"harness_platform_secret_winrm":  config.IdentifierFromProvider,

	// ── Platform: GitOps ──────────────────────────────────────────────────────

	// terraform import harness_platform_gitops_agent.example <org_id>/<project_id>/<agent_id>
	"harness_platform_gitops_agent": config.IdentifierFromProvider,
	// terraform import harness_platform_gitops_app_project.example <org_id>/<project_id>/<agent_id>/<app_project_name>
	"harness_platform_gitops_app_project": config.IdentifierFromProvider,
	// terraform import harness_platform_gitops_app_project_mapping.example <org_id>/<project_id>/<agent_id>/<app_project_name>
	"harness_platform_gitops_app_project_mapping": config.IdentifierFromProvider,
	// terraform import harness_platform_gitops_applications.example <org_id>/<project_id>/<agent_id>/<app_name>
	"harness_platform_gitops_applications": config.IdentifierFromProvider,
	// terraform import harness_platform_gitops_cluster.example <org_id>/<project_id>/<agent_id>/<cluster_id>
	"harness_platform_gitops_cluster": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_gitops_filters": config.IdentifierFromProvider,
	// terraform import harness_platform_gitops_gnupg.example <org_id>/<project_id>/<agent_id>/<gnupg_key_id>
	"harness_platform_gitops_gnupg": config.IdentifierFromProvider,
	// terraform import harness_platform_gitops_repo_cert.example <org_id>/<project_id>/<agent_id>
	"harness_platform_gitops_repo_cert": config.IdentifierFromProvider,
	// terraform import harness_platform_gitops_repo_cred.example <org_id>/<project_id>/<agent_id>/<repo_cred_id>
	"harness_platform_gitops_repo_cred": config.IdentifierFromProvider,
	// terraform import harness_platform_gitops_repository.example <org_id>/<project_id>/<agent_id>/<repo_id>
	"harness_platform_gitops_repository": config.IdentifierFromProvider,

	// ── Platform: IDP ─────────────────────────────────────────────────────────

	// terraform import harness_platform_idp_catalog_entity.example <account_id>/<entity_id>
	"harness_platform_idp_catalog_entity": config.IdentifierFromProvider,
	// terraform import harness_platform_idp_environment.example <account_id>/<env_id>
	"harness_platform_idp_environment": config.IdentifierFromProvider,
	// terraform import harness_platform_idp_environment_blueprint.example <account_id>/<blueprint_id>
	"harness_platform_idp_environment_blueprint": config.IdentifierFromProvider,

	// ── Platform: Infra (IaCM / Module Registry) ──────────────────────────────

	// terraform import harness_platform_infra_module.example <org_id>/<project_id>/<module_id>
	"harness_platform_infra_module": config.IdentifierFromProvider,
	// terraform import harness_platform_infra_module_testing.example <org_id>/<project_id>/<module_id>/<version_id>/<testing_id>
	"harness_platform_infra_module_testing": config.IdentifierFromProvider,
	// terraform import harness_platform_infra_provider.example <org_id>/<project_id>/<provider_id>
	"harness_platform_infra_provider": config.IdentifierFromProvider,
	// terraform import harness_platform_infra_provider_signing_key.example <org_id>/<project_id>/<provider_id>/<key_id>
	"harness_platform_infra_provider_signing_key": config.IdentifierFromProvider,
	// terraform import harness_platform_infra_provider_version.example <org_id>/<project_id>/<provider_id>/<version_id>
	"harness_platform_infra_provider_version": config.IdentifierFromProvider,
	// terraform import harness_platform_infra_provider_version_file.example <org_id>/<project_id>/<provider_id>/<version_id>/<filename>
	"harness_platform_infra_provider_version_file": config.IdentifierFromProvider,
	// terraform import harness_platform_infra_provider_version_publish.example <org_id>/<project_id>/<provider_id>/<version_id>
	"harness_platform_infra_provider_version_publish": config.IdentifierFromProvider,
	// terraform import harness_platform_infra_variable_set.example <org_id>/<project_id>/<variable_set_id>
	"harness_platform_infra_variable_set": config.IdentifierFromProvider,

	// ── Platform: Core resources ──────────────────────────────────────────────

	// terraform import harness_platform_apikey <parent_id>/<apikey_id>/<apikey_type>
	"harness_platform_apikey": config.IdentifierFromProvider,
	// terraform import harness_platform_ccm_filters.example <filter_id>/<type>
	"harness_platform_ccm_filters": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_central_notification_channel": config.IdentifierFromProvider,
	"harness_platform_central_notification_rule":    config.IdentifierFromProvider,
	// terraform import harness_platform_dashboard_folders.example <folder_id>
	"harness_platform_dashboard_folders": config.IdentifierFromProvider,
	// terraform import harness_platform_dashboards.example <dashboard_id>
	"harness_platform_dashboards": config.IdentifierFromProvider,
	// terraform import harness_platform_db_instance.example <org_id>/<project_id>/<db_instance_id>
	"harness_platform_db_instance": config.IdentifierFromProvider,
	// terraform import harness_platform_db_schema.example <org_id>/<project_id>/<db_schema_id>
	"harness_platform_db_schema": config.IdentifierFromProvider,
	// terraform import harness_platform_default_images.example <account_id>
	"harness_platform_default_images": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_default_notification_template_set": config.IdentifierFromProvider,
	// terraform import harness_platform_delegatetoken.example <org_id>/<project_id>/<delegate_token_id>
	"harness_platform_delegatetoken": config.IdentifierFromProvider,
	// terraform import harness_platform_environment.example <org_id>/<project_id>/<environment_id>
	"harness_platform_environment": config.IdentifierFromProvider,
	// terraform import harness_platform_environment_clusters_mapping.example <org_id>/<project_id>/<env_id>
	"harness_platform_environment_clusters_mapping": config.IdentifierFromProvider,
	// terraform import harness_platform_environment_group.example <org_id>/<project_id>/<env_group_id>
	"harness_platform_environment_group": config.IdentifierFromProvider,
	// terraform import harness_platform_environment_service_overrides.example <org_id>/<project_id>/<env_id>
	"harness_platform_environment_service_overrides": config.IdentifierFromProvider,
	// terraform import harness_platform_feature_flag.example <org_id>/<project_id>/<feature_flag_id>
	"harness_platform_feature_flag": config.IdentifierFromProvider,
	// terraform import harness_platform_feature_flag_target.example <org_id>/<project_id>/<env_id>/<target_id>
	"harness_platform_feature_flag_target": config.IdentifierFromProvider,
	// terraform import harness_platform_feature_flag_target_group.example <org_id>/<project_id>/<env_id>/<target_group_id>
	"harness_platform_feature_flag_target_group": config.IdentifierFromProvider,
	// terraform import harness_platform_ff_api_key.example <org_id>/<project_id>/<env_id>/<key_id>
	"harness_platform_ff_api_key": config.IdentifierFromProvider,
	// terraform import harness_platform_file_store_file.example <org_id>/<project_id>/<file_id>
	"harness_platform_file_store_file": config.IdentifierFromProvider,
	// terraform import harness_platform_file_store_folder.example <org_id>/<project_id>/<folder_id>
	"harness_platform_file_store_folder": config.IdentifierFromProvider,
	// terraform import harness_platform_filters.example <filter_id>/<type>
	"harness_platform_filters": config.IdentifierFromProvider,
	// terraform import harness_platform_gitx_webhook.example <org_id>/<project_id>/<webhook_id>
	"harness_platform_gitx_webhook": config.IdentifierFromProvider,
	// terraform import harness_platform_har_registry.example <space_ref>/<registry_identifier>
	"harness_platform_har_registry": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_iacm_default_pipeline": config.IdentifierFromProvider,
	// terraform import harness_platform_infrastructure.example <org_id>/<project_id>/<env_id>/<infrastructure_id>
	"harness_platform_infrastructure": config.IdentifierFromProvider,
	// terraform import harness_platform_input_set.example <org_id>/<project_id>/<pipeline_id>/<input_set_id>
	"harness_platform_input_set": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_ip_allowlist": config.IdentifierFromProvider,
	// terraform import harness_platform_manual_freeze.example <org_id>/<project_id>/<freeze_id>
	"harness_platform_manual_freeze": config.IdentifierFromProvider,
	// terraform import harness_platform_monitored_service.example <org_id>/<project_id>/<monitored_service_id>
	"harness_platform_monitored_service": config.IdentifierFromProvider,
	// terraform import harness_platform_notification_rule.example <org_id>/<project_id>/<notification_rule_id>
	"harness_platform_notification_rule": config.IdentifierFromProvider,
	// terraform import harness_platform_overrides.example <org_id>/<project_id>/<override_id>
	"harness_platform_overrides": config.IdentifierFromProvider,
	// terraform import harness_platform_pipeline.example <org_id>/<project_id>/<pipeline_id>
	"harness_platform_pipeline": config.IdentifierFromProvider,
	// terraform import harness_platform_pipeline_central_notification_rule.example <org_id>/<project_id>/<notification_rule_id>
	"harness_platform_pipeline_central_notification_rule": config.IdentifierFromProvider,
	// terraform import harness_platform_pipeline_filters.example <filter_id>/<type>
	"harness_platform_pipeline_filters": config.IdentifierFromProvider,
	// terraform import harness_platform_policy.example <policy_id>
	"harness_platform_policy": config.IdentifierFromProvider,
	// terraform import harness_platform_policyset.example <policyset_id>
	"harness_platform_policyset": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_provider": config.IdentifierFromProvider,
	// terraform import harness_platform_repo.example <org_id>/<project_id>/<identifier>
	"harness_platform_repo": config.IdentifierFromProvider,
	// no import syntax documented
	"harness_platform_repo_rule_branch": config.IdentifierFromProvider,
	"harness_platform_repo_webhook":     config.IdentifierFromProvider,
	// terraform import harness_platform_resource_group.example <org_id>/<project_id>/<resource_group_id>
	"harness_platform_resource_group": config.IdentifierFromProvider,
	// terraform import harness_platform_role_assignments.example <org_id>/<project_id>/<role_assignments_id>
	"harness_platform_role_assignments": config.IdentifierFromProvider,
	// terraform import harness_platform_roles.example <org_id>/<project_id>/<roles_id>
	"harness_platform_roles": config.IdentifierFromProvider,
	// terraform import harness_platform_service.example <org_id>/<project_id>/<service_id>
	"harness_platform_service": config.IdentifierFromProvider,
	// terraform import harness_platform_service_account.example <org_id>/<project_id>/<service_account_id>
	"harness_platform_service_account": config.IdentifierFromProvider,
	// terraform import harness_platform_service_overrides_v2.example <org_id>/<project_id>/<override_id>
	"harness_platform_service_overrides_v2": config.IdentifierFromProvider,
	// terraform import harness_platform_slo.example <org_id>/<project_id>/<slo_id>
	"harness_platform_slo": config.IdentifierFromProvider,
	// terraform import harness_platform_template.example <org_id>/<project_id>/<template_id>
	"harness_platform_template": config.IdentifierFromProvider,
	// terraform import harness_platform_template_filters.example <filter_id>/<type>
	"harness_platform_template_filters": config.IdentifierFromProvider,
	// terraform import harness_platform_token <org_id>/<project_id>/<parent_id>/<apikey_id>/<apikey_type>/<token_id>
	"harness_platform_token": config.IdentifierFromProvider,
	// terraform import harness_platform_triggers.example <org_id>/<project_id>/<target_id>/<triggers_id>
	"harness_platform_triggers": config.IdentifierFromProvider,
	// terraform import harness_platform_user.john_doe <email_id>
	"harness_platform_user": config.IdentifierFromProvider,
	// terraform import harness_platform_usergroup.example <org_id>/<project_id>/<usergroup_id>
	"harness_platform_usergroup": config.IdentifierFromProvider,
	// terraform import harness_platform_variables.example <org_id>/<project_id>/<variable_id>
	"harness_platform_variables": config.IdentifierFromProvider,
	// terraform import harness_platform_workspace.example <org_id>/<project_id>/<workspace_id>
	"harness_platform_workspace": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
