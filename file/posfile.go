package file

type PosFile struct {
	Inode    int
	LastLine int
	Path     string
	FileName string
}

func (p *PosFile) write() {

}

func (p *PosFile) Update(lastline int) {
	p.LastLine = lastline
	p.write()
}

func (p *PosFile) getInode(lastline int) int {
	if p.Inode != 0 {
		return p.Inode
	}
	return 0
}
