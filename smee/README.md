# Boots

[![Build Status](https://cloud.drone.io/api/badges/tinkerbell/boots/status.svg)](https://cloud.drone.io/tinkerbell/boots)
![](https://img.shields.io/badge/Stability-Experimental-red.svg)

This services handles DHCP, PXE, tftp, and iPXE for provisions.

This repository is [Experimental](https://github.com/packethost/standards/blob/master/experimental-statement.md) meaning that it's based on untested ideas or techniques and not yet established or finalized or involves a radically new and innovative style!
This means that support is best effort (at best!) and we strongly encourage you to NOT use this in production.

### Local Setup

First, you need to make sure you have [git-lfs](https://github.com/git-lfs/git-lfs/wiki/Installation) installed:

```
# install "git-lfs" package for your OS. On Ubuntu, for instance:
# curl https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | sudo bash
# apt install git-lfs

# then run these two commands:
git lfs install
git lfs pull
```

Running the Tests

```
# make the files
make all
# run the tests
go test
```

Build/Run Boots

```
# run boots
./boots
```

You can use NixOS shell, which will have the Git-LFS, Go and others

`nix-shell`

Note: for mac users, you will need to comment out the line `pkgsCross.aarch64-multiplatform.buildPackages.gcc` in order for the build to work
