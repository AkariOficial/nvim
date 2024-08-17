package termux

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsTermux() bool {

    fmt.Println("Installing setup of the neovim, wait...")

	out, err := exec.Command("bash", "-c", "uname -o").Output()
	if err != nil {
		return false
	}
    return strings.TrimSpace(string(out)) == "Android"
}

func TermuxCommands() []string {
	return []string{
		"apt update",
		"pkg in libllvm lua53 clang rust nodejs-lts git python neovim -y",
		"pip install --upgrade python-pip",
		"pip install wheel",
		"pkg in python-pynvim -y",
		"mkdir -p $HOME/.config/nvim",
		"curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
		"git clone -v https://github.com/whyakari/nvim $HOME/nvim_tmp",
		"cp $HOME/nvim_tmp/init.lua $HOME/.config/nvim/",
        "cp -r $HOME/nvim_tmp/lua $HOME/.config/nvim/",
        "cp $HOME/nvim_tmp/lazy-lock.json $HOME/.config/nvim/",
        "cd $HOME",
		"rm -rf nvim",
		"rm -rf $HOME/nvim_tmp",
		"rm $HOME/README.md",
        "rm $HOME/*.tar.gz",
	}
}

