********************************
Sampo Server Daemon (``sampod``)
********************************

.. image:: https://img.shields.io/badge/license-Public%20Domain-blue.svg
   :alt: Project license
   :target: https://unlicense.org

.. image:: https://img.shields.io/badge/godoc-reference-blue.svg
   :alt: GoDoc reference
   :target: https://godoc.org/github.com/sampoapp/sampo-server

.. image:: https://goreportcard.com/badge/github.com/sampoapp/sampo-server
   :alt: Go Report Card score
   :target: https://goreportcard.com/report/github.com/sampoapp/sampo-server

.. image:: https://img.shields.io/travis/sampoapp/sampo-server/master.svg
   :alt: Travis CI build status
   :target: https://travis-ci.org/sampoapp/sampo-server

|

Installation
============

::

   $ go get -u github.com/sampoapp/sampo-server/sampod

Reference
=========

::

   Sampo is a personal information manager (PIM) app.
   This is the server daemon for Sampo.

   Usage:
     sampod [flags]

   Flags:
     -C, --config string   Set config file (default: $HOME/.sampo/config.yaml)
     -d, --debug           Enable debugging
     -h, --help            help for sampod
     -v, --verbose         Be verbose
         --version         version for sampod
