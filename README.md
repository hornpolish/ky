
# ky - a kubernetes yaml swiss army knife

compare two yamls with `ky diff file1 file2`

split a monster yaml into parts with `ky split big-ole.yaml`

## install

Releases are posted by [goreleaser](https://goreleaser.com/) to the project [releases page](https://github.com/hornpolish/ky/releases). On Mac, you may prefer brew:

```lang=bash
brew install hornpolish/brew/ky
```

## credits

The diff subroutines came from [Sahil Muthoo](https://github.com/sahilm/yamldiff)

The split subroutines came from [Frederik Mogensen](https://github.com/mogensen/kubernetes-split-yaml)

Their programming style was harmonized somewhat to mine, and the combined works published as KY, which operates in the kubectl style of {command} {verb} {noun}

The github actions for golang came from/were inspired by [Bruno Paz](https://dev.to/brpaz/building-a-basic-ci-cd-pipeline-for-a-golang-application-using-github-actions-icj)

## contributing

The KY project welcomes contributions.  

* fork the repo
* clone/pull from the fork
* create a branch for the improvement
* commit to the fork/branch
* raise a merge-request to splice your improvement into the KY project
* "cheers!"

There is a nice discussion of this [contribution workflow](https://github.com/freeCodeCamp/how-to-contribute-to-open-source/blob/master/CONTRIBUTING.md) on github.


## Release
```lang=bash
vi VERSION
ver=$(cat VERSION)
git tag -a $ver -m "announce message"
git push origin $ver
```

## TODO
* some coverage for main.go
* badges in README?
