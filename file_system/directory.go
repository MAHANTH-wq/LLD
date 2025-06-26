package main

type directory struct {
	name        string
	fileSystems []fileSystem
}

func getNewDirectory(name string) *directory {
	return &directory{
		name: name,
	}
}

func (d *directory) addToDirectory(fileSystem fileSystem) {

	d.fileSystems = append(d.fileSystems, fileSystem)

}
func (d *directory) ls() {

	for _, fileSystem := range d.fileSystems {
		fileSystem.ls()
	}
}
