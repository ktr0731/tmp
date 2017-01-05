# tmp
[![CircleCI](https://circleci.com/gh/lycoris0731/clip.svg?style=svg&circle-token=0e33c0cfb7bb1105ff821abbe845483d269145f8)](https://circleci.com/gh/lycoris0731/clip)
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
`-n` option decides item's name. If nothing, be used `tmp` for name.

``` sh
$ tmp make 
$ tmp mk
```

Look all temporary directories.  

``` sh
$ tmp list
$ tmp ls
```

Remove the directory.
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
