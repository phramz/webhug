package security

import "testing"

func TestGithubSign(t *testing.T) {
	var tests = []struct {
		signature string
		secret    string
		payload   string
	}{
		{
			"sha1=21e3409b544189b1f12a0fe9ba8270685e093edb",
			"231a9f49a1592ec348663c37e53885164fe840a0b5e464e017c2ed6f06686",
			"{\"zen\":\"Design for failure.\",\"hook_id\":178086120,\"hook\":{\"type\":\"Repository\",\"id\":178086120,\"name\":\"web\",\"active\":true,\"events\":[\"*\"],\"config\":{\"content_type\":\"json\",\"insecure_ssl\":\"0\",\"secret\":\"********\",\"url\":\"http://node.krautdrive.com:20000/github\"},\"updated_at\":\"2020-01-26T15:30:47Z\",\"created_at\":\"2020-01-26T15:30:47Z\",\"url\":\"https://api.github.com/repos/phramz/kd-worker-api/hooks/178086120\",\"test_url\":\"https://api.github.com/repos/phramz/kd-worker-api/hooks/178086120/test\",\"ping_url\":\"https://api.github.com/repos/phramz/kd-worker-api/hooks/178086120/pings\",\"last_response\":{\"code\":null,\"status\":\"unused\",\"message\":null}},\"repository\":{\"id\":224494642,\"node_id\":\"MDEwOlJlcG9zaXRvcnkyMjQ0OTQ2NDI=\",\"name\":\"kd-worker-api\",\"full_name\":\"phramz/kd-worker-api\",\"private\":true,\"owner\":{\"login\":\"phramz\",\"id\":2319696,\"node_id\":\"MDQ6VXNlcjIzMTk2OTY=\",\"avatar_url\":\"https://avatars3.githubusercontent.com/u/2319696?v=4\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/phramz\",\"html_url\":\"https://github.com/phramz\",\"followers_url\":\"https://api.github.com/users/phramz/followers\",\"following_url\":\"https://api.github.com/users/phramz/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/phramz/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/phramz/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/phramz/subscriptions\",\"organizations_url\":\"https://api.github.com/users/phramz/orgs\",\"repos_url\":\"https://api.github.com/users/phramz/repos\",\"events_url\":\"https://api.github.com/users/phramz/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/phramz/received_events\",\"type\":\"User\",\"site_admin\":false},\"html_url\":\"https://github.com/phramz/kd-worker-api\",\"description\":null,\"fork\":false,\"url\":\"https://api.github.com/repos/phramz/kd-worker-api\",\"forks_url\":\"https://api.github.com/repos/phramz/kd-worker-api/forks\",\"keys_url\":\"https://api.github.com/repos/phramz/kd-worker-api/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/phramz/kd-worker-api/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/phramz/kd-worker-api/teams\",\"hooks_url\":\"https://api.github.com/repos/phramz/kd-worker-api/hooks\",\"issue_events_url\":\"https://api.github.com/repos/phramz/kd-worker-api/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/phramz/kd-worker-api/events\",\"assignees_url\":\"https://api.github.com/repos/phramz/kd-worker-api/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/phramz/kd-worker-api/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/phramz/kd-worker-api/tags\",\"blobs_url\":\"https://api.github.com/repos/phramz/kd-worker-api/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/phramz/kd-worker-api/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/phramz/kd-worker-api/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/phramz/kd-worker-api/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/phramz/kd-worker-api/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/phramz/kd-worker-api/languages\",\"stargazers_url\":\"https://api.github.com/repos/phramz/kd-worker-api/stargazers\",\"contributors_url\":\"https://api.github.com/repos/phramz/kd-worker-api/contributors\",\"subscribers_url\":\"https://api.github.com/repos/phramz/kd-worker-api/subscribers\",\"subscription_url\":\"https://api.github.com/repos/phramz/kd-worker-api/subscription\",\"commits_url\":\"https://api.github.com/repos/phramz/kd-worker-api/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/phramz/kd-worker-api/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/phramz/kd-worker-api/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/phramz/kd-worker-api/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/phramz/kd-worker-api/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/phramz/kd-worker-api/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/phramz/kd-worker-api/merges\",\"archive_url\":\"https://api.github.com/repos/phramz/kd-worker-api/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/phramz/kd-worker-api/downloads\",\"issues_url\":\"https://api.github.com/repos/phramz/kd-worker-api/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/phramz/kd-worker-api/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/phramz/kd-worker-api/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/phramz/kd-worker-api/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/phramz/kd-worker-api/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/phramz/kd-worker-api/releases{/id}\",\"deployments_url\":\"https://api.github.com/repos/phramz/kd-worker-api/deployments\",\"created_at\":\"2019-11-27T18:37:46Z\",\"updated_at\":\"2020-01-26T13:39:25Z\",\"pushed_at\":\"2020-01-26T13:39:23Z\",\"git_url\":\"git://github.com/phramz/kd-worker-api.git\",\"ssh_url\":\"git@github.com:phramz/kd-worker-api.git\",\"clone_url\":\"https://github.com/phramz/kd-worker-api.git\",\"svn_url\":\"https://github.com/phramz/kd-worker-api\",\"homepage\":null,\"size\":8644,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"HTML\",\"has_issues\":true,\"has_projects\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":0,\"mirror_url\":null,\"archived\":false,\"disabled\":false,\"open_issues_count\":0,\"license\":null,\"forks\":0,\"open_issues\":0,\"watchers\":0,\"default_branch\":\"develop\"},\"sender\":{\"login\":\"phramz\",\"id\":2319696,\"node_id\":\"MDQ6VXNlcjIzMTk2OTY=\",\"avatar_url\":\"https://avatars3.githubusercontent.com/u/2319696?v=4\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/phramz\",\"html_url\":\"https://github.com/phramz\",\"followers_url\":\"https://api.github.com/users/phramz/followers\",\"following_url\":\"https://api.github.com/users/phramz/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/phramz/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/phramz/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/phramz/subscriptions\",\"organizations_url\":\"https://api.github.com/users/phramz/orgs\",\"repos_url\":\"https://api.github.com/users/phramz/repos\",\"events_url\":\"https://api.github.com/users/phramz/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/phramz/received_events\",\"type\":\"User\",\"site_admin\":false}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.signature, func(t *testing.T) {
			s := githubSign([]byte(tt.payload), []byte(tt.secret))
			if s != tt.signature {
				t.Errorf("got %q, want %q", s, tt.signature)
			}
		})
	}

}