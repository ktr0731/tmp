# tmp
[![CircleCI](https://circleci.com/gh/lycoris0731/tmp.svg?style=svg&circle-token=b6ab9824b5dd9a252a19d98c64a3aaa31361f2db)](https://circleci.com/gh/lycoris0731/tmp)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)  
Manage your all temporary directories easily

===

## Description  
When we want to experiments, tests, and so on, we'll create temporary directory.  
However, almost temporary directories are not deleted that became unnecessary.  
Accordingly, many unnecessary directories are scattered around here and there.  

`tmp` manages that directories.  
You can ...
- look all existing temporary directories.  
- remove the directory by using `tmp`. Or remove all.  
- jump to the directory.

## Installation
``` sh
$ go get github.com/lycoris0731/tmp
```

## Usage
Add new temporary directory(it is called `item`).  
If there is no arguments, be used `tmp` for name.

``` sh
$ tmp make 
$ tmp mk
$ tmp mk temp
$ tmp mk ../
```

Look all temporary directories.  
If `-n` argument is passed, the command shows identifier number with paths.  

``` sh
$ tmp list
$ tmp ls
$ tmp ls -n
```

Remove the directory.  
You can pass arguments by full path, identifier number or both.  

``` sh
$ tmp remove [target path] [target id]
$ tmp rm [target path...] [target id...]
```

Remove all directries.
``` sh
$ tmp rm --all
```

## License
Please see [LICENSE](./LICENSE).
