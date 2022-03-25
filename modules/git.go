package modules

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type gitModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

const (
	gitStashedIcon   = "S"
	gitUntrackedIcon = "?"
	gitModifiedIcon  = "M"
	gitStagedIcon    = "A"
	gitConflictsIcon = "!"
)

const maxGitRecursion = 10

func init() {
	register("git", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &gitModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label":          "\uE725 %s",
		"icon_stashed":   gitStashedIcon,
		"icon_untracked": gitUntrackedIcon,
		"icon_modified":  gitModifiedIcon,
		"icon_staged":    gitStagedIcon,
		"icon_conflicts": gitConflictsIcon,
	})
}

func (e *gitModule) Render(w *powerline.Writer) bool {
	baseStyle := e.mc.Style(e.gc)
	path, err := e.findGitPath(e.state.WorkingDir, 0)
	if err != nil {
		return false
	}

	output, clean := e.gitInfo(path)

	if clean {
		baseStyle.Foreground = e.mc.Fg("clean_fg", baseStyle.Foreground.String())
		baseStyle.Background = e.mc.Bg("clean_bg", baseStyle.Background.String())
	} else {
		baseStyle.Foreground = e.mc.Fg("dirty_fg", baseStyle.Foreground.String())
		baseStyle.Background = e.mc.Bg("dirty_bg", baseStyle.Background.String())
	}

	w.Printf(
		baseStyle,
		"%s",
		output,
	)
	return false
}

func (e *gitModule) findGitPath(start string, count int) (string, error) {
	if info, err := os.Stat(filepath.Join(start, ".git")); err == nil {
		if info.IsDir() {
			return start, nil
		}
	}
	above := filepath.Dir(start)
	if above == start || count >= maxGitRecursion {
		return "", fmt.Errorf("no .git directory found")
	}
	return e.findGitPath(above, count+1)
}

// see https://git-scm.com/docs/git-status
func (e *gitModule) gitInfo(path string) (string, bool) {

	head, err := os.ReadFile(filepath.Join(path, ".git", "HEAD"))
	if err != nil {
		return "", true
	}
	var branch string
	for _, line := range strings.Split(string(head), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "ref: ") {
			branch = strings.TrimPrefix(strings.TrimSpace(line[5:]), "refs/heads/")
			break
		}
		if len(line) > 8 {
			branch = line[:8]
		}
	}

	var stashed, untracked, modified, staged, conflicts int
	var ahead, behind int

	stashPath := filepath.Join(path, ".git", "refs", "stash")
	if stashData, err := os.ReadFile(stashPath); err == nil {
		stashed = len(strings.Split(strings.TrimSpace(string(stashData)), "\n"))
	}

	cmd := exec.Command("git", "status", "-s", "-b", "--ignore-submodules")
	cmd.Dir = path
	if output, err := cmd.Output(); err == nil {
		for _, line := range strings.Split(string(output), "\n") {
			if len(line) < 2 {
				continue
			}
			switch line[:2] {
			case "##":
				// TODO: parse ahead/behind
				parts := strings.Split(line, "[")
				if len(parts) > 1 {
					distance := strings.Split(parts[1], "]")[0]
					parts = strings.Split(distance, ",")
					for _, part := range parts {
						bits := strings.Split(strings.TrimSpace(part), " ")
						if len(bits) == 2 {

							count, err := strconv.Atoi(bits[1])
							if err != nil {
								continue
							}
							switch bits[0] {
							case "ahead":
								ahead = count
							case "behind":
								behind = count
							}
						}
					}
				}
			case "UU", "DD", "AU", "UD", "UA", "DU", "AA":
				conflicts++
			case "??":
				untracked++
			default:
				switch line[1] {
				case 'M', 'T', 'D', 'R', 'C':
					modified++
				case ' ':
					staged++
				}
			}
		}
	}

	var icons []string
	if stashed > 0 {
		icons = append(icons, e.mc.String("icon_stashed", gitStashedIcon))
	}
	if untracked > 0 {
		icons = append(icons, e.mc.String("icon_untracked", gitUntrackedIcon))
	}
	if modified > 0 {
		icons = append(icons, e.mc.String("icon_modified", gitModifiedIcon))
	}
	if staged > 0 {
		icons = append(icons, e.mc.String("icon_staged", gitStagedIcon))
	}
	if conflicts > 0 {
		icons = append(icons, e.mc.String("icon_conflicts", gitConflictsIcon))
	}

	output := branch
	if len(icons) > 0 {
		output += " " + strings.Join(icons, " ")
	}
	switch {
	case ahead > 0 && behind > 0:
		output = fmt.Sprintf("%s +%d/-%d", output, ahead, behind)
	case ahead > 0 && behind == 0:
		output = fmt.Sprintf("%s +%d", output, ahead)
	case ahead == 0 && behind > 0:
		output = fmt.Sprintf("%s -%d", output, behind)
	}

	clean := untracked+modified+staged+conflicts == 0
	return output, clean
}
