<p align="center">
  <img src="https://cis.cornell.edu/sites/default/files/styles/square_thumbnail/public/Screen%20Shot%202018-09-04%20at%2010.17.51%20AM.png?itok=dxbUcd1u" width="80" alt="arXiv">
  <img src="https://upload.wikimedia.org/wikipedia/commons/1/17/GraphQL_Logo.svg" width="80" alt="GraphQL">
</p>

# arxiv-graphql

GraphQL API for arXiv: e-print service by Cornell University

## About

### arXiv

arXiv® is a free distribution service and an open archive for scholarly articles in the fields of physics, mathematics, computer science, quantitative biology, quantitative finance, statistics, electrical engineering and systems science, and economics. arXiv is a collaboratively funded, community-supported resource founded by Paul Ginsparg in 1991 and maintained and operated by Cornell University.

Operations are maintained by the arXiv Leadership Team and arXiv staff at Cornell, with the help of numerous volunteer subject moderators. Governance of arXiv is led by the Leadership Team with guidance from the arXiv Scientific Advisory Board and the arXiv Member Advisory Board. arXiv is funded by Cornell University, the Simons Foundation, member institutions, and donors.

Registered users may submit articles to be announced by arXiv. Submissions to arXiv are subject to a moderation process that verifies material is topical to subject areas and has scholarly value. Material is not peer-reviewed by arXiv.

### GraphQL

GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data. GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools.

## Getting Started

### Installation and Setup

- Download and Install [Golang](https://golang.org/dl/)

- Install dependencies

```bash
make install
```

### Running

```bash
export PORT=<PORT> && go run main.go
```

### Testing

```bash
make test
```

### Code Coverage

```bash
make cover
```

## Issues

Issues are managed via [GitHub Issues](https://github.com/prabhuomkar/arxiv-graphql/issues).
