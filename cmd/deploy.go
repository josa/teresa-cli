package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/fatih/color"
	"github.com/luizalabs/teresa-cli/tar"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

type TeresaYaml struct {
	Hooks *struct {
		Pre  []string `yaml:"pre,omitempty"`
		Post []string `yaml:"post,omitempty"`
	} `yaml:"hooks,omitempty"`
}

var deployCmd = &cobra.Command{
	Use:   "deploy <app folder>",
	Short: "Deploy an app",
	// 	Long: `Deploy an application.
	//
	// To deploy an app you have to pass it's name, the team the app
	// belongs and the path to the source code. You might want to
	// describe your deployments through --description, as that'll
	// eventually help on rollbacks.
	//
	// eg.:
	//
	//   $ teresa deploy . --app webapi --team site --description "release 1.2 with new checkout"
	// 	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cluster, err := getCurrentClusterName()
		if err != nil {
			return newCmdError("You have to select a cluster first, check the config help: teresa config")
		}
		if len(args) != 1 {
			return newUsageError("You should provide the app folder in order to continue")
		}
		appFolder := args[0]
		appName, _ := cmd.Flags().GetString("app")
		deployDescription, _ := cmd.Flags().GetString("description")
		// showing warning message to the user
		fmt.Printf("Deploying app %s to the cluster %s...\n", color.CyanString(`"%s"`, appName), color.YellowString(`"%s"`, cluster))
		noinput, _ := cmd.Flags().GetBool("no-input")
		if !noinput {
			fmt.Print("Are you sure? (yes/NO)? ")
			// Waiting for the user answer...
			s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			if s = strings.ToLower(strings.TrimRight(s, "\r\n")); s != "yes" {
				return nil
			}
		}

		// create and get the archive
		fmt.Println("Generating tarball of:", appFolder)
		tarPath, err := createTempArchiveToUpload(appName, appFolder)
		if err != nil {
			// TODO: check what happen when the app folder is not valid
			return err
		}
		file, err := os.Open(*tarPath)
		if err != nil {
			return err
		}
		defer file.Close()

		tc := NewTeresa()
		err = tc.CreateDeploy(appName, deployDescription, file, os.Stdout)
		if err != nil {
			return err
		}
		return nil
	},
}

// // Writer to be used on deployment, as Write() is very specific and
// // should be implemented some other way -- moving out the deployment
// // error checking from it's Write method.
// type deploymentWriter struct {
// 	w io.Writer
// }
//
// // Write the buffer out to logger, return an error when the string
// // `----------deployment-error----------` is found on the buffer.
// func (tw *deploymentWriter) Write(p []byte) (n int, err error) {
// 	s := strings.Replace(string(p), deploymentErrorMark, "", -1)
// 	s = strings.Replace(s, deploymentSuccessMark, "", -1)
// 	// log.Info(strings.Trim(fmt.Sprintf("%s", s), "\n"))
// 	if strings.Contains(string(p), deploymentErrorMark) {
// 		return len(p), errors.New("Deploy failed")
// 	}
// 	return len(p), nil
// }

// create a temporary archive file of the app to deploy and return the path of this file
func createTempArchiveToUpload(appName, source string) (path *string, err error) {
	id := uuid.NewV4()
	source, err = filepath.Abs(source)
	if err != nil {
		return nil, err
	}
	p := filepath.Join(os.TempDir(), fmt.Sprintf("%s_%s.tar.gz", appName, id))
	if err = createArchive(source, p); err != nil {
		return nil, err
	}
	return &p, nil
}

// create an archive of the source folder
func createArchive(source, target string) error {
	dir, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("Dir not found to create an archive. %s", err)
	} else if !dir.IsDir() {
		return errors.New("Path to create the app archive isn't a directory")
	}

	hookFiles, err := createHookFiles(source)
	if err != nil {
		return fmt.Errorf("Can't process hooks of teresa.yaml file. %s", err)
	}
	defer cleanHookFiles(hookFiles)

	ignorePatterns, err := getIgnorePatterns(source)
	if err != nil {
		return errors.New("Invalid file '.teresaignore'")
	}

	t, err := tar.New(target)
	if err != nil {
		return err
	}
	defer t.Close()

	if ignorePatterns != nil {
		if err = addFiles(source, t, ignorePatterns); err != nil {
			return err
		}
	} else {
		t.AddAll(source)
	}

	return nil
}

func getIgnorePatterns(source string) ([]string, error) {
	fPath := filepath.Join(source, ".teresaignore")
	if _, err := os.Stat(fPath); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	file, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	patterns := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if text := scanner.Text(); text != "" {
			patterns = append(patterns, text)
		}
	}

	if len(patterns) == 0 {
		return nil, nil
	}

	return patterns, nil
}

// get teresa.yaml configuration
func getAppConfig(source string) (*TeresaYaml, error) {
	teresaYamlPath := filepath.Join(source, "teresa.yaml")
	if _, err := os.Stat(teresaYamlPath); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile(teresaYamlPath)
	if err != nil {
		return nil, err
	}

	teresaYaml := new(TeresaYaml)
	if err = yaml.Unmarshal(data, teresaYaml); err != nil {
		return nil, err
	}

	return teresaYaml, nil
}

func addFiles(source string, tar tar.Writer, ignorePatterns []string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, ip := range ignorePatterns {
			if matched, _ := filepath.Match(ip, info.Name()); matched {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if info.IsDir() {
			return nil
		}

		filename := strings.Replace(path, fmt.Sprintf("%s/", source), "", 1)
		return tar.AddFile(path, filename)
	})
}

// create hook files
func createHookFiles(source string) (hookFiles []string, err error) {
	hookFiles = make([]string, 0, 3)
	hooksPath := filepath.Join(source, "bin")

	teresaYaml, err := getAppConfig(source)
	if err != nil {
		return
	}
	if teresaYaml == nil || teresaYaml.Hooks == nil {
		return
	}

	if _, err = os.Stat(hooksPath); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(hooksPath, 0755); err != nil {
				return
			}
			hookFiles = append(hookFiles, hooksPath)
		}
	}
	if teresaYaml.Hooks.Pre != nil && len(teresaYaml.Hooks.Pre) > 0 {
		preCompileHook := filepath.Join(hooksPath, "pre-compile")
		if err = writeHookFile(preCompileHook, teresaYaml.Hooks.Pre); err != nil {
			return
		}
		hookFiles = append(hookFiles, preCompileHook)
	}
	if teresaYaml.Hooks.Post != nil && len(teresaYaml.Hooks.Post) > 0 {
		postCompileHook := filepath.Join(hooksPath, "post-compile")
		if err = writeHookFile(postCompileHook, teresaYaml.Hooks.Post); err != nil {
			return
		}
		hookFiles = append(hookFiles, postCompileHook)
	}
	return
}

func cleanHookFiles(hookFiles []string) {
	for i := len(hookFiles) - 1; i >= 0; i-- {
		if f, _ := os.Stat(hookFiles[i]); f.IsDir() {
			os.RemoveAll(hookFiles[i])
		} else {
			os.Remove(hookFiles[i])
		}
	}
}

func writeHookFile(fileName string, hooks []string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, s := range hooks {
		if _, err := f.WriteString(fmt.Sprintf("%s\n", s)); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	RootCmd.AddCommand(deployCmd)

	deployCmd.Flags().String("app", "", "app name (required)")
	deployCmd.Flags().String("description", "", "deploy description (required)")
	deployCmd.Flags().Bool("no-input", false, "deploy app without warning")

}
