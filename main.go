package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

var (
	Options map[string][]string = map[string][]string{
		"delete_all": {
			"--delete-all",
			"-d",
		},
	}
)

func PrintHelp() {
	fmt.Println("godel <filename : string>")
}

func CheckFileExist(path string) (string, bool) {
	file, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error on file exist check:", err)
		return err.Error(), false
	}
	return file.Name(), true
}

type InputFlag struct {
	IsDeleteAll bool
	Files       map[string]string
}

func (inputFlag *InputFlag) ParseFlag(args []string) error {
	for k, v := range args {
		if k == 0 {
			if v == "-d" || v == "--delete-all" {
				inputFlag.IsDeleteAll = true
				continue
			} else {
				inputFlag.IsDeleteAll = false
			}
		}
		filename_or_mesg, ok := CheckFileExist(v)
		if ok {
			inputFlag.Files[filename_or_mesg] = v
		} else {
			return errors.New(filename_or_mesg)
		}
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Println("Please provide filename that you want to delete")
		PrintHelp()
		return
	}

	inputs := InputFlag{Files: map[string]string{}}

	err := inputs.ParseFlag(args[1:])
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	for k, v := range inputs.Files {
		if inputs.IsDeleteAll {
			f, err := os.Stat(v)
			if err != nil {
				fmt.Println("Error checking on file or directory:", err.Error())
				continue
			}
			if f.IsDir() {
				os.RemoveAll(v)
			} else {
				os.Remove(v)
			}
		} else {
			f, err := os.Stat(v)
			if err != nil {
				fmt.Println("Error checking on file or directory:", err.Error())
				continue
			}
			temp_file := path.Join(os.TempDir(), fmt.Sprintf("%v.tar.gz", k))
			file, err := os.Create(temp_file)
			if err != nil {
				fmt.Println("Error on creating temp tar.gz file", err.Error())
				continue
			}
			defer file.Close()
			gf := gzip.NewWriter(file)
			defer gf.Close()
			tf := tar.NewWriter(gf)
			defer tf.Close()
			rfile, err := os.Open(v)
			if err != nil {
				fmt.Println("Error on reading :", err.Error())
				continue
			}
			if err != nil {
				fmt.Println("Error on file info reading :", err.Error())
				continue
			}
			header, err := tar.FileInfoHeader(f, f.Name())
			if err != nil {
				fmt.Println("Error on getting tar header :", err.Error())
				continue
			}
			header.Name = k
			err = tf.WriteHeader(header)
			if err != nil {
				fmt.Println("Error on writing tar header :", err.Error())
				continue
			}
			_, err = io.Copy(tf, rfile)
			if err != nil {
				fmt.Println("Error on copying to tar.gz :", err.Error())
				continue
			}
			os.Remove(k)
		}
	}
}
