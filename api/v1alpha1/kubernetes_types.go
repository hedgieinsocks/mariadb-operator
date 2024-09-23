// nolint:lll
package v1alpha1

import (
	kadapter "github.com/mariadb-operator/mariadb-operator/pkg/kubernetes/adapter"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#volume-v1-core.
type VolumeSource struct {
	// +optional
	EmptyDir *corev1.EmptyDirVolumeSource `json:"emptyDir,omitempty" protobuf:"bytes,2,opt,name=emptyDir"`
	// +optional
	NFS *corev1.NFSVolumeSource `json:"nfs,omitempty" protobuf:"bytes,7,opt,name=nfs"`
	// +optional
	CSI *corev1.CSIVolumeSource `json:"csi,omitempty" protobuf:"bytes,28,opt,name=csi"`
	// +optional
	PersistentVolumeClaim *corev1.PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty" protobuf:"bytes,10,opt,name=persistentVolumeClaim"`
}

func VolumeSourceFromKubernetesType(kv corev1.VolumeSource) VolumeSource {
	return VolumeSource{
		EmptyDir:              kv.EmptyDir,
		NFS:                   kv.NFS,
		CSI:                   kv.CSI,
		PersistentVolumeClaim: kv.PersistentVolumeClaim,
	}
}

func (v VolumeSource) ToKubernetesType() corev1.VolumeSource {
	return corev1.VolumeSource{
		EmptyDir:              v.EmptyDir,
		NFS:                   v.NFS,
		CSI:                   v.CSI,
		PersistentVolumeClaim: v.PersistentVolumeClaim,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#volume-v1-core.
type Volume struct {
	Name         string `json:"name" protobuf:"bytes,1,opt,name=name"`
	VolumeSource `json:",inline" protobuf:"bytes,2,opt,name=volumeSource"`
}

func (v Volume) ToKubernetesType() corev1.Volume {
	return corev1.Volume{
		Name:         v.Name,
		VolumeSource: v.VolumeSource.ToKubernetesType(),
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#volumemount-v1-core.
type VolumeMount struct {
	// This must match the Name of a Volume.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// +optional
	ReadOnly  bool   `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
	MountPath string `json:"mountPath" protobuf:"bytes,3,opt,name=mountPath"`
	// +optional
	SubPath string `json:"subPath,omitempty" protobuf:"bytes,4,opt,name=subPath"`
}

func (v VolumeMount) ToKubernetesType() corev1.VolumeMount {
	return corev1.VolumeMount{
		Name:      v.Name,
		ReadOnly:  v.ReadOnly,
		MountPath: v.MountPath,
		SubPath:   v.SubPath,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#persistentvolumeclaimspec-v1-core.
type PersistentVolumeClaimSpec struct {
	// +optional
	// +listType=atomic
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes,omitempty" protobuf:"bytes,1,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,4,opt,name=selector"`
	// +optional
	Resources corev1.VolumeResourceRequirements `json:"resources,omitempty" protobuf:"bytes,2,opt,name=resources"`
	// +optional
	StorageClassName *string `json:"storageClassName,omitempty" protobuf:"bytes,3,opt,name=storageClassName"`
}

func (p PersistentVolumeClaimSpec) ToKubernetesType() corev1.PersistentVolumeClaimSpec {
	return corev1.PersistentVolumeClaimSpec{
		AccessModes:      p.AccessModes,
		Selector:         p.Selector,
		Resources:        p.Resources,
		StorageClassName: p.StorageClassName,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.
type PodAffinityTerm struct {
	// +optional
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,1,opt,name=labelSelector"`
	TopologyKey   string                `json:"topologyKey" protobuf:"bytes,2,opt,name=topologyKey"`
}

func (p PodAffinityTerm) ToKubernetesType() corev1.PodAffinityTerm {
	return corev1.PodAffinityTerm{
		LabelSelector: p.LabelSelector,
		TopologyKey:   p.TopologyKey,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#weightedpodaffinityterm-v1-core.
type WeightedPodAffinityTerm struct {
	Weight          int32           `json:"weight" protobuf:"varint,1,opt,name=weight"`
	PodAffinityTerm PodAffinityTerm `json:"podAffinityTerm" protobuf:"bytes,2,opt,name=podAffinityTerm"`
}

func (p WeightedPodAffinityTerm) ToKubernetesType() corev1.WeightedPodAffinityTerm {
	return corev1.WeightedPodAffinityTerm{
		Weight:          p.Weight,
		PodAffinityTerm: p.PodAffinityTerm.ToKubernetesType(),
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.
type PodAntiAffinity struct {
	// +optional
	// +listType=atomic
	RequiredDuringSchedulingIgnoredDuringExecution []PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,rep,name=requiredDuringSchedulingIgnoredDuringExecution"`
	// +optional
	// +listType=atomic
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution"`
}

func (p PodAntiAffinity) ToKubernetesType() corev1.PodAntiAffinity {
	return corev1.PodAntiAffinity{
		RequiredDuringSchedulingIgnoredDuringExecution:  kadapter.ToKubernetesSlice(p.RequiredDuringSchedulingIgnoredDuringExecution),
		PreferredDuringSchedulingIgnoredDuringExecution: kadapter.ToKubernetesSlice(p.PreferredDuringSchedulingIgnoredDuringExecution),
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#affinity-v1-core.
type Affinity struct {
	// +optional
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty" protobuf:"bytes,1,opt,name=podAntiAffinity"`
}

func (a Affinity) ToKubernetesType() corev1.Affinity {
	var affinity corev1.Affinity
	if a.PodAntiAffinity != nil {
		affinity.PodAntiAffinity = ptr.To(a.PodAntiAffinity.ToKubernetesType())
	}
	return affinity
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#topologyspreadconstraint-v1-core.
type TopologySpreadConstraint struct {
	MaxSkew           int32                                `json:"maxSkew" protobuf:"varint,1,opt,name=maxSkew"`
	TopologyKey       string                               `json:"topologyKey" protobuf:"bytes,2,opt,name=topologyKey"`
	WhenUnsatisfiable corev1.UnsatisfiableConstraintAction `json:"whenUnsatisfiable" protobuf:"bytes,3,opt,name=whenUnsatisfiable,casttype=UnsatisfiableConstraintAction"`
	// +optional
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,4,opt,name=labelSelector"`
	// +optional
	MinDomains *int32 `json:"minDomains,omitempty" protobuf:"varint,5,opt,name=minDomains"`
	// +optional
	NodeAffinityPolicy *corev1.NodeInclusionPolicy `json:"nodeAffinityPolicy,omitempty" protobuf:"bytes,6,opt,name=nodeAffinityPolicy"`
	// +optional
	NodeTaintsPolicy *corev1.NodeInclusionPolicy `json:"nodeTaintsPolicy,omitempty" protobuf:"bytes,7,opt,name=nodeTaintsPolicy"`
	// +optional
	MatchLabelKeys []string `json:"matchLabelKeys,omitempty" protobuf:"bytes,8,opt,name=matchLabelKeys"`
}

func (t TopologySpreadConstraint) ToKubernetesType() corev1.TopologySpreadConstraint {
	return corev1.TopologySpreadConstraint{
		MaxSkew:            t.MaxSkew,
		TopologyKey:        t.TopologyKey,
		WhenUnsatisfiable:  t.WhenUnsatisfiable,
		LabelSelector:      t.LabelSelector,
		MinDomains:         t.MinDomains,
		NodeAffinityPolicy: t.NodeAffinityPolicy,
		NodeTaintsPolicy:   t.NodeTaintsPolicy,
		MatchLabelKeys:     t.MatchLabelKeys,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.
type LocalObjectReference struct {
	// +optional
	// +default=""
	// +kubebuilder:default=""
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

func (r LocalObjectReference) ToKubernetesType() corev1.LocalObjectReference {
	return corev1.LocalObjectReference{
		Name: r.Name,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectreference-v1-core.
type ObjectReference struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
}

func (r ObjectReference) ToKubernetesType() corev1.ObjectReference {
	return corev1.ObjectReference{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.
// +structType=atomic
type SecretKeySelector struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Key                  string `json:"key" protobuf:"bytes,2,opt,name=key"`
}

func (s SecretKeySelector) ToKubernetesType() corev1.SecretKeySelector {
	return corev1.SecretKeySelector{
		LocalObjectReference: s.LocalObjectReference.ToKubernetesType(),
		Key:                  s.Key,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.
// +structType=atomic
type ConfigMapKeySelector struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Key                  string `json:"key" protobuf:"bytes,2,opt,name=key"`
}

func (s ConfigMapKeySelector) ToKubernetesType() corev1.ConfigMapKeySelector {
	return corev1.ConfigMapKeySelector{
		LocalObjectReference: s.LocalObjectReference.ToKubernetesType(),
		Key:                  s.Key,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.
// +structType=atomic
type ObjectFieldSelector struct {
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,1,opt,name=apiVersion"`
	FieldPath  string `json:"fieldPath" protobuf:"bytes,2,opt,name=fieldPath"`
}

func (s ObjectFieldSelector) ToKubernetesType() corev1.ObjectFieldSelector {
	return corev1.ObjectFieldSelector{
		APIVersion: s.APIVersion,
		FieldPath:  s.FieldPath,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.
type EnvVarSource struct {
	// +optional
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" protobuf:"bytes,1,opt,name=fieldRef"`
	// +optional
	ConfigMapKeyRef *ConfigMapKeySelector `json:"configMapKeyRef,omitempty" protobuf:"bytes,3,opt,name=configMapKeyRef"`
	// +optional
	SecretKeyRef *SecretKeySelector `json:"secretKeyRef,omitempty" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

func (e EnvVarSource) ToKubernetesType() corev1.EnvVarSource {
	var env corev1.EnvVarSource
	if e.FieldRef != nil {
		env.FieldRef = ptr.To(e.FieldRef.ToKubernetesType())
	}
	if e.ConfigMapKeyRef != nil {
		env.ConfigMapKeyRef = ptr.To(e.ConfigMapKeyRef.ToKubernetesType())
	}
	if e.SecretKeyRef != nil {
		env.SecretKeyRef = ptr.To(e.SecretKeyRef.ToKubernetesType())
	}
	return env
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.
type EnvVar struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
	// +optional
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty" protobuf:"bytes,3,opt,name=valueFrom"`
}

func (e EnvVar) ToKubernetesType() corev1.EnvVar {
	env := corev1.EnvVar{
		Name:  e.Name,
		Value: e.Value,
	}
	if e.ValueFrom != nil {
		env.ValueFrom = ptr.To(e.ValueFrom.ToKubernetesType())
	}
	return env
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envfromsource-v1-core.
type EnvFromSource struct {
	// +optional
	Prefix string `json:"prefix,omitempty" protobuf:"bytes,1,opt,name=prefix"`
	// +optional
	ConfigMapRef *LocalObjectReference `json:"configMapRef,omitempty" protobuf:"bytes,2,opt,name=configMapRef"`
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef"`
}

func (e EnvFromSource) ToKubernetesType() corev1.EnvFromSource {
	env := corev1.EnvFromSource{
		Prefix: e.Prefix,
	}
	if e.ConfigMapRef != nil {
		env.ConfigMapRef = ptr.To(corev1.ConfigMapEnvSource{
			LocalObjectReference: e.ConfigMapRef.ToKubernetesType(),
		})
	}
	if e.SecretRef != nil {
		env.SecretRef = ptr.To(corev1.SecretEnvSource{
			LocalObjectReference: e.SecretRef.ToKubernetesType(),
		})
	}
	return env
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.
type ExecAction struct {
	// +optional
	// +listType=atomic
	Command []string `json:"command,omitempty" protobuf:"bytes,1,rep,name=command"`
}

func (e ExecAction) ToKubernetesType() corev1.ExecAction {
	return corev1.ExecAction{
		Command: e.Command,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.
type HTTPGetAction struct {
	// +optional
	Path string             `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	Port intstr.IntOrString `json:"port" protobuf:"bytes,2,opt,name=port"`
	// +optional
	Host string `json:"host,omitempty" protobuf:"bytes,3,opt,name=host"`
	// +optional
	Scheme corev1.URIScheme `json:"scheme,omitempty" protobuf:"bytes,4,opt,name=scheme,casttype=URIScheme"`
}

func (e HTTPGetAction) ToKubernetesType() corev1.HTTPGetAction {
	return corev1.HTTPGetAction{
		Path:   e.Path,
		Port:   e.Port,
		Host:   e.Host,
		Scheme: e.Scheme,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#probe-v1-core.
type ProbeHandler struct {
	// +optional
	Exec *ExecAction `json:"exec,omitempty" protobuf:"bytes,1,opt,name=exec"`
	// +optional
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty" protobuf:"bytes,2,opt,name=httpGet"`
}

func (p ProbeHandler) ToKubernetesType() corev1.ProbeHandler {
	var probe corev1.ProbeHandler
	if p.Exec != nil {
		probe.Exec = ptr.To(p.Exec.ToKubernetesType())
	}
	if p.HTTPGet != nil {
		probe.HTTPGet = ptr.To(p.HTTPGet.ToKubernetesType())
	}
	return probe
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#probe-v1-core.
type Probe struct {
	ProbeHandler `json:",inline" protobuf:"bytes,1,opt,name=handler"`
	// +optional
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty" protobuf:"varint,2,opt,name=initialDelaySeconds"`
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,3,opt,name=timeoutSeconds"`
	// +optional
	PeriodSeconds int32 `json:"periodSeconds,omitempty" protobuf:"varint,4,opt,name=periodSeconds"`
	// +optional
	SuccessThreshold int32 `json:"successThreshold,omitempty" protobuf:"varint,5,opt,name=successThreshold"`
	// +optional
	FailureThreshold int32 `json:"failureThreshold,omitempty" protobuf:"varint,6,opt,name=failureThreshold"`
}

func (p Probe) ToKubernetesType() corev1.Probe {
	return corev1.Probe{
		ProbeHandler:        p.ProbeHandler.ToKubernetesType(),
		InitialDelaySeconds: p.InitialDelaySeconds,
		TimeoutSeconds:      p.TimeoutSeconds,
		PeriodSeconds:       p.PeriodSeconds,
		SuccessThreshold:    p.SuccessThreshold,
		FailureThreshold:    p.FailureThreshold,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#resourcerequirements-v1-core.
type ResourceRequirements struct {
	// +optional
	Limits corev1.ResourceList `json:"limits,omitempty" protobuf:"bytes,1,rep,name=limits,casttype=ResourceList,castkey=ResourceName"`
	// +optional
	Requests corev1.ResourceList `json:"requests,omitempty" protobuf:"bytes,2,rep,name=requests,casttype=ResourceList,castkey=ResourceName"`
}

func (r ResourceRequirements) ToKubernetesType() corev1.ResourceRequirements {
	return corev1.ResourceRequirements{
		Limits:   r.Limits,
		Requests: r.Requests,
	}
}

// Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#securitycontext-v1-core.
type SecurityContext struct {
	// +optional
	Capabilities *corev1.Capabilities `json:"capabilities,omitempty" protobuf:"bytes,1,opt,name=capabilities"`
	// +optional
	Privileged *bool `json:"privileged,omitempty" protobuf:"varint,2,opt,name=privileged"`
	// +optional
	RunAsUser *int64 `json:"runAsUser,omitempty" protobuf:"varint,4,opt,name=runAsUser"`
	// +optional
	RunAsGroup *int64 `json:"runAsGroup,omitempty" protobuf:"varint,8,opt,name=runAsGroup"`
	// +optional
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty" protobuf:"varint,5,opt,name=runAsNonRoot"`
	// +optional
	ReadOnlyRootFilesystem *bool `json:"readOnlyRootFilesystem,omitempty" protobuf:"varint,6,opt,name=readOnlyRootFilesystem"`
	// +optional
	AllowPrivilegeEscalation *bool `json:"allowPrivilegeEscalation,omitempty" protobuf:"varint,7,opt,name=allowPrivilegeEscalation"`
}

func (r SecurityContext) ToKubernetesType() corev1.SecurityContext {
	return corev1.SecurityContext{
		Capabilities:             r.Capabilities,
		Privileged:               r.Privileged,
		RunAsUser:                r.RunAsUser,
		RunAsGroup:               r.RunAsGroup,
		RunAsNonRoot:             r.RunAsNonRoot,
		ReadOnlyRootFilesystem:   r.ReadOnlyRootFilesystem,
		AllowPrivilegeEscalation: r.AllowPrivilegeEscalation,
	}
}