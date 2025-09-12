package v1alpha1

import (
	appclassapi "github.com/diktyo-io/appclass-api/pkg/apis/appclass"
	_ "github.com/gogo/protobuf/gogoproto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Constants for AppClass
const (
	// appClassLabel is the default label of the AppClass
	appClassLabel = appclassapi.GroupName
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Namespaced,shortName=ac

// AppClass is a collection of Pods belonging to the same application.
// +protobuf=true
type AppClass struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// AppGroupSpec defines the number of Pods and which Pods belong to the group.
	// +optional
	Spec AppClassSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// AppClassStatus defines the observed use.
	// +optional
	Status AppClassStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// AppClassSpec represents the template of a App Class.
// +protobuf=true
type AppClassSpec struct {
	// +required
	Predicates []Predicate `json:"predicates,omitempty" protobuf:"bytes,1,opt,name=predicates, casttype=Predicate"`

	// +required
	ApplicationClasses []ApplicationClassSpecEntry `json:"applicationClasses,omitempty" protobuf:"bytes,2,rep,name=applicationClasses, casttype=ApplicationClassSpecEntry"`

	// +optional
	GlobalClassification *GlobalSpecClassificationDescriptor `json:"globalClassification,omitempty" protobuf:"bytes,3,opt,name=globalClassification, casttype=AppGroupWorkloadList"`
}

// ApplicationClassSpecEntry defines a class and its app group workloads for spec
// +protobuf=true
type ApplicationClassSpecEntry struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// +optional
	AppGroupWorkloads []AppGroupWorkloadReference `json:"appGroupWorkloads,omitempty" protobuf:"bytes,2,opt,name=appGroupWorkloads"`
}

// AppGroupWorkloadReference defines the app group workloads for spec classification
// +protobuf=true
type AppGroupWorkloadReference struct {
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,1,opt,name=namespace"`

	// +optional
	AppGroup string `json:"appGroup,omitempty" protobuf:"bytes,2,opt,name=appGroup"`

	// +optional
	AppGroupWorkloads []string `json:"appGroupWorkloads,omitempty" protobuf:"bytes,3,opt,name=appGroupWorkloads"`
}

// AppClassStatus defines the observed state of AppClass
// +protobuf=true
type AppClassStatus struct {
	// +optional
	ApplicationClasses []ApplicationClass `json:"applicationClasses,omitempty" protobuf:"bytes,1,opt,name=applicationClasses, casttype=ApplicationClass"`

	// +optional
	GlobalClassification *GlobalClassificationDescriptor `json:"globalClassification,omitempty" protobuf:"bytes,2,opt,name=globalClassification, casttype=globalClassification"`
}

// GlobalSpecClassificationDescriptor holds the spec.globalClassification
type GlobalSpecClassificationDescriptor struct {
	AppGroups []GlobalSpecAppGroupClassification `json:"appGroups,omitempty" protobuf:"bytes,1,opt,name=appGroups, casttype=GlobalSpecAppGroupClassification"`
}

// Predicate defines a scheduling predicate with a name and expression
// +protobuf=true
type Predicate struct {
	// +required
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// +required
	Expression string `json:"expression" protobuf:"bytes,2,opt,name=expression"`
}

// GlobalClassificationDescriptor holds globally classified app groups
// +protobuf=true
type GlobalClassificationDescriptor struct {
	AppGroups []GlobalAppGroupClassification `json:"appGroups,omitempty"`
}

// GlobalAppGroupClassification defines a globally classified app group
// +protobuf=true
type GlobalAppGroupClassification struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// +optional
	AppInfo string `json:"appInfo,omitempty" protobuf:"bytes,2,opt,name=appInfo"`

	// +optional
	Class string `json:"class,omitempty" protobuf:"bytes,3,opt,name=class"`

	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,4,opt,name=namespace"`

	// +optional
	AppGroup string `json:"appGroup,omitempty" protobuf:"bytes,5,opt,name=appGroup"`
}

// GlobalSpecAppGroupClassification defines authoritative classification in spec
// +protobuf=true
type GlobalSpecAppGroupClassification struct {
	// +optional
	Class string `json:"class,omitempty" protobuf:"bytes,1,opt,name=class"`

	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`

	// +optional
	AppGroup string `json:"appGroup,omitempty" protobuf:"bytes,3,opt,name=appGroup"`
}

// AppInfoReference holds reference metadata to an AppInfo
// +protobuf=true
type AppInfoReference struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`

	// +optional
	AppGroup string `json:"appGroup,omitempty" protobuf:"bytes,3,opt,name=appGroup"`

	// +optional
	AppGroupWorkloads []string `json:"appGroupWorkloads,omitempty" protobuf:"bytes,4,opt,name=appGroupWorkloads"`
}

// ApplicationClass represents a classification group with its app infos
// +protobuf=true
type ApplicationClass struct {
	Name     string             `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	AppInfos []AppInfoReference `json:"appInfos,omitempty" protobuf:"bytes,1,opt,name=appInfos, casttype=appInfos"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppClassList is a collection of app classes.
type AppClassList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the list of AppClass
	Items []AppClass `json:"items"`
}
