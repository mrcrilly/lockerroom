# TinySafe API

This RESTful API is used for storing received secrets in a SQLite DB.

## Goals

- RESTful API;
- Client certificates for authentication (optional);
- Encrypted DB file;
- 2-Factor Authentication (using another of my projects: Gate2);
- Self-hosted password manager;
- Simple and multi-platform;
- Vagrantfile for easily bringing up a local copy for testing, development, etc;
- On the back of the Vagrantfile, an Ansible Playbook for downloading and deploying the code anywhere;

## Overall Architecture

1. The TinySafe API is brought online on a remote UNIX system;
1. [TinySafe]() is configured locally on your Mac, Windows, or Linux desktop/laptop;
1. You then visit http://localhost:4869 locally and begin adding safes for storing secrets;

### Further

The TinySafe API application is the remote system. The TinySafe application is a local system run on your home computer, but is in fact also an API its self. It acts as a service, silently running in the background, waiting for you to connect to it and configure it. As well as being an API, it also serves a simple Bootstrap powered UI (on port 4869) locally.

The rationale behind this is simply to allow anyone to write whatever client they need, be it a Python CLI client, or a Swift iOS app connecting on the local network (or over a WAN?)

## Author

- Michael Crilly
- http://mcrilly.me/
- @mrmcrilly
