package k8s_client

type command interface {
	Execute(cli *KubeClient) (any, error)
}

type BaseCommand struct {
	Namespace string
}
