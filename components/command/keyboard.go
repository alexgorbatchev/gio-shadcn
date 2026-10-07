package command

import (
	"slices"
	"strings"

	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
)

func (c *Command) keyboardEvents(gtx layout.Context) []key.Event {
	filters := make([]event.Filter, 0, (len(c.Items)+1)*6)
	for _, tag := range append([]event.Tag{c.searchEditor}, c.itemTags()...) {
		for _, name := range []key.Name{key.NameUpArrow, key.NameDownArrow, key.NameHome, key.NameEnd, key.NameReturn, key.NameEnter} {
			filters = append(filters, key.Filter{Focus: tag, Name: name})
		}
	}
	var keys []key.Event
	for {
		e, ok := gtx.Event(filters...)
		if !ok {
			break
		}
		if e := e.(key.Event); e.State == key.Press {
			keys = append(keys, e)
		}
	}
	return keys
}

func (c *Command) itemTags() []event.Tag {
	tags := make([]event.Tag, 0, len(c.Items))
	for _, item := range c.Items {
		if !item.Disabled {
			tags = append(tags, item.clickable)
		}
	}
	return tags
}

func (c *Command) updateSelection(gtx layout.Context, query string, modal bool, keys []key.Event) {
	eligible := make([]int, 0, len(c.Items))
	itemFocused := false
	for i, item := range c.Items {
		if !item.Disabled && strings.Contains(strings.ToLower(item.Label), query) {
			eligible = append(eligible, i)
			if gtx.Focused(item.clickable) {
				c.activeIndex = i
				itemFocused = true
			}
		}
	}
	pos := slices.Index(eligible, c.activeIndex)
	if query != c.lastQuery || pos < 0 {
		pos = 0
	}
	c.lastQuery = query
	c.activeIndex = -1
	if len(eligible) == 0 {
		return
	}
	c.activeIndex = eligible[pos]
	for _, e := range keys {
		switch e.Name {
		case key.NameDownArrow:
			pos++
		case key.NameUpArrow:
			pos--
		case key.NameHome:
			pos = 0
		case key.NameEnd:
			pos = len(eligible) - 1
		case key.NameReturn, key.NameEnter:
			c.activate(c.activeIndex, modal)
		}
		if c.Loop {
			pos = (pos + len(eligible)) % len(eligible)
		} else {
			pos = max(0, min(len(eligible)-1, pos))
		}
		c.activeIndex = eligible[pos]
		if itemFocused && e.Name != key.NameReturn && e.Name != key.NameEnter {
			gtx.Execute(key.FocusCmd{Tag: c.Items[c.activeIndex].clickable})
		}
	}
}

func (c *Command) activate(index int, modal bool) {
	if index < 0 || index >= len(c.Items) || c.Items[index].Disabled {
		return
	}
	c.activeIndex = index
	if modal {
		c.Open = false
	}
	if c.OnSelectItem != nil {
		c.OnSelectItem(index)
	}
}
