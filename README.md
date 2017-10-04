

<img src="https://github.com/okkur/reposeed/blob/master/media/logo.svg" width="500"/>



Create your next open source repository with batteries included

[![state](https://img.shields.io/badge/state-beta-blue.svg)]() [![release](https://img.shields.io/github/release/okkur/reposeed.svg)](https://github.com/okkur/reposeed/releases) [![license](https://img.shields.io/github/license/okkur/reposeed.svg)](LICENSE) 

**NOTE: This is a beta release, we do not consider it completely production ready yet. Use at your own risk.**



# RepoSeed
Start with the base layer necessary to focus on your project and not on the repository.  
Licensing, structure, documentation and more boilerplate to get you started from your first commit.



## Using RepoSeed


  ```
git clone https://github.com/okkur/reposeed.git  
mkdir YOUR-PROJECT
cp reposeed/templates/seed-config.example.yaml YOUR-PROJECT/.seed-config.yaml
```  
Change .seed-config.yaml to your needs.
```
go run cmd/generator.go --input=reposeed/templates --output=YOUR-PROJECT
```




Take a look at our full [documentation](/docs).


## Support
For detailed information on support options see our [support guide](/SUPPORT.md).

## Helping out
Best place to start is our [contribution guide](/CONTRIBUTING.md).

----

*Code is licensed under the [Apache License, Version 2.0](/LICENSE)*  
*Documentation is licensed under [Creative Commons BY-SA 4.0](/docs/LICENSE)*  

*Illustrations, trademarks and third-party resources are owned by their respective party and are subject to different licensing.*

---

Copyright 2017 - The RepoSeed authors
