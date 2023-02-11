package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type Response struct {
	ID          int    `json:"id"`
	WebURL      string `json:"web_url"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Projects    []struct {
		ID                                        int           `json:"id"`
		Path                                      string        `json:"path"`
		PathWithNamespace                         string        `json:"path_with_namespace"`
		DefaultBranch                             string        `json:"default_branch"`
		TagList                                   []interface{} `json:"tag_list"`
		Topics                                    []interface{} `json:"topics"`
		SSHURLToRepo                              string        `json:"ssh_url_to_repo"`
		HTTPURLToRepo                             string        `json:"http_url_to_repo"`
		WebURL                                    string        `json:"web_url"`
		ReadmeURL                                 string        `json:"readme_url"`
		AvatarURL                                 interface{}   `json:"avatar_url"`
		ForksCount                                int           `json:"forks_count"`
		StarCount                                 int           `json:"star_count"`
		LastActivityAt                            time.Time     `json:"last_activity_at"`
		PackagesEnabled                           bool          `json:"packages_enabled"`
		EmptyRepo                                 bool          `json:"empty_repo"`
		Archived                                  bool          `json:"archived"`
		Visibility                                string        `json:"visibility"`
		ResolveOutdatedDiffDiscussions            bool          `json:"resolve_outdated_diff_discussions"`
		IssuesEnabled                             bool          `json:"issues_enabled"`
		MergeRequestsEnabled                      bool          `json:"merge_requests_enabled"`
		WikiEnabled                               bool          `json:"wiki_enabled"`
		JobsEnabled                               bool          `json:"jobs_enabled"`
		SnippetsEnabled                           bool          `json:"snippets_enabled"`
		ContainerRegistryEnabled                  bool          `json:"container_registry_enabled"`
		ServiceDeskEnabled                        bool          `json:"service_desk_enabled"`
		ServiceDeskAddress                        interface{}   `json:"service_desk_address"`
		CanCreateMergeRequestIn                   bool          `json:"can_create_merge_request_in"`
		IssuesAccessLevel                         string        `json:"issues_access_level"`
		RepositoryAccessLevel                     string        `json:"repository_access_level"`
		MergeRequestsAccessLevel                  string        `json:"merge_requests_access_level"`
		ForkingAccessLevel                        string        `json:"forking_access_level"`
		WikiAccessLevel                           string        `json:"wiki_access_level"`
		BuildsAccessLevel                         string        `json:"builds_access_level"`
		SnippetsAccessLevel                       string        `json:"snippets_access_level"`
		PagesAccessLevel                          string        `json:"pages_access_level"`
		OperationsAccessLevel                     string        `json:"operations_access_level"`
		AnalyticsAccessLevel                      string        `json:"analytics_access_level"`
		ContainerRegistryAccessLevel              string        `json:"container_registry_access_level"`
		SecurityAndComplianceAccessLevel          string        `json:"security_and_compliance_access_level"`
		ReleasesAccessLevel                       string        `json:"releases_access_level"`
		EnvironmentsAccessLevel                   string        `json:"environments_access_level"`
		FeatureFlagsAccessLevel                   string        `json:"feature_flags_access_level"`
		InfrastructureAccessLevel                 string        `json:"infrastructure_access_level"`
		MonitorAccessLevel                        string        `json:"monitor_access_level"`
		EmailsDisabled                            interface{}   `json:"emails_disabled"`
		SharedRunnersEnabled                      bool          `json:"shared_runners_enabled"`
		LfsEnabled                                bool          `json:"lfs_enabled"`
		CreatorID                                 int           `json:"creator_id"`
		ImportURL                                 interface{}   `json:"import_url"`
		ImportType                                interface{}   `json:"import_type"`
		ImportStatus                              string        `json:"import_status"`
		OpenIssuesCount                           int           `json:"open_issues_count"`
		CiDefaultGitDepth                         int           `json:"ci_default_git_depth"`
		CiForwardDeploymentEnabled                bool          `json:"ci_forward_deployment_enabled"`
		CiJobTokenScopeEnabled                    bool          `json:"ci_job_token_scope_enabled"`
		CiSeparatedCaches                         bool          `json:"ci_separated_caches"`
		CiOptInJwt                                bool          `json:"ci_opt_in_jwt"`
		CiAllowForkPipelinesToRunInParentProject  bool          `json:"ci_allow_fork_pipelines_to_run_in_parent_project"`
		PublicJobs                                bool          `json:"public_jobs"`
		BuildTimeout                              int           `json:"build_timeout"`
		AutoCancelPendingPipelines                string        `json:"auto_cancel_pending_pipelines"`
		CiConfigPath                              interface{}   `json:"ci_config_path"`
		SharedWithGroups                          []interface{} `json:"shared_with_groups"`
		OnlyAllowMergeIfPipelineSucceeds          bool          `json:"only_allow_merge_if_pipeline_succeeds"`
		AllowMergeOnSkippedPipeline               interface{}   `json:"allow_merge_on_skipped_pipeline"`
		RestrictUserDefinedVariables              bool          `json:"restrict_user_defined_variables"`
		RequestAccessEnabled                      bool          `json:"request_access_enabled"`
		OnlyAllowMergeIfAllDiscussionsAreResolved bool          `json:"only_allow_merge_if_all_discussions_are_resolved"`
		RemoveSourceBranchAfterMerge              bool          `json:"remove_source_branch_after_merge"`
		PrintingMergeRequestLinkEnabled           bool          `json:"printing_merge_request_link_enabled"`
		MergeMethod                               string        `json:"merge_method"`
		SquashOption                              string        `json:"squash_option"`
		EnforceAuthChecksOnUploads                bool          `json:"enforce_auth_checks_on_uploads"`
		SuggestionCommitMessage                   interface{}   `json:"suggestion_commit_message"`
		MergeCommitTemplate                       interface{}   `json:"merge_commit_template"`
		SquashCommitTemplate                      interface{}   `json:"squash_commit_template"`
		IssueBranchTemplate                       interface{}   `json:"issue_branch_template"`
		AutoDevopsEnabled                         bool          `json:"auto_devops_enabled"`
		AutoDevopsDeployStrategy                  string        `json:"auto_devops_deploy_strategy"`
		AutocloseReferencedIssues                 bool          `json:"autoclose_referenced_issues"`
		KeepLatestArtifact                        bool          `json:"keep_latest_artifact"`
		RunnerTokenExpirationInterval             interface{}   `json:"runner_token_expiration_interval"`
	} `json:"projects"`
	SharedProjects []interface{} `json:"shared_projects"`
}

type SubGroups []struct {
	ID                             int         `json:"id"`
	WebURL                         string      `json:"web_url"`
	Name                           string      `json:"name"`
	Path                           string      `json:"path"`
	Description                    string      `json:"description"`
	Visibility                     string      `json:"visibility"`
	ShareWithGroupLock             bool        `json:"share_with_group_lock"`
	RequireTwoFactorAuthentication bool        `json:"require_two_factor_authentication"`
	TwoFactorGracePeriod           int         `json:"two_factor_grace_period"`
	ProjectCreationLevel           string      `json:"project_creation_level"`
	AutoDevopsEnabled              interface{} `json:"auto_devops_enabled"`
	SubgroupCreationLevel          string      `json:"subgroup_creation_level"`
	EmailsDisabled                 interface{} `json:"emails_disabled"`
	MentionsDisabled               interface{} `json:"mentions_disabled"`
	LfsEnabled                     bool        `json:"lfs_enabled"`
	DefaultBranchProtection        int         `json:"default_branch_protection"`
	AvatarURL                      interface{} `json:"avatar_url"`
	RequestAccessEnabled           bool        `json:"request_access_enabled"`
	FullName                       string      `json:"full_name"`
	FullPath                       string      `json:"full_path"`
	CreatedAt                      time.Time   `json:"created_at"`
	ParentID                       int         `json:"parent_id"`
}

func clone(idGroup string, token string, root_dir string, privateKeyFile string, password string) {
	_, err := os.Stat(privateKeyFile)
	if err != nil {
		Warning("read file %s failed %s\n", privateKeyFile, err.Error())
		return
	}

	// Clone the given repository to the given directory
	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, password)
	if err != nil {
		Warning("generate publickeys failed: %s\n", err.Error())
		return
	}

	resp, err := http.Get("https://git.preprod.mangopay.com/api/v4/groups/" + idGroup + "?private_token=" + token + "&per_page=100000")
	if err != nil {
		fmt.Print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, s := range result.Projects {
		_, err := git.PlainClone(root_dir+s.Path, false, &git.CloneOptions{
			Auth:     publicKeys,
			URL:      s.SSHURLToRepo,
			Progress: os.Stdout,
		})
		fmt.Println("clone " + s.SSHURLToRepo)
		fmt.Println("in  " + root_dir + s.Path)
		if err != nil {
			fmt.Println("error " + s.SSHURLToRepo)
		}
	}

	// SUBGROUB
	resp_sub, err := http.Get("https://git.preprod.mangopay.com/api/v4/groups/" + idGroup + "/subgroups?private_token=" + token + "&per_page=100000")
	if err != nil {
		fmt.Print(err)
	}
	body_sub, err := ioutil.ReadAll(resp_sub.Body)

	var result_sub SubGroups
	if err := json.Unmarshal(body_sub, &result_sub); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, s := range result_sub {
		clone(strconv.Itoa(s.ID), token, root_dir+s.Path+string(os.PathSeparator), privateKeyFile, password)
	}
}

func main() {
	idRoot := os.Args[1]
	token := os.Args[2]
	root_dir := os.Args[3]
	privateKeyFile := os.Args[4]
	password := os.Args[5]
	clone(idRoot, token, root_dir+string(os.PathSeparator), privateKeyFile, password)
}
