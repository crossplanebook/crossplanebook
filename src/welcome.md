# Welcome

Welcome to the Crossplane book! This book exists for anyone who wants to learn
more about the open source [Crossplane](https://crossplane.io) project. It
explores concepts from a technical perspective with a goal of enabling a reader
to contribute to and extend the project for their use-case. However, it does not
assume that the reader is already familiar with [Go](https://golang.org/),
[Kubernetes](https://kubernetes.io/), or the [controller
pattern](https://kubernetes.io/docs/concepts/architecture/controller/#controller-pattern).
It starts with the fundamentals and builds to deeper concepts.

Where possible, we will always refer to existing guides and documentation. There
are copious resources available on every topic we will cover, and there is no
need to re-invent the wheel. We take the same stance on using existing
frameworks and libraries. The only reason not to use existing code is if you
know or suspect that it will not meet your needs in the short to medium term.

As you learn more about Crossplane, you will understand how different it is from
other Kubernetes components. However, Crossplane integrates with with Kubernetes
because:

1. It allows for standardization on a single API, making it easy to integrate
   into your existing workflow, and to play nice with other projects.
2. It provides a production-grade scheduling engine out-of-the-box.

If you are already familiar with Go, Kubernetes, and the controller pattern, you
can jump right into learning about [Crossplane](crossplane/crossplane.md).
Otherwise, let's [get started](go/go.md)!