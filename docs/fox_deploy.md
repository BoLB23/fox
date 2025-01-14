## fox deploy

Deploy KubeFox App using the component code from the currently checked out Git commit

```
fox deploy [NAME] [flags]
```

### Options

```
  -t, --create-tag         create Git tag using the AppDeployment version
      --dry-run            submit server-side request without persisting the resource
  -h, --help               help for deploy
  -n, --namespace string   namespace of KubeFox Platform
  -p, --platform string    name of KubeFox Platform to utilize
  -s, --version string     version to assign to the AppDeployment, making it immutable
      --wait duration      wait up the specified time for components to be ready
```

### Options inherited from parent commands

```
  -a, --app string      path to directory containing KubeFox App
  -i, --info            enable info output
  -o, --output string   output format, one of ["json", "yaml"] (default "yaml")
  -v, --verbose         enable verbose output
```

### SEE ALSO

* [fox](fox.md)	 - CLI for interacting with KubeFox

