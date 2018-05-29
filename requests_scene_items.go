package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// GetSceneItemPropertiesRequest : Gets the scene specific properties of the specified source item.
// Since obs-websocket version: 4.3.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsceneitemproperties
type GetSceneItemPropertiesRequest struct {
	// the name of the scene that the source item belongs to.
	// Defaults to the current scene.
	// Required: No.
	Scene string `json:"scene"`
	// The name of the source.
	// Required: Yes.
	ItemID string `json:"item.id"`
	// The name of the source.
	// Required: Yes.
	ItemName string `json:"item.name"`
	_request
}

// NewGetSceneItemPropertiesRequest returns a new GetSceneItemPropertiesRequest.
func NewGetSceneItemPropertiesRequest(
	scene string,
	itemID string,
	itemName string,
) GetSceneItemPropertiesRequest {
	return GetSceneItemPropertiesRequest{
		scene,
		itemID,
		itemName,
		_request{
			MessageID:   getMessageID(),
			RequestType: "GetSceneItemProperties",
		},
	}

}

// ID returns the request's message ID.
func (r GetSceneItemPropertiesRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r GetSceneItemPropertiesRequest) Type() string { return r.RequestType }

// GetSceneItemPropertiesResponse : Response for GetSceneItemPropertiesRequest.
// Since obs-websocket version: 4.3.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsceneitemproperties
type GetSceneItemPropertiesResponse struct {
	// The name of the scene.
	// Required: Yes.
	Scene string `mapstructure:"scene"`
	// The name of the source.
	// Required: Yes.
	ItemName string `mapstructure:"item.name"`
	// The id of the scene item.
	// Required: Yes.
	ItemID string `mapstructure:"item.id"`
	// The x position of the source from the left.
	// Required: Yes.
	ItemPositionX int `mapstructure:"item.position.x"`
	// The y position of the source from the top.
	// Required: Yes.
	ItemPositionY int `mapstructure:"item.position.y"`
	// The point on the source that the item is manipulated from.
	// Required: Yes.
	ItemPositionAlignment int `mapstructure:"item.position.alignment"`
	// The clockwise rotation of the item in degrees around the point of alignment.
	// Required: Yes.
	ItemRotation float64 `mapstructure:"item.rotation"`
	// The x-scale factor of the source.
	// Required: Yes.
	ItemScaleX float64 `mapstructure:"item.scale.x"`
	// The y-scale factor of the source.
	// Required: Yes.
	ItemScaleY float64 `mapstructure:"item.scale.y"`
	// The number of pixels cropped off the top of the source before scaling.
	// Required: Yes.
	ItemCropTop int `mapstructure:"item.crop.top"`
	// The number of pixels cropped off the right of the source before scaling.
	// Required: Yes.
	ItemCropRight int `mapstructure:"item.crop.right"`
	// The number of pixels cropped off the bottom of the source before scaling.
	// Required: Yes.
	ItemCropBottom int `mapstructure:"item.crop.bottom"`
	// The number of pixels cropped off the left of the source before scaling.
	// Required: Yes.
	ItemCropLeft int `mapstructure:"item.crop.left"`
	// If the source is visible.
	// Required: Yes.
	ItemVisible bool `mapstructure:"item.visible"`
	// If the source is locked.
	// Required: Yes.
	ItemLocked bool `mapstructure:"item.locked"`
	// Type of bounding box.
	// Required: Yes.
	ItemBoundsType string `mapstructure:"item.bounds.type"`
	// Alignment of the bounding box.
	// Required: Yes.
	ItemBoundsAlignment int `mapstructure:"item.bounds.alignment"`
	// Width of the bounding box.
	// Required: Yes.
	ItemBoundsX float64 `mapstructure:"item.bounds.x"`
	// Height of the bounding box.
	// Required: Yes.
	ItemBoundsY float64 `mapstructure:"item.bounds.y"`
	_response   `mapstructure:",squash"`
}

// ID returns the response's message ID.
func (r GetSceneItemPropertiesResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r GetSceneItemPropertiesResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r GetSceneItemPropertiesResponse) Err() string { return r.Error }

// SetSceneItemPropertiesRequest : Sets the scene specific properties of a source
// Unspecified properties will remain unchanged.
// Since obs-websocket version: 4.3.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemproperties
type SetSceneItemPropertiesRequest struct {
	// the name of the scene that the source item belongs to.
	// Defaults to the current scene.
	// Required: No.
	Scene string `json:"scene"`
	// The name of the item.
	// Required: Yes.
	ItemName string `json:"item.name"`
	// The id of the item.
	// Required: Yes.
	ItemID int `json:"item.id"`
	// The new x position of the item.
	// Required: Yes.
	ItemPositionX int `json:"item.position.x"`
	// The new y position of the item.
	// Required: Yes.
	ItemPositionY int `json:"item.position.y"`
	// The new alignment of the item.
	// Required: Yes.
	ItemPositionAlignment int `json:"item.position.alignment"`
	// The new clockwise rotation of the item in degrees.
	// Required: Yes.
	ItemRotation float64 `json:"item.rotation"`
	// The new x scale of the item.
	// Required: Yes.
	ItemScaleX float64 `json:"item.scale.x"`
	// The new y scale of the item.
	// Required: Yes.
	ItemScaleY float64 `json:"item.scale.y"`
	// The new amount of pixels cropped off the top of the source before scaling.
	// Required: Yes.
	ItemCropTop int `json:"item.crop.top"`
	// The new amount of pixels cropped off the bottom of the source before scaling.
	// Required: Yes.
	ItemCropBottom int `json:"item.crop.bottom"`
	// The new amount of pixels cropped off the left of the source before scaling.
	// Required: Yes.
	ItemCropLeft int `json:"item.crop.left"`
	// The new amount of pixels cropped off the right of the source before scaling.
	// Required: Yes.
	ItemCropRight int `json:"item.crop.right"`
	// The new visibility of the item.
	// 'true' shows source, 'false' hides source.
	// Required: Yes.
	ItemVisible bool `json:"item.visible"`
	// The new locked of the item.
	// 'true' is locked, 'false' is unlocked.
	// Required: Yes.
	ItemLocked bool `json:"item.locked"`
	// The new bounds type of the item.
	// Required: Yes.
	ItemBoundsType string `json:"item.bounds.type"`
	// The new alignment of the bounding box.
	// (0-2, 4-6, 8-10).
	// Required: Yes.
	ItemBoundsAlignment int `json:"item.bounds.alignment"`
	// The new width of the bounding box.
	// Required: Yes.
	ItemBoundsX float64 `json:"item.bounds.x"`
	// The new height of the bounding box.
	// Required: Yes.
	ItemBoundsY float64 `json:"item.bounds.y"`
	_request
}

// NewSetSceneItemPropertiesRequest returns a new SetSceneItemPropertiesRequest.
func NewSetSceneItemPropertiesRequest(
	scene string,
	itemName string,
	itemID int,
	itemPositionX int,
	itemPositionY int,
	itemPositionAlignment int,
	itemRotation float64,
	itemScaleX float64,
	itemScaleY float64,
	itemCropTop int,
	itemCropBottom int,
	itemCropLeft int,
	itemCropRight int,
	itemVisible bool,
	itemLocked bool,
	itemBoundsType string,
	itemBoundsAlignment int,
	itemBoundsX float64,
	itemBoundsY float64,
) SetSceneItemPropertiesRequest {
	return SetSceneItemPropertiesRequest{
		scene,
		itemName,
		itemID,
		itemPositionX,
		itemPositionY,
		itemPositionAlignment,
		itemRotation,
		itemScaleX,
		itemScaleY,
		itemCropTop,
		itemCropBottom,
		itemCropLeft,
		itemCropRight,
		itemVisible,
		itemLocked,
		itemBoundsType,
		itemBoundsAlignment,
		itemBoundsX,
		itemBoundsY,
		_request{
			MessageID:   getMessageID(),
			RequestType: "SetSceneItemProperties",
		},
	}

}

// ID returns the request's message ID.
func (r SetSceneItemPropertiesRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetSceneItemPropertiesRequest) Type() string { return r.RequestType }

// SetSceneItemPropertiesResponse : Response for SetSceneItemPropertiesRequest.
// Since obs-websocket version: 4.3.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemproperties
type SetSceneItemPropertiesResponse _response

// ID returns the response's message ID.
func (r SetSceneItemPropertiesResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetSceneItemPropertiesResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetSceneItemPropertiesResponse) Err() string { return r.Error }

// ResetSceneItemRequest : Reset a scene item.
// Since obs-websocket version: 4.2.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#resetsceneitem
type ResetSceneItemRequest struct {
	// Name of the scene the source belogns to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// Name of the source item.
	// Required: Yes.
	Item string `json:"item"`
	_request
}

// NewResetSceneItemRequest returns a new ResetSceneItemRequest.
func NewResetSceneItemRequest(
	sceneName string,
	item string,
) ResetSceneItemRequest {
	return ResetSceneItemRequest{
		sceneName,
		item,
		_request{
			MessageID:   getMessageID(),
			RequestType: "ResetSceneItem",
		},
	}

}

// ID returns the request's message ID.
func (r ResetSceneItemRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r ResetSceneItemRequest) Type() string { return r.RequestType }

// ResetSceneItemResponse : Response for ResetSceneItemRequest.
// Since obs-websocket version: 4.2.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#resetsceneitem
type ResetSceneItemResponse _response

// ID returns the response's message ID.
func (r ResetSceneItemResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r ResetSceneItemResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r ResetSceneItemResponse) Err() string { return r.Error }

// SetSceneItemRenderRequest : Show or hide a specified source item in a specified scene.
// Since obs-websocket version: 0.3.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemrender
type SetSceneItemRenderRequest struct {
	// Scene item name in the specified scene.
	// Required: Yes.
	Source string `json:"source"`
	// true = shown ; false = hidden.
	// Required: Yes.
	Render bool `json:"render"`
	// Name of the scene where the source resides.
	// Defaults to the currently active scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	_request
}

// NewSetSceneItemRenderRequest returns a new SetSceneItemRenderRequest.
func NewSetSceneItemRenderRequest(
	source string,
	render bool,
	sceneName string,
) SetSceneItemRenderRequest {
	return SetSceneItemRenderRequest{
		source,
		render,
		sceneName,
		_request{
			MessageID:   getMessageID(),
			RequestType: "SetSceneItemRender",
		},
	}

}

// ID returns the request's message ID.
func (r SetSceneItemRenderRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetSceneItemRenderRequest) Type() string { return r.RequestType }

// SetSceneItemRenderResponse : Response for SetSceneItemRenderRequest.
// Since obs-websocket version: 0.3.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemrender
type SetSceneItemRenderResponse _response

// ID returns the response's message ID.
func (r SetSceneItemRenderResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetSceneItemRenderResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetSceneItemRenderResponse) Err() string { return r.Error }

// SetSceneItemPositionRequest : Sets the coordinates of a specified source item.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemposition
type SetSceneItemPositionRequest struct {
	// The name of the scene that the source item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// The name of the source item.
	// Required: Yes.
	Item string `json:"item"`
	// X coordinate.
	// Required: Yes.
	X float64 `json:"x"`
	// Y coordinate.
	// Required: Yes.
	Y float64 `json:"y"`
	_request
}

// NewSetSceneItemPositionRequest returns a new SetSceneItemPositionRequest.
func NewSetSceneItemPositionRequest(
	sceneName string,
	item string,
	x float64,
	y float64,
) SetSceneItemPositionRequest {
	return SetSceneItemPositionRequest{
		sceneName,
		item,
		x,
		y,
		_request{
			MessageID:   getMessageID(),
			RequestType: "SetSceneItemPosition",
		},
	}

}

// ID returns the request's message ID.
func (r SetSceneItemPositionRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetSceneItemPositionRequest) Type() string { return r.RequestType }

// SetSceneItemPositionResponse : Response for SetSceneItemPositionRequest.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemposition
type SetSceneItemPositionResponse _response

// ID returns the response's message ID.
func (r SetSceneItemPositionResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetSceneItemPositionResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetSceneItemPositionResponse) Err() string { return r.Error }

// SetSceneItemTransformRequest : Set the transform of the specified source item.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemtransform
type SetSceneItemTransformRequest struct {
	// The name of the scene that the source item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// The name of the source item.
	// Required: Yes.
	Item string `json:"item"`
	// Width scale factor.
	// Required: Yes.
	XScale float64 `json:"x-scale"`
	// Height scale factor.
	// Required: Yes.
	YScale float64 `json:"y-scale"`
	// Source item rotation (in degrees).
	// Required: Yes.
	Rotation float64 `json:"rotation"`
	_request
}

// NewSetSceneItemTransformRequest returns a new SetSceneItemTransformRequest.
func NewSetSceneItemTransformRequest(
	sceneName string,
	item string,
	xScale float64,
	yScale float64,
	rotation float64,
) SetSceneItemTransformRequest {
	return SetSceneItemTransformRequest{
		sceneName,
		item,
		xScale,
		yScale,
		rotation,
		_request{
			MessageID:   getMessageID(),
			RequestType: "SetSceneItemTransform",
		},
	}

}

// ID returns the request's message ID.
func (r SetSceneItemTransformRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetSceneItemTransformRequest) Type() string { return r.RequestType }

// SetSceneItemTransformResponse : Response for SetSceneItemTransformRequest.
// Since obs-websocket version: 4.0.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemtransform
type SetSceneItemTransformResponse _response

// ID returns the response's message ID.
func (r SetSceneItemTransformResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetSceneItemTransformResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetSceneItemTransformResponse) Err() string { return r.Error }

// SetSceneItemCropRequest : Sets the crop coordinates of the specified source item.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemcrop
type SetSceneItemCropRequest struct {
	// the name of the scene that the source item belongs to.
	// Defaults to the current scene.
	// Required: No.
	SceneName string `json:"scene-name"`
	// The name of the source.
	// Required: Yes.
	Item string `json:"item"`
	// Pixel position of the top of the source item.
	// Required: Yes.
	Top int `json:"top"`
	// Pixel position of the bottom of the source item.
	// Required: Yes.
	Bottom int `json:"bottom"`
	// Pixel position of the left of the source item.
	// Required: Yes.
	Left int `json:"left"`
	// Pixel position of the right of the source item.
	// Required: Yes.
	Right int `json:"right"`
	_request
}

// NewSetSceneItemCropRequest returns a new SetSceneItemCropRequest.
func NewSetSceneItemCropRequest(
	sceneName string,
	item string,
	top int,
	bottom int,
	left int,
	right int,
) SetSceneItemCropRequest {
	return SetSceneItemCropRequest{
		sceneName,
		item,
		top,
		bottom,
		left,
		right,
		_request{
			MessageID:   getMessageID(),
			RequestType: "SetSceneItemCrop",
		},
	}

}

// ID returns the request's message ID.
func (r SetSceneItemCropRequest) ID() string { return r.MessageID }

// Type returns the request's message type.
func (r SetSceneItemCropRequest) Type() string { return r.RequestType }

// SetSceneItemCropResponse : Response for SetSceneItemCropRequest.
// Since obs-websocket version: 4.1.0.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemcrop
type SetSceneItemCropResponse _response

// ID returns the response's message ID.
func (r SetSceneItemCropResponse) ID() string { return r.MessageID }

// Stat returns the response's status.
func (r SetSceneItemCropResponse) Stat() string { return r.Status }

// Err returns the response's error.
func (r SetSceneItemCropResponse) Err() string { return r.Error }
