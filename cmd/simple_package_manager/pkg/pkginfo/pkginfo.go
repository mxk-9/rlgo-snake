package pkginfo

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	log "spm/pkg/slog"
)

type FetchPath struct {
	From string   `json:"from"`
	To   []string `json:"to"`
}

type UnpackTask struct {
	Type          string   `json:"type"`
	Src           []string `json:"src"`
	Dest          []string `json:"dest"`
	SelectedItems []string `json:"selected_items"` // Full path to needed object, extacting recursively
}

type InstallTaskCommand string

type InstallTask struct {
	Command InstallTaskCommand `json:"command"`
	Args    []string           `json:"args"`
	From    []string           `json:"from"`
	To      []string           `json:"to"`
	Path    []string           `json:"path"`
}

const (
	Exec               InstallTaskCommand = "exec"
	Copy               InstallTaskCommand = "copy"
	AddToPath          InstallTaskCommand = "add_to_path"
	ExtCmd             InstallTaskCommand = "ext_cmd"
	Remove             InstallTaskCommand = "rm"
	CallPackageManager InstallTaskCommand = "pkgman"
)

type Package struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Dependendencies []string `json:"dependencies"`

	FetchPhase []FetchPath `json:"fetch_phase"`

	UnpackPhase []UnpackTask `json:"unpack_phase"`

	InstallPhase []InstallTask `json:"install_phase"`
}

// pathToPackage is a directory where package stores
// target is for what OS package should be builded
func NewPackageItem(pathToPackage, target string) (pkg *Package, err error) {
	var (
		filePackage         *os.File
		targetPackage       *os.File
		pathToTargetPackage string = path.Join(pathToPackage, "package_"+target+".json")
	)

	pkg = &Package{}

	if filePackage, err = os.Open(path.Join(pathToPackage, "package.json")); err != nil {
		err = fmt.Errorf("Failed to open package's file:\n%s\n", err)
		return
	}
	defer filePackage.Close()

	if err = json.NewDecoder(filePackage).Decode(pkg); err != nil {
		err = fmt.Errorf("Failed to decode package.json:\n%s\n", err)
		return
	}

	if _, err = os.Stat(pathToTargetPackage); err == nil && target != "" {
		log.Debugln("OS-depend package found:", pathToTargetPackage)
		phases := &Package{}

		if targetPackage, err = os.Open(pathToTargetPackage); err != nil {
			err = fmt.Errorf("Failed to open '%s':\n%s\n", pathToTargetPackage, err)
			return
		}
		defer targetPackage.Close()

		if err = json.NewDecoder(targetPackage).Decode(phases); err != nil {
			err = fmt.Errorf("Failed to decode '%s':\n%s\n", pathToTargetPackage, err)
			return
		}

		pkg.concatPkgPhases(phases)
		log.Debug(
			"F: %v\nU: %v\nI: %v\n",
			pkg.FetchPhase, pkg.UnpackPhase, pkg.InstallPhase,
		)
	} else if err != nil {
		err = nil
	}

	return
}

func phasesInfo(data *string, phase ...any) {
	if len(phase) != 0 {
		for _, item := range phase {
			*data = fmt.Sprintf("%s\t%v\n", *data, item)
		}
	}
}

func (pkg *Package) Info() (data string) {
	data = fmt.Sprintf("Name: %s\n", pkg.Name)
	data += fmt.Sprintf("Description: %s\n", pkg.Description)
	data += fmt.Sprintf("Depends on: %v\n", pkg.Dependendencies)

	data += fmt.Sprintln("Fetching:")
	phasesInfo(&data, pkg.FetchPhase)

	data += fmt.Sprintln("Unpacking:")
	phasesInfo(&data, pkg.UnpackPhase)

	data += fmt.Sprintln("Installation:")
	phasesInfo(&data, pkg.InstallPhase)

	return
}

func (pkg *Package) concatPkgPhases(phases *Package) {
	if len(phases.FetchPhase) > 0 {
		log.Debugln("Fetch phase found")
		pkg.FetchPhase = append(pkg.FetchPhase, phases.FetchPhase...)
	}

	if len(phases.UnpackPhase) > 0 {
		log.Debugln("Unpack phase found")
		pkg.UnpackPhase = append(pkg.UnpackPhase, phases.UnpackPhase...)
	}

	if len(phases.InstallPhase) > 0 {
		log.Debugln("Install phase found")
		pkg.InstallPhase = append(pkg.InstallPhase, phases.InstallPhase...)
	}

	return
}
