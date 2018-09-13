package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// 使用编辑器编辑内容
func Editorbuffer() string {

	editor := os.Getenv("EDITOR")
	tmpDir := os.TempDir()
	tmpFile, tmpFileErr := ioutil.TempFile(tmpDir, "tempFilePrefix")
	if tmpFileErr != nil {
		fmt.Printf("Error %s while creating tempFile", tmpFileErr)
	}
	defer os.Remove(tmpFile.Name())

	path, err := exec.LookPath(editor)
	if err != nil {
		fmt.Printf("Error %s while looking up for %s!!", path, editor)
	}
	fmt.Printf("%s is available at %s\nCalling it with file %s \n", editor, path, tmpFile.Name())

	cmd := exec.Command(path, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Start failed: %s", err)
	}
	fmt.Printf("Waiting for command to finish.\n")
	err = cmd.Wait()
	fmt.Printf("Command finished with error: %v\n", err)

	dat, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		fmt.Printf("Error %s while open tempFile", err)
	}
	return string(dat)
}

// 创建Issue
// POST /repos/:owner/:repo/issues, 需要登录
// Issues内容通过编辑器输入,输出到临时文件，然后读取临时文件的内容
// 输入 owner repo title
func CreateIssues(owner, repo, title string) (*Issue, error) {

	CIssuesURL := "https://api.github.com/repos/" + owner + "/" + repo + "/issues"
	issueBody := Editorbuffer()

	values := map[string]string{"title": title, "body": issueBody}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", CIssuesURL, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &result, nil
}
