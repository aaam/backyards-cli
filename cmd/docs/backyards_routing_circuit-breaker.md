## backyards routing circuit-breaker

Manage circuit-breaker configurations

### Synopsis

Manage circuit-breaker configurations

### Options

```
  -h, --help   help for circuit-breaker
```

### Options inherited from parent commands

```
      --context string      name of the kubeconfig context to use
      --interactive         ask questions interactively even if stdin or stdout is non-tty
  -c, --kubeconfig string   path to the kubeconfig file to use for CLI requests
  -n, --namespace string    namespace in which Backyards is installed [$BACKYARDS_NAMESPACE] (default "backyards-system")
      --non-interactive     never ask questions interactively
  -o, --output string       output format (table|yaml|json) (default "table")
  -v, --verbose             turn on debug logging
```

### SEE ALSO

* [backyards routing](backyards_routing.md)	 - Manage service routing configurations
* [backyards routing circuit-breaker delete](backyards_routing_circuit-breaker_delete.md)	 - Delete circuit breaker rules of a service
* [backyards routing circuit-breaker get](backyards_routing_circuit-breaker_get.md)	 - Get circuit breaker rules for a service
* [backyards routing circuit-breaker graph](backyards_routing_circuit-breaker_graph.md)	 - Show graph
* [backyards routing circuit-breaker set](backyards_routing_circuit-breaker_set.md)	 - Set traffic shifting rules for a service

###### Auto generated by spf13/cobra on 3-Oct-2019