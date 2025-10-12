# Houston Linux User's Group Website Sources

The Houston Linux User's Group is the longest lived, most active, Linux User's
Group in Houston, Texas.  Founded in 1995, it has moved locations a number of
times, and currently holds weekly meetings.

This repository contains the source content and generation tools to render
the current statically presented Houston Linux User's Group website, located
at http://houstonlinux.org/, as well as the soon to be new site with hugo (not
yet hosted anywhere).

Since the repo contains both the maven site and hugo site there are 2 ways to
run and develop the site. The following information is on how to run the new
development hugo site and not the maven one since it will not get any updates
now.

# Motivation

To simplify security, we opted to host static web pages only.  However, the
costs of maintaining static web pages written in HTML with modern standards
and decent styling are high.  To this end, we leveraged the use of bootstrap
combined with hugo to create stylish static sites.

# Usage

The casual user should browse to http://houstonlinux.org/ and see the current
publication of these sources.

The bold may fork or clone these files to leverage the project in the generation
of their own website, however we ask that you not misrepresent your efforts as
those of the Houston Linux User's Group.

# Requirements

The following are required to build and run the site:

- [Go](https://go.dev/doc/install) v1.24.0 or later
- [Hugo Extended](https://github.com/gohugoio/hugo) v0.149.0 or later

# Development

To run the development site clone the repo, cd into the directory and run the
`hugo server` command:

```bash
git clone https://github.com/edwbuck/hlug-website.git
cd hlug-website
hugo server
```
Note:

We recommend using the flags `--noHTTPCache` and `--disableFastRender` to
help mitigate any issues of things not getting updated because of cached
files.

<details>
<summary>Maven Instructions (Legacy)</summary>

# Installation

The official web site is installed via the mvn command:

    mvn clean site:site site:stage site:deploy -P website

note that this command will not work unless you have the appropriate keys
registered for the ssh channel to the live website.  

To build a local rendering of the web site for testing purposes:

    mvn clean site:site site:stage -P website

and open a web browser to view the 

    file:///path/to/your/root/target/website-site/index.html

index.html file.
</details>
