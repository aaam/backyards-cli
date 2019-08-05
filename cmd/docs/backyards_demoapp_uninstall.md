## backyards demoapp uninstall

Output or delete Kubernetes resources to uninstall Canary feature

### Synopsis

Output or delete Kubernetes resources to uninstall Canary feature.

The command provides the resources that can be deleted manually or
it can delete the resources automatically with --delete-resources option.

```
backyards demoapp uninstall [flags]
```

### Examples

```
  # Default uninstall.
  backyards canary uninstall | kubectl delete -f -

  # Uninstall Canary feature from a non-default namespace.
  backyards canary uninstall install -n custom-istio-ns | kubectl delete -f -
```

### Options

```
      --demo-namespace string   Namespace for demo application (default "backyards-demo")
  -d, --dump-resources          Dump resources to stdout instead of applying them
  -h, --help                    help for uninstall
```

### Options inherited from parent commands

```
      --context string      Name of the kubeconfig context to use
  -c, --kubeconfig string   Path to the kubeconfig file to use for CLI requests
  -n, --namespace string    Namespace in which Backyards is installed [$BACKYARDS_NAMESPACE] (default "backyards-system")
  -v, --verbose             Turn on debug logging
```

### SEE ALSO

* [backyards demoapp](backyards_demoapp.md)	 - install and manage demo application

###### Auto generated by spf13/cobra on 5-Aug-2019