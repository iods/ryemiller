Personal Website
================

Development
-----------

Vue frontend, Go binary.

API running in Go for a health dashboard and interactive CV, among other future fun apis.

it is compiled, then a static file server runs the vue app, to show an example of the cross-functional API, the forms
are actually JS post functions acting as forms. They are still forms; however, JS handles the submission, gets the data
and posts it to a Go URL. That url then gets the data, and uses a Sendmail API to email me that someone stopped by. Nothing
special but it is a start.

### Blockers

One blocker was the understanding of how to integrate a server side app into a binary one. most developers used completely 
separate project environments, and I didn't like not being able to compile everything as one.