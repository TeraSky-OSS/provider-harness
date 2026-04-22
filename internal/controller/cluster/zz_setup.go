// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alert "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/alert"
	awsalb "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/awsalb"
	awsproxy "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/awsproxy"
	azuregateway "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/azuregateway"
	azureproxy "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/azureproxy"
	gcpproxy "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/gcpproxy"
	ruleecs "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/ruleecs"
	rulek8s "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/rulek8s"
	rulerds "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/rulerds"
	rulescalegroup "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/rulescalegroup"
	rulevm "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/rulevm"
	schedule "github.com/terasky-oss/provider-harness/internal/controller/cluster/autostopping/schedule"
	actiontemplate "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/actiontemplate"
	experiment "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/experiment"
	experimenttemplate "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/experimenttemplate"
	faulttemplate "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/faulttemplate"
	hub "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/hub"
	hubsync "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/hubsync"
	hubv2 "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/hubv2"
	imageregistry "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/imageregistry"
	infrastructure "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/infrastructure"
	infrastructurev2 "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/infrastructurev2"
	probetemplate "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/probetemplate"
	securitygovernancecondition "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/securitygovernancecondition"
	securitygovernancerule "github.com/terasky-oss/provider-harness/internal/controller/cluster/chaos/securitygovernancerule"
	aws "github.com/terasky-oss/provider-harness/internal/controller/cluster/cloudprovider/aws"
	azure "github.com/terasky-oss/provider-harness/internal/controller/cluster/cloudprovider/azure"
	datacenter "github.com/terasky-oss/provider-harness/internal/controller/cluster/cloudprovider/datacenter"
	gcp "github.com/terasky-oss/provider-harness/internal/controller/cluster/cloudprovider/gcp"
	kubernetes "github.com/terasky-oss/provider-harness/internal/controller/cluster/cloudprovider/kubernetes"
	spot "github.com/terasky-oss/provider-harness/internal/controller/cluster/cloudprovider/spot"
	tanzu "github.com/terasky-oss/provider-harness/internal/controller/cluster/cloudprovider/tanzu"
	orchestrator "github.com/terasky-oss/provider-harness/internal/controller/cluster/cluster/orchestrator"
	orchestratorconfig "github.com/terasky-oss/provider-harness/internal/controller/cluster/cluster/orchestratorconfig"
	appdynamics "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/appdynamics"
	artifactory "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/artifactory"
	awsconnector "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/aws"
	awscc "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/awscc"
	awskms "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/awskms"
	awssecretmanager "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/awssecretmanager"
	azureartifacts "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/azureartifacts"
	azurecloudcost "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/azurecloudcost"
	azurecloudprovider "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/azurecloudprovider"
	azurekeyvault "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/azurekeyvault"
	azurerepo "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/azurerepo"
	bitbucket "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/bitbucket"
	customhealthsource "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/customhealthsource"
	customsecretmanager "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/customsecretmanager"
	datadog "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/datadog"
	docker "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/docker"
	dynatrace "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/dynatrace"
	elasticsearch "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/elasticsearch"
	gcpconnector "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/gcp"
	gcpcloudcost "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/gcpcloudcost"
	gcpkms "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/gcpkms"
	gcpsecretmanager "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/gcpsecretmanager"
	git "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/git"
	github "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/github"
	gitlab "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/gitlab"
	helm "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/helm"
	jdbc "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/jdbc"
	jenkins "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/jenkins"
	jira "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/jira"
	kubernetesconnector "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/kubernetes"
	kubernetescloudcost "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/kubernetescloudcost"
	newrelic "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/newrelic"
	nexus "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/nexus"
	ocihelm "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/ocihelm"
	pagerduty "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/pagerduty"
	pdc "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/pdc"
	prometheus "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/prometheus"
	rancher "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/rancher"
	servicenow "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/servicenow"
	splunk "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/splunk"
	spotconnector "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/spot"
	sumologic "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/sumologic"
	tas "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/tas"
	terraformcloud "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/terraformcloud"
	vault "github.com/terasky-oss/provider-harness/internal/controller/cluster/connector/vault"
	agent "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/agent"
	applications "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/applications"
	appproject "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/appproject"
	appprojectmapping "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/appprojectmapping"
	cluster "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/cluster"
	filters "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/filters"
	gnupg "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/gnupg"
	repocert "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/repocert"
	repocred "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/repocred"
	repository "github.com/terasky-oss/provider-harness/internal/controller/cluster/gitops/repository"
	rule "github.com/terasky-oss/provider-harness/internal/controller/cluster/governance/rule"
	ruleenforcement "github.com/terasky-oss/provider-harness/internal/controller/cluster/governance/ruleenforcement"
	ruleset "github.com/terasky-oss/provider-harness/internal/controller/cluster/governance/ruleset"
	addusertogroup "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/addusertogroup"
	application "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/application"
	applicationgitsync "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/applicationgitsync"
	delegateapproval "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/delegateapproval"
	encryptedtext "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/encryptedtext"
	environment "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/environment"
	gitconnector "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/gitconnector"
	infrastructuredefinition "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/infrastructuredefinition"
	sshcredential "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/sshcredential"
	user "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/user"
	usergroup "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/usergroup"
	usergrouppermissions "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/usergrouppermissions"
	yamlconfig "github.com/terasky-oss/provider-harness/internal/controller/cluster/harness/yamlconfig"
	catalogentity "github.com/terasky-oss/provider-harness/internal/controller/cluster/idp/catalogentity"
	environmentidp "github.com/terasky-oss/provider-harness/internal/controller/cluster/idp/environment"
	environmentblueprint "github.com/terasky-oss/provider-harness/internal/controller/cluster/idp/environmentblueprint"
	module "github.com/terasky-oss/provider-harness/internal/controller/cluster/infra/module"
	moduletesting "github.com/terasky-oss/provider-harness/internal/controller/cluster/infra/moduletesting"
	provider "github.com/terasky-oss/provider-harness/internal/controller/cluster/infra/provider"
	providersigningkey "github.com/terasky-oss/provider-harness/internal/controller/cluster/infra/providersigningkey"
	providerversion "github.com/terasky-oss/provider-harness/internal/controller/cluster/infra/providerversion"
	providerversionfile "github.com/terasky-oss/provider-harness/internal/controller/cluster/infra/providerversionfile"
	providerversionpublish "github.com/terasky-oss/provider-harness/internal/controller/cluster/infra/providerversionpublish"
	variableset "github.com/terasky-oss/provider-harness/internal/controller/cluster/infra/variableset"
	apikey "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/apikey"
	ccmfilters "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/ccmfilters"
	centralnotificationchannel "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/centralnotificationchannel"
	centralnotificationrule "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/centralnotificationrule"
	dashboardfolders "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/dashboardfolders"
	dashboards "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/dashboards"
	dbinstance "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/dbinstance"
	dbschema "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/dbschema"
	defaultimages "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/defaultimages"
	defaultnotificationtemplateset "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/defaultnotificationtemplateset"
	delegatetoken "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/delegatetoken"
	environmentplatform "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/environment"
	environmentclustersmapping "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/environmentclustersmapping"
	environmentgroup "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/environmentgroup"
	environmentserviceoverrides "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/environmentserviceoverrides"
	featureflag "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/featureflag"
	featureflagtarget "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/featureflagtarget"
	featureflagtargetgroup "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/featureflagtargetgroup"
	ffapikey "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/ffapikey"
	filestorefile "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/filestorefile"
	filestorefolder "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/filestorefolder"
	filtersplatform "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/filters"
	gitxwebhook "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/gitxwebhook"
	harregistry "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/harregistry"
	iacmdefaultpipeline "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/iacmdefaultpipeline"
	infrastructureplatform "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/infrastructure"
	inputset "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/inputset"
	ipallowlist "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/ipallowlist"
	manualfreeze "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/manualfreeze"
	monitoredservice "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/monitoredservice"
	notificationrule "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/notificationrule"
	organization "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/organization"
	overrides "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/overrides"
	pipeline "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/pipeline"
	pipelinecentralnotificationrule "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/pipelinecentralnotificationrule"
	pipelinefilters "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/pipelinefilters"
	policy "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/policy"
	policyset "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/policyset"
	project "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/project"
	providerplatform "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/provider"
	repo "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/repo"
	reporulebranch "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/reporulebranch"
	repowebhook "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/repowebhook"
	resourcegroup "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/resourcegroup"
	roleassignments "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/roleassignments"
	roles "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/roles"
	service "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/service"
	serviceaccount "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/serviceaccount"
	serviceoverridesv2 "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/serviceoverridesv2"
	slo "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/slo"
	template "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/template"
	templatefilters "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/templatefilters"
	token "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/token"
	triggers "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/triggers"
	userplatform "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/user"
	usergroupplatform "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/usergroup"
	variables "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/variables"
	workspace "github.com/terasky-oss/provider-harness/internal/controller/cluster/platform/workspace"
	providerconfig "github.com/terasky-oss/provider-harness/internal/controller/cluster/providerconfig"
	file "github.com/terasky-oss/provider-harness/internal/controller/cluster/secret/file"
	sshkey "github.com/terasky-oss/provider-harness/internal/controller/cluster/secret/sshkey"
	text "github.com/terasky-oss/provider-harness/internal/controller/cluster/secret/text"
	winrm "github.com/terasky-oss/provider-harness/internal/controller/cluster/secret/winrm"
	ami "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/ami"
	awscodedeploy "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/awscodedeploy"
	awslambda "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/awslambda"
	discoveryagent "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/discoveryagent"
	discoverysetting "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/discoverysetting"
	ecs "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/ecs"
	helmservice "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/helm"
	kubernetesservice "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/kubernetes"
	ssh "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/ssh"
	tanzuservice "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/tanzu"
	winrmservice "github.com/terasky-oss/provider-harness/internal/controller/cluster/service/winrm"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alert.Setup,
		awsalb.Setup,
		awsproxy.Setup,
		azuregateway.Setup,
		azureproxy.Setup,
		gcpproxy.Setup,
		ruleecs.Setup,
		rulek8s.Setup,
		rulerds.Setup,
		rulescalegroup.Setup,
		rulevm.Setup,
		schedule.Setup,
		actiontemplate.Setup,
		experiment.Setup,
		experimenttemplate.Setup,
		faulttemplate.Setup,
		hub.Setup,
		hubsync.Setup,
		hubv2.Setup,
		imageregistry.Setup,
		infrastructure.Setup,
		infrastructurev2.Setup,
		probetemplate.Setup,
		securitygovernancecondition.Setup,
		securitygovernancerule.Setup,
		aws.Setup,
		azure.Setup,
		datacenter.Setup,
		gcp.Setup,
		kubernetes.Setup,
		spot.Setup,
		tanzu.Setup,
		orchestrator.Setup,
		orchestratorconfig.Setup,
		appdynamics.Setup,
		artifactory.Setup,
		awsconnector.Setup,
		awscc.Setup,
		awskms.Setup,
		awssecretmanager.Setup,
		azureartifacts.Setup,
		azurecloudcost.Setup,
		azurecloudprovider.Setup,
		azurekeyvault.Setup,
		azurerepo.Setup,
		bitbucket.Setup,
		customhealthsource.Setup,
		customsecretmanager.Setup,
		datadog.Setup,
		docker.Setup,
		dynatrace.Setup,
		elasticsearch.Setup,
		gcpconnector.Setup,
		gcpcloudcost.Setup,
		gcpkms.Setup,
		gcpsecretmanager.Setup,
		git.Setup,
		github.Setup,
		gitlab.Setup,
		helm.Setup,
		jdbc.Setup,
		jenkins.Setup,
		jira.Setup,
		kubernetesconnector.Setup,
		kubernetescloudcost.Setup,
		newrelic.Setup,
		nexus.Setup,
		ocihelm.Setup,
		pagerduty.Setup,
		pdc.Setup,
		prometheus.Setup,
		rancher.Setup,
		servicenow.Setup,
		splunk.Setup,
		spotconnector.Setup,
		sumologic.Setup,
		tas.Setup,
		terraformcloud.Setup,
		vault.Setup,
		agent.Setup,
		applications.Setup,
		appproject.Setup,
		appprojectmapping.Setup,
		cluster.Setup,
		filters.Setup,
		gnupg.Setup,
		repocert.Setup,
		repocred.Setup,
		repository.Setup,
		rule.Setup,
		ruleenforcement.Setup,
		ruleset.Setup,
		addusertogroup.Setup,
		application.Setup,
		applicationgitsync.Setup,
		delegateapproval.Setup,
		encryptedtext.Setup,
		environment.Setup,
		gitconnector.Setup,
		infrastructuredefinition.Setup,
		sshcredential.Setup,
		user.Setup,
		usergroup.Setup,
		usergrouppermissions.Setup,
		yamlconfig.Setup,
		catalogentity.Setup,
		environmentidp.Setup,
		environmentblueprint.Setup,
		module.Setup,
		moduletesting.Setup,
		provider.Setup,
		providersigningkey.Setup,
		providerversion.Setup,
		providerversionfile.Setup,
		providerversionpublish.Setup,
		variableset.Setup,
		apikey.Setup,
		ccmfilters.Setup,
		centralnotificationchannel.Setup,
		centralnotificationrule.Setup,
		dashboardfolders.Setup,
		dashboards.Setup,
		dbinstance.Setup,
		dbschema.Setup,
		defaultimages.Setup,
		defaultnotificationtemplateset.Setup,
		delegatetoken.Setup,
		environmentplatform.Setup,
		environmentclustersmapping.Setup,
		environmentgroup.Setup,
		environmentserviceoverrides.Setup,
		featureflag.Setup,
		featureflagtarget.Setup,
		featureflagtargetgroup.Setup,
		ffapikey.Setup,
		filestorefile.Setup,
		filestorefolder.Setup,
		filtersplatform.Setup,
		gitxwebhook.Setup,
		harregistry.Setup,
		iacmdefaultpipeline.Setup,
		infrastructureplatform.Setup,
		inputset.Setup,
		ipallowlist.Setup,
		manualfreeze.Setup,
		monitoredservice.Setup,
		notificationrule.Setup,
		organization.Setup,
		overrides.Setup,
		pipeline.Setup,
		pipelinecentralnotificationrule.Setup,
		pipelinefilters.Setup,
		policy.Setup,
		policyset.Setup,
		project.Setup,
		providerplatform.Setup,
		repo.Setup,
		reporulebranch.Setup,
		repowebhook.Setup,
		resourcegroup.Setup,
		roleassignments.Setup,
		roles.Setup,
		service.Setup,
		serviceaccount.Setup,
		serviceoverridesv2.Setup,
		slo.Setup,
		template.Setup,
		templatefilters.Setup,
		token.Setup,
		triggers.Setup,
		userplatform.Setup,
		usergroupplatform.Setup,
		variables.Setup,
		workspace.Setup,
		providerconfig.Setup,
		file.Setup,
		sshkey.Setup,
		text.Setup,
		winrm.Setup,
		ami.Setup,
		awscodedeploy.Setup,
		awslambda.Setup,
		discoveryagent.Setup,
		discoverysetting.Setup,
		ecs.Setup,
		helmservice.Setup,
		kubernetesservice.Setup,
		ssh.Setup,
		tanzuservice.Setup,
		winrmservice.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alert.SetupGated,
		awsalb.SetupGated,
		awsproxy.SetupGated,
		azuregateway.SetupGated,
		azureproxy.SetupGated,
		gcpproxy.SetupGated,
		ruleecs.SetupGated,
		rulek8s.SetupGated,
		rulerds.SetupGated,
		rulescalegroup.SetupGated,
		rulevm.SetupGated,
		schedule.SetupGated,
		actiontemplate.SetupGated,
		experiment.SetupGated,
		experimenttemplate.SetupGated,
		faulttemplate.SetupGated,
		hub.SetupGated,
		hubsync.SetupGated,
		hubv2.SetupGated,
		imageregistry.SetupGated,
		infrastructure.SetupGated,
		infrastructurev2.SetupGated,
		probetemplate.SetupGated,
		securitygovernancecondition.SetupGated,
		securitygovernancerule.SetupGated,
		aws.SetupGated,
		azure.SetupGated,
		datacenter.SetupGated,
		gcp.SetupGated,
		kubernetes.SetupGated,
		spot.SetupGated,
		tanzu.SetupGated,
		orchestrator.SetupGated,
		orchestratorconfig.SetupGated,
		appdynamics.SetupGated,
		artifactory.SetupGated,
		awsconnector.SetupGated,
		awscc.SetupGated,
		awskms.SetupGated,
		awssecretmanager.SetupGated,
		azureartifacts.SetupGated,
		azurecloudcost.SetupGated,
		azurecloudprovider.SetupGated,
		azurekeyvault.SetupGated,
		azurerepo.SetupGated,
		bitbucket.SetupGated,
		customhealthsource.SetupGated,
		customsecretmanager.SetupGated,
		datadog.SetupGated,
		docker.SetupGated,
		dynatrace.SetupGated,
		elasticsearch.SetupGated,
		gcpconnector.SetupGated,
		gcpcloudcost.SetupGated,
		gcpkms.SetupGated,
		gcpsecretmanager.SetupGated,
		git.SetupGated,
		github.SetupGated,
		gitlab.SetupGated,
		helm.SetupGated,
		jdbc.SetupGated,
		jenkins.SetupGated,
		jira.SetupGated,
		kubernetesconnector.SetupGated,
		kubernetescloudcost.SetupGated,
		newrelic.SetupGated,
		nexus.SetupGated,
		ocihelm.SetupGated,
		pagerduty.SetupGated,
		pdc.SetupGated,
		prometheus.SetupGated,
		rancher.SetupGated,
		servicenow.SetupGated,
		splunk.SetupGated,
		spotconnector.SetupGated,
		sumologic.SetupGated,
		tas.SetupGated,
		terraformcloud.SetupGated,
		vault.SetupGated,
		agent.SetupGated,
		applications.SetupGated,
		appproject.SetupGated,
		appprojectmapping.SetupGated,
		cluster.SetupGated,
		filters.SetupGated,
		gnupg.SetupGated,
		repocert.SetupGated,
		repocred.SetupGated,
		repository.SetupGated,
		rule.SetupGated,
		ruleenforcement.SetupGated,
		ruleset.SetupGated,
		addusertogroup.SetupGated,
		application.SetupGated,
		applicationgitsync.SetupGated,
		delegateapproval.SetupGated,
		encryptedtext.SetupGated,
		environment.SetupGated,
		gitconnector.SetupGated,
		infrastructuredefinition.SetupGated,
		sshcredential.SetupGated,
		user.SetupGated,
		usergroup.SetupGated,
		usergrouppermissions.SetupGated,
		yamlconfig.SetupGated,
		catalogentity.SetupGated,
		environmentidp.SetupGated,
		environmentblueprint.SetupGated,
		module.SetupGated,
		moduletesting.SetupGated,
		provider.SetupGated,
		providersigningkey.SetupGated,
		providerversion.SetupGated,
		providerversionfile.SetupGated,
		providerversionpublish.SetupGated,
		variableset.SetupGated,
		apikey.SetupGated,
		ccmfilters.SetupGated,
		centralnotificationchannel.SetupGated,
		centralnotificationrule.SetupGated,
		dashboardfolders.SetupGated,
		dashboards.SetupGated,
		dbinstance.SetupGated,
		dbschema.SetupGated,
		defaultimages.SetupGated,
		defaultnotificationtemplateset.SetupGated,
		delegatetoken.SetupGated,
		environmentplatform.SetupGated,
		environmentclustersmapping.SetupGated,
		environmentgroup.SetupGated,
		environmentserviceoverrides.SetupGated,
		featureflag.SetupGated,
		featureflagtarget.SetupGated,
		featureflagtargetgroup.SetupGated,
		ffapikey.SetupGated,
		filestorefile.SetupGated,
		filestorefolder.SetupGated,
		filtersplatform.SetupGated,
		gitxwebhook.SetupGated,
		harregistry.SetupGated,
		iacmdefaultpipeline.SetupGated,
		infrastructureplatform.SetupGated,
		inputset.SetupGated,
		ipallowlist.SetupGated,
		manualfreeze.SetupGated,
		monitoredservice.SetupGated,
		notificationrule.SetupGated,
		organization.SetupGated,
		overrides.SetupGated,
		pipeline.SetupGated,
		pipelinecentralnotificationrule.SetupGated,
		pipelinefilters.SetupGated,
		policy.SetupGated,
		policyset.SetupGated,
		project.SetupGated,
		providerplatform.SetupGated,
		repo.SetupGated,
		reporulebranch.SetupGated,
		repowebhook.SetupGated,
		resourcegroup.SetupGated,
		roleassignments.SetupGated,
		roles.SetupGated,
		service.SetupGated,
		serviceaccount.SetupGated,
		serviceoverridesv2.SetupGated,
		slo.SetupGated,
		template.SetupGated,
		templatefilters.SetupGated,
		token.SetupGated,
		triggers.SetupGated,
		userplatform.SetupGated,
		usergroupplatform.SetupGated,
		variables.SetupGated,
		workspace.SetupGated,
		providerconfig.SetupGated,
		file.SetupGated,
		sshkey.SetupGated,
		text.SetupGated,
		winrm.SetupGated,
		ami.SetupGated,
		awscodedeploy.SetupGated,
		awslambda.SetupGated,
		discoveryagent.SetupGated,
		discoverysetting.SetupGated,
		ecs.SetupGated,
		helmservice.SetupGated,
		kubernetesservice.SetupGated,
		ssh.SetupGated,
		tanzuservice.SetupGated,
		winrmservice.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
