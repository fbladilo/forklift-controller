package plan

import (
	libref "github.com/konveyor/controller/pkg/ref"
	api "github.com/konveyor/virt-controller/pkg/apis/virt/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type PlanPredicate struct {
	predicate.Funcs
}

func (r PlanPredicate) Create(e event.CreateEvent) bool {
	_, cast := e.Object.(*api.Plan)
	if cast {
		libref.Mapper.Create(e)
		return true
	}

	return false
}

func (r PlanPredicate) Update(e event.UpdateEvent) bool {
	object, cast := e.ObjectNew.(*api.Plan)
	if !cast {
		return false
	}
	changed := object.Status.ObservedGeneration < object.Generation
	if changed {
		libref.Mapper.Update(e)
	}

	return changed
}

func (r PlanPredicate) Delete(e event.DeleteEvent) bool {
	_, cast := e.Object.(*api.Plan)
	if cast {
		libref.Mapper.Delete(e)
		return true
	}

	return false
}

type ProviderPredicate struct {
	predicate.Funcs
}

func (r ProviderPredicate) Create(e event.CreateEvent) bool {
	p, cast := e.Object.(*api.Provider)
	if cast {
		reconciled := p.Status.ObservedGeneration == p.Generation
		return reconciled
	}

	return false
}

func (r ProviderPredicate) Update(e event.UpdateEvent) bool {
	p, cast := e.ObjectNew.(*api.Provider)
	if cast {
		reconciled := p.Status.ObservedGeneration == p.Generation
		return reconciled
	}

	return false
}

func (r ProviderPredicate) Delete(e event.DeleteEvent) bool {
	p, cast := e.Object.(*api.Provider)
	if cast {
		reconciled := p.Status.ObservedGeneration == p.Generation
		return reconciled
	}

	return false
}

func (r ProviderPredicate) Generic(e event.GenericEvent) bool {
	p, cast := e.Object.(*api.Provider)
	if cast {
		reconciled := p.Status.ObservedGeneration == p.Generation
		return reconciled
	}

	return false
}