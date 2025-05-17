package main

func main() {

	var kindOrder = map[string]int{}

	kinds := []string{
		"Namespace",
		"NetworkPolicy",
		"ResourceQuota",
		"LimitRange",
		"PodSecurityPolicy",
		"PodDisruptionBudget",
		"ServiceAccount",
		"Secret",
		"SecretList",
		"ConfigMap",
		"StorageClass",
		"PersistentVolume",
		"PersistentVolumeClaim",
		"CustomResourceDefinition",
		"ClusterRole",
		"ClusterRoleList",
		"ClusterRoleBinding",
		"ClusterRoleBindingList",
		"Role",
		"RoleList",
		"RoleBinding",
		"RoleBindingList",
		"Service",
		"DaemonSet",
		"Pod",
		"ReplicationController",
		"ReplicaSet",
		"Deployment",
		"HorizontalPodAutoscaler",
		"StatefulSet",
		"Job",
		"CronJob",
		"Ingress",
		"APIService",
	}
	for i, kind := range kinds {
		// make sure none of the above entries are zero, we need that for custom resources
		kindOrder[kind] = i - len(kinds)
	}

	// Loop though all items in kindOrder with range loop and print key and value
	for kind, order := range kindOrder {
		println(kind, order)
	}

}
