# typed
[![Go](https://github.com/JohnMurray/typed/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/JohnMurray/typed/actions/workflows/go.yml)

Typed is a simple and pragmatic set of type wrappers for common use-cases with Go 1.18+. This library
aims to be simple and __usable in production__ environments.

The Go team decided [_not to update libraries_][no_change] in 1.18 alongside the release of generics.
I strongly agree with this move, and so this library aims to provide wrappers for common operations
and types as both an experiment in typing in Go 1.18+ and also a transitionary package as the standard
library evolves over future releases.


  [no_change]: https://github.com/golang/go/issues/48918