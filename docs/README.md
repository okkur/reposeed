<!--
Copyright 2017 - The RepoSeed authors
This work is licensed under a Creative Commons Attribution-ShareAlike 4.0 International License;
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://creativecommons.org/licenses/by-sa/4.0/legalcode
Unless required by applicable law or agreed to in writing, documentation
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->


## Development
Run the following commands to install reposeed and change your directory to reposeed's source code directory.
```
go get -v -u github.com/okkur/reposeed/cmd/reposeed/...
cd $GOPATH/src/github.com/okkur/reposeed
```
Then run following command to install **packr**
```
make packr
```
Then in case you made any changes on  templates, just run ```make``` or ```packr install ./cmd/reposeed``` to bundle templates in the generated binary.
# Documentation

No documentation available yet. Start your first contribution with some documentation.

See how you can contribute with our [contribution guide](/CONTRIBUTING.md).
