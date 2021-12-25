/*
Copyright © 2021 luckyhappycat Beijing China <luckyhappycat@126.com>

*/
package linux

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/luckyhappycat/vscode-configer/pkg/common"
	"github.com/luckyhappycat/vscode-configer/pkg/cpp"
	"github.com/pkg/errors"
)

func NewMacOS() *mac {
	return &mac{}
}

type mac struct {
}

func (c *mac) Create() error {
	log.Print("pkg/cpp/cpp.create()")
	if err := c.createCCppProperties(); err != nil {
		return errors.WithStack(err)
	}
	if err := c.createLaunch(); err != nil {
		return errors.WithStack(err)
	}
	if err := c.createSettings(); err != nil {
		return errors.WithStack(err)
	}
	if err := c.createTasks(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *mac) createCCppProperties() error {
	log.Print("pkg/cpp/cpp.createCCppProperties()")
	paths := []string{"${workspaceFolder}/**"}
	includePath := "/usr/include"
	if common.IsDir(includePath) {
		paths = append(paths, includePath)
	}
	includePath = "/usr/local/include"
	if common.IsDir(includePath) {
		paths = append(paths, includePath)
	}
	includePath = "/Library/Developer/CommandLineTools/usr/include/"
	if common.IsDir(includePath) {
		paths = append(paths, includePath)
	}
	includePath = "/Library/Developer/CommandLineTools/usr/include/c++/v1"
	if common.IsDir(includePath) {
		paths = append(paths, includePath)
	}
	includePath = "/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include"
	if common.IsDir(includePath) {
		paths = append(paths, includePath)
	}
	log.Print("/Library/Developer/CommandLineTools/usr/lib/clang/n")
	versions, err := common.GetDirs("/Library/Developer/CommandLineTools/usr/lib/clang/")
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(versions)
	if err == nil {
		if len(versions) != 0 {
			includePath = fmt.Sprintf("/Library/Developer/CommandLineTools/usr/lib/clang/%s/include", versions[len(versions)-1])
			if common.IsDir(includePath) {
				paths = append(paths, includePath)
			} else {
				log.Printf("%s is not dir\n", includePath)
			}
		}
	}
	log.Print("/usr/include/c++/")
	versions, err = common.GetDirs("/usr/include/c++/")
	if err == nil {
		if len(versions) != 0 {
			includePath = fmt.Sprintf("/usr/include/c++/%s/", versions[len(versions)-1])
			if common.IsDir(includePath) {
				paths = append(paths, includePath)
			}
			includePath = fmt.Sprintf("/usr/include/c++/%s/backward", versions[len(versions)-1])
			if common.IsDir(includePath) {
				paths = append(paths, includePath)
			}
			includePath = fmt.Sprintf("/usr/include/c++/%s/x86_64-redhat-linux", versions[len(versions)-1])
			if common.IsDir(includePath) {
				paths = append(paths, includePath)
			}
		}
	}
	ccpp := cpp.CCppProperties{
		Configurations: []cpp.CCppPropertiesConfiguration{
			{
				Name:             "Linux",
				IncludePath:      paths,
				Defines:          []string{},
				CompilerPath:     "/usr/bin/gcc",
				CStandard:        "c11",
				CppStandard:      "c++11",
				IntelliSenseMode: "gcc-x64",
			},
		},
		Version: 4,
	}
	indentCcpp, err := json.MarshalIndent(ccpp, "", "    ")
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.MkdirAll(path.Dir("./.vscode/c_cpp_properties.json"), os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create("./.vscode/c_cpp_properties.json")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	f.Write(indentCcpp)
	return nil
}

func (c *mac) createLaunch() error {
	log.Print("pkg/cpp/cpp.createLaunch()")
	cfg := cpp.Configuration{
		Name:            "g++调试",
		Type:            "cppdbg",
		Request:         "launch",
		Program:         "${fileDirname}/${fileBasenameNoExtension}",
		Args:            []string{},
		StopAtEntry:     false,
		Cwd:             "${workspaceFolder}",
		Environment:     []string{},
		ExternalConsole: false,
		MIMode:          "gdb",
		PreLaunchTask:   "C/C++: g++ build",
		PostDebugTask:   "C/C++: g++ clean",
	}
	launch := cpp.Launch{
		Version:        "0.2.0",
		Configurations: []cpp.Configuration{cfg},
	}
	indentLaunch, err := json.MarshalIndent(launch, "", "    ")
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.MkdirAll(path.Dir("./.vscode/launch.json"), os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create("./.vscode/launch.json")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	f.Write(indentLaunch)

	return nil
}

func (c *mac) createSettings() error {
	log.Print("pkg/cpp/cpp.createSettings()")
	fileAssociations := cpp.FilesAssociations{
		Tcc:             "cpp",
		Streambuf:       "cpp",
		String:          "cpp",
		Array:           "cpp",
		Cctype:          "cpp",
		Clocale:         "cpp",
		Cmath:           "cpp",
		Cstdarg:         "cpp",
		Cstdint:         "cpp",
		Cstdio:          "cpp",
		Cstdlib:         "cpp",
		Cstring:         "cpp",
		Ctime:           "cpp",
		Cwchar:          "cpp",
		Deque:           "cpp",
		List:            "cpp",
		UnorderedMap:    "cpp",
		Vector:          "cpp",
		Exception:       "cpp",
		Fstream:         "cpp",
		Functional:      "cpp",
		InitializerList: "cpp",
		Iosfwd:          "cpp",
		Iostream:        "cpp",
		Istream:         "cpp",
		Limits:          "cpp",
		New:             "cpp",
		Ostream:         "cpp",
		Numeric:         "cpp",
		Sstream:         "cpp",
		Stdexcept:       "cpp",
		Cinttypes:       "cpp",
		Tuple:           "cpp",
		TypeTraits:      "cpp",
		Utility:         "cpp",
		Typeinfo:        "cpp",
	}
	settings := cpp.Settings{
		CmakeConfigureOnOpen:        false,
		FilesAssociations:           fileAssociations,
		CCppClangFormatSortIncludes: false,
	}
	indentSettings, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.MkdirAll(path.Dir("./.vscode/settings.json"), os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create("./.vscode/settings.json")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	f.Write(indentSettings)
	return nil
}

func (c *mac) createTasks() error {
	log.Print("pkg/cpp/cpp.createTasks()")
	gppBuild := cpp.Task{
		Type:    "shell",
		Label:   "C/C++: g++ build",
		Command: "/usr/bin/g++",
		Args: []string{"-std=c++11",
			"-g",
			"${file}",
			"-o",
			"${fileDirname}/${fileBasenameNoExtension}",
			"&&",
			"clear",
			"&&",
			"${fileDirname}/${fileBasenameNoExtension}",
		},
		Options: cpp.Options{
			Cwd: "${workspaceFolder}",
		},
		ProblemMatcher: []string{"$gcc"},
		Group: cpp.Group{
			Kind:      "build",
			IsDefault: true,
		},
		Presentation: cpp.Presentation{
			Echo:             false,
			Reveal:           "always",
			Focus:            false,
			Panel:            "shared",
			Clear:            true,
			ShowReuseMessage: false,
		},
	}

	gppClean := cpp.Task{
		Type:    "shell",
		Label:   "C/C++: g++ clean",
		Command: "rm",
		Args: []string{"-rf",
			"${fileDirname}/${fileBasenameNoExtension}.dSYM",
			"&&",
			"rm",
			"-rf",
			"${fileDirname}/${fileBasenameNoExtension}",
		},
		Options: cpp.Options{
			Cwd: "${workspaceFolder}",
		},
		Group: cpp.Group{
			Kind:      "build",
			IsDefault: true,
		},
		Presentation: cpp.Presentation{
			Echo:             false,
			Reveal:           "always",
			Focus:            false,
			Panel:            "shared",
			ShowReuseMessage: false,
		},
	}

	cmakeBuild := cpp.Task{
		Type:    "shell",
		Label:   "C/C++: cmake build",
		Command: "cd ./build; cmake ..; make; ",
		Options: cpp.Options{
			Cwd: "${workspaceFolder}",
		},
		Group: cpp.Group{
			Kind:      "build",
			IsDefault: true,
		},
		Presentation: cpp.Presentation{
			Echo:             false,
			Reveal:           "always",
			Focus:            false,
			Panel:            "shared",
			ShowReuseMessage: false,
		},
	}

	cmakeClean := cpp.Task{
		Type:    "shell",
		Label:   "C/C++: cmake clean",
		Command: "${fileDirname}/bin/${fileBasenameNoExtension}",
		Args: []string{
			"&&",
			"rm",
			"-rf",
			"./build/*",
		},
		Options: cpp.Options{
			Cwd: "${workspaceFolder}",
		},
		Group: cpp.Group{
			Kind:      "build",
			IsDefault: true,
		},
		Presentation: cpp.Presentation{
			Echo:             false,
			Reveal:           "always",
			Focus:            false,
			Panel:            "shared",
			ShowReuseMessage: false,
		},
	}

	tasks := cpp.Tasks{
		Tasks:   []cpp.Task{gppBuild, gppClean, cmakeBuild, cmakeClean},
		Version: "2.0.0",
	}
	indentTasks, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return errors.WithStack(err)
	}
	indentTasks = bytes.Replace(indentTasks, []byte("\\u0026"), []byte("&"), -1)
	err = os.MkdirAll(path.Dir("./.vscode/tasks.json"), os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create("./.vscode/tasks.json")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	f.Write(indentTasks)
	return nil
}