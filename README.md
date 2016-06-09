# simpleslide - slideshows made simple

**This is a proof of concept, almost anything is still to be done**

This is for anyone who wants to create browser-based slideshows without installing the whole javascript universe.
Slides are written in Markdown and the generator is a single go binary.

## API

```sh
$ ./simpeslide --help
	build slides.md [slides.html]	- generates the slideshow as a self-contained HTML file
	show [localhost:12345]			- runs a local webserver and serves the generated slideshow
```
