package cmd

import (
	"fmt"
	"github.com/xzly21k/d_cli/ask"
	"github.com/xzly21k/d_cli/constants"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

const (
	gitRepoURL = "https://github.com/xzly21k/d_cli.git"
	RepoUrl    = "github.com/xzly21k/d_cli"
)

func getLatestTag() (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-Command",
			"git ls-remote --tags --sort='v:refname' --refs "+gitRepoURL+" | Select-Object -Last 1")
	} else {
		cmd = exec.Command("bash", "-c",
			"git ls-remote --tags --sort='v:refname' --refs "+gitRepoURL+" | tail -n 1")
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	var latestTag string
	if runtime.GOOS == "windows" {
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		latestTagLine := lines[len(lines)-1]
		latestTag = strings.Split(latestTagLine, "/")[2]
	} else {
		latestTagLine := strings.TrimSpace(string(output))
		latestTag = strings.Split(latestTagLine, "/")[2]
	}
	latestTag = strings.ReplaceAll(latestTag, "v", "")
	return latestTag, nil
}

func installLatestVersion(repo string) error {
	var cmd *exec.Cmd
	modulePath := fmt.Sprintf("%s@%s", repo, "latest")
	cmd = exec.Command("go", "install", modulePath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("执行命令时发生错误：%v", err)
	}
	return nil
}

func UpdateLatestVersion() (isDone bool) {
	isDone = false
	var (
		latestVersion string
		err           error
	)
	if latestVersion, err = getLatestTag(); err != nil {
		log.Printf("[获取最新版本失败]:" + err.Error())
		return
	}
	if constants.Version != latestVersion {
		log.Printf((fmt.Sprintf("[目前的版本]version:%s", constants.Version)))
		log.Println(fmt.Sprintf("[发现新版本]version:%s", latestVersion))
		if ok, _ := ask.ConfirmYes("是否需要更新版本"); !ok {
			return
		}
		if err := installLatestVersion(RepoUrl); err != nil {
			log.Printf("[安装最新版本失败]:" + err.Error())
			return
		}
		log.Printf("安装最新版本成功,请重新执行命令")
		isDone = true
	}
	return
}
