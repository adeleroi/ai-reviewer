# AI Reviewer

AI Reviewer is a code review system I'm building in Go to explore how LLMs can reason about source-code changes with richer context than a raw diff alone.

The project is being built incrementally, with an emphasis on understanding and implementing the underlying mechanisms rather than relying entirely on existing agent or code-analysis frameworks.

## Motivation

A diff shows what changed, but reviewing a change often requires understanding code outside the modified lines: the functions being called, surrounding types, related declarations, and eventually dependencies across files.

The goal of this project is to experiment with how a review system can construct that context before asking an LLM to reason about a change.

## Current Architecture

The current implementation focuses on the foundations of the review pipeline:

1. Read and represent source-code changes.
2. Parse modified source files.
3. Map changed regions back to their surrounding code structures.
4. Build structured context that can later be provided to an LLM for review.

The project is currently under active development.

## Why Go?

I'm using Go both to deepen my experience with the language and because its standard tooling exposes useful primitives for working with source code, including token positions and ASTs.

## Current Focus

I'm currently working on improving the source-context construction layer and the boundaries between diff parsing, source parsing, and the review pipeline.

Upcoming work includes:

- expanding test coverage around parsing and context construction
- improving handling of multi-file changes
- integrating model-based review on top of the structured context
- experimenting with context-selection strategies
- evaluating review quality and failure modes

## Project Status

This is an active learning and engineering project rather than a finished product. I'm intentionally building the system in stages so that each layer can be understood, tested, and evaluated independently before adding more sophisticated LLM and agent capabilities.