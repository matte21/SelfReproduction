# SelfReproduction

This repo contains a tiny program whose output is an exact copy of its source
code (AKA a self-reproducing program).

I did it out of boredom after reading about the idea in [Ken Thompson's turing award lecture](https://www.archive.ece.cmu.edu/~ganger/712.fall02/papers/p761-thompson.pdf).

I picked Go because that's the language I've been using the most recently.

To build the binary, simply navigate to the root directory of the repo and run

```bash
go build
```