# godel

In unix system, we use rm to remove file completely. Once delete it is hard to recovery.
To solve this, I created godel using golang language.

## Purpose
Everytime you delete files, there is backup gzip archive file in the /tmp directory.
you can recover anytime until you shutdown or reboot the system.

## Prequisites
Make sure golang is installed in your machine.
```
wget -v https://go.dev/dl/go1.17.7.linux-amd64.tar.gz && tar -xzvf go1.17.7.linux-amd64.tar.gz && tar -xzvf go1.17.7.linux-amd64.tar.gz && sudo mv go/ /var/local/. && echo "export PATH=$PATH:/var/local/go/bin" >> ~/.profile
```
Then restart.
after this you can write this command
```
$go version
go version go1.17.6 linux/amd64
```

## Installation
```
go build -o godel main.go && sudo mv godel /usr/local/bin/.
```

## Command
For example(1), to delete "test.txt"
> godel test.txt

If you don't want to create backup, just add "-d" or "--delete-all"
> godel -d test.txt
or 
> godel --delete-all text.txt

## Versions
 - v0.0.1 (released on 3rd Mar 2022)

## Note
This is just beginning.
