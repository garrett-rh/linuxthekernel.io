---
id: "UserDrivenDesign"
title: "User Driven Design"
date: "2025-1-31"
summary: "My opinions on user driven design"
tags: ["User", "Driven", "Design", "User Driven Design"]
---

# User-Driven Design (UDD)

One of my favorite ways to design a product is user-driven. I've found success in designing my
software with the end user directly in mind. This design methodology is handy when 
creating a product for highly technical individuals.

## Why This is The Way

Feedback is the greatest driver of good software. You have to create a short feedback loop with your customer base if 
your product will be continually useful to them. As their needs change, you'll have to adapt the software to 
reflect the latest requests. The hardest part of it all is finding enthusiastic users who will provide you with frequent 
and helpful feedback. User driven design helps drive users to your product and helps you create the feedback loop.

User-driven design has the goal of creating that feedback loop. Getting new users is hard and convincing people to change
already existing workflows is even harder. What can smooth it all over is a product that has great ergonomics with no 
exception to performance.

UDD is particularly great when designing products for other engineers. I enjoy utilizing this design philosophy when 
writing software for highly specialized engineers like those with Machine Learning, Artificial Intelligence and Data focuses. 
These folks are incredibly talented at their subset of computer engineering, but more often than not, do not care about 
the nuance involved in modern day infrastructure and developer environments. For these audiences, I've found UDD to work 
well as it allows the designer to abstract away things like authentication and default configuration in favor of reasonable 
defaults.

<img src="https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExdWF4NDk1MHBoM2JyZjJ0cmFtbWYxamEwOWIyZGUweGRvcDU3ZTNnciZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/aCatQNctAK7PC1H4zh/giphy.gif"  class="img-fluid" alt="The Way"/>

## How to Implement UDD

In my experience, UDD is best implemented utilizing conventions over configuration. When writing software for a limited 
audience (like your fellow engineers at work), you can make safe assumptions about the environments you exist in. 
These types of environments are great places to implement tools that utilize UDD. For example, take a project that has 
the following structure for generating Sphinx Documentation:

```
myproject
├── main.py
└── ...
docs
├── build
├── make.bat
├── Makefile
└── source
   ├── conf.py
   ├── index.rst
   ├── _static
   └── _templates
ci 
└── ci-pipeline.yml
```

All in all, this repo will probably have 200+ lines of boilerplate documentation generation that will require constant upkeep
by each individual team. Much of the upkeep will be on the `conf.py` and the `ci-pipeline.yml`. In case you are 
unfamiliar with what these files can contain, you can find a link to these files inside the `pydata-sphinx-theme` repository.

- [CI Pipeline](https://github.com/pydata/pydata-sphinx-theme/blob/main/.github/workflows/CI.yml)
- [Sphinx Config File](https://github.com/pydata/pydata-sphinx-theme/blob/main/docs/conf.py)
 
As you can see, these files can get pretty hefty, pretty fast.

This also doesn't include all the extra work associated with hosting and publishing your docs site. 
The most immediate impact would be template out the `./docs/source/conf.py` file in order to create some standard docs 
formatting. This will also reduce developer burden for maintaining that file. Next, you would likely want to template the docs 
generation/publishing stage in the `./ci/ci-pipeline.yml` file. Offloading the logic of this file to a separate repository 
allows other potential users to become a part of the ecosystem.

The pseudocode for the new CI would look something like this:

```yaml
stages:
  - docs-gen

docs-gen:
  - template: docs-generator
  - ref: main
```

The old CI file would be migrated to the `docs-generator` template repository. This template repository would contain 
a lot of the same information from the original repository, but with the added benefit of being repository agnostic.

You'll also be able to completely replace the bespoke conf.py with one that is centrally managed and template'd out. This 
conf.py template would exist inside the `docs-generator` repository with widely agreed upon defaults. All repository specific 
customization, like project names or authors, would be pulled from the calling repository and injected into the template.
