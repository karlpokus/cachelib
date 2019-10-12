package cachelib

import "testing"

func TestCache(t *testing.T) {
	cache, err := New("1s")
	if err != nil {
		t.Errorf("expected err to be %v, got %s", nil, err)
	}
	if cache.Fresh() { // no content
		t.Errorf("expected fresh to be %t, got %t", false, cache.Fresh())
	}
	got := cache.Update(1)
	want := Response{Cached:false,Data:1}
	if got != want {
		t.Errorf("expected contents to be %s, got %s", want, got)
	}
	if !cache.Fresh() { // has content
		t.Errorf("expected fresh to be %t, got %t", true, cache.Fresh())
	}
	got = cache.Contents(true)
	want = Response{Cached:true,Data:1}
	if got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}
