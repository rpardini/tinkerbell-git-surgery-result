# Tinkerbell

[![Build Status](https://cloud.drone.io/api/badges/tinkerbell/tink/status.svg)](https://cloud.drone.io/tinkerbell/tink)
[![codecov](https://codecov.io/gh/tinkerbell/tink/branch/master/graph/badge.svg)](https://codecov.io/gh/tinkerbell/tink)
![](https://img.shields.io/badge/Stability-Experimental-red.svg)

This repository is [Experimental](https://github.com/packethost/standards/blob/master/experimental-statement.md) meaning that it's based on untested ideas or techniques and not yet established or finalized or involves a radically new and innovative style! This means that support is best effort (at best!) and we strongly encourage you to NOT use this in production.

It is comprised of following five major components:

1.  A DHCP server ([boots](https://github.com/tinkerbell/boots))
2.  A workflow engine (tink, this repository)
3.  A metadata service ([hegel](https://github.com/tinkerbell/hegel))
4.  An in-memory installation environment([osie](https://github.com/tinkerbell/osie))
5.  A controller/handler of BMC interactions([pbnj](https://github.com/tinkerbell/pbnj))

The workflow engine is comprised of a server and a CLI, which communicates over gRPC.
The CLI is used to create a workflow and its building blocks: templates and targeted hardware.

## Workflow

A workflow is a framework responsible for handling flexible, bare metal provisioning, that is...

-   standalone and does not need the Packet API to function
-   contains `Boots`, `Tink`, `Hegel`, `OSIE`, `PBnJ` and workers
-   can bootstrap any remote worker using `Boots + Hegel + OSIE + PBnJ`
-   can run any set of actions as Docker container runtimes
-   receive, manipulate, and save runtime data

## Website

For complete documentation, please visit the Tinkerbell project hosted at [tinkerbell.org](https://tinkerbell.org).
