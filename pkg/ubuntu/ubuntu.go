package ubuntu

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsUbuntu() bool {

    fmt.Println("Installing setup of the neovim, wait...")

	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
    return strings.TrimSpace(string(out)) == "ubuntu"
}

func UbuntuCommands() []string {
	return []string{
		"apt update",
        "wget https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz",
        "tar -xnzf nvim-linux64.tar.gz",
        "cp nvim-linux64/bin/nvim /usr/bin/",
        "rm -rf nvim-linux64",
        "rm nvim-linux64.tar.gz",
        "apt install -y curl npm python3-pip nodejs git",
		"pip install --upgrade pip",
		"pip install wheel --break-system-packages",
		"pip install pynvim neovim --break-system-packages",
		"mkdir -p $HOME/.config/nvim",
		"curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
		"rm -rf $HOME/nvim",
		"git clone -v https://github.com/whyakari/nvim $HOME/nvim",
		"cp $HOME/nvim/init.vim $HOME/.config/nvim/",
        "rm README.md",
		"rm -rf $HOME/nvim",
	}
}
