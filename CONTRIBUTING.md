Neighborhood Manager Contribution Guidelines
============================================

## Dependency Management  
### Adding a library dependency  
When a new library is used for the first time, it must be added to Glide. Open the file `glide.yaml` and add a new line under the `import` heading:
```
import:
- package: github.com/king-jam/gossdp
```
replacing the example GitHub link with the path that `go get` needs to grab the library.

When this is done, save `glide.yaml` and run `glide update` to copy the dependent source code to the `vendor/` folder.  

After this, the steps in the next section should be done to lock in the specific version of the dependent library.

### Updating a library dependency
When a dependent library has an updated commit that is desired (such as a big fix, added feature, etc), the library's entry in `glide.lock` should be updated. Open the file and find its entry, such as
```
- name: github.com/hashicorp/consul
  version: 36dc9201f2e006d4b5db1f0446b17357811297bf
```
Replace the existing commit hash with the hash of the new desired commit.  
Save and close the file, then run `glide install`  
