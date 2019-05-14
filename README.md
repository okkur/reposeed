<a href='https://reposeed.okkur.org'><img src='https://raw.githubusercontent.com/resumic/schema/master/media/logo.svg?sanitize=true' width='500'/></a>

Extensive repository base files up and running in an instant

 [![state](https://img.shields.io/badge/state-beta-blue.svg)]() [![release](https://img.shields.io/github/release/okkur/reposeed.svg)](https://github.com/okkur/reposeed/releases) [![license](https://img.shields.io/github/license/okkur/reposeed.svg)](LICENSE)

**NOTE: This is a beta release, we do not consider it completely production ready yet. Use at your own risk.**

# Reposeed
Start with the base layer necessary to focus on your project and not on the repository.  
Licensing, structure, documentation and more boilerplate to get you started from your first commit.

## Using Reposeed
```
go get -v -u github.com/okkur/reposeed/cmd/reposeed/...
mkdir YOUR-PROJECT
```

Run the following command to create the sample .seed-config.yaml file.
```
reposeed init YOUR-PROJECT
```

Change .seed-config.yaml to your needs.
```
reposeed --conf=YOUR-PROJECT/.seed-config.yaml --output=YOUR-PROJECT
```
Take a look at our full [documentation](/docs).

## Support
For detailed information on support options see our [support guide](/SUPPORT.md).

## Helping out
Best place to start is our [contribution guide](/CONTRIBUTING.md).

----

*Code is licensed under the [Apache License, Version 2.0](/LICENSE).*  
*Documentation/examples are licensed under [Creative Commons BY-SA 4.0](/docs/LICENSE).*  
*Illustrations, trademarks and third-party resources are owned by their respective party and are subject to different licensing.*

*The Reposeed logo was created by [Florin Luca](https://99designs.com/profiles/florinluca)*

---

Copyright 2017 - The Reposeed authors
