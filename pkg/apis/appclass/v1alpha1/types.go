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
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,shortName=ac

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
	// Type defines the classification strategy. One of: Property, Explicit, Predicate.
	// +required
	Type string `json:"type" protobuf:"bytes,1,opt,name=type"`

	// PropertyDefinition holds the spec for Property-based classification.
	// +optional
	PropertyDefinition *PropertyDefinition `json:"propertyDefinition,omitempty" protobuf:"bytes,2,opt,name=propertyDefinition"`

	// ExplicitDefinition holds the spec for Explicit classification.
	// +optional
	ExplicitDefinition *ExplicitDefinition `json:"explicitDefinition,omitempty" protobuf:"bytes,3,opt,name=explicitDefinition"`

	// Predicates defines the scheduling predicates used for Predicate-based classification.
	// +optional
	Predicates []Predicate `json:"predicates,omitempty" protobuf:"bytes,4,rep,name=predicates,casttype=Predicate"`

	// ApplicationClasses defines the named classes and their app group workloads.
	// +optional
	ApplicationClasses []ApplicationClassSpecEntry `json:"applicationClasses,omitempty" protobuf:"bytes,5,rep,name=applicationClasses,casttype=ApplicationClassSpecEntry"`

	// GlobalClassification optionally defines authoritative class assignments per app group.
	// +optional
	GlobalClassification *GlobalSpecClassificationDescriptor `json:"globalClassification,omitempty" protobuf:"bytes,6,opt,name=globalClassification"`
}

// PropertyDefinition holds the spec for Property-based classification.
// +protobuf=true
type PropertyDefinition struct {
	// AppInfoProperties lists the AppInfo property names used for classification.
	// +optional
	AppInfoProperties []string `json:"appInfoProperties,omitempty" protobuf:"bytes,1,rep,name=appInfoProperties"`
}

// ExplicitDefinition holds the spec for Explicit classification.
// +protobuf=true
type ExplicitDefinition struct {
	// Other specifies the default class name for unmatched app groups.
	// +required
	Other string `json:"other" protobuf:"bytes,1,opt,name=other"`

	// ClassDefs defines the list of class definitions with their matching clauses.
	// +required
	ClassDefs []ClassDef `json:"classDefs" protobuf:"bytes,2,rep,name=classDefs,casttype=ClassDef"`
}

// ClassDef defines a class and the conditions under which it is assigned.
// +protobuf=true
type ClassDef struct {
	// Class is the name of the class being defined.
	// +required
	Class string `json:"class" protobuf:"bytes,1,opt,name=class"`

	// Clauses is a list of OR-ed clauses; an app group matches this class if any clause matches.
	// +required
	Clauses []ExplicitClause `json:"clauses" protobuf:"bytes,2,rep,name=clauses,casttype=ExplicitClause"`
}

// ExplicitClause is a list of AND-ed conditions.
// +protobuf=true
type ExplicitClause struct {
	// Conditions is the list of conditions that must all be true for this clause to match.
	// +required
	Conditions []ExplicitCondition `json:"conditions" protobuf:"bytes,1,rep,name=conditions,casttype=ExplicitCondition"`
}

// ExplicitCondition is a single condition comparing a property to a value or another property.
// +protobuf=true
type ExplicitCondition struct {
	// Property is the AppInfo property name to evaluate.
	// +required
	Property string `json:"property" protobuf:"bytes,1,opt,name=property"`

	// Operator is the comparison operator. One of: =, !=, <, <=, >, >=
	// +required
	Operator string `json:"operator" protobuf:"bytes,2,opt,name=operator"`

	// Value is the literal value to compare against (mutually exclusive with OtherProperty).
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,3,opt,name=value"`

	// OtherProperty is the other AppInfo property to compare against (mutually exclusive with Value).
	// +optional
	OtherProperty string `json:"otherProperty,omitempty" protobuf:"bytes,4,opt,name=otherProperty"`
}

// ApplicationClassSpecEntry defines a class and its app group workloads for spec
// +protobuf=true
type ApplicationClassSpecEntry struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// InGroup lists class names whose members are considered inside this class.
	// +optional
	InGroup []string `json:"inGroup,omitempty" protobuf:"bytes,2,rep,name=inGroup"`

	// OutGroup lists class names whose members are considered outside this class.
	// +optional
	OutGroup []string `json:"outGroup,omitempty" protobuf:"bytes,3,rep,name=outGroup"`

	// +optional
	AppGroupWorkloads []AppGroupWorkloadReference `json:"appGroupWorkloads,omitempty" protobuf:"bytes,4,rep,name=appGroupWorkloads"`
}

// AppGroupWorkloadReference defines the app group workloads for spec classification
// +protobuf=true
type AppGroupWorkloadReference struct {
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,1,opt,name=namespace"`

	// +optional
	AppGroup string `json:"appGroup,omitempty" protobuf:"bytes,2,opt,name=appGroup"`

	// +optional
	AppGroupWorkloads []string `json:"appGroupWorkloads,omitempty" protobuf:"bytes,3,rep,name=appGroupWorkloads"`
}

// AppClassStatus defines the observed state of AppClass
// +protobuf=true
type AppClassStatus struct {
	// +optional
	ApplicationClasses []ApplicationClass `json:"applicationClasses,omitempty" protobuf:"bytes,1,rep,name=applicationClasses,casttype=ApplicationClass"`

	// +optional
	GlobalClassification *GlobalClassificationDescriptor `json:"globalClassification,omitempty" protobuf:"bytes,2,opt,name=globalClassification"`
}

// GlobalSpecClassificationDescriptor holds the spec.globalClassification
// +protobuf=true
type GlobalSpecClassificationDescriptor struct {
	AppGroups []GlobalSpecAppGroupClassification `json:"appGroups,omitempty" protobuf:"bytes,1,rep,name=appGroups,casttype=GlobalSpecAppGroupClassification"`
}

// Predicate defines a scheduling predicate with a name and structured clauses.
// +protobuf=true
type Predicate struct {
	// +required
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// Clauses is a list of OR-ed clauses; the predicate matches if any clause matches.
	// +required
	Clauses []Clause `json:"clauses" protobuf:"bytes,2,rep,name=clauses,casttype=Clause"`
}

// Clause is a list of AND-ed condition nodes.
// +protobuf=true
type Clause struct {
	// AndClauses is the list of condition nodes that must all be true.
	// +required
	AndClauses []CondNode `json:"andClauses" protobuf:"bytes,1,rep,name=andClauses,casttype=CondNode"`
}

// CondNode is a single condition node, one of: left, right, or rel.
// +protobuf=true
type CondNode struct {
	// Left is a condition applied to the left-hand AppInfo.
	// +optional
	Left *Condition `json:"left,omitempty" protobuf:"bytes,1,opt,name=left"`

	// Right is a condition applied to the right-hand AppInfo.
	// +optional
	Right *Condition `json:"right,omitempty" protobuf:"bytes,2,opt,name=right"`

	// Rel is a relational condition comparing a property across two AppInfos.
	// +optional
	Rel *RelCondition `json:"rel,omitempty" protobuf:"bytes,3,opt,name=rel"`
}

// Condition is a single-AppInfo condition, comparing a property to a value or a relational condition.
// +protobuf=true
type Condition struct {
	// Property is the AppInfo property name.
	// +optional
	Property string `json:"property,omitempty" protobuf:"bytes,1,opt,name=property"`

	// Operator is the comparison operator. One of: =, !=, <, <=, >, >=
	// +optional
	Operator string `json:"operator,omitempty" protobuf:"bytes,2,opt,name=operator"`

	// Value is the literal value to compare against.
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,3,opt,name=value"`

	// Rel is an embedded relational condition.
	// +optional
	Rel *RelCondition `json:"rel,omitempty" protobuf:"bytes,4,opt,name=rel"`
}

// RelCondition compares a property on one AppInfo against a property on another.
// +protobuf=true
type RelCondition struct {
	// Property is the property name on the first AppInfo.
	// +required
	Property string `json:"property" protobuf:"bytes,1,opt,name=property"`

	// Operator is the comparison operator. One of: =, !=, <, <=, >, >=
	// +required
	Operator string `json:"operator" protobuf:"bytes,2,opt,name=operator"`

	// OtherProperty is the property name on the second AppInfo.
	// +required
	OtherProperty string `json:"otherProperty" protobuf:"bytes,3,opt,name=otherProperty"`
}

// GlobalClassificationDescriptor holds globally classified app groups
// +protobuf=true
type GlobalClassificationDescriptor struct {
	AppGroups []GlobalAppGroupClassification `json:"appGroups,omitempty" protobuf:"bytes,1,rep,name=appGroups"`
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
	AppGroupWorkloads []string `json:"appGroupWorkloads,omitempty" protobuf:"bytes,4,rep,name=appGroupWorkloads"`
}

// ApplicationClass represents a classification group with its app infos
// +protobuf=true
type ApplicationClass struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// InGroup lists class names whose members are considered inside this class.
	// +optional
	InGroup []string `json:"inGroup,omitempty" protobuf:"bytes,2,rep,name=inGroup"`

	// OutGroup lists class names whose members are considered outside this class.
	// +optional
	OutGroup []string `json:"outGroup,omitempty" protobuf:"bytes,3,rep,name=outGroup"`

	// +optional
	AppInfos []AppInfoReference `json:"appInfos,omitempty" protobuf:"bytes,4,rep,name=appInfos,casttype=AppInfoReference"`
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
