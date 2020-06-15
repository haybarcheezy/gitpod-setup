package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "init" {
			if len(os.Args) >= 3 {
				switch os.Args[2] {
				case "julia":
					juliaInit()
					return
				case "nim":
					nimInit()
					return
				case "hy":
					hyInit()
					return
				case "clojure":
					clojureInit()
					return
				case "haskell":
					haskellInit()
					return
				case "dotnet":
					dotNetInit()
					return
				case "zsh":
					zshInit()
					return
				case "kotlin":
					kotlinInit()
					return
				default:
					fmt.Println("Invalid argument '" + strings.Join(os.Args[2:], " ") + "'")
					return
				}
			}
			initInteractive()
			return
		}
		fmt.Println("Help:\n    init: initialize gitpod")
		return
	}
	return
}

func initInteractive() {
	isError := false
START:
	if isError {
		fmt.Println("\033[1;31mError Reading Option\033[0m")
		isError = false
	}
	{
		startPrompt := promptui.Select{
			Label: "What are you configuring",
			Items: []string{"Language", "Shell", "Never Mind"},
		}
		_, result, err := startPrompt.Run()
		if err != nil {
			exit()
		}
		switch result {
		case "Never Mind":
			exit()
		case "Shell":
			goto SHELL
		case "Language":
			goto Language
		}
	}
SHELL:
	{
		shellPrompt := promptui.Select{
			Label: "What shell are you configuring",
			Items: []string{"ZSH", "Back", "Never Mind"},
		}
		_, result1, err1 := shellPrompt.Run()
		if err1 != nil {
			exit()
		}
		switch result1 {
		case "ZSH":
			zshInit()
			return
		case "Never Mind":
			exit()
		case "Back":
			goto START
		default:
			isError = true
			goto START
		}
	}
Language:
	{
		langPrompt := promptui.Select{
			Label: "What Language are you configuring",
			Items: []string{"Julia", "Nim", "Hy", "Clojure", "Haskell", ".NET", "Kotlin", "Back", "Never Mind"},
		}
		_, result2, err2 := langPrompt.Run()
		if err2 != nil {
			exit()
		}
		switch result2 {
		case "Julia":
			juliaInit()
		case "Nim":
			nimInit()
		case "Hy":
			hyInit()
		case "Clojure":
			clojureInit()
		case "Haskell":
			haskellInit()
		case ".NET":
			dotNetInit()
		case "Kotlin":
			kotlinInit()
		case "Never Mind":
			exit()
		case "Back":
			goto START
		default:
			isError = true
			goto START
		}
	}
}
func juliaInit() {
	initBase(juliaDockerFile, juliaYaml)
	fmt.Println("Julia Setup Complete!")
}
func nimInit() {
	initBase(nimDockerFile, nimYaml)
	fmt.Println("Nim Setup Complete!")
}
func hyInit() {
	initBase(hyDockerfile, hyYaml)
	fmt.Println("Hy Setup Complete!")
}
func clojureInit() {
	initBase(clojureDockerfile, clojureYaml)
	fmt.Println("Clojure Setup Complete!")
}
func haskellInit() {
	initBase(haskellDockerfile, haskellYaml)
	fmt.Println("Haskell Setup Complete!")
}
func dotNetInit() {
	initBase(dotNetDockerfile, dotNetYaml)
	fmt.Println(".NET Setup Complete!")
}
func zshInit() {
	initBase(zshDockerfile, zshYaml)
	fmt.Println("ZSH Setup Complete!")
}
func kotlinInit() {
	initBase(kotlinDockerfile, kotlinYaml)
	fmt.Println("Kotlin Setup Complete!")
}
func initBase(dockerFile, Yaml string) {
	gitpodDockerfile, _ := os.Create(".gitpod.Dockerfile")
	gitpodYaml, _ := os.Create(".gitpod.yml")
	gitpodDockerfile.WriteString(dockerFile)
	gitpodYaml.WriteString(Yaml)
}
func exit() {
	prompt := promptui.Prompt{
		Label:     "Are you sure",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return
	}
	if strings.ToLower(result) == "y" {
		os.Exit(0)
	} else {
		return
	}
	fmt.Printf("You choose %q\n", result)
}

var juliaDockerFile string = `FROM gitpod/workspace-full

USER gitpod

# Install Julia
RUN sudo apt-get update \
	&& sudo apt-get install -y \
		build-essential \
		libatomic1 \
		python \
		gfortran \
		perl \
		wget \
		m4 \
		cmake \
		pkg-config \
		julia \
	&& sudo rm -rf /var/lib/apt/lists/*

# Give control back to Gitpod Layer
USER root`

var juliaYaml string = `image:
  file: .gitpod.Dockerfile

tasks:
- command: julia --version

vscode:
  extensions:
    - julialang.language-julia@0.12.3:lgRyBd8rjwUpMGG0C5GAig==
`

var nimDockerFile string = `FROM gitpod/workspace-full

USER gitpod

RUN sudo apt-get update \
	&& sudo apt-get install -y \
		nim`

var nimYaml string = `image:
  file: .gitpod.Dockerfile

tasks:
  - command: nimc --version

vscode:
  extensions:
  	- kosz78.nim@0.6.3:w7n1wKOFVkz9yIqgRYT7lQ==`

var hyDockerfile string = `FROM gitpod/workspace-full

USER gitpod

RUN pip3 install hy --user
`

var hyYaml string = `image:
  file: .gitpod.Dockerfile

vscode:
  extensions:
  	  - xuqinghan.vscode-hy@0.0.4:Utf282betrZISZjOJLTZlg==
`
var clojureDockerfile string = `FROM gitpod/workspace-full

USER gitpod

# Install Clojure
RUN curl -O https://download.clojure.org/install/linux-install-1.10.1.492.sh \
    && chmod +x linux-install-1.10.1.492.sh  \
    && sudo ./linux-install-1.10.1.492.sh

# Give access back to gitpod image builder
USER root
`
var clojureYaml string = `image:
  file: .gitpod.Dockerfile

vscode:
  extensions:
  	  - avli.clojure@0.11.1:LAV1SbBlP0gU7J8kduhQvQ==
`
var haskellDockerfile string = `FROM gitpod/workspace-full

USER gitpod

# Installing Haskell
RUN sudo add-apt-repository -y ppa:hvr/ghc \
    && sudo apt-get update && \
	&& sudo apt-get install -y \
		cabal-install \
		ghc

# Give control back to gitpod layer
USER root
`
var haskellYaml string = `image:
  file: .gitpod.Dockerfile

vscode:
  extensions:
      - alanz.vscode-hie-server@0.0.28:j/YAJtXUGGbb8xSSz1i/CQ==
      - justusadam.language-haskell@2.6.0:CvYnp3YmQPTuto0m1di+1A==
      - phoityne.phoityne-vscode@0.0.24:FTkd1r93lYs3z95fjRROAg==
      - hoovercj.haskell-linter@0.0.6:VpJluXvOyr9Iw7TIKg2Oyg==
      - dramforever.vscode-ghc-simple@0.1.13:X3A6Dr3LYAP8MxXBh/hb1A==
`
var dotNetDockerfile string = `FROM gitpod/workspace-full
USER gitpod

# Install .NET
RUN wget -q https://packages.microsoft.com/config/ubuntu/16.04/packages-microsoft-prod.deb -O packages-microsoft-prod.deb \
	&& sudo dpkg -i packages-microsoft-prod.deb \
	&& sudo apt-get update \
	&& sudo apt-get install -y \
        fsharp \
		apt-transport-https \
		dotnet-sdk-3.1 \
        aspnetcore-runtime-3.1 \
        dotnet-runtime-3.1
`
var dotNetYaml string = `image:
  file: .gitpod.Dockerfile

vscode:
  extensions:
    - Ionide.Ionide-fsharp@4.1.0:vk6avJmuBqlMwZEelzdnZQ==
    - ms-vscode.csharp@1.21.4:lLV3lBwYKRTJ3QAQjtNMaQ==

`
var zshDockerfile string = `FROM gitpod/workspace-full

USER root

RUN apt-get update \
   && apt-get install -y zsh \
   && apt-get clean \
   && rm -rf /var/cache/apt/* \
   && rm -rf /var/lib/apt/lists/* \
   && rm -rf /tmp/*
`
var zshYaml string = `image:
  file: .gitpod.Dockerfile

tasks:
  - command: zsh
`
var kotlinDockerfile string = `FROM gitpod/workspace-full

USER gitpod

RUN brew install kotlin
`
var kotlinYaml string = `image:
  file: .gitpod.Dockerfile

vscode:
  extensions:
    - mathiasfrohlich.Kotlin@1.7.0:9xQZtwTUg4bdXHCMyxM7vQ==
    - fwcd.kotlin@0.2.11:moh8IDanzsIlhtK2IeiLmQ==`
