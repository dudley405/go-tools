package types

import (
    "time"
)

type KubeResponse struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			GenerateName      string    `json:"generateName"`
			Namespace         string    `json:"namespace"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			Generation        int       `json:"generation"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				App             string `json:"app"`
				PodTemplateHash string `json:"pod-template-hash"`
			} `json:"labels"`
			OwnerReferences []struct {
				APIVersion         string `json:"apiVersion"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				UID                string `json:"uid"`
				Controller         bool   `json:"controller"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
			} `json:"ownerReferences"`
			ManagedFields []struct {
				Manager    string    `json:"manager"`
				Operation  string    `json:"operation"`
				APIVersion string    `json:"apiVersion"`
				Time       time.Time `json:"time"`
				FieldsType string    `json:"fieldsType"`
				FieldsV1   struct {
					FMetadata struct {
						FGenerateName struct {
						} `json:"f:generateName"`
						FLabels struct {
							NAMING_FAILED struct {
							} `json:"."`
							FApp struct {
							} `json:"f:app"`
							FPodTemplateHash struct {
							} `json:"f:pod-template-hash"`
						} `json:"f:labels"`
						FOwnerReferences struct {
							NAMING_FAILED struct {
							} `json:"."`
							KUIDC3Bf4Da35D344705A7Dc13D8F7Ec7E79 struct {
							} `json:"k:{"uid":"c3bf4da3-5d34-4705-a7dc-13d8f7ec7e79"}"`
						} `json:"f:ownerReferences"`
					} `json:"f:metadata"`
					FSpec struct {
						FContainers struct {
							KNameNginx struct {
								NAMING_FAILED struct {
								} `json:"."`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FName struct {
								} `json:"f:name"`
								FPorts struct {
									NAMING_FAILED struct {
									} `json:"."`
									KContainerPort80ProtocolTCP struct {
										NAMING_FAILED struct {
										} `json:"."`
										FContainerPort struct {
										} `json:"f:containerPort"`
										FProtocol struct {
										} `json:"f:protocol"`
									} `json:"k:{"containerPort":80,"protocol":"TCP"}"`
								} `json:"f:ports"`
								FResources struct {
								} `json:"f:resources"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
							} `json:"k:{"name":"nginx"}"`
						} `json:"f:containers"`
						FDNSPolicy struct {
						} `json:"f:dnsPolicy"`
						FEnableServiceLinks struct {
						} `json:"f:enableServiceLinks"`
						FRestartPolicy struct {
						} `json:"f:restartPolicy"`
						FSchedulerName struct {
						} `json:"f:schedulerName"`
						FSecurityContext struct {
						} `json:"f:securityContext"`
						FTerminationGracePeriodSeconds struct {
						} `json:"f:terminationGracePeriodSeconds"`
					} `json:"f:spec"`
				} `json:"fieldsV1"`
				Subresource string `json:"subresource,omitempty"`
			} `json:"managedFields"`
		} `json:"metadata"`
		Spec struct {
			Volumes []struct {
				Name      string `json:"name"`
				Projected struct {
					Sources []struct {
						ServiceAccountToken struct {
							ExpirationSeconds int    `json:"expirationSeconds"`
							Path              string `json:"path"`
						} `json:"serviceAccountToken,omitempty"`
						ConfigMap struct {
							Name  string `json:"name"`
							Items []struct {
								Key  string `json:"key"`
								Path string `json:"path"`
							} `json:"items"`
						} `json:"configMap,omitempty"`
						DownwardAPI struct {
							Items []struct {
								Path     string `json:"path"`
								FieldRef struct {
									APIVersion string `json:"apiVersion"`
									FieldPath  string `json:"fieldPath"`
								} `json:"fieldRef"`
							} `json:"items"`
						} `json:"downwardAPI,omitempty"`
					} `json:"sources"`
					DefaultMode int `json:"defaultMode"`
				} `json:"projected"`
			} `json:"volumes"`
			Containers []struct {
				Name  string `json:"name"`
				Image string `json:"image"`
				Ports []struct {
					ContainerPort int    `json:"containerPort"`
					Protocol      string `json:"protocol"`
				} `json:"ports"`
				Resources struct {
				} `json:"resources"`
				VolumeMounts []struct {
					Name      string `json:"name"`
					ReadOnly  bool   `json:"readOnly"`
					MountPath string `json:"mountPath"`
				} `json:"volumeMounts"`
				TerminationMessagePath   string `json:"terminationMessagePath"`
				TerminationMessagePolicy string `json:"terminationMessagePolicy"`
				ImagePullPolicy          string `json:"imagePullPolicy"`
			} `json:"containers"`
			RestartPolicy                 string `json:"restartPolicy"`
			TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
			DNSPolicy                     string `json:"dnsPolicy"`
			ServiceAccountName            string `json:"serviceAccountName"`
			ServiceAccount                string `json:"serviceAccount"`
			NodeName                      string `json:"nodeName"`
			SecurityContext               struct {
			} `json:"securityContext"`
			SchedulerName string `json:"schedulerName"`
			Tolerations   []struct {
				Key               string `json:"key"`
				Operator          string `json:"operator"`
				Effect            string `json:"effect"`
				TolerationSeconds int    `json:"tolerationSeconds"`
			} `json:"tolerations"`
			Priority           int    `json:"priority"`
			EnableServiceLinks bool   `json:"enableServiceLinks"`
			PreemptionPolicy   string `json:"preemptionPolicy"`
		} `json:"spec"`
		Status struct {
			ObservedGeneration int    `json:"observedGeneration"`
			Phase              string `json:"phase"`
			Conditions         []struct {
				Type               string      `json:"type"`
				ObservedGeneration int         `json:"observedGeneration"`
				Status             string      `json:"status"`
				LastProbeTime      interface{} `json:"lastProbeTime"`
				LastTransitionTime time.Time   `json:"lastTransitionTime"`
			} `json:"conditions"`
			HostIP  string `json:"hostIP"`
			HostIPs []struct {
				IP string `json:"ip"`
			} `json:"hostIPs"`
			PodIP  string `json:"podIP"`
			PodIPs []struct {
				IP string `json:"ip"`
			} `json:"podIPs"`
			StartTime         time.Time `json:"startTime"`
			ContainerStatuses []struct {
				Name  string `json:"name"`
				State struct {
					Running struct {
						StartedAt time.Time `json:"startedAt"`
					} `json:"running"`
				} `json:"state"`
				LastState struct {
					Terminated struct {
						ExitCode    int       `json:"exitCode"`
						Reason      string    `json:"reason"`
						StartedAt   time.Time `json:"startedAt"`
						FinishedAt  time.Time `json:"finishedAt"`
						ContainerID string    `json:"containerID"`
					} `json:"terminated"`
				} `json:"lastState"`
				Ready        bool   `json:"ready"`
				RestartCount int    `json:"restartCount"`
				Image        string `json:"image"`
				ImageID      string `json:"imageID"`
				ContainerID  string `json:"containerID"`
				Started      bool   `json:"started"`
				Resources    struct {
				} `json:"resources"`
				VolumeMounts []struct {
					Name              string `json:"name"`
					MountPath         string `json:"mountPath"`
					ReadOnly          bool   `json:"readOnly"`
					RecursiveReadOnly string `json:"recursiveReadOnly"`
				} `json:"volumeMounts"`
			} `json:"containerStatuses"`
			QosClass string `json:"qosClass"`
		} `json:"status"`
	} `json:"items"`
}