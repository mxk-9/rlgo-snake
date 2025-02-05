package deptree

import (
	"fmt"
	"path"
	"spm/pkg/pkginfo"
	log "spm/pkg/slog"
)

type PkgData struct {
	PkgsPath string
	Target   string
}

type Node struct {
	Data    *PkgData
	Pkg     *pkginfo.Package
	Depends []*Node // If len(Depends) is 0, we reach the end
}

type Tree struct {
	Nodes *Node
	Data  PkgData
}

// pkgs — is a path to packages
func NewDepTree(pkgs, packageName, target string) (depTree *Tree,
	err error) {

	depTree = &Tree{
		Data: PkgData{
			PkgsPath: pkgs,
			Target:   target,
		},
	}

	log.Debugln("Creating dependency tree")
	if depTree.Nodes, err = NewNode(&depTree.Data, packageName); err != nil {
		err = fmt.Errorf("Failed to build dependency tree:\n%s\n", err)
	}

	return
}

func NewNode(data *PkgData, pkgName string) (depNode *Node, err error) {
	// Creates simple dependency node
	depNode = &Node{
		Data: data,
	}

	pkg := &pkginfo.Package{}

	log.Debug("Creating package item '%s'", pkgName)
	if pkg, err = pkginfo.NewPackageItem(path.Join(data.PkgsPath, pkgName),
		data.Target); err != nil {
		return
	}

	log.Debugln("Looking for dependencies...")
	for _, item := range pkg.Dependendencies {
		log.Debug("Found '%s', appending to list", item)
		if err = depNode.Append(depNode.Data, item); err != nil {
			return
		}
	}
	log.Debug("Success fetching dependencies for package '%s'", pkgName)

	depNode.Pkg = pkg

	return
}

func (dn *Node) Append(data *PkgData, pkgName string) (err error) {
	// Expand dependency nodes array
	localNode := &Node{}

	log.Debug("Creating node '%s'", pkgName)
	if localNode, err = NewNode(dn.Data, pkgName); err != nil {
		err = fmt.Errorf("Failed to create node: %s\n", err)
		return
	}

	log.Debugln("Appending")
	dn.Depends = append(dn.Depends, localNode)

	return
}

func (dn *Node) ShowNode() {
	if dn.Pkg == nil {
		return
	}

	for _, item := range dn.Depends {
		item.ShowNode()
	}

	log.Infoln(dn.Pkg.Info())
	fmt.Println()
}

func (dn *Node) InstallNode() (err error) {
	if dn.Pkg == nil {
		err = fmt.Errorf("Empty node")
		return
	}

	for _, item:= range dn.Depends {
		item.InstallNode()
	}

	// TODO: Fetch phase
	// TODO: Unpack phase
	// TODO: Install phase
	// OPTIONAL: Uninstall phase
	return
}

func (dp *Tree) ShowTree() {
	log.Infoln("Packages path:", dp.Data.PkgsPath)
	log.Infoln("Target:", dp.Data.Target)

	dp.Nodes.ShowNode()
}

func (dp *Tree) Install() (err error) {
	// Going through all dependencies packages until len(Depends) == 0 then
	// recursively install everything.
	// For each phase we have one package

	return
}
