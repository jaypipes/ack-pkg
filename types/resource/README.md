# `types/resource`

The `types/resource` package contains interfaces that *define* AWS resources.

These interfaces are specifically designed in such a way as to *not* have any
dependencies on Kubernetes libraries. This is because the generated
implementations of these interfaces are intended to be used by any client that
wishes to standardize its calling of AWS service APIs using a declarative
resource management strategy but isn't dependent on the Kubernetes
controller-runtime or reconciliation logic.
