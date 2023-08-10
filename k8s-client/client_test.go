package k8s_client

import (
	"encoding/json"
	v1 "k8s.io/api/core/v1"
	"testing"
)

func TestNewKubeClient(t *testing.T) {
	g := &GetPodCommand{
		BaseCommand: BaseCommand{
			Namespace: "default",
		},
		PodName: "nginx-77b4fdf86c-rpl5x",
	}
	client := NewKubeClient("")
	r, err := client.ExecCommand(g)

	if err != nil {
		panic(err)
	}

	pod := r.(*v1.Pod)

	if d, err := json.Marshal(pod); err == nil {
		t.Logf("%v", string(d))
	}

}
