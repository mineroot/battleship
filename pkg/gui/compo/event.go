package compo

type Event string
type Handler func(data ...any)

const (
	Click      Event = "click"
	MouseLDown Event = "mouse_l_down"
	MouseLUp   Event = "mouse_l_up"
	MouseEnter Event = "mouse_enter"
	MouseMove  Event = "mouse_move"
	MouseLeave Event = "mouse_leave"
	DragStart  Event = "drag_start"
	Dragging   Event = "dragging"
	DragEnd    Event = "drag_end"
	FocusIn    Event = "focus_in"
	FocusOut   Event = "focus_our"
)
