package common

type SensorType string

const (
	SensorTypeSidecar SensorType = "falcon-container"
	SensorTypeKac     SensorType = "falcon-kac"
	SensorTypeNode    SensorType = "falcon-sensor"
)

const (
	FalconContainerInjection               = "sensor.falcon-system.crowdstrike.com/injection"
	FalconContainerInjectorTLSName         = "injector-tls"
	FalconHostInstallDir                   = "/opt"
	FalconInitHostInstallDir               = "/host_opt"
	FalconDataDir                          = "/opt/CrowdStrike"
	FalconInitDataDir                      = "/host_opt/CrowdStrike/"
	FalconStoreFile                        = "/opt/CrowdStrike/falconstore"
	FalconInitStoreFile                    = "/host_opt/CrowdStrike/falconstore"
	FalconDaemonsetInitBinary              = "/opt/CrowdStrike/falcon-daemonset-init -i"
	FalconDaemonsetInitBinaryInvocation    = "falcon-daemonset-init -i"
	FalconDaemonsetCleanupBinaryInvocation = "falcon-daemonset-init -u"
	FalconDaemonsetCleanupBinary           = "/opt/CrowdStrike/falcon-daemonset-init -u"
	FalconContainerProbePath               = "/live"
	FalconAdmissionClientStartupProbePath  = "/startz"
	FalconAdmissionClientLivenessProbePath = "/livez"
	FalconAdmissionStartupProbePath        = "/startz-kac"
	FalconAdmissionLivenessProbePath       = "/livez-kac"
	FalconAdmissionServiceHTTPSName        = "webhook-port"
	FalconServiceHTTPSName                 = "https"
	FalconServiceHTTPSPort                 = 443

	FalconInstanceNameKey    = "crowdstrike.com/name"
	FalconInstanceKey        = "crowdstrike.com/instance"
	FalconComponentKey       = "crowdstrike.com/component"
	FalconManagedByKey       = "crowdstrike.com/managed-by"
	FalconPartOfKey          = "crowdstrike.com/part-of"
	FalconProviderKey        = "crowdstrike.com/provider"
	FalconCreatedKey         = "crowdstrike.com/created-by"
	FalconAdmissionReviewKey = "falcon.crowdstrike.com/admission-review"

	FalconKernelSensor        = "kernel_sensor"
	FalconSidecarSensor       = "container_sensor"
	FalconAdmissionController = "admission_controller"
	FalconFinalizer           = "falcon.crowdstrike.com/finalizer"
	FalconProviderValue       = "crowdstrike"
	FalconPartOfValue         = "Falcon"
	FalconCreatedValue        = "falcon-operator"
	FalconManagedByValue      = "controller-manager"
	FalconPriorityClassName   = "system-cluster-critical"

	SidecarServiceAccountName   = "falcon-operator-sidecar-sensor"
	FalconPullSecretName        = "crowdstrike-falcon-pull-secret"
	NodeServiceAccountName      = "falcon-operator-node-sensor"
	AdmissionServiceAccountName = "falcon-operator-admission-controller"
	NodeClusterRoleBindingName  = "crowdstrike-falcon-node-sensor"
)
