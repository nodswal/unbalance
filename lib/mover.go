package lib

import (
	"apertoire.net/unbalance/message"
	"github.com/golang/glog"
	"io"
	"os"
	"path/filepath"
)

type PassThru struct {
	io.Reader
	Progress *message.ProgressStatus
	Ch       chan *message.ProgressStatus
}

func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)

	if err == nil {
		pt.Progress.CurrentCopied += uint64(n)
		pt.Ch <- pt.Progress
	}

	return n, err
}

type Mover struct {
	Src      string
	Dst      string
	err      error
	Progress *message.ProgressStatus

	ProgressCh chan *message.ProgressStatus
	DoneCh     chan bool
}

func NewMover() *Mover {
	progressCh := make(chan *message.ProgressStatus)
	doneCh := make(chan bool)

	return &Mover{ProgressCh: progressCh, DoneCh: doneCh}
}

func (self *Mover) visit(path string, info os.FileInfo, err error) (e error) {
	if info.IsDir() {
		out := "Path: " + path + " --- Dest: " + self.Dst
		newPath := filepath.Join(self.Dst, path[len(self.Src):])
		glog.Infof("A (%s) -- B (%s)", out, newPath)
	} else {
		out := "Path: " + path + " --- nDest: " + self.Dst
		newPath := filepath.Join(self.Dst, path[len(self.Src):])
		glog.Infof("A (%s) -- B (%s)", out, newPath)
		self.DoneCh <- true

		// self.Progress.CurrentFile = path
		// self.Progress.CurrentSize = uint64(info.Size())

		// in, err := os.Open(self.Src)
		// if err != nil {
		// 	return
		// }

		// defer in.Close()
		// out, err := os.Create(self.Dst)
		// if err != nil {
		// 	return
		// }
		// defer func() {
		// 	cerr := out.Close()
		// 	if err == nil {
		// 		err = cerr
		// 	}
		// }()
		// if _, err = io.Copy(out, in); err != nil {
		// 	return
		// }
		// err = out.Sync()

		self.ProgressCh <- self.Progress
	}

	return nil
}

// func (self *Mover) Copy() (chan *message.ProgressStatus, chan bool) {
func (self *Mover) Copy() {
	// self.progressCh = make(chan *message.ProgressStatus)
	// self.doneCh = make(chan bool)

	go func() {
		filepath.Walk(self.Src, self.visit)
		// defer close(self.progressCh)
		// defer close(self.doneCh)
	}()

	// return self.progressCh, self.doneCh
}
