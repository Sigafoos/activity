package streams

import (
	"context"
	"errors"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// TypePredicatedResolver resolves ActivityStreams values if the value satisfies a
// predicate condition based on its type.
type TypePredicatedResolver struct {
	delegate  Resolver
	predicate interface{}
}

// NewTypePredicatedResolver creates a new Resolver that applies a predicate to an
// ActivityStreams value to determine whether to Resolve or not. The
// ActivityStreams value's type is examined to determine if the predicate can
// apply itself to the value. This guarantees the predicate will receive a
// concrete value whose underlying ActivityStreams type matches the concrete
// interface name. The predicate function must be of the form:
//
//   func(context.Context, <TypeInterface>) (bool, error)
//
// where TypeInterface is the code-generated interface for an ActivityStreams
// type. An error is returned if the predicate does not match this signature.
func NewTypePredicatedResolver(delegate Resolver, predicate interface{}) (*TypePredicatedResolver, error) {
	// The predicate must satisfy one known predicate function signature, or else we will generate a runtime error instead of silently fail.
	switch predicate.(type) {
	case func(context.Context, vocab.ActivityStreamsAccept) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsActivity) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsAdd) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsAnnounce) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsApplication) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsArrive) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsArticle) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsAudio) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsBlock) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsCollection) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsCollectionPage) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsCreate) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsDelete) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsDislike) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsDocument) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsEvent) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsFlag) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsFollow) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsGroup) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsIgnore) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsImage) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsIntransitiveActivity) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsInvite) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsJoin) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsLeave) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsLike) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsLink) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsListen) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsMention) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsMove) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsNote) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsObject) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsOffer) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsOrderedCollection) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsOrderedCollectionPage) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsOrganization) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsPage) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsPerson) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsPlace) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsProfile) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.W3IDSecurityV1PublicKey) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsQuestion) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsRead) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsReject) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsRelationship) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsRemove) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsService) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsTentativeAccept) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsTentativeReject) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsTombstone) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsTravel) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsUndo) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsUpdate) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsVideo) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsView) (bool, error):
		// Do nothing, this predicate has a correct signature.
	default:
		return nil, errors.New("the predicate function is of the wrong signature and would never be called")
	}
	return &TypePredicatedResolver{
		delegate:  delegate,
		predicate: predicate,
	}, nil
}

// Apply uses a predicate to determine whether to resolve the ActivityStreams
// value. The predicate's signature is matched with the ActivityStreams
// value's type. This strictly assures that the predicate will only be passed
// ActivityStream objects whose type matches its interface. Returns an error
// if the ActivityStreams type does not match the predicate, is not a type
// handled by the generated code, or the resolver returns an error. Returns
// true if the predicate returned true.
func (this TypePredicatedResolver) Apply(ctx context.Context, o ActivityStreamsInterface) (bool, error) {
	var predicatePasses bool
	var err error
	if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Accept" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsAccept) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsAccept); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Activity" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsActivity) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsActivity); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Add" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsAdd) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsAdd); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Announce" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsAnnounce) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsAnnounce); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Application" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsApplication) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsApplication); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Arrive" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsArrive) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsArrive); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Article" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsArticle) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsArticle); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Audio" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsAudio) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsAudio); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Block" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsBlock) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsBlock); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Collection" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsCollection) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsCollection); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "CollectionPage" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsCollectionPage) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsCollectionPage); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Create" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsCreate) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsCreate); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Delete" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsDelete) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsDelete); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Dislike" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsDislike) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsDislike); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Document" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsDocument) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsDocument); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Event" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsEvent) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsEvent); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Flag" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsFlag) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsFlag); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Follow" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsFollow) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsFollow); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Group" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsGroup) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsGroup); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Ignore" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsIgnore) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsIgnore); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Image" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsImage) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsImage); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "IntransitiveActivity" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsIntransitiveActivity) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsIntransitiveActivity); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Invite" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsInvite) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsInvite); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Join" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsJoin) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsJoin); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Leave" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsLeave) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsLeave); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Like" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsLike) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsLike); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Link" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsLink) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsLink); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Listen" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsListen) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsListen); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Mention" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsMention) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsMention); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Move" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsMove) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsMove); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Note" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsNote) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsNote); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Object" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsObject) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsObject); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Offer" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsOffer) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsOffer); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "OrderedCollection" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsOrderedCollection) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsOrderedCollection); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "OrderedCollectionPage" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsOrderedCollectionPage) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsOrderedCollectionPage); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Organization" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsOrganization) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsOrganization); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Page" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsPage) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsPage); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Person" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsPerson) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsPerson); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Place" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsPlace) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsPlace); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Profile" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsProfile) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsProfile); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://w3id.org/security/v1" && o.GetTypeName() == "PublicKey" {
		if fn, ok := this.predicate.(func(context.Context, vocab.W3IDSecurityV1PublicKey) (bool, error)); ok {
			if v, ok := o.(vocab.W3IDSecurityV1PublicKey); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Question" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsQuestion) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsQuestion); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Read" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsRead) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsRead); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Reject" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsReject) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsReject); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Relationship" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsRelationship) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsRelationship); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Remove" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsRemove) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsRemove); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Service" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsService) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsService); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "TentativeAccept" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsTentativeAccept) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsTentativeAccept); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "TentativeReject" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsTentativeReject) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsTentativeReject); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Tombstone" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsTombstone) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsTombstone); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Travel" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsTravel) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsTravel); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Undo" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsUndo) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsUndo); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Update" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsUpdate) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsUpdate); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "Video" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsVideo) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsVideo); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/ns/activitystreams" && o.GetTypeName() == "View" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsView) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsView); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else {
		return false, ErrUnhandledType
	}
	if err != nil {
		return predicatePasses, err
	}
	if predicatePasses {
		return true, this.delegate.Resolve(ctx, o)
	} else {
		return false, nil
	}
}
